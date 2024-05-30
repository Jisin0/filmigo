// (c) Jisin0
// File-based caching used for scraping operations like GetMovie, GetPerson etc.
package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Cache is a struct that manages file-based caching
type Cache struct {
	directory string
	timeout   time.Duration
}

// NewCache creates a new Cache instance
func NewCache(directory string, timeout time.Duration) *Cache {
	return &Cache{directory: directory, timeout: timeout}
}

// CacheData is a struct to hold the actual data and the timestamp
type CacheData struct {
	Timestamp time.Time   `json:"timestamp"`
	Data      interface{} `json:"data"`
}

// Save saves a JSON object to a file identified by a unique ID
func (c *Cache) Save(id string, data interface{}) error {
	// Ensure the cache directory exists
	if err := os.MkdirAll(c.directory, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create cache directory: %v", err)
	}

	// Wrap data with timestamp
	cacheData := CacheData{
		Timestamp: time.Now(),
		Data:      data,
	}

	// Serialize the data to JSON
	filePath := filepath.Join(c.directory, id+".json")

	jsonData, err := json.Marshal(cacheData)
	if err != nil {
		return fmt.Errorf("failed to serialize data: %v", err)
	}

	// Write the JSON data to a file
	if err := os.WriteFile(filePath, jsonData, 0o644); err != nil {
		return fmt.Errorf("failed to write data to file: %v", err)
	}

	return nil
}

// Load loads a JSON object from a file identified by a unique ID
func (c *Cache) Load(id string, data interface{}) error {
	// Read the JSON data from the file
	filePath := filepath.Join(c.directory, id+".json")

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read data from file: %v", err)
	}

	// Deserialize the JSON data
	var cacheData CacheData

	if err = json.Unmarshal(jsonData, &cacheData); err != nil {
		return fmt.Errorf("failed to deserialize data: %v", err)
	}

	// Check if the data is within the allowed timeout
	if time.Since(cacheData.Timestamp) > c.timeout {
		return fmt.Errorf("cached data has expired")
	}

	// Copy the actual data
	dataBytes, err := json.Marshal(cacheData.Data)
	if err != nil {
		return fmt.Errorf("failed to re-marshal inner data: %v", err)
	}

	if err := json.Unmarshal(dataBytes, data); err != nil {
		return fmt.Errorf("failed to unmarshal inner data: %v", err)
	}

	return nil
}
