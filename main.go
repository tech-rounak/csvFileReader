package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	f, err := os.Open("netflix_titles.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 4, '.', tabwriter.Debug)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line

		for _, row := range rec {
			fmt.Fprintf(w, "%s\t", row)
		}
		fmt.Fprintf(w, "\n")
		// fmt.Printf("%+v\n", rec)
		// fmt.Printf("%s\n", rec[0])
		// fmt.Printf("%d\n", len(rec))
	}
	w.Flush()
}
