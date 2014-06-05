---
title: Home
layout: default
---

<section id="downloads" class="clearfix">
	<a href="https://github.com/adampresley/mailslurper-go/releases/tag/4.0" class="button"><span>Get Release 4.0!</span></a>
	<a href="https://github.com/adampresley/mailslurper-go/zipball/master" id="download-zip" class="button"><span>Download .zip</span></a>
	<a href="https://github.com/adampresley/mailslurper-go" id="view-on-github" class="button"><span>View on GitHub</span></a>
</section>

<hr />

### What Is MailSlurper?
MailSlurper is a simple SMTP server that slurps mail into oblivion! Having MailSlurper installed on your local
or shared development environment allows you to test sending email without having to go through the trouble
of configuring a real mail server. You also don't have to worry about emails actually going out to unintended
recipients.

Simply configure your application and/or application server stack to the host and port you are running
MailSlurper on and it will capture emails you send out into a database for later viewing. Then open
up your favorite web browser to http://localhost:8080 and you can view captured emails and their attachments!

### How To Run MailSluper
Make sure you have the following:

* An operating system (Windows, Linux, FreeBSD)
* A modern browser (Chrome, Firefox)

Then download the latest release of MailSlurper. You can find releases at
https://github.com/adampresley/mailslurper-go/releases. Extract the contents
of the ZIP file to any location you like and then run the executable.

**Windows**
* Open Explorer to where you extracted this contents of the ZIP file
* Double-click on *mailslurper.exe*. This will open the console window
* Open your favorite browser and browse to **http://localhost:8080**

**Ubuntu**
* Open a terminal
* Change directory to where you extract the ZIP file
   * *cd /path/to/mailslurper*
* Execute the program
   * *./mailslurper*
* Open your favorite browser and browse to **http://localhost:8080**

