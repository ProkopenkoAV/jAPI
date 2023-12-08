package create

import (
	"encoding/base64"
	"fmt"
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
	cmd.Flags().Set("xml_f", tmpfile.Name())

	cfg := config.InitConfig()
	cfg.URL = "fakeURL"
	cfg.PORT = "fakePort"
	cfg.USER = "fakeUser"
	cfg.TOKEN = "fakeToken"
	cfg.JOB = "fakeJob"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/createItem?name=fakeJob", r.URL.Path)
		assert.Equal(t, "Basic "+base64.StdEncoding.EncodeToString([]byte("fakeUser:fakeToken")), r.Header.Get("Authorization"))
		body, _ := ioutil.ReadAll(r.Body)
		assert.Equal(t, xmlData, body)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg.URL = server.URL[7:]

	configValues := []string{cfg.URL, cfg.PORT, cfg.USER, cfg.TOKEN, cfg.JOB}
	err = runCreateCmd(cmd, configValues)
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
		URL:   server.URL[7:],
		PORT:  server.URL[17:],
		USER:  "fakeUser",
		TOKEN: "fakeToken",
		JOB:   "fakeJob",
	}

	fmt.Println(server.URL[7:])
	fmt.Println(server.URL[17:])

	cfg.URL = "127.0.0.1"

	err := createJob(cfg, []byte("<xml>...</xml>"), "fakeJob")
	assert.NoError(t, err)
}
