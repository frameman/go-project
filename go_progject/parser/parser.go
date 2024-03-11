package parser

import (
	//"go_project/variable"
	"strings"
)

// Reset the struct by dereferencing the pointer and assigning a new struct
func (m *variable.whois_base_info) Reset() {
	*m = variable.base_info{}
}
func (m *variable.whois_org_info) Reset() {
	*m = variable.org_info{}
}

func get_data(text string) {
	variable.whois_base_info.Reset()
	variable.whois_org_info.Reset()
	//split the line
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		/*paring the first part and data
		eg. the result will be
		NetRange:       208.102.0.0 - 208.102.255.255
		it will divide into parts "NetRange" and  "208.102.0.0 - 208.102.255.255"
		and it will save according to the data struct build in the package variable
		*/
		part := strings.SplitN(line, ":", 2)
		switch part[0] {
		case "NetRange":
			variable.whois_base_info.NetRange = strings.TrimSpace(part[1])
		case "OriginAS":
			variable.whois_base_info.OriginAS = strings.TrimSpace(part[1])
		case "NetName":
			variable.whois_base_info.NetName = strings.TrimSpace(part[1])
		case "CIDR":
			variable.whois_base_info.CIDR = strings.TrimSpace(part[1])
		case "Parent":
			variable.whois_base_info.Parent = strings.TrimSpace(part[1])
		case "NetType":
			variable.whois_base_info.NetType = strings.TrimSpace(part[1])
		case "Organization":
			variable.whois_base_info.Organization = strings.TrimSpace(part[1])
		case "RegDate":
			if variable.whois_base_info.RegDate == nil {
				variable.whois_base_info.RegDate = strings.TrimSpace(part[1])
			} else {
				variable.whois_org_info.RegDate = strings.TrimSpace(part[1])
			}
		case "Updated":
			if variable.whois_base_info.Updated == nil {
				variable.whois_base_info.Updated = strings.TrimSpace(part[1])
			} else {
				variable.whois_org_info.Updated = strings.TrimSpace(part[1])
			}
		case "Commen":
			variable.whois_base_info.Commen = strings.TrimSpace(part[1])
		case "Ref":
			if variable.whois_base_info.Ref == nil {
				variable.whois_base_info.Ref = strings.TrimSpace(part[1])
			} else {
				variable.whois_org_info.Ref = strings.TrimSpace(part[1])
			}
		case "OrgName":
			variable.whois_org_info.OrgAbuseName = strings.TrimSpace(part[1])
		case "OrgId":
			variable.whois_org_info.OrgId = strings.TrimSpace(part[1])
		case "Address":
			variable.whois_org_info.Address = strings.TrimSpace(part[1])
		case "City":
			variable.whois_org_info.City = strings.TrimSpace(part[1])
		case "StateProv":
			variable.whois_org_info.StateProv = strings.TrimSpace(part[1])
		case "PostalCode":
			variable.whois_org_info.PostalCode = strings.TrimSpace(part[1])
		case "Comment":
			variable.whois_org_info.Updated = strings.TrimSpace(part[1])
		case "ReferralServer":
			variable.whois_org_info.ReferralServer = strings.TrimSpace(part[1])
		}
	}
}
