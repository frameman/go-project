package variable

type AfrinicBaseInfo struct {
	Inetnum     string
	Netname     string
	Descr       string
	Country     string
	Org         string
	Admin_c     string
	Tech_c      string
	Status      string
	Mnt_by      string
	Mnt_lower   string
	Mnt_domains string
	Mnt_routes  string
	Mnt_irt     string
	Source      string
	Parent      string
}

func (m *AfrinicBaseInfo) Reset() {
	*m = AfrinicBaseInfo{}
}

type AfrinicOrgInfo struct {
	Organisation string
	Org_name     string
	Org_type     string
	Country      string
	Address      string
	Phone        string
	Admin_c      string
	Tech_c       string
	Mnt_ref      string
	Mnt_by       string
	Remarks      string
	Source       string
}

func (m *AfrinicOrgInfo) Reset() {
	*m = AfrinicOrgInfo{}
}

type AfrinicNIC struct {
	Person  string
	Nic_hdl string
	Address string
	Phone   string
	Mnt_by  string
	Source  string
}

func (m *AfrinicNIC) Reset() {
	*m = AfrinicNIC{}
}

/* % This is the AfriNIC Whois server.
% The AFRINIC whois database is subject to the following terms of Use. See https://afrinic.net/whois/terms

% Note: this output has been filtered.
% To receive output for a database update, use the "-B" flag.

% Information related to '196.216.2.0 - 196.216.3.255'

% Abuse contact for '196.216.2.0 - 196.216.3.255' is 'abuse{At]afrinic.net'

inetnum: 196.216.2.0 - 196.216.3.255
netname: AFRINIC
descr: AfriNIC - Internal Use
country: ZA
org: ORG-AFNC1-AFRINIC
admin-c: CA15-AFRINIC
tech-c: IT7-AFRINIC
status: ASSIGNED PI
mnt-by: AFRINIC-HM-MNT
mnt-lower: AFRINIC-IT-MNT
mnt-domains: AFRINIC-IT-MNT
mnt-routes: AFRINIC-IT-MNT
mnt-irt: IRT-AFRINIC-IT
source: AFRINIC # Filtered
parent: 196.0.0.0 - 196.255.255.255

organisation: ORG-AFNC1-AFRINIC
org-name: African Network Information Center - ( AfriNIC Ltd. )
org-type: RIR
country: MU
address: 11th Floor, Standard Chartered Tower
address: 19, Cybercity
address: Ebène
phone: tel:+230-466-6758
phone: tel:+230-403-5100
admin-c: CA15-AFRINIC
tech-c: IT7-AFRINIC
mnt-ref: AFRINIC-HM-MNT
mnt-ref: AFRINIC-DB-MNT
mnt-ref: AFRINIC-IT-MNT
mnt-by: AFRINIC-HM-MNT
remarks: =======================================
remarks: For more information on AFRINIC assigned blocks,
remarks: querry whois.afrinic.net port 43, or the web based
remarks: query at http://whois.afrinic.net or www.afrinic.net
remarks: website: www.afrinic.net
remarks: Other Contacts:
remarks: ===============
remarks: hostmaster{At]afrinic.net - for IP resources
remarks: new-member{At]afrinic.net - for new members and other
remarks: inquiries.
source: AFRINIC # Filtered

person: CTO AFRINIC
address: 11th Floor, Standard Chartered Tower
address: Cybercity, Ebène
address: Mauritius
phone: tel:+230-403-5100
nic-hdl: CA15-AFRINIC
mnt-by: CTO-MNT
source: AFRINIC # Filtered

person: Infrastructure Team
nic-hdl: IT7-AFRINIC
address: AFRINIC Ltd
address: 11th Floor, Standard Chartered Tower
address: 19, Cybercity
address: Ebène
address: Mauritius
phone: tel:+230-403-5100
phone: tel:+230-5943-8680
mnt-by: AFRINIC-IT-MNT
source: AFRINIC # Filtered

% Information related to '196.216.2.0/23AS33764'

route: 196.216.2.0/23
descr: AFRINIC-ZA-JNB-AS
origin: AS33764
mnt-by: AFRINIC-IT-MNT
source: AFRINIC # Filtered*/
