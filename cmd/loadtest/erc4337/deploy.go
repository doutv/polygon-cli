package erc4337loadtest

import (
	"context"
	"fmt"
	"math/big"
	"reflect"

	"github.com/0xPolygon/polygon-cli/bindings/4337/accountfactory"
	"github.com/0xPolygon/polygon-cli/bindings/4337/config"
	"github.com/0xPolygon/polygon-cli/bindings/4337/entryPoint/core/entrypoint"
	"github.com/0xPolygon/polygon-cli/bindings/4337/modules/fallbackhandlers/tokenreceiver"
	"github.com/0xPolygon/polygon-cli/bindings/4337/modules/validators/webauthnandecdsavalidator"
	"github.com/0xPolygon/polygon-cli/bindings/4337/payableaccount"
	"github.com/0xPolygon/polygon-cli/bindings/4337/test/helper"
	"github.com/0xPolygon/polygon-cli/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

type (
	// ERC4337Config represents the whole ERC4337 configuration (contracts and addresses)
	ERC4337Config struct {
		UopBatchSize      uint32
		EntryPoint        ContractConfig[entrypoint.EntryPoint]
		AccountFactory    ContractConfig[accountfactory.AccountFactory]
		PayableAccount    ContractConfig[payableaccount.PayableAccount]
		WebAuthnValidator ContractConfig[webauthnandecdsavalidator.WebAuthnAndECDSAValidator]
		Config            ContractConfig[config.Config]
		Helper            ContractConfig[helper.Helper]
		TokenReceiver     ContractConfig[tokenreceiver.TokenReceiver]
	}

	// ERC4337Addresses is a subset of ERC4337Config. It represents the addresses of the whole
	// ERC4337 configuration.
	ERC4337Addresses struct {
		EntryPoint, AccountFactory, PayableAccount, WebAuthnValidator, Config, Helper, TokenReceiver common.Address
	}

	// ContractConfig represents a contract and its address.
	ContractConfig[T Contract] struct {
		Address  common.Address
		Contract *T
	}

	// Contract represents an ERC4337 contract
	Contract interface {
		entrypoint.EntryPoint | accountfactory.AccountFactory | payableaccount.PayableAccount | webauthnandecdsavalidator.WebAuthnAndECDSAValidator | config.Config | helper.Helper | tokenreceiver.TokenReceiver
	}
)

func DeployContracts(ctx context.Context, client *ethclient.Client, tops *bind.TransactOpts, cops *bind.CallOpts, knownAddresses ERC4337Addresses, fromAddress common.Address) (cfg ERC4337Config, err error) {
	log.Debug().Msg("Deploying EntryPoint")
	cfg.EntryPoint.Address, cfg.EntryPoint.Contract, err = deployOrInstantiateContract(
		ctx, client, tops, cops,
		knownAddresses.EntryPoint,
		entrypoint.DeployEntryPoint,
		entrypoint.NewEntryPoint,
		func(contract *entrypoint.EntryPoint) (err error) {
			_, err = contract.BalanceOf(cops, fromAddress)
			return
		},
	)
	if err != nil {
		panic(err)
	}

	log.Debug().Msg("Deploying Helper")
	cfg.Helper.Address, cfg.Helper.Contract, err = deployOrInstantiateContract(
		ctx, client, tops, cops,
		knownAddresses.Helper,
		helper.DeployHelper,
		helper.NewHelper,
		func(contract *helper.Helper) (err error) {
			_, err = contract.GetBlocktimeStamp(cops)
			return
		},
	)
	if err != nil {
		panic(err)
	}

	log.Debug().Msg("Deploying TokenReceiver")
	cfg.TokenReceiver.Address, cfg.TokenReceiver.Contract, err = deployOrInstantiateContract(
		ctx, client, tops, cops,
		knownAddresses.TokenReceiver,
		tokenreceiver.DeployTokenReceiver,
		tokenreceiver.NewTokenReceiver,
		func(contract *tokenreceiver.TokenReceiver) (err error) {
			_, err = contract.IsModuleType(cops, big.NewInt(0))
			return
		},
	)
	if err != nil {
		panic(err)
	}

	log.Debug().Msg("Deploying Config")
	cfg.Config.Address, cfg.Config.Contract, err = deployOrInstantiateContract(
		ctx, client, tops, cops,
		knownAddresses.Config,
		func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *config.Config, error) {
			return config.DeployConfig(tops, client, cfg.TokenReceiver.Address, fromAddress)
		},
		config.NewConfig,
		func(contract *config.Config) (err error) {
			owner, err := contract.Owner(cops)
			if owner != fromAddress {
				return fmt.Errorf("expected owner address %s, got %s", fromAddress, owner)
			}
			return
		},
	)
	if err != nil {
		panic(err)
	}

	log.Debug().Msg("Deploying WebAuthnAndECDSAValidator")
	cfg.WebAuthnValidator.Address, cfg.WebAuthnValidator.Contract, err = deployOrInstantiateContract(
		ctx, client, tops, cops,
		knownAddresses.WebAuthnValidator,
		func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *webauthnandecdsavalidator.WebAuthnAndECDSAValidator, error) {
			return webauthnandecdsavalidator.DeployWebAuthnAndECDSAValidator(tops, client, cfg.Config.Address)
		},
		webauthnandecdsavalidator.NewWebAuthnAndECDSAValidator,
		func(contract *webauthnandecdsavalidator.WebAuthnAndECDSAValidator) (err error) {
			cfgAddress, err := contract.CONFIG(cops)
			if cfgAddress != cfg.Config.Address {
				return fmt.Errorf("expected config address %s, got %s", cfg.Config.Address, cfgAddress)
			}
			return
		},
	)
	if err != nil {
		panic(err)
	}

	log.Debug().Msg("Deploying PayableAccount")
	cfg.PayableAccount.Address, cfg.PayableAccount.Contract, err = deployOrInstantiateContract(
		ctx, client, tops, cops,
		knownAddresses.PayableAccount,
		func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *payableaccount.PayableAccount, error) {
			return payableaccount.DeployPayableAccount(tops, client, cfg.EntryPoint.Address, cfg.Config.Address)
		},
		payableaccount.NewPayableAccount,
		func(contract *payableaccount.PayableAccount) (err error) {
			cfgAddress, err := contract.CONFIG(cops)
			if cfgAddress != cfg.Config.Address {
				return fmt.Errorf("expected config address %s, got %s", cfg.Config.Address, cfgAddress)
			}
			entrypointAddress, err := contract.ENTRYPOINT(cops)
			if entrypointAddress != cfg.EntryPoint.Address {
				return fmt.Errorf("expected entrypoint address %s, got %s", cfg.EntryPoint.Address, entrypointAddress)
			}
			return
		},
	)
	if err != nil {
		panic(err)
	}

	log.Debug().Msg("Deploying AccountFactory")
	cfg.AccountFactory.Address, cfg.AccountFactory.Contract, err = deployOrInstantiateContract(
		ctx, client, tops, cops,
		knownAddresses.AccountFactory,
		func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *accountfactory.AccountFactory, error) {
			return accountfactory.DeployAccountFactory(tops, client, cfg.Config.Address, fromAddress)
		},
		accountfactory.NewAccountFactory,
		func(contract *accountfactory.AccountFactory) (err error) {
			cfgAddress, err := contract.CONFIG(cops)
			if cfgAddress != cfg.Config.Address {
				return fmt.Errorf("expected config address %s, got %s", cfg.Config.Address, cfgAddress)
			}
			owner, err := contract.Owner(cops)
			if owner != fromAddress {
				return fmt.Errorf("expected owner address %s, got %s", fromAddress, owner)
			}
			return
		},
	)
	if err != nil {
		panic(err)
	}

	// Configure contracts
	log.Debug().Msg("Configuring contracts")
	_, err = cfg.Config.Contract.AddSafeSingleton(tops, cfg.PayableAccount.Address)
	if err != nil {
		panic(err)
	}
	_, err = cfg.Config.Contract.AddWhitelistedBundlers(tops, []common.Address{fromAddress})
	if err != nil {
		panic(err)
	}

	return
}

// deployOrInstantiateContract deploys or instantiates a UniswapV3 contract.
// If knownAddress is empty, it deploys the contract; otherwise, it instantiates it.
func deployOrInstantiateContract[T Contract](
	ctx context.Context,
	c *ethclient.Client,
	tops *bind.TransactOpts,
	cops *bind.CallOpts,
	knownAddress common.Address,
	deploy func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *T, error),
	instantiate func(common.Address, bind.ContractBackend) (*T, error),
	call func(*T) error,
) (address common.Address, contract *T, err error) {
	if knownAddress == (common.Address{}) {
		// Deploy the contract if known address is empty.
		address, _, contract, err = deploy(tops, c)
		if err != nil {
			log.Error().Err(err).Msg("Unable to deploy contract")
			return
		}
		reflectedContractName := reflect.TypeOf(contract).Elem().Name()
		log.Debug().Str("name", reflectedContractName).Interface("address", address).Msg("Contract deployed")
	} else {
		// Otherwise, instantiate the contract.
		address = knownAddress
		contract, err = instantiate(address, c)
		if err != nil {
			log.Error().Err(err).Msg("Unable to instantiate contract")
			return
		}
		reflectedContractName := reflect.TypeOf(contract).Elem().Name()
		log.Debug().Str("name", reflectedContractName).Msg("Contract instantiated")
	}

	// Check that the contract can be called.
	err = util.BlockUntilSuccessful(ctx, c, func() error {
		log.Trace().Msg("Contract is not available yet")
		return call(contract)
	})
	return
}

// Return contracts addresses from the ERC4337 configuration.
func (c *ERC4337Config) GetAddresses() ERC4337Addresses {
	return ERC4337Addresses{
		EntryPoint:        c.EntryPoint.Address,
		AccountFactory:    c.AccountFactory.Address,
		PayableAccount:    c.PayableAccount.Address,
		WebAuthnValidator: c.WebAuthnValidator.Address,
		Config:            c.Config.Address,
		Helper:            c.Helper.Address,
		TokenReceiver:   c.TokenReceiver.Address,
	}
}
