# How to send an E-mail via GMail SMTP Server using GO

## Introduction

Using the Gmail's SMTP Server, you can send emails to any domain using your Gmail Credentials. Following are some email sending limit criteria:
+ Google limits the number of recipients in a single email and the number of emails that can be sent per day.
+ The current limit is 500 Emails in a day or 500 recipients in a single email.
+ On reaching threshold limits, you won't be able to send messages for the next 24 hours.
+ After the suspension period, the counter gets reset automatically, and the user can resume sending Emails.
+ For more information about email sending limits refer the following links:
 + [Email sending limits](https://support.google.com/a/answer/166852 "Email sending limits")
 + [Error messages once limit is crossed](https://support.google.com/mail/answer/22839 "Error messages once limit is crossed")

## Settings to be updated on Google

1. Before sending emails using the Gmail's SMTP Server, Change the required settings under your Google Account Security Settings or [Click here](https://myaccount.google.com/security "Click here").
![Google Account Security Settings](https://i.imgur.com/6Hxmb2G.png)

2. Make sure that 2-Step-Verification is disabled.
![2-Step Virification Disabled](https://i.imgur.com/6Hxmb2G.png)

3. Turn ON the "Less Secure App" access or Click [here](https://myaccount.google.com/u/0/lesssecureapps "here").
![Less Secure App Access](https://i.imgur.com/hymkYJ6.png)

4. If 2-step-verification is enabled, then you will have to create app password for your application or device.
![2-Step Virification Enabled](https://i.imgur.com/vcQYoGo.png)

5. For security measures, Google may require you to complete this additional step while signing-in. Click here to allow access to your Google account using the new device/app.
![New Device-App](https://i.imgur.com/mEGa22F.png)

*Note: It may take an hour or more to reflect any security changes*


## Writing the GO Code to Send Email using Gmail SMTP Server

SMTP/NET package implements the Simple Mail Transfer Protocol as defined in RFC 5321.

```
func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
```

**Parameters**

+ **addr**  is a Host Server Address along with Port Number separated by Column ':'
+ **a** is an Authentication response from Gmail
+ **from** is an Email Address using which we are authenticating and sending Email
+ **to** can a single Email Address or an array of Email Address to which we want to send an Email
+ **msg** 
  - This parameter should be an RFC 822-style with headers first, a blank line, and then the message body.
  - The content should be CRLF terminated.
  - The headers part includes fields such as "From", "To", "Subject", and "Cc".
  - Sending "Bcc" messages is accomplished by including an email address in the to parameter but not including it in the msg headers. 
  - This function and the net/SMTP package are low-level mechanisms and do not provide support for DKIM signing, MIME attachments and other features.

<br>
### GO Code for sending Email as HTML

[Click here](https://github.com/gaurangmacharya/pepithon/blob/master/send-email-via-gmail-smtp-server-using-go.go) to download the complete working code for sending email as HTML.

**Explaining the working code:**

**Step 1:** Import required packages
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
**Step 2:** Set required parameters to authenticating access to SMTP
``` go
from_email:= "from-email@domain"
password  := "gmail-app-password"
host      := "smtp.gmail.com:587"
auth      := smtp.PlainAuth("", from_email, password, "smtp.gmail.com")
```

**Step 3:** Set required Email header parameters like From, To and Subject
``` go
header := make(map[string]string)
to_email        := "recipient-email@domain"
header["From"]   = from_email
header["To"]     = to_email
header["Subject"]= "Write Your Subject Here"
```

**Step 4:** Set header parameters to define type of Email content.
``` go
header["MIME-Version"]              = "1.0"
header["Content-Type"]              = fmt.Sprintf("%s; charset=\"utf-8\"", "text/html")
header["Content-Disposition"]       = "inline"
header["Content-Transfer-Encoding"] = "quoted-printable"
```

**Step 5:** Prepare Formatted header string by looping all Header parameters.
``` go
header_message := ""
for key, value := range header {
    header_message += fmt.Sprintf("%s: %s\r\n", key, value)
}
```

**Step 6:** Prepare Quoted-Printable Email body. 
``` go
body := "<h1>This is your HTML Body</h1>"
var body_message bytes.Buffer
temp := quotedprintable.NewWriter(&body_message)
temp.Write([]byte(body))
temp.Close()
```

**Step 7:** Prepare final Email message by concatenating header and body.
``` go
final_message := header_message + "\r\n" + body_message.String()
```

**Step 8:** Send Email and print log accordingly
``` go
status  := smtp.SendMail(host, auth, from_email, []string{to_email}, []byte(final_message))
if status != nil {
    log.Printf("Error from SMTP Server: %s", status)
}
log.Print("Email Sent Successfully")
```
<br>
### GO Code for sending Email as TEXT
In case, you just want to send a TEXT based email then use the below code:

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


## List of Possible Errors And Exceptions

Following are some of the errors which you may encounter while testing Gmail SMTP Module:

**Error 1:**. If you have entered wrong credentials
```
2019/09/18 12:21:51 Error from SMTP Server: 535 5.7.8 Username and Password not accepted. Learn more at
5.7.8  https://support.google.com/mail/?p=BadCredentials c8sm5954072pfi.117 - gsmtp
```
**Error 2:**. If you have not enabled App Password
```
2019/09/18 11:46:49 Error from SMTP Server: 534 5.7.9 Application-specific password required. Learn more at
5.7.9  https://support.google.com/mail/?p=InvalidSecondFactor s141sm5130851pfs.13 - gsmtp
```
**Error 3:**. If you have entered wrong Email Address
```
2019/09/18 13:16:06 Error from SMTP Server: 553 5.1.2 The recipient address <recipient-email> is not a valid RFC-5321
5.1.2 address. w6sm8782758pfj.17 - gsmtp
```

#### You can also try package named [Gomail](https://github.com/go-gomail/gomail) for sending mail via Gmail.

## Conclusion
Hope the steps explained above were useful and you were able to successfully send mail from your Gmail SMTP server using GO. Feel free to contribute, in case you encountered some issue which is not listed as a part of this tutorials. Use below comments section to ask/share any feedback.

<-- Happy Coding -->
