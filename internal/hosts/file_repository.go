package hosts

import (
	"errors"
	"strings"
)

type FileRepository struct {
	Path string
}

func NewFileRepository(path string) *FileRepository {
	return &FileRepository{
		Path: path,
	}
}

func (r *FileRepository) List() ([]Host, error) {
	return ReadFile(r.Path)
}

func (r *FileRepository) Exists(mac string) (bool, error) {

	list, err := r.List()
	if err != nil {
		return false, err
	}

	for _, h := range list {
		if strings.EqualFold(h.HWAddress, mac) {
			return true, nil
		}
	}

	return false, nil
}

func (r *FileRepository) Add(host Host) error {

	list, err := r.List()
	if err != nil {
		return err
	}

	for _, h := range list {
		if strings.EqualFold(h.HWAddress, host.HWAddress) {
			return errors.New("MAC already exists")
		}
	}

	if len(host.ClientClasses) == 0 {
		host.ClientClasses = []string{"reserved_class"}
	}

	list = append(list, host)

	return WriteFile(r.Path, list)
}

func (r *FileRepository) Delete(mac string) error {

	list, err := r.List()
	if err != nil {
		return err
	}

	out := make([]Host, 0, len(list))

	for _, h := range list {

		if strings.EqualFold(h.HWAddress, mac) {
			continue
		}

		out = append(out, h)
	}

	return WriteFile(r.Path, out)
}
