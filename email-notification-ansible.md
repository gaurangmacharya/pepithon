# How to send an Email Notification using Ansible?

**Synopsis**

Ansible is a radically simple IT automation engine that automates cloud provisioning, configuration management, application deployment, intra-service orchestration, and many other IT needs. Ansible comes with 20+ Built in modules with more than 3000+ functionality.  

[Click here](https://docs.ansible.com/ansible/latest/installation_guide/index.html) to go to Ansible Installation Guide
and for installing Ansible on RHEL/CentOS run:
```
$ sudo yum install ansible
```

After completion of any Automated Task i.e. Playbook It is very important to record and send detailed report to stake holders so that they act and play their role. For this Ansible can send Emails with attachment via
- Local Configured Email Server say Postfix
- Remote Email Server with required access and credentials
- Sending a mail using SMTP Services like Pepipost, Gmail, Mandrill, Mailjet, SendGrid etc
[Click here](https://docs.ansible.com/ansible/latest/modules/mail_module.html) to know more about Email module in Ansible.

**Important Parameters:**

```
host    : The mail server. Default is localhost.
port    : The mail server port. Mostly 25, 465, 587
username: If SMTP requires username.
password: If SMTP requires password.
timeout : Sets the timeout in seconds for connection attempts.
headers : A list of headers which should be added to the message.

from    : The email-address the mail is sent from. Default is root.
to      : The email-address(es) the mail is being sent to.
cc      : The email-address(es) the mail is being copied to
bcc     : The list of email-address(es) the mail is being 'blind' copied to.

body    : The body of the email being sent.
charset : The character set of email being sent. Default is UTF-8
subject : The subject of the email being sent. This is a mandatory field.
subtype : The minor mime type, can be either plain or html. The major type is always text.
attach  : A list of path-names of files to attach to the message. Attached files will have their content-type set to application/octet-stream
```


**Example-1: With minimum parameters without attachement**

``` yml
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

**Example-2: Sending Report as attachment using In-House SMTP Server**
``` yml
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
==**Note:** Please check Host-Port Connection using Ping, Telnet command or else you may get an error saying socket.error: [Errno 111] Connection refused==

**Error 1:** When mailing service was not running on Port 25 OR connecion was refused by server
```
An exception occurred during task execution. To see the full traceback, use -vvv. The error was: socket.error: [Errno 111] Connection refused
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
```
fatal: [localhost -> localhost]: FAILED! => {"changed": false, "msg": "Authentication to smtp.gmail.com:587 failed, please check your username and/or password", "rc": 1}
```

**Error 3:** When file to be attached was either missing at defined location or was inaccessible.
```
An exception occurred during task execution. To see the full traceback, use -vvv. The error was: IOError: [Errno 2] No such file or directory: '/tmp/ansible.logs'
fatal: [localhost -> localhost]: FAILED! => 
array (
  'changed' => false,
  'msg' => 'Failed to send mail: can\'t attach file /tmp/ansible.logs: [Errno 2] No such file or directory: \'/tmp/ansible.logs\'',
  'rc' => 1,
)
```

**Error 4:** When recipient email address was incorrect.
```
An exception occurred during task execution. To see the full traceback, use -vvv. The error was: SMTPRecipientsRefused: {'to-email-1': (550, '5.1.1 <to-email-1>: Recipient address rejected: User unknown in local recipient table')}
fatal: [localhost -> localhost]: FAILED! => 
array (
  'changed' => false,
  'msg' => 'Failed to send mail to \'to-email-1\': {\'to-email-1\': (550, \'5.1.1 : Recipient address rejected: User unknown in local recipient table\')}',
  'rc' => 1,
)
```
**Few other notification modules in Ansible**
+ hipchat – Send a message to Hipchat
+ jabber – Send a message to jabber user or chat room
+ rabbitmq_publish – Publish a message to a RabbitMQ queue
+ rocketchat – Send notifications to Rocket Chat
+ say – Makes a computer to speak
+ sendgrid – Sends an email with the SendGrid API
+ slack – Send Slack notifications
+ telegram – module for sending notifications via telegram
