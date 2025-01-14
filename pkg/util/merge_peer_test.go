package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergePeers(t *testing.T) {
	t1 := MergePeers([]string{}, []string{"a"})
	assert.Equal(t, []string{"a"}, t1)

	t2 := MergePeers([]string{"a"}, []string{"a"})
	assert.Equal(t, []string{"a"}, t2)

	t3 := MergePeers([]string{"a", "b"}, []string{"a", "c"})
	assert.Equal(t, []string{"a", "b", "c"}, t3)
}
