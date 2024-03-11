package main

import (
	"fmt"

	"github.com/likexian/whois"
)

func main() {
	out, err := whois.Whois("208.102.51.6")
	if err != nil {
		fmt.Print(err)
	}
	parser.get_data(out)
	fmt.Print(variable.whois_base_info.NetRange)
}
