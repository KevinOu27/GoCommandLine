Assignment 3 for MSDS 431 to create a command line application using Go that will read csv text file and create JSON lines file for the same data as output. 

## Installation

Clone the repository and run the application found in the repository (ensure terminal is in the directory github repository is downloaded to)
Use ./csvtojsonl -input=path/to/input.csv (this is your input csv file) -output=path/to/output.jsonl (this is the name of the output file you want to convert to)

Example:
./csvtojsonl -input=housesInput.csv -output-housesOutput.jsonl

## Usage
This will efficiently parse CSV files and each record from the CSV file will convert into a single JSON object outputting as one line in a .jsonl file. 

##Build
Created a new Go file in Terminal using "touch csvtojsonl.go"
Write code found in the csvtojsonl.go file
Build the application using "go build csvtojsonl.go"
Run the application using "./csvtojsonl -input=path/to/input.csv"
in this case, ./csvtojsonl -input=housesInput.csv -output-housesOutput.jsonl
