package godatainterface

type DataSource interface {
	Get(key string) (any, error)
	List() ([]string, error)
}
