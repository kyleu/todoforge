// Content managed by Project Forge, see [projectforge.md] for details.
package database

import (
	"slices"
	"sync"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app/util"
)

var (
	serviceRegistry   = map[string]*Service{}
	serviceRegistryMu = sync.Mutex{}
)

func register(s *Service, logger util.Logger) {
	serviceRegistryMu.Lock()
	defer serviceRegistryMu.Unlock()
	if _, ok := serviceRegistry[s.Key]; ok {
		logger.Warnf("double registration for database [%s]", s.Key)
	}
	serviceRegistry[s.Key] = s
}

func unregister(s *Service) {
	delete(serviceRegistry, s.Key)
}

func RegistryGet(key string) (*Service, error) {
	ret, ok := serviceRegistry[key]
	if !ok {
		return nil, errors.Errorf("no database registered with key [%s]", key)
	}
	return ret, nil
}

func RegistryKeys() []string {
	ret := lo.Keys(serviceRegistry)
	slices.Sort(ret)
	return ret
}
