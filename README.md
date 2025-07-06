# Progress Meter

This is a demo of using the bubbletea progress meter.

## Example 1

    go run cmd/ex1/main.go

The first example works by writing Go code to feed the widget a percentage. To do that you need to implement the Progress interface:

```go
type Progress interface {
	// return a value between 0 and 1
	Done() (float64, error)
}
```

In the following video I'm using a func with a goroutine to generate the percentages:

```go
var mut sync.Mutex
var value float64
go func() {
    for i := range 100 {
        time.Sleep(10 * time.Millisecond)
        mut.Lock()
        value = float64(i+1) / 100
        mut.Unlock()
    }
}()
progMeter := newMeter(ProgressFunc(func() (float64, error) {
    mut.Lock()
    defer mut.Unlock()
    return value, nil
}))
```

https://github.com/user-attachments/assets/56b95210-c27d-4c64-9e3c-7a7d6f41fdc4

## Example 2

    go run cmd/ex2/main.go

The second example works by parsing out floats from stdin. As long as your shell pipeline generates numbers between 0 and 1 it will update the widget.

https://github.com/user-attachments/assets/e0da8966-dc26-448f-8183-7d769f362239

