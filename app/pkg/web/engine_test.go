package web_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/getfider/fider/app/models"

	"github.com/getfider/fider/app/pkg/web"

	. "github.com/onsi/gomega"
)

func TestEngine_StartRequestStop(t *testing.T) {
	RegisterTestingT(t)
	w := web.New(&models.SystemSettings{})
	group := w.Group()
	{
		group.Get("/hello", func(c web.Context) error {
			return c.Ok(web.Map{})
		})
	}

	go w.Start(":8080")
	time.Sleep(time.Second)
	resp, err := http.Get("http://127.0.0.1:8080/hello")
	Expect(err).To(BeNil())
	resp.Body.Close()
	Expect(resp.StatusCode).To(Equal(http.StatusOK))

	resp, err = http.Get("http://127.0.0.1:8080/world")
	Expect(err).To(BeNil())
	resp.Body.Close()
	Expect(resp.StatusCode).To(Equal(http.StatusNotFound))

	err = w.Stop()
	Expect(err).To(BeNil())

	resp, err = http.Get("http://127.0.0.1:8080/hello")
	Expect(err).NotTo(BeNil())
	Expect(resp).To(BeNil())
}