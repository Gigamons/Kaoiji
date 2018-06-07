package api

import (
	"net/http"

	"github.com/Gigamons/Kaoiji/global"
)

func SendUserstats(w http.ResponseWriter, r *http.Request) {
	if r.ParseForm() != nil {
		finish(w, 500, "SERVERSIDE")
		return
	}
	if r.FormValue("APIKey") != global.CONFIG.API.APIKey {
		finish(w, 403, "APIKEY")
		return
	}
	finish(w, 200, "OK")
}
