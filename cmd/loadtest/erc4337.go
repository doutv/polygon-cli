package loadtest

import (
	"bufio"
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
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
	UopBatchSize                                                                                                         *uint32
	EntryPoint, AccountFactory, Config, Helper, TokenReceiver, WebAuthnAndECDSAValidator, PayableAccount, Pay, TestERC20 *string
	CallDataFile                                                                                                         *string
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
	params.Pay = erc4337LoadTestCmd.Flags().String("pay", "", "The address of a pre-deployed Pay contract")
	params.TestERC20 = erc4337LoadTestCmd.Flags().String("test-erc20", "", "The address of a pre-deployed TestERC20 contract")
	params.CallDataFile = erc4337LoadTestCmd.Flags().String("calldata-file", "", "The file containing the calldata to be used for the user operations")
	erc4337LoadTestParams = *params
}

func initERC4337Loadtest(ctx context.Context, c *ethclient.Client, tops *bind.TransactOpts, cops *bind.CallOpts, erc4337Addresses erc4337loadtest.ERC4337Addresses, fromAddress common.Address, erc4337LoadTestParams erc4337params) (erc4337Config erc4337loadtest.ERC4337Config, err error) {
	log.Debug().Msg("Initializing ERC4337 contracts...")
	erc4337Config, err = erc4337loadtest.DeployContracts(ctx, c, tops, cops, erc4337Addresses, fromAddress)
	erc4337Config.UopBatchSize = *erc4337LoadTestParams.UopBatchSize
	erc4337Config.CallDataList, err = readLastFields(*erc4337LoadTestParams.CallDataFile)
	if err != nil {
		log.Error().Msgf("Failed to read call data file: %v, file:%v", err, *erc4337LoadTestParams.CallDataFile)
		return
	}
	if err != nil {
		panic(err)
	}
	log.Debug().Interface("addresses", erc4337Config.GetAddresses()).Msg("ERC4337 contracts deployed")
	return
}

func readLastFields(filePath string) ([]string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return nil, err
	}
	newFilePath := filepath.Join(cwd, filePath)
	file, err := os.Open(newFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var callDataList []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := strings.Split(line, ",")
		lastField := fields[len(fields)-1]
		callDataList = append(callDataList, lastField)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return callDataList, nil
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

	cops := &bind.CallOpts{
		From: tops.From,
	}

	t1 = time.Now()
	defer func() { t2 = time.Now() }()

	// Generate new user operations each time
	userOps, err := erc4337loadtest.GenerateUops(c, ctx, tops, cops, privateKey, &config, nil, config.UopBatchSize)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate user operations")
		return
	}

	// Send user operation
	if _, err = config.EntryPoint.Contract.HandleOps(tops, userOps, tops.From); err != nil {
		log.Error().Err(err).Msg("Failed to send user operations")
	}
	return
}
