package signer

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SignerClaims struct {
    Digest string `json:"digest"`
    jwt.RegisteredClaims
}

func Sign(payload []byte) string {
    chekcsum := sha256.New()
    chekcsum.Write(payload)
    bodySHA256 := chekcsum.Sum(nil)
    digest := fmt.Sprintf("%x", bodySHA256)
    privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
    claims := SignerClaims{
        digest,
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims) 
    s,_ := token.SignedString(privateKey)
    return s
}
