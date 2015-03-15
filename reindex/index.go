package reindex

import "log"

// index takes a ScrollResult and generates & executes a bulk api indexing
// request for all hits in the ScrollResult.
func index(scrollRes *ScrollResult) error {
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

	// TODO: Execute the bulk api request
}

// generateBulkRequest maps a Hit into a slice of the meta portion as well as
// the source portion of a bulk api indexing request.
func generateBulkRequest(hit *Hit) ([]string, error) {
	meta := hit.GenerateBulkMeta()
	src, err := hit.GenerateBulkSource()
	if err != nil {
		return err
	}
	return []string{meta, src}, nil
}
