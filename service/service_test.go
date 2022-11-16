package service

import (
	"fmt"
	"testing"
)

func TestMeasure(t *testing.T) {
	for _, provider := range []string{"speedtest", "fast"} {
		func(provider string) {
			t.Run(
				fmt.Sprintf("%s provider test", provider),
				func(t *testing.T) {
					t.Parallel()

					p, err := ParseProvider(provider)
					assertNilError(err, t)

					s, err := New(p)
					assertNilError(err, t)

					res, err := s.Measure()
					assertNilError(err, t)

					t.Log(res)
				},
			)
		}(provider)
	}
}

func TestParseProviderError(t *testing.T) {
	_, err := ParseProvider("wrong-provider")
	if err != nil {
		return
	}

	t.Fatalf("unknown provider is not handled")
}

func TestNewProviderError(t *testing.T) {
	_, err := New(Provider(0))
	if err != nil {
		return
	}

	t.Fatalf("unknown provider is not handled")
}

func BenchmarkMeasure(b *testing.B) {
	for _, provider := range []string{"speedtest", "fast"} {
		func(provider string) {
			b.Run(
				fmt.Sprintf("%s provider benchmark", provider),
				func(b *testing.B) {
					p, err := ParseProvider(provider)
					assertNilError(err, b)

					s, err := New(p)
					assertNilError(err, b)

					for i := 0; i < b.N; i++ {
						_, err = s.Measure()
						assertNilError(err, b)
					}
				},
			)
		}(provider)
	}
}

func assertNilError(err error, t testing.TB) {
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
