package controllers

import (
	"net/http"

	"github.com/oduortoni/odrserver/src/server/view"
)

func Index(res http.ResponseWriter, req *http.Request) {
	view.Send(res, "./templates/index.html", nil)
}
