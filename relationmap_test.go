package disgordbot

import (
	"testing"
)

func HammerRead(p *relationMap, loops int, cdone chan<- bool) {
	for i := 0; i < loops; i++ {
		p.getRaw("")
	}
	cdone <- true
}

func HammerWrite(p *relationMap, loops int, cdone chan<- bool) {
	for i := 0; i < loops; i++ {
		p.setRaw("", true)
	}
	cdone <- true
}

func Test_permissionMap_getRaw(t *testing.T) {
	c := make(chan bool)
	p := newPermissionMap()
	for i := 0; i < 10; i++ {
		go HammerRead(p, 1000, c)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func Test_permissionMap_setRaw(t *testing.T) {
	c := make(chan bool)
	p := newPermissionMap()
	for i := 0; i < 10; i++ {
		go HammerWrite(p, 1000, c)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func Test_permissionMap_mixedReadWrites(t *testing.T) {
	c := make(chan bool)
	p := newPermissionMap()
	for i := 0; i < 10; i++ {
		go HammerWrite(p, 1000, c)
		go HammerRead(p, 1000, c)
	}
	for i := 0; i < 20; i++ {
		<-c
	}
}
