package main

import (
	"flag"
	"fmt"
	"os"
	"qualabs/src"
	"strings"
	"time"
)

var (
	source    *string
	operation *string
)

func init() {
	source = flag.String("source", "./data", "data directory")
	operation = flag.String("option", "", "set option type to run, options: A or B")
	flag.Usage = func() {
		info := fmt.Sprintln("\n[+] Qualabs Challenge\n\tOperations:\n\t\tA (chanllenge A)\n\t\tB (challenge B)\nExample: go run . -option A")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "%s", info)
	}
}

func main() {
	flag.Parse()
	user := src.NewUser()
	start := time.Now()
	fmt.Println("[*] Loading data...")
	if !user.LoadData(*source) {
		panic(fmt.Errorf("[-] Not found data directory"))
	}
	fmt.Printf("[*] Data loaded time taken %d ms\n", time.Now().Sub(start).Milliseconds())

	switch strings.ToUpper(*operation) {
	case "A":
		resp, err := user.SolutionA();
		if  err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", resp)
		break
	case "B":
		fmt.Printf("%s\n", user.SolutionB())
		break
	default:
		fmt.Println("[-] Operation not selected")
	}
}
