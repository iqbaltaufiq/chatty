package database

type ErrNoConfig struct{}

func (e *ErrNoConfig) Error() string {
	return "error: config is nil"
}
