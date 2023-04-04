package godatainterface

type DataEntry interface {
	Value() any
}

type DataEntryVersioned interface {
	DataEntry
	Key() int64
}

type DataEntryVersionedLinked interface {
	DataEntryVersioned
	GetAt(timestamp int64) any
	Array() []DataEntry
	Map() map[int64]any
}

type DataEntryVersionedLinkedArray []DataEntry

type DataEntryVersionedLinkedMap map[int64]any
