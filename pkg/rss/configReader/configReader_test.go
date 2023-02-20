package configReader

import (

	// "news-aggregator/pkg/rss/configReader"
	"testing"
)

func TestConfigReader_Read(t *testing.T) {
	config := Read("./../../../config.json")
	t.Log(config)
}
