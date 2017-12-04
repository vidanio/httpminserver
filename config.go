package main

const (
	// config variables for the HTTP server
	MAX_MEMORY             = 16 * 1024 * 1024                             // 16 MB
	rootdir                = "/usr/local/bin/html/"                       // website root ??????????????????
	session           bool = false                                        // session control by cookies enabled
	session_timeout        = 1200                                         // timeout for a session (20 min?)
	first_page             = "index"                                      // login page (default root page - always .html)
	enter_page             = "enter.html"                                 // enter page after login
	sent_page              = "sent"                                       // sent mail page
	http_port              = "80"                                         // HTTP server port ??????????????????? "80"
	name_username          = "usr"                                        // name of input username in the login page
	name_password          = "pwd"                                        // name of input password in the login page
	CookieName             = "GOSESSID"                                   // cookie name used for user/admin sessions
	login_cgi              = "/login.cgi"                                 // action cgi login in login page
	logout_cgi             = "/logout.cgi"                                // logout link at any page
	session_value_len      = 26                                           // len of the value of the session cookie
	spanHTMLlogerr         = "<span id='loginerr'></span>"                // <span> where to publish the login error
	ErrorText              = "Login Error"                                // message to show when error login
	logFile                = "/usr/local/bin/html/logs/httpminserver.log" // error logs file path ???????????????? "/var/log/hlserver.log"
	settingsFile           = "/usr/local/bin/settings.reg"                // file with some settings inside ?????????????? "/usr/local/bin/settings.reg"
	metalFile              = "/tmp/httpminserver"                         // file to know if system rebooted or not when systemd falls
	recvmail               = "sales@streamrus.com"
	subject                = "Full Metal Player's Contact Form"
)
