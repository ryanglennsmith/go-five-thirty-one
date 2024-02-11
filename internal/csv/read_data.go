package csv

import (
	"encoding/csv"
	"os"
	"strconv"
)

//LIFT,WEEKSET,WEIGHT,REPS,DATE,COMMENTS,ONERM

type LiftData struct {
	Lift string
	WeekSet string
	Weight float64
	Reps int64
	Date string
	Comments string
	OneRM int64
}

func ReadData(filePath string) ([]LiftData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var liftDataList []LiftData
	for _, record := range records[1:]{
		weightStr := record[2]
		repsStr := record[3]
		oneRMStr := record[6]

		if weightStr == "" {
			weightStr = "0"
		}
		if repsStr == "" {
			repsStr = "0"
		}
		if oneRMStr == "" {
			oneRMStr = "0"
		}
		
		weight, err := strconv.ParseFloat(weightStr, 64)
		if err != nil {
			return nil, err
		}

		reps, err := strconv.ParseInt(repsStr, 10, 64)
		if err != nil {
			return nil, err
		}

		oneRM, err := strconv.ParseInt(oneRMStr, 10, 64)
		if err != nil {
			return nil, err
		}
		
		liftDataList = append(liftDataList, LiftData{
			Lift: record[0],
			WeekSet: record[1],
			Weight: weight,
			Reps: reps,
			Date: record[4],
			Comments: record[5],
			OneRM: oneRM,
		})
	}

	return liftDataList, nil
}