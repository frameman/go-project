package whois

import (
	"fmt"
	"net"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

// Constants for WHOIS servers of different RIRs
const (
	arinWhoisServer    = "whois.arin.net"
	ripeWhoisServer    = "whois.ripe.net"
	apnicWhoisServer   = "whois.apnic.net"
	lacnicWhoisServer  = "whois.lacnic.net"
	afrinicWhoisServer = "whois.afrinic.net"
	defaultWhoisPort   = "43"
	defaultTimeout     = 30 * time.Second
)

var DefaultClient = NewClient()

type Client struct {
	dialer  proxy.Dialer
	timeout time.Duration
}

func NewClient() *Client {
	return &Client{
		dialer:  &net.Dialer{Timeout: defaultTimeout},
		timeout: defaultTimeout,
	}
}

func (c *Client) SetDialer(dialer proxy.Dialer) *Client {
	c.dialer = dialer
	return c
}

func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.timeout = timeout
	return c
}

func (c *Client) Whois(query string, servers ...string) (result string, err error) {
	server := servers[0]
	result, err = c.rawQuery(query, server, defaultWhoisPort)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (c *Client) rawQuery(query, server, port string) (string, error) {
	conn, err := c.dialer.Dial("tcp", net.JoinHostPort(server, port))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(query + "\r\n"))
	if err != nil {
		return "", err
	}

	var response []byte
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			break
		}
		response = append(response, buffer[:n]...)
	}

	return string(response), nil
}

// Simple function to determine the RIR WHOIS server based on the IP prefix
func getRIRWhoisServer(ip string) string {
	if strings.HasPrefix(ip, "192.") || strings.HasPrefix(ip, "193.") {
		return ripeWhoisServer // RIPE NCC
	} else if strings.HasPrefix(ip, "202.") {
		return apnicWhoisServer // APNIC
	} else if strings.HasPrefix(ip, "200.") || strings.HasPrefix(ip, "201.") {
		return lacnicWhoisServer // LACNIC
	} else if strings.HasPrefix(ip, "41.") || strings.HasPrefix(ip, "105.") {
		return afrinicWhoisServer // AFRINIC
	}
	return arinWhoisServer // Default to ARIN
}

func InputIP(ip string) (result string) {
	client := NewClient()
	server := getRIRWhoisServer(ip)
	result, err := client.Whois(ip, server)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	return result
}
