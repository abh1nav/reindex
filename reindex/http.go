package reindex

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

var client = &http.Client{}

// execJSONHTTPReq executes an HTTP request that sends and/or receives JSON.
func execJSONHTTPReq(httpVerb, url string, byteArrayPayload []byte) ([]byte, error) {
	req, err := http.NewRequest(httpVerb, url, bytes.NewReader(byteArrayPayload))
	if err != nil {
		return nil, err
	}
	req.Close = true
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	return resBody, err
}
