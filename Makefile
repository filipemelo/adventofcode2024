build: 
	@for day in $(shell ls days); do \
		go build -o days/$$day/day$$day days/$$day/main.go
	done

run-day:
	@read -p "Input the day number format (dd): " day; \
		go run ./days/$$day/main.go

run-day-01:
	go run ./days/01/main.go

run-day-02:
	go run ./days/02/main.go

run-day-03:
	go run ./days/03/main.go

clean:
	rm -f days/*/day*

test:
	go test ./tests/... -v