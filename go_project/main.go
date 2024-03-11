package main

import (
	"fmt"
	"go_project/parser"

	"github.com/likexian/whois"
)

func main() {
	out, err := whois.Whois("208.102.51.6")
	if err != nil {
		fmt.Print(err)
	}
	BaseInfo, OrgInfo := parser.GetData(out)
	fmt.Println(BaseInfo, "\n", OrgInfo)

	BaseInfo.Reset()
	OrgInfo.Reset()
}
