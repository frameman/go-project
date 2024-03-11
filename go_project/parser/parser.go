package parser

import (
	"go_project/variable"
	"strings"
)

func GetData(text string) (variable.BaseInfo, variable.OrgInfo) {
	Base_info := variable.BaseInfo{}
	Org_info := variable.OrgInfo{}
	//split the line
	lines := strings.Split(text, "\n")
	/*paring the first part and data
	eg. the result will be
	NetRange:       208.102.0.0 - 208.102.255.255
	it will divide into parts "NetRange" and  "208.102.0.0 - 208.102.255.255"
	and it will save according to the data struct build in the package variable
	*/
	for _, line := range lines {
		part := strings.SplitN(line, ":", 2)
		switch part[0] {
		case "NetRange":
			Base_info.NetRange = strings.TrimSpace(part[1])
		case "OriginAS":
			Base_info.OriginAS = strings.TrimSpace(part[1])
		case "NetName":
			Base_info.NetName = strings.TrimSpace(part[1])
		case "CIDR":
			Base_info.CIDR = strings.TrimSpace(part[1])
		case "Parent":
			Base_info.Parent = strings.TrimSpace(part[1])
		case "NetType":
			Base_info.NetType = strings.TrimSpace(part[1])
		case "Organization":
			Base_info.Organization = strings.TrimSpace(part[1])
		case "RegDate":
			if Base_info.RegDate == "" {
				Base_info.RegDate = strings.TrimSpace(part[1])
			} else {
				Org_info.RegDate = strings.TrimSpace(part[1])
			}
		case "Updated":
			if Base_info.Updated == "" {
				Base_info.Updated = strings.TrimSpace(part[1])
			} else {
				Org_info.Updated = strings.TrimSpace(part[1])
			}
		case "Commen":
			Base_info.Commen = strings.TrimSpace(part[1])
		case "Ref":
			if Base_info.Ref == "" {
				Base_info.Ref = strings.TrimSpace(part[1])
			} else {
				Org_info.Ref = strings.TrimSpace(part[1])
			}
		case "OrgName":
			Org_info.OrgName = strings.TrimSpace(part[1])
		case "OrgId":
			Org_info.OrgId = strings.TrimSpace(part[1])
		case "Address":
			Org_info.Address = strings.TrimSpace(part[1])
		case "City":
			Org_info.City = strings.TrimSpace(part[1])
		case "StateProv":
			Org_info.StateProv = strings.TrimSpace(part[1])
		case "PostalCode":
			Org_info.PostalCode = strings.TrimSpace(part[1])
		case "Comment":
			Org_info.Updated = strings.TrimSpace(part[1])
		case "ReferralServer":
			Org_info.ReferralServer = strings.TrimSpace(part[1])
			break
		}
	}
	return Base_info, Org_info
}
