package tmhash_test

import (
	"crypto/sha256"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/crypto/tmhash"
)

func TestHash(t *testing.T) {
	testVector := []byte("abc")
	hasher := tmhash.New()
	_, err := hasher.Write(testVector)
	require.NoError(t, err)
	bz := hasher.Sum(nil)

	bz2 := tmhash.Sum(testVector)

	hasher = sha256.New()
	_, err = hasher.Write(testVector)
	require.NoError(t, err)
	bz3 := hasher.Sum(nil)

	assert.Equal(t, bz, bz2)
	assert.Equal(t, bz, bz3)
}

func TestHashTruncated(t *testing.T) {
	testVector := []byte("abc")
	hasher := tmhash.NewTruncated()
	_, err := hasher.Write(testVector)
	require.NoError(t, err)
	bz := hasher.Sum(nil)

	bz2 := tmhash.SumTruncated(testVector)

	hasher = sha256.New()
	_, err = hasher.Write(testVector)
	require.NoError(t, err)
	bz3 := hasher.Sum(nil)
	bz3 = bz3[:tmhash.TruncatedSize]

	assert.Equal(t, bz, bz2)
	assert.Equal(t, bz, bz3)
}

func TestValidateSHA256(t *testing.T) {
	// Valid 64-character lowercase hex string
	validLower := strings.Repeat("a1b2c3d4", 8) // 64 chars

	// Valid 64-character uppercase hex string
	validUpper := strings.Repeat("A1B2C3D4", 8) // 64 chars

	// Valid 64-character mixed case hex string
	validMixed := "aAbBcCdDeEfF0123456789aAbBcCdDeEfF0123456789aAbBcCdDeEfF01234567"

	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid lowercase hex",
			input:   validLower,
			wantErr: false,
		},
		{
			name:    "valid uppercase hex",
			input:   validUpper,
			wantErr: false,
		},
		{
			name:    "valid mixed case hex",
			input:   validMixed,
			wantErr: false,
		},
		{
			name:    "valid all zeros",
			input:   strings.Repeat("0", 64),
			wantErr: false,
		},
		{
			name:    "valid all f's",
			input:   strings.Repeat("f", 64),
			wantErr: false,
		},
		{
			name:    "too short - 63 chars",
			input:   strings.Repeat("a", 63),
			wantErr: true,
			errMsg:  "expected 64 characters, but have 63",
		},
		{
			name:    "too long - 65 chars",
			input:   strings.Repeat("a", 65),
			wantErr: true,
			errMsg:  "expected 64 characters, but have 65",
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
			errMsg:  "expected 64 characters, but have 0",
		},
		{
			name:    "invalid char g in middle",
			input:   strings.Repeat("a", 31) + "g" + strings.Repeat("a", 32),
			wantErr: true,
			errMsg:  "contains non-hexadecimal characters",
		},
		{
			name:    "invalid char z at start",
			input:   "z" + strings.Repeat("a", 63),
			wantErr: true,
			errMsg:  "contains non-hexadecimal characters",
		},
		{
			name:    "space in middle",
			input:   strings.Repeat("a", 32) + " " + strings.Repeat("a", 31),
			wantErr: true,
			errMsg:  "contains non-hexadecimal characters",
		},
		{
			name:    "special character",
			input:   strings.Repeat("a", 63) + "!",
			wantErr: true,
			errMsg:  "contains non-hexadecimal characters",
		},
		{
			name:    "newline character",
			input:   strings.Repeat("a", 63) + "\n",
			wantErr: true,
			errMsg:  "contains non-hexadecimal characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tmhash.ValidateSHA256(tt.input)

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
