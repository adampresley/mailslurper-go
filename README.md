MailSlurper-Go
==============

Simple mail SMTP server that slurps mail into oblivion! Useful only for local development
MailSlurper runs on port 8000 (though it can be changed) and listens for outgoing mail send
requests. When one is received the mail item is stored in a local SQLite database.

To view the mail items that were slurped up you can open your web browser and go
to the address *http://localhost:8080* (note that this port can be changed too).

Requirements (to build)
-----------------------
* Google Go (http://golang.org/)

Requirements (to run)
---------------------
* An operating system (Windows, Mac OS X, Linux, FreeBSD)
* A modern browser (Chrome, Firefox)

How to Build
------------
Assuming you have Go installed on your system and your GOPATH environment
variable setup, place this source code, the whole **mailslurper** folder,
in *$GOPATH/src/github.com/adampresley*.

Once you've done this open up a terminal and execute the following:

```bash
$ go install github.com/adampresley/mailslurper
```

Executing the above will create an executable for your OS/platform
into the *$GOPATH/bin* folder. Here you may want to create a folder
somewhere easily accessible. Copy the following items to this new folder.

* *$GOPATH/bin/mailslurper* (mailslurper.exe on Windows)
* *$GOPATH/src/github.com/adampresley/mailslurper/www*

How to Run
----------
From a terminal:

* Windows: mailslurper.exe
* Linux: ./mailslurper

To see what options are available on the terminal execute the following:

```bash
$ ./mailslurper -help
```

The following options are available.

* **-smtpport** - Port number to bind to for the SMTP server. Defaults to 8000
* **-smtpaddress** - Address to bind the SMTP server to. Defaults to **127.0.0.1**
* **-wwwport** - Port number to bind to for the web-based administrator. Defaults to 8080
* **-www** - Path to the web administrator directory. Defaults to **www/**

So, for example, to run MailSlurper on different ports, try this.

```bash
$ ./mailslurper -smtpport=2500 -wwwport=8083
```

Configuration
-------------
MailSlurper can be configured by providing settings in a file called **config.json**.
This is a text-based file with a JSON structure in it containing four configuration
settings. It looks like this.

```javascript
{
	"www": "www/",
	"wwwPort": 8080,
	"smtpAddress": "127.0.0.1",
	"smtpPort": 8000
}
```

* **www** - Path to the web administrator directory.
* **wwwPort** - Port number to bind to for the web-based administrator.
* **smtpAddress** - Address to bind the SMTP server to.
* **smtpPort** - Port number to bind to for the SMTP server.

Please note that these provide MailSlurper the settings it needs to run and the file
must be configured properly for the application to function. Also note that if you
provide command line flag settings when running the server these configuration
values will be superceded by the command line flags.

Documentation
-------------
Wanna see the documentation? Open up a terminal and try the following (Linux. Windows will vary slightly).

```bash
$ cd $GOPATH
$ godoc -http=:6060
```

Then open up your favorite browser to *http://localhost:6060* and you will see a Go page.
Click on the button named **Packages** at the top, and you will be presented with
a list of packages. Find **github.com** and under that you will find the path **adampresley/mailslurper**.
The package documentation is all there.

Release Notes
-------------

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

Copyright Information
---------------------

Mailslurper makes use of many libraries and tools. They are copyright of their respective owners.

* jQuery - Copyright 2014 The jQuery Foundation
* jQuery UI - Copyright 2014 The jQuery Foundation
* Ractive - Copyright 2012-2014 Rich Harris
* Bootstrap - Copyright 2014 Twitter
* RequireJS - Copyright 2010-2013, The Dojo Foundation
* jQuery BlockUI - Copyright 2007-2009 M. Alsup
* jQuery UI Layout - Copyright 2013 Kevin Dalman
* Moment.js - Copyright 2011-2014 Tim Wood, Iskren Chernev, Moment.js contributors
* Google Go - Copyright 2012 The Go Authors
* go-sqlite3 - Copyright 2012-2014 Yasuhiro Matsumoto
* Gorilla Web Toolkit - Copyright 2012 Rodrigo Moraes
* MailSlurper logo uses:
	* Go gopher - Created by and copyright Renee French
	* Mail icon copyright David Hopkins, http://semlabs.co.uk

License
-------
The MIT License (MIT)
Copyright (c) 2013-2014 Adam Presley

Permission is hereby granted, free of charge, to any person obtaining a copy of this
software and associated documentation files (the "Software"), to deal in the Software
without restriction, including without limitation the rights to use, copy, modify,
merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or
substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE
FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.