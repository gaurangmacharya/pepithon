# How to send an E-mail via GMail SMTP Server using GO

**1. Introduction**
SMTP/NET Package implements the Simple Mail Transfer Protocol as defined in RFC 5321.
```
func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
```
**2. Parameters**
+ **addr**  is a Host Server Address along with Port Number separated by Column ':'
+ **a** is a Authentication response from Gmail
+ **from** is an Email Address using which we are authenticating and sending Email
+ **to** can a single Email Address or an array of Email Address to which we want to send an Email
+ **msg** 
  - This parameter should be an RFC 822-style with headers first, a blank line, and then the message body.
  - The content should be CRLF terminated.
  - The headers part includes fields such as "From", "To", "Subject", and "Cc".
  - Sending "Bcc" messages is accomplished by including an email address in the to parameter but not including it in the msg headers. 
  - This function and the net/smtp package are low-level mechanisms and do not provide support for DKIM signing, MIME attachments and  other features.

<br>

**3. Google Account Security Settings**

Before sending Emails using Gmail SMTP Server, Change the required setting using Google Account Security Settings or [Click Here](https://myaccount.google.com/security)

**3.1:** Make sure that 2-Step-Verification is Enabled

Screen: Google Account Security Settings 2-Step-Verification is Disabled

![Google Account Security Settings 2-Step-Verification is Disabled](https://i.imgur.com/6Hxmb2G.png)

Screen: Google Account Security Settings 2-Step-Verification is Enabled

![Google Account Security Settings 2-Step-Verification Enabled](https://i.imgur.com/vcQYoGo.png)

**3.2:** Turn ON/OFF to toggle the Less Secure App Access or [Click Here](https://myaccount.google.com/u/0/lesssecureapps)

![Less Secure App Access](https://i.imgur.com/mEGa22F.png)

**3.3:** Create APP Password for your application or device.

![Generate App Password](https://i.imgur.com/LHfCxdH.png)

**3.4:** For security precaution, Google may require you to complete this additional step while signing-in. [Click Here](https://accounts.google.com/DisplayUnlockCaptcha) to Allow access to your Google account using new Device/App.

![New Device-App](https://i.imgur.com/hymkYJ6.png)

==**Note**: It may take an hour or more to reflect this security changes==

<br>

**4. Errors**
Following are some of errors which you may encounter while testing Gmail SMTP Module

**4.1**. If you have entered wrong credentials
```
2019/09/18 12:21:51 Error from SMTP Server: 535 5.7.8 Username and Password not accepted. Learn more at
5.7.8  https://support.google.com/mail/?p=BadCredentials c8sm5954072pfi.117 - gsmtp
```
**4.2**. If you have not enabled App Password
```
2019/09/18 11:46:49 Error from SMTP Server: 534 5.7.9 Application-specific password required. Learn more at
5.7.9  https://support.google.com/mail/?p=InvalidSecondFactor s141sm5130851pfs.13 - gsmtp
```
**4.3**. If you have entered wrong Email Address
```
2019/09/18 13:16:06 Error from SMTP Server: 553 5.1.2 The recipient address <recipient-email> is not a valid RFC-5321
5.1.2 address. w6sm8782758pfj.17 - gsmtp
```
<br>

**5. Sample Code [Email as HTML]**
[Click here](https://github.com/gaurangmacharya/pepithon/blob/master/send-email-via-gmail-smtp-server-using-go.go) To download complete code.

**5.1:** Import required packages
- [log](https://golang.org/pkg/log/) :: log.Print() to print important stages and errors
- [fmt](https://golang.org/pkg/fmt/) :: fmt.Sprintf() To print formatted text
- [net/smpt](https://golang.org/pkg/net/smtp/) :: smtp.PlainAuth() is to authenticate account and smtp.SendMail() is to send Email using SMTP Protocol
- [mime/quotedprintable](https://golang.org/pkg/mime/quotedprintable/) :: quotedprintable.NewWriter() Is convert Email body in "Quoted Printable" Format. [Click here](https://en.wikipedia.org/wiki/Quoted-printable) To know more about the format.
``` Go
package main
import (
    "log"
    "fmt"
    "bytes"
    "net/smtp"
    "mime/quotedprintable"
)
```
**5.2:** Set required parameters to authenticating access to SMTP
``` go
from_email:= "from-email@domain"
password  := "gmail-app-password"
host      := "smtp.gmail.com:587"
auth      := smtp.PlainAuth("", from_email, password, "smtp.gmail.com")
```
**5.3:** Set required Email header parameters like From, To and Subject
``` go
header := make(map[string]string)
to_email        := "recipient-email@domain"
header["From"]   = from_email
header["To"]     = to_email
header["Subject"]= "Write Your Subject Here"
```
**5.4:** Set header parameters to define type of Email content.
``` go
header["MIME-Version"]              = "1.0"
header["Content-Type"]              = fmt.Sprintf("%s; charset=\"utf-8\"", "text/html")
header["Content-Disposition"]       = "inline"
header["Content-Transfer-Encoding"] = "quoted-printable"
```
**5.5:** Prepare Formatted header string by looping all Header parameters.
``` go
header_message := ""
for key, value := range header {
    header_message += fmt.Sprintf("%s: %s\r\n", key, value)
}
```
**5.6:** Prepare Quoted-Printable Email body. 
``` go
body := "<h1>This is your HTML Body</h1>"
var body_message bytes.Buffer
temp := quotedprintable.NewWriter(&body_message)
temp.Write([]byte(body))
temp.Close()
```
**5.7:** Prepare final Email message by concatenating header and body.
``` go
final_message := header_message + "\r\n" + body_message.String()
```
**5.8:** Send Email and print log accordingly
``` go
status  := smtp.SendMail(host, auth, from_email, []string{to_email}, []byte(final_message))
if status != nil {
    log.Printf("Error from SMTP Server: %s", status)
}
log.Print("Email Sent Successfully")
```
<br>

**6. Sample Code [Email as Plain Text]**
``` go
package main
import (
    "log"
    "net/smtp"
)
func main() {
    to_email     := "recipient-email@domain"
    from_email   := "from-email@domain"
    password     := "gmail-app-password"
    subject_body := "Subject: Write Your Subject\n\n" + "This is your Email Body"
    status       := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from_email, password, "smtp.gmail.com"), from_email, []string{to_email}, []byte(subject_body))
    if status != nil {
        log.Printf("Error from SMTP Server: %s", status)
    }
    log.Print("Email Sent Successfully")
}
```

#### You can also try package named [Gomail](https://github.com/go-gomail/gomail) for sending mail via Gmail.
