package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (app *application) payloadSign(w http.ResponseWriter, r *http.Request) {
    body, _ := ioutil.ReadAll(r.Body)
    chekcsum := sha256.New()
    chekcsum.Write(body)
    bodySHA256 := chekcsum.Sum(nil)
    digest := fmt.Sprintf("%x", bodySHA256)
    privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
    claims := jwt.RegisteredClaims{
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
        IssuedAt:  jwt.NewNumericDate(time.Now()),
        NotBefore: jwt.NewNumericDate(time.Now()),
        Issuer:    "wesley",
    }
    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims) 
    s,_ := token.SignedString(privateKey)
    println(s)
    println(digest)
}

func ping(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
}
