package godatainterface

type DataCollection interface {
	Create(store string) error
	Get(store string, key string) (any, error)
	Put(store string, key string, value any) error
}

type DataCollectionVersioned interface {
	DataCollection
	GetAt(store string, key string, timestamp int64) (any, error)
	PutAt(store string, key string, value any, timestamp int64) error
}

type DataCollectionVersionedRange interface {
	DataCollectionVersioned
	Range(store string, key string, start int64, end int64) (map[int64]any, error)
}

type DataCollectionVersionedRangeFrom interface {
	DataCollectionVersionedRange
	From(store string, key string, start int64) (map[int64]any, error)
}

type DataCollectionVersionedRangeFromInterval interface {
	DataCollectionVersionedRangeFrom
	RangeInterval(store string, key string, start int64, end int64, interval int64) (map[int64]any, error)
	FromInterval(store string, key string, start int64, interval int64) (map[int64]any, error)
}
