package reindex

import (
	"encoding/json"
	"fmt"
	"log"
)

// createScroll initiates a Scan and Scroll operation with Elasticsearch
// and returns the Scroll ID that will be used to pull data out.
func createScroll(server, index, timeout string) string {
	url := fmt.Sprintf("%s/%s/_search?search_type=scan&scroll=%s", server, index, timeout)
	body, err := execJSONHTTPReq("GET", url, []byte(""))
	if err != nil {
		log.Fatal("Failed to intiate Scan and Scroll operation: " + err.Error())
	}

	var scrollRes ScrollResult
	err = json.Unmarshal(body, &scrollRes)
	if err != nil {
		log.Fatal("Failed to parse HTTP response while intiating Scan and " +
			"Scroll operation: " + err.Error())
	}

	return scrollRes.ScrollId
}

// fetchScrollPage executes a Scan and Scroll request with the given params.
// It returns a fully populated ScrollResult.
func fetchScrollPage(server, timeout, scrollID string) (*ScrollResult, error) {
	url := fmt.Sprintf("%s/_search/scroll?scroll=%s", server, timeout)
	body, err := execJSONHTTPReq("POST", url, []byte(scrollID))
	if err != nil {
		return nil, err
	}

	var scrollRes ScrollResult
	err = json.Unmarshal(body, &scrollRes)
	if err != nil {
		return nil, err
	}

	return &scrollRes
}
