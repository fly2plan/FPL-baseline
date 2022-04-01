package main

import (
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"os"
)

func exportFPLVerifier() error {
	err := setupFPLCircuit()
	if err != nil {
		return err
	}
	f, err := os.Create("../contracts/FPLVerifier.sol")
	if err != nil {
		return nil
	}
	vkf, err := os.Open("../keys/fpl_vk")
	if err != nil {
		return nil
	}
	vk := groth16.NewVerifyingKey(ecc.BN254)
	_, _ = vk.ReadFrom(vkf)
	err = vk.ExportSolidity(f)
	if err != nil {
		return err
	}
	vkf.Close()
	fmt.Println("written fpl verifier as .sol")

	return nil
}

func exportACKVerifier() error {
	err := setupACKCircuit()
	if err != nil {
		return err
	}
	f, err := os.Create("../contracts/ACKVerifier.sol")
	if err != nil {
		return nil
	}
	vkf, err := os.Open("../keys/ack_vk")
	if err != nil {
		return nil
	}

	vk := groth16.NewVerifyingKey(ecc.BN254)
	_, _ = vk.ReadFrom(vkf)
	err = vk.ExportSolidity(f)
	if err != nil {
		return err
	}
	vkf.Close()
	fmt.Println("written ack verifier as .sol")

	return nil
}

func main() {
	os.Mkdir("../keys", os.ModePerm)
	err := exportFPLVerifier()
	if err != nil {
		fmt.Println(err)
	}
	err = exportACKVerifier()
	if err != nil {
		fmt.Println(err)
	}
}