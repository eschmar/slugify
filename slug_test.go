package slug

import (
	"fmt"
	"testing"
)

func truncate(in string, max int) string {
	if len(in) < max {
		return in
	}

	return in[:max] + "..."
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		in, verify string
	}{
		{
			" Importànt Dôcument (³)_copy final [36900fe]-compressed (((((8)))))---",
			"Important-Document-3-copy-final-36900fe-compressed-8",
		},
		{
			"Schwiizerdütsch & svensk omljud ä ö å behandling",
			"Schwiizerdutsch-and-svensk-omljud-a-o-a-behandling",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Test '%s'", truncate(test.in, 10)), func(t *testing.T) {
			out := Ify(test.in)

			if out != test.verify {
				t.Errorf("got '%s', want '%s'", out, test.verify)
			}
		})
	}
}

// 3871 ns/op
func BenchmarkSlugify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Ify(" Importànt Dôcument (³)_copy final [36900fe]-compressed (((((8)))))---")
	}
}
