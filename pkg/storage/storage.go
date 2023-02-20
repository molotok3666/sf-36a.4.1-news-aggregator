package storage

// News - новость.
type News struct {
	GUID    string
	Title   string
	Content string
	PubTime int64
	Link    string
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	News(int) ([]News, error)
}

// Конструктор новости
func NewNews(guid string, title string, content string, pubTime int64, link string) News {
	return News{
		guid,
		title,
		content,
		pubTime,
		link,
	}
}
