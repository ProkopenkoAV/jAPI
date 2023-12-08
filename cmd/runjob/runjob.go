package runjob

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"jAPI/cmd"
	"net/http"
)

var runJobCmd = &cobra.Command{
	Use:   "runjob",
	Short: "run Job",
	Long:  "Running Jenkins Job",
	Run: func(cmd *cobra.Command, args []string) {
		URL := viper.GetString("url")
		PORT := viper.GetString("port")
		USER := viper.GetString("user")
		TOKEN := viper.GetString("token")

		job, _ := cmd.Flags().GetString("job")
		if job == "" {
			fmt.Println("Job name is empty")
			return
		}
		fullUrl := fmt.Sprintf("http://%s:%s/job/%s/build", URL, PORT, job)

		client := &http.Client{}
		req, err := http.NewRequest("POST", fullUrl, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		auth := USER + ":" + TOKEN
		authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Add("Authorization", "Basic "+authEncoded)

		data, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer data.Body.Close()

		if data.StatusCode != http.StatusOK {
			fmt.Println(data.Status)
			return
		}
		responseData, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(responseData))
	},
}

func init() {
	cmd.RootCmd.AddCommand(runJobCmd)
	runJobCmd.Flags().StringP("job", "j", "", "Job name")
}
