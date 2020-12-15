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

package signal

import (
	"context"
	"fmt"
	"sync"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNotifyContext_MultipleContexts_Parallel(t *testing.T) {
	var (
		ctx1, ctx2, ctx3 context.Context
		ctxGroup         sync.WaitGroup
	)

	ctxGroup.Add(3)

	go func() {
		defer ctxGroup.Done()

		ctx1 = NotifyContext(context.Background())
		require.NotNil(t, ctx1)
	}()

	go func() {
		defer ctxGroup.Done()

		ctx2 = NotifyContext(context.Background())
		require.NotNil(t, ctx2)
	}()

	go func() {
		defer ctxGroup.Done()

		ctx3 = NotifyContext(context.Background())
		require.NotNil(t, ctx3)
	}()

	ctxGroup.Wait()

	err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	require.NoError(t, err)

	<-ctx1.Done()
	require.Equal(t, context.Canceled, ctx1.Err())

	<-ctx2.Done()
	require.Equal(t, context.Canceled, ctx2.Err())

	<-ctx3.Done()
	require.Equal(t, context.Canceled, ctx3.Err())
}

func ExampleNotifyContext() {
	ctx1 := NotifyContext(context.Background())
	ctx2 := NotifyContext(context.Background())

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
