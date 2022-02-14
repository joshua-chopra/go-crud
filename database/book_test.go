package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyBook(t *testing.T) {
	var b Book
	assert.True(t, b.IsEmpty())
}
