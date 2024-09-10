package strings

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func ToLower(s string) string {
	// switch len(s) {
	// case 0:
	// 	// Fast path for empty strings
	// 	return ""
	// case 1:
	// 	// Fast path for single-byte strings that cannot contain UTF8.
	// 	if s[0] < 'A' || s[0] > 'Z' {
	// 		// No modification necessary.
	// 		return s
	// 	}
	// 	return string(s[0] + 'a' - 'A')
	// case 2:
	// 	// Fast path for either of:
	// 	//
	// 	//   s[0]  | s[1]
	// 	//   ------|------
	// 	//   UTF8  | UTF8
	// 	//   ASCII | ASCII

	// 	if s[0] >= utf8.RuneSelf {
	// 		//  1 UTF8 rune only.
	// 		r, size := utf8.DecodeLastRuneInString(s)
	// 		switch size {
	// 		case 1, 3: // Invalid encoding.
	// 			return strings.Map(unicode.ToLower, s)
	// 		case 2:
	// 			rl := unicode.To(unicode.LowerCase, r)
	// 			if rl == r {
	// 				// No modification necessary.
	// 				return s
	// 			}
	// 			return string(rl)
	// 		}
	// 	}

	// 	// 2 ASCII characters.
	// 	if (s[0] < 'A' || s[0] > 'Z') && (s[1] < 'A' || s[1] > 'Z') {
	// 		// No modification necessary.
	// 		return s
	// 	}

	// 	// One of the characters requires modification.
	// 	var b [2]byte
	// 	if s[0] >= 'A' && s[0] <= 'Z' {
	// 		b[0] = s[0] + 'a' - 'A'
	// 	} else {
	// 		b[0] = s[0]
	// 	}
	// 	if s[1] >= 'A' && s[1] <= 'Z' {
	// 		b[1] = s[1] + 'a' - 'A'
	// 	} else {
	// 		b[1] = s[1]
	// 	}
	// 	return string(b[:])
	// case 3:
	// 	// Fast path for either of:
	// 	//
	// 	//   s[0]  | s[1]  | s[2]
	// 	//   ------|-------|------
	// 	//   UTF8  | UTF8  | UTF8
	// 	//   UTF8  | UTF8  | ASCII
	// 	//   ASCII | UTF8  | UTF8
	// 	//   ASCII | ASCII | UTF8
	// 	//   ASCII | ASCII | ASCII

	// 	var b [3]byte
	// 	if s[0] >= utf8.RuneSelf {
	// 		// First byte is starting a UTF8 rune.
	// 		r, size := utf8.DecodeRuneInString(s)
	// 		switch size {
	// 		case 1: // Invalid encoding.
	// 			return strings.Map(unicode.ToLower, s)
	// 		case 2:
	// 			rl := unicode.To(unicode.LowerCase, r)
	// 			if rl == r && (s[2] < 'A' || s[2] > 'Z') {
	// 				// No modification necessary.
	// 				return s
	// 			}
	// 			// Write first and second bytes of the 2-byte UTF8 sequence.
	// 			b[0] = byte(rl>>6) | 0xC0
	// 			b[1] = byte(rl&0x3F) | 0x80
	// 			if s[2] >= 'A' && s[2] <= 'Z' {
	// 				// Write the last ASCII character with changed case.
	// 				b[2] = s[2] + 'a' - 'A'
	// 			} else {
	// 				b[2] = s[2]
	// 			}
	// 			return string(b[:])
	// 		}
	// 		// Single-rune string (3 byte UTF8 sequence).
	// 		rl := unicode.To(unicode.LowerCase, r)
	// 		if rl == r {
	// 			// No modification necessary.
	// 			return s
	// 		}
	// 		return string(rl)
	// 	} else if s[1] >= utf8.RuneSelf {
	// 		// Second byte is starting a UTF8 rune.
	// 		r, size := utf8.DecodeRuneInString(s[1:])
	// 		if size == 1 {
	// 			// Invalid encoding.
	// 			return strings.Map(unicode.ToLower, s)
	// 		}
	// 		if s[0] >= 'A' && s[0] <= 'Z' {
	// 			b[0] = s[0] + 'a' - 'A'
	// 		} else {
	// 			b[0] = s[0]
	// 		}
	// 		rl := unicode.To(unicode.LowerCase, r)
	// 		b[1] = byte(rl>>6) | 0xC0
	// 		b[2] = byte(rl&0x3F) | 0x80
	// 		if b[0] == s[0] && b[1] == s[1] && b[2] == s[2] {
	// 			return s
	// 		}
	// 		return string(b[:])
	// 	} else {
	// 		if (s[0] < 'A' || s[0] > 'Z') &&
	// 			(s[1] < 'A' || s[1] > 'Z') &&
	// 			(s[2] < 'A' || s[2] > 'Z') {
	// 			// No modification necessary.
	// 			return s
	// 		}

	// 		// 3 ASCII characters.
	// 		if s[0] >= 'A' && s[0] <= 'Z' {
	// 			b[0] = s[0] + 'a' - 'A'
	// 		} else {
	// 			b[0] = s[0]
	// 		}
	// 		if s[1] >= 'A' && s[1] <= 'Z' {
	// 			b[1] = s[1] + 'a' - 'A'
	// 		} else {
	// 			b[1] = s[1]
	// 		}
	// 		if s[2] >= 'A' && s[2] <= 'Z' {
	// 			b[2] = s[2] + 'a' - 'A'
	// 		} else {
	// 			b[2] = s[2]
	// 		}
	// 		return string(b[:])
	// 	}
	// }

	i, hasUpper := 0, false
	// Use a simpler algorithm for the first 32 bytes
	// which is more efficient for short strings.
	var buf [32]byte
	for ; i < len(s) && i < 32; i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			// Upper case ASCII
			hasUpper, buf[i] = true, s[i]+'a'-'A'
			continue
		} else if s[i] >= utf8.RuneSelf {
			return strings.Map(unicode.ToLower, s)
		}
		buf[i] = s[i]
	}
	if len(s) < 32 {
		if hasUpper {
			return string(buf[:i])
		}
		return s
	}
	var b strings.Builder
	pos := 0
	if hasUpper {
		pos = i
		b.Grow(len(s))
		b.Write(buf[:i])
	}
	// At this point, the first 32 bytes were processed.

	for i < len(s) {
		// Try to skip 4 ASCII characters effectively at once.
		sx := s[i:]
		if len(sx) > 4 {
			if lut[sx[0]] > 0 {
				goto CHECK
			}
			if lut[sx[1]] > 0 {
				i++
				goto CHECK
			}
			if lut[sx[2]] > 0 {
				i += 2
				goto CHECK
			}
			if lut[sx[3]] > 0 {
				i += 3
				goto CHECK
			}
			i += 4
			continue
		}

	CHECK:
		if s[i] >= 'A' && s[i] <= 'Z' {
			if b.Cap() == 0 {
				b.Grow(len(s))
			}
			if pos < i {
				b.WriteString(s[pos:i])
			}
			b.WriteByte(s[i] + 'a' - 'A')
			i++
			pos = i
			continue
		}
		if s[i] >= utf8.RuneSelf {
			r, size := utf8.DecodeRuneInString(s[i:])
			if r == utf8.RuneError {
				return strings.Map(unicode.ToLower, s)
			}
			l := unicode.To(unicode.LowerCase, r)
			if l == r {
				i += size
				continue
			}
			if b.Cap() == 0 {
				b.Grow(len(s))
			}
			if pos < i {
				b.WriteString(s[pos:i])
			}
			b.WriteRune(l)
			i += size // Advance by the size of the multi-byte character
			pos = i
			continue
		}
		i++
	}
	if pos < len(s) {
		if b.Cap() == 0 {
			return s
		}
		b.WriteString(s[pos:])
	}
	return b.String()
}

// lut maps upper-case ASCII to 1 and non-ASCII to 2.
// Other bytes are mapped to 0.
var lut = [256]uint8{
	'A': 1, 'B': 1, 'C': 1, 'D': 1, 'E': 1, 'F': 1, 'G': 1,
	'H': 1, 'I': 1, 'J': 1, 'K': 1, 'L': 1, 'M': 1, 'N': 1,
	'O': 1, 'P': 1, 'Q': 1, 'R': 1, 'S': 1, 'T': 1, 'U': 1,
	'V': 1, 'W': 1, 'X': 1, 'Y': 1, 'Z': 1,

	128: 2, // utf8.RuneSelf
	129: 2, 130: 2, 131: 2, 132: 2, 133: 2, 134: 2, 135: 2,
	136: 2, 137: 2, 138: 2, 139: 2, 140: 2, 141: 2, 142: 2,
	143: 2, 144: 2, 145: 2, 146: 2, 147: 2, 148: 2, 149: 2,
	150: 2, 151: 2, 152: 2, 153: 2, 154: 2, 155: 2, 156: 2,
	157: 2, 158: 2, 159: 2, 160: 2, 161: 2, 162: 2, 163: 2,
	164: 2, 165: 2, 166: 2, 167: 2, 168: 2, 169: 2, 170: 2,
	171: 2, 172: 2, 173: 2, 174: 2, 175: 2, 176: 2, 177: 2,
	178: 2, 179: 2, 180: 2, 181: 2, 182: 2, 183: 2, 184: 2,
	185: 2, 186: 2, 187: 2, 188: 2, 189: 2, 190: 2, 191: 2,
	192: 2, 193: 2, 194: 2, 195: 2, 196: 2, 197: 2, 198: 2,
	199: 2, 200: 2, 201: 2, 202: 2, 203: 2, 204: 2, 205: 2,
	206: 2, 207: 2, 208: 2, 209: 2, 210: 2, 211: 2, 212: 2,
	213: 2, 214: 2, 215: 2, 216: 2, 217: 2, 218: 2, 219: 2,
	220: 2, 221: 2, 222: 2, 223: 2, 224: 2, 225: 2, 226: 2,
	227: 2, 228: 2, 229: 2, 230: 2, 231: 2, 232: 2, 233: 2,
	234: 2, 235: 2, 236: 2, 237: 2, 238: 2, 239: 2, 240: 2,
	241: 2, 242: 2, 243: 2, 244: 2, 245: 2, 246: 2, 247: 2,
	248: 2, 249: 2, 250: 2, 251: 2, 252: 2, 253: 2, 254: 2,
	255: 2,
}
