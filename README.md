# go-lang-assignment

Overview
The Fixlet Manager is a command-line application written in Go for managing entries in a CSV file. It provides functionalities to list, query, sort, add, and delete fixlet entries from a CSV file. The application is tailored to handle fixlet data with fields like SiteID, FxiletID, Name, Criticality, and RelevantComputerCount.

Features

1. List Entries: Displays all entries in the CSV file.
2. Query Entry: Searches for an entry by FxiletID.
3. Sort Entries: Sorts entries by RelevantComputerCount in descending order.
4. Add Entry: Adds a new fixlet entry to the list.
5. Delete Entry: Removes an entry by FxiletID.
6. Save Changes: Automatically persists changes to the CSV file.

Prerequisites

1. Go Programming Language: Ensure Go is installed on your system.
    Download Go from golang.org.

Menu Options

List all entries
Prints all fixlet entries in the CSV file.

Query an entry by FxiletID
Prompts for a FxiletID and searches for the corresponding entry.

Sort entries by RelevantComputerCount
Sorts entries in descending order of RelevantComputerCount and displays them.

Add a new entry
Prompts for details of a new fixlet and adds it to the list.

Delete an entry by FxiletID
Prompts for a FxiletID and deletes the corresponding entry from the list.

Exit
Terminates the application.

Notes

Changes are automatically saved to the CSV file whenever entries are added or deleted.
If the program encounters invalid data, it will skip those rows and continue processing.

Limitations

The application currently assumes the CSV file structure remains constant.
No validation is done to check for duplicate FxiletIDs during addition.
