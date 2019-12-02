package main

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/poions/GoDomainScan/model"
)

func main() {
	figure.NewFigure("GoDomainScan", "", true).Print()
	fmt.Println("\n")
	//fmt.Println(os.Args)
	for i := 0; i < len(os.Args); i++ {
		switch {
		case os.Args[i] == "-u":
			model.GetAPIRequestResponseData(os.Args[2])
		case os.Args[i] == "-h":
			fmt.Println("-h/--help \n-u main.go test.com \n-o main.go -u test.com -o /tmp/d.txt")
		case os.Args[i] == "-o":
			fmt.Println("test")
		case os.Args[i] == "--h":
			fmt.Println("-h/--help \n-u main.go test.com \n-o main.go -u test.com -o /tmp/d.txt")
		case os.Args[i] == "--help":
			fmt.Println("-h/--help \n-u main.go test.com \n-o main.go -u test.com -o /tmp/d.txt")
		case os.Args[i] == "help":
			fmt.Println("-h/--help \n-u main.go test.com \n-o main.go -u test.com -o /tmp/d.txt")
		}
	}
}
