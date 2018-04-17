package story

import (
	"context"

	driver "github.com/arangodb/go-driver"
	"github.com/kumparan/go-lib/logger"
	"github.com/kumparan/kumgo-stack/buff"
	"github.com/kumparan/kumgo-stack/repository/connector"
)

func GetStories() (stories []buff.Story) {
	ctx := context.Background()
	query := "FOR d IN v1_stories LIMIT 10 RETURN d"
	cursor, err := connector.ArangoDB.Query(ctx, query, nil)
	if err != nil {
		logger.Err("Query error: ", err)
	}
	defer cursor.Close()
	for {
		var story buff.Story
		_, err := cursor.ReadDocument(ctx, &story)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			logger.Err(err)
		}
		stories = append(stories, story)
	}
	return
}
