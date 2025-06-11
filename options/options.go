package options

import (
	"fmt"
	"time"

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

// Credit options
func WithCreditID(credit_id string) func(*types.TMDBReqProps) {
	return func(props *types.TMDBReqProps) {
		props.URL += fmt.Sprintf("%s/%s", consts.CreditURL, credit_id)
	}
}

// Date Options
func WithEndDate(date time.Time) func(*types.TMDBReqProps) {
	return func(props *types.TMDBReqProps) {
		props.URL += fmt.Sprintf("%s%s", consts.EndDateURL, date.Format("2006-01-02"))
	}
}

func WithStartDate(date time.Time) func(*types.TMDBReqProps) {
	return func(props *types.TMDBReqProps) {
		props.URL += fmt.Sprintf("%s%s", consts.StartDateURL, date.Format("2006-01-02"))
	}
}
