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
	"os"
	"os/signal"
	"sync"
)

var (
	mu = &sync.RWMutex{}
	// cancelFuncs contains all context cancel functions.
	cancelFuncs []context.CancelFunc
)

// WithCancelSigInt returns a context that is canceled when an interrupt
// signal is received. It is allowed to call WithCancelSigInt multiple times.
func WithCancelSigInt() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	mu.Lock()
	defer mu.Unlock()

	if len(cancelFuncs) == 0 {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)

		go listen(sig)
	}

	cancelFuncs = append(cancelFuncs, cancel)

	return ctx
}

func listen(sig chan os.Signal) {
	<-sig

	signal.Stop(sig)

	cancelContexts()
}

func cancelContexts() {
	mu.Lock()
	defer mu.Unlock()

	for _, cancelFunc := range cancelFuncs {
		cancelFunc()
	}

	cancelFuncs = []context.CancelFunc{}
}
