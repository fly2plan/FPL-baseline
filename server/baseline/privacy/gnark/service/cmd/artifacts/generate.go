package main

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"os"
	"service/pkg/prove"
)

func setupFPLCircuit() error {
	var circuit prove.FPLCircuit
	r1cs, err := frontend.Compile(ecc.BN254, backend.GROTH16, &circuit)
	if err != nil {
		return err
	}

	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		return err
	}

	r1csf, err := os.Create("../keys/fpl_r1cs")
	if err != nil {
		return err
	}
	r1cs.WriteTo(r1csf)
	r1csf.Close()

	pkf, err := os.Create("../keys/fpl_pk")
	if err != nil {
		return err
	}
	pk.WriteTo(pkf)
	pkf.Close()

	vkf, err := os.Create("../keys/fpl_vk")
	if err != nil {
		return err
	}
	vk.WriteTo(vkf)
	vkf.Close()

	return nil
}

func setupACKCircuit() error {
	var circuit prove.ACKCircuit
	r1cs, err := frontend.Compile(ecc.BN254, backend.GROTH16, &circuit)
	if err != nil {
		return err
	}

	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		return err
	}

	r1csf, err := os.Create("../keys/ack_r1cs")
	if err != nil {
		return err
	}
	r1cs.WriteTo(r1csf)
	r1csf.Close()

	pkf, err := os.Create("../keys/ack_pk")
	if err != nil {
		return err
	}
	pk.WriteTo(pkf)
	pkf.Close()

	vkf, err := os.Create("../keys/ack_vk")
	if err != nil {
		return err
	}
	vk.WriteTo(vkf)
	vkf.Close()

	return nil
}
