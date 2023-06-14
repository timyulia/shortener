package service

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
		short = generateShortURL(short) //в случае коллизии генерируется новая короткая ссылка по той короткой ссылке, что была сгенерирована ранее
		existed, err = s.repo.GetLongURL(short)
	}
}

func (s *Service) GetLongURL(short string) (string, error) {
	return s.repo.GetLongURL(short)
}
