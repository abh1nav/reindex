package reindex

import (
	"fmt"
	"log"
	"strings"
)

// index takes a ScrollResult and generates & executes a bulk api indexing
// request for all hits in the ScrollResult.
//
// Elasticsearch Bulk API docs:
// http://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
func Index(scrollRes *ScrollResult) error {
	var bulkReq []string
	for _, hit := range scrollRes.Hits.Hits {
		req, err := generateBulkRequest(&hit)
		if err != nil {
			// TODO: Push this into an error channel that logs to a file
			log.Println("Failed to generate bulk API indexing request for " +
				hit.ID + " - reason: " + err.Error())
			continue
		}
		bulkReq = append(bulkReq, req...)
	}

	reqBody := strings.Join(bulkReq, "\n")
	conf := GetConf()
	url := fmt.Sprintf("%s/_bulk", conf.DestServer)
	_, err := execJSONHTTPReq("POST", url, []byte(reqBody))
	if err != nil {
		// TODO: Push this into an error channel that logs to a file
		return err
	}

	// TODO: Go through the response body and check for errors
	return nil
}

// generateBulkRequest maps a Hit into a slice of the meta portion as well as
// the source portion of a bulk api indexing request.
func generateBulkRequest(hit *Hit) ([]string, error) {
	meta := hit.GenerateBulkMeta()
	src := hit.GenerateBulkSource()
	return []string{meta, src}, nil
}
