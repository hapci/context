<pre>
                        __            __ 
      _________  ____  / /____  _  __/ /_
     / ___/ __ \/ __ \/ __/ _ \| |/_/ __/
    / /__/ /_/ / / / / /_/  __//  // /_  
    \___/\____/_/ /_/\__/\___/_/|_|\__/
    
</pre>

[![Build Status](https://github.com/hapci/context/workflows/CI/badge.svg)](https://github.com/hapci/context/actions?query=workflow%3ACI)
[![Coverage Status](https://codecov.io/gh/hapci/context/branch/master/graph/badge.svg)](https://codecov.io/gh/hapci/context)
[![Go Report Card](https://goreportcard.com/badge/github.com/hapci/context)](https://goreportcard.com/report/github.com/hapci/context)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hapci/context)](https://pkg.go.dev/github.com/hapci/context)

__Concurrently safe context handling of OS signals.__

## Getting started

```go
package main

import (
	"fmt"
	"syscall"
	
	"github.com/hapci/context"
)

func main() {
    ctx1 := context.WithCancelSigInt()
    ctx2 := context.WithCancelSigInt()
    
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