package goton

import (
	"github.com/move-ton/ever-client-go/domain"
	clientgw "github.com/move-ton/ever-client-go/gateway/client"
	"github.com/move-ton/ever-client-go/usecase/abi"
	"github.com/move-ton/ever-client-go/usecase/boc"
	"github.com/move-ton/ever-client-go/usecase/crypto"
	"github.com/move-ton/ever-client-go/usecase/debot"
	"github.com/move-ton/ever-client-go/usecase/net"
	"github.com/move-ton/ever-client-go/usecase/processing"
	"github.com/move-ton/ever-client-go/usecase/proofs"
	"github.com/move-ton/ever-client-go/usecase/tvm"
	"github.com/move-ton/ever-client-go/usecase/utils"
)

// Ton ...
type Ever struct {
	Abi        domain.AbiUseCase
	Boc        domain.BocUseCase
	Client     domain.ClientGateway
	Crypto     domain.CryptoUseCase
	Debot      domain.DebotUseCase
	Net        domain.NetUseCase
	Processing domain.ProcessingUseCase
	Proofs     domain.ProofsUseCase
	Tvm        domain.TvmUseCase
	Utils      domain.UtilsUseCase
}

// NewEverWithConfig ...
func NewEverWithConfig(config domain.Config) (*Ever, error) {
	client, err := clientgw.NewClientGateway(config)
	if err != nil {
		return nil, err
	}
	ever := &Ever{
		Abi:        abi.NewAbi(config, client),
		Boc:        boc.NewBoc(config, client),
		Client:     client,
		Crypto:     crypto.NewCrypto(config, client),
		Debot:      debot.NewDebot(config, client),
		Net:        net.NewNet(config, client),
		Proofs:     proofs.NewProofs(config, client),
		Processing: processing.NewProcessing(config, client),
		Tvm:        tvm.NewTvm(config, client),
		Utils:      utils.NewUtils(config, client),
	}
	return ever, nil
}

// NewTon ...
func NewTon(address string, endPoints []string) (*Ever, error) {
	conf := domain.NewDefaultConfig(address, endPoints)
	return NewEverWithConfig(conf)
}
