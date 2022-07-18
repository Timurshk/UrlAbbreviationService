package hanglers

import "testing"

func TestShortening(t *testing.T) {
	type args struct {
		Url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Shortening(tt.args.Url); got != tt.want {
				t.Errorf("Shortening() = %v, want %v", got, tt.want)
			}
		})
	}
}
