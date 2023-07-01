package createdatacsv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func FileExists(name string) bool {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
	}

	if file != nil {
		return true
	}

	return false
}

func AddToCSVFile(name string, data []string) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}
