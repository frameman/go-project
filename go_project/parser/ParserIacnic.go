package parser

import (
	"go_project/variable"
	"strings"
)

func ParserIacnicData(text string) ([]variable.IacnicBaseInfo, []variable.IacnicNIC) {
	baseInfo := variable.IacnicBaseInfo{}
	Base_Info := []variable.IacnicBaseInfo{}

	personInfo := variable.IacnicNIC{}
	Person_Info := []variable.IacnicNIC{}

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		part := strings.SplitN(line, ":", 2)
		switch part[0] {

		//parser baseInfo
		case "inetnum":
			baseInfo.Inetnum = strings.TrimSpace(part[1])

		case "status":
			baseInfo.Status = strings.TrimSpace(part[1])

		case "aut-num":
			baseInfo.Aut_num = strings.TrimSpace(part[1])

		case "owner":
			baseInfo.Owner = strings.TrimSpace(part[1])

		case "ownerid":
			baseInfo.Ownerid = strings.TrimSpace(part[1])

		case "responsible":
			baseInfo.Responsible = strings.TrimSpace(part[1])

		case "address":
			if baseInfo.Inetnum != "" {
				if baseInfo.Address != "" {
					baseInfo.Address += " " + strings.TrimSpace(part[1])
				} else {
					baseInfo.Address = strings.TrimSpace(part[1])
				}
			}
			if personInfo.Nic_hdl != "" {
				if personInfo.Address != "" {
					personInfo.Address += " " + strings.TrimSpace(part[1])
				} else {
					personInfo.Address = strings.TrimSpace(part[1])
				}
			}

		case "country":
			if baseInfo.Inetnum != "" {
				baseInfo.Country = strings.TrimSpace(part[1])
			}
			if personInfo.Nic_hdl != "" {
				personInfo.Country = strings.TrimSpace(part[1])
			}

		case "phone":
			if baseInfo.Inetnum != "" {
				if baseInfo.Phone != "" {
					baseInfo.Phone += " " + strings.TrimSpace(part[1])
				} else {
					baseInfo.Phone = strings.TrimSpace(part[1])
				}
			}
			if personInfo.Nic_hdl != "" {
				personInfo.Phone = strings.TrimSpace(part[1])
			}

		case "owner-c":
			baseInfo.Owner_c = strings.TrimSpace(part[1])

		case "tech-c":
			if baseInfo.Inetnum != "" {
				baseInfo.Tech_c = strings.TrimSpace(part[1])
			}
			if personInfo.Person != "" {
				if personInfo.Phone != "" {
					personInfo.Phone += " " + strings.TrimSpace(part[1])
				} else {
					personInfo.Phone = strings.TrimSpace(part[1])
				}
			}

		case "abuse-c":
			baseInfo.Abuse_c = strings.TrimSpace(part[1])

		case "intrev":
			baseInfo.Inetrev = strings.TrimSpace(part[1])

		case "nserver":
			if baseInfo.Nserver != "" {
				baseInfo.Nserver += " " + strings.TrimSpace(part[1])
			} else {
				baseInfo.Nserver = strings.TrimSpace(part[1])
			}

		case "nslastaa":
			if baseInfo.Nslastaa != "" {
				baseInfo.Nslastaa += " " + strings.TrimSpace(part[1])
			} else {
				baseInfo.Nslastaa = strings.TrimSpace(part[1])
			}
		case "nsstat":
			if baseInfo.Nsstat != "" {
				baseInfo.Nsstat += " " + strings.TrimSpace(part[1])
			} else {
				baseInfo.Nsstat = strings.TrimSpace(part[1])
			}

		case "created":
			if baseInfo.Inetnum != "" {
				baseInfo.Created = strings.TrimSpace(part[1])
			}
			if personInfo.Nic_hdl != "" {
				personInfo.Created = strings.TrimSpace(part[1])
			}

		case "changed":

			if baseInfo.Inetnum != "" {
				baseInfo.Changed = strings.TrimSpace(part[1])
				Base_Info = append(Base_Info, baseInfo)
				baseInfo.Reset()
			}
			if personInfo.Nic_hdl != "" {
				personInfo.Changed = strings.TrimSpace(part[1])
				Person_Info = append(Person_Info, personInfo)
				personInfo.Reset()
			}

			//parser nic person
		case "nic-hdl":
			personInfo.Nic_hdl = strings.TrimSpace(part[1])

		case "person":
			personInfo.Person = strings.TrimSpace(part[1])

		case "e-mail":
			personInfo.Email = strings.TrimSpace(part[1])
		}
	}

	return Base_Info, Person_Info
}
