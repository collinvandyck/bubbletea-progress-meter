# Progress Meter

This is a demo of using the bubbletea progress meter.

To run:

    go run ./main.go

You'd implement the Progress interface:

```go
type Progress interface {
	// return a value between 0 and 1
	Done() (float64, error)
}
```

to supply your own percentage.

Demo:

https://github.com/user-attachments/assets/56b95210-c27d-4c64-9e3c-7a7d6f41fdc4

