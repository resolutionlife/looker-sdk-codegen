package rtl

type HeaderOptions struct {
	Headers map[string]string
}

// http2 has lower case headers, look in go sdk
func (h *HeaderOptions) setAccept() {

	if _, ok := h.Headers["Accept"]; !ok {
		h.Headers["Accept"] = "application/json"
	}
}

// use type http.Header
func NewHeaderOptions(header *map[string]string) *HeaderOptions {
	h := new(HeaderOptions)
	if header != nil {
		h.Headers = *header
	} else {
		h.Headers = make(map[string]string, 0)
	}

	h.setAccept()

	return h
}
