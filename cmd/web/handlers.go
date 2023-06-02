package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wesleynepo/payload-signer-go/internal/signer"
)

type envelope = map[string]interface{}

func (app *application) payloadSign(w http.ResponseWriter, r *http.Request) {
    body, _ := ioutil.ReadAll(r.Body)
    sign := signer.Sign(body)
    data := envelope{"token": sign}
    js, err := json.MarshalIndent(data, "", "\t")

    if err != nil {
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    if _, err := w.Write(js); err != nil {
    }
}

func ping(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
}
