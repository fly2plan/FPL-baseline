package chain

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"math/big"
	"service/pkg/solidity"
)

type Chain struct {
	vFPL   *solidity.FPLVerifier
	vACK   *solidity.ACKVerifier
	Shield *solidity.Shield
	auth   *bind.TransactOpts
}

func (ch *Chain) Initialise() error {
	const gasLimit uint64 = 80000029

	// setup simulated backend
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	auth.GasPrice=big.NewInt(875000000)
	genesis := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: big.NewInt(100000000000000000)},
	}
	backend := backends.NewSimulatedBackend(genesis, gasLimit)
	ch.auth = auth

	// deploy verifier contracts
	vFPLaddress, _, vFPL, err := solidity.DeployFPLVerifier(auth, backend)
	if err != nil {
		return err
	}
	ch.vFPL = vFPL
	vACKaddress, _, vACK, err := solidity.DeployACKVerifier(auth, backend)
	if err != nil {
		return err
	}
	ch.vACK = vACK
	// deploy shield contract
	_, _, shield, err := solidity.DeployShield(auth, backend, vFPLaddress, vACKaddress, big.NewInt(0))
	if err != nil {
		return err
	}
	ch.Shield = shield
	backend.Commit()
	return nil
}

func (ch *Chain) VerifyAndPush(t string, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int, inputs []*big.Int) (*types.Transaction, error) {
	var pCommitment []byte
	for _, input := range inputs {
		pCommitment = append(pCommitment, input.Bytes()...)
	}
	h := sha3.New256()
	h.Write(pCommitment)
	var hash [32]byte
	copy(hash[:], h.Sum(nil))
	switch t {
	case "fpl":
		var input [6]*big.Int
		copy(input[:], inputs)
		txn, err := ch.Shield.VerifyAndPushFPL(ch.auth, a, b, c, input, hash)
		if err != nil {
			return nil, err
		}
		return txn, nil
	case "ack":
		var input [12]*big.Int
		copy(input[:], inputs)
		txn, err := ch.Shield.VerifyAndPushACK(ch.auth, a, b, c, input, hash)
		if err != nil {
			return nil, err
		}
		return txn, nil
	default:
		return nil, errors.New("unknown verifier type")
	}
}



