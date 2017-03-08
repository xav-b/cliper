# Cliper [![GoDoc Widget]][walker] [![Go Report Card](https://goreportcard.com/badge/github.com/hackliff/cliper)](https://goreportcard.com/report/github.com/hackliff/cliper) [![Circle CI](https://circleci.com/gh/hackliff/cliper.svg?style=svg)](https://circleci.com/gh/hackliff/cliper)

> Light clipboard manager from the command line

**Cliper** allows you to get back in your clipboard what you copied
earlier. It can become handy when you need to recall a command you
needed a few minutes ago, or go to an url you pasted in the morning for
example.


## Usage

The tool works in client/server mode : a daemon watches the clipboard
and stores its history, while you can query it from the client.

Typically you will want to fire the server as a startup daemon and
forget about it, and only uses the client when needed

**This is only tested on MacOSX for now, but it could work elsewhere**

```Bash
$ # monitor the clipboard
$ cliper \
  -db /tmp/clip.db \  # where to store the data
  -refresh 1s      \  # how long between polling clipboard update
  -reset           \  # start with a fresh history
  watch               # the command to trigger server mode
```

Then go ctrl-c some stuff and interact with your clipboard history.

```Bash
$ # inspect the history
$ cliper \
  -db /tmp/clip.db \  # make sure to point on the same DB (or use default)
  ls
2017/03/06 08:19:05 initializing data backend [driver=sqlite3 path=/tmp/clip.db]
[ 52  ] i'm a unicorn
[ 49  ] how to: conquer the world with Go

$ # copy back the entry you need
$ cliper cp 49
```


## [Installation][releases]

- One liner you can trust: 

```
$ CLIPER_VERSION="0.1.1" PROJECT_URL="https://raw.githubusercontent.com"
$ curl "${GH_CONTENT}"/hackliff/cliper/blob/master/scripts/bootstrap.sh | bash
$ cliper -help
```

- Or DIY:

```Sh
local version="0.1.1"
local platform="darwin-amd64"
local binary="cliper"

curl \
  -ksL \
  -o /usr/local/bin/${binary} \
  https://github.com/hackliff/${binary}/releases/download/v${version}/${binary}-${platform}
  chmod +x /usr/local/bin/${binary}

cliper -help
```

For the cutting edge version (but probably stable), compile from source: `go get -t -u
github.com/hackliff/cliper`

## API Documentation

Check it out on [gowalker][walker], [godoc][GoDoc], or browse it
locally:

```console
$ make godoc
$ $BROWSER localhost:6060/pkg/github.com/hackliff/cliper
```


## Conventions

**cliper** follows some wide-accepted guidelines

* [Semantic Versioning known as SemVer][semver]
* [Git commit messages][commit]


## Licence

Copyright 2017 Xavier Bruhiere.

**cliper** is available under the MIT Licence.

---

<p align="center">
  <img src="https://raw.github.com/hivetech/hivetech.github.io/master/images/pilotgopher.jpg" alt="gopher" width="200px"/>
</p>


[GoDoc]: https://godoc.org/github.com/hackliff/cliper
[walker]: http://gowalker.org/github.com/hackliff/cliper
[GoDoc Widget]: https://godoc.org/hackliff/cliper?status.svg
[releases]: https://github.com/hackliff/cliper/releases

[semver]: http://semver.org
[commit]: https://chris.beams.io/posts/git-commit/
