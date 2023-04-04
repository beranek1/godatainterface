package godatainterface

type DataStore interface {
	Get(key string) (any, error)
	Put(key string, value any) error
}

type DataStoreVersioned interface {
	DataStore
	GetAt(key string, timestamp int64) (any, error)
	PutAt(key string, value any, timestamp int64) error
}

type DataStoreVersionedRange interface {
	DataStoreVersioned
	Range(key string, start int64, end int64) (map[int64]any, error)
}

type DataStoreVersionedRangeFrom interface {
	DataStoreVersionedRange
	From(key string, start int64) (map[int64]any, error)
}

type DataStoreVersionedRangeFromInterval interface {
	DataStoreVersionedRangeFrom
	RangeInterval(key string, start int64, end int64, interval int64) (map[int64]any, error)
	FromInterval(key string, start int64, interval int64) (map[int64]any, error)
}

type DataStoreVersionedRangeFromIntervalArray interface {
	DataStoreVersionedRangeFrom
	RangeArray(key string, start int64, end int64) ([]any, error)
	FromArray(key string, start int64) ([]any, error)
	RangeIntervalArray(key string, start int64, end int64, interval int64) ([]any, error)
	FromIntervalArray(key string, start int64, interval int64) ([]any, error)
}
