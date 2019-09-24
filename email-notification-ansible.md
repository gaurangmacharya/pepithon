# How to Send an Email Notifications using Ansible?

## Introduction

Ansible is a radically simple IT automation engine that automates cloud provisioning, configuration management, application deployment, intra-service orchestration, and many other IT needs. Ansible comes with 20+ Built-in modules with more than 3000+ functionality.
Click [here](https://docs.ansible.com/ansible/latest/installation_guide/index.html "here") to go to Ansible installation guide and for installing Ansible on RHEL/CentOS run:

```command
$ sudo yum install ansible
```
It is very important to record the completion of any automated task, i.e. Playbook, and send a detailed report to stakeholders so that they act accordingly. Few cases can be:
- After installation of any application
- After any activity in GIT
- When a specific job has been completed

Ansible can send emails with attachment via;
- Local Configured Email Server say Postfix
- Remote Email Server with required access and credentials
- Sending a mail using SMTP Services like Pepipost, Gmail, Mandrill, Mailjet, SendGrid. Click [here](https://docs.ansible.com/ansible/latest/modules/mail_module.html "here") to know more about Email module in Ansible.

## Parameters

Let's understand the different parameters used for making an SMTP connection and to send an email.

#### 1. SMTP Connection
```
host: The address of the mail server. Default is localhost.
port: The port number of the mail server to connect. Mostly 25, 465, 587.
username: If SMTP requires a username.
password: If SMTP requires a password.
timeout: Sets the timeout in seconds for connection attempts.
```

#### 2. Connection Security
```
secure= always. The connection sends email only if the connection is encrypted. It fails if the server doesn't accept the encrypted connection.
secure=never. Before sending an email, the connection doesn't attempt to set up a secure SSL/TLS session. 
secure=try. Before trying to send an email, the connection attempts to set up a secure SSL/TLS session.
secure=starttls. Before sending an email, the connection tries to upgrade to a secure SSL/TLS connection. The connection fails in case if unable to do so.
```

#### 3. Email Headers
```
headers : A list of headers that needs to go with the message.
from    : The email address from which the mail is sent. Default is root.
to      : The email address(es) of the recipient to whom the mail sent.
cc      : The email address(es) of the recipient to whom the mail copied.
bcc     : The email address(es) of the recipient to whom the mail 'blind' copied.
```

#### 4. Subject
```
subject : The subject of the email to send (mandatory).
subtype : The minor mime type, can be either plain or HTML. The major type is always text.
```

#### 5. Email Body
```
charset : The character set of the email to send. Default is UTF-8
body    : The body of the email to send.
attach  : A list of files (full path) to attach to the mail. The content-type should be set to "application/octet-stream" for all the attached files.
```

## Code to Send Email Without Attachment Using Gmail SMTP Server
```
- hosts:
    - localhost
  tasks:
    - name: Sending an e-mail using Gmail SMTP servers
      mail:
        host: smtp.gmail.com
        port: 587
        username: username@gmail.com
        password: your-password
        to: recipient-name <recipient-email@domain>
        subject: Ansible Report
        body: System {{ ansible_hostname }} has been successfully provisioned.
      delegate_to: localhost
```

## Code to Send Email With Attachment Using In-House Email Server
```
- hosts:
    - localhost
  tasks:
    - name: Send Emails to a bunch of users, with Playbook Report as an attachment.
      mail:
        host: localhost
        port: 25
        subject: Ansible Playbook Report
        body: This is an Email generated using Ansible after execution of task.
        from: from-email@domain (Ansible Automates)
        to:
        - to-email-name-1 <to-email-1@domain>
        - to-email-name-2 <to-email-1@domain>
        cc: cc-name <cc-email@domain>
        attach:
        - <enter-path-of-your-file-to-be-attached>
        headers:
        - Reply-To=reply-to-email@domain
        - X-Special="Write something special about this Email"
        charset: us-ascii
      delegate_to: localhost
```
> Note: Please check Host-Port Connection using Ping, Telnet command or else you may get an error saying socket.error: [Errno 111] connection refused

Hope, you are now able to send email notifications from your Ansible set up. In case you are unable to send emails and getting some error then, please refer the below list:

## Possible Errors/Exceptions

**Error 1:** When mailing service was not running on Port 25 OR connection was refused by server
An exception occurred during task execution. To see the full traceback, use -vvv. The error was: 
```
socket.error: [Errno 111] Connection refused
fatal: [localhost -> localhost]: FAILED! => 
array (
  'changed' => false,
  'module_stderr' => 'Traceback (most recent call last):
  File ".ansible/tmp/ansible-tmp-1569048408.38-18927787866074/AnsiballZ_mail.py", line 114, in _ansiballz_main()
  File ".ansible/tmp/ansible-tmp-1569048408.38-18927787866074/AnsiballZ_mail.py", line 106, in _ansiballz_main invoke_module(zipped_mod, temp_path, ANSIBALLZ_PARAMS)
  File ".ansible/tmp/ansible-tmp-1569048408.38-18927787866074/AnsiballZ_mail.py", line 49, in invoke_module imp.load_module(\'__main__\', mod, module, MOD_DESC)
  File "/tmp/ansible_mail_payload_fiQfEe/__main__.py", line 397, in 
  File "/tmp/ansible_mail_payload_fiQfEe/__main__.py", line 286, in main
  File "/usr/lib64/python2.7/smtplib.py", line 315, in connect self.sock = self._get_socket(host, port, self.timeout)
  File "/usr/lib64/python2.7/smtplib.py", line 290, in _get_socket return socket.create_connection((host, port), timeout)
  File "/usr/lib64/python2.7/socket.py", line 571, in create_connection raise err socket.error: [Errno 111] Connection refused',
  'module_stdout' => '',
  'msg' => 'MODULE FAILURE
See stdout/stderr for the exact error',
  'rc' => 1,
)
```
**Error 2:** When Gmail account credentials were incorrect
fatal: 
```
[localhost -> localhost]: FAILED! => {"changed": false, "msg": "Authentication to smtp.gmail.com:587 failed, please check your username and/or password", "rc": 1}
```

**Error 3:** When the file to be attached was either missing at a defined location or was inaccessible
An exception occurred during task execution. To see the full traceback, use -vvv. The error was: 
```
IOError: [Errno 2] No such file or directory: '/tmp/ansible.logs'
fatal: [localhost -> localhost]: FAILED! => 
array (
  'changed' => false,
  'msg' => 'Failed to send mail: can\'t attach file /tmp/ansible.logs: [Errno 2] No such file or directory: \'/tmp/ansible.logs\'',
  'rc' => 1,
)
```
**Error 4:** When the recipient email address was incorrect
An exception occurred during task execution. To see the full traceback, use -vvv. The error was: 
```
SMTPRecipientsRefused: {'to-email-1': (550, '5.1.1 <to-email-1>: Recipient address rejected: User unknown in local recipient table')}
fatal: [localhost -> localhost]: FAILED! => 
array (
  'changed' => false,
  'msg' => 'Failed to send mail to \'to-email-1\': {\'to-email-1\': (550, \'5.1.1 : Recipient address rejected: User unknown in local recipient table\')}',
  'rc' => 1,
)
```
**Error 5:** When any mandatory parameters is missing like here "Subject" was missing
fatal: 
```
[localhost -> localhost]: FAILED! => {"changed": false, "msg": "missing required arguments: subject"}
```

Few Other Important Notification Sending Modules In Ansible:

- hipchat - Send a message to Hipchat
- jabber - Send a message to jabber user or chat room
- rabbitmq_publish - Publish a message to a RabbitMQ queue
- rocketchat - Send notifications to Rocket Chat
- say - Makes a computer to speak
- slack - Send Slack notifications
- telegram - module for sending notifications via telegram

**In case you are facing some issues which are not listed above in the tutorial, or you have some suggestions, then please feel free to contribute below in comments.**
