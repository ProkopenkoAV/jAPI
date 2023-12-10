package del

import (
	"encoding/base64"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"jAPI/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO : fix this test
func TestRunDelJob(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/job/fakeJob/config.xml", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/job/fakeJob/config.xml", r.URL.Path)
		assert.Equal(t, "Basic "+base64.StdEncoding.EncodeToString([]byte("fakeUser:fakeToken")), r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")
	r.HandleFunc("/job/fakeJob/doDelete", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "Basic "+base64.StdEncoding.EncodeToString([]byte("fakeUser:fakeToken")), r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
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

	err := runDelJob(cmd, configValues)
	assert.NoError(t, err)
}

func TestDelJob(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/job/fakeJob/doDelete", r.URL.Path)
		assert.Equal(t, "Basic "+base64.StdEncoding.EncodeToString([]byte("fakeUser:fakeToken")), r.Header.Get("Authorization"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := &config.Config{
		URL:   server.URL[7:16],
		PORT:  server.URL[17:],
		USER:  "fakeUser",
		TOKEN: "fakeToken",
		JOB:   "fakeJob",
	}

	err := delJob(cfg, cfg.JOB)
	assert.NoError(t, err)
}
