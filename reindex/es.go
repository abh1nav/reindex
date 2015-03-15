package reindex

type ScrollResult struct {
	Hits     *Hits      `json:"hits"`
	ScrollId string     `json:"_scroll_id"`
	Shards   *ShardInfo `json:"_shards"`
	TimedOut bool       `json:"timed_out"`
	Took     uint32     `json:"took"`
}

type ShardInfo struct {
	Failed     uint8 `json:"failed"`
	Successful uint8 `json:"successful"`
	Total      uint8 `json:"total"`
}

type Hits struct {
	Hits []Hit `json:"hits"`
}

type Hit struct {
	Index  string     `json:"_index"`
	Type   string     `json:"_type"`
	ID     string     `json:"_id"`
	Source RawMessage `json:"_source"`
}
