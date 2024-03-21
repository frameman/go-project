package parser

import (
	"go_project/variable"
	"strings"
)

func ParserRipeData(text string) ([]variable.RipeBaseInfo, []variable.RipeOrgInfo, []variable.RipePContact) {
	baseInfo := variable.RipeBaseInfo{}
	Base_Info := []variable.RipeBaseInfo{}

	orgInfo := variable.RipeOrgInfo{}
	Org_Info := []variable.RipeOrgInfo{}

	personInfo := variable.RipePContact{}
	Person_Info := []variable.RipePContact{}

	roleInfo := variable.RipeRole{}
	Role_Info := []variable.RipeRole{}

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		part := strings.SplitN(line, ":", 2)
		switch part[0] {

		//parser baseInfo
		case "inetnum":
			baseInfo.Inetnum = strings.TrimSpace(part[1])

		case "netname":
			baseInfo.Netname = strings.TrimSpace(part[1])

		case "country":
			if baseInfo.Inetnum != "" {
				baseInfo.Country = strings.TrimSpace(part[1])

			} else if orgInfo.Organisation != "" {
				orgInfo.Country = strings.TrimSpace(part[1])
			}

		case "admin-c":
			if baseInfo.Inetnum != "" {
				baseInfo.Admin_c = strings.TrimSpace(part[1])
			} else if orgInfo.Organisation != "" {
				if orgInfo.Admin_c != "" {
					orgInfo.Admin_c += " " + strings.TrimSpace(part[1])
				}
				orgInfo.Admin_c = strings.TrimSpace(part[1])
			}

		case "tech-c":
			baseInfo.Tech_c = strings.TrimSpace(part[1])

		case "status":
			baseInfo.Status = strings.TrimSpace(part[1])

		case "org":
			baseInfo.Org = strings.TrimSpace(part[1])

		case "mnt-by":
			if baseInfo.Inetnum != "" {
				baseInfo.Mnt_by = strings.TrimSpace(part[1])
			}
			if orgInfo.Organisation != "" {
				if orgInfo.Mnt_by != "" {
					orgInfo.Mnt_by += " " + strings.TrimSpace(part[1])
				} else {
					orgInfo.Mnt_by = strings.TrimSpace(part[1])
				}
			}
			if personInfo.Person != "" {
				personInfo.Mnt_by = strings.TrimSpace(part[1])
			}
			if roleInfo.Role != "" {
				roleInfo.Mnt_by = strings.TrimSpace(part[1])
			}

		case "mnt-routes":
			baseInfo.Mnt_routes = strings.TrimSpace(part[1])

		case "mnt-domains":
			baseInfo.Mnt_domains = strings.TrimSpace(part[1])

		case "created":
			if baseInfo.Inetnum != "" {
				baseInfo.Created = strings.TrimSpace(part[1])
			}
			if orgInfo.Organisation != "" {
				orgInfo.Created = strings.TrimSpace(part[1])
			}
			if roleInfo.Role != "" {
				roleInfo.Created = strings.TrimSpace(part[1])
			}

		case "last-modified":
			if baseInfo.Inetnum != "" {
				baseInfo.Last_modified = strings.TrimSpace(part[1])
			}
			if orgInfo.Organisation != "" {
				orgInfo.Last_modified = strings.TrimSpace(part[1])
			}
			if roleInfo.Role != "" {
				roleInfo.Last_modified = strings.TrimSpace(part[1])
			}

		case "source":
			if baseInfo.Inetnum != "" {
				baseInfo.Source = strings.TrimSpace(part[1])

				Base_Info = append(Base_Info, baseInfo)
				baseInfo.Reset()
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
			if roleInfo.Role != "" {
				roleInfo.Source = strings.TrimSpace(part[1])
				Role_Info = append(Role_Info, roleInfo)
				roleInfo.Reset()
			}

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
			if roleInfo.Role != "" {
				if roleInfo.Address != "" {
					roleInfo.Address += " " + strings.TrimSpace(part[1])
				} else {
					roleInfo.Address = strings.TrimSpace(part[1])
				}
			}

		case "phone":
			if orgInfo.Organisation != "" {
				if orgInfo.Phone != "" {
					orgInfo.Phone += " ," + strings.TrimSpace(part[1])
				} else {
					orgInfo.Phone = strings.TrimSpace(part[1])
				}
			}
			if personInfo.Person != "" {
				if personInfo.Phone != "" {
					personInfo.Phone += " " + strings.TrimSpace(part[1])
				} else {
					personInfo.Phone = strings.TrimSpace(part[1])
				}
			}

		case "fax-no":
			orgInfo.Fax_no = strings.TrimSpace(part[1])

		case "abuse-c":
			orgInfo.Abuse_c = strings.TrimSpace(part[1])

		case "mnt-ref":
			if orgInfo.Organisation != "" {
				if orgInfo.Mnt_ref != "" {
					orgInfo.Mnt_ref += " " + strings.TrimSpace(part[1])
				} else {
					orgInfo.Mnt_ref = strings.TrimSpace(part[1])
				}
			}

			//parser nic person
		case "person":
			personInfo.Person = strings.TrimSpace(part[1])

		case "nic-hdl":
			if personInfo.Person != "" {
				personInfo.Nic_hdl = strings.TrimSpace(part[1])
			}
			if roleInfo.Role != "" {
				roleInfo.Nic_hdl = strings.TrimSpace(part[1])
			}

			//parser riperole
		case "role":
			roleInfo.Role = strings.TrimSpace(part[1])
		}
	}
	return Base_Info, Org_Info, Person_Info
}
