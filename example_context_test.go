/*
   Copyright The HAPCI Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package signal_test

import (
	"context"
	"fmt"
	"syscall"

	"github.com/hapci/signal"
)

func ExampleNotifyContext() {
	ctx1 := signal.NotifyContext(context.Background())
	ctx2 := signal.NotifyContext(context.Background())

	fmt.Println("Sending interrupt signal")

	err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	if err != nil {
		fmt.Println(err)
	}

	<-ctx1.Done()
	fmt.Println("Context 1 is done")

	<-ctx2.Done()
	fmt.Println("Context 2 is done")
}
