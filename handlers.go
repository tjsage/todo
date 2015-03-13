package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"net/http"
)

var (
	whiteList []string
	loginUrl  string = "/login"
)

func SetupLoginHandler(goodPaths []string) martini.Handler {
	whiteList := goodPaths

	return func(res http.ResponseWriter, r *http.Request, c martini.Context, session sessions.Session) {
		v := session.Get("user_id")

		if v == nil {
			for _, value := range whiteList {
				if value == r.URL.Path {
					return
				}
			}
			http.Redirect(res, r, loginUrl, http.StatusFound)
		}
	}
}
