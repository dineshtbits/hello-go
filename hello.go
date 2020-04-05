package main

import (
        "fmt"
        "github.com/dineshtbits/hello/morestrings"
        "github.com/google/go-cmp/cmp"
)

func main() {
   fmt.Println(morestrings.ReverseRunes("Hello World."))
   fmt.Println(cmp.Diff("Hello", "Hell"))
}
