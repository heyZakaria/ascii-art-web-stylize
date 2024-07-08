package web

import (
	"net/http"

	ascii_gen "web/packages/ascii_gen/src"
	data "web/packages/common_data"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data.Result.Text = r.FormValue("text")
		data.Result.Banner = r.FormValue("banner")
		path := "./packages/ascii_gen/banners/" + data.Result.Banner + ".txt"
		if data.Result.Text == "" || data.Result.Banner == "" || !ascii_gen.CheckBanner(path) || len(data.Result.Text) > 2000 {
			http.Error(w, "Bad Request Error", http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
