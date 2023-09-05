[![Go Report Card](https://goreportcard.com/badge/github.com/hedarikun/rogu)](https://goreportcard.com/report/github.com/hedarikun/rogu)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# Introduction
This is yet another simple logging library written in go to fit what I need and the style I like. yes, there might be a better option out there but what this includes fits me more <3

# Installation
```
go get github.com/hedarikun/rogu
```

# Features
- [x] Having the ability to disable/enable some types of logging as well as define the output using `io.Writer` interface
- [x] Supports Optional Dates
- [x] Supports Optional Log Stack Trace (it tells you from where the log is being called)
- [x] Supports being able to create a Log with custom type 
- [ ] Supports Coloring of messages for better viewing using different methods depending on the platform
- [ ] Supporting multiple `io.Writer` interfaces to able to log to console and to a seperate file

# Example
```go
package main
import (
  "github.com/hedarikun/rogu"
)

func main() {
  r := rogu.New()
  r.Log("this is a Log message")
  r.Error("this is an Error message")
  r.Warn("this is a Warning message")
}
```
the result of the above code is
```
[LOG][main.go:8 - main][2023-09-05 10:22:05 Tue]: this is a Log message
[ERROR][main.go:9 - main][2023-09-05 10:22:05 Tue]: this is an Error message
[WARN][main.go:10 - main][2023-09-05 10:22:05 Tue]: this is a Warning message
```
> [!NOTE]
> By default showing log stack and current date is on by default but you can change this behavior

it supports creating some custom made tags for the logging like in the following example
```go
package main
import (
  "github.com/hedarikun/rogu"
)

func main() {
  r := rogu.New()
  notify := r.Logger("NOTIFY")
  notify("this is a notify log!!")
}
```
it will show as 
```
[NOTIFY][main.go:9 - main][2023-09-05 11:29:05 Tue]: this is a notify log!!
```
