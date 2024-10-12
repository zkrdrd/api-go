package locker

import (
	"sync"
)

type Locker struct {
	// для безопасного обращения
	workMX *sync.Mutex

	// для безопасного обращения к карте (изменения баланса адресов)
	value map[string]*sync.Mutex
}

func (l *Locker) Lock(address string) {
	l.workMX.Lock()

	v, ok := l.value[address]
	if !ok {
		v = &sync.Mutex{}
		l.value[address] = v
	}

	l.workMX.Unlock()

	v.Lock()
}

func (l *Locker) Unlock(address string) {
	l.workMX.Lock()

	v, ok := l.value[address]
	if !ok {
		return
	}

	l.workMX.Unlock()

	v.Unlock()
}
