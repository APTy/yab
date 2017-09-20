package transport

func (r *Request) DeepCopy() *Request {
	return &Request{
		Method:           r.Method,
		Timeout:          r.Timeout,
		Headers:          deepCopyMap(r.Headers),
		Baggage:          deepCopyMap(r.Baggage),
		TransportHeaders: deepCopyMap(r.TransportHeaders),
		ShardKey:         r.ShardKey,
		Body:             r.Body,
		TargetService:    r.TargetService,
	}
}

func deepCopyMap(src map[string]string) map[string]string {
	dest := make(map[string]string, len(src))
	for k, v := range src {
		dest[k] = v
	}
	return dest
}

func deepCopyBytes(src []byte) []byte {
	dest := make([]byte, len(src))
	copy(dest, src)
	return dest
}
