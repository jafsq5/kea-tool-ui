package hosts

import (
	"errors"
	"strings"
)

type Repository interface {
    List() ([]Host, error)
    Add(Host) error
    Delete(string) error
    Exists(string) (bool, error)
}

func New(path string) *Repository {
	return &Repository{
		Path: path,
	}
}

func (r *Repository) List() ([]Host, error) {
	return ReadFile(r.Path)
}

func (r *Repository) Exists(mac string) (bool, error) {

	hosts, err := r.List()
	if err != nil {
		return false, err
	}

	for _, h := range hosts {

		if strings.EqualFold(h.HWAddress, mac) {
			return true, nil
		}
	}

	return false, nil
}

func (r *Repository) Add(host Host) error {

	hosts, err := r.List()
	if err != nil {
		return err
	}

	for _, h := range hosts {

		if strings.EqualFold(h.HWAddress, host.HWAddress) {
			return errors.New("MAC already exists")
		}
	}

	host.ClientClasses = []string{"reserved_class"}

	hosts = append(hosts, host)

	return WriteFile(r.Path, hosts)
}

func (r *Repository) Delete(mac string) error {

	hosts, err := r.List()
	if err != nil {
		return err
	}

	out := make([]Host, 0, len(hosts))

	for _, h := range hosts {

		if strings.EqualFold(h.HWAddress, mac) {
			continue
		}

		out = append(out, h)
	}

	return WriteFile(r.Path, out)
}
