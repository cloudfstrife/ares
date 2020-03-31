package utils

import (
	"testing"
)

func TestJudge(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(Judge(10))
	}
}

func BenchmarkJudge(b *testing.B) {
	Judge(0)
	Judge(10)
	Judge(20)
	Judge(30)
	Judge(40)
	Judge(50)
	Judge(60)
	Judge(70)
	Judge(80)
	Judge(90)
	Judge(100)
}
