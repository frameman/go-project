package parser

import (
	"go_project/variable"
	"strings"
)

func ParserAfrinicData(text string) ([]variable.AfrinicBaseInfo, []variable.AfrinicOrgInfo, []variable.AfrinicNIC) {
	baseInfo := variable.AfrinicBaseInfo{}
	Base_Info := []variable.AfrinicBaseInfo{}

	orgInfo := variable.AfrinicOrgInfo{}
	Org_Info := []variable.AfrinicOrgInfo{}

	personInfo := variable.AfrinicNIC{}
	Person_Info := []variable.AfrinicNIC{}

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		part := strings.SplitN(line, ":", 2)
		switch part[0] {

		//parser baseInfo
		case "inetnum":
			baseInfo.Inetnum = strings.TrimSpace(part[1])

		case "netname":
			baseInfo.Netname = strings.TrimSpace(part[1])

		case "descr":
			baseInfo.Descr = strings.TrimSpace(part[1])

		case "country":
			if baseInfo.Inetnum != "" {
				baseInfo.Country = strings.TrimSpace(part[1])
			} else if orgInfo.Organisation != "" {
				orgInfo.Country = strings.TrimSpace(part[1])
			}

		case "org":
			baseInfo.Org = strings.TrimSpace(part[1])

		case "admin-c":
			if baseInfo.Inetnum != "" {
				baseInfo.Admin_c = strings.TrimSpace(part[1])
			} else if orgInfo.Organisation != "" {
				orgInfo.Admin_c = strings.TrimSpace(part[1])
			}

		case "tech-c":
			if baseInfo.Inetnum != "" {
				baseInfo.Tech_c = strings.TrimSpace(part[1])
			} else if orgInfo.Organisation != "" {
				orgInfo.Tech_c = strings.TrimSpace(part[1])
			}

		case "status":
			baseInfo.Status = strings.TrimSpace(part[1])

		case "mnt-by":
			if baseInfo.Inetnum != "" {
				baseInfo.Mnt_by = strings.TrimSpace(part[1])
			}
			if orgInfo.Organisation != "" {
				orgInfo.Mnt_by = strings.TrimSpace(part[1])
			}
			if personInfo.Person != "" {
				personInfo.Mnt_by = strings.TrimSpace(part[1])
			}

		case "mnt-lower":
			baseInfo.Mnt_lower = strings.TrimSpace(part[1])

		case "mnt-domains":
			baseInfo.Mnt_domains = strings.TrimSpace(part[1])

		case "mnt-routes":
			baseInfo.Mnt_routes = strings.TrimSpace(part[1])

		case "mnt-irt":
			baseInfo.Mnt_irt = strings.TrimSpace(part[1])

		case "source":
			if baseInfo.Inetnum != "" {
				baseInfo.Source = strings.TrimSpace(part[1])
			}
			if orgInfo.Organisation != "" {
				orgInfo.Source = strings.TrimSpace(part[1])
				Org_Info = append(Org_Info, orgInfo)
				orgInfo.Reset()
			}
			if personInfo.Person != "" {
				personInfo.Source = strings.TrimSpace(part[1])
				Person_Info = append(Person_Info, personInfo)
				personInfo.Reset()
			}

		case "parent":
			baseInfo.Parent = strings.TrimSpace(part[1])
			Base_Info = append(Base_Info, baseInfo)
			baseInfo.Reset()

			//parser organisation
		case "organisation":
			orgInfo.Organisation = strings.TrimSpace(part[1])

		case "org-name":
			orgInfo.Org_name = strings.TrimSpace(part[1])

		case "org-type":
			orgInfo.Org_type = strings.TrimSpace(part[1])

		case "address":
			if orgInfo.Organisation != "" {
				if orgInfo.Address != "" {
					orgInfo.Address += " " + strings.TrimSpace(part[1])
				} else {
					orgInfo.Address = strings.TrimSpace(part[1])
				}
			}
			if personInfo.Person != "" {
				if personInfo.Address != "" {
					personInfo.Address += " " + strings.TrimSpace(part[1])
				} else {
					personInfo.Address = strings.TrimSpace(part[1])
				}
			}

		case "phone":
			if orgInfo.Organisation != "" {
				if orgInfo.Phone != "" {
					orgInfo.Phone += " / " + strings.TrimSpace(part[1])
				} else {
					orgInfo.Phone = strings.TrimSpace(part[1])
				}
			}
			if personInfo.Person != "" {
				if personInfo.Phone != "" {
					personInfo.Phone += "  / " + strings.TrimSpace(part[1])
				} else {
					personInfo.Phone = strings.TrimSpace(part[1])
				}
			}

		case "mnt-ref":
			if orgInfo.Organisation != "" {
				if orgInfo.Mnt_ref != "" {
					orgInfo.Mnt_ref += " / " + strings.TrimSpace(part[1])
				} else {
					orgInfo.Mnt_ref = strings.TrimSpace(part[1])
				}
			}

		case "remarks":
			if orgInfo.Organisation != "" {
				if orgInfo.Remarks != "" {
					orgInfo.Remarks += " " + strings.TrimSpace(part[1])
				} else {
					orgInfo.Remarks = strings.TrimSpace(part[1])
				}
			}
			//parser nic person
		case "person":
			personInfo.Person = strings.TrimSpace(part[1])

		case "nic-hdl":
			personInfo.Nic_hdl = strings.TrimSpace(part[1])

		}

	}
	return Base_Info, Org_Info, Person_Info
}
