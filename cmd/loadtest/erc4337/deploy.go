package erc4337loadtest

import (
	"context"
	"fmt"

	"github.com/0xPolygon/polygon-cli/bindings/4337/accountfactory"
	"github.com/0xPolygon/polygon-cli/bindings/4337/config"
	"github.com/0xPolygon/polygon-cli/bindings/4337/entryPoint/core/entrypoint"
	"github.com/0xPolygon/polygon-cli/bindings/4337/modules/fallbackhandlers/tokenreceiver"
	"github.com/0xPolygon/polygon-cli/bindings/4337/modules/validators/webauthnandecdsavalidator"
	"github.com/0xPolygon/polygon-cli/bindings/4337/payableaccount"
	"github.com/0xPolygon/polygon-cli/bindings/4337/test/helper"
	"github.com/0xPolygon/polygon-cli/bindings/4337/test/mock/mockrecoverymodule"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

type (
	// ERC4337Config represents the whole ERC4337 configuration (contracts and addresses)
	ERC4337Config struct {
		UopBatchSize       uint32
		EntryPoint         ContractConfig[entrypoint.EntryPoint]
		AccountFactory     ContractConfig[accountfactory.AccountFactory]
		PayableAccount     ContractConfig[payableaccount.PayableAccount]
		WebAuthnValidator  ContractConfig[webauthnandecdsavalidator.WebAuthnAndECDSAValidator]
		Config             ContractConfig[config.Config]
		Helper             ContractConfig[helper.Helper]
		TokenReceiver      ContractConfig[tokenreceiver.TokenReceiver]
		MockRecoveryModule ContractConfig[mockrecoverymodule.MockRecoveryModule]
		Sender             common.Address // computed counterfactual address
	}

	// ERC4337Addresses is a subset of ERC4337Config. It represents the addresses of the whole
	// ERC4337 configuration.
	ERC4337Addresses struct {
		EntryPoint, AccountFactory, PayableAccount, WebAuthnValidator, Config, Helper, TokenReceiver, MockRecoveryModule common.Address
	}

	// ContractConfig represents a contract and its address.
	ContractConfig[T Contract] struct {
		Address  common.Address
		Contract *T
	}

	// Contract represents an ERC4337 contract
	Contract interface {
		entrypoint.EntryPoint | accountfactory.AccountFactory | payableaccount.PayableAccount | webauthnandecdsavalidator.WebAuthnAndECDSAValidator | config.Config | helper.Helper | tokenreceiver.TokenReceiver | mockrecoverymodule.MockRecoveryModule
	}
)

func DeployContracts(ctx context.Context, client *ethclient.Client, tops *bind.TransactOpts, cops *bind.CallOpts, knownAddresses ERC4337Addresses, fromAddress common.Address) (cfg ERC4337Config, err error) {
	log.Debug().Msg("Instantiating EntryPoint")
	cfg.EntryPoint.Address = knownAddresses.EntryPoint
	cfg.EntryPoint.Contract, err = entrypoint.NewEntryPoint(knownAddresses.EntryPoint, client)
	if err != nil {
		return cfg, fmt.Errorf("failed to instantiate EntryPoint: %w", err)
	}

	log.Debug().Msg("Instantiating Helper")
	cfg.Helper.Address = knownAddresses.Helper
	cfg.Helper.Contract, err = helper.NewHelper(knownAddresses.Helper, client)
	if err != nil {
		return cfg, fmt.Errorf("failed to instantiate Helper: %w", err)
	}

	log.Debug().Msg("Instantiating TokenReceiver")
	cfg.TokenReceiver.Address = knownAddresses.TokenReceiver
	cfg.TokenReceiver.Contract, err = tokenreceiver.NewTokenReceiver(knownAddresses.TokenReceiver, client)
	if err != nil {
		return cfg, fmt.Errorf("failed to instantiate TokenReceiver: %w", err)
	}

	log.Debug().Msg("Instantiating Config")
	cfg.Config.Address = knownAddresses.Config
	cfg.Config.Contract, err = config.NewConfig(knownAddresses.Config, client)
	if err != nil {
		return cfg, fmt.Errorf("failed to instantiate Config: %w", err)
	}

	log.Debug().Msg("Instantiating WebAuthnAndECDSAValidator")
	cfg.WebAuthnValidator.Address = knownAddresses.WebAuthnValidator
	cfg.WebAuthnValidator.Contract, err = webauthnandecdsavalidator.NewWebAuthnAndECDSAValidator(knownAddresses.WebAuthnValidator, client)
	if err != nil {
		return cfg, fmt.Errorf("failed to instantiate WebAuthnAndECDSAValidator: %w", err)
	}

	log.Debug().Msg("Instantiating PayableAccount")
	cfg.PayableAccount.Address = knownAddresses.PayableAccount
	cfg.PayableAccount.Contract, err = payableaccount.NewPayableAccount(knownAddresses.PayableAccount, client)
	if err != nil {
		return cfg, fmt.Errorf("failed to instantiate PayableAccount: %w", err)
	}

	log.Debug().Msg("Instantiating AccountFactory")
	cfg.AccountFactory.Address = knownAddresses.AccountFactory
	cfg.AccountFactory.Contract, err = accountfactory.NewAccountFactory(knownAddresses.AccountFactory, client)
	if err != nil {
		return cfg, fmt.Errorf("failed to instantiate AccountFactory: %w", err)
	}

	log.Debug().Msg("Instantiating MockRecoveryModule")
	cfg.MockRecoveryModule.Address = knownAddresses.MockRecoveryModule
	cfg.MockRecoveryModule.Contract, err = mockrecoverymodule.NewMockRecoveryModule(knownAddresses.MockRecoveryModule, client)
	if err != nil {
		return cfg, fmt.Errorf("failed to instantiate MockRecoveryModule: %w", err)
	}

	// Calculate the sender address (counterfactual address)
	sender, err := cfg.AccountFactory.Contract.ComputeAddress(
		cops,
		cfg.PayableAccount.Address,
		salt,
	)
	if err != nil {
		return cfg, fmt.Errorf("failed to compute sender address: %w", err)
	}
	cfg.Sender = sender
	log.Info().Msgf("Sender address: %s", cfg.Sender.Hex())

	return cfg, nil
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
		TokenReceiver:     c.TokenReceiver.Address,
		MockRecoveryModule: c.MockRecoveryModule.Address,
	}
}
