//go:build !disable_pgv

package corev3

import "testing"

func TestHTTPHealthCheckPatternsUseExplicitControlCharacterRanges(t *testing.T) {
	patterns := map[string]interface{ MatchString(string) bool }{
		"host": _HealthCheck_HttpHealthCheck_Host_Pattern,
		"path": _HealthCheck_HttpHealthCheck_Path_Pattern,
	}

	for name, pattern := range patterns {
		t.Run(name, func(t *testing.T) {
			if !pattern.MatchString("example.com\t/health") {
				t.Fatal("pattern rejected visible text and the allowed horizontal tab")
			}
			for _, input := range []string{"bad\x00value", "bad\bvalue", "bad\nvalue", "bad\x7fvalue"} {
				if pattern.MatchString(input) {
					t.Fatalf("pattern accepted forbidden control character in %q", input)
				}
			}
		})
	}
}
