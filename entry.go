package godatainterface

type DataEntry interface {
	Data() any
}

type DataEntryVersioned interface {
	DataEntry
	Timestamp() int64
}

type DataEntryVersionedLinked interface {
	DataEntryVersioned
	GetAt(timestamp int64) any
	Array() []DataEntry
	Map() map[int64]any
}

type DataEntryVersionedLinkedArray []DataEntry

type DataEntryVersionedLinkedMap map[int64]any
