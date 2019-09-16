# Installing, Configuring nad using Sendmail Command

++**Synopsis**++
<br>
Sendmail is the default SMTP (Simple mail transfer protocol) server installed on CentOS and although it can be slightly more complicated for beginners to learn than similar alternatives such as Postfix, it can be quite powerful and useful to learn. Sendmail by itself as the name suggests is a MTA (Mail transfer agent) which is useful for sending emails from your server to other servers. However in order to create a proper mail server you will also need a few other pieces of software such as a MUA (Mail user agent) to connect to the MTA and a POP3/IMAP (Post office protocol / Internet message access protocol) server to listen for incoming emails. These will all be covered in later tutorials, but in this tutorial we will just cover some of the basics of Sendmail to help you become familiar with email servers.

**++Terminologies++**

This all might be a bit confusing to understand at first, but it’s actually quite straight forward once you get to know all of the terms. I have included some more information below which will help you to remember some of the terms and their role in a mail server.

+ MTA  – Mail Transfer Agent – Sendmail, Postfix
+ MDA  – Mail Delivery Agent – Procmail
+ MUA  – Mail User Agent – Outlook, Thunderbird
+ SMTP – Simple Mail Transfer Protocol – Sendmail, Postfix
+ POP3 – Post Office Protocol 3 – Dovecot
+ IMAP – Internet Message Access Protocol – Dovecot

**++Installation++**
```
yum install sendmail
```

**++Starting Service++**
```
/etc/init.d/sendmail start
or
systemctl start sendmail
```

**++DNS Records++**
<br>
Next you will need to modify your DNS records so that you can use your domain name as an email address. First you will need to create a sub-domain under your domain name, for example mail.example.com then you will need to forward the MX records of mail.example.com to point to the IP address of your server. Modifying the records of your DNS server particularly if you are using an external hosting company may take 24 to 48 hours to propagate, so make sure that your settings are correct before you modify them.

**++Accepting Incoming Email++**
You will also need to specify on your server which domains you would like to accept mail from. You can do this by editing the /etc/mail/local-host-names file and adding the domain names as well as the sub-domains that you wish to accept. For example you can add the following lines to the file.
example.com
mail.example.com

**++Using Your Domain Name++**
When you check your email the first thing you will notice is that your email came from your hostname not your domain name. If you want a quick fix to change this you can simply change your hostname to become your domain name either through the /etc/hosts file or by entering the following command.
```
hostname example.com
```
You can also masquerade as another domain name through the sendmail.mc file. Adding the following lines to your sendmail.mc file will allow all your emails to appear as though they are coming from the domain name example.com.
```
MASQUERADE_AS(`example.com’)dnl
MASQUERADE_DOMAIN(localhost)dnl
FEATURE(masquerade_entire_domain)dnl
```

**++Forwarding Emails++**
<br>
Sendmail also has the ability to forward emails to multiple addresses; you can do this by editing the alias file located at /etc/aliases. The aliases file allows you to forward emails to local accounts or other emails. This is especially useful for people who want to create a simple mailing list or create email groups. For example if you have a set of users to answer support emails you could add the following line to /etc/aliases to allow them all to receive supports emails.
support: bob, henry, john, support@example.com

**++Setting Relay Access++**
<br>
By default only the localhost is setup to be able to relay emails to other servers. You can give relay access to users by specifying the servers you’d like them to connect to or the IP that they are connecting from.
```
To modify the relay access open the /etc/mail/access file
```
```
You can relay or reject domain names by adding the lines below
To:gmail.com RELAY
To:spam.com REJECT
```
```
You can also specify IP addresses of hosts you would like to relay or reject
Connect:127.0.0.1 RELAY
Connect:123.123.123.123 REJECT
```

**++Command Line Parameters++**
```
-Ac 	Use submit.cf even if the operation mode does not indicate an initial mail submission.
-Am 	Use sendmail.cf even if the operation mode indicates an initial mail submission.
-Btype 	Set the body type to type. Current legal values are 7BIT or 8BITMIME.
-ba 	Go into ARPANET mode. All input lines must end with a CR-LF, and all messages will be generated with a CR-LF at the end. Also, the "From:" and "Sender:" fields are examined for the name of the sender.
-bd 	Run as a daemon. sendmail will fork and run in background listening on socket 25 for incoming SMTP connections. This is normally run from /etc/rc.
-bD 	Same as -bd, except it runs in foreground.
-bh 	Print the persistent host status database.
-bH 	Purge expired entries from the persistent host status database.
-bi 	Initialize the alias database.
-bm 	Deliver mail in the usual way (This is the default).
-bp 	Print a listing of the queue(s).
-bP 	Print number of entries in the queue(s); only available with shared memory support.
-bs 	Use the SMTP protocol as described in RFC 821 on standard input and output. This flag implies all the operations of the -ba flag that are compatible with SMTP.
-bt 	Run in address test mode. This mode reads addresses and shows the steps in parsing; it is used for debugging configuration tables.
-bv 	Verify names only; do not try to collect or deliver a message. Verify mode is normally used for validating users or mailing lists.
-Cfile 	Use alternate configuration file. sendmail gives up any enhanced (set-user-ID or set-group-ID) privileges if an alternate configuration file is specified.
-Dlogfile 	Send debugging output to the indicated log file instead of stdout.
-dcategory.level... 	Set the debugging flag for category to level. The category is either an integer or a name specifying the topic, and level an integer specifying the level of debugging output desired. Higher levels generally mean more output. More than one flag can be specified by separating them with commas. A list of numeric debugging categories can be found in the TRACEFLAGS file in the sendmail source distribution. The option -d0.1 prints the version of sendmail and the options used during the compile. Most other categories are only useful with, and documented in, sendmail's source code.
-Ffullname 	Set the full name of the sender.
-fname 	Sets the name of the "from" person (i.e., the envelope sender of the mail). This address may also be used in the "From:" header if that header is missing during initial submission. The envelope sender address is used as the recipient for delivery status notifications and may also appear in a "Return-Path:" header. -f should only be used by "trusted" users (normally root, daemon, and network) or if the person you are trying to become is the same as the person you are. Otherwise, an "X-Authentication-Warning" header will be added to the message.
-G 	Relay (gateway) submission of a message, e.g., when rmail calls sendmail.
-hN 	Set the hop count to N. The hop count is incremented every time the mail is processed. When it reaches a limit, the mail is returned with an error message, most likely the victim of an aliasing loop. If not specified, "Received:" lines in the message are counted.
-i 	Ignore dots alone on lines by themselves in incoming messages. This should be set if you are reading data from a file.
-L tag 	Set the identifier used in syslog messages to the supplied tag.
-N dsn 	Set delivery status notification conditions to dsn, which can be "never" for no notifications, or a comma separated list of the values "failure" to be notified if delivery failed, "delay" to be notified if delivery is delayed, and "success" to be notified when the message is successfully delivered.
-n 	Don't do aliasing.
-O option=value 	Set option option to the specified value. This form uses long names. See below for more details.
-ox value 	Set option x to the specified value. This form uses single character names only.
-pprotocol 	Set the name of the protocol used to receive the message. This can be a simple protocol name such as "UUCP" or a protocol and hostname, such as "UUCP:ucbvax".
-q[time] 	Process saved messages in the queue at given intervals. If time is omitted, process the queue once. The time is given as a tagged number, with "s" being seconds, "m" being minutes (default), "h" being hours, "d" being days, and "w" being weeks. For example, "-q1h30m" or "-q90m" would both set the timeout to one hour thirty minutes. By default, sendmail will run in the background. This option can be used safely with -bd.
-qp[time] 	Similar to -qtime, except that instead of periodically forking a child to process the queue, sendmail forks a single persistent child for each queue that alternates between processing the queue and sleeping. The sleep time is given as the argument; it defaults to 1 second. The process will always sleep at least 5 seconds if the queue was empty in the previous queue run.
-qf 	Process saved messages in the queue once and do not fork, but run in the foreground.
-qGname 	Process jobs in queue group called name only.
-q[!]Isubstr 	Limit processed jobs to those containing substr as a substring of the queue id or not when ! is specified.
-q[!]Qsubstr 	Limit processed jobs to quarantined jobs containing substr as a substring of the quarantine reason or not when ! is specified.
-q[!]Rsubstr 	Limit processed jobs to those containing substr as a substring of one of the recipients or not when ! is specified.
-q[!]Ssubstr 	Limit processed jobs to those containing substr as a substring of the sender or not when ! is specified.
-Q[reason] 	Quarantine a normal queue items with the given reason or unquarantine quarantined queue items if no reason is given. This should only be used with some sort of item matching using as described above.
-R return 	Set the amount of the message to be returned if the message bounces. The return parameter can be "full" to return the entire message or "hdrs" to return only the headers. In the latter case, local bounces return only the headers.
-rname 	An alternate and obsolete form of the -f flag.
-t 	Read message for recipients. To:, Cc:, and Bcc: lines will be scanned for recipient addresses. The Bcc: line will be deleted before transmission.
-V envid 	Set the original envelope id. This is propagated across SMTP to servers that support DSNs and is returned in DSN-compliant error messages.
-v 	Go into verbose mode. Alias expansions will be announced, etc.
-X logfile 	Log all traffic in and out of mailers in the indicated log file. This should only be used as a last resort for debugging mailer bugs. It will log a lot of data very quickly.
-- 	Stop processing command flags and use the rest of the arguments as addresses.
```

**++Command Line Options++**
```
AliasFile=file 	Use alternate alias file.
HoldExpensive 	On mailers that are considered "expensive" to connect to, don't initiate immediate connection. This requires queueing.
CheckpointInterval=N 	Checkpoint the queue file after every N successful deliveries (default 10). This avoids excessive duplicate deliveries when sending to long mailing lists interrupted by system crashes.
DeliveryMode=x 	Set the delivery mode to x. Delivery modes are "i" for interactive (synchronous) delivery, "b" for background (asynchronous) delivery, "q" for queue only; i.e., actual delivery is done the next time the queue is run, and "d" for deferred: the same as "q" except that database lookups for maps which have set the -D option (default for the host map) are avoided.
ErrorMode=x 	Set error processing to mode x. Valid modes are "m" to mail back the error message, "w" to "write" back the error message (or mail it back if the sender is not logged in), "p" to print the errors on the terminal (default), "q" to throw away error messages (only exit status is returned), and "e" to do special processing for BerkNet. If the text of the message is not mailed back by modes "m" or "w" and if the sender is local to this machine, a copy of the message is appended to the file dead.letter in the sender's home directory.
SaveFromLine 	Save Unix-style From lines at the front of messages.
MaxHopCount=N 	The maximum number of times a message is allowed to "hop" before we decide it is in a loop.
IgnoreDots 	Do not take dots on a line by themselves as a message terminator.
SendMimeErrors 	Send error messages in MIME format. If not set, the DSN (Delivery Status Notification) SMTP extension is disabled.
ConnectionCacheTimeout=timeout 	Set connection cache timeout.
ConnectionCacheSize=N 	Set connection cache size.
LogLevel=n 	The log level.
MeToo=False 	Don't send to "me" (the sender) if I am in an alias expansion.
CheckAliases 	Validate the right side of aliases during a newaliases command.
OldStyleHeaders 	If set, this message may have old style headers. If not set, this message is guaranteed to have new style headers (i.e., commas instead of spaces between addresses). If set, an adaptive algorithm is used that will correctly determine the header format in most cases.
QueueDirectory=queuedir 	Select the directory in which to queue messages.
StatusFile=file 	Save statistics in the named file.
Timeout.queuereturn=time 	Set the timeout on undelivered messages in the queue to the specified time. After delivery has failed (e.g., because of a host being down) for this amount of time, failed messages will be returned to the sender. The default is five days.
UserDatabaseSpec=userdatabase 	If set, a user database is consulted to get forwarding information. You can consider this an adjunct to the aliasing mechanism, except that the database is intended to be distributed; aliases are local to a particular host. This may not be available if your sendmail does not have the USERDB option compiled in.
ForkEachJob 	Fork each job during queue runs. May be convenient on memory-poor machines.
SevenBitInput 	Strip incoming messages to seven bits.
EightBitMode=mode 	Set the handling of eight bit input to seven bit destinations to mode: m (mimefy) will convert to seven-bit MIME format, p (pass) will pass it as eight bits (but violates protocols), and s (strict) will bounce the message.
MinQueueAge=timeout 	Sets how long a job must ferment in the queue between attempts to send it.
DefaultCharSet=charset 	Sets the default character set used to label 8-bit data that is not otherwise labelled.
DialDelay=sleeptime 	If opening a connection fails, sleep for sleeptime seconds and try again. Useful on dial-on-demand sites.
NoRecipientAction=action 	Set the behaviour when there are no recipient headers (To:, Cc: or Bcc:) in the message to action: none leaves the message unchanged, add-to adds a To: header with the envelope recipients, add-apparently-to adds an Apparently-To: header with the envelope recipients, add-bcc adds an empty Bcc: header, and add-to-undisclosed adds a header reading `To: undisclosed-recipients:;'.
MaxDaemonChildren=N 	Sets the maximum number of children that an incoming SMTP daemon will allow to spawn at any time to N.
ConnectionRateThrottle=N 	Sets the maximum number of connections per second to the SMTP port to N.
```