package grifts

import (
	"github.com/dlsniper/buffalonew/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
