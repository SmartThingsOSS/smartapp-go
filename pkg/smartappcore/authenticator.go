package smartappcore

import (
	"github.com/spacemonkeygo/httpsig"
	"net/http"
)

type AuthConfig struct {
	PublicKeyBase64 string `mapstructure:"publicKeyBase64"`
	KeyApiHost      string `mapstructure:"keyApiHost"`
}

type Authenticator interface {
	Verify(req *http.Request) error
}

type DefaultAuthenticator struct {
	Verifier *httpsig.Verifier
}

func (a *DefaultAuthenticator) Verify(req *http.Request) error {
	return a.Verifier.Verify(req)
}

func NewAuthenticator(config *AuthConfig) Authenticator {
	resolver := NewKeyGetter(config)
	verifier := httpsig.NewVerifier(resolver)
	return &DefaultAuthenticator{
		Verifier: verifier,
	}
}
