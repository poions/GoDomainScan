package main

import (
	"fmt"
	"os"

	"github.com/poions/GoDomainScan/model"
)

func main() {
	fmt.Println(os.Args)
	for i := 0; i < len(os.Args); i++ {
		//fmt.Println(os.Args[i])
		model.GetApiRequestResponseData(os.Args[i])
	}
}
