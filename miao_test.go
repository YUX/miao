package miao

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiao(t *testing.T) {
	s := "ffe 测试🤯、n\n s!!!！！gsler"
	assert.Equal(t, s, Decode(Encode(s)))
}
