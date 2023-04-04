package godatainterface

type DataEntry interface {
	Get() any
}

type DataEntryVersioned interface {
	DataEntry
	GetTimestamp() int64
}

type DataEntryVersionedLinked interface {
	DataEntryVersioned
	GetAt(timestamp int64) any
	Array() DataEntryVersionedLinkedArray
	Map() DataEntryVersionedLinkedMap
}

type DataEntryVersionedLinkedArray []any

type DataEntryVersionedLinkedMap map[int64]any
