package controllers

import (
	"github.com/codegangsta/martini-contrib/render"
)

var (
	SIMPLE_LAYOUT = render.HTMLOptions{
		Layout: "simple",
	}
)
