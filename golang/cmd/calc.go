package main

import (
	"flag"
	"fmt"
	"kyoh-dev/15-exercises/internal/p1"
	"log"
	"os"
)

func main() {
	sumCmd := flag.NewFlagSet("sum", flag.ExitOnError)
	prodCmd := flag.NewFlagSet("product", flag.ExitOnError)
	if len(os.Args) < 2 {
		fmt.Println("ERR: expected a subcommand: [sum, product, mean, sqrt]")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "sum":
		err := sumCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		nums := sumCmd.Args()
		if nums != nil {
			n := p1.StrToIntSlice(&nums)
			s := p1.Sum(n)
			fmt.Printf("SUM: %d\n", s)
		} else {
			fmt.Println("ERR: expected a sequence of space-separated integers")
			os.Exit(1)
		}
	case "product":
		err := prodCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		nums := prodCmd.Args()
		if nums != nil {
			n := p1.StrToIntSlice(&nums)
			p := p1.Product(n)
			fmt.Printf("PRODUCT: %d\n", p)
		} else {
			fmt.Println("ERR: expected a sequence of space-separated integers")
			os.Exit(1)
		}
	default:
		fmt.Println("ERR: expected a subcommand: [sum, product, mean, sqrt]")
		os.Exit(1)
	}
}
