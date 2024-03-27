package variable

type ArinBaseInfo struct {
	NetRange     string
	OriginAS     string
	NetName      string
	CIDR         string
	NetHandle    string
	Parent       string
	NetType      string
	Organization string
	RegDate      string
	Updated      string
	Comment      string
	Ref          string
}

func (m *ArinBaseInfo) Reset() {
	*m = ArinBaseInfo{}
}

type ArinOrgInfo struct {
	OrgName        string
	OrgId          string
	Address        string
	City           string
	StateProv      string
	PostalCode     string
	Country        string
	RegDate        string
	Updated        string
	Ref            string
	ReferralServer string
	Comment        string
}

func (m *ArinOrgInfo) Reset() {
	*m = ArinOrgInfo{}
}

// the above is option
type ArinContactPOrg struct {
	OrgTechHandle string
	OrgTechName   string
	OrgTechPhone  string
	OrgTechEmail  string
	OrgTechRef    string
}

func (m *ArinContactPOrg) Reset() {
	*m = ArinContactPOrg{}
}

type ArinContactPR struct {
	RTechHandle string
	RTechName   string
	RTechPhone  string
	RTechEmail  string
	RTechRef    string
}

func (m *ArinContactPR) Reset() {
	*m = ArinContactPR{}
}

type ArinOrgNOC struct {
	ArinOrgNOCHandle string
	ArinOrgNOCName   string
	ArinOrgNOCPhone  string
	ArinOrgNOCEmail  string
	ArinOrgNOCRef    string
}

func (m *ArinOrgNOC) Reset() {
	*m = ArinOrgNOC{}
}

type ArinContactPAbuse struct {
	OrgAbuseHandle string
	OrgAbuseName   string
	OrgAbusePhone  string
	OrgAbuseEmail  string
	OrgAbuseRef    string
}

func (m *ArinContactPAbuse) Reset() {
	*m = ArinContactPAbuse{}
}

/*#
# ARIN WHOIS data and services are subject to the Terms of Use
# available at: https://www.arin.net/resources/registry/whois/tou/
#
# If you see inaccuracies in the results, please report at
# https://www.arin.net/resources/registry/whois/inaccuracy_reporting/
#
# Copyright 1997-2024, American Registry for Internet Numbers, Ltd.
#


NetRange: 155.94.128.0 - 155.94.255.255
CIDR: 155.94.128.0/17
NetName: QUADRANET
NetHandle: NET-155-94-128-0-1
Parent: NET155 (NET-155-0-0-0-0)
NetType: Direct Allocation
OriginAS: AS8100
Organization: QuadraNet Enterprises LLC (QEL-5)
RegDate: 2014-06-11
Updated: 2018-08-30
Ref: https://rdap.arin.net/registry/ip/155.94.128.0


OrgName: QuadraNet Enterprises LLC
OrgId: QEL-5
Address: 530 W. 6th ST
Address: STE #901
City: Los Angeles
StateProv: CA
PostalCode: 90014
Country: US
RegDate: 2018-06-07
Updated: 2023-02-14
Ref: https://rdap.arin.net/registry/entity/QEL-5

ReferralServer: rwhois://rwhois.quadranet.com:4321

OrgAbuseHandle: QUADR4-ARIN
OrgAbuseName: QuadraNet Abuse
OrgAbusePhone: +1-213-614-8371
OrgAbuseEmail: abuse[At]quadranet.com
OrgAbuseRef: https://rdap.arin.net/registry/entity/QUADR4-ARIN

OrgTechHandle: QNO6-ARIN
OrgTechName: QuadraNet Network Operations
OrgTechPhone: +1-213-614-9371
OrgTechEmail: support[At]quadranet.com
OrgTechRef: https://rdap.arin.net/registry/entity/QNO6-ARIN


#
# ARIN WHOIS data and services are subject to the Terms of Use
# available at: https://www.arin.net/resources/registry/whois/tou/
#
# If you see inaccuracies in the results, please report at
# https://www.arin.net/resources/registry/whois/inaccuracy_reporting/
#
# Copyright 1997-2024, American Registry for Internet Numbers, Ltd.
# */
