package echosvc

import (
	"context"
	"testing"

	"github.com/kumparan/kumgo-stack/buff"
)

func TestEcho(t *testing.T) {
	s := &svc{}
	// set up test cases
	tests := []struct {
		message string
		want    string
	}{
		{
			message: "world",
			want:    "world",
		},
		{
			message: "123",
			want:    "123",
		},
	}

	for _, tt := range tests {
		req := &buff.EchoRequest{Message: tt.message}
		resp, err := s.Echo(context.Background(), req)
		if err != nil {
			t.Errorf("TestEcho got unexpected error")
		}
		if resp.Message != tt.want {
			t.Errorf("TestEcho(%v)=%v, wanted %v", tt.message, resp.Message, tt.want)
		}
	}
}
