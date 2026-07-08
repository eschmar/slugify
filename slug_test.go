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
		style      Style
		in, verify string
	}{
		{
			in:     " Importànt Dôcument (³)_copy final [36900fe]-compressed (((((8)))))---",
			verify: "Important-Document-3-copy-final-36900fe-compressed-8",
		},
		{
			in:     "Schwiizerdütsch & svensk omljud ä ö å behandling",
			verify: "Schwiizerdutsch-and-svensk-omljud-a-o-a-behandling",
		},
		{
			style:  German,
			in:     "Schwiizerdütsch & svensk omljud ä ö å behandling",
			verify: "Schwiizerduetsch-and-svensk-omljud-ae-oe-a-behandling",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Test '%s'", truncate(test.in, 10)), func(t *testing.T) {
			out := test.style.Ify(test.in)

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
