package postgres

import "fmt"

func (s *Storage) SaveAlias(url, alias string) error {
	const op = "postgres.Storage.SaveAlias"

	_, err := s.db.Exec("insert into public.short_url (url,alias) values ($1,$2)", url, alias)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "postgres.Storage.GetURL"

	var url string
	if err := s.db.QueryRow("select url from public.short_url where alias = $1", alias).
		Scan(&url); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return url, nil
}
