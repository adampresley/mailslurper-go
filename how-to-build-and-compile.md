---
title: How To Build and Compile
layout: default
---

### Build and Compile

If you are more adventurous and wish to compile this yourself then this
section is for you. Assuming you have Go installed on your system and
your GOPATH environment variable setup, place this source code, the whole **mailslurper** folder,
in *$GOPATH/src/github.com/adampresley*. Please note that the name of the folder must be
**mailslurper**. Not **mailslurper-go** or anything else. For more information on setting up
Google Go visit http://golang.org/doc/install.

Once you've done this open up a terminal and execute the following:

```bash
$ go get
$ go install github.com/adampresley/mailslurper
```

Executing the above will download and compile any dependencies and create
an executable for your OS/platform into the *$GOPATH/bin* folder. Here
you may want to create a folder somewhere easily accessible. Copy the
following items to this new folder.

* *$GOPATH/bin/mailslurper* (mailslurper.exe on Windows)
* *$GOPATH/src/github.com/adampresley/mailslurper/www*
* *$GOPATH/src/github.com/adampresley/mailslurper/config.json*

#### To Run

If you don't want to build, but instead just want to run from source
then the steps are simplier. You still need to make sure your GOPATH
is setup, and then ensure you have downloaded and compiled dependencies.
For example if your your Go source lives at **~/me/go**, you would do:

```bash
$ export GOPATH=~/me/go
$ cd ~/me/go
$ go get
$ cd ./src/github.com/adampresley/mailslurper
$ go run ./mailslurper.go
```

And that's about it.


[Back](index.html)