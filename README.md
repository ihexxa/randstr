# randstr

[![Build Status](https://travis-ci.org/ihexxa/q-radix.svg?branch=master)](https://travis-ci.org/ihexxa/q-radix)
[![Go Report](https://goreportcard.com/badge/github.com/ihexxa/randstr)](https://goreportcard.com/report/github.com/ihexxa/randstr)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ihexxa/randstr)](https://pkg.go.dev/github.com/ihexxa/randstr)

_Customizable random string generator in Go/Golang._

### Install

Install randstr: `go get github.com/ihexxa/randstr`

### Examples

Document is [here](https://pkg.go.dev/github.com/ihexxa/randstr).
**Note**: This library should NOT be used in security-critical system.

```go
import "github.com/ihexxa/randstr"

var randomString string

// unfixed length (at most 10 alphabets) string generator
unfixedLenAlphabets := randstr.NewAlphabetStr(false, 10)
randomString = unfixedLenAlphabets.Gen()

// fixed length (10 numbers) generator generator
fixedLenNums := randstr.NewNumberStr(true, 10) 
randomString = fixedLenNums.Gen()

// fixed length (5 alphabets or numbers) string generator
fixedLenAlnums := randstr.NewAlnumStr(true, 5)
randomString = fixedLenAlnums.Gen() 

// customized generator 
candidates := []string{"你", "好", "世", "界"}
randStr := randstr.NewRandStr(candidates, true, 8)
randomString = randStr.Gen() 
// output could be: 世世界你你你你好

```
