package mynet

import (
	"testing"
)

func TestSendMsg(t *testing.T) {
	x := setupTestCase(t)
	SendMsg()
	x(t)
}

func setupTestCase(t *testing.T) func(t *testing.T) {
	go StartServer()
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
		StopServer()
	}
}
