package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"service/pkg/fpl"
	"service/pkg/hash"
	"service/pkg/prove"
	"strings"
)

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type FileStream struct {
	Files  []*File `json:"files"`
	State  string  `json:"state"`
}

func contractsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fileStream := &FileStream{}
		currentPath := "./"
		err := filepath.Walk("./contracts", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatalf(err.Error())
			}
			//if info.IsDir() {
			//	fmt.Printf("skipping a directory without errors: %+v \n", info.Name())
			//	return filepath.SkipDir
			//}
			if info.IsDir() {
				currentPath = currentPath + info.Name() + "/"
			}
			if !info.IsDir() {
				fmt.Printf("current path: %s\n", currentPath)
				fmt.Printf("found file: %s\n", info.Name())
				bytes, err := os.ReadFile(currentPath + info.Name())
				if err != nil {
					fmt.Println(err)
				}
				fileStream.Files = append(fileStream.Files, &File{
					Name: info.Name(),
					Content: strings.Replace(string(bytes), "\n", "", -1),
				})
			}
			return nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			fileStream.State = "finished"
			j, _ := json.Marshal(fileStream)
			w.Write(j)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not supported")
	}
}

func keyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		keys := generateKeys()
		fmt.Println("keys", keys)
		j, err := json.Marshal(keys)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(j)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not supported")
	}
}

func hashHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		vars := mux.Vars(r)
		switch vars["type"] {
		case "fpl":
			dFPL := json.NewDecoder(r.Body)
			pFPL := &fpl.FPLForm{}
			err := dFPL.Decode(pFPL)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			parsedFPL := fpl.ParseFPL(*pFPL)
			FPLhash := hash.HashFPL(parsedFPL)
			pHash := &hash.Hash{
				Hash: hex.EncodeToString(FPLhash),
			}
			fmt.Printf("\n%+v", pFPL)
			fmt.Printf("\n%+v", parsedFPL)
			j, err := json.Marshal(pHash)
			w.Write(j)
			break
		case "ack":
			dHash := json.NewDecoder(r.Body)
			pProof := &prove.FPLProveInput{}
			err := dHash.Decode(pProof)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			h, _ := hex.DecodeString(pProof.EHash)
			ackHash := hash.HashACK(h)
			j, err := json.Marshal(ackHash)
			w.Write(j)
			break
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not supported")
	}
}

func signHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	msg := query.Get("msg")
	if msg == "" {
		fmt.Fprintf(w, "msg to sign not present")
	}
	switch r.Method {
	case "POST":
		d := json.NewDecoder(r.Body)
		k := &Keys{}
		_ = d.Decode(k)
		sk, _ := hex.DecodeString(k.Sk)
		pk, _ := hex.DecodeString(k.Pk)
		fmt.Println("msg is", msg)
		fmt.Println("sk is", k.Sk)
		fmt.Println("pk is", k.Pk)
		msgBin, _ := hex.DecodeString(msg)
		sig := signMessage(sk, msgBin)
		s := &Signature{
			Sig: sig,
		}
		fmt.Println("sig is", sig)
		j, err := json.Marshal(s)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		if verifySig(sig, pk, msgBin) {
			w.Write(j)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not supported")
	}
}

func proveHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		vars := mux.Vars(r)
		switch vars["type"] {
		case "fpl":
			dFPLInput := json.NewDecoder(r.Body)
			pFPLInput := &prove.FPLInput{}
			err := dFPLInput.Decode(pFPLInput)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			parsedFPL := fpl.ParseFPL(*pFPLInput.FPLForm)
			FPLhash := hash.HashFPL(parsedFPL)
			Ehash := hex.EncodeToString(FPLhash)
			skBytes, _ := hex.DecodeString(pFPLInput.EPrivateKey)
			AOSig := signMessage(skBytes, FPLhash)
			AOPublicKey := getPublicKey(skBytes)
			pProveInput := &prove.FPLProveInput{
				FPLForm:    pFPLInput.FPLForm,
				EHash:      Ehash,
				ESignature: AOSig,
				EPublicKey: AOPublicKey,
			}
			proof := prove.ProveFPL(pProveInput)
			j, err := json.Marshal(proof)
			w.Write(j)
			break
		case "ack":
			dACKInput := json.NewDecoder(r.Body)
			pACKInput := &prove.ACKInput{}
			err := dACKInput.Decode(pACKInput)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			FPLhashBytes, _ := hex.DecodeString(pACKInput.Proof.EHash)
			ACKhashBytes := hash.HashACK(FPLhashBytes)
			Ehash := hex.EncodeToString(ACKhashBytes)
			skBytes, _ := hex.DecodeString(pACKInput.EPrivateKey)
			NMSig := signMessage(skBytes, ACKhashBytes)
			NMPublicKey := getPublicKey(skBytes)
			pkBytes, _ := hex.DecodeString(NMPublicKey)
			fmt.Printf("\n verified sig %+v", verifySig(NMSig, pkBytes, ACKhashBytes))
			pProveInput := &prove.ACKProveInput{
				Proof: pACKInput.Proof,
				EHash: Ehash,
				ESignature: NMSig,
				EPublicKey: NMPublicKey,
			}
			proof := prove.ProveACK(pProveInput)
			j, err := json.Marshal(proof)
			w.Write(j)
			break
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not supported")
	}
}

func formatHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		vars := mux.Vars(r)
		switch vars["type"] {
		case "fpl":
			dFormatInput := json.NewDecoder(r.Body)
			pFormatInput := &prove.FormattedFPLProof{}
			err := dFormatInput.Decode(pFormatInput)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			inputs := prove.GetFPLProofInputs(pFormatInput.Inputs)
			j, err := json.Marshal(&prove.FPLProof{
				Proof:  pFormatInput.Proof,
				Inputs: inputs,
			})
			fmt.Println(inputs.Hash, inputs.PublicKey, inputs.Signature)
			w.Write(j)
		case "ack":
			dFormatInput := json.NewDecoder(r.Body)
			pFormatInput := &prove.FormattedACKProof{}
			err := dFormatInput.Decode(pFormatInput)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			inputs := prove.GetACKProofInputs(pFormatInput.Inputs)
			j, err := json.Marshal(&prove.ACKProof{
				Proof:  pFormatInput.Proof,
				Inputs: inputs,
			})
			w.Write(j)
			break
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not supported")
	}
}

func main() {

	rtr := mux.NewRouter()
	rtr.HandleFunc("/contracts", contractsHandler)
	rtr.HandleFunc("/keys", keyHandler)
	rtr.HandleFunc("/hash/{type:.*}", hashHandler)
	rtr.HandleFunc("/sign", signHandler)
	rtr.HandleFunc("/prove/{type:.*}", proveHandler)
	rtr.HandleFunc("/format/{type:.*}", formatHandler)
	http.Handle("/", rtr)

	log.Println("Go!")
	http.ListenAndServe(":3000", nil)
}
