package createdatacsv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV() {

	file, err := os.Open("NoisyPolynomialData.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	fmt.Println(records)
}
