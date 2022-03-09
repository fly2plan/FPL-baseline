package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func keyHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		keys := generateKeys()
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
			// Decode the JSON in the body and overwrite 'tom' with it
			dFPL := json.NewDecoder(r.Body)
			pFPL := &FPLForm{}
			err := dFPL.Decode(pFPL)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			parsedFPL := parseFPL(*pFPL)
			FPLhash := hashFPL(parsedFPL)
			pHash := &Hash{
				Hash: hex.EncodeToString(FPLhash),
			}
			fmt.Printf("\n%+v", pFPL)
			fmt.Printf("\n%+v", parsedFPL)
			j, err := json.Marshal(pHash)
			w.Write(j)
			break
		case "ack":
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
		// Decode the JSON in the body and overwrite 'tom' with it
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
			// Decode the JSON in the body and overwrite 'tom' with it
			dProveInput := json.NewDecoder(r.Body)
			pProveInput := &ProveFPLInput{}
			err := dProveInput.Decode(pProveInput)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			proof := proveFPL(pProveInput)
			fmt.Printf("%+v", proof)
			//pHash := &Hash{
			//	Hash: hex.EncodeToString(FPLhash),
			//}
			//fmt.Printf("\n")
			//j, err := json.Marshal(pHash)
			//w.Write(j)
			break
		case "ack":
			break
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not supported")
	}
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/keys", keyHandler)
	rtr.HandleFunc("/hash/{type:.*}", hashHandler)
	rtr.HandleFunc("/sign", signHandler)
	rtr.HandleFunc("/prove/{type:.*}", proveHandler)
	http.Handle("/", rtr)
	//rtr.HandleFunc("/number/{id:[0-9]+}", pageHandler)

	log.Println("Go!")
	http.ListenAndServe(":3000", nil)
}
