package telegraph

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
	"unicode/utf8"

	"github.com/brianvoe/gofakeit/v7"
)

// ShortName represent account name, helps users with several accounts remember
// which they are currently using. Displayed to the user above the
// "Edit/Publish" button on Telegra.ph, other users don't see this name.
type ShortName struct {
	shortName string // 1-32 characters
}

var ErrShortNameLength error = errors.New("unsupported length")

// NewShortName parse raw string as [ShortName] and validate it's length.
func NewShortName(raw string) (*ShortName, error) {
	if count := utf8.RuneCountInString(raw); count < 1 || 32 < count {
		return nil, fmt.Errorf("ShortName: %w: want 1-32 characters, got %d", ErrShortNameLength, count)
	}

	return &ShortName{raw}, nil
}

func (sn *ShortName) UnmarshalJSON(v []byte) error {
	unquoted, err := strconv.Unquote(string(v))
	if err != nil {
		return fmt.Errorf("ShortName: UnmarshalJSON: cannot unquote value '%s': %w", string(v), err)
	}

	result, err := NewShortName(unquoted)
	if err != nil {
		return fmt.Errorf("ShortName: UnmarshalJSON: cannot parse value '%s': %w", string(v), err)
	}

	*sn = *result

	return nil
}

func (sn ShortName) MarshalJSON() ([]byte, error) {
	if sn.shortName != "" {
		return []byte(strconv.Quote(sn.shortName)), nil
	}

	return nil, nil
}

func (sn ShortName) String() string {
	return sn.shortName
}

func (sn ShortName) GoString() string {
	return "telegraph.ShortName(" + sn.String() + ")"
}

func TestShortName(tb testing.TB) *ShortName {
	tb.Helper()

	return &ShortName{gofakeit.Username()}
}