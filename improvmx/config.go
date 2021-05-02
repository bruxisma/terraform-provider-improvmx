package improvmx

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

type Config struct {
	Token   string
	BaseURL string
}
