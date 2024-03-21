package variable

type IacnicBaseInfo struct {
	Inetnum     string
	Status      string
	Aut_num     string
	Owner       string
	Ownerid     string
	Responsible string
	Address     string
	Country     string
	Phone       string
	Owner_c     string
	Tech_c      string
	Abuse_c     string
	Inetrev     string
	Nserver     string
	Nsstat      string
	Nslastaa    string
	Created     string
	Changed     string
	Inetnum_up  string
}

func (m *IacnicBaseInfo) Reset() {
	*m = IacnicBaseInfo{}
}

type IacnicNIC struct {
	Nic_hdl string
	Person  string
	Email   string
	Address string
	Country string
	Phone   string
	Created string
	Changed string
}

func (m *IacnicNIC) Reset() {
	*m = IacnicNIC{}
}

/*
% IP Client: 155.94.155.220

% Joint Whois - whois.lacnic.net
% This server accepts single ASN, IPv4 or IPv6 queries

% LACNIC resource: whois.lacnic.net


% Copyright LACNIC lacnic.net
% The data below is provided for information purposes
% and to assist persons in obtaining information about or
% related to AS and IP numbers registrations
% By submitting a whois query, you agree to use this data
% only for lawful purposes.
% 2024-03-17 11:16:26 (-03 -03:00)

inetnum: 200.3.12.0/22
status: assigned
aut-num: AS28001
owner: LACNIC - Latin American and Caribbean IP address
ownerid: UY-LACN-LACNIC
responsible: Ernesto Majó
address: Rambla República de México, 6125, Esquina 6 de Abril
address: 11400 - Montevideo -
country: UY
phone: +598 26042222 [4401]
owner-c: AIL
tech-c: AIL
abuse-c: AIL
inetrev: 200.3.12.0/22
nserver: NS.LACNIC.NET
nsstat: 20240314 AA
nslastaa: 20240314
nserver: NS2.LACNIC.NET
nsstat: 20240314 AA
nslastaa: 20240314
nserver: DNS.ANYCAST.LACNIC.NET
nsstat: 20240314 AA
nslastaa: 20240314
created: 20080902
changed: 20080902

nic-hdl: AIL
person: Carlos M Martínez
e-mail: ipadmin(!)lacnic.net
address: Rambla Rep. Mexico, 6125, Casa de Internet
address: 11400 - Montevideo - Montevideo
country: UY
phone: +598 26042222 [4401]
created: 20080125
changed: 20240208

% whois.lacnic.net accepts only direct match queries.
% Types of queries are: POCs, ownerid, CIDR blocks, IP
% and AS numbers
*/
