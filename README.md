![Monocle](http://dsincl12.b-cdn.net/github/monocle-small.png)

[![GoDoc](https://godoc.org/github.com/dsincl12/monocle?status.svg)](https://godoc.org/github.com/dsincl12/monocle)
[![Go Report Card](https://goreportcard.com/badge/github.com/dsincl12/monocle)](https://goreportcard.com/report/github.com/dsincl12/monocle)
[![GitHub license](https://img.shields.io/github/license/dsincl12/monocle.svg)](https://github.com/dsincl12/monocle/blob/master/LICENSE)

# Monocle
Monocle is a fast, efficient and thread-safe unique id generator written in Go. Choose the timestamp resolution and the length of the random string to suit your specific requirements.

## Sortable
The generated values are constructed so that they are always in chronological order when sorted alphabetically.

## Installation

```
go get github.com/dsincl12/monocle
```

## Usage example

```go
import "github.com/dsincl12/monocle"

func main() {
    // create instance with configuration
    m := monocle.New(monocle.Config{
        TimestampResolution:      time.Millisecond,
        NumberOfRandomCharacters: 8,
    })
    
    // generate value
    value := m.Next() // => RSMJ6sUpQxNE1eN

    // parse timestamp from value
    timestamp := m.ParseTimestamp(value) // => 2019-06-03 18:48:34.13 +0000 UTC
    
    // parse random string from value
    randomString := m.ParseRandomString(value) // => pQxNE1eN
}
```

### Benchmarks
```
BenchmarkNext-4          10000000          212 ns/op          32 B/op          3 allocs/op
```

## Credits
Written by [David Sinclair](https://github.com/dsincl12)  
Original fast random string algorithm written by [András Belicza](https://github.com/icza)  

## License

Apache License 2.0

Copyright (c) 2019 David Sinclair

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.