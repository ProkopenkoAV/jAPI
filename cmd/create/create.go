package create

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"jAPI/config"
	"log"
	"net/http"
	"os"
	"strings"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create Job",
	Long:  "Create Jenkins Job",
	Run:   runCreateCmd,
}

func init() {
	CreateCmd.Flags().StringP("xml_f", "f", "", "File xml for create Job")
	viper.BindPFlag("xml_f", CreateCmd.Flags().Lookup("xml_f"))
}

func runCreateCmd(_ *cobra.Command, _ []string) {
	cfg := config.InitConfig()

	fileContents := make([]string, 0)

	fileContents, isFile := fileOrString(cfg)

	if fileContents == nil {
		fmt.Fprintf(os.Stdout, "%s is empty or does not exist\n", cfg.JOB)
		return
	}

	if !isFile && len(fileContents) == 0 {
		fmt.Println("Flag mode")
		fileContents = trimString(cfg.JOB)
	}

	for _, jobName := range fileContents {
		xmlData, err := readXMLFile(viper.GetString("xml_f"))
		if err != nil {
			log.Println(err)
			continue
		}

		if jobExists(cfg, jobName) {
			fmt.Fprintf(os.Stdout, "%s already exists\n", jobName)
			continue
		}

		if err := createJob(cfg, xmlData, jobName); err != nil {
			log.Println(err)
			continue
		}

		fmt.Fprintf(os.Stdout, "%s created successfully\n", jobName)
	}

}

func readXMLFile(filename string) ([]byte, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	if fileInfo.IsDir() {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	xmlData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return xmlData, nil
}

func jobExists(cfg *config.Config, job string) bool {
	client := &http.Client{}

	existsURL := fmt.Sprintf("http://%s:%s/job/%s/config.xml", cfg.URL, cfg.PORT, job)
	req, err := http.NewRequest(http.MethodGet, existsURL, nil)
	if err != nil {
		log.Println(err)
		return false
	}
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(cfg.USER+":"+cfg.TOKEN)))

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func createJob(cfg *config.Config, xmlData []byte, job string) error {
	client := &http.Client{}

	fullURL := fmt.Sprintf("http://%s:%s/createItem?name=%s", cfg.URL, cfg.PORT, job)
	req, err := http.NewRequest(http.MethodPost, fullURL, bytes.NewReader(xmlData))
	req.Header.Set("Content-Type", "application/xml")
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(cfg.USER+":"+cfg.TOKEN)))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to create job: %s", resp.Status)
	}

	return nil
}

func fileOrString(cfg *config.Config) ([]string, bool) {
	fileInfo, err := os.Stat(cfg.JOB)
	if err != nil {
		log.Println(err)
		return []string{}, false
	}

	if fileInfo.IsDir() {
		return nil, false
	}

	file, err := os.Open(cfg.JOB)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return nil, false
	}

	fileLines := strings.Split(string(fileData), "\n")

	if len(fileLines) == 1 && fileLines[0] == "" {
		return nil, false
	}

	return fileLines, true
}

func trimString(str string) []string {
	return strings.Split(strings.TrimSpace(str), " ")
}
