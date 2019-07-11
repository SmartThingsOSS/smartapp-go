package smartappcore

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
)

var _ = Describe("KeyGetter", func() {

	var (
		TestPemFile  string
		TestCertFile string
	)

	BeforeEach(func() {
		TestPemFile, _ = filepath.Abs("../../test/data/test.pem")
		TestCertFile, _ = filepath.Abs("../../test/data/test.crt")
	})

	Context("With a LocalKeyGetter", func() {
		It("should be able to create a local key getter", func() {
			keyGetter := NewLocalKeyGetter(&AuthConfig{
				PublicKeyBase64: ReadPublicKeyToBase64(TestPemFile),
			})
			Expect(keyGetter).To(Not(BeNil()))
			Expect(keyGetter.GetKey("")).To(BeAssignableToTypeOf(&rsa.PublicKey{}))
		})
	})

	Context("With a HttpKeyGetter", func() {
		It("should be able to create a http key getter", func() {
			keyId := "/pl/useast1/myKeyId"
			server := NewCertHttpServer(keyId, ReadCert(TestCertFile))
			defer server.Close()

			keyGetter := NewHttpKeyGetter(&AuthConfig{
				KeyApiHost: fmt.Sprintf("%s%%s", server.URL),
			})

			Expect(keyGetter).To(Not(BeNil()))
			Expect(keyGetter.GetKey(keyId)).To(BeAssignableToTypeOf(&rsa.PublicKey{}))
		})
	})

	Context("With a KeyGetter", func() {
		It("should be able to create a new key getter", func() {
			keyId := "/pl/useast1/myKeyId"
			server := NewCertHttpServer(keyId, ReadCert(TestCertFile))
			defer server.Close()

			keyGetter := NewKeyGetter(&AuthConfig{
				KeyApiHost:      fmt.Sprintf("%s%%s", server.URL),
				PublicKeyBase64: ReadPublicKeyToBase64(TestPemFile),
			})

			Expect(keyGetter).To(Not(BeNil()))
			Expect(keyGetter.GetKey(keyId)).To(BeAssignableToTypeOf(&rsa.PublicKey{}))
			Expect(keyGetter.GetKey("/SmartThings")).To(BeAssignableToTypeOf(&rsa.PublicKey{}))
			Expect(keyGetter.GetKey("/garbage")).To(BeNil())
		})
	})
})

func ReadCert(filename string) []byte {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("missing_test_fixture error=%v", err)
	}
	return file
}

func ReadPublicKeyToBase64(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("missing_test_fixture error=%v", err)
	}
	return base64.StdEncoding.EncodeToString(file)
}

func NewCertHttpServer(keyId string, cert []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.RequestURI != keyId {
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Type", "application/pkix-cert")
		_, _ = w.Write(cert)
	}))
}
