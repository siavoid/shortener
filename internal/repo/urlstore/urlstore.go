package urlstore

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
)

// URLStore представляет собой структуру для хранения полных и сокращённых URL
type URLStore struct {
	fileStorePath string
	mu            sync.Mutex
	//	- shortURL - сокращенная ссылка
	//	- url      - исходный url
	storeLongShort map[string]string // Хранит пары  url : shortURL
	storeShortLong map[string]string // Хранит пары  shortURL : url
}

// NewURLStore создает новый экземпляр URLStore
func NewURLStore(fileStorePath string) (*URLStore, error) {
	store := &URLStore{
		fileStorePath:  fileStorePath,
		storeLongShort: make(map[string]string),
		storeShortLong: make(map[string]string),
	}
	err := store.init()
	return store, err
}

// init инициализирует структуру из файла
func (u *URLStore) init() error {
	file, err := os.OpenFile(u.fileStorePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("open|create file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var data map[string]string
		err := json.Unmarshal(scanner.Bytes(), &data)
		if err != nil {
			return fmt.Errorf("unmarshal data: %w", err)
		}
		u.storeShortLong[data["short_url"]] = data["original_url"]
		u.storeLongShort[data["original_url"]] = data["short_url"]
	}
	return nil
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
func (u *URLStore) Put(url, shortURL string) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.storeShortLong[shortURL] = url
	u.storeLongShort[url] = shortURL

	data := map[string]string{
		"uuid":         uuid.NewString(), // генерация UUID
		"short_url":    shortURL,
		"original_url": url,
	}

	file, err := os.OpenFile(u.fileStorePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal data: %w", err)
	}

	jsonData = append(jsonData, []byte("\n")...)
	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("write data: %w", err)
	}
	return nil
}
