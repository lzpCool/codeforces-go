// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_b(t *testing.T) {
	targetCaseNum :=2
	if err := testutil.RunLeetCodeFuncWithFile(t, partitionString, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-310/problems/optimal-partition-of-string/