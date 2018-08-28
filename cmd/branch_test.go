package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareBranchName(t *testing.T) {
	branchName := "__%$test__$$$___name__"
	expectedBranchName := "test_name"

	preparedBrancName := prepareBranchName(branchName)

	assert.Equal(t, expectedBranchName, preparedBrancName, "should remove not acceptable symbols")
}
