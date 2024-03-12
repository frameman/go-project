package main

import (
	"fmt"
	"go_project/parser"

	"github.com/likexian/whois"
)

func main() {
	out, err := whois.Whois("107.89.184.71")
	if err != nil {
		fmt.Print(err)
	}
	BaseInfo, OrgInfo, Contact_Abuse, Contact_NOC, Contact_ORG, Contact_PR := parser.GetData(out)
	//fmt.Print(out)
	fmt.Println(Contact_Abuse)
	fmt.Println(Contact_NOC)
	fmt.Println(Contact_ORG)
	fmt.Println(Contact_PR)
	fmt.Println(OrgInfo)
	fmt.Println(BaseInfo.Comment)
	BaseInfo.Reset()
	OrgInfo.Reset()
}
