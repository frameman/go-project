package variable

type BaseInfo struct {
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

func (m *BaseInfo) Reset() {
	*m = BaseInfo{}
}

type OrgInfo struct {
	OrgName    string
	OrgId      string
	Address    string
	City       string
	StateProv  string
	PostalCode string
	Country    string
	RegDate    string
	Updated    string
	//Comment        string //same result with the baseinfo.comment
	Ref            string
	ReferralServer string
}

func (m *OrgInfo) Reset() {
	*m = OrgInfo{}
}

// the above is option
type ContactPOrg struct {
	OrgTechHandle string
	OrgTechName   string
	OrgTechPhone  string
	OrgTechEmail  string
	OrgTechRef    string
}

func (m *ContactPOrg) Reset() {
	*m = ContactPOrg{}
}

type ContactPR struct {
	RTechHandle string
	RTechName   string
	RTechPhone  string
	RTechEmail  string
	RTechRef    string
}

func (m *ContactPR) Reset() {
	*m = ContactPR{}
}

type OrgNOC struct {
	OrgNOCHandle string
	OrgNOCName   string
	OrgNOCPhone  string
	OrgNOCEmail  string
	OrgNOCRef    string
}

func (m *OrgNOC) Reset() {
	*m = OrgNOC{}
}

type ContactPAbuse struct {
	OrgAbuseHandle string
	OrgAbuseName   string
	OrgAbusePhone  string
	OrgAbuseEmail  string
	OrgAbuseRef    string
}

func (m *ContactPAbuse) Reset() {
	*m = ContactPAbuse{}
}
