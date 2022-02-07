package middleware

import (
	"myapp/data"

	"github.com/youngjae-lim/gosnel"
)

type Middleware struct {
	App *gosnel.Gosnel
	Models data.Models
}