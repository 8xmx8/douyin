package utils

import "testing"

func TestGetId(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log("ID获取测试: ", GetId(2, 20230724))
	}
}
func TestRandString(t *testing.T) {
	for i := 1; i < 6; i++ {
		t.Logf("长度: %d,随机字符串测试: %s\n", i*11, RandString(i*11))
	}
}
