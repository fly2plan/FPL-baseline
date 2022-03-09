// Welcome to the gnark playground!
package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	edwardsbn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	eddsabn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/consensys/gnark-crypto/signature"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/hint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/twistededwards"
	"github.com/consensys/gnark/std/hash/mimc"
	"github.com/consensys/gnark/std/signature/eddsa"
	"math/big"
	"strings"
)

type ProveFPLInput struct {
	FPLForm    `json:"fpl"`
	EHash      string `json:"hash"`
	ESignature string `json:"sig"`
	EPublicKey string `json:"pk"`
}

//func reverseSlice(_ ecc.ID, inputs []*big.Int, results []*big.Int) error {
//	for i, j := 0, len(inputs)-1; i < j; i, j = i+1, j-1 {
//		inputs[i], inputs[j] = inputs[j], inputs[i]
//		results[i] = inputs[i]
//		results[j] = inputs[j]
//	}
//	return nil
//}
//
//var reverseArray2 = hint.NewStaticHint(reverseSlice, 2, 2)
//var reverseArray4 = hint.NewStaticHint(reverseSlice, 4, 4)
//var reverseArray5 = hint.NewStaticHint(reverseSlice, 5, 5)
//var reverseArray6 = hint.NewStaticHint(reverseSlice, 6, 6)
//var reverseArray7 = hint.NewStaticHint(reverseSlice, 7, 7)
//var reverseArray10 = hint.NewStaticHint(reverseSlice, 10, 10)
//var reverseArray12 = hint.NewStaticHint(reverseSlice, 12, 12)
//var reverseArray19 = hint.NewStaticHint(reverseSlice, 19, 19)
//var reverseArray38 = hint.NewStaticHint(reverseSlice, 38, 38)
//var reverseArray49 = hint.NewStaticHint(reverseSlice, 49, 49)
//var reverseArray76 = hint.NewStaticHint(reverseSlice, 76, 76)
//var reverseArray152 = hint.NewStaticHint(reverseSlice, 152, 152)

func generateMultiplier(_ ecc.ID, inputs []*big.Int, results []*big.Int) error {
	multiplier := initialMultiplier(int(inputs[0].Int64()))
	*results[0] = *multiplier
	return nil
}

var pMultiplier = hint.NewStaticHint(generateMultiplier, 1, 1)

//func DigitsToInt(api frontend.API, digits [1]frontend.Variable) frontend.Variable {
//	int := frontend.Variable(0)
//	multiplier, _ := api.NewHint(pMultiplier, []frontend.Variable{ 1 })
//	int = api.Add(int, api.Mul(digits[0], multiplier))
//	api.Println(int)
//	return int
//}

func CharactersToInt2(api frontend.API, characters [2]frontend.Variable) frontend.Variable {
	cAsInt := frontend.Variable(0)
	multiplier, _ := api.NewHint(pMultiplier, []frontend.Variable{2}...)
	for i := 0; i < 2; i++ {
		cAsInt = api.Add(cAsInt, api.Mul(characters[i], multiplier[0]))
		multiplier[0] = api.Div(multiplier[0], 100)
	}
	return cAsInt
}

func CharactersToInt4(api frontend.API, characters [4]frontend.Variable) frontend.Variable {
	cAsInt := frontend.Variable(0)
	multiplier, _ := api.NewHint(pMultiplier, []frontend.Variable{4}...)
	for i := 0; i < 4; i++ {
		cAsInt = api.Add(cAsInt, api.Mul(characters[i], multiplier[0]))
		multiplier[0] = api.Div(multiplier[0], 100)
	}
	return cAsInt
}

func CharactersToInt5(api frontend.API, characters [5]frontend.Variable) frontend.Variable {
	cAsInt := frontend.Variable(0)
	multiplier, _ := api.NewHint(pMultiplier, []frontend.Variable{5}...)
	for i := 0; i < 5; i++ {
		cAsInt = api.Add(cAsInt, api.Mul(characters[i], multiplier[0]))
		multiplier[0] = api.Div(multiplier[0], 100)
	}
	return cAsInt
}

func CharactersToInt6(api frontend.API, characters [6]frontend.Variable) frontend.Variable {
	cAsInt := frontend.Variable(0)
	multiplier, _ := api.NewHint(pMultiplier, []frontend.Variable{6}...)
	for i := 0; i < 6; i++ {
		cAsInt = api.Add(cAsInt, api.Mul(characters[i], multiplier[0]))
		multiplier[0] = api.Div(multiplier[0], 100)
	}
	return cAsInt
}

func CharactersToInt7(api frontend.API, characters [7]frontend.Variable) frontend.Variable {
	cAsInt := frontend.Variable(0)
	multiplier, _ := api.NewHint(pMultiplier, []frontend.Variable{7}...)
	for i := 0; i < 7; i++ {
		cAsInt = api.Add(cAsInt, api.Mul(characters[i], multiplier[0]))
		multiplier[0] = api.Div(multiplier[0], 100)
	}
	return cAsInt
}

func CharactersToInt10(api frontend.API, characters [10]frontend.Variable) frontend.Variable {
	cAsInt := frontend.Variable(0)
	multiplier, _ := api.NewHint(pMultiplier, []frontend.Variable{10}...)
	for i := 0; i < 10; i++ {
		cAsInt = api.Add(cAsInt, api.Mul(characters[i], multiplier[0]))
		multiplier[0] = api.Div(multiplier[0], 100)
	}
	return cAsInt
}

func CharactersToInt12(api frontend.API, characters [12]frontend.Variable) frontend.Variable {
	cAsInt := frontend.Variable(0)
	multiplier, _ := api.NewHint(pMultiplier, []frontend.Variable{12}...)
	for i := 0; i < 12; i++ {
		cAsInt = api.Add(cAsInt, api.Mul(characters[i], multiplier[0]))
		multiplier[0] = api.Div(multiplier[0], 100)
	}
	return cAsInt
}

func CharactersToInt19(api frontend.API, characters [19]frontend.Variable) frontend.Variable {
	cAsInt := frontend.Variable(0)
	multiplier, _ := api.NewHint(pMultiplier, []frontend.Variable{19}...)
	for i := 0; i < 19; i++ {
		cAsInt = api.Add(cAsInt, api.Mul(characters[i], multiplier[0]))
		multiplier[0] = api.Div(multiplier[0], 100)
	}
	return cAsInt
}

func CharactersToInt37(api frontend.API, characters [37]frontend.Variable) frontend.Variable {
	cAsInt := frontend.Variable(0)
	multiplier, _ := api.NewHint(pMultiplier, []frontend.Variable{37}...)
	for i := 0; i < 37; i++ {
		cAsInt = api.Add(cAsInt, api.Mul(characters[i], multiplier[0]))
		multiplier[0] = api.Div(multiplier[0], 100)
	}
	return cAsInt
}

func CharactersToInt49(api frontend.API, characters [49]frontend.Variable) [2]frontend.Variable {
	ints := [2]frontend.Variable{0, 0}
	multiplier1, _ := api.NewHint(pMultiplier, []frontend.Variable{25}...)
	multiplier2, _ := api.NewHint(pMultiplier, []frontend.Variable{24}...)
	multipliers := [2]frontend.Variable{multiplier1[0], multiplier2[0]}
	for i := 0; i < 25; i++ {
		ints[0] = api.Add(ints[0], api.Mul(characters[i], multipliers[0]))
		multipliers[0] = api.Div(multipliers[0], 100)
	}
	for i := 25; i < 49; i++ {
		ints[1] = api.Add(ints[1], api.Mul(characters[i], multipliers[1]))
		multipliers[1] = api.Div(multipliers[1], 100)
	}
	return ints
}

func CharactersToInt74(api frontend.API, characters [74]frontend.Variable) [2]frontend.Variable {
	ints := [2]frontend.Variable{big.NewInt(0), big.NewInt(0)}
	multiplier1, _ := api.NewHint(pMultiplier, []frontend.Variable{37}...)
	multiplier2, _ := api.NewHint(pMultiplier, []frontend.Variable{37}...)
	multipliers := [2]frontend.Variable{multiplier1[0], multiplier2[0]}
	for i := 0; i < 37; i++ {
		ints[0] = api.Add(ints[0], api.Mul(characters[i], multipliers[0]))
		multipliers[0] = api.Div(multipliers[0], 100)
	}
	for i := 37; i < 74; i++ {
		ints[1] = api.Add(ints[1], api.Mul(characters[i], multipliers[1]))
		multipliers[1] = api.Div(multipliers[1], 100)
	}
	return ints
}

func CharactersToInt148(api frontend.API, characters [148]frontend.Variable) [4]frontend.Variable {
	ints := [4]frontend.Variable{big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	multiplier1, _ := api.NewHint(pMultiplier, []frontend.Variable{37}...)
	multiplier2, _ := api.NewHint(pMultiplier, []frontend.Variable{37}...)
	multiplier3, _ := api.NewHint(pMultiplier, []frontend.Variable{37}...)
	multiplier4, _ := api.NewHint(pMultiplier, []frontend.Variable{37}...)
	multipliers := [4]frontend.Variable{multiplier1[0], multiplier2[0], multiplier3[0], multiplier4[0]}
	for i := 0; i < 37; i++ {
		ints[0] = api.Add(ints[0], api.Mul(characters[i], multipliers[0]))
		multipliers[0] = api.Div(multipliers[0], 100)
	}
	for i := 37; i < 74; i++ {
		ints[1] = api.Add(ints[1], api.Mul(characters[i], multipliers[1]))
		multipliers[1] = api.Div(multipliers[1], 100)
	}
	for i := 74; i < 111; i++ {
		ints[2] = api.Add(ints[2], api.Mul(characters[i], multipliers[2]))
		multipliers[2] = api.Div(multipliers[2], 100)
	}
	for i := 111; i < 148; i++ {
		ints[3] = api.Add(ints[3], api.Mul(characters[i], multipliers[3]))
		multipliers[3] = api.Div(multipliers[3], 100)
	}
	return ints
}

func generateMultipliers(_ ecc.ID, inputs []*big.Int, results []*big.Int) error {
	var sb strings.Builder
	for i, input := range inputs {
		sb.WriteString("1")
		for j := 0; j < int(input.Int64()); j++ {
			sb.WriteString("0")
		}
		multiplier, _ := big.NewInt(0).SetString(sb.String(), 10)
		sb.Reset()
		*results[i] = *multiplier
	}
	return nil
}

var multipliers1 = hint.NewStaticHint(generateMultipliers, 1, 1)
var multipliers4 = hint.NewStaticHint(generateMultipliers, 4, 4)
var multipliers7 = hint.NewStaticHint(generateMultipliers, 7, 7)
var multipliers11 = hint.NewStaticHint(generateMultipliers, 11, 11)

func ToInt62(api frontend.API, ints [2]frontend.Variable) frontend.Variable {
	multipliers, _ := api.NewHint(multipliers1, []frontend.Variable{24}...)
	int62 := api.Add(api.Mul(ints[0], multipliers[0]), ints[1])
	return int62
}

func ToInt70(api frontend.API, ints [5]frontend.Variable) frontend.Variable {
	//multipliers := [4]frontend.Variable{big.NewInt(1e50), big.NewInt(1e36), big.NewInt(1e22), big.NewInt(1e10)}
	multipliers, _ := api.NewHint(multipliers4, []frontend.Variable{50, 36, 22, 10}...)
	int70 := api.Mul(ints[0], multipliers[0])
	for i := 1; i < 4; i++ {
		int70 = api.Add(int70, api.Mul(ints[i], multipliers[i]))
	}
	int70 = api.Add(int70, ints[4])
	return int70
}

func ToInt68(api frontend.API, ints [8]frontend.Variable) frontend.Variable {
	//multipliers := [8]frontend.Variable{
	//	big.NewInt(1e66),
	//	big.NewInt(1e56),
	//	big.NewInt(1e48),
	//	big.NewInt(1e40),
	//	big.NewInt(1e32),
	//	big.NewInt(1e24),
	//	big.NewInt(1e16),
	//	big.NewInt(1e8),
	//}
	multipliers, _ := api.NewHint(multipliers7, []frontend.Variable{58, 48, 40, 32, 24, 16, 8}...)
	int76 := api.Mul(ints[0], multipliers[0])
	for i := 1; i < 7; i++ {
		int76 = api.Add(int76, api.Mul(ints[i], multipliers[i]))
	}
	int76 = api.Add(int76, ints[7])
	api.Println(int76)
	return int76
}

func ToInt60(api frontend.API, ints [12]frontend.Variable) frontend.Variable {
	multipliers, _ := api.NewHint(multipliers11, []frontend.Variable{52, 44, 36, 28, 20, 16, 12, 8, 6, 4, 2}...)
	int52 := api.Mul(ints[0], multipliers[0])
	for i := 1; i < 11; i++ {
		int52 = api.Add(int52, api.Mul(ints[i], multipliers[i]))
	}
	api.AssertIsEqual(ints[11], ints[11])
	int52 = api.Add(int52, ints[11])
	api.Println(int52)
	return int52
}

func paddingInt(prefix *big.Int, multipleOf10 int) *big.Int {
	sPrefix := prefix.String()
	sb := strings.Builder{}
	sb.WriteString(sPrefix)
	for j := 0; j < multipleOf10; j++ {
		sb.WriteString("0")
	}
	pInt, _ := big.NewInt(0).SetString(sb.String(), 10)
	return pInt
}

func generatePaddingInt(_ ecc.ID, inputs []*big.Int, results []*big.Int) error {
	pInt := paddingInt(inputs[0], int(inputs[1].Int64()))
	*results[0] = *pInt
	return nil
}

var padIntHint = hint.NewStaticHint(generatePaddingInt, 2, 1)

func Pad48(api frontend.API, int frontend.Variable) frontend.Variable {
	//prepend '19999999999999999999999999999' -> 254 bits
	prefix, _ := big.NewInt(0).SetString("19999999999999999999999999999", 10)
	padding, _ := api.NewHint(padIntHint, []frontend.Variable{prefix, 48}...)
	compareTo, _ := api.NewHint(padIntHint, []frontend.Variable{1, 48}...)
	api.AssertIsEqual(api.Mul(compareTo[0], prefix), padding[0])
	pInt := api.Add(padding[0], int)
	api.Println(pInt)
	//multiplier := big.NewInt(2)
	//for i := 0; i < 87; i++ {
	//	pInt = api.Mul(pInt, multiplier)
	//}
	return pInt
}

func Pad50(api frontend.API, int frontend.Variable) frontend.Variable {
	//prepend '199999999999999999999999999' -> 254 bits
	prefix, _ := big.NewInt(0).SetString("199999999999999999999999999", 10)
	padding, _ := api.NewHint(padIntHint, []frontend.Variable{prefix, 50}...)
	compareTo, _ := api.NewHint(padIntHint, []frontend.Variable{1, 50}...)
	api.AssertIsEqual(api.Mul(compareTo[0], prefix), padding[0])
	pInt := api.Add(padding[0], int)
	api.Println(pInt)
	//for i := 0; i < 81; i++ {
	//	pInt = api.Mul(pInt, big.NewInt(2))
	//	api.Println(pInt)
	//}
	//api.Println(pInt)
	return pInt
}

func Pad60(api frontend.API, int frontend.Variable) frontend.Variable {
	//prepend '19999999999999999' -> 254 bits
	prefix, _ := big.NewInt(0).SetString("19999999999999999", 10)
	padding, _ := api.NewHint(padIntHint, []frontend.Variable{prefix, 60}...)
	compareTo, _ := api.NewHint(padIntHint, []frontend.Variable{1, 60}...)
	api.AssertIsEqual(api.Mul(compareTo[0], prefix), padding[0])
	pInt := api.Add(padding[0], int)
	api.Println(pInt)
	//for i := 0; i < 48; i++ {
	//	pInt = api.Mul(pInt, big.NewInt(2))
	//}
	//api.Println(pInt)
	return pInt
}

func Pad62(api frontend.API, int frontend.Variable) frontend.Variable {
	//prepend '199999999999999' -> 254 bits
	prefix, _ := big.NewInt(0).SetString("199999999999999", 10)
	padding, _ := api.NewHint(padIntHint, []frontend.Variable{prefix, 62}...)
	compareTo, _ := api.NewHint(padIntHint, []frontend.Variable{1, 62}...)
	api.AssertIsEqual(api.Mul(compareTo[0], prefix), padding[0])
	pInt := api.Add(padding[0], int)
	api.Println(pInt)
	//for i := 0; i < 41; i++ {
	//	pInt = api.Mul(pInt, big.NewInt(2))
	//}
	//api.Println(pInt)
	return pInt
}

func Pad68(api frontend.API, int frontend.Variable) frontend.Variable {
	//prepend '199999999' -> 254 bits
	prefix, _ := big.NewInt(0).SetString("199999999", 10)
	padding, _ := api.NewHint(padIntHint, []frontend.Variable{prefix, 68}...)
	compareTo, _ := api.NewHint(padIntHint, []frontend.Variable{1, 68}...)
	api.AssertIsEqual(api.Mul(compareTo[0], prefix), padding[0])
	pInt := api.Add(padding[0], int)
	api.Println(pInt)
	//for i := 0; i < 21; i++ {
	//	pInt = api.Mul(pInt, big.NewInt(2))
	//}
	//api.Println(pInt)
	return pInt
}

func Pad70(api frontend.API, int frontend.Variable) frontend.Variable {
	//prepend '1999999' -> 254 bits
	prefix, _ := big.NewInt(0).SetString("1999999", 10)
	padding, _ := api.NewHint(padIntHint, []frontend.Variable{prefix, 70}...)
	compareTo, _ := api.NewHint(padIntHint, []frontend.Variable{1, 70}...)
	api.AssertIsEqual(api.Mul(compareTo[0], prefix), padding[0])
	pInt := api.Add(padding[0], int)
	api.Println(pInt)
	//for i := 0; i < 14; i++ {
	//	pInt = api.Mul(pInt, big.NewInt(2))
	//}
	//api.Println(pInt)
	return pInt
}

func Pad74(api frontend.API, int frontend.Variable) frontend.Variable {
	//prepend '199' -> 254 bits
	prefix, _ := big.NewInt(0).SetString("199", 10)
	padding, _ := api.NewHint(padIntHint, []frontend.Variable{prefix, 74}...)
	compareTo, _ := api.NewHint(padIntHint, []frontend.Variable{1, 74}...)
	api.AssertIsEqual(api.Mul(compareTo[0], prefix), padding[0])
	pInt := api.Add(padding[0], int)
	api.Println(pInt)
	//for i := 0; i < 4; i++ {
	//	pInt = api.Mul(pInt, big.NewInt(2))
	//}
	return pInt
}

// gnark is a zk-SNARK library written in Go. Circuits are regular structs.
// The inputs must be of type frontend.Variable and make up the witness.
// The witness has a
//       * secret part --> known to the prover only
//       * public part --> known to the prover and the verifier
type FPLCircuit struct {
	FPLForm   CircuitFPLForm    `gnark:",secret"`
	PublicKey eddsa.PublicKey   `gnark:",public"`
	Signature eddsa.Signature   `gnark:",public"`
	Hash      frontend.Variable `gnark:",public"`
}

// Define declares the circuit logic. The compiler then produces a list of constraints
// which must be satisfied (valid witness) in order to create a valid zk-SNARK
// This circuit verifies an EdDSA signature.
func (circuit *FPLCircuit) Define(api frontend.API) error {
	// set the twisted edwards curve to use
	circuit.PublicKey.Curve, _ = twistededwards.NewEdCurve(api.Curve())

	beforeHash := api.Tag("before hash")

	AI := CharactersToInt7(api, circuit.FPLForm.AI)
	FR := circuit.FPLForm.FR[0]
	TOF := circuit.FPLForm.TOF[0]
	NOA := CharactersToInt2(api, circuit.FPLForm.NOA)
	TOA := CharactersToInt4(api, circuit.FPLForm.TOA)
	WTC := circuit.FPLForm.WTC[0]
	EANDC_RCNAAE_S := CharactersToInt2(api, circuit.FPLForm.EANDC.RCNAAE.S)
	EANDC_RCNAAE_A := CharactersToInt49(api, circuit.FPLForm.EANDC.RCNAAE.A)
	EANDC_Surveillance_S := circuit.FPLForm.EANDC.Surveillance.S[0]
	EANDC_Surveillance_SSR_ModeAC_A := CharactersToInt2(api, circuit.FPLForm.EANDC.Surveillance.SSR.ModeAC.A)
	EANDC_Surveillance_SSR_ModeS_A := CharactersToInt7(api, circuit.FPLForm.EANDC.Surveillance.SSR.ModeS.A)
	EANDC_Surveillance_ADS_C_A := CharactersToInt4(api, circuit.FPLForm.EANDC.Surveillance.ADS.C.A)
	EANDC_Surveillance_ADS_B_A := CharactersToInt12(api, circuit.FPLForm.EANDC.Surveillance.ADS.B.A)
	T := CharactersToInt4(api, circuit.FPLForm.T)
	DA := CharactersToInt4(api, circuit.FPLForm.DA)
	CS := CharactersToInt5(api, circuit.FPLForm.CS)
	CL := CharactersToInt5(api, circuit.FPLForm.CL)
	R := CharactersToInt148(api, circuit.FPLForm.R)
	DE := CharactersToInt4(api, circuit.FPLForm.DE)
	TEET := CharactersToInt4(api, circuit.FPLForm.TEET)
	ADA := CharactersToInt4(api, circuit.FPLForm.ADA)
	ADAS_0 := CharactersToInt4(api, circuit.FPLForm.ADAS[0])
	ADAS_1 := CharactersToInt4(api, circuit.FPLForm.ADAS[1])
	OI := CharactersToInt74(api, circuit.FPLForm.OI)
	E := CharactersToInt6(api, circuit.FPLForm.E)
	POB := CharactersToInt5(api, circuit.FPLForm.POB)
	ER := CharactersToInt4(api, circuit.FPLForm.ER)
	SE := CharactersToInt4(api, circuit.FPLForm.SE)
	J := CharactersToInt10(api, circuit.FPLForm.J)
	D := CharactersToInt37(api, circuit.FPLForm.D)
	ACM := CharactersToInt37(api, circuit.FPLForm.ACM)
	RE := CharactersToInt37(api, circuit.FPLForm.RE)
	PN := CharactersToInt19(api, circuit.FPLForm.PN)

	preimage := [15]frontend.Variable{
		Pad74(api, R[0]),
		Pad74(api, R[1]),
		Pad74(api, R[2]),
		Pad74(api, R[3]),
		Pad74(api, OI[0]),
		Pad74(api, OI[1]),
		Pad74(api, D),
		Pad74(api, ACM),
		Pad74(api, RE),
		Pad50(api, EANDC_RCNAAE_A[0]),
		Pad48(api, EANDC_RCNAAE_A[1]),
		Pad62(api, ToInt62(api,
			[2]frontend.Variable{
				PN,
				EANDC_Surveillance_ADS_B_A,
			})),
		Pad70(api, ToInt70(
			api,
			[5]frontend.Variable{
				J,
				AI,
				EANDC_Surveillance_SSR_ModeS_A,
				E,
				CS,
			})),
		Pad68(api, ToInt68(
			api,
			[8]frontend.Variable{
				CL,
				POB,
				TOA,
				EANDC_Surveillance_ADS_C_A,
				T,
				DA,
				DE,
				TEET,
			})),
		Pad60(api, ToInt60(
			api,
			[12]frontend.Variable{
				ADA,
				ADAS_0,
				ADAS_1,
				ER,
				SE,
				NOA,
				EANDC_RCNAAE_S,
				EANDC_Surveillance_SSR_ModeAC_A,
				FR,
				TOF,
				WTC,
				EANDC_Surveillance_S,
			})),
	}

	//for i := 0; i < 15; i++ {
	//	api.Println(preimage[i], i)
	//}

	hash, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}
	for i := 0; i < 15; i++ {
		hash.Write(preimage[i])
	}
	h := hash.Sum()
	api.AssertIsEqual(h, circuit.Hash)

	api.AddCounter(beforeHash, api.Tag("after hash"))

	//circuit.Message =  make([]frontend.Variable, 2)

	// tip: gnark counters enable circuit developers to measure the number of constraints
	// generated by a part of the circuit

	//beforeInt := api.Tag("before int")
	//
	//Test(api, circuit.Test)
	//AssertDigits(api, circuit.AI)
	//api.Println("digitsAsInt", DigitsToInt(api, circuit.AI))
	//
	//api.AddCounter(beforeInt, api.Tag("after int"))

	//assert that the digitsToInt of aircraft identifier is calculated correctly

	beforeVerify := api.Tag("before verify")
	// verify the EdDSA signature
	eddsa.Verify(api, circuit.Signature, circuit.Hash, circuit.PublicKey)

	api.AddCounter(beforeVerify, api.Tag("after verify"))

	return nil
}

func proveFPL(input *ProveFPLInput) groth16.Proof {

	parsedFPL := parseFPL(input.FPLForm)
	circuitFPL := parsedFPLToCircuitFPL(parsedFPL)
	dHash, _ := hex.DecodeString(input.EHash)
	dSig, _ := hex.DecodeString(input.ESignature)
	dPk, _ := hex.DecodeString(input.EPublicKey)

	fmt.Printf("\npk is %+v", input.EPublicKey)
	fmt.Printf("\nsig is %+v", input.ESignature)
	fmt.Printf("\nhash is %+v", input.EHash)

	signature.Register(signature.EDDSA_BN254, eddsabn254.GenerateKeyInterfaces)
	privateKey, _ := signature.EDDSA_BN254.New(rand.Reader)
	publicKey := privateKey.Public()
	publicKey.SetBytes(dPk)

	isValid := verifySig(input.ESignature, dPk, dHash)
	fmt.Println(isValid)

	var circuit FPLCircuit
	r1cs, err := frontend.Compile(ecc.BN254, backend.GROTH16, &circuit)
	if err != nil {
		fmt.Printf("%+v", err.Error())
	}
	//fmt.Printf("first r1cs %+v \n", r1cs)
	fmt.Printf("counters %+v \n", r1cs.GetCounters())

	pk, vk, _ := groth16.Setup(r1cs)

	// declare the witness
	var assignment FPLCircuit

	assignment.FPLForm = circuitFPL
	assignment.Hash = dHash

	// public key bytes
	_publicKey := publicKey.Bytes()

	// temporary point
	var p edwardsbn254.PointAffine

	// assign public key values
	p.SetBytes(_publicKey[:32])
	axb := p.X.Bytes()
	ayb := p.Y.Bytes()
	assignment.PublicKey.A.X = axb[:]
	assignment.PublicKey.A.Y = ayb[:]

	// assign signature values
	p.SetBytes(dSig[:32])
	rxb := p.X.Bytes()
	ryb := p.Y.Bytes()
	assignment.Signature.R.X = rxb[:]
	assignment.Signature.R.Y = ryb[:]
	assignment.Signature.S = dSig[32:]

	witness, err := frontend.NewWitness(&assignment, ecc.BN254)
	publicWitness, err := witness.Public()

	// generate the proof
	proveOpts := []backend.ProverOption{
		backend.WithHints([]hint.Function{
			pMultiplier,
			multipliers1,
			multipliers4,
			multipliers7,
			multipliers11,
			padIntHint}...),
	}
	proof, err := groth16.Prove(r1cs, pk, witness, proveOpts...)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// verify the proof
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		// invalid proof
		fmt.Printf("%v", err)
	} else {
		fmt.Printf("Proof is verified!")
	}

	return proof
}
