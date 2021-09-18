// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`"aaabcccd"`, `2`, 
			`4`,
		},
		{
			`"aabbaa"`, `2`, 
			`2`,
		},
		{
			`"aaaaaaaaaaa"`, `0`, 
			`3`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, getLengthOfOptimalCompression, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-199/problems/string-compression-ii/