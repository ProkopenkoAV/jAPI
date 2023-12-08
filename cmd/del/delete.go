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

var DelJobCmd = &cobra.Command{
	Use:   "del",
	Short: "del Job",
	Long:  "Delete Jenkins Job",
	Run:   runDelCmd,
}

func runDelCmd(_ *cobra.Command, _ []string) {
	cfg := config.InitConfig()

	fileContents := make([]string, len(cfg.JOB))

	fileContents, isFile := common.FileOrString(cfg)

	if fileContents == nil {
		log.Printf("%s is empty or does not exist\n", cfg.JOB)
		return
	}

	if !isFile && len(fileContents) == 0 {
		fmt.Println("Flag mode")
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
		} else {
			log.Printf("%s deleted!\n", jobName)
		}
	}
}

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
