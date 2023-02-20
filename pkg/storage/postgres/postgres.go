package postgres

import (
	"context"
	"news-aggregator/pkg/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DbStorage struct {
	db *pgxpool.Pool
}

// Констурктор БД
func New(connstr string) (*DbStorage, error) {
	db, err := pgxpool.Connect(context.Background(), connstr)
	if err != nil {
		return nil, err
	}
	st := DbStorage{
		db: db,
	}
	return &st, nil
}

// Получется новости из БД
func (s *DbStorage) News(count int) ([]storage.News, error) {
	rows, err := s.db.Query(context.Background(),
		`
			SELECT *
			FROM news
			ORDER BY pub_time
			LIMIT $1
		`,
		count,
	)
	if err != nil {
		return nil, err
	}

	var news []storage.News

	for rows.Next() {
		var n storage.News
		err = rows.Scan(
			&n.GUID,
			&n.Title,
			&n.Content,
			&n.PubTime,
			&n.Link,
		)

		if err != nil {
			return nil, err
		}
		news = append(news, n)
	}

	return news, nil
}

// Добавляет новость в БД
func (s *DbStorage) AddNews(n storage.News) error {
	_, err := s.db.Exec(context.Background(), `
	INSERT INTO news (guid, title, content, pub_time, link)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT
	DO NOTHING
	RETURNING guid;
`,
		n.GUID,
		n.Title,
		n.Content,
		n.PubTime,
		n.Link,
	)
	if err != nil {
		return err
	}

	return nil
}
