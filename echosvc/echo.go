package echosvc

import (
	"context"

	"github.com/kumparan/kumgo-stack/buff"
)

// Echo :nodoc:
func (m *svc) Echo(c context.Context, s *buff.EchoRequest) (*buff.EchoResponse, error) {
	return &buff.EchoResponse{Message: s.Message}, nil
}
