run:
	go run .

bench:
	go test . -bench .

test:
	go test -v -race .
