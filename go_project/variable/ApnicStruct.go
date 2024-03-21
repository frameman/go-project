package variable

type ApnicBaseInfo struct {
	Inetnum       string
	Netname       string
	Descr         string
	Country       string
	Admin_c       string
	Tech_c        string
	Abuse_c       string
	Status        string
	Mnt_by        string
	Mnt_lower     string
	Mnt_routes    string
	Mnt_irt       string
	Last_modified string
	Source        string
	Remarks       string
}

func (m *ApnicBaseInfo) Reset() {
	*m = ApnicBaseInfo{}
}

// IRT = Incident Response Teams
type ApnicIRT struct {
	Irt           string
	Address       string
	Email         string
	Abuse_mailbox string
	Admin_c       string
	Tech_c        string
	Auth          string
	Mnt_by        string
	Last_modified string
	Source        string
	Remarks       string
}

func (m *ApnicIRT) Reset() {
	*m = ApnicIRT{}
}

type ApnicRole struct {
	Role          string
	Address       string
	Country       string
	Phone         string
	Email         string
	Admin_c       string
	Tech_c        string
	Nic_hdl       string
	Remarks       string
	Abuse_mailbox string
	Mnt_by        string
	Last_modified string
	Source        string
}

func (m *ApnicRole) Reset() {
	*m = ApnicRole{}
}

type ApnicPerson struct {
	Person        string
	Address       string
	Country       string
	Phone         string
	Fax_no        string
	Email         string
	Nic_hdl       string
	Remarks       string
	Mnt_by        string
	Last_modified string
	Source        string
}

func (m *ApnicPerson) Reset() {
	*m = ApnicPerson{}
}

/*
% [whois.apnic.net]
% Whois data copyright terms http://www.apnic.net/db/dbcopyright.html

% Information related to '113.120.0.0 - 113.127.255.255'

% Abuse contact for '113.120.0.0 - 113.127.255.255' is 'anti-spam<|)chinatelecom.cn'

inetnum: 113.120.0.0 - 113.127.255.255
netname: CHINANET-SD
descr: CHINANET SHANDONG PROVINCE NETWORK
descr: China Telecom
descr: No.31,jingrong street
descr: Beijing 100032
country: CN
admin-c: CH93-AP
tech-c: XR55-AP
abuse-c: AC1573-AP
status: ALLOCATED PORTABLE
remarks: service provider
remarks: --------------------------------------------------------
remarks: To report network abuse, please contact mnt-irt
remarks: For troubleshooting, please contact tech-c and admin-c
remarks: Report invalid contact via www.apnic.net/invalidcontact
remarks: --------------------------------------------------------
mnt-by: APNIC-HM
mnt-lower: MAINT-CHINANET-SD
mnt-routes: MAINT-CHINANET-SD
mnt-irt: IRT-CHINANET-CN
last-modified: 2021-06-15T08:05:37Z
source: APNIC

irt: IRT-CHINANET-CN
address: No.31 ,jingrong street,beijing
address: 100032
e-mail: anti-spam<|)chinatelecom.cn
abuse-mailbox: anti-spam<|)chinatelecom.cn
admin-c: CH93-AP
tech-c: CH93-AP
auth: # Filtered
remarks: anti-spam<|)chinatelecom.cn was validated on 2023-10-08
mnt-by: MAINT-CHINANET
last-modified: 2023-10-08T08:55:58Z
source: APNIC

role: ABUSE CHINANETCN
address: No.31 ,jingrong street,beijing
address: 100032
country: ZZ
phone: +000000000
e-mail: anti-spam<|)chinatelecom.cn
admin-c: CH93-AP
tech-c: CH93-AP
nic-hdl: AC1573-AP
remarks: Generated from irt object IRT-CHINANET-CN
remarks: anti-spam<|)chinatelecom.cn was validated on 2023-10-08
abuse-mailbox: anti-spam<|)chinatelecom.cn
mnt-by: APNIC-ABUSE
last-modified: 2023-10-08T08:56:49Z
source: APNIC

person: Chinanet Hostmaster
nic-hdl: CH93-AP
e-mail: anti-spam<|)chinatelecom.cn
address: No.31 ,jingrong street,beijing
address: 100032
phone: +86-10-58501724
fax-no: +86-10-58501724
country: CN
mnt-by: MAINT-CHINANET
last-modified: 2022-02-28T06:53:44Z
source: APNIC

person: Xin Ruosheng
nic-hdl: XR55-AP
e-mail: ipreport.sd<|)chinatelecom.cn
address: No.999, road Shunhua, Jinan, Shandong province,China
phone: +86-531-83190000
fax-no: +86-531-83190000
country: CN
mnt-by: MAINT-CHINANET-SD
last-modified: 2019-12-20T07:11:49Z
source: APNIC
*/
