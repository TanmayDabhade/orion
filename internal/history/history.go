package history

import (
	"database/sql"
	"fmt"
	"time"

	"orion/storage"
)

type Store struct {
	db *sql.DB
}

type Usage struct {
	Key      string
	Count    int
	LastUnix int64
}

func Open(path string) (*Store, error) {
	db, err := storage.Open(path)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}

func (s *Store) Record(input string, key string, success bool) error {
	if s == nil || s.db == nil {
		return fmt.Errorf("history store unavailable")
	}
	ts := time.Now().Unix()
	successInt := 0
	if success {
		successInt = 1
	}

	if _, err := s.db.Exec(`INSERT INTO history (input, success, ts) VALUES (?, ?, ?)`, input, successInt, ts); err != nil {
		return err
	}

	if _, err := s.db.Exec(
		`INSERT INTO usage (key, count, last_ts) VALUES (?, 1, ?)
		 ON CONFLICT(key) DO UPDATE SET count = count + 1, last_ts = excluded.last_ts`,
		key, ts,
	); err != nil {
		return err
	}

	return nil
}

func (s *Store) Usage(keys []string) (map[string]Usage, error) {
	result := make(map[string]Usage)
	if s == nil || s.db == nil || len(keys) == 0 {
		return result, nil
	}

	placeholders := make([]string, 0, len(keys))
	args := make([]interface{}, 0, len(keys))
	for _, key := range keys {
		placeholders = append(placeholders, "?")
		args = append(args, key)
	}

	query := fmt.Sprintf(`SELECT key, count, last_ts FROM usage WHERE key IN (%s)`, join(placeholders, ","))
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var usage Usage
		if err := rows.Scan(&usage.Key, &usage.Count, &usage.LastUnix); err != nil {
			return nil, err
		}
		result[usage.Key] = usage
	}
	return result, rows.Err()
}

func join(items []string, sep string) string {
	if len(items) == 0 {
		return ""
	}
	out := items[0]
	for i := 1; i < len(items); i++ {
		out += sep + items[i]
	}
	return out
}
