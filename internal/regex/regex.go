package regex

import "regexp"

type Matcher interface {
	GodaddyKey(s string) bool
	GodaddySecret(s string) bool
	DuckDNSToken(s string) bool
	NamecheapPassword(s string) bool
	DreamhostKey(s string) bool
	CloudflareKey(s string) bool
	CloudflareUserServiceKey(s string) bool
}

type matcher struct {
	goDaddyKey, goDaddySecret, duckDNSToken, namecheapPassword, dreamhostKey, cloudflareKey, cloudflareUserServiceKey *regexp.Regexp
}

//nolint:gocritic
func NewMatcher() (m Matcher, err error) {
	matcher := &matcher{}
	matcher.goDaddyKey, err = regexp.Compile(`^[A-Za-z0-9]{10,14}\_[A-Za-z0-9]{22}$`)
	if err != nil {
		return nil, err
	}
	matcher.goDaddySecret, err = regexp.Compile(`^[A-Za-z0-9]{22}$`)
	if err != nil {
		return nil, err
	}
	matcher.duckDNSToken, err = regexp.Compile(`^[a-f0-9]{8}\-[a-f0-9]{4}\-[a-f0-9]{4}\-[a-f0-9]{4}\-[a-f0-9]{12}$`)
	if err != nil {
		return nil, err
	}
	matcher.namecheapPassword, err = regexp.Compile(`^[a-f0-9]{32}$`)
	if err != nil {
		return nil, err
	}
	matcher.dreamhostKey, err = regexp.Compile(`^[a-zA-Z0-9]{16}$`)
	if err != nil {
		return nil, err
	}
	matcher.cloudflareKey, err = regexp.Compile(`^[a-zA-Z0-9]+$`)
	if err != nil {
		return nil, err
	}
	matcher.cloudflareUserServiceKey, err = regexp.Compile(`^v1\.0.+$`)
	if err != nil {
		return nil, err
	}
	return matcher, nil
}

func (m *matcher) GodaddyKey(s string) bool        { return m.goDaddyKey.MatchString(s) }
func (m *matcher) GodaddySecret(s string) bool     { return m.goDaddySecret.MatchString(s) }
func (m *matcher) DuckDNSToken(s string) bool      { return m.duckDNSToken.MatchString(s) }
func (m *matcher) NamecheapPassword(s string) bool { return m.namecheapPassword.MatchString(s) }
func (m *matcher) DreamhostKey(s string) bool      { return m.dreamhostKey.MatchString(s) }
func (m *matcher) CloudflareKey(s string) bool     { return m.cloudflareKey.MatchString(s) }
func (m *matcher) CloudflareUserServiceKey(s string) bool {
	return m.cloudflareUserServiceKey.MatchString(s)
}