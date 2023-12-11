// Package create provides a Cobra command for creating Jenkins Jobs.
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

// CreateJobCmd represents the Cobra command for creating Jenkins Jobs.
var CreateJobCmd = &cobra.Command{
	Use:   "create",
	Short: "create Job",
	Long:  "Create Jenkins Job",
	Run:   createCmdWrapper,
}

// init sets up flags for the CreateJobCmd command.
func init() {
	CreateJobCmd.Flags().StringP("xml_f", "f", "", "File xml for create Job")
	if err := viper.BindPFlag("xml_f", CreateJobCmd.Flags().Lookup("xml_f")); err != nil {
		log.Fatalf("Failed to bind flag: %v", err)
	}
}

// createCmdWrapper executes the creation of Jenkins Jobs.
func createCmdWrapper(cmd *cobra.Command, args []string) {
	err := runCreateJob(cmd, args)
	if err != nil {
		log.Println(err)
	}
}

// runCreateJob performs the creation of Jenkins Jobs.
func runCreateJob(_ *cobra.Command, args []string) error {
	cfg := config.InitConfig()
	if len(args) > 0 {
		cfg = config.UpdateConfigFromArgs(args)
	}

	fileContents := make([]string, len(cfg.JOB))
	fileContents, isFile, err := common.FileOrString(cfg)
	if isFile && err != nil {
		return err
	}
	if !isFile && err != nil {
		log.Printf("File not found, using content from command line: %s\n", cfg.JOB)
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
	return nil
}

// readXMLFile reads data from an XML file.
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

// createJob creates Jenkins Jobs.
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
