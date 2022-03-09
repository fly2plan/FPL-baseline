package main

import (
	"fmt"
	"github.com/consensys/gnark-crypto/hash"
	"math/big"
	"strings"
)

var hFunc = hash.MIMC_BN254.New()

type Hash struct {
	Hash string `json:"hash"'`
}

func concat(arrays [][]int) [][]int {
	if len(arrays) == 1 {
		return arrays
	}
	var array []int
	array = append(array, arrays[0]...)
	array = append(array, arrays[1]...)
	arrays = arrays[2:]
	arrays = append([][]int{array}, arrays...)
	return concat(arrays)
}

func concatArraysOfInts(arrays [][]int) []int {
	arrays = append([][]int{}, arrays...)
	array := concat(arrays)[0]
	return array
}

//func digitsToInt(digits []int) big.Int {
//	var result big.Int
//	multiplier := new(big.Int).SetInt64(1)
//	for i := 0; i < len(digits); i++ {
//		var addition big.Int
//		digit := new(big.Int).SetInt64(int64(digits[i]))
//		addition.Mul(digit, multiplier)
//		result.Add(&result, &addition)
//		multiplier.Mul(multiplier, new(big.Int).SetInt64(10))
//	}
//	return result
//}

func initialMultiplier(lenOfCharacters int) *big.Int {
	var sb strings.Builder
	sb.WriteString("1")
	for j := 0; j < lenOfCharacters-1; j++ {
		sb.WriteString("00")
	}
	multiplier, _ := big.NewInt(0).SetString(sb.String(), 10)
	return multiplier
}

func charactersToInt(characters []int) *big.Int {
	var result big.Int
	//multiplier := new(big.Int).SetInt64(1)
	multiplier := initialMultiplier(len(characters))
	//fmt.Println(multiplier)
	for i := 0; i < len(characters); i++ {
		var addition big.Int
		character := new(big.Int).SetInt64(int64(characters[i]))
		addition.Mul(character, multiplier)
		result.Add(&result, &addition)
		multiplier.Div(multiplier, new(big.Int).SetInt64(100))
		//fmt.Println(result, multiplier)
	}

	switch len(characters) {
	case 37:
		prefix, _ := big.NewInt(0).SetString("199", 10)
		result.Add(&result, paddingInt(prefix, 74))
		break
	case 25:
		prefix, _ := big.NewInt(0).SetString("199999999999999999999999999", 10)
		result.Add(&result, paddingInt(prefix, 50))
		break
	case 24:
		prefix, _ := big.NewInt(0).SetString("19999999999999999999999999999", 10)
		result.Add(&result, paddingInt(prefix, 48))
		break
	case 31:
		prefix, _ := big.NewInt(0).SetString("199999999999999", 10)
		result.Add(&result, paddingInt(prefix, 62))
		break
	case 35:
		prefix, _ := big.NewInt(0).SetString("1999999", 10)
		result.Add(&result, paddingInt(prefix, 70))
		break
	case 34:
		prefix, _ := big.NewInt(0).SetString("199999999", 10)
		result.Add(&result, paddingInt(prefix, 68))
		break
	case 30:
		prefix, _ := big.NewInt(0).SetString("19999999999999999", 10)
		result.Add(&result, paddingInt(prefix, 60))
		break
	}
	return &result
}

//Generate a hash from a field element as big int
func generateHash(element *big.Int) []byte {
	bytes := element.Bytes()
	hFunc.Write(bytes[:])
	return hFunc.Sum(nil)
}

//Generate a hash from hashes which add to more than 254 bits
func compressToHash(hashes [][]byte) []byte {
	var sum []byte
	for _, s := range hashes {
		hFunc.Write(s)
		sum = hFunc.Sum(nil)
	}
	//fmt.Printf("%+v", hex.EncodeToString(sum))
	return sum
}

func hashPreimage(preimage []*big.Int) []byte {
	for _, i := range preimage {
		fmt.Println(" ", i.String(), i.BitLen())
		hFunc.Write(i.Bytes())
	}
	return hFunc.Sum(nil)
}

func hashFPL(parsedForm ParsedFPLForm) []byte {
	preimage := make([]*big.Int, 15)
	preimage[0] = charactersToInt(parsedForm.R[0:37])
	preimage[1] = charactersToInt(parsedForm.R[37:74])
	preimage[2] = charactersToInt(parsedForm.R[74:111])
	preimage[3] = charactersToInt(parsedForm.R[111:148])
	preimage[4] = charactersToInt(parsedForm.OI[0:37])
	preimage[5] = charactersToInt(parsedForm.OI[37:74])
	preimage[6] = charactersToInt(parsedForm.D[0:37])
	preimage[7] = charactersToInt(parsedForm.ACM[0:37])
	preimage[8] = charactersToInt(parsedForm.RE[0:37])
	preimage[9] = charactersToInt(parsedForm.EANDC.RCNAAE.A[0:25])
	preimage[10] = charactersToInt(parsedForm.EANDC.RCNAAE.A[25:49])
	preimage[11] = charactersToInt(concatArraysOfInts(
		[][]int{
			parsedForm.PN[0:19],
			parsedForm.EANDC.Surveillance.ADS.B.A[0:12],
		}))
	preimage[12] = charactersToInt(concatArraysOfInts(
		[][]int{
			parsedForm.J[0:10],
			parsedForm.AI[0:7],
			parsedForm.EANDC.Surveillance.SSR.ModeS.A[0:7],
			parsedForm.E[0:6],
			parsedForm.CS[0:5],
		}))
	preimage[13] = charactersToInt(concatArraysOfInts(
		[][]int{
			parsedForm.CL[0:5],
			parsedForm.POB[0:5],
			parsedForm.TOA[0:4],
			parsedForm.EANDC.Surveillance.ADS.C.A[0:4],
			parsedForm.T[0:4],
			parsedForm.DA[0:4],
			parsedForm.DE[0:4],
			parsedForm.TEET[0:4],
		}))
	preimage[14] = charactersToInt(concatArraysOfInts(
		[][]int{
			parsedForm.ADA[0:4],
			parsedForm.ADAS[0][0:4],
			parsedForm.ADAS[1][0:4],
			parsedForm.ER[0:4],
			parsedForm.SE[0:4],
			parsedForm.NOA[0:2],
			parsedForm.EANDC.RCNAAE.S[0:2],
			parsedForm.EANDC.Surveillance.SSR.ModeAC.A[0:2],
			parsedForm.FR[0:1],
			parsedForm.TOF[0:1],
			parsedForm.WTC[0:1],
			parsedForm.EANDC.Surveillance.S[0:1],
		}))
	//hashes := make([][]byte, 15)
	//for i := 0; i < 15; i++ {
	//	hashes[i] = generateHash(preimage[i])
	//	//fmt.Printf("\n%+v", hex.EncodeToString(hashes[i]))
	//}
	////fmt.Printf("%+v", hashes)
	//hash := compressToHash(hashes)
	hFunc.Reset()
	return hashPreimage(preimage)
}
