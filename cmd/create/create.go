package create

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"jAPI/common"
	"jAPI/config"
	"log"
	"net/http"
	"os"
)

var CreateJobCmd = &cobra.Command{
	Use:   "create",
	Short: "create Job",
	Long:  "Create Jenkins Job",
	Run:   runCreateCmd,
}

func init() {
	CreateJobCmd.Flags().StringP("xml_f", "f", "", "File xml for create Job")
	if err := viper.BindPFlag("xml_f", CreateJobCmd.Flags().Lookup("xml_f")); err != nil {
		log.Fatalf("Failed to bind flag: %v", err)
	}
}

func runCreateCmd(_ *cobra.Command, _ []string) {
	cfg := config.InitConfig()

	fileContents := make([]string, len(cfg.JOB))

	fileContents, isFile := common.FileOrString(cfg)

	if fileContents == nil {
		log.Printf("%s is empty or does not exist\n", cfg.JOB)
		return
	}

	if !isFile && len(fileContents) == 0 {
		fmt.Printf("")
		fileContents = common.TrimString(cfg.JOB)
	}

	for _, jobName := range fileContents {

		xmlData, err := readXMLFile(viper.GetString("xml_f"))
		if err != nil {
			log.Println(err)
			continue
		}

		if common.JobExists(cfg, jobName) {
			log.Printf("%s already exists\n", jobName)
			continue
		}

		if err := createJob(cfg, xmlData, jobName); err != nil {
			log.Println(err)
			continue
		}

		log.Printf("%s created!\n", jobName)
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
	defer func() {
		_ = file.Close()
	}()

	xmlData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return xmlData, nil
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
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create job: %s", resp.Status)
	}

	return nil
}
