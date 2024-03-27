package parser

import (
	"go_project/variable"
	"strings"
)

func ParserApnicData(text string) ([]variable.ApnicBaseInfo, []variable.ApnicIRT, []variable.ApnicRole, []variable.ApnicPerson) {

	Base_Info := []variable.ApnicBaseInfo{}
	baseInfo := variable.ApnicBaseInfo{}

	Irt_Info := []variable.ApnicIRT{}
	irtInfo := variable.ApnicIRT{}

	Role_info := []variable.ApnicRole{}
	roleInfo := variable.ApnicRole{}

	NICPerson_Info := []variable.ApnicPerson{}
	nicPersonInfo := variable.ApnicPerson{}

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		part := strings.SplitN(line, ":", 2)
		switch part[0] {

		// paraser the base info first
		case "inetnum":
			baseInfo.Inetnum = strings.TrimSpace(part[1])

		case "netname":
			baseInfo.Netname = strings.TrimSpace(part[1])

		case "descr":
			if baseInfo.Descr != "" {
				baseInfo.Descr += " / " + strings.TrimSpace(part[1])
			} else {
				baseInfo.Descr = strings.TrimSpace(part[1])
			}

		case "country":
			if baseInfo.Inetnum != "" {
				baseInfo.Country = strings.TrimSpace(part[1])
			}
			if roleInfo.Role != "" {
				roleInfo.Country = strings.TrimSpace(part[1])
			}
			if nicPersonInfo.Person != "" {
				nicPersonInfo.Country = strings.TrimSpace(part[1])
			}

		case "admin-c":
			if baseInfo.Inetnum != "" {
				baseInfo.Admin_c = strings.TrimSpace(part[1])
			}
			if irtInfo.Irt != "" {
				irtInfo.Admin_c = strings.TrimSpace(part[1])
			}
			if roleInfo.Role != "" {
				roleInfo.Admin_c = strings.TrimSpace(part[1])
			}

		case "tech-c":
			if baseInfo.Inetnum != "" {
				baseInfo.Tech_c = strings.TrimSpace(part[1])
			}
			if irtInfo.Irt != "" {
				irtInfo.Tech_c = strings.TrimSpace(part[1])
			}
			if roleInfo.Role != "" {
				roleInfo.Tech_c = strings.TrimSpace(part[1])
			}

		case "abuse-c":
			baseInfo.Abuse_c = strings.TrimSpace(part[1])
		case "status":
			baseInfo.Status = strings.TrimSpace(part[1])

		case "remarks":
			if baseInfo.Inetnum != "" {
				if baseInfo.Remarks != "" {
					baseInfo.Remarks += " " + strings.TrimSpace(part[1])
				} else {
					baseInfo.Remarks = strings.TrimSpace(part[1])
				}
			}
			if irtInfo.Irt != "" {
				if irtInfo.Remarks != "" {
					irtInfo.Remarks += " " + strings.TrimSpace(part[1])
				} else {
					irtInfo.Remarks = strings.TrimSpace(part[1])
				}
			}
			if roleInfo.Role != "" {
				if roleInfo.Remarks != "" {
					roleInfo.Remarks += " " + strings.TrimSpace(part[1])
				} else {
					roleInfo.Remarks = strings.TrimSpace(part[1])
				}
			}

		case "mnt-by":
			if baseInfo.Inetnum != "" {
				baseInfo.Mnt_by = strings.TrimSpace(part[1])
			}
			if irtInfo.Irt != "" {
				irtInfo.Mnt_by = strings.TrimSpace(part[1])
			}
			if roleInfo.Role != "" {
				roleInfo.Mnt_by = strings.TrimSpace(part[1])
			}
			if nicPersonInfo.Person != "" {
				nicPersonInfo.Mnt_by = strings.TrimSpace(part[1])
			}

		case "mnt-lower":
			baseInfo.Mnt_lower = strings.TrimSpace(part[1])

		case "mnt-routes":
			baseInfo.Mnt_routes = strings.TrimSpace(part[1])

		case "mnt-irt":
			baseInfo.Mnt_irt = strings.TrimSpace(part[1])

		case "last-modified":
			if baseInfo.Inetnum != "" {
				baseInfo.Last_modified = strings.TrimSpace(part[1])
			}
			if irtInfo.Irt != "" {
				irtInfo.Last_modified = strings.TrimSpace(part[1])
			}
			if roleInfo.Role != "" {
				roleInfo.Last_modified = strings.TrimSpace(part[1])
			}
			if nicPersonInfo.Person != "" {
				nicPersonInfo.Last_modified = strings.TrimSpace(part[1])
			}

		case "source":
			if baseInfo.Inetnum != "" {
				baseInfo.Source = strings.TrimSpace(part[1])
				Base_Info = append(Base_Info, baseInfo)
				baseInfo.Reset()
			}
			if irtInfo.Irt != "" {
				irtInfo.Source = strings.TrimSpace(part[1])
				Irt_Info = append(Irt_Info, irtInfo)
				irtInfo.Reset()
			}
			if roleInfo.Role != "" {
				roleInfo.Source = strings.TrimSpace(part[1])
				Role_info = append(Role_info, roleInfo)
				roleInfo.Reset()
			}
			if nicPersonInfo.Person != "" {
				nicPersonInfo.Source = strings.TrimSpace(part[1])
				NICPerson_Info = append(NICPerson_Info, nicPersonInfo)
				nicPersonInfo.Reset()
			}

			//parser IrtInfo
		case "irt":
			irtInfo.Irt = strings.TrimSpace(part[1])

		case "address":
			if irtInfo.Irt != "" {
				if irtInfo.Address != "" {
					irtInfo.Address += " " + strings.TrimSpace(part[1])
				} else {
					irtInfo.Address = strings.TrimSpace(part[1])
				}
			} else if roleInfo.Role != "" {
				if roleInfo.Address != "" {
					roleInfo.Address += " " + strings.TrimSpace(part[1])
				} else {
					roleInfo.Address = strings.TrimSpace(part[1])
				}
			} else if nicPersonInfo.Person != "" {
				if nicPersonInfo.Address != "" {
					nicPersonInfo.Address += " " + strings.TrimSpace(part[1])
				} else {
					nicPersonInfo.Address = strings.TrimSpace(part[1])
				}
			}

		case "e-mail":
			if irtInfo.Irt != "" {
				irtInfo.Email = strings.TrimSpace(part[1])
			}
			if roleInfo.Role != "" {
				roleInfo.Email = strings.TrimSpace(part[1])
			}
			if nicPersonInfo.Person != "" {
				nicPersonInfo.Email = strings.TrimSpace(part[1])
			}

		case "abuse-mailbox":
			if irtInfo.Irt != "" {
				irtInfo.Abuse_mailbox = strings.TrimSpace(part[1])
			}
			if roleInfo.Role != "" {
				roleInfo.Abuse_mailbox = strings.TrimSpace(part[1])
			}

		case "auth":
			irtInfo.Auth = strings.TrimSpace(part[1])

			//parser roleInfo
		case "role":
			roleInfo.Role = strings.TrimSpace(part[1])

		case "phone":
			if roleInfo.Role != "" {
				roleInfo.Phone = strings.TrimSpace(part[1])
			} else if nicPersonInfo.Person != "" {
				nicPersonInfo.Phone = strings.TrimSpace(part[1])
			}

		case "nic-hdl":
			if roleInfo.Role != "" {
				roleInfo.Nic_hdl = strings.TrimSpace(part[1])
			}
			if nicPersonInfo.Person != "" {
				nicPersonInfo.Nic_hdl = strings.TrimSpace(part[1])
			}

			//parser nicPersonInfo
		case "person":
			nicPersonInfo.Person = strings.TrimSpace(part[1])
		case "fax-no":
			nicPersonInfo.Fax_no = strings.TrimSpace(part[1])

		}
	}
	return Base_Info, Irt_Info, Role_info, NICPerson_Info
}
