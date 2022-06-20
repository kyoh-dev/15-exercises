package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"kyoh-dev/15-exercises/internal/p1"
)

func main() {
	sumCmd := flag.NewFlagSet("sum", flag.ExitOnError)
	prodCmd := flag.NewFlagSet("product", flag.ExitOnError)
	meanCmd := flag.NewFlagSet("mean", flag.ExitOnError)
	sqrtCmd := flag.NewFlagSet("sqrt", flag.ExitOnError)
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
			fmt.Printf("sum of %v: %d\n", n, s)
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
			fmt.Printf("product of %v: %d\n", n, p)
		} else {
			fmt.Println("ERR: expected a sequence of space-separated integers")
			os.Exit(1)
		}
	case "mean":
		err := meanCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		nums := meanCmd.Args()
		if nums != nil {
			n := p1.StrToIntSlice(&nums)
			m := p1.Mean(n)
			fmt.Printf("mean of %v: %d\n", n, m)
		} else {
			fmt.Println("ERR: expected a sequence of space-separated integers")
			os.Exit(1)
		}
	case "sqrt":
		err := sqrtCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		num := os.Args[2]
		i, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		s := p1.Sqrt(float64(i))
		fmt.Printf("sqrt of %d: %f\n", i, s)
	default:
		fmt.Println("ERR: expected a subcommand: [sum, product, mean, sqrt]")
		os.Exit(1)
	}
}
