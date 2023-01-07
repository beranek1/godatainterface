package godatainterface

type DataDestination interface {
	Put(key string, value any) error
}
