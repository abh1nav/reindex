package reindex

// ReindexConf keeps the args around for easy access.
type ReindexConf struct {
	SrcServer     string
	SrcIndex      string
	DestServer    string
	DestIndex     string
	ScrollTimeout string
}

// reindexConf is where the config is stored.
var reindexConf *ReindexConf

// SetConf is called by main to save the config.
func SetConf(c *ReindexConf) {
	reindexConf = c
}

// GetConf is called by anyone who needs access to the config.
func GetConf() ReindexConf {
	return *reindexConf
}
