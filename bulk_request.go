package elastic

type bulkRequest interface {
	OpType() string
	Source() ([]string, error)
}