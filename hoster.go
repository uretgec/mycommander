package mycommander

// hoster
type Hoster struct {
	hosts map[string]struct{}
}

// create new hoster
func NewHoster(hosts ...string) *Hoster {
	t := &Hoster{
		hosts: map[string]struct{}{},
	}

	if len(hosts) > 0 {
		for _, host := range hosts {
			t.hosts[host] = struct{}{}
		}
	}

	return t
}

// valid host
func (h *Hoster) valid(host string) bool {
	_, ok := h.hosts[host]
	return ok
}
