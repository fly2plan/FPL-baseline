package prove

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"service/pkg/fpl"
	"service/pkg/hash"
	"strings"
	"time"

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
)

func generateMultiplier(_ ecc.ID, inputs []*big.Int, results []*big.Int) error {
	multiplier := hash.InitialMultiplier(int(inputs[0].Int64()))
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


func generatePaddingInt(_ ecc.ID, inputs []*big.Int, results []*big.Int) error {
	pInt := hash.PaddingInt(inputs[0], int(inputs[1].Int64()))
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
	FPLForm   fpl.CircuitFPLForm `gnark:",secret"`
	PublicKey eddsa.PublicKey    `gnark:",public"`
	Signature eddsa.Signature    `gnark:",public"`
	Hash      frontend.Variable  `gnark:",public"`
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

	h, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}
	for i := 0; i < 15; i++ {
		h.Write(preimage[i])
	}
	hh := h.Sum()
	api.AssertIsEqual(hh, circuit.Hash)

	api.AddCounter(beforeHash, api.Tag("after hash"))

	beforeVerify := api.Tag("before verify")
	// verify the EdDSA signature
	eddsa.Verify(api, circuit.Signature, circuit.Hash, circuit.PublicKey)

	api.AddCounter(beforeVerify, api.Tag("after verify"))

	return nil
}

type ACKCircuit struct {
	FPLForm     fpl.CircuitFPLForm `gnark:",secret"`
	AOPublicKey eddsa.PublicKey    `gnark:",public"`
	AOSignature eddsa.Signature    `gnark:",public"`
	NMPublicKey eddsa.PublicKey    `gnark:",public"`
	NMSignature eddsa.Signature    `gnark:",public"`
	FPLHash     frontend.Variable  `gnark:",public"`
	ACKHash     frontend.Variable  `gnark:",public"`
}

func (circuit *ACKCircuit) Define(api frontend.API) error {
	// set the twisted edwards curve to use
	circuit.AOPublicKey.Curve, _ = twistededwards.NewEdCurve(api.Curve())
	circuit.NMPublicKey.Curve, _ = twistededwards.NewEdCurve(api.Curve())

	beforeTotal := api.Tag("before total")

	beforeFPLHash := api.Tag("before hash")

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

	h, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}
	for i := 0; i < 15; i++ {
		h.Write(preimage[i])
	}
	fplh := h.Sum()
	api.AssertIsEqual(fplh, circuit.FPLHash)
	api.AddCounter(beforeFPLHash, api.Tag("after fpl hash"))

	beforeACKHash := api.Tag("before ack hash")
	h.Reset()
	h.Write(circuit.FPLHash)
	ack := frontend.Variable(big.NewInt(hash.ACK))
	padding, _ := big.NewInt(0).SetString(hash.ACKPadding, 10)
	ackPadding := frontend.Variable(padding)
	h.Write(api.Mul(ack, ackPadding))
	ackh := h.Sum()
	api.AssertIsEqual(ackh, circuit.ACKHash)
	api.AddCounter(beforeACKHash, api.Tag("after ack hash"))

	beforeVerify := api.Tag("before verify")
	// verify the EdDSA signatures
	eddsa.Verify(api, circuit.AOSignature, circuit.FPLHash, circuit.AOPublicKey)
	eddsa.Verify(api, circuit.NMSignature, circuit.ACKHash, circuit.NMPublicKey)
	api.AddCounter(beforeVerify, api.Tag("after verify"))

	api.AddCounter(beforeTotal, api.Tag("after total"))
	api.Println(api.Tag("after total"))
	return nil
}

type Verifier struct {
	State       string `json:"state"`
	FPLVerifier string
}

type FPLInput struct {
	*fpl.FPLForm 	   `json:"fpl"`
	EPrivateKey string `json:"sk"`
}

type ACKInput struct {
	*ACKProveInput
	EPrivateKey string `json:"sk"`
}

type FPLProveInput struct {
	*fpl.FPLForm      `json:"fpl"`
	EHash      string `json:"hash"`
	ESignature string `json:"sig"`
	EPublicKey string `json:"pk"`
}

type ACKProveInput struct {
	Proof      *FPLProveInput `json:"proof"`
	EHash      string 	 	  `json:"hash"`
	ESignature string         `json:"sig"`
	EPublicKey string         `json:"pk"`
}

type Proof struct {
	a [2]*big.Int
	b [2][2]*big.Int
	c [2]*big.Int
}

type FormattedProof struct {
	A [2]string    `json:"a"`
	B [2][2]string `json:"b"`
	C [2]string    `json:"c"`
}

type FPLProofInput struct {
	Hash      string    `json:"hash"`
	Signature [3]string `json:"sig"`
	PublicKey [2]string `json:"pk"`
}

type FormattedFPLProofInput struct {
	Hash      string `json:"hash"`
	Signature string `json:"sig"`
	PublicKey string `json:"pk"`
}

type FormattedFPLProof struct {
	Proof  *FormattedProof         `json:"proof"`
	Inputs *FormattedFPLProofInput `json:"inputs""`
}

type FPLProof struct {
	Proof   *FormattedProof `json:"proof"`
	Inputs  *FPLProofInput  `json:"inputs"`
}

type ACKProof struct {
	Proof   *FormattedProof `json:"proof"`
	Inputs  *ACKProofInput  `json:"inputs"`
}

type ACKProofInput struct {
	AOPublicKey [2]string `json:"ao_pk"`
	AOSignature [3]string `json:"ao_sig"`
	NMPublicKey [2]string `json:"nm_pk"`
	NMSignature [3]string `json:"nm_sig"`
	FPLHash      string   `json:"fpl_hash"`
	ACKHash      string   `json:"ack_hash"`
}

type FormattedACKProofInput struct {
	AOPublicKey string `json:"ao_pk"`
	AOSignature string `json:"ao_sig"`
	NMPublicKey string `json:"nm_pk"`
	NMSignature string `json:"nm_sig"`
	FPLHash     string `json:"fpl_hash"`
	ACKHash     string `json:"ack_hash"`
}

type FormattedACKProof struct {
	Proof  *FormattedProof         `json:"proof"`
	Inputs *FormattedACKProofInput `json:"inputs""`
}

func ProveFPL(input *FPLProveInput) *FormattedFPLProof {

	start := time.Now()
	parsedFPL := fpl.ParseFPL(*input.FPLForm)
	circuitFPL := fpl.ParsedFPLToCircuitFPL(parsedFPL)
	dHash, _ := hex.DecodeString(input.EHash)
	dSig, _ := hex.DecodeString(input.ESignature)
	dPk, _ := hex.DecodeString(input.EPublicKey)

	r1csf, err := os.Open("./keys/fpl_r1cs")
	if err != nil {
		fmt.Println(err)
	}
	pkf, err := os.Open("./keys/fpl_pk")
	if err != nil {
		fmt.Println(err)
	}
	r1cs := groth16.NewCS(ecc.BN254)
	_, err = r1cs.ReadFrom(r1csf)
	if err != nil {
		fmt.Println(err)
	}
	pk := groth16.NewProvingKey(ecc.BN254)
	_, _ = pk.ReadFrom(pkf)

	signature.Register(signature.EDDSA_BN254, eddsabn254.GenerateKeyInterfaces)
	privateKey, _ := signature.EDDSA_BN254.New(rand.Reader)
	publicKey := privateKey.Public()
	publicKey.SetBytes(dPk)

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

	// generate the proof
	proveOpts := []backend.ProverOption{
		backend.WithHints([]hint.Function{
			pMultiplier,
			multipliers1,
			multipliers4,
			multipliers7,
			multipliers11,
			padIntHint,
		}...),
	}

	proof, err := groth16.Prove(r1cs, pk, witness, proveOpts...)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// get proof bytes
	const fpSize = 4 * 8
	var buf bytes.Buffer
	proof.WriteRawTo(&buf)
	proofBytes := buf.Bytes()

	formattedProof := &FormattedProof{}
	// solidity contract inputs

	// proof.Ar, proof.Bs, proof.Krs
	formattedProof.A[0] = new(big.Int).SetBytes(proofBytes[fpSize*0 : fpSize*1]).String()
	formattedProof.A[1] = new(big.Int).SetBytes(proofBytes[fpSize*1 : fpSize*2]).String()
	formattedProof.B[0][0] = new(big.Int).SetBytes(proofBytes[fpSize*2 : fpSize*3]).String()
	formattedProof.B[0][1] = new(big.Int).SetBytes(proofBytes[fpSize*3 : fpSize*4]).String()
	formattedProof.B[1][0] = new(big.Int).SetBytes(proofBytes[fpSize*4 : fpSize*5]).String()
	formattedProof.B[1][1] = new(big.Int).SetBytes(proofBytes[fpSize*5 : fpSize*6]).String()
	formattedProof.C[0] = new(big.Int).SetBytes(proofBytes[fpSize*6 : fpSize*7]).String()
	formattedProof.C[1] = new(big.Int).SetBytes(proofBytes[fpSize*7 : fpSize*8]).String()

	FormattedFPLProof := &FormattedFPLProof{
		Proof: formattedProof,
		Inputs: &FormattedFPLProofInput{
			Hash: input.EHash,
			Signature: input.ESignature,
			PublicKey: input.EPublicKey,
		},
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("elapsed time is", elapsed)

	return FormattedFPLProof
}

func GetFPLProofInputs(proofInput *FormattedFPLProofInput) *FPLProofInput {
	dHash, _ := hex.DecodeString(proofInput.Hash)
	dSig, _ := hex.DecodeString(proofInput.Signature)
	dPk, _ := hex.DecodeString(proofInput.PublicKey)

	signature.Register(signature.EDDSA_BN254, eddsabn254.GenerateKeyInterfaces)
	privateKey, _ := signature.EDDSA_BN254.New(rand.Reader)
	publicKey := privateKey.Public()
	publicKey.SetBytes(dPk)

	// public key bytes
	_publicKey := publicKey.Bytes()

	// temporary point
	var p edwardsbn254.PointAffine

	// assign public key values
	p.SetBytes(_publicKey[:32])
	axb := p.X.Bytes()
	ayb := p.Y.Bytes()

	// assign signature values
	p.SetBytes(dSig[:32])
	rxb := p.X.Bytes()
	ryb := p.Y.Bytes()

	// solidity contract input
	FPLProofInput :=  &FPLProofInput{
		PublicKey: [2]string{
			big.NewInt(0).SetBytes(axb[:]).String(),
			big.NewInt(0).SetBytes(ayb[:]).String(),
		},
		Signature: [3]string{
			big.NewInt(0).SetBytes(rxb[:]).String(),
			big.NewInt(0).SetBytes(ryb[:]).String(),
			big.NewInt(0).SetBytes(dSig[32:]).String(),
		},
		Hash: big.NewInt(0).SetBytes(dHash).String(),
	}

	return FPLProofInput
}

func GetACKProofInputs(proofInput *FormattedACKProofInput) *ACKProofInput {
	dFPLHash, _ := hex.DecodeString(proofInput.FPLHash)
	dACKHash, _ := hex.DecodeString(proofInput.ACKHash)
	dAOSig, _ := hex.DecodeString(proofInput.AOSignature)
	dAOPk, _ := hex.DecodeString(proofInput.AOPublicKey)
	dNMSig, _ := hex.DecodeString(proofInput.NMSignature)
	dNMPk, _ := hex.DecodeString(proofInput.NMPublicKey)

	signature.Register(signature.EDDSA_BN254, eddsabn254.GenerateKeyInterfaces)
	privateKey, _ := signature.EDDSA_BN254.New(rand.Reader)
	publicKey := privateKey.Public()

	// temporary point
	var p edwardsbn254.PointAffine

	// AO public key and sig
	publicKey.SetBytes(dAOPk)
	_publicKey := publicKey.Bytes()

	// assign public key values
	p.SetBytes(_publicKey[:32])
	AOaxb := p.X.Bytes()
	AOayb := p.Y.Bytes()

	// assign signature values
	p.SetBytes(dAOSig[:32])
	AOrxb := p.X.Bytes()
	AOryb := p.Y.Bytes()

	// NM public key and sig
	publicKey.SetBytes(dNMPk)
	_publicKey = publicKey.Bytes()

	// assign public key values
	p.SetBytes(_publicKey[:32])
	NMaxb := p.X.Bytes()
	NMayb := p.Y.Bytes()

	// assign signature values
	p.SetBytes(dNMSig[:32])
	NMrxb := p.X.Bytes()
	NMryb := p.Y.Bytes()

	// solidity contract input
	ACKProofInput :=  &ACKProofInput{
		FPLHash: big.NewInt(0).SetBytes(dFPLHash).String(),
		ACKHash: big.NewInt(0).SetBytes(dACKHash).String(),
		AOSignature: [3]string{
			big.NewInt(0).SetBytes(AOrxb[:]).String(),
			big.NewInt(0).SetBytes(AOryb[:]).String(),
			big.NewInt(0).SetBytes(dAOSig[32:]).String(),
		},
		AOPublicKey: [2]string{
			big.NewInt(0).SetBytes(AOaxb[:]).String(),
			big.NewInt(0).SetBytes(AOayb[:]).String(),
		},
		NMSignature: [3]string{
			big.NewInt(0).SetBytes(NMrxb[:]).String(),
			big.NewInt(0).SetBytes(NMryb[:]).String(),
			big.NewInt(0).SetBytes(dNMSig[32:]).String(),
		},
		NMPublicKey: [2]string{
			big.NewInt(0).SetBytes(NMaxb[:]).String(),
			big.NewInt(0).SetBytes(NMayb[:]).String(),
		},
	}

	return ACKProofInput
}

func ProveACK(input *ACKProveInput) *FormattedACKProof {

	parsedFPL := fpl.ParseFPL(*input.Proof.FPLForm)
	circuitFPL := fpl.ParsedFPLToCircuitFPL(parsedFPL)
	dFPLHash, _ := hex.DecodeString(input.Proof.EHash)
	dAOSig, _ := hex.DecodeString(input.Proof.ESignature)
	dAOPk, _ := hex.DecodeString(input.Proof.EPublicKey)
	dACKHash, _ := hex.DecodeString(input.EHash)
	dNMSig, _ := hex.DecodeString(input.ESignature)
	dNMPk, _ := hex.DecodeString(input.EPublicKey)

	r1csf, _ := os.Open("./keys/ack_r1cs")
	pkf, _ := os.Open("./keys/ack_pk")

	r1cs := groth16.NewCS(ecc.BN254)
	_, _ = r1cs.ReadFrom(r1csf)
	pk := groth16.NewProvingKey(ecc.BN254)
	_, _ = pk.ReadFrom(pkf)

	signature.Register(signature.EDDSA_BN254, eddsabn254.GenerateKeyInterfaces)

	AOPrivateKey, _ := signature.EDDSA_BN254.New(rand.Reader)
	AOPublicKey := AOPrivateKey.Public()
	AOPublicKey.SetBytes(dAOPk)

	NMPrivateKey, _ := signature.EDDSA_BN254.New(rand.Reader)
	NMPublicKey := NMPrivateKey.Public()
	NMPublicKey.SetBytes(dNMPk)

	// declare the witness
	var assignment ACKCircuit

	assignment.FPLForm = circuitFPL
	assignment.FPLHash = dFPLHash
	assignment.ACKHash = dACKHash

	// AO public key bytes
	_AOpublicKey := AOPublicKey.Bytes()
	// NM public key bytes
	_NMpublicKey := NMPublicKey.Bytes()

	// temporary point
	var p edwardsbn254.PointAffine

	// assign public key values of AO
	p.SetBytes(_AOpublicKey[:32])
	axb := p.X.Bytes()
	ayb := p.Y.Bytes()
	assignment.AOPublicKey.A.X = axb[:]
	assignment.AOPublicKey.A.Y = ayb[:]

	// assign signature values of AO
	p.SetBytes(dAOSig[:32])
	rxb := p.X.Bytes()
	ryb := p.Y.Bytes()
	assignment.AOSignature.R.X = rxb[:]
	assignment.AOSignature.R.Y = ryb[:]
	assignment.AOSignature.S = dAOSig[32:]

	// assign public key values of NM
	p.SetBytes(_NMpublicKey[:32])
	axb = p.X.Bytes()
	ayb = p.Y.Bytes()
	assignment.NMPublicKey.A.X = axb[:]
	assignment.NMPublicKey.A.Y = ayb[:]

	// assign signature values of NM
	p.SetBytes(dNMSig[:32])
	rxb = p.X.Bytes()
	ryb = p.Y.Bytes()
	assignment.NMSignature.R.X = rxb[:]
	assignment.NMSignature.R.Y = ryb[:]
	assignment.NMSignature.S = dNMSig[32:]

	witness, err := frontend.NewWitness(&assignment, ecc.BN254)
	if err != nil {
		fmt.Printf("\n witness error: %+v", err)
	}

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
		fmt.Printf("\n prove error: %+v", err)
	}

	fmt.Printf("\n proof %+v", err)

	// get proof bytes
	const fpSize = 4 * 8
	var buf bytes.Buffer
	_, err = proof.WriteRawTo(&buf)
	if err != nil {
		fmt.Printf("\n buffer error: %+v", err)
	}
	proofBytes := buf.Bytes()

	formattedProof := &FormattedProof{}

	// proof.Ar, proof.Bs, proof.Krs
	formattedProof.A[0] = new(big.Int).SetBytes(proofBytes[fpSize*0 : fpSize*1]).String()
	formattedProof.A[1] = new(big.Int).SetBytes(proofBytes[fpSize*1 : fpSize*2]).String()
	formattedProof.B[0][0] = new(big.Int).SetBytes(proofBytes[fpSize*2 : fpSize*3]).String()
	formattedProof.B[0][1] = new(big.Int).SetBytes(proofBytes[fpSize*3 : fpSize*4]).String()
	formattedProof.B[1][0] = new(big.Int).SetBytes(proofBytes[fpSize*4 : fpSize*5]).String()
	formattedProof.B[1][1] = new(big.Int).SetBytes(proofBytes[fpSize*5 : fpSize*6]).String()
	formattedProof.C[0] = new(big.Int).SetBytes(proofBytes[fpSize*6 : fpSize*7]).String()
	formattedProof.C[1] = new(big.Int).SetBytes(proofBytes[fpSize*7 : fpSize*8]).String()

	ACKProof := &FormattedACKProof{
		Proof: formattedProof,
		Inputs: &FormattedACKProofInput{
			AOPublicKey: input.Proof.EPublicKey,
			AOSignature: input.Proof.ESignature,
			NMPublicKey: input.EPublicKey,
			NMSignature: input.ESignature,
			FPLHash: input.Proof.EHash,
			ACKHash: input.EHash,
		},
	}

	return ACKProof

}
