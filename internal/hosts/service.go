package hosts



type Service struct {
    repo     *Repository
    reloader Reloader
}

func NewService(repo *Repository, r Reloader) *Service {

    return &Service{
        repo: repo,
        reloader: r,
    }
}

func (s *Service) List() ([]Host, error) {
	return s.repo.List()
}

func (s *Service) Exists(mac string) (bool, error) {
	return s.repo.Exists(mac)
}

func (s *Service) Add(host Host) error {

	if err := s.repo.Add(host); err != nil {
		return err
	}

	return s.reloader.Reload()
}

func (s *Service) Delete(mac string) error {

	if err := s.repo.Delete(mac); err != nil {
		return err
	}

	return s.reloader.Reload()
}
