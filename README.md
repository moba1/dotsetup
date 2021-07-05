# dotsetup

dotsetup is a library for seting up dotfiles.

## Quick examples

```go
import (
	"github.com/moba1/dotsetup"
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
s := dotsetup.NewScript([]dotsetup.Command{sp, c})
if err := s.Execute(); err != nil {
	log.Fatal(err)
}
```

## Runnning Tests
```bash
$ sudo docker-compose -f test/docker-compose.yml up
```
