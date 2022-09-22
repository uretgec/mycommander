package mycommander

// transporter
type Transporter struct {
	transports map[string]struct{}
}

// create new transporter
func NewTransporter(transports ...string) *Transporter {
	t := &Transporter{
		transports: map[string]struct{}{},
	}

	if len(transports) > 0 {
		for _, transport := range transports {
			t.transports[transport] = struct{}{}
		}
	}

	return t
}

// valid scheme
func (t *Transporter) valid(scheme string) bool {
	_, ok := t.transports[scheme]
	return ok
}
