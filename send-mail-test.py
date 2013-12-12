#
# Use this script to quickly send a bunch of mails. Useful for testing.
#
import smtplib
import datetime
import time
from email.mime.text import MIMEText

numMails = 10
address = "localhost"
smtpPort = 8000

me = "someone@another.com"
to = "bob@bobtestingmailslurper.com"
body = """Hello,

This is a test message. It is here to test the SMTP server and
admin called MailSlurper. I sure hope it all works!!

Sincerely,
Adam Presley"""

try:

	for index in range(numMails):
		msg = MIMEText(body)
		msg["Subject"] = "Test Message #{0}".format(index)
		msg["From"] = me
		msg["To"] = to
		msg["Date"] = datetime.datetime.now().strftime("%a, %d %b %Y %H:%M:%S +0000 UTC")

		server = smtplib.SMTP("{0}:{1}".format(address, smtpPort))
		server.sendmail(me, [to], msg.as_string())
		server.quit()

		time.sleep(3)


except Exception as e:
	print("An error occurred while trying to connect and send the email: {0}".format(e.message))
