package hello_test

import (
	"testing"

	"github.com/vhugo/hello"
)

func TestHello(t *testing.T) {
	for _, tc := range []struct {
		name      string
		arg, want string
	}{
		{
			name: "empty",
			want: "Hello, World!",
		},
		{
			name: "with argument",
			arg:  "Gophers",
			want: "Hello, Gophers!",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := hello.Hello(tc.arg)
			if tc.want != got {
				t.Fatalf("want %q, got %q", tc.want, got)
			}
		})
	}
}
