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
	// req.Header.Set("Content-Type", "application/json")
	// Disabling chunked requests is recommended by Elasticsearch, especially
	// for Bulk API requests. Chunked requests are disabled by setting the
	// correct Transfer-Encoding and Content-Length headers.
	// req.Header.Set("Transfer-Encoding", "identity")
	// if len(byteArrayPayload) > 0 {
	// 	req.Header.Set("Content-Length", string(len(byteArrayPayload)))
	// }

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	return resBody, err
}
