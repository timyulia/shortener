package service

// GetShortURL ...
func (s *Service) GetShortURL(long string) (string, error) {
	short := generateShortURL(long)
	existed, err := s.repo.GetLongURL(short)
	for {
		if err != nil {
			return "", err
		}
		if existed == "" {
			err := s.repo.AddLinksPair(short, long)
			return short, err
		}
		if existed == long {
			return short, nil
		}
		short = generateShortURL(short)
		existed, err = s.repo.GetLongURL(short)
	}
}

// GetLongURL ...
func (s *Service) GetLongURL(short string) (string, error) {
	return s.repo.GetLongURL(short)
}

//todo check if it is link
