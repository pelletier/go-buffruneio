package buffruneio

import (
	"strings"
	"testing"
)

func assumeRune(t *testing.T, rd *Reader, r rune) {
	gotRune, err := rd.ReadRune()
	if err != nil {
		t.Fatal("unexpected error", err)
	}
	if gotRune != r {
		t.Fatal("got", string(gotRune),
			"(", []byte(string(gotRune)), ")",
			"expected", string(r),
			"(", []byte(string(r)), ")")
	}
}

func TestReadString(t *testing.T) {
	s := "hello"
	rd := NewReader(strings.NewReader(s))

	assumeRune(t, rd, 'h')
	assumeRune(t, rd, 'e')
	assumeRune(t, rd, 'l')
	assumeRune(t, rd, 'l')
	assumeRune(t, rd, 'o')
	assumeRune(t, rd, EOF)
}

func TestMultipleEOF(t *testing.T) {
	s := ""
	rd := NewReader(strings.NewReader(s))

	assumeRune(t, rd, EOF)
	assumeRune(t, rd, EOF)
}

func TestUnread(t *testing.T) {
	s := "ab"
	rd := NewReader(strings.NewReader(s))

	assumeRune(t, rd, 'a')
	assumeRune(t, rd, 'b')
	rd.UnreadRune()
	assumeRune(t, rd, 'b')
	assumeRune(t, rd, EOF)
}

func TestUnreadEOF(t *testing.T) {
	s := ""
	rd := NewReader(strings.NewReader(s))

	rd.UnreadRune()
	assumeRune(t, rd, EOF)
	assumeRune(t, rd, EOF)
	rd.UnreadRune()
	assumeRune(t, rd, EOF)
}

func TestForget(t *testing.T) {
	s := "hello"
	rd := NewReader(strings.NewReader(s))

	assumeRune(t, rd, 'h')
	assumeRune(t, rd, 'e')
	assumeRune(t, rd, 'l')
	assumeRune(t, rd, 'l')
	rd.Forget()
	if rd.UnreadRune() != ErrNoRuneToUnread {
		t.Fatal("no rune should be available")
	}
}

func TestForgetEmpty(t *testing.T) {
	s := ""
	rd := NewReader(strings.NewReader(s))

	rd.Forget()
	assumeRune(t, rd, EOF)
	rd.Forget()
}
