package utils

import (
	"testing"
	"time"
)

func TestGetId(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			t.Log("ID获取测试: ", GetId(2, 20230724))
		}()
	}
	time.Sleep(3 * time.Second)
}

func TestRandString(t *testing.T) {
	for i := 1; i < 6; i++ {
		t.Logf("长度: %d,随机字符串测试: %s", i*11, RandString(i*11))
	}
}
