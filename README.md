# Kaoiji
a Bancho written in Golang (our current Bancho)

# Build Status
[![Build Status](https://travis-ci.org/Gigamons/Kaoiji.svg?branch=master)](https://travis-ci.org/Gigamons/Kaoiji)[![CodeFactor](https://www.codefactor.io/repository/github/gigamons/kaoiji/badge/master)](https://www.codefactor.io/repository/github/gigamons/kaoiji/overview/master)[![Appveyor](https://ci.appveyor.com/api/projects/status/github/Gigamons/kaoiji?svg=true&retina=true&branch=master&passingText=master%20-%20OK)](https://ci.appveyor.com/project/Mempler/kaoiji)

# Requirements
* [MariaDB 10.2](https://downloads.mariadb.org/)
* [Redis 4+](https://redis.io/download)
* [Golang 1.10.x (For Building)](https://golang.org/dl/)

# Building
Install the latest Golang version if not already [Here](https://golang.org/dl/).

Download Kaoiji via
```
go get -d -v github.com/Gigamons/Kaoiji
```

Go into the Source code located at \
Windows: `%GOPATH%/src/Gigamons/Kaoiji` \
Unix: `$GOPATH/src/Gigamons/Kaoiji`

Then run \
`go build -v kaoiji.go`

the Compiled binary is located in the Source directory.

# Installation

Download the LATEST pre-compiled binary [here](https://github.com/Gigamons/Kaoiji/releases) or build it by yourself.

run `Kaoiji` or `Kaoiji.exe` once to generate a config file. \
Edit the Config file then rerun `Kaoiji` or `Kaoiji.exe`.

Done, The database structure will be generated automaticly into the given Database.

# Todo
* README Fix english.
