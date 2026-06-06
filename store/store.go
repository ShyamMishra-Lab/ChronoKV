package store

import (
	"fmt"
	"os"
	"time"
)


// a struct that stores the Log File and a map for Key-value pairs.
type Store struct {
	Value map[string]string 
	LogFile *os.File
}

// a constructor function that initializes the Store struct, creates a log file and returns a pointer to the Store struct. 
// If there is an error opening the log file, it returns nil.
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

// Set command---adds the key value pair to the map, appends the log file with the timestamp, command, key and value.
func (s *Store) Set(key, value string) {
	t := time.Now().Unix()

	s.Value[key] = value
	fmt.Fprintf(s.LogFile, "%d SET %s %s\n", t, key, value)
}

// Get command---retrieves the value of a key from the map, if the key is not found it returns an error message.
func (s *Store) Get(key string) (string, error) {
	val, ok := s.Value[key]
	// returning an error message if the key is not found
	if !ok {
		return "", fmt.Errorf("key not found: %s", key)
	}

	return val, nil
}

// Delete command---removes a key from the map and appends the log file with the timestamp, command, and key.
func (s *Store) Delete(key string) {
	t := time.Now().Unix()

	delete(s.Value, key)
	fmt.Fprintf(s.LogFile, "%d DELETE %s\n", t, key)
}

func (s *Store) GetAt(timestamp int64) string {
	// GetAt--- a function that retrieves the value of a key at a specific timestamp.
	return "GetAt function is not implemented yet"
}