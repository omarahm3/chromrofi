package browser

type Browser interface {
	Close() error
	GetHistoryLocation() string
	GetLocalState() (*LocalState, error)
}
