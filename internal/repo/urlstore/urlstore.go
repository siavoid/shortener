package urlstore

import "sync"

// URLStore представляет собой структуру для хранения полных и сокращённых URL
type URLStore struct {
	mu sync.Mutex
	//	- shortURL - сокращенная ссылка
	//	- url      - исходный url
	storeLongShort map[string]string // Хранит пары  url : shortURL
	storeShortLong map[string]string // Хранит пары  shortURL : url
}

// NewURLStore создает новый экземпляр URLStore
func NewURLStore() *URLStore {
	return &URLStore{
		storeLongShort: make(map[string]string),
		storeShortLong: make(map[string]string),
	}
}

// Get получает полное значение URL по сокращённому ключу
func (u *URLStore) GetLongURL(shortURL string) (string, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()
	url, exists := u.storeShortLong[shortURL]
	return url, exists
}

// Get получает полное значение URL по сокращённому ключу
func (u *URLStore) GetShortURL(url string) (string, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()
	shortURL, exists := u.storeLongShort[url]
	return shortURL, exists
}

// Put добавляет новую пару ключ-значение в хранилище
func (u *URLStore) Put(url, shortURL string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.storeShortLong[shortURL] = url
	u.storeLongShort[url] = shortURL
}
