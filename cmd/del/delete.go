// Package del provides a Cobra command for deleting Jenkins Jobs.
package del

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"jAPI/common"
	"jAPI/config"
	"log"
	"net/http"
)

// DelJobCmd represents the Cobra command for deleting Jenkins Jobs.
var DelJobCmd = &cobra.Command{
	Use:   "delete",
	Short: "del Job",
	Long:  "Delete Jenkins Job",
	Run:   delCmdWrapper,
}

// delCmdWrapper executes the deletion of Jenkins Jobs.
func delCmdWrapper(cmd *cobra.Command, args []string) {
	err := runDelJob(cmd, args)
	if err != nil {
		log.Println(err)
	}
}

// runDelJob performs the deletion of Jenkins Jobs.
func runDelJob(_ *cobra.Command, args []string) error {
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

		if !common.JobExists(cfg, jobName) {
			log.Printf("%s don't exists\n", jobName)
			continue
		}

		if err := delJob(cfg, jobName); err != nil {
			log.Println(err)
			continue
		}
		log.Printf("%s deleted!\n", jobName)
	}
	return nil
}

// delJob deletes Jenkins Jobs.
func delJob(cfg *config.Config, job string) error {
	client := &http.Client{}

	fullURL := fmt.Sprintf("http://%s:%s/job/%s/doDelete", cfg.URL, cfg.PORT, job)
	req, err := http.NewRequest(http.MethodPost, fullURL, nil)
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
		return fmt.Errorf("failed to del job: %s", resp.Status)
	}

	return nil
}
