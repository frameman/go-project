package variable

type RipeBaseInfo struct {
	Inetnum       string
	Netname       string
	Country       string
	Admin_c       string
	Tech_c        string
	Status        string
	Org           string
	Mnt_by        string
	Mnt_routes    string
	Mnt_domains   string
	Created       string
	Last_modified string
	Source        string
	Geofeed       string
}

func (m *RipeBaseInfo) Reset() {
	*m = RipeBaseInfo{}
}

type RipeOrgInfo struct {
	Organisation  string
	Org_name      string
	Org_type      string
	Address       string
	Phone         string
	Fax_no        string
	Admin_c       string
	Country       string
	Abuse_c       string
	Mnt_ref       string
	Mnt_by        string
	Created       string
	Last_modified string
	Source        string
}

func (m *RipeOrgInfo) Reset() {
	*m = RipeOrgInfo{}
}

type RipePContact struct {
	Person        string
	Address       string
	Phone         string
	Nic_hdl       string
	Mnt_by        string
	Created       string
	Last_modified string
	Source        string
}

func (m *RipePContact) Reset() {
	*m = RipePContact{}
}

type RipeRole struct {
	Role          string
	Address       string
	Nic_hdl       string
	Mnt_by        string
	Created       string
	Last_modified string
	Source        string
}

func (m *RipeRole) Reset() {
	*m = RipeRole{}
}

/*
% This is the RIPE Database query service.
% The objects are in RPSL format.
%
% The RIPE Database is subject to Terms and Conditions.
% See https://apps.db.ripe.net/docs/HTML-Terms-And-Conditions

% Note: this output has been filtered.
% To receive output for a database update, use the "-B" flag.

% Information related to '84.205.64.0 - 84.205.95.255'

% Abuse contact for '84.205.64.0 - 84.205.95.255' is 'abuse(&>ripe.net'

inetnum: 84.205.64.0 - 84.205.95.255
netname: RIPE-NCC-RIS-BEACON
org: ORG-RIEN1-RIPE
country: EU
remarks: RIPE NCC RIS anchors and beacons for BGP studies
admin-c: RISM-RIPE
tech-c: RISM-RIPE
status: ASSIGNED PI
mnt-by: RIPE-NCC-END-MNT
mnt-by: RIPE-NCC-MNT
mnt-routes: RIPE-NCC-RIS-MNT
mnt-domains: RIPE-NCC-RIS-MNT
created: 2004-11-05T14:49:06Z
last-modified: 2024-03-18T18:50:08Z
source: RIPE

organisation: ORG-RIEN1-RIPE
org-name: Reseaux IP Europeens Network Coordination Centre (RIPE NCC)
country: NL
org-type: LIR
descr: RIPE NCC Operations
address: P.O. Box 10096
address: 1001 EB
address: Amsterdam
address: NETHERLANDS
phone: +31205354444
fax-no: +31205354445
admin-c: MENN1-RIPE
admin-c: MDIR-RIPE
abuse-c: ops4-ripe
mnt-ref: RIPE-NCC-HM-MNT
mnt-ref: RIPE-NCC-MNT
mnt-by: RIPE-NCC-HM-MNT
mnt-by: RIPE-NCC-MNT
created: 2012-03-09T13:21:52Z
last-modified: 2023-12-05T13:40:03Z
source: RIPE # Filtered

role: RIPE NCC RIS Operations
address: Stationsplein 11
address: 1012 AB Amsterdam
address: The Netherlands
phone: +31 20 535 4444
fax-no: +31 20 535 4445
admin-c: GII-RIPE
tech-c: GII-RIPE
nic-hdl: RISM-RIPE
mnt-by: RIPE-NCC-RIS-MNT
created: 2002-09-27T09:14:06Z
last-modified: 2017-08-15T11:52:31Z
source: RIPE # Filtered

% Information related to '84.205.64.0/24AS12654'

route: 84.205.64.0/24
descr: RIPE-NCC-RIS BGP Anycast Beacon Prefix (&> RRC00,RRC25
origin: AS12654
mnt-by: RIPE-NCC-RIS-MNT
created: 2005-01-31T13:08:09Z
last-modified: 2024-02-02T13:56:47Z
source: RIPE

% This query was served by the RIPE Database Query Service version 1.109.1 (ABERDEEN)
*/
