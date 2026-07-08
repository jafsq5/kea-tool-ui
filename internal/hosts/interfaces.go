package hosts

type Repository interface {
	List() ([]Host, error)
	Add(Host) error
	Delete(string) error
	Exists(string) (bool, error)
}

type Reloader interface {
	Reload() error
}
