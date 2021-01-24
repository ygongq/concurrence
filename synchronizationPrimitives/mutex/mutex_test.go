package mutex

import (
	"testing"
)

// go test -v mutex_test.go mutex.go -run TestDemo
func TestDemo(t *testing.T) {
	t.Log("====== NOT Mutex ======")
	NotMutex()
	t.Log("====== Mutex ======")
	Mutex()
	t.Log("====== Postting ======")
	GetCount(10, 10000)
}

func TestNotMutex(t *testing.T) {
	NotMutex()
}
