package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/smtp"
	"net/url"
	"strings"
	"time"
)

func contact(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // recupera campos del form tanto GET como POST r.Method(string)

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
			c.Mail(r.FormValue("email"))
			c.Rcpt(recvmail)
			// Send the email body.
			wc, err := c.Data()
			if err != nil {
				http.Error(w, "Internal Server Error", 500)
				return
			}
			defer wc.Close()
			header := fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\nMime-Version: 1.0\r\nDate: %s\r\nContent-Type: text/plain; charset=UTF-8\r\nContent-Transfer-Encoding: quoted-printable\r\n\r\n",
				recvmail, r.FormValue("email"), subject, time.Now().Format(time.RFC1123Z))
			content := fmt.Sprintf("[Name]: %s\r\n[Company]: %s\r\n[IP]: %s\r\n[Message]: %s\r\n", r.FormValue("name"), r.FormValue("company"), r.RemoteAddr, r.FormValue("message"))
			buf := bytes.NewBufferString(header + content)
			if _, err = buf.WriteTo(wc); err != nil {
				http.Error(w, "Internal Server Error", 500)
				return
			}
			fmt.Println("Mail sent successfully:\n" + header + content)
			http.Redirect(w, r, "/"+sent_page+".html", http.StatusFound)
		} else {
			http.Error(w, "Internal Server Error", 500)
			return
		}
	}

	// END OF FUNC
	http.Redirect(w, r, r.Referer(), http.StatusFound) // vuelve al formulario que la llamó
}
