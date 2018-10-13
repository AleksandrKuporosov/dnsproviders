// Package yandex adapts the lego Yandex DNS
// provider for Caddy. Importing this package plugs it in.
package yandex

import (
	"errors"

	"github.com/AleksandrKuporosov/lego/providers/dns/yandex"
	"github.com/mholt/caddy/caddytls"
)

func init() {
	caddytls.RegisterDNSProvider("yandex", NewDNSProvider)
}

// NewDNSProvider returns a new Yandex DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = Email
//         credentials[1] = Application Key
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return yandex.NewDNSProvider()
	case 2:
		return yandex.NewDNSProviderCredentials(credentials[0], credentials[1])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
