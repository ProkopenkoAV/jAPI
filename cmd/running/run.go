package running

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"jAPI/config"
	"net/http"
	"os"
)

var RunJobCmd = &cobra.Command{
	Use:   "run",
	Short: "running Job",
	Long:  "Running Jenkins Job",
	Run: func(c *cobra.Command, args []string) {
		config := config.InitConfig()

		fullUrl := fmt.Sprintf("http://%s:%s/job/%s/build", config.URL, config.PORT, config.JOB)
		fmt.Println(fullUrl)

		client := &http.Client{}
		req, err := http.NewRequest("POST", fullUrl, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		auth := config.USER + ":" + config.TOKEN
		authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Add("Authorization", "Basic "+authEncoded)

		exists, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s:%s/job/%s/config.xml", config.URL, config.PORT, config.JOB), nil)
		exists.Header.Add("Authorization", "Basic "+authEncoded)
		ex, _ := client.Do(exists)
		if ex.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stdout, "%s not exists\n", config.JOB)
			return
		}
		defer ex.Body.Close()

		data, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer data.Body.Close()

		if data.StatusCode != http.StatusCreated {
			fmt.Println(data.Status)
			return
		} else {
			fmt.Fprintf(os.Stdout, "%s running...\n", config.JOB)
		}
	},
}
