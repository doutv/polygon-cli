package loadtest

import (
	"context"
	"math/big"
	"time"

	erc4337loadtest "github.com/0xPolygon/polygon-cli/cmd/loadtest/erc4337"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	erc4337Usage          string
	erc4337LoadTestParams erc4337params
)

type erc4337params struct {
	UopBatchSize          *uint32
	EntryPoint, AccountFactory, Config, Helper, TokenReceiver, WebAuthnAndECDSAValidator, PayableAccount *string
}

var erc4337LoadTestCmd = &cobra.Command{
	Use:   "erc4337",
	Short: "Run ERC4337-like load test against an Eth/EVm style JSON-RPC endpoint.",
	Long:  erc4337Usage,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputLoadTestParams.Modes = &[]string{"4337"}

		err := runLoadTest(cmd.Context())
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	params := new(erc4337params)
	params.UopBatchSize = erc4337LoadTestCmd.Flags().Uint32("uop-batch-size", 1, "The batch size of user operations")
	params.EntryPoint = erc4337LoadTestCmd.Flags().String("entry-point", "", "The address of a pre-deployed EntryPoint contract")
	params.AccountFactory = erc4337LoadTestCmd.Flags().String("account-factory", "", "The address of a pre-deployed AccountFactory contract")
	params.Config = erc4337LoadTestCmd.Flags().String("config", "", "The address of a pre-deployed Config contract")
	params.Helper = erc4337LoadTestCmd.Flags().String("helper", "", "The address of a pre-deployed Helper contract")
	params.TokenReceiver = erc4337LoadTestCmd.Flags().String("token-receiver", "", "The address of a pre-deployed TokenReceiver contract")
	params.WebAuthnAndECDSAValidator = erc4337LoadTestCmd.Flags().String("validator", "", "The address of a pre-deployed WebAuthnAndECDSAValidator contract")
	params.PayableAccount = erc4337LoadTestCmd.Flags().String("payable-account", "", "The address of a pre-deployed PayableAccount contract")
	erc4337LoadTestParams = *params
}

func initERC4337Loadtest(ctx context.Context, c *ethclient.Client, tops *bind.TransactOpts, cops *bind.CallOpts, erc4337Addresses erc4337loadtest.ERC4337Addresses, fromAddress common.Address, uopBatchSize uint32) (erc4337Config erc4337loadtest.ERC4337Config, err error) {
	log.Debug().Msg("Deploying ERC4337 contracts...")
	erc4337Config, err = erc4337loadtest.DeployContracts(ctx, c, tops, cops, erc4337Addresses, fromAddress)
	erc4337Config.UopBatchSize = uopBatchSize
	if err != nil {
		panic(err)
	}
	log.Debug().Interface("addresses", erc4337Config.GetAddresses()).Msg("ERC4337 contracts deployed")

	// Deposit 1 ETH to EntryPoint
	tops.Value = big.NewInt(1000000000000000000)
	if _, err = erc4337Config.EntryPoint.Contract.DepositTo(tops, fromAddress); err != nil {
		panic(err)
	}

	// Create AA for sender, send init UOP
	privateKey := inputLoadTestParams.ECDSAPrivateKey
	if err = erc4337loadtest.SendInitUop(c, ctx, tops, cops, privateKey, &erc4337Config); err != nil {
		panic(err)
	}
	return
}

func runERC4337Loadtest(ctx context.Context, c *ethclient.Client, nonce uint64, config erc4337loadtest.ERC4337Config) (t1 time.Time, t2 time.Time, err error) {
	ltp := inputLoadTestParams
	chainID := new(big.Int).SetUint64(*ltp.ChainID)
	privateKey := ltp.ECDSAPrivateKey

	tops, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Error().Err(err).Msg("Unable create transaction signer")
		return
	}
	tops.Nonce = new(big.Int).SetUint64(nonce)
	tops = configureTransactOpts(tops)

	t1 = time.Now()
	defer func() { t2 = time.Now() }()

	// send user operation
	cops := new(bind.CallOpts)
	if err = erc4337loadtest.SendUops(c, ctx, tops, cops, privateKey, &config); err != nil {
		panic(err)
	}
	return
}
