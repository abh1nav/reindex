package main

import (
	"flag"

	"github.com/abh1nav/reindex/reindex"
)

var (
	srcES     = flag.CommandLine.String("src-es", "", "URL for source Elasticsearch server. Example: http://es1.example.com:9200")
	srcIndex  = flag.CommandLine.String("src-index", "", "Source index name. Example: index1")
	destES    = flag.CommandLine.String("dest-es", "", "URL for destination Elasticsearch server. Example: https://es2.example.com:9200")
	destIndex = flag.CommandLine.String("dest-index", "", "Destination index name. Example: index2")
)

func main() {
	flag.Parse()
	c := &reindex.ReindexConf{
		SrcServer:  srcES,
		SrcIndex:   srcIndex,
		DestServer: destES,
		destIndex:  destIndex}
	reindex.SetConf(c)
}
