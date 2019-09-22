# How to send an E-mail via GMail SMTP Server using PHP

**1. INTRODUCTION**

Using GMail SMTP Server you can send E-mails to any domain using your Gmail Credentials. Following are some Emails sending limit criterias.
+ Google limits the number of recipients in a single Email and number of Emails can be sent per day.
+ Current limit is 500 Emails in a day or 500 recipients in a single Email.
+ On reaching threshold limits, You can not send messages for 1 to 24 hours.
+ After Suspension Period counters will get reset automatically and the user can resume sending Emails.
+ For more information about Email sending limits refer following links:
  - Link 1: [Email sending limits](https://support.google.com/a/answer/166852)
  - Link 2: [Error messages once limit is crossed](https://support.google.com/mail/answer/22839)

**2. SETTINGS**

**2.1:** Before sending Emails using Gmail SMTP Server, Change the required setting using Google Account Security Settings or [Click Here](https://myaccount.google.com/security)

![Google Account Security Settings](https://i.imgur.com/6Hxmb2G.png))

**2.2:** Make sure that 2-Step-Verification is Disabled

![2-Step Virification Disabled](https://i.imgur.com/6Hxmb2G.png)

**2.3:** Turn ON the Less Secure App Access or [Click Here](https://myaccount.google.com/u/0/lesssecureapps)

![Less Secure App Access](https://i.imgur.com/hymkYJ6.png)

**2.4:** If 2-Step-Verification is Enabled, then you will have to create APP Password for your application or device.

![2-Step Virification Enabled](https://i.imgur.com/vcQYoGo.png)

![Generate App Password](https://i.imgur.com/LHfCxdH.png)

**2.5:** For security precaution, Google may require you to complete this additional step while signing-in. [Click Here](https://accounts.google.com/DisplayUnlockCaptcha) to Allow access to your Google account using new Device/App.

![New Device-App](https://i.imgur.com/mEGa22F.png)

==**Note**: It may take an hour or more to reflect any security changes==

**3. DOWNLOAD PHP LIBRARY**
+ [Click here](https://github.com/PHPMailer/PHPMailer/) for more details about PHP Mailer
+ [Click here](https://github.com/PHPMailer/PHPMailer/archive/master.zip) to download .zip file
+ Unzip the master.zip in your application directory and run following command from your application Directory.

```
composer require phpmailer/phpmailer
```
[Composer](https://getcomposer.org/) is the recommended way to install PHPMailer.

**4. PHP CODE**
+ Using your Gmail Credentials, Connect to Host smtp.gmail.com
  - On port 465, if you’re using SSL
  - On port 587, if you’re using TLS
+ [Click here](https://github.com/PHPMailer/PHPMailer/wiki/Tutorial) for some more Examples and Tutorials of PHPMailer


**4.1:** Include packages and files for PHP Mailers and SMTP Protocol
``` php
use PHPMailer\PHPMailer\PHPMailer;
use PHPMailer\PHPMailer\Exception;
require 'PHPMailer-master/src/Exception.php';
require 'PHPMailer-master/src/PHPMailer.php';
require 'PHPMailer-master/src/SMTP.php';
```

**4.2:** Initialize PHP Mailer and set SMTP as mailing Protocol 
``` php
$mail = new PHPMailer();
$mail->IsSMTP();
$mail->Mailer = "smtp";
```

**4.3:** Set required Parameters for SMTP Connection like Server, Port and account credentials. SSL and TLS are both Cryptographic protocols that provide authentication and data encryption between servers, machines and applications operating over a network. SSL is the predecessor to TLS.

``` php
$mail->SMTPDebug  = 1;  
$mail->SMTPAuth   = TRUE;
$mail->SMTPSecure = "tls";
$mail->Port       = 587;
$mail->Host       = "smtp.gmail.com";
$mail->Username   = "your-email@gmail.com";
$mail->Password   = "your-gmail-password";
```
**4.4:** Set required parameters for Email Header and Body
``` php
$mail->IsHTML(true);
$mail->AddAddress("recipient-email@domain", "recipient-name");
$mail->SetFrom("from-email@gmail.com", "from-name");
$mail->AddReplyTo("reply-to-email@domain", "reply-to-name");
$mail->AddCC("cc-recipient-email@domain", "cc-recipient-name");
$mail->Subject = "Test is Test Email sent via Gmail SMTP Server using PHP Mailer";
$content = "<b>This is a Test Email sent via Gmail SMTP Server using PHP mailer class.</b>";
```
**4.5:** Send the Email and catch required exceptions
``` php
$mail->MsgHTML($content); 
if(!$mail->Send()) {
  echo "Error while sending Email.";
  var_dump($mail);
} else {
  echo "Email sent successfully";
}
```
[Click here](https://github.com/gaurangmacharya/pepithon/blob/master/send-email-via-gmail-smtp-server-using-phpmailer.php) to download the complete code.

**5. ERRORS**

**5.1:** When SMTP Mail Server Credentials were correct but Application Specific Password was not provided
```
SMTP ERROR: Password command failed: 534-5.7.9 Application-specific password required. Learn more at 534 5.7.9  https://support.google.com/mail/?p=InvalidSecondFactor z2sm11041738pfq.58 - gsmtp
SMTP Error: Could not authenticate.
CLIENT -> SERVER: QUIT
SMTP connect() failed. https://github.com/PHPMailer/PHPMailer/wiki/Troubleshooting
Problem sending email.
```

**5.2:** When application specific password is incorrect
```
SMTP ERROR: Password command failed: 535-5.7.8 Username and Password not accepted. Learn more at 535 5.7.8  https://support.google.com/mail/?p=BadCredentials f3sm5807314pgj.62 - gsmtp
SMTP Error: Could not authenticate.
CLIENT -> SERVER: QUIT
SMTP connect() failed. https://github.com/PHPMailer/PHPMailer/wiki/Troubleshooting
Problem sending email.
```

**5.3:** When recipient email is invalid
```
Invalid address:  (to): recipient-email
Problem sending email.
```

**5.4:** Debugger output after sending email successfully
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

<br>

**6. ALTERNATE SOLUTIONS**
+ **G Suite SMTP Relay Service:** Send mail from your organization by authenticating with the IP addresses. You can send messages to anyone inside or outside of your domain.
+ **Restricted Gmail SMTP Server:** Send messages to Gmail or G Suite users only. This option does not require you to authenticate.
+ [Click here](https://support.google.com/a/answer/176600) to see detailed comparison of all 3 services.
