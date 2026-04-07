package ixgo_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/bytes"
	_ "github.com/goplus/ixgo/pkg/context"
	_ "github.com/goplus/ixgo/pkg/fmt"
	_ "github.com/goplus/ixgo/pkg/time"
)

func TestInterpreter_ConcurrentRun1(t *testing.T) {
	source := `
package main

import (
	"fmt"
	"bytes"
)

type T struct {
	*bytes.Buffer
}

func main() {
	t := &T{ bytes.NewBufferString("Hello World") }
	fmt.Println(t)
}
`

	var wg sync.WaitGroup
	numGoroutines := 100
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_, err := ixgo.RunFile("main.go", source, nil, ixgo.SupportMultipleInterp)
			if err != nil {
				t.Errorf("goroutine %d: RunFile failed: %v", id, err)
			}
		}(i)
	}

	wg.Wait()
}

func TestInterpreter_ConcurrentRun2(t *testing.T) {
	source := `
package main

import (
	"fmt"
	"bytes"
)

type T struct {
	*bytes.Buffer
}

func main() {
	t := &T{ bytes.NewBufferString("Hello World") }
	fmt.Println(t)
}
`

	var wg sync.WaitGroup
	numGoroutines := 100
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ctx := ixgo.NewContext(ixgo.SupportMultipleInterp)
			defer ctx.UnsafeRelease()
			pkg, err := ctx.LoadFile("main.go", source)
			if err != nil {
				t.Errorf("goroutine %d: Load failed: %v", id, err)
			}
			interp, err := ctx.NewInterp(pkg)
			if err != nil {
				t.Errorf("goroutine %d: NewInterp failed: %v", id, err)
			}
			defer interp.UnsafeRelease()
			_, err = interp.RunMain()
			if err != nil {
				t.Errorf("goroutine %d: RunMain failed: %v", id, err)
			}
		}(i)
	}

	wg.Wait()
}

func TestInterpreter_ConcurrentRun3(t *testing.T) {
	source := `
package main

import (
	"fmt"
	"bytes"
)

type T struct {
	*bytes.Buffer
}

func main() {
	t := &T{ bytes.NewBufferString("Hello World") }
	fmt.Println(t)
}
`

	ctx := ixgo.NewContext(ixgo.SupportMultipleInterp)
	var wg sync.WaitGroup
	numGoroutines := 100

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_, err := ctx.RunFile("main.go", source, nil)
			if err != nil {
				t.Errorf("goroutine %d: RunFile failed: %v", id, err)
			}
		}(i)
	}

	wg.Wait()
}

func TestInterpreter_ConcurrentRun4(t *testing.T) {
	source := `
package main

import (
	"fmt"
	"bytes"
)

type T struct {
	*bytes.Buffer
}

func main() {
	t := &T{ bytes.NewBufferString("Hello World") }
	fmt.Println(t)
}
`
	ctx := ixgo.NewContext(ixgo.SupportMultipleInterp)
	var wg sync.WaitGroup
	numGoroutines := 100
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			pkg, err := ctx.LoadFile("main.go", source)
			if err != nil {
				t.Errorf("goroutine %d: Load failed: %v", id, err)
			}
			interp, err := ctx.NewInterp(pkg)
			if err != nil {
				t.Errorf("goroutine %d: NewInterp failed: %v", id, err)
			}
			defer interp.UnsafeRelease()
			_, err = interp.RunMain()
			if err != nil {
				t.Errorf("goroutine %d: RunMain failed: %v", id, err)
			}
		}(i)
	}

	wg.Wait()
}

func TestContextCancelRace(t *testing.T) {
	source := `
package main
import (
	"context"
	"time"
)
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	
	select {
	case <-ctx.Done():
		return
	case <-time.After(10 * time.Second):
		return
	}
}
`

	_, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := ixgo.RunFile("ctx.go", source, nil, ixgo.SupportMultipleInterp)
			if err != nil {
				t.Logf("Expected error or timeout: %v", err)
			}
		}()
	}

	time.Sleep(5 * time.Millisecond)
	cancel()

	wg.Wait()
}
