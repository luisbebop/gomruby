// +build ignore
package main

import (
	"fmt"
	"os"
)

func main() {
	for _, file := range []string{"parse.y", "keywords", "lex.def"} {
		fi, err := os.Stat("mruby/src/" + file)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%10s - %s (%d)\n", fi.Name(), fi.ModTime(), fi.ModTime().UnixNano())
	}
}
