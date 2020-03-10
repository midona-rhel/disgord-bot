package disgordbot

import (
	"sync"
)

type relationMap struct {
	sync.RWMutex
	m map[string]interface{}
}

func newPermissionMap() *relationMap {
	return &relationMap{m: map[string]interface{}{}}
}

func (p *relationMap) getRaw(s string) (i interface{}) {
	p.RLock()
	i = p.m[s]
	p.RUnlock()
	return
}

func (p *relationMap) setRaw(s string, i interface{}) {
	p.Lock()
	p.m[s] = i
	p.Unlock()
}
