builder := go
exe := replacementSimulator
src := main.go

build:
	$(builder) build -o $(exe) $(src)

run:
	$(builder) run $(src)

clean:
	$(builder) clean