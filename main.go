package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
    r := gin.Default()

    r.POST("/sign", func(ctx *gin.Context) {
        body, _ := ioutil.ReadAll(ctx.Request.Body)
        chekcsum := sha256.New()
        chekcsum.Write(body)
        bodySHA256 := chekcsum.Sum(nil)
        digest := fmt.Sprintf("%x", bodySHA256)
        privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
        claims := SignClaims{digest, 	jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
            Issuer:    "wesley",
        }}
        token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims) 
        s,_ := token.SignedString(privateKey)
        ctx.IndentedJSON(http.StatusOK, SignedResponse{Token: s}) 
    })

    r.Run()
}

type SignedResponse struct {
    Token string `json:"token"`
}

type SignClaims struct {
    Digest string `json:"digest"`
    jwt.RegisteredClaims
}
