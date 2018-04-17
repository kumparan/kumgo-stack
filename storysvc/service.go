package storysvc

type svc struct{}

// NewServer :nodoc:
func NewServer() *svc {
	return new(svc)
}
