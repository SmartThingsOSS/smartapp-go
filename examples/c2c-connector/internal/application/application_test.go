package application_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/SmartThingsOSS/smartapp-go/examples/c2c-connector/internal/application"
	"net/http"
	"net/http/httptest"
)

var (
	app     application.Application
	version = "test"
)

var _ = BeforeSuite(func() {
	app = application.GetInstance(version)
})

var _ = Describe("Application", func() {
	Context("After startup", func() {
		It("should have initialized a server", func() {
			Expect(app.GetConfig().GetString("server.port")).To(Equal("5555"))
			Expect(len(app.GetEngine().Routes())).To(Equal(3))
		})
	})

	Context("With a route of /buildinfo", func() {
		It("should successfully return a version", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/buildinfo", nil)
			app.GetEngine().ServeHTTP(w, req)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("{\"version\":\"" + version + "\"}"))
		})
	})

	Context("With a route of /health", func() {
		It("should successfully return a healthcheck", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/health", nil)
			app.GetEngine().ServeHTTP(w, req)

			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("{}"))
		})
	})
})
