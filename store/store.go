package store

import (
	"fmt"
	"os"
	"time"
)

type Store struct {
	Value map[string]string 
	LogFile *os.File
}

func NewStore() *Store {
	data := make(map[string]string)

	// err is a common variable name used in Go to represent an error, fuctions that can fail typically return an error as the second return value.
	logFile, err := os.OpenFile("store.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		// In go functions, it's common to return an error as the second return value to indicate if something went wrong. 
		// If the file fails to open, we can return nil for the Store and an error to indicate the failure.
		return nil
	}

	return &Store{
		Value:   data,
		LogFile: logFile,
	}

}

func (s *Store) Set(key, value string) {
	t := time.Now().Unix()

	s.Value[key] = value
	fmt.Fprintf(s.LogFile, "%d SET %s %s\n", t, key, value)
}

func (s *Store) Get(key string) (string, error) {
	val, ok := s.Value[key]
	// returning an error message if the key is not found
	if !ok {
		return "", fmt.Errorf("key not found: %s", key)
	}

	return val, nil
}

func (s *Store) Delete(key string) {
	t := time.Now().Unix()

	delete(s.Value, key)
	fmt.Fprintf(s.LogFile, "%d DELETE %s\n", t, key)
}

func (s *Store) GetAt(timestamp int64) string {
	// GetAt--- a function that retrieves the value of a key at a specific timestamp.
}