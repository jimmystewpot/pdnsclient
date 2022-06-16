package pdns

import (
	"net/http"
	"testing"
	"time"
)

func TestCheckHost(t *testing.T) {
	type args struct {
		hostname string
		port     string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid hostname",
			args: args{
				hostname: "foo.org",
				port:     "8888",
			},
			want:    "http://foo.org:8888/",
			wantErr: false,
		},
		{
			name: "hostname that will not url.parse",
			args: args{
				hostname: ":\\^",
				port:     "8888",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkHost(tt.args.hostname, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultTransport(t *testing.T) {
	type want struct {
		IdleTimeout time.Duration
		IdleConns   int
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "Check Return Type",
			want: want{
				IdleTimeout: defaultTransportTimeout,
				IdleConns:   defaultTransportIdleConns,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DefaultTransport()
			if got.IdleConnTimeout != defaultTransportTimeout {
				t.Errorf("DefaultTransport() idle connections timeout = %v, want %v", got, tt.want)
			}
			if got.MaxIdleConns != defaultTransportIdleConns {
				t.Errorf("DefaultTransport() max idle connections = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClientWithDefaults(t *testing.T) {
	type args struct {
		hostname string
		port     string
		apikey   string
	}
	tests := []struct {
		name    string
		args    args
		want    Client
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				hostname: "foo.org",
				port:     "8888",
				apikey:   "changeme",
			},
			want: Client{
				Host:   "http://foo.org:8888/",
				APIKey: "changeme",
			},
			wantErr: false,
		},
		{
			name: "expected error with host parse",
			args: args{
				hostname: ":\\^",
				port:     "8888",
				apikey:   "changeme",
			},
			want:    Client{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClientWithDefaults(tt.args.hostname, tt.args.port, tt.args.apikey)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("NewClientWithDefaults() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				return
			}
			if got.Host != tt.want.Host {
				t.Errorf("NewClientWithDefaults() Host = %v, want %v", got.Host, tt.want.Host)
			}
			if got.APIKey != tt.want.APIKey {
				t.Errorf("NewClientWithDefaults() APIKey = %v, want %v", got.APIKey, tt.want.APIKey)
			}
			if got.UserAgent != useragent {
				t.Errorf("NewClientWithDefaults() UserAgent = %v, want %v", got.APIKey, tt.want.APIKey)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		hostname  string
		port      string
		apikey    string
		useragent string
		transport *http.Transport
	}
	tests := []struct {
		name    string
		args    args
		want    Client
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				hostname:  "foo.org",
				port:      "8888",
				apikey:    "changeme",
				useragent: useragent,
				transport: DefaultTransport(),
			},
			want: Client{
				Host:      "http://foo.org:8888/",
				APIKey:    "changeme",
				UserAgent: useragent,
			},
			wantErr: false,
		},
		{
			name: "happy path different useragent",
			args: args{
				hostname:  "foo.org",
				port:      "8888",
				apikey:    "changeme",
				useragent: "my foo agent",
				transport: DefaultTransport(),
			},
			want: Client{
				Host:      "http://foo.org:8888/",
				APIKey:    "changeme",
				UserAgent: "my foo agent",
			},
			wantErr: false,
		},
		{
			name: "expected error with host parse",
			args: args{
				hostname:  ":\\^",
				port:      "8888",
				apikey:    "changeme",
				useragent: useragent,
				transport: DefaultTransport(),
			},
			want:    Client{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.hostname, tt.args.port, tt.args.apikey, tt.args.useragent, tt.args.transport)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				return
			}
			if got.Host != tt.want.Host {
				t.Errorf("NewClient() Host = %v, want %v", got.Host, tt.want.Host)
			}
			if got.APIKey != tt.want.APIKey {
				t.Errorf("NewClient() APIKey = %v, want %v", got.APIKey, tt.want.APIKey)
			}
			if got.UserAgent != tt.want.UserAgent {
				t.Errorf("NewClient() UserAgent = %v, want %v", got.APIKey, tt.want.APIKey)
			}
		})
	}
}
