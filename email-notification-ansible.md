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
- tasks:
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

**Few other notification modules in Ansible**
+ hipchat – Send a message to Hipchat
+ jabber – Send a message to jabber user or chat room
+ rabbitmq_publish – Publish a message to a RabbitMQ queue
+ rocketchat – Send notifications to Rocket Chat
+ say – Makes a computer to speak
+ sendgrid – Sends an email with the SendGrid API
+ slack – Send Slack notifications
+ telegram – module for sending notifications via telegram
