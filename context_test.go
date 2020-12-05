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

package context

import (
	"context"
	"sync"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInterrupt_MultipleContexts_Parallel(t *testing.T) {
	var ctx1 context.Context
	var ctx2 context.Context
	var ctx3 context.Context
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()

		ctx1 = Interrupt()
		require.NotNil(t, ctx1)
	}()

	go func() {
		defer wg.Done()

		ctx2 = Interrupt()
		require.NotNil(t, ctx2)
	}()

	go func() {
		defer wg.Done()

		ctx3 = Interrupt()
		require.NotNil(t, ctx3)
	}()

	wg.Wait()

	err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	require.NoError(t, err)

	<-ctx1.Done()
	require.Equal(t, context.Canceled, ctx1.Err())

	<-ctx2.Done()
	require.Equal(t, context.Canceled, ctx2.Err())

	<-ctx3.Done()
	require.Equal(t, context.Canceled, ctx3.Err())
}
