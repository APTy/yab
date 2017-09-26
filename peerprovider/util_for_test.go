package peerprovider

import "net/url"

func mustParseURL(s string) *url.URL {
	u, _ := url.Parse(s)
	return u
}

func stubRegistry() func() {
	lock.Lock()
	oldRegistry := registry
	registry = make(map[string]PeerProvider)
	lock.Unlock()
	return func() {
		lock.Lock()
		registry = oldRegistry
		lock.Unlock()
	}
}
