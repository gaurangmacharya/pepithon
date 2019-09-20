# How to send an E-mail via GMail SMTP Server using GO

SMTP/NET Package implements the Simple Mail Transfer Protocol as defined in RFC 5321.

```
func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
```
Following is explanation about parameters needs to be passed.
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

**Google Account Security Settings**

Before sending Emails using Gmail SMTP Server, Change the required setting using Google Account Security Settings or [Click Here](https://myaccount.google.com/security)

1. Make sure that **2-Step-Verification** is Enabled

++Screen: Google Account Security Settings 2-Step-Verification is Disabled++
![Google Account Security Settings 2-Step-Verification is Disabled](https://i.imgur.com/6Hxmb2G.png)

++Screen: Google Account Security Settings 2-Step-Verification is Enabled++
![Google Account Security Settings 2-Step-Verification Enabled](https://i.imgur.com/vcQYoGo.png)


2. Turn ON/OFF to toggle the **Less Secure App Access** or [Click Here](https://myaccount.google.com/u/0/lesssecureapps)

![Less Secure App Access](https://i.imgur.com/mEGa22F.png)

3. Create APP Password for your application or device.

![Generate App Password](https://i.imgur.com/LHfCxdH.png)

2. For security precaution, Google may require you to complete this additional step while signing-in. [Click Here](https://accounts.google.com/DisplayUnlockCaptcha) to Allow access to your Google account using new Device/App.

![New Device-App](https://i.imgur.com/hymkYJ6.png)

==**Note**: It may take an hour or more to reflect this security changes==

 
Following are some of errors which you may get from Gmail SMTP Module

**Error 1**. If you have entered wrong credentials
```
2019/09/18 12:21:51 Error from SMTP Server: 535 5.7.8 Username and Password not accepted. Learn more at
5.7.8  https://support.google.com/mail/?p=BadCredentials c8sm5954072pfi.117 - gsmtp
```
**Error 2**. If you have not enabled App Password
```
2019/09/18 11:46:49 Error from SMTP Server: 534 5.7.9 Application-specific password required. Learn more at
5.7.9  https://support.google.com/mail/?p=InvalidSecondFactor s141sm5130851pfs.13 - gsmtp
```
**Error 3**. If you have entered wrong Email Address
```
2019/09/18 13:16:06 Error from SMTP Server: 553 5.1.2 The recipient address <recipient-email> is not a valid RFC-5321
5.1.2 address. w6sm8782758pfj.17 - gsmtp
```

**Sample Code [Email as HTML]**

``` Go
package main

import (
    "log"
    "net/smtp"
    "bytes"
    "fmt"
    "mime/quotedprintable"
)

func main() {
    to_email     := "recipient-email@domain"
    from_email   := "from-email@domain"
    password     := "gmail-app-password"

    header := make(map[string]string)
    header["From"]     = from_email
    header["To"]       = to_email
    header["Subject"]  = "Write Your Subject Here"

    header["MIME-Version"]        = "1.0"
    header["Content-Type"]        = fmt.Sprintf("%s; charset=\"utf-8\"", "text/html")
    header["Content-Disposition"] = "inline"
    header["Content-Transfer-Encoding"] = "quoted-printable"

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
    host := "smtp.gmail.com:587"
    auth := smtp.PlainAuth("", from_email, password, "smtp.gmail.com")
    status  := smtp.SendMail(host, auth, from_email, []string{to_email}, []byte(final_message))
    if status != nil {
        log.Printf("Error from SMTP Server: %s", status)
    }
    log.Print("Email Sent Successfully")
}
```

**Sample Code [Email as Plain Text]**

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
You can also try package named [Gomail](https://github.com/go-gomail/gomail) for sending mail via Gmail.
