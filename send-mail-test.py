#
# Use this script to quickly send a bunch of mails. Useful for testing.
#
import sys
import json
import time
import smtplib
import urllib2
import datetime

from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart

def getQuote():
	response = urllib2.urlopen("http://www.iheartquotes.com/api/v1/random?format=json")
	obj = json.loads(response.read())

	quoteLines = obj["quote"].split("--")

	if len(quoteLines) > 0:
		return {
			"quote": quoteLines[0].strip(),
			"source": "Unknown" if len(quoteLines) <= 1 else quoteLines[1].strip(),
		}
	else:
		return {
			"quote": "No quote",
			"source": "Adam Presley"
		}

if __name__ == "__main__":
	numMails = 5
	address = "127.0.0.1"
	smtpPort = 8000

	me = "someone@another.com"
	to = "bob@bobtestingmailslurper.com"

	try:
		for index in range(numMails):
			quote = getQuote()

			textBody = "Hello,\nHere is today's quote.\n\n{0}\n  -- {1}\n\nSincerely,\nAdam Presley".format(quote["quote"], quote["source"])
			htmlBody = "<p>Hello,</p><p>Here is today's quote.</p><p><em>{0}</em><br />&nbsp;&nbsp;-- {1}</p><p>Sincerely,<br />Adam Presley</p>".format(quote["quote"], quote["source"],)

			text = MIMEText(textBody, "plain")
			html = MIMEText(htmlBody, "html")

			msg = MIMEMultipart("alternative")

			msg["Subject"] = "Quote From {0}".format(quote["source"])
			msg["From"] = me
			msg["To"] = to
			msg["Date"] = datetime.datetime.now().strftime("%a, %d %b %Y %H:%M:%S +0000 UTC")

			msg.attach(text)
			msg.attach(html)

			server = smtplib.SMTP("{0}:{1}".format(address, smtpPort))
			server.sendmail(me, [to], msg.as_string())
			server.quit()

			time.sleep(2)

		for index in range(numMails):
			textBody = "Hello,\nI am mail #{0}.\n\nSincerely,\nAdam Presley".format(index,)

			msg = MIMEText(textBody)

			msg["Subject"] = "Test Mail #{0}".format(index,)
			msg["From"] = me
			msg["To"] = to
			msg["Date"] = datetime.datetime.now().strftime("%a, %d %b %Y %H:%M:%S -0700 (UTC)")

			server = smtplib.SMTP("{0}:{1}".format(address, smtpPort))
			server.sendmail(me, [to], msg.as_string())
			server.quit()

			time.sleep(1)


	except Exception as e:
		print("An error occurred while trying to connect and send the email: {0}".format(e.message))
		print(sys.exc_info())
