package service

import "context"

func (s *Service) GetShortURL(ctx context.Context, long string) (string, error) {
	short := generateShortURL(long)
	existed, err := s.repo.GetLongURL(ctx, short)
	for {
		if err != nil {
			return "", err
		}
		if existed == "" {
			err := s.repo.AddLinksPair(ctx, short, long)
			return short, err
		}
		if existed == long {
			return short, nil
		}
		short = generateShortURL(short) //в случае коллизии генерируется новая короткая ссылка по той короткой ссылке, что была сгенерирована ранее
		existed, err = s.repo.GetLongURL(ctx, short)
	}
}

func (s *Service) GetLongURL(ctx context.Context, short string) (string, error) {
	return s.repo.GetLongURL(ctx, short)
}
