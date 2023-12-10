package create

import (
	"encoding/base64"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"io/ioutil"
	"jAPI/config"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestRunCreateCmd(t *testing.T) {
	xmlData := []byte("<xml>...</xml>")
	tmpfile, err := ioutil.TempFile("", "test.xml")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()
	tmpfile.Write(xmlData)

	cmd := &cobra.Command{}
	viper.Set("xml_f", tmpfile.Name())

	r := mux.NewRouter()
	r.HandleFunc("/createItem", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "Basic "+base64.StdEncoding.EncodeToString([]byte("fakeUser:fakeToken")), r.Header.Get("Authorization"))
		body, _ := ioutil.ReadAll(r.Body)
		assert.Equal(t, xmlData, body)
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	server := httptest.NewServer(r)
	defer server.Close()

	cfg := config.InitConfig()
	cfg.URL = server.URL[7:16]
	cfg.PORT = server.URL[17:]
	cfg.USER = "fakeUser"
	cfg.TOKEN = "fakeToken"
	cfg.JOB = "fakeJob"
	configValues := []string{cfg.URL, cfg.PORT, cfg.USER, cfg.TOKEN, cfg.JOB}

	err = runCreateJob(cmd, configValues)
	assert.NoError(t, err)
}

func TestReadXMLFile(t *testing.T) {
	xmlData := []byte("<xml>...</xml>")
	tmpfile, err := ioutil.TempFile("", "test.xml")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()
	tmpfile.Write(xmlData)

	data, err := readXMLFile(tmpfile.Name())
	assert.NoError(t, err)
	assert.Equal(t, xmlData, data)
}

func TestCreateJob(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/createItem", r.URL.Path)
		assert.Equal(t, "fakeJob", r.URL.Query().Get("name"))
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

	err := createJob(cfg, []byte("<xml>...</xml>"), "fakeJob")
	assert.NoError(t, err)
}
