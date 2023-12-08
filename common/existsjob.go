package common

import (
	"encoding/base64"
	"fmt"
	"jAPI/config"
	"log"
	"net/http"
)

func JobExists(cfg *config.Config, job string) bool {
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
	defer func() {
		_ = resp.Body.Close()
	}()

	return resp.StatusCode == http.StatusOK
}
