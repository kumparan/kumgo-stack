package connector

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/kumparan/go-lib/logger"
	"github.com/kumparan/go-lib/utils"
	"github.com/kumparan/kumgo-stack/config"
)

// ArangoPost nodoc
func ArangoPost(query string) (body []byte, err error) {
	finalQuery := []byte(utils.StandardizeSpaces(query))

	req, err := http.NewRequest("POST", config.ArangoHost(), bytes.NewBuffer(finalQuery))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logger.Err("Http Post failed: ", err)
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Err("Query arango failed: ", err)
	}

	return
}
