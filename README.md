# golog

`golog` is a simple Go Structured Leveled Logging Library.

It is designed to be simple and hold absolutely nodependencies. It provides a subset of the fantastic [zerolog](https://github.com/rs/zerolog) library, and thus is somewhat compatible with it.

## Why Another Logging Library?

We developed one because strongly desired to avoid logging discrepancies. That way we can standardize our log output format and locations.

Also, the choice to copy `zerolog`'s interface implementation was because the stadard library `log` does not provide structured or Leveled logging.

## Show Me The Money

OK, enough talking. The API is simple and straightforward, and a subset of `zerolog`. You can find several logging examples in the `examples` directory.

To run them, just type:

```(bash)
$ go run examples/examples.go
```
