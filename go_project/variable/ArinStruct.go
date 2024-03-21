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
