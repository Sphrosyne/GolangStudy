package auth

import "testing"

func Test_authMethod(t *testing.T) {
	type args struct {
		method string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"1", args{"cas"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authMethod(tt.args.method)
		})
	}
}
