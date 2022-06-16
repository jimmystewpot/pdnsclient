package pdns

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultTransportTimeout   time.Duration = 15 * time.Second
	defaultTransportIdleConns int           = 10
	useragent                 string        = "go-powerdns-client"
	// apiStatsPath              string        = "api/v1/servers/localhost/statistics"
)

type Client struct {
	Host      string
	APIKey    string
	UserAgent string
	conn      *http.Client
}

// NewClient will return a PowerDNS client with the flexibility to set your own http.Transport. This allows you to configure
// the behaviour of the client. To set a client with a custom transport you could do this
// transport := &http.Transport{
//	      IdleConnTimeout: 200 * time.Second,
//	}
// client, err := NewClient("foobar.org", "8080", "changeme", "my-user-agent", transport)
//
func NewClient(hostname string, port string, apikey string, useragent string, transport *http.Transport) (*Client, error) {
	host, err := checkHost(hostname, port)
	if err != nil {
		return &Client{}, err
	}
	return &Client{
		Host:      host,
		APIKey:    apikey,
		UserAgent: useragent,
		conn: &http.Client{
			Transport: transport,
		},
	}, nil
}

// NewClientWithDefaults will return a PowerDNS client with the default transport, useragent and idle timeouts set.
func NewClientWithDefaults(hostname string, port string, apikey string) (*Client, error) {
	host, err := checkHost(hostname, port)
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		Host:      host,
		APIKey:    apikey,
		UserAgent: useragent,
		conn: &http.Client{
			Transport: DefaultTransport(),
		},
	}, nil
}

// DefaultTransport sets the default values for the HTTP transport. This could be used if you wished to set the
// user-agent and not change the default transport values.
// client, err := NewClient("foobar.org", "8080", "changeme", "my-user-agent", DefaultTransport())
func DefaultTransport() *http.Transport {
	return &http.Transport{
		MaxIdleConns:       defaultTransportIdleConns,
		IdleConnTimeout:    defaultTransportTimeout,
		DisableCompression: true,
	}
}

// checkHost will join the host and ensure that it matches a valid URL.
func checkHost(hostname, port string) (string, error) {
	host := fmt.Sprintf("http://%s/", net.JoinHostPort(hostname, port))
	_, err := url.Parse(host)
	if err != nil {
		return "", err
	}
	return host, nil
}
