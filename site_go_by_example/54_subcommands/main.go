package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable foo logic")
	fooName := fooCmd.String("name", "Mike", "name of foo")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level of bar")

	if len(os.Args) < 2 {
		fmt.Println("`foo` or `bar` subcommand expected")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand `foo`:")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand `bar`:")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", fooCmd.Args())
	default:
		fmt.Printf("command `%s` is not supported - use `foo` or `bar`\n", os.Args[1])
		os.Exit(1)
	}
}
