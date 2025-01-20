package test

import (
	"go-gin/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTrue(t *testing.T) {
	assert.True(t, util.IsTrue(11), "number not zero should be true")
	assert.True(t, util.IsFalse(0), "number zero should be false")
	assert.True(t, util.IsTrue("hello"), "string should be true")
	assert.True(t, util.IsFalse(""), "empty string should be false")
	assert.True(t, util.IsTrue(1.0), "float number should be true")
	assert.True(t, util.IsFalse(0.0), "float number should be false")
}

func TestConditional(t *testing.T) {

	assert.Equal(t, util.When(true, "hello", ""), "hello", "condtional should return first value")
	assert.Equal(t, util.When(false, "hello", "aa"), "aa", "condtional should return second value")
}
