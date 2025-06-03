package options

import (
	"fmt"

	"github.com/aamgayle/gotmdb/consts"
	"github.com/aamgayle/gotmdb/types"
)

func WithPages(page int) func(*types.TMDBReqProps) {
	if page <= 1 {
		return func(props *types.TMDBReqProps) {
			props.URL += fmt.Sprintf("%s%d", consts.PageURL, 1)
		}
	}
	return func(props *types.TMDBReqProps) {
		props.URL += fmt.Sprintf("%s%d", consts.PageURL, page)
	}
}

func WithBaseURL(url string) func(*types.TMDBReqProps) {
	return func(props *types.TMDBReqProps) {
		props.BaseURL = url
	}
}

func WithKeyword(keyword string) func(*types.TMDBReqProps) {
	return func(props *types.TMDBReqProps) {
		props.URL += fmt.Sprintf("%s%s%s", consts.KeywordURL, consts.QueryURL, keyword)
	}
}
