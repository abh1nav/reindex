package reindex

import (
	"encoding/json"
	"fmt"
)

// generateScanScrollBody returns the query that's executed on the
// intial Scan Scroll setup request.
func generateScanScrollBody() string {
	return `{
	    "query": {"match_all": {}},
	    "size":  5
	}`
}

// ScrollResult represents the Elasticsearch response returned when a
// Scan and Scroll request is created or a page is retrieved using a
// Scan and Scroll call.
type ScrollResult struct {
	Hits     *Hits      `json:"hits"`
	ScrollId string     `json:"_scroll_id"`
	Shards   *ShardInfo `json:"_shards"`
	TimedOut bool       `json:"timed_out"`
	Took     uint32     `json:"took"`
}

// ShardInfo indicates the shard-level success and failure information
// of the request.
type ShardInfo struct {
	Failed     uint8 `json:"failed"`
	Successful uint8 `json:"successful"`
	Total      uint8 `json:"total"`
}

// Hits is a list containing the actual list of hits returned for the request.
type Hits struct {
	Hits []Hit `json:"hits"`
}

// Hit represents a document in Elasticsearch.
type Hit struct {
	Index  string          `json:"_index"`
	Type   string          `json:"_type"`
	ID     string          `json:"_id"`
	Source json.RawMessage `json:"_source"`
}

// bulkMetaTemplate is the template used for the first line of a bulk api
// indexing request. Elasticsearch Bulk API docs:
// http://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
var bulkMetaTemplate = `{"index": {"_index": "%s", "_type": "%s", "_id": "%s"}}`

// GenerateBulkMeta generates the meta line for the bulk api indexing request
// from the given hit by extracting the index, type and ID info.
func (hit *Hit) GenerateBulkMeta() string {
	return fmt.Sprintf(bulkMetaTemplate, hit.Index, hit.Type, hit.ID)
}

// GenerateBulkSource serializes the source field into a JSON string for the
// second line of a bulk api indexing request.
func (hit *Hit) GenerateBulkSource() string {
	return string(hit.Source)
}
