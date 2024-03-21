package main

import (
	"fmt"
	"go_project/parser"
	"strings"

	"github.com/likexian/whois"
)

func main() {
	//lacnic
	//ip := "200.3.12.0"

	//ripe
	//ip := "84.205.64.0"

	//afrinic
	//ip := "196.216.2.0"

	//apnic
	ip := "113.125.208.171"

	//arin
	//ip := "155.94.155.220"
	out, err := whois.Whois(ip)
	if err != nil {
		fmt.Print(err)
	} else {

		lines := strings.Split(out, "\n")
	loop:
		for _, line := range lines {
			if strings.Contains(line, "RIPE") {
				RipeBaseInfo, RipeOrgInfo, RiprPContact := parser.ParserRipeData(out)
				fmt.Printf("%+v \n", RipeBaseInfo)
				fmt.Printf("%+v \n", RipeOrgInfo)
				fmt.Printf("%+v \n", RiprPContact)
				break loop
			} else if strings.Contains(line, "lacnic") {
				IacnicBaseInfo, IacnicNIC := parser.ParserIacnicData(out)
				fmt.Printf("%+v \n", IacnicBaseInfo)
				fmt.Printf("%+v \n", IacnicNIC)
				break loop
			} else if strings.Contains(line, "AfriNIC") {
				AfrinicBaseInfo, AfrinicOrgInfo, AfrinicNIC := parser.ParserAfrinicData(out)
				fmt.Printf("%+v \n", AfrinicBaseInfo)
				fmt.Printf("%+v \n", AfrinicOrgInfo)
				fmt.Printf("%+v \n", AfrinicNIC)
				break loop
			} else if strings.Contains(line, "apnic") {
				ApnicBaseInfo, ApnicIRT, ApnicRole, ApnicPerson := parser.ParserApnicData(out)
				fmt.Printf("%+v \n ", ApnicBaseInfo)
				fmt.Printf("%+v \n", ApnicIRT)
				fmt.Printf("%+v \n", ApnicRole)
				fmt.Printf("%+v \n", ApnicPerson)
				break loop
			} else if strings.Contains(line, "ARIN") {
				ArinBaseInfo, ArinOrgInfo, ArinContactPAbuse, ArinContactPOrg, ArinOrgNOC, ArinContactPR := parser.ParserArinData(out)
				fmt.Printf("%+v \n ", ArinBaseInfo)
				fmt.Printf("%+v \n", ArinOrgInfo)
				fmt.Printf("%+v \n", ArinContactPAbuse)
				fmt.Printf("%+v \n", ArinContactPOrg)
				fmt.Printf("%+v \n", ArinOrgNOC)
				fmt.Printf("%+v \n", ArinContactPR)
				break loop
			}
		}

	}
	//fmt.Println(out)
}
