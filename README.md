### Go Module Logger
Logger modules
Faltar Logger use [Logger v1.8.1](https://github.com/sirupsen/logrus) as logging backend. 

### Installation
```sh
go get gitlab.com/fajardiyanto/flt-go-logger
```

###### Upgrading to the latest version
```sh
go get -u gitlab.com/fajardiyanto/flt-go-logger
```

###### Upgrade or downgrade with tag version if available
```sh
go get -u gitlab.com/fajardiyanto/flt-go-logger@v1.0.0
```

### Usage
```go
package main

import (
	"gitlab.com/fajardiyanto/flt-go-logger/lib"
)

type Message struct {
	msg string
}

func main() {
	logger := lib.NewLib().Init()
	logger.SetFormat("text").SetReportCaller(true)
	logger.Info(Message{
		msg: "Also working with map struct",
	})
	logger.Info("Simple testing logger")
}

```

#### Run Example
```sh
make run
```

#### Tips
Maybe it would be better to do some basic code scanning before pushing to the repository.
```sh
# for *.nix users just run gosec.sh
# curl is required
# more information https://github.com/securego/gosec
make scan
```