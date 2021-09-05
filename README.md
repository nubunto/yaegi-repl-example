# Yaegi as your REPL

I had a tough time using Yaegi as a REPL to my project. I was able to (mostly) sort it out.
You can use this as an example.

## Usage 

You need `rlwrap` installed.
If you have it, run `make repl`.

Otherwise, just `go build ./cmd/repl` and run that directly.

Either way you should see a REPL pop up. Test it out with:

```go
> a := foo.Capability.DoSomething()
> b := bar.Capability.DoSomething()
> a == b
> 1 + 1
```

As you can see, `foo` `bar`s `Capability` struct is instantiated and automatically imported by the REPL.

## Why?

I like REPLs. I like Go. Why not both?

## Takeaways

- Yaegi has some [limitations](https://github.com/traefik/yaegi/issues/276), so don't expect everything to work perfectly
- I couldn't make `yaegi extract` work the way I wanted to. I've did it by hand, as you can check in `cmd/repl/imports.go`. Maybe there's a better way... feel free to reach out through the issues or, even better, with a PR
- I would pretty much prefer a Go tool REPL
- I prefer to use REPLs to quickly test some API that I'm building, or just running code for "bookkeeping" tasks related to the application I'm developing. This looks like a good approach for the latter, but not the former. Maybe if `yaegi extract` worked we could have both? Maybe I want a [task runner](https://github.com/markbates/grift)? Yes, this is rambling
- I *really* want the Go tool to have a REPL
