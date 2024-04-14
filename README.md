Assignment 3 for MSDS 431 to create a command line application using Go that will read csv text file and create JSON lines file for the same data as output. 

## Installation

Clone the repository and run the application found in the repository
Use ./csvtojsonl -input=path/to/input.csv (this is your input csv file) -output=path/to/output.jsonl (this is the name of the output file you want to convert to)

Example:
./csvtojsonl -input=housesInput.csv -output-housesOutput.jsonl

## Usage
This will efficiently parse CSV files and each record from the CSV file will convert into a single JSON object outputting as one line in a .jsonl file. 
