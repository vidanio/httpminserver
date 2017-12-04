package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/smtp"
	"net/url"
	"strings"
)

func contact(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // recupera campos del form tanto GET como POST r.Method(string)
	for k, v := range r.Form {
		fmt.Printf("%s -> %s\n", k, v)
	}
	fmt.Printf("IP: %s\t Referrer: %s\n", r.RemoteAddr, r.Referer())

	if r.FormValue("send") != "" {
		var ref *url.URL
		var err error

		if ref, err = url.Parse(r.Referer()); err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}
		if strings.Contains(ref.Hostname(), "fullmetalplayer") {
			// Connect to the remote SMTP server.
			c, err := smtp.Dial("localhost:25")
			if err != nil {
				http.Error(w, "Internal Server Error", 500)
				return
			}
			defer c.Close()
			// Set the sender and recipient.
			c.Mail("info@todostreaming.es")
			c.Rcpt(r.FormValue("email"))
			// Send the email body.
			wc, err := c.Data()
			if err != nil {
				http.Error(w, "Internal Server Error", 500)
				return
			}
			defer wc.Close()
			content := fmt.Sprintf("[Name]: %s\n[Company]: %s\n[IP]: %s\n[Message]: %s\n", r.FormValue("name"), r.FormValue("company"), r.RemoteAddr, r.FormValue("message"))
			buf := bytes.NewBufferString(content)
			if _, err = buf.WriteTo(wc); err != nil {
				http.Error(w, "Internal Server Error", 500)
				return
			}
			fmt.Println("Mail sent successfully:\n" + content)
		} else {
			http.Error(w, "Internal Server Error", 500)
			return
		}
	}

	// END OF FUNC
	http.Redirect(w, r, r.Referer(), http.StatusFound) // vuelve al formulario que la llamó
}
