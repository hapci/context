<pre>
           _____                      ______
    __________(_)______ _____________ ___  /
    __  ___/_  /__  __ `/_  __ \  __ `/_  / 
    _(__  )_  / _  /_/ /_  / / / /_/ /_  /  
    /____/ /_/  _\__, / /_/ /_/\__,_/ /_/   
                /____/
</pre>

[![Build Status](https://github.com/hapci/signal/workflows/CI/badge.svg)](https://github.com/hapci/signal/actions?query=workflow%3ACI)
[![Coverage Status](https://codecov.io/gh/hapci/signal/branch/master/graph/badge.svg)](https://codecov.io/gh/hapci/signal)
[![Go Report Card](https://goreportcard.com/badge/github.com/hapci/signal)](https://goreportcard.com/report/github.com/hapci/signal)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hapci/signal)](https://pkg.go.dev/github.com/hapci/signal)

__Concurrently safe context handling of OS signals.__

### Installation

```
$ go get github.com/hapci/signal
```

### Getting started

```go
package main

import (
	"context"
	"fmt"
	"syscall"

	"github.com/hapci/signal"
)

func main() {
	ctx1 := signal.NotifyContext(context.Background())
	ctx2 := signal.NotifyContext(context.Background())

	fmt.Println("Sending interrupt signal")

	err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	if err != nil {
		panic(err)
	}

	<-ctx1.Done()
	fmt.Println("Context 1 is done")

	<-ctx2.Done()
	fmt.Println("Context 2 is done")
}
```