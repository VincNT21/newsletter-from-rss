package main

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// Load env variable if any
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, using default filepaths")
	}

	// Start logging system
	setupLogging()

	infoLogger.Println("=====================")
	infoLogger.Println("Application started")

	// Get global config
	globalConfig, err := getGlobalConfig()
	if err != nil {
		logError(err.Error(), "Read JSON file")
	} else {
		debugLogger.Println("config.json reading ok")
	}

	// Get Newsletter data
	newsletterData, err := getNewsletterData(globalConfig)
	if err != nil {
		logError(err.Error(), "Fetch data from RSS")
	} else {
		debugLogger.Println("Newsletter Data fetching ok")
	}

	// Create Newsletter HTML file
	filename, err := generateNewsletter(newsletterData, globalConfig.DateFormat)
	if err != nil {
		errMessage := fmt.Sprintf("Couldn't create HTML file: %v", err)
		logError(errMessage, "HTML file creation")
	} else {
		debugLogger.Println("Newsletter HTML creation ok")
	}

	// Tell user about errors
	if hadError {
		if len(errorComponents) == 1 {
			fmt.Printf(" !! There has been an error with: '%s'. You could check for details in 'app.log' file. !!\n", errorComponents[0])
		} else {
			fmt.Printf(" !! There has been several errors with: '%s'. You could check for details in 'app.log' file !!\n", strings.Join(errorComponents, ", "))
		}
	} else {
		fmt.Printf(" vv Everything went smoothly ! Your newsletter is in file '%s' vv\n", filename)
	}

	// Manage exit
	fmt.Println("Press Enter to exit...")
	infoLogger.Printf("Waiting for user to press enter to exit")
	fmt.Scanln()
	if hadError {
		infoLogger.Printf("Application exit with errors in: %s", strings.Join(errorComponents, ", "))
	} else {
		infoLogger.Printf("Application exit without errors")
	}
	infoLogger.Println("=====================")
}
