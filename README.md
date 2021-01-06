# randstr

[![Build Status](https://travis-ci.org/ihexxa/q-radix.svg?branch=master)](https://travis-ci.org/ihexxa/q-radix)
[![Go Report](https://goreportcard.com/badge/github.com/ihexxa/randstr)](https://goreportcard.com/report/github.com/ihexxa/randstr)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ihexxa/randstr)](https://pkg.go.dev/github.com/ihexxa/randstr)

_Customizable random string generator in Go/Golang._

### Examples

Document is [here](https://pkg.go.dev/github.com/ihexxa/randstr).

**Note**: This library should NOT be used in security-critical system.

# Generate random strings with numbers/alphabets/alnums/candidates:

```go
import "github.com/ihexxa/randstr"

// customized generator
candidates := []string{"你", "好", "世", "界"}
randStr := randstr.New(candidates)

fmt.Println(randStr.Gen()) // output could be: 世世界你你你你好
fmt.Println(randStr.Alphabets())
fmt.Println(randStr.Numbers())
fmt.Println(randStr.Alnums())

```

# Generate fixed-length random strings:

```go
import "github.com/ihexxa/randstr"

candidates := []string{"你", "好", "世", "界"}
randStr := randstr.New(candidates)
randStr.SetLenFixed(true) // set sample size fixed
randStr.SetLenMax(8) // set sample size to 8
// above 3 lines can be replaced with randstr.NewRandstr(candidates, true, 8)

fmt.Println(randStr.Gen())
...

```

# Generate random slices:

```go
import "github.com/ihexxa/randstr"

randStr := randstr.New([]string{"你", "好", "世", "界"})

fmt.Println(randStr.GenSlice())
fmt.Println(randStr.AlphabetSlice())
fmt.Println(randStr.NumberSlice())
fmt.Println(randStr.AlnumSlice())

```

### License

MIT
