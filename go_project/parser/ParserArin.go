package parser

import (
	"go_project/variable"
	"strings"
)

// this is parser with the sever whois server :Arin
func ParserArinData(text string) ([]variable.ArinBaseInfo, []variable.ArinOrgInfo, []variable.ArinContactPAbuse, []variable.ArinContactPOrg, []variable.ArinOrgNOC, []variable.ArinContactPR) {
	Base_info := []variable.ArinBaseInfo{}
	baseInfo := variable.ArinBaseInfo{}

	Org_info := []variable.ArinOrgInfo{}
	orgInfo := variable.ArinOrgInfo{}

	Contact_ORG := []variable.ArinContactPOrg{} // the array store all the struct of ArinContactPOrg
	contactORG := variable.ArinContactPOrg{}

	Contact_PR := []variable.ArinContactPR{} //the array store all the struct of ArinContactPR
	contactPR := variable.ArinContactPR{}

	Contact_Abuse := []variable.ArinContactPAbuse{} //the array store all the struct of ArinContactPAbuse
	contactAbuse := variable.ArinContactPAbuse{}

	Contact_NOC := []variable.ArinOrgNOC{} //the array store all the struct of contactNOC
	contactNOC := variable.ArinOrgNOC{}

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

		//handling the ArinBaseInfo
		case "NetRange":
			baseInfo.NetRange = strings.TrimSpace(part[1])
		case "CIDR":
			baseInfo.CIDR = strings.TrimSpace(part[1])
		case "NetName":
			baseInfo.NetName = strings.TrimSpace(part[1])
		case "NetHandle":
			baseInfo.NetHandle = strings.TrimSpace(part[1])
		case "Parent":
			baseInfo.Parent = strings.TrimSpace(part[1])
		case "NetType":
			baseInfo.NetType = strings.TrimSpace(part[1])
		case "OriginAS":
			baseInfo.OriginAS = strings.TrimSpace(part[1])
		case "Organization":
			baseInfo.Organization = strings.TrimSpace(part[1])
		case "RegDate":
			if baseInfo.NetRange != "" {
				baseInfo.RegDate = strings.TrimSpace(part[1])
			} else {
				orgInfo.RegDate = strings.TrimSpace(part[1])
			}
		case "Updated":
			if baseInfo.Updated == "" {
				baseInfo.Updated = strings.TrimSpace(part[1])
			} else {
				orgInfo.Updated = strings.TrimSpace(part[1])
			}
		case "Comment":
			if baseInfo.NetRange != "" {
				baseInfo.Comment = strings.TrimSpace(part[1])
				if baseInfo.Comment != "" {
					baseInfo.Comment += " " + strings.TrimSpace(part[1])
				}
			} else if orgInfo.OrgName != "" {
				orgInfo.Comment = strings.TrimSpace(part[1])
				if orgInfo.Comment != "" {
					orgInfo.Comment += " " + strings.TrimSpace(part[1])
				}
			}

		case "Ref":
			if baseInfo.NetRange != "" {
				baseInfo.Ref = strings.TrimSpace(part[1])
				Base_info = append(Base_info, baseInfo)
				baseInfo.Reset()
			} else if orgInfo.OrgName != "" {
				orgInfo.Ref = strings.TrimSpace(part[1])

			}

		//handling the ArinOrgInfo
		case "OrgName":
			orgInfo.OrgName = strings.TrimSpace(part[1])
		case "OrgId":
			orgInfo.OrgId = strings.TrimSpace(part[1])
		case "Address":
			orgInfo.Address = strings.TrimSpace(part[1])
		case "City":
			orgInfo.City = strings.TrimSpace(part[1])
		case "StateProv":
			orgInfo.StateProv = strings.TrimSpace(part[1])
		case "PostalCode":
			orgInfo.PostalCode = strings.TrimSpace(part[1])
		case "Country":
			orgInfo.Country = strings.TrimSpace(part[1])
		case "ReferralServer":
			orgInfo.ReferralServer = strings.TrimSpace(part[1])
			Org_info = append(Org_info, orgInfo)
			orgInfo.Reset()

		// Handling the OrgTech ArinOrgNOC RTech
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

		//ArinOrgNOC
		case "ArinOrgNOCHandle":
			contactNOC.ArinOrgNOCHandle = strings.TrimSpace(part[1])

		case "ArinOrgNOCName":
			contactNOC.ArinOrgNOCName = strings.TrimSpace(part[1])

		case "ArinOrgNOCPhone":
			contactNOC.ArinOrgNOCPhone = strings.TrimSpace(part[1])

		case "ArinOrgNOCEmail":
			contactNOC.ArinOrgNOCEmail = strings.TrimSpace(part[1])

		case "ArinOrgNOCRef":
			contactNOC.ArinOrgNOCRef = strings.TrimSpace(part[1])
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
			contactPR.RTechHandle = strings.TrimSpace(part[1])

		case "RTechName":
			contactPR.RTechName = strings.TrimSpace(part[1])

		case "RTechPhone":
			contactPR.RTechPhone = strings.TrimSpace(part[1])

		case "RTechEmail":
			contactPR.RTechEmail = strings.TrimSpace(part[1])

		case "RTechRef":
			contactPR.RTechRef = strings.TrimSpace(part[1])
			Contact_PR = append(Contact_PR, contactPR)
			contactPR.Reset()

		}
	}
	return Base_info, Org_info, Contact_Abuse, Contact_ORG, Contact_NOC, Contact_PR
}
