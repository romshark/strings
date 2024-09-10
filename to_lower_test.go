package strings_test

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"

	romsharkstrings "github.com/romshark/strings"

	"github.com/stretchr/testify/require"
)

var implementationsToLower = []struct {
	name string
	fn   func(string) string
}{
	{"std", strings.ToLower},
	{"opt", romsharkstrings.ToLower},
}

func TestToLower(t *testing.T) {
	f := func(expect, input string) {
		t.Helper()
		std := strings.ToLower(input)
		require.Equal(t, expect, std,
			"incorrect expectation when testing against Go standard library "+
				"expected %v, received: %v", []byte(expect), []byte(std))
		for _, fn := range implementationsToLower {
			mod := fn.fn(input)
			require.Equal(t, expect, mod,
				"incorrect expectation for %q when testing against Go standard library "+
					"expected %v, received: %v", fn.name, []byte(expect), []byte(mod))
		}
	}

	// Zero byte
	f("\u0000", "\u0000")
	f("\u0000\u0000", "\u0000\u0000")
	f("\u0000\u0000\u0000", "\u0000\u0000\u0000")
	f("\u0000\u0000\u0000\u0000", "\u0000\u0000\u0000\u0000")

	// ASCII (no modification)
	f("", "")
	f(" ", " ")
	f("?", "?")
	f("??", "??")
	f("???", "???")
	f("a", "a")
	f("ab", "ab")
	f("aaa", "aaa")
	f("aaaa", "aaaa")

	// ASCII (case change)
	f("a", "A")
	f("ac", "aC")
	f("db", "Db")
	f("123", "123")
	f("aab", "aaB")
	f("aca", "aCa")
	f("daa", "Daa")
	f("aee", "aEE")
	f("ffa", "FFa")
	f("gaa", "Gaa")
	f("hah", "HaH")
	f("iii", "III")
	f("jjjj", "jjjj")
	f("jjjk", "jjjK")
	f("jjkk", "jjKK")
	f("jkkk", "jKKK")
	f("kkkk", "KKKK")
	f("kkkj", "KKKj")
	f("kkjj", "KKjj")
	f("kjjj", "Kjjj")
	f("kjjk", "KjjK")
	f("jkkj", "jKKj")
	f("jjkj", "jjKj")
	f("jkjj", "jKjj")

	// UTF8 (no modification)
	f("€", "€")
	f("€€", "€€")
	f("€€€", "€€€")
	f("🐭", "🐭")
	f("🧠", "🧠")
	f("🧠🧠", "🧠🧠")
	f("🦄🧠🚀", "🦄🧠🚀")

	// UTF8 (case change)
	f("ж", "Ж")
	f("ы", "Ы")
	f("йй", "йй")
	f("щы", "щЫ")
	f("шы", "Шы")
	f("дд", "ДД")
	f("яяя", "ЯЯЯ")
	f("яця", "ЯцЯ")
	f("южю", "юЖю")
	f("пппп", "пппп")
	f("пппü", "пппÜ")
	f("ппü", "ппÜ")
	f("пü", "пÜ")
	f("ü", "Ü")
	f("üп", "Üп")
	f("üпп", "Üпп")
	f("üппп", "Üппп")
	f("üппü", "ÜппÜ")
	f("пüп", "пÜп")
	f("ппüп", "ппÜп")
	f("пüпп", "пÜпп")

	// Mixed UTF8 and ASCII
	f("xö", "xö")
	f("xö", "xÖ")
	f("xö", "XÖ")
	f("xö", "Xö")

	f("öx", "öx")
	f("öx", "Öx")
	f("öx", "ÖX")
	f("öx", "öX")

	f("öxx", "öxx")
	f("öxx", "öxX")
	f("öxx", "öXx")
	f("öxx", "öXX")
	f("öxx", "ÖXX")
	f("öxx", "ÖXx")
	f("öxx", "Öxx")
	f("öxx", "ÖxX")

	f("test text that's longer than 8 bytes.", "test text that's longer than 8 bytes.")
	f("test text that's longer than 8 bytes.", "Test Text That's Longer Than 8 Bytes.")

	f("utf-8 текст длиннее 8 байт (mixed)", "UTF-8 текст длиннее 8 байт (Mixed)")

	// Invalid UTF8
	invalidFirstByte := "\xF5\x80\x80\x80" // First byte is out of valid UTF-8 range
	f("����", invalidFirstByte)
}

var GS string

//go:embed testdata/lorem_ipsum.txt
var txtLoremIpsum string

//go:embed testdata/lorem_ipsum_lowercase.txt
var txtLoremIpsumLowercase string

//go:embed testdata/lorem_ipsum_utf8end.txt
var txtLoremIpsumUTF8End string

//go:embed testdata/lorem_ipsum_lowercase_utf8end.txt
var txtLoremIpsumUTF8EndLowercase string

//go:embed testdata/mixed-764b.txt
var txtMixed764b string

//go:embed testdata/mixed-764b-low.txt
var txtMixed764bLowercase string

//go:embed testdata/wiki-japan-en.html.txt
var txtWikiJapanEnHTML string

//go:embed testdata/wiki-japan-jp.html.txt
var txtWikiJapanJpHTML string

//go:embed testdata/romeojuliet.txt
var txtRomeoJuliet string

//go:embed testdata/romeojuliet-low.txt
var txtRomeoJulietLowercase string

func init() {
	if utf8.RuneCountInString(txtRomeoJuliet) != len(txtRomeoJuliet) {
		panic(fmt.Errorf("txtRomeoJuliet contains UTF8"))
	}
	if utf8.RuneCountInString(txtRomeoJulietLowercase) != len(txtRomeoJulietLowercase) {
		panic(fmt.Errorf("txtRomeoJulietLowercase contains UTF8"))
	}
	if strings.ToLower(txtRomeoJulietLowercase) != txtRomeoJulietLowercase {
		panic(fmt.Errorf("txtRomeoJulietLowercase is not all lower case"))
	}
}

var benchmarks = []struct {
	name  string
	input string
}{
	{"empty________________", ""},
	{"ascii-1______________", "A"},
	{"ascii-1-low__________", "a"},
	{"ascii-2______________", "AB"},
	{"ascii-2-low__________", "ab"},
	{"ascii-3______________", "ABC"},
	{"ascii-3-low__________", "abc"},
	{"ascii-7______________", "VAR_ENV"},
	{"ascii-7-low__________", "var_env"},
	{"ascii-8______________", "VAR_ENV1"},
	{"ascii-8-low__________", "var_env1"},
	{"ascii-9______________", "VAR_ENV_2"},
	{"ascii-9-low__________", "var_env_2"},
	{"ascii-33-capital_____", "This Is A Capitalized String Test"},
	{"ascii-33-most-up_____", "THIS IS A CAPITALIZED STRING TEST"},
	{"ascii-33-up__________", "THISXISXAXCAPITALIZEDXSTRINGXTEST"},
	{"ascii-33-low_________", "this is a capitalized string test"},
	{"ascii-49-capital_____", "This Is A Capitalized String For Benchmarks (A-Z)"},
	{"ascii-49-up---_______", "THISISANALLUPPERCASESTRINGFORBENCHMARKINGPURPOSES"},
	{"ascii-49-low---------", "this is a capitalized string for benchmarks (a-z)"},
	{"ascii-loremipsum_____", txtLoremIpsum},
	{"ascii-loremipsum-low_", txtLoremIpsumLowercase},
	{"loremipsum_u8end_____", txtLoremIpsumUTF8End},
	{"loremipsum-low_u8end_", txtLoremIpsumUTF8EndLowercase},
	{"romeo-juliet_________", txtRomeoJuliet},
	{"romeo-juliet-low_____", txtRomeoJulietLowercase},
	{"utf8_4-1_____________", "😎"},
	{"utf8_2-1_____________", "Ж"},
	{"utf8_2-1-low_________", "ж"},
	{"utf8_3-1_____________", "€"},
	{"utf8_3-3_____________", "€€€"},
	{
		"utf8-japanese________",
		"日本語 は、日本国内や、かつての日本領だった国、そして国外移民や移住者を含む日本人同士の間で使",
	},
	{"hallo________________", txtMixed764b},
	{"mixed-764b___________", txtMixed764b},
	{"32-ascii_1-utf8______", "12345678901234567890123456789012て"},
	{"mixed-764b-low_______", txtMixed764bLowercase},
	{"wiki-japan-en-html___", txtWikiJapanEnHTML},
	{"wiki-japan-jp-html___", txtWikiJapanJpHTML},
}

func TestBenchmarks(t *testing.T) {
	for _, td := range benchmarks {
		t.Run(td.name, func(t *testing.T) {
			for _, fn := range implementationsToLower {
				t.Run(fn.name, func(t *testing.T) {
					std := strings.ToLower(td.input)
					mod := fn.fn(td.input)
					require.Equal(t, std, mod)
				})
			}
		})
	}
}

func BenchmarkAll(b *testing.B) {
	for _, bd := range benchmarks {
		b.Run(bd.name, func(b *testing.B) {
			fmt.Println("")
			for _, fn := range implementationsToLower {
				b.Run(fn.name, func(b *testing.B) {
					for range b.N {
						GS = fn.fn(bd.input)
					}
				})
			}
		})
	}
}

var fBenchFunc = flag.String("func", "std", "implementation function name")

func BenchmarkToLower(b *testing.B) {
	var fn func(string) string
	switch *fBenchFunc {
	case "std":
		fn = strings.ToLower
	case "opt":
		fn = romsharkstrings.ToLower
	default:
		panic(fmt.Errorf("unknown function: %q", *fBenchFunc))
	}
	// This benchmark makes it easy to use benchstat.
	// First run with `strings.ToLower` using `go test -run ToLower -count 10 > old.txt`,
	// then replace it with the other function
	for _, bd := range benchmarks {
		b.Run(bd.name, func(b *testing.B) {
			for range b.N {
				// Replace this with the
				GS = fn(bd.input)
			}
		})
	}
}
