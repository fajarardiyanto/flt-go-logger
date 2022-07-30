### Go Module Logger
Logger modules
Faltar Logger use [Logger v1.8.1](https://github.com/sirupsen/logrus) as logging backend. 

### Installation
```sh
go get github.com/fajarardiyanto/flt-go-logger
```

###### Upgrading to the latest version
```sh
go get -u github.com/fajarardiyanto/flt-go-logger
```

###### Upgrade or downgrade with tag version if available
```sh
go get -u github.com/fajarardiyanto/flt-go-logger@v1.0.0
```

### Usage
```go
package main

import (
	"github.com/fajarardiyanto/flt-go-logger/lib"
)

type Message struct {
	msg string
}

func main() {
	logger := lib.NewLib()
	logger.Init("Testing modules")
	logger.Debug("Lorem Ipsum is simply dummy text of the printing and typesetting %s.", "industry")
	logger.Info(map[string]interface{}{
		"name": "Flt Go Logger",
	})
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