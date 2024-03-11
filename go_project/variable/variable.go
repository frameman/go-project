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
	Commen       string
	Ref          string
}

func (m *BaseInfo) Reset() {
	*m = BaseInfo{}
}

type OrgInfo struct {
	OrgName        string
	OrgId          string
	Address        string
	City           string
	StateProv      string
	PostalCode     string
	Country        string
	RegDate        string
	Updated        string
	Comment        string
	Ref            string
	ReferralServer string
}

func (m *OrgInfo) Reset() {
	*m = OrgInfo{}
}

// the above is option
type ContractPOrg struct {
	OrgTechHandle string
	OrgTechName   string
	OrgTechPhone  string
	OrgTechEmail  string
	OrgTechRef    string
}

func (m *ContractPOrg) Reset() {
	*m = ContractPOrg{}
}

type ContractPR struct {
	RTechHandle string
	RTechName   string
	RTechPhone  string
	RTechEmail  string
	RTechRef    string
}

func (m *ContractPR) Reset() {
	*m = ContractPR{}
}

type ContractPAbuse struct {
	OrgAbuseHandle string
	OrgAbuseName   string
	OrgAbusePhone  string
	OrgAbuseEmail  string
	OrgAbuseRef    string
}

func (m *ContractPAbuse) Reset() {
	*m = ContractPAbuse{}
}
