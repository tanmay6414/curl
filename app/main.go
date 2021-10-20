package main

import (
    "fmt"
	  "log"
    "net"
    "net/mail"
	  "net/smtp"
    "crypto/tls"
	"time"
	"os"
	"net/http"
)

// SSL/TLS Email Example

func main() {

    // from := mail.Address{"", os.Getenv("SENDER_USERNAME")}
    // to   := mail.Address{"", os.Getenv("RECEIVER_MAIL")}
	from := mail.Address{"", os.Getenv("SENDER_USERNAME")}
    to   := mail.Address{"", os.Getenv("RECEIVER_MAIL")}
    subj := "DAILY BACKUP"
	currentTime := time.Now()
    body := "Backup process completed at " + currentTime.Format("2006.01.02 15:04:05") + ". And the Backup is successfully synced with S3"
    // Setup headers
    headers := make(map[string]string)
    headers["From"] = from.String()
    headers["To"] = to.String()
    headers["Subject"] = subj

    // Setup message
    message := ""
    for k,v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body

    // Connect to the SMTP Server
    servername := "smtp.gmail.com:465"

    host, _, _ := net.SplitHostPort(servername)

    auth := smtp.PlainAuth("",os.Getenv("SENDER_USERNAME"), os.Getenv("SENDER_PASSWORD"), host)

    // TLS config
    tlsconfig := &tls.Config {
        InsecureSkipVerify: true,
        ServerName: host,
    }
	log.Printf("Mail send Successfully at %s", currentTime.Format("2006.01.02 15:04:05"))
    // Here is the key, you need to call tls.Dial instead of smtp.Dial
    // for smtp servers running on 465 that require an ssl connection
    // from the very beginning (no starttls)
    conn, err := tls.Dial("tcp", servername, tlsconfig)
    if err != nil {
        log.Panic(err)
    }

    c, err := smtp.NewClient(conn, host)
    if err != nil {
        log.Panic(err)
    }

    // Auth
    if err = c.Auth(auth); err != nil {
        log.Panic(err)
    }

    // To && From
    if err = c.Mail(from.Address); err != nil {
        log.Panic(err)
    }

    if err = c.Rcpt(to.Address); err != nil {
        log.Panic(err)
    }

    // Data
    w, err := c.Data()
    if err != nil {
        log.Panic(err)
    }

    _, err = w.Write([]byte(message))
    if err != nil {
        log.Panic(err)
    }

    err = w.Close()
    if err != nil {
        log.Panic(err)
    }

    c.Quit()


	http.HandleFunc("/", helloHandler) // Update this line of code


    fmt.Printf("Starting server at port 80\n")
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal(err)
    }


}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

	currentTime := time.Now()
    fmt.Fprintf(w, "Mail send Successfully at %s", currentTime.Format("2006.01.02 15:04:05"))
}

