package smartappcore

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	SmartThingsKeyApiHost = "https://key.smartthings.com%s"
)

type KeyGetter interface {
	GetKey(id string) interface{}
}

type DefaultKeyGetter struct {
	LocalKeyGetter KeyGetter
	HttpKeyGetter  KeyGetter
}

func (g *DefaultKeyGetter) GetKey(id string) interface{} {
	if strings.HasPrefix(id, "/SmartThings") {
		return g.LocalKeyGetter.GetKey(id)
	}
	return g.HttpKeyGetter.GetKey(id)
}

func NewKeyGetter(config *AuthConfig) KeyGetter {
	return &DefaultKeyGetter{
		LocalKeyGetter: NewLocalKeyGetter(config),
		HttpKeyGetter:  NewHttpKeyGetter(config),
	}
}

type LocalKeyGetter struct {
	PublicKey interface{}
}

func (g *LocalKeyGetter) GetKey(id string) interface{} {
	return g.PublicKey
}

func NewLocalKeyGetter(config *AuthConfig) KeyGetter {
	return &LocalKeyGetter{
		PublicKey: ParsePublicKeyBase64(config.PublicKeyBase64),
	}
}

type HttpKeyGetter struct {
	Host   string
	Cache  KeyStore
	Client *http.Client
}

func (g *HttpKeyGetter) GetKey(id string) interface{} {
	key, ok := g.Cache.Get(id)
	if ok {
		return key
	}
	host := fmt.Sprintf(g.Host, id)

	resp, err := g.Client.Get(host)

	if err != nil {
		log.Printf("key_resolver.http_client_error keyId=%v, error=%v", id, err)
		return nil
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Printf("key_resolver.http_status_error keyId=%v status=%v", id, resp.StatusCode)
		return nil
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("key_resolver.io_read_error keyId=%v, error=%v", id, err)
		return nil
	}

	publicKey := ParsePublicKey(data)

	g.Cache.Put(id, publicKey)

	return publicKey
}

func NewHttpKeyGetter(config *AuthConfig) KeyGetter {
	tr := &http.Transport{}
	client := &http.Client{
		Timeout:   time.Second * 5,
		Transport: tr,
	}

	host := config.KeyApiHost
	if host == "" {
		host = SmartThingsKeyApiHost
	}

	return &HttpKeyGetter{
		Host:   host,
		Cache:  NewKeyStore(),
		Client: client,
	}
}

/**
 * Builds an rsa.PublicKey object for a pem formatted base64 encoded public key string.
 */
func ParsePublicKeyBase64(publicKeyBase64 string) interface{} {
	if publicKeyBase64 == "" {
		return nil
	}

	bytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		log.Printf("key_resolver.decode_base64_error error=%v", err)
		return nil
	}

	return ParsePublicKey(bytes)
}

func ParsePublicKey(bytes []byte) interface{} {
	block, _ := pem.Decode(bytes)

	if block == nil {
		log.Printf("key_resolver.decode_pem_error")
		return nil
	}

	// ParsePKIXPublicKey parses a DER encoded public key. These values are
	// typically found in PEM blocks with "BEGIN PUBLIC KEY".
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err == nil {
		return pub
	}

	// ParseCertificate parses a single certificate from the given ASN.1 DER data.
	cert, err := x509.ParseCertificate(block.Bytes)

	if err == nil {
		return cert.PublicKey
	}

	log.Printf("key_resolver.unsupported_key_format")
	return nil
}

type KeyStore interface {
	Get(key string) (interface{}, bool)
	Put(key string, publicKey interface{})
}

func NewKeyStore() KeyStore {
	return &DefaultKeyStore{
		Data: make(map[string]interface{}),
	}
}

type DefaultKeyStore struct {
	Data map[string]interface{}
	Lock sync.RWMutex
}

func (m *DefaultKeyStore) Get(key string) (interface{}, bool) {
	m.Lock.RLock()
	result, ok := m.Data[key]
	m.Lock.RUnlock()
	return result, ok
}

func (m *DefaultKeyStore) Put(key string, publicKey interface{}) {
	m.Lock.Lock()
	m.Data[key] = publicKey
	m.Lock.Unlock()
}
