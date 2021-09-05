.PHONY: repl
repl:
	@go build -o repl.bin ./cmd/repl
	@rlwrap ./repl.bin
