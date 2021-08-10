![test](https://github.com/moba1/dotsetup/actions/workflows/test.yml/badge.svg)

# dotsetup

dotsetup is a library for seting up dotfiles.

This library use `sudo` & `curl` command internally.

This library setup dotfiles by combining `task` and finally executing all tasks..
`task` represents execution entity corresponding one line shell command.
For exmaple, `dotsetup.Curl` corresponds to `curl` command.

## Quick examples

This sample executes series of process to

- install `sample-package`
- fetch `index.html` from `example.com`

```go
import (
	"github.com/moba1/dotsetup/v2"
	"log"
)

// set `sample-package` installing task
sp := dotsetup.Package{
	Name: "sample-package"
}
// set fetching index.html to /tmp
c := dotsetup.Curl{
	Args: []string{"-o", "/tmp/index.html", "https://example.com/index.html"}
}
// execute all tasks
// order: sp -> c
s := dotsetup.NewScript([]dotsetup.Task{sp, c})
// enable debug mode
s.Debug = true
if err := s.Execute("sudo password"); err != nil {
	log.Fatal(err)
}
```

## Runnning Tests
```bash
$ go test -v ./...
```

## Task

`task` represents shell command.

### `Package` task

install package.

| Property | type | description |
|:--------:|:----:|:-----------:|
| Name     | string | target package name |

```go
import "github.com/moba1/dotsetup/v2"

// install `sample-package`
p := dotsetup.Package{
	Name: "sample-package"
}
```

### `Curl` task

This task represents `curl` command.

| Property | type | description |
|:--------:|:----:|:-----------:|
| Args     | []string | `curl` command arguments |

```go
import "github.com/moba1/dotsetup/v2"

// execute `curl -o /tmp/sample.txt https://github.com`
c := dotsetup.Curl{
	Args: []string{
		"-o", "/tmp/sample.txt", "https://github.com"
	}
}
```

### `Directory` task

create directory.

| Property | type | description |
|:--------:|:----:|:-----------:|
| Path     | string | directory path |
| Mode     | string | directory mode |

```go
import "github.com/moba1/dotsetup/v2"

// create `/tmp/directory` directory with mode "rwxr-xr-x"
d := dotsetup.Directory{
	Path: "/tmp/directory"
	Mode: "755"
}
```

### `Execute` task

execute shell command.

| Property | type | description |
|:--------:|:----:|:-----------:|
| RawCommands | []dotsetup.ExecuteCommand | shell commands |

```go
import "github.com/moba1/dotsetup/v2"

// execute `sudo -S ls -l`
e := dotsetup.Execute{
	RawCommands: []dotsetup.ExecuteCommand{
		RawCommand: dotsetup.RawCommand{"ls", "-l"},
		DoRoot: true
	},
}
```

### `Link` task

create symbolic link.

| Property | type | description |
|:--------:|:----:|:-----------:|
| Source   | string | source path |
| Destination | string | destination path |
| Force    | set force mode |

```go
import "github.com/moba1/dotsetup/v2"

// put symbolic link from `/dev/null` to `/tmp/null`
l := dotsetup.Link{
	Source: "/dev/null"
	Destination: "/tmp/null"
	Force: true
}
```

## `Os` var
`Os` var is os name.

If runtime is Linux, `Os` var is equal to `/etc/os-release`'s `ID` var.
In other runtime, equal to `runtime.GOOS`.
