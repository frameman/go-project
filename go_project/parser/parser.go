package parser

import (
	"go_project/variable"
	"strings"
)

func GetData(text string) (variable.BaseInfo, variable.OrgInfo, []variable.ContactPAbuse, []variable.ContactPOrg, []variable.OrgNOC, []variable.ContactPR) {
	Base_info := variable.BaseInfo{}
	Org_info := variable.OrgInfo{}

	Contact_ORG := []variable.ContactPOrg{} // the array store all the struct of contactPOrg
	contactORG := variable.ContactPOrg{}

	Contact_PR := []variable.ContactPR{} //the array store all the struct of contactPR
	contractPR := variable.ContactPR{}

	Contact_Abuse := []variable.ContactPAbuse{} //the array store all the struct of contactPAbuse
	contactAbuse := variable.ContactPAbuse{}

	Contact_NOC := []variable.OrgNOC{} //the array store all the struct of contactNOC
	contactNOC := variable.OrgNOC{}

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

		//handling the baseInfo
		case "NetRange":
			Base_info.NetRange = strings.TrimSpace(part[1])
		case "CIDR":
			Base_info.CIDR = strings.TrimSpace(part[1])
		case "NetName":
			Base_info.NetName = strings.TrimSpace(part[1])
		case "NetHandle":
			Base_info.NetHandle = strings.TrimSpace(part[1])
		case "Parent":
			Base_info.Parent = strings.TrimSpace(part[1])
		case "NetType":
			Base_info.NetType = strings.TrimSpace(part[1])
		case "OriginAS":
			Base_info.OriginAS = strings.TrimSpace(part[1])
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
		case "Comment":
			Base_info.Comment = strings.TrimSpace(part[1])
			if Base_info.Comment != "" {
				Base_info.Comment += " " + strings.TrimSpace(part[1])
			}

		case "Ref":
			if Base_info.Ref == "" {
				Base_info.Ref = strings.TrimSpace(part[1])
			} else {
				Org_info.Ref = strings.TrimSpace(part[1])
			}

		//handling the OrgInfo
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
		case "Country":
			Org_info.Country = strings.TrimSpace(part[1])
		case "ReferralServer":
			Org_info.ReferralServer = strings.TrimSpace(part[1])

		// Handling the OrgTech OrgNOC RTech
		//OrgTech
		case "OrgTechHandle":
			contactORG.OrgTechHandle = strings.TrimSpace(part[1])

		case "OrgTechName":
			contactORG.OrgTechName = strings.TrimSpace(part[1])

		case "OrgTechPhone":
			contactORG.OrgTechPhone = strings.TrimSpace(part[1])

		case "OrgTechEmail":
			contactORG.OrgTechEmail = strings.TrimSpace(part[1])

		case "OrgTechRef":
			contactORG.OrgTechRef = strings.TrimSpace(part[1])
			Contact_ORG = append(Contact_ORG, contactORG)
			contactORG.Reset()

		//OrgNOC
		case "OrgNOCHandle":
			contactNOC.OrgNOCHandle = strings.TrimSpace(part[1])

		case "OrgNOCName":
			contactNOC.OrgNOCName = strings.TrimSpace(part[1])

		case "OrgNOCPhone":
			contactNOC.OrgNOCPhone = strings.TrimSpace(part[1])

		case "OrgNOCEmail":
			contactNOC.OrgNOCEmail = strings.TrimSpace(part[1])

		case "OrgNOCRef":
			contactNOC.OrgNOCRef = strings.TrimSpace(part[1])
			Contact_NOC = append(Contact_NOC, contactNOC)
			contactNOC.Reset()

		//OrgAbuse
		case "OrgAbuseHandle":
			contactAbuse.OrgAbuseHandle = strings.TrimSpace(part[1])

		case "OrgAbuseName":
			contactAbuse.OrgAbuseName = strings.TrimSpace(part[1])

		case "OrgAbusePhone":
			contactAbuse.OrgAbusePhone = strings.TrimSpace(part[1])

		case "OrgAbuseEmail":
			contactAbuse.OrgAbuseEmail = strings.TrimSpace(part[1])

		case "OrgAbuseRef":
			contactAbuse.OrgAbuseRef = strings.TrimSpace(part[1])
			Contact_Abuse = append(Contact_Abuse, contactAbuse)
			contactAbuse.Reset()

		//RTech
		case "RTechHandle":
			contractPR.RTechHandle = strings.TrimSpace(part[1])

		case "RTechName":
			contractPR.RTechName = strings.TrimSpace(part[1])

		case "RTechPhone":
			contractPR.RTechPhone = strings.TrimSpace(part[1])

		case "RTechEmail":
			contractPR.RTechEmail = strings.TrimSpace(part[1])

		case "RTechRef":
			contractPR.RTechRef = strings.TrimSpace(part[1])
			Contact_PR = append(Contact_PR, contractPR)
			contactAbuse.Reset()

		}
	}
	return Base_info, Org_info, Contact_Abuse, Contact_ORG, Contact_NOC, Contact_PR
}
