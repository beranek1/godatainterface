package godatainterface

type DataVersionEntry struct {
	Data      any   `json:"d"`
	Timestamp int64 `json:"t"`
}

type DataVersionLinked interface {
	Array() DataVersionLinkedArray
	Map() DataVersionLinkedMap
}

type DataVersionLinkedArray []DataVersionEntry
type DataVersionLinkedMap map[int64]any
