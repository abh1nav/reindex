package reindex

import (
	"encoding/json"
	"fmt"
	"log"
)

// CreateScroll initiates a Scan and Scroll operation with Elasticsearch
// and returns the Scroll ID that will be used to pull data out.
func CreateScroll() string {
	conf := GetConf()
	url := fmt.Sprintf("%s/%s/_search?search_type=scan&scroll=%s",
		conf.SrcServer, conf.SrcIndex, conf.ScrollTimeout)
	reqBody := generateScanScrollBody()
	body, err := execJSONHTTPReq("GET", url, []byte(reqBody))
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

// FetchScrollPage executes a Scan and Scroll request with the given params.
// It returns a fully populated ScrollResult.
func FetchScrollPage(scrollID string) (*ScrollResult, error) {
	conf := GetConf()
	url := fmt.Sprintf("%s/_search/scroll?scroll=%s", conf.SrcServer, conf.ScrollTimeout)
	body, err := execJSONHTTPReq("POST", url, []byte(scrollID))
	if err != nil {
		return nil, err
	}

	var scrollRes ScrollResult
	err = json.Unmarshal(body, &scrollRes)
	if err != nil {
		return nil, err
	}

	return &scrollRes, nil
}
