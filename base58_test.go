package base58_test

import (
	base58 "github.com/jamieabc/base58-encode-decode"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {
	result := base58.Encode("Cat")
	assert.Equal(t, "PdgX", result, "wrong encode")
}

func TestDecode(t *testing.T) {
	result := base58.Decode("PdgX")
	assert.Equal(t, "Cat", result, "wrong decode")
}
