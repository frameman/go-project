package main

import (
	"fmt"
	"go_project/parser"
	"reflect"
	"strings"

	"github.com/likexian/whois"
)

// beacase the output format for the parser is hard to read so
// i ask gpt to generate a reading the element in the []data struct
func printSliceData(slice interface{}) {
	sliceVal := reflect.ValueOf(slice)
	for i := 0; i < sliceVal.Len(); i++ {
		elem := sliceVal.Index(i)
		fmt.Printf("Element %d: \n", i)
		for j := 0; j < elem.NumField(); j++ {
			field := elem.Type().Field(j)
			value := elem.Field(j).Interface()
			fmt.Printf("  %s: %v\n", field.Name, value)
		}
	}
}

func main() {
	//ip := "90.173.152.199"

	//lacnic
	//ip := "200.3.12.0"

	//ripe
	ip := "84.205.64.0"

	//afrinic
	//ip := "196.216.2.0"

	//apnic
	//ip := "113.125.208.171"

	//arin
	//ip := "155.94.155.220"

	out, err := whois.Whois(ip)
	//server_name := whois.DefaultClient.GetSever(ip)
	//fmt.Print(server_name)
	if err != nil {
		fmt.Print(err)
	} else {

		lines := strings.Split(out, "\n")

		for _, line := range lines {
			if strings.Contains(line, "ripe") {
				RipeBaseInfo, RipeOrgInfo, RipePContact, RipeRole := parser.ParserRipeData(out)
				printSliceData(RipeBaseInfo)
				printSliceData(RipeOrgInfo)
				printSliceData(RipePContact)
				printSliceData(RipeRole)
				break
			} else if strings.Contains(line, "lacnic") {
				IacnicBaseInfo, IacnicNIC := parser.ParserIacnicData(out)
				printSliceData(IacnicBaseInfo)
				printSliceData(IacnicNIC)
				break
			} else if strings.Contains(line, "afrinic") {
				AfrinicBaseInfo, AfrinicOrgInfo, AfrinicNIC := parser.ParserAfrinicData(out)
				printSliceData(AfrinicBaseInfo)
				printSliceData(AfrinicOrgInfo)
				printSliceData(AfrinicNIC)
				break
			} else if strings.Contains(line, "apnic") {
				ApnicBaseInfo, ApnicIRT, ApnicRole, ApnicPerson := parser.ParserApnicData(out)
				printSliceData(ApnicBaseInfo)
				printSliceData(ApnicIRT)
				printSliceData(ApnicRole)
				printSliceData(ApnicPerson)
				break
			} else if strings.Contains(line, "arin") {
				ArinBaseInfo, ArinOrgInfo, ArinContactPAbuse, ArinContactPOrg, ArinOrgNOC, ArinContactPR := parser.ParserArinData(out)
				printSliceData(ArinBaseInfo)
				printSliceData(ArinOrgInfo)
				printSliceData(ArinContactPAbuse)
				printSliceData(ArinContactPOrg)
				printSliceData(ArinOrgNOC)
				printSliceData(ArinContactPR)
				break
			}
		}

	}

}
