package gotmdb

import (
	"go/types"
	"net/http"
)

type TMDbCLient struct {
	Config types.Config
	Client http.Client
}
