package remote_package

import "testing"
import cm "src/github.com/easierway/concurrent_map"

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(10)
	m.Set(cm.StrKey("zhang"), 10)
	t.Log(m.Get(cm.StrKey("zhang")))
}