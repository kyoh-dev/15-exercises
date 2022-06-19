package main

import (
	"flag"
	"fmt"
	"kyoh-dev/15-exercises/internal/p1"
	"log"
	"os"
	"strconv"

	_ "kyoh-dev/15-exercises/internal/p1"
)

func main() {
	sumCmd := flag.NewFlagSet("sum", flag.ExitOnError)
	if len(os.Args) < 2 {
		fmt.Println("ERR: expected a subcommand: sum")
		os.Exit(1)
	}
	err := sumCmd.Parse(os.Args[2:])
	if err != nil {
		log.Fatal(err)
	}
	nums := sumCmd.Args()
	if nums != nil {
		n := strToIntSlice(&nums)
		s := p1.Sum(n)
		fmt.Printf("SUM: %d\n", s)
	} else {
		fmt.Println("ERR: expected a series of numbers")
		os.Exit(1)
	}
}

func strToIntSlice(s *[]string) []int {
	slc := *s
	retSlc := []int{}
	for _, str := range slc {
		conv, _ := strconv.Atoi(str)
		retSlc = append(retSlc, conv)
	}
	return retSlc
}
