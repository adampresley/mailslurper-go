---
title: Release Notes
layout: default
---

### Release Notes

**2014-06-04**
* MailSlurper now supports storing mail and attachments in one of three engines:
   * SQlite
   * MySQL
   * Microsoft SQL Server 2008+

**2014-06-02**
* Rewrote the mail header and body parsing routine.
* Attachments are now parsed correctly, stored, and can be viewed from the interface
* Body contents are no longer stored with the mail item row. They are retrieved on demain (when clicking on the row)
* Added a favicon

**2014-05-05**
* Mails that contain HTML or are multipart text and HTML now display HTML in the viewer
* Added ability to search the subject and bodies of mails to filter mail list
* Added sorting of mail items
* Addressed a date parsing issue with mails that have the timezone wrapped in parentheses
* Addressed browser resize issue. Layout now is resizable and more responsive
* Removed unneeded code
* Updated several libraries

**2013-12-15**
* Added new interface to the administrator to change settings to *config.json*
* Some code cleanup

**2013-12-12**
* Options can now be configured through settings in a file named *config.json*

**2013-12-11**
* New command line flag **-smtpaddress** allows you to specify address to bind the SMTP server to

**2013-12-06**
* Mails now display in descending order
* SMTP server now writes new mail items to a websocket to update the UI
* Updated version to 1.1

**2013-12-04**
* Initial conversion of the Groovy+Grails version of MailSlurper to Google Go. It is still pretty rough, but the mechanics are there

[Back](index.html)