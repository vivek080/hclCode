package services

import (
	"bufio"
	"encoding/csv"
	"fmt"

	"os"
	"strconv"
	"strings"
	"time"

	con "github.com/vivek080/hclCode/gosrc/constants"
	ml "github.com/vivek080/hclCode/gosrc/model"
)

func ReadCSVFile(filepath string) ([]ml.ComputerObject, error) {
	compObject := make([]ml.ComputerObject, 0)

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error in opeing file : ", err)
		return nil, err
	}
	defer file.Close()

	tStart := time.Now()

	reader := csv.NewReader(bufio.NewReader(file))
	for {
		record, err := reader.Read()
		if err != nil {
			break
		} else {
			if record[0] != "" && record[1] != "" && record[2] != "" {
				computerID, _ := strconv.Atoi(record[0])
				userID, _ := strconv.Atoi(record[1])
				appID, _ := strconv.Atoi(record[2])
				compObject = append(compObject, ml.ComputerObject{
					ComputerID:    computerID,
					UserID:        userID,
					ApplicationID: appID,
					ComputerType:  strings.ToUpper(record[3]),
					Comment:       record[4],
				})
			}
		}
	}

	fmt.Println("Time taken to process CSV data :", time.Since(tStart).Seconds())
	return compObject, nil
}

func CalculateMinimumCopy(compObject []ml.ComputerObject, applicationID int) int {
	compTypeCount := make(map[int]map[int]string, 0)

	for _, comp := range compObject {
		if comp.ApplicationID == applicationID {
			if _, ok := compTypeCount[comp.UserID]; ok {
				compTypeCount[comp.UserID][comp.ComputerID] = comp.ComputerType
			} else {
				compTypeCount[comp.UserID] = map[int]string{comp.ComputerID: comp.ComputerType}
			}
		}
	}

	copiesCount := 0
	for _, rec := range compTypeCount {
		isLaptop := false
		for _, i := range rec {
			if i == con.ComputerType {
				isLaptop = true
			}
		}
		if isLaptop {
			copiesCount++
		} else {
			copiesCount += len(rec)
		}
	}

	return copiesCount
}
