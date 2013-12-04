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

Release Notes
-------------

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