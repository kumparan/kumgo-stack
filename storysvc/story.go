package storysvc

import (
	"context"

	"github.com/kumparan/kumgo-stack/buff"
	"github.com/kumparan/kumgo-stack/repository/story"
)

// GetStories :nodoc:
func (m *svc) GetStories(c context.Context, e *buff.Empty) (*buff.GetStoriesResponse, error) {
	var storiesResp []*buff.Story
	stories := story.GetStories()

	for _, story := range stories {
		storiesResp = append(storiesResp, &buff.Story{
			Author: story.Author,
			Editor: story.Editor,
		})
	}
	return &buff.GetStoriesResponse{Stories: storiesResp}, nil
}
