package dfa

import (
	"GinTemplate/common/dfa/filter"
	"GinTemplate/common/dfa/store"
	"fmt"
	"sync"
)

type Manager struct {
	store  store.Model
	filter filter.Model

	filterMux sync.RWMutex
}

func NewFilter(storeOption StoreOption, filterOption FilterOption) *Manager {
	var filterStore store.Model
	var myFilter filter.Model

	switch storeOption.Type {
	case StoreMemory:
		filterStore = store.NewMemoryModel()
	}

	switch filterOption.Type {
	case FilterDfa:
		dfaModel := filter.NewDfaModel()

		go dfaModel.Listen(filterStore.GetAddChan(), filterStore.GetDelChan())

		myFilter = dfaModel
	}

	return &Manager{
		store:  filterStore,
		filter: myFilter,
	}
}

func (m *Manager) GetStore() store.Model {
	return m.store
}

func (m *Manager) GetFilter() filter.Model {
	m.filterMux.RLock()
	myFilter := m.filter
	m.filterMux.RUnlock()
	return myFilter
}

// 提供全局调用的接口

func CheckWordByDFA(word string) bool {

	filterManager := NewFilter(
		StoreOption{
			Type: StoreMemory,
		},
		FilterOption{
			Type: FilterDfa,
		},
	)

	// 加载字典

	err := filterManager.GetStore().LoadDictPath("common/dfa/dict/default_dict.txt")
	if err != nil {
		fmt.Println(err)
		return false
	}

	return filterManager.GetFilter().IsSensitive(word)
}
