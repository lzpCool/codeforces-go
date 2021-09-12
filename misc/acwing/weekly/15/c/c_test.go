// Code generated by copypasta/template/acwing/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`1 2
QW`,
			`none`,
		},
		{
			`2 2
ER
WQ`,
			`infinity`,
		},
		{
			`5 5
QWERQ
QWERW
QWERE
QQERR
RREWQ`,
			`4`,
		},
		
	}
	target :=  -1
	testutil.AssertEqualStringCase(t, testCases, target, run)
}