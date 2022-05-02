package str_rand

import (
	"testing"
)

func TestString(t *testing.T) {
	r1 := String(6)
	r2 := String(6)
	len1 := len(r1)
	len2 := len(r2)

	if len1 != 6 || len2 != 6 {
		t.Errorf("生成随机字符串长度不正确")
		return
	}

	if r1 == r2 {
		t.Errorf("两次生成的随机字符串一样！")
		return
	}

	t.Logf("结果正确: r1: %s; r2: %s", r1, r2)
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = String(10)
	}
}
