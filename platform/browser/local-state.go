package browser

type LocalState interface {
	HasProfile(id string) bool
	GetProfileKey(id string) string
}
