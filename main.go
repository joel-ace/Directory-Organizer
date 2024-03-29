package main

import (
	"desktop-cleaner/organizer"
	"fmt"
)

func main() {
	var targetDirectory string
	var userInput string
	fileCategories := organizer.GetFileCategories()
	desktopDirectory := organizer.GetDesktopDirectory()
	downloadDirectory := organizer.GetDownloadDirectory()

	fmt.Println("===================================================================================")
	fmt.Println("                           WELCOME TO DIRECTORY ORGANIZER                          ")
	fmt.Println("===================================================================================")
	fmt.Println("CHOOSE DIRECTORY TO BE CLEANED \n Enter 1 for Desktop \n Enter 2 for Downloads")

	fmt.Scan(&userInput)

	switch userInput {
	case "1":
		targetDirectory = desktopDirectory
	case "2":
		targetDirectory = downloadDirectory
	default:
		targetDirectory = desktopDirectory
	}

	arrangedFileDirectory := organizer.GetArrangedFileDirectory(targetDirectory)

	organizer.CreateDirectories(fileCategories, arrangedFileDirectory)

	files := organizer.ReadDirectory(targetDirectory)

	organizer.OrganizeFiles(files, fileCategories, targetDirectory, arrangedFileDirectory)
}
