# randstr

[![Build Status](https://travis-ci.org/ihexxa/q-radix.svg?branch=master)](https://travis-ci.org/ihexxa/q-radix)
[![Go Report](https://goreportcard.com/badge/github.com/ihexxa/randstr)](https://goreportcard.com/report/github.com/ihexxa/randstr)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ihexxa/randstr)](https://pkg.go.dev/github.com/ihexxa/randstr)

_Customizable random string generator in Go/Golang._

### Examples

Document is [here](https://pkg.go.dev/github.com/ihexxa/randstr).
**Note**: This library should NOT be used in security-critical system.

```go
import "github.com/ihexxa/randstr"

// customized generator
candidates := []string{"你", "好", "世", "界"}
randStr := randstr.NewRandStr(candidates, true, 8)

fmt.Println(randStr.Gen()) // output could be: 世世界你你你你好
fmt.Println(randStr.Alphabets())
fmt.Println(randStr.Numbers())
fmt.Println(randStr.Alnums())

```
