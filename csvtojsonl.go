package main

import (
    "encoding/csv"
    "encoding/json"
    "flag"
    "fmt"
    "os"
    "strconv"
)

// House represents the structure of our input data
type House struct {
    Value     float64 `json:"value"`
    Income    float64 `json:"income"`
    Age       int     `json:"age"`
    Rooms     int     `json:"rooms"`
    Bedrooms  int     `json:"bedrooms"`
    Population int    `json:"population"`
    Households int    `json:"households"`
}

func main() {
    inputFilePath := flag.String("input", "housesInput.csv", "Path to input CSV file")
    outputFilePath := flag.String("output", "housesOutput.jsonl", "Path to output JSON Lines file")
    flag.Parse()

    file, err := os.Open(*inputFilePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)
    reader.Comma = ',' // Set the delimiter to comma
    reader.Read()      // Read and discard the header

    outputFile, err := os.Create(*outputFilePath)
    if err != nil {
        fmt.Println("Error creating output file:", err)
        return
    }
    defer outputFile.Close()

    for {
        record, err := reader.Read()
        if err != nil {
            if err == csv.ErrFieldCount {
                fmt.Println("Warning: skipping bad CSV row:", record)
                continue
            }
            break // EOF or an actual read error
        }

        house := House{}
        house.Value, _ = strconv.ParseFloat(record[0], 64)
        house.Income, _ = strconv.ParseFloat(record[1], 64)
        house.Age, _ = strconv.Atoi(record[2])
        house.Rooms, _ = strconv.Atoi(record[3])
        house.Bedrooms, _ = strconv.Atoi(record[4])
        house.Population, _ = strconv.Atoi(record[5])
        house.Households, _ = strconv.Atoi(record[6])

        jsonLine, err := json.Marshal(house)
        if err != nil {
            fmt.Println("Error marshalling JSON:", err)
            continue
        }
        outputFile.WriteString(string(jsonLine) + "\n")
    }
}
