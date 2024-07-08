package web

import (
	"html/template"
	"net/http"

	ascii_gen "web/packages/ascii_gen/src"
	data "web/packages/common_data"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the URL is Correct if Not We return 404 Page Not Found Error
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Check if the Method is  if Not We return 405 Method not allowed
	if r.Method != "GET" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "500 :( Server Error", http.StatusInternalServerError)
		return
	}
	if data.Result.Banner != "" && data.Result.Text != "" {
		path := "packages/ascii_gen/banners/" + data.Result.Banner + ".txt"
		art, err := ascii_gen.AsciiGenerator(data.Result.Text, path)
		if err != nil {
			http.Error(w, "500 Internal Server error", http.StatusInternalServerError)
			return
		}
		data.Result.Ascii_art = art

		if data.Result.NotAllowedChars != nil {
			data.Result.Error = data.HandleErrorData(data.Result.NotAllowedChars)
		}
	}

	err = tmp.Execute(w, data.Result)
	if err != nil {
		http.Error(w, "500 :( Server Error", http.StatusInternalServerError)
		return
	}

	data.CleanData(&data.Result)
}
