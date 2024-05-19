package main

import (
	"fmt"
	"sharemail/internal/config"
	"sharemail/internal/url"
	"strconv"
)

func main() {
	config.LoadEnv()

	length, err := strconv.Atoi(config.AppConfig["UNIQUE_STRING_LENGTH"])
	if err != nil {
		fmt.Printf("Error converting string length to int: %s\n", err)
		return
	}

	uniqueStrings := url.GenerateUniqueStrings(length)
	fmt.Printf("Generated %d unique strings\n", len(uniqueStrings))

	err = url.SyncUrls(&uniqueStrings)
	if err != nil {
		fmt.Printf("Error syncing urls: %s\n", err)
		return
	}

	fmt.Println("Urls synced successfully")

}
