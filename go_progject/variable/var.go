package var_

type base_info struct {
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
	Commen       string
	Ref          string
}

var whois_base_info = base_info{}

type org_info struct {
	OrgName        string
	OrgId          string
	Address        string
	City           string
	StateProv      string
	PostalCode     int16
	Country        string
	RegDate        string
	Updated        string
	Comment        string
	Ref            string
	ReferralServer string
}

var whois_org_info = org_info{}

// the above is option
type contact_p_org struct {
	OrgTechHandle string
	OrgTechName   string
	OrgTechPhone  string
	OrgTechEmail  string
	OrgTechRef    string
}

var whois_contact_p_org = contact_p_org{}

type contact_p_r struct {
	RTechHandle string
	RTechName   string
	RTechPhone  string
	RTechEmail  string
	RTechRef    string
}

var whois_contact_p_r = contact_p_r{}

type contact_p_abuse struct {
	OrgAbuseHandle string
	OrgAbuseName   string
	OrgAbusePhone  string
	OrgAbuseEmail  string
	OrgAbuseRef    string
}

var whois_contact_p_abuse = contact_p_abuse{}
