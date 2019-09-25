# How to send an E-mail via GMail SMTP Server using PHP

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

## Writing the PHP Code to Send Email using Gmail SMTP

**Step 1:** Download PHP Library
+ [Click here](https://github.com/PHPMailer/PHPMailer/) for more details about PHPMailer
+ [Click here](https://github.com/PHPMailer/PHPMailer/archive/master.zip) to download .zip file
+ Unzip the master.zip in your application directory and run following command from your application directory.

```
composer require phpmailer/phpmailer
```
[Composer](https://getcomposer.org/) is the recommended way to install PHPMailer.

**Step 2:** Writing the PHP Code to make an SMTP connection
+ Using your Gmail credentials, connect to host **"smtp.gmail.com"**
  - Use port 465, if want SSL
  - Use port 587, if want TLS
+ [Click here](https://github.com/PHPMailer/PHPMailer/wiki/Tutorial) for some more Examples and Tutorials of PHPMailer

**Step 3:** Include packages and files for PHPMailer and SMTP protocol:
``` php
use PHPMailer\PHPMailer\PHPMailer;
use PHPMailer\PHPMailer\Exception;
require 'PHPMailer-master/src/Exception.php';
require 'PHPMailer-master/src/PHPMailer.php';
require 'PHPMailer-master/src/SMTP.php';
```

**Step 4:** Initialize PHP Mailer and set SMTP as mailing protocol:
``` php
$mail = new PHPMailer();
$mail->IsSMTP();
$mail->Mailer = "smtp";
```

**Step 5:** Set required parameters for making an SMTP connection like server, port and account credentials. SSL and TLS are both cryptographic protocols that provide authentication and data encryption between servers, machines and applications operating over a network. SSL is the predecessor to TLS.

``` php
$mail->SMTPDebug  = 1;  
$mail->SMTPAuth   = TRUE;
$mail->SMTPSecure = "tls";
$mail->Port       = 587;
$mail->Host       = "smtp.gmail.com";
$mail->Username   = "your-email@gmail.com";
$mail->Password   = "your-gmail-password";
```
**Step 6:** Set required parameters for email header and body:
``` php
$mail->IsHTML(true);
$mail->AddAddress("recipient-email@domain", "recipient-name");
$mail->SetFrom("from-email@gmail.com", "from-name");
$mail->AddReplyTo("reply-to-email@domain", "reply-to-name");
$mail->AddCC("cc-recipient-email@domain", "cc-recipient-name");
$mail->Subject = "Test is Test Email sent via Gmail SMTP Server using PHP Mailer";
$content = "<b>This is a Test Email sent via Gmail SMTP Server using PHP mailer class.</b>";
```
**Step 7:** Send the email and catch required exceptions:
``` php
$mail->MsgHTML($content); 
if(!$mail->Send()) {
  echo "Error while sending Email.";
  var_dump($mail);
} else {
  echo "Email sent successfully";
}
```


## Working Code to Send Email via Gmail SMTP using PHP
[Click here](https://github.com/gaurangmacharya/pepithon/blob/master/send-email-via-gmail-smtp-server-using-phpmailer.php) to download the complete code.

## List of Possible Errors And Exceptions

**Error 1:** When SMTP Mail Server Credentials were correct but Application Specific Password was not provided
```
SMTP ERROR: Password command failed: 534-5.7.9 Application-specific password required. Learn more at 534 5.7.9  https://support.google.com/mail/?p=InvalidSecondFactor z2sm11041738pfq.58 - gsmtp
SMTP Error: Could not authenticate.
CLIENT -> SERVER: QUIT
SMTP connect() failed. https://github.com/PHPMailer/PHPMailer/wiki/Troubleshooting
Problem sending email.
```

**Error 2:** When application specific password is incorrect
```
SMTP ERROR: Password command failed: 535-5.7.8 Username and Password not accepted. Learn more at 535 5.7.8  https://support.google.com/mail/?p=BadCredentials f3sm5807314pgj.62 - gsmtp
SMTP Error: Could not authenticate.
CLIENT -> SERVER: QUIT
SMTP connect() failed. https://github.com/PHPMailer/PHPMailer/wiki/Troubleshooting
Problem sending email.
```

**Error 3:** When recipient email is invalid
```
Invalid address:  (to): recipient-email
Problem sending email.
```

**Error 4:** Debugger step-by-step output after sending email successfully
```
CLIENT -> SERVER: EHLO NL616
CLIENT -> SERVER: STARTTLS
CLIENT -> SERVER: EHLO NL616
CLIENT -> SERVER: AUTH LOGIN
CLIENT -> SERVER: <credentials hidden>
CLIENT -> SERVER: <credentials hidden>
CLIENT -> SERVER: MAIL FROM:<from-email@gmail.com>
CLIENT -> SERVER: RCPT TO:<recipient-email@domain>
CLIENT -> SERVER: RCPT TO:<cc-recipient-email@domain>
CLIENT -> SERVER: DATA
CLIENT -> SERVER: Date: Sun, 22 Sep 2019 05:11:15 +0000
CLIENT -> SERVER: To: recipient-name <recipient-email@domain>
CLIENT -> SERVER: From: PHP SMTP Mailer <from-email@gmail.com>
CLIENT -> SERVER: Cc: cc-recipient-name <cc-recipient-email@domain>
CLIENT -> SERVER: Reply-To: reply-to-name <reply-to-email@domain>
CLIENT -> SERVER: Subject: Test email using PHP mailer
CLIENT -> SERVER: Message-ID: <UlnH3mCpHcFVNBY3Lb3PR2tVs6tvdJlu2F8g5sPN4@NL616>
CLIENT -> SERVER: X-Mailer: PHPMailer 6.0.7 (https://github.com/PHPMailer/PHPMailer)
CLIENT -> SERVER: MIME-Version: 1.0
CLIENT -> SERVER: Content-Type: multipart/alternative;
CLIENT -> SERVER:  boundary="b1_UlnH3mCpHcFVNBY3Lb3PR2tVs6tvdJlu2F8g5sPN4"
CLIENT -> SERVER: Content-Transfer-Encoding: 8bit
CLIENT -> SERVER: This is a multi-part message in MIME format.
CLIENT -> SERVER: --b1_UlnH3mCpHcFVNBY3Lb3PR2tVs6tvdJlu2F8g5sPN4
CLIENT -> SERVER: Content-Type: text/plain; charset=us-ascii
CLIENT -> SERVER: This is a test email using PHP mailer class.
CLIENT -> SERVER: --b1_UlnH3mCpHcFVNBY3Lb3PR2tVs6tvdJlu2F8g5sPN4
CLIENT -> SERVER: Content-Type: text/html; charset=us-ascii
CLIENT -> SERVER: <b>This is a test email using PHP mailer class.</b>
CLIENT -> SERVER: --b1_UlnH3mCpHcFVNBY3Lb3PR2tVs6tvdJlu2F8g5sPN4--
CLIENT -> SERVER: QUIT
email sent.
```

## Conclusion
Hope the steps explained above were useful and you were able to successfully send mail from your Gmail SMTP server using PHP. Feel free to contribute, in case you encountered some issue which is not listed as a part of this tutorials. Use below comments section to ask/share any feedback.

<? Happy Coding ?>
