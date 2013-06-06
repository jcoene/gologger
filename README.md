# Logger

A simple stdout logger with levels.

## Usage example

```go
log := *logger.NewLogger(logger.LOG_LEVEL_INFO, "myapp")

log.Debug "too much information"      // no output
log.Info "things are going well"      // [info]  (myapp) things are going well
log.Warn "it's getting hot in here"   // [warn]  (myapp) it's getting hot in here
log.Error "it got too hot in here"    // [error] (myapp) it got too hot in here
log.Fatal "the system has melted"     // [fatal] (myapp) the system has melted
```

## License

MIT license, see LICENSE for more information.
