package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Fixlet struct {
	SiteID                int
	FxiletID              string
	Name                  string
	Criticality           string
	RelevantComputerCount int
}

var fixlets []Fixlet
var filePath = "C:/Users/anubhav.arora/Desktop/golang/Day-3-fixlets.csv"

func main() {
	// Load the CSV data
	err := loadCSV()
	if err != nil {
		fmt.Printf("Error loading CSV file: %v\n", err)
		return
	}

	// Menu-driven program
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. List entries")
		fmt.Println("2. Find by FxiletID")
		fmt.Println("3. Sort entries using ComputerCount")
		fmt.Println("4. Add entry")
		fmt.Println("5. Delete Entry")
		fmt.Println("6. Exit the program")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			listEntries()
		case 2:
			queryEntry()
		case 3:
			sortEntries()
		case 4:
			addEntry()
		case 5:
			deleteEntry()
		case 6:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func loadCSV() error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Parse CSV data into Fixlet structs
	for i, record := range records {
		// Skip the header row
		if i == 0 {
			continue
		}

		// Parse SiteID
		siteID, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Printf("Invalid SiteID at row %d: %v\n", i+1, err)
			continue
		}

		// Parse RelevantComputerCount
		relevantCount, err := strconv.Atoi(record[4])
		if err != nil {
			fmt.Printf("Invalid RelevantComputerCount at row %d: %v\n", i+1, err)
			continue
		}

		// Create a Fixlet struct
		fixlet := Fixlet{
			SiteID:                siteID,
			FxiletID:              record[1],
			Name:                  record[2],
			Criticality:           record[3],
			RelevantComputerCount: relevantCount,
		}
		fixlets = append(fixlets, fixlet)
	}
	return nil
}

func listEntries() {
	fmt.Println("\nListing all entries:")
	for _, fixlet := range fixlets {
		fmt.Printf("SiteID: %d, FxiletID: %s, Name: %s, Criticality: %s, RelevantComputerCount: %d\n",
			fixlet.SiteID, fixlet.FxiletID, fixlet.Name, fixlet.Criticality, fixlet.RelevantComputerCount)
	}
}

func queryEntry() {
	fmt.Print("Enter FxiletID to query: ")
	var fxiletID string
	fmt.Scan(&fxiletID)

	for _, fixlet := range fixlets {
		if fixlet.FxiletID == fxiletID {
			fmt.Printf("Found Entry - SiteID: %d, FxiletID: %s, Name: %s, Criticality: %s, RelevantComputerCount: %d\n",
				fixlet.SiteID, fixlet.FxiletID, fixlet.Name, fixlet.Criticality, fixlet.RelevantComputerCount)
			return
		}
	}
	fmt.Println("Entry not found.")
}

func sortEntries() {
	sort.Slice(fixlets, func(i, j int) bool {
		return fixlets[i].RelevantComputerCount > fixlets[j].RelevantComputerCount
	})
	fmt.Println("Entries sorted by RelevantComputerCount (descending).")
	listEntries()
}

func addEntry() {
	var fixlet Fixlet
	fmt.Print("Enter SiteID: ")
	fmt.Scan(&fixlet.SiteID)
	fmt.Print("Enter FxiletID: ")
	fmt.Scan(&fixlet.FxiletID)
	fmt.Print("Enter Name: ")
	fmt.Scan(&fixlet.Name)
	fmt.Print("Enter Criticality: ")
	fmt.Scan(&fixlet.Criticality)
	fmt.Print("Enter RelevantComputerCount: ")
	fmt.Scan(&fixlet.RelevantComputerCount)

	fixlets = append(fixlets, fixlet)
	fmt.Println("New entry added.")
}

func deleteEntry() {
	fmt.Print("Enter FxiletID to delete: ")
	var fxiletID string
	fmt.Scan(&fxiletID)

	for i, fixlet := range fixlets {
		if fixlet.FxiletID == fxiletID {
			fixlets = append(fixlets[:i], fixlets[i+1:]...)
			fmt.Println("Entry deleted.")
			return
		}
	}
	fmt.Println("Entry not found.")
}

func saveCSV() error {
	// Open file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"SiteID", "FxiletID", "Name", "Criticality", "RelevantComputerCount"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write entries
	for _, fixlet := range fixlets {
		record := []string{
			strconv.Itoa(fixlet.SiteID),
			fixlet.FxiletID,
			fixlet.Name,
			fixlet.Criticality,
			strconv.Itoa(fixlet.RelevantComputerCount),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
