package autocomplete

import "testing"

func TestRebuildColonSeparatedArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{name: "standalone colon", args: []string{"a", ":", "b"}, want: []string{"a:b"}},
		{name: "trailing colon", args: []string{"config:", "get"}, want: []string{"config:get"}},
		{name: "repeated colon", args: []string{"a", ":", ":", "b"}, want: []string{"a::b"}},
		{name: "ordinary args", args: []string{"a", "b"}, want: []string{"a", "b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rebuildColonSeparatedArgs(tt.args)
			if len(got) != len(tt.want) {
				t.Fatalf("rebuildColonSeparatedArgs(%q) = %q, want %q", tt.args, got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("rebuildColonSeparatedArgs(%q) = %q, want %q", tt.args, got, tt.want)
				}
			}
		})
	}
}
