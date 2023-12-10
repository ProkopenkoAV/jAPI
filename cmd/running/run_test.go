package running

import (
	"encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"jAPI/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO : fix this test
func TestRunRunningJob(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/job/fakeJob/config.xml", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/job/fakeJob/config.xml", r.URL.Path)
		//fmt.Println(r.URL.)
		assert.Equal(t, "Basic "+base64.StdEncoding.EncodeToString([]byte("fakeUser:fakeToken")), r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")
	r.HandleFunc("/job/fakeJob/build", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "Basic "+base64.StdEncoding.EncodeToString([]byte("fakeUser:fakeToken")), r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusCreated)
	}).Methods("POST")

	server := httptest.NewServer(r)
	defer server.Close()

	cmd := &cobra.Command{}
	cfg := config.InitConfig()
	cfg.URL = server.URL[7:16]
	cfg.PORT = server.URL[17:]
	cfg.USER = "fakeUser"
	cfg.TOKEN = "fakeToken"
	cfg.JOB = "fakeJob"
	configValues := []string{cfg.URL, cfg.PORT, cfg.USER, cfg.TOKEN, cfg.JOB}

	err := runRunningJob(cmd, configValues)
	assert.NoError(t, err)
}

func TestRunJob(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		fmt.Println(r.URL.Path)
		assert.Equal(t, "/job/fakeJob/build", r.URL.Path)
		assert.Equal(t, "Basic "+base64.StdEncoding.EncodeToString([]byte("fakeUser:fakeToken")), r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	cfg := &config.Config{
		URL:   server.URL[7:16],
		PORT:  server.URL[17:],
		USER:  "fakeUser",
		TOKEN: "fakeToken",
		JOB:   "fakeJob",
	}

	err := runJob(cfg, cfg.JOB)
	assert.NoError(t, err)
}
