package entity

type URL struct {
	ShortURL    string `gorm:"primaryKey"`
	OriginalURL string
}
