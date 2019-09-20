package main
import (
    "log"
    "fmt"
    "bytes"
    "net/smtp"
    "mime/quotedprintable"
)

func main() {
    from_email := "from-email@domain"
    password   := "gmail-app-password"
    host       := "smtp.gmail.com:587"
    auth       := smtp.PlainAuth("", from_email, password, "smtp.gmail.com")

    header := make(map[string]string)
    to_email           := "recipient-email@domain"
    header["From"]     = from_email
    header["To"]       = to_email
    header["Subject"]  = "Write Your Subject Here"

    header["MIME-Version"]              = "1.0"
    header["Content-Type"]              = fmt.Sprintf("%s; charset=\"utf-8\"", "text/html")
    header["Content-Transfer-Encoding"] = "quoted-printable"
    header["Content-Disposition"]       = "inline"

    header_message := ""
    for key, value := range header {
        header_message += fmt.Sprintf("%s: %s\r\n", key, value)
    }

    body := "<h1>This is your HTML Body</h1>"
    var body_message bytes.Buffer
    temp := quotedprintable.NewWriter(&body_message)
    temp.Write([]byte(body))
    temp.Close()

    final_message := header_message + "\r\n" + body_message.String()
    status  := smtp.SendMail(host, auth, from_email, []string{to_email}, []byte(final_message))
    if status != nil {
        log.Printf("Error from SMTP Server: %s", status)
    }
    log.Print("Email Sent Successfully")
}
