package create

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"jAPI/config"
	"net/http"
	"os"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create Job",
	Long:  "Create Jenkins Job",
	Run: func(c *cobra.Command, args []string) {
		config := config.InitConfig()
		FILE := viper.GetString("xml_f")

		f, err := os.Open(FILE)
		if err != nil {
			fmt.Println(err)
			return
		}
		xmlData, err := io.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		fullUrl := fmt.Sprintf("http://%s:%s/createItem?name=%s", config.URL, config.PORT, config.JOB)
		fmt.Println(fullUrl)

		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPost, fullUrl, bytes.NewReader(xmlData))
		req.Header.Set("Content-Type", "application/xml")

		if err != nil {
			fmt.Println(err)
			return
		}
		auth := config.USER + ":" + config.TOKEN
		authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Add("Authorization", "Basic "+authEncoded)

		exists, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s:%s/job/%s/config.xml", config.URL, config.PORT, config.JOB), nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		exists.Header.Add("Authorization", "Basic "+authEncoded)
		ex, _ := client.Do(exists)
		if ex.StatusCode == http.StatusOK {
			fmt.Fprintf(os.Stdout, "%s exists\n", config.JOB)
			return
		}
		defer ex.Body.Close()

		data, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer data.Body.Close()

		if data.StatusCode != http.StatusOK {
			fmt.Println(data.Status)
			return
		} else {
			fmt.Fprintf(os.Stdout, "%s created...\n", config.JOB)
		}
	},
}

func init() {
	CreateCmd.Flags().StringP("xml_f", "f", "", "File xml for create Job")
	viper.BindPFlag("xml_f", CreateCmd.Flags().Lookup("xml_f"))
}
