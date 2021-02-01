# Croptop ![Github Actions](https://github.com/sf9v/croptop/workflows/test/badge.svg) [![Coverage Status](https://coveralls.io/repos/github/sf9v/croptop/badge.svg?branch=main)](https://coveralls.io/github/sf9v/croptop?branch=main) [![Go Report Card](https://goreportcard.com/badge/github.com/sf9v/croptop)](https://goreportcard.com/report/github.com/sf9v/croptop)

An easy to use image cropping library built on top of [imaging](https://github.com/disintegration/imaging). This package is just a wrapper to [imaging](https://github.com/disintegration/imaging).

## Installation

```console
$ go get github.com/sf9v/croptop
```

## Example

1. Decode image from a reader.

```go
img, err := croptop.Decode(r)
// handle error
```

2. Specify crop options and encode to a writer.

```go
err = img.Height(height).Width(width).
    OffsetX(offsetX).OffsetY(offsetY).
    Crop().Encode(w)
// handle error
```

## License

MIT