# logger

> A simple log library

[![License][License-Image]][License-Url] [![ReportCard][ReportCard-Image]][ReportCard-Url] [![GoDoc][GoDoc-Image]][GoDoc-Url]

## Get

``` bash
go get github.com/LyricTian/logger
```

## Usage

```go
package main

import (
	"github.com/LyricTian/logger"
)

func main() {
	log := logger.NewStdLogger(true, true, true, true, true)
	log.Debugf("hello,%s", "world")
	log.Tracef("hello,%s", "world")
	log.Infof("hello,%s", "world")
	log.Errorf("hello,%s", "world")
	log.Fatalf("hello,%s", "world")
}
```

## Features

* Based on the standard library log
* Support output to the console and file

## MIT License

```
Copyright (c) 2017 LyricTian
```

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[ReportCard-Url]: https://goreportcard.com/report/github.com/LyricTian/logger
[ReportCard-Image]: https://goreportcard.com/badge/github.com/LyricTian/logger
[GoDoc-Url]: https://godoc.org/github.com/LyricTian/logger
[GoDoc-Image]: https://godoc.org/github.com/LyricTian/logger?status.svg
