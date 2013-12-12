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
* **-wwwport** - Port number to bind to for the web-based administrator. Defaults to 8080
* **-www** - Path to the web administrator directory. Defaults to **www/**

So, for example, to run MailSlurper on different ports, try this.

```bash
$ ./mailslurper -smtpport=2500 -wwwport=8083
```

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

**2013-12-11**
* New command line flag **-smtpaddress** allows you to specify address to bind the SMTP server to

**2013-12-06**
* Mails now display in descending order
* SMTP server now writes new mail items to a websocket to update the UI
* Updated version to 1.1

**2013-12-04**
* Initial conversion of the Groovy+Grails version of MailSlurper to Google Go. It is still pretty rough, but the mechanics are there


License
-------
The MIT License (MIT)
Copyright (c) 2013 Adam Presley

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