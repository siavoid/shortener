package pgrepo

func (p *PostgresRepo) GetLongURL(shortURL string) (string, bool) {
	return "", true
}

func (p *PostgresRepo) GetShortURL(url string) (string, bool) {
	return "", true
}

func (p *PostgresRepo) Put(url string, shortURL string) error {
	return nil
}
