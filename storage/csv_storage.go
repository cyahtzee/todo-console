package storage

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"todo-console/constants"
)

type CSVStorage struct {
	Items    *[]Item
	filePath string
}

func NewCSVStorage(filePath string) *CSVStorage {
	items := loadItems(filePath)

	return &CSVStorage{filePath: filePath, Items: items}
}

// Create(todo *Item)
// Find(id int) *Item
// Remove(id int)
// FindAll() *[]Item

func loadItems(filePath string) *[]Item {
	csvFile, err := os.Create(filePath)

	if err != nil {
		log.Fatalf("Failed to create CSV file: %v", err)
	}

	defer csvFile.Close()

	rows, err := csv.NewReader(csvFile).ReadAll()

	items := make([]Item, len(rows))

	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	for i, item := range rows {
		id, _ := strconv.Atoi(item[0])
		items[i] = &Todo{
			ID:          id,
			Title:       item[1],
			Description: item[2],
			Completed:   constants.Status(item[3]),
		}
	}

	return &items
}

func (s *CSVStorage) Create(item Item) {
	id := len(*s.Items) + 1
	item.SetID(id)
	*s.Items = append(*s.Items, item)
	s.Commit()
}

func (s *CSVStorage) Find(id int) *Item {
	for _, item := range *s.Items {
		if item.GetID() == id {
			return &item
		}
	}

	return nil
}

func (s *CSVStorage) Remove(id int) {
	for i, item := range *s.Items {
		if item.GetID() == id {
			*s.Items = append((*s.Items)[:i], (*s.Items)[i+1:]...)
		}
	}

	s.Commit()
}

func (s *CSVStorage) FindAll() *[]Item {
	return s.Items
}

func (s *CSVStorage) Commit() {
	csvFile, err := os.Create(s.filePath)

	if err != nil {
		log.Fatalf("Failed to create CSV file: %v", err)
	}

	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	csvFile.Truncate(0)
	csvFile.Seek(0, io.SeekStart)
	writer.Write([]string{"ID", "Title", "Description", "Status"})

	for _, item := range *s.Items {
		writer.Write([]string{strconv.Itoa(item.GetID()), item.GetTitle(), item.GetDescription(), item.GetStatus()})
	}

	writer.Flush()
}
