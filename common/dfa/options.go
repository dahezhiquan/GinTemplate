package dfa

const (
	StoreMemory = iota
	StoreMysql
	StoreMongo
)

const (
	FilterDfa = iota
)

type StoreOption struct {
	Type uint32
	//MysqlConfig *store.MysqlConfig
	//MongoConfig *store.MongoConfig
}

type FilterOption struct {
	Type uint32
}
