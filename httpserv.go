package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

// benchmarks: ab -r -k -l -n 30000 -c 5000 [uri]
func root(w http.ResponseWriter, r *http.Request) {
	// request uri = "http://localhost/live/luztv-livestream.w8889.m3u8?id=0x449484abb&wid=0xbc677870"
	// r.URL.Path[1:] = "live/luztv-livestream.w8889.m3u8" <=> r.URL.RawQuery = "id=0x449484abb&wid=0xbc677870"
	mypath := r.URL.Path[1:] // live/luztv-livestream.m3u8

	file := rootdir + mypath
	fileinfo, err := os.Stat(file)
	if err != nil { // does not exist the path (file nor dir)
		http.NotFound(w, r)
		return
	} else if fileinfo.IsDir() { // it is a dir
		if strings.HasSuffix(file, "/") { // add /index.html to the end
			file = file + first_page + ".html"
		} else {
			file = file + "/" + first_page + ".html"
		}
	}
	fr, err := os.Open(file)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	defer fr.Close()
	if session {
		if strings.Contains(r.URL.String(), "?err") {
			// replace <span id="loginerr"></span> with an error text to show
			buf, _ := ioutil.ReadAll(fr)
			html := string(buf)
			html = strings.Replace(html, spanHTMLlogerr, ErrorText, -1)
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Content-Type", mime.TypeByExtension(".html"))
			fmt.Fprint(w, html)
		} else {
			// Get the cookies
			if (path.Ext(file) != ".html") || (file == (rootdir + first_page + ".html")) {
				w.Header().Set("Cache-Control", "no-cache")
				http.ServeContent(w, r, file, fileinfo.ModTime(), fr)
			} else {
				cookie, err := r.Cookie(CookieName)
				if err != nil {
					Warning.Println("Cookie not found in the browser")
					http.Redirect(w, r, "/"+first_page+".html", http.StatusFound)
				} else {
					key := cookie.Value
					mu_user.RLock()
					_, ok := user_[key]
					mu_user.RUnlock()
					if ok {
						cookie.Expires = time.Now().Add(time.Duration(session_timeout) * time.Second)
						http.SetCookie(w, cookie)
						mu_user.Lock()
						time_[cookie.Value] = cookie.Expires
						mu_user.Unlock()
						buf, _ := ioutil.ReadAll(fr)
						html := string(buf)
						// apply template
						tmpl := template.Must(template.New("template").Funcs(template.FuncMap{
							"eq": func(a, b string) bool {
								return a == b
							}}).Parse(`{{define "T"}}` + html + `{{end}}`))
						tmpl.ExecuteTemplate(w, "T", settings)
					} else {
						Warning.Println("Cookie not found in the server")
						http.Redirect(w, r, "/"+first_page+".html", http.StatusFound)
					}
				}
			}
		}
	} else {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeContent(w, r, file, fileinfo.ModTime(), fr)
	}

}
