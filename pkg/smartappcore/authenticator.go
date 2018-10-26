package smartappcore

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/spacemonkeygo/httpsig"
	"log"
)

type KeyResolver struct {
	PublicKey *rsa.PublicKey
}

type Authenticator struct {
	Verifier *httpsig.Verifier
}

func NewAuthenticator(publicKey *rsa.PublicKey) *Authenticator {
	resolver := &KeyResolver{
		PublicKey: publicKey,
	}
	verifier := httpsig.NewVerifier(resolver)
	return &Authenticator{
		Verifier: verifier,
	}
}

func (k KeyResolver) GetKey(id string) interface{} {
	return k.PublicKey
}

/**
 * Builds an rsa.PublicKey object for a pem formatted base64 encoded public key string.
 */
func ParsePublicKeyBase64(publicKeyBase64 string) *rsa.PublicKey {
	bytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		log.Panic("Failed to parse public key", err)
	}
	return parsePublicKey(bytes)
}

func parsePublicKey(bytes []byte) *rsa.PublicKey {
	block, _ := pem.Decode(bytes)

	if block == nil {
		log.Panic("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		log.Panic("Failed to parse public key", err)
	}

	return pub.(*rsa.PublicKey)
}
