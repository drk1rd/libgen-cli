package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/drk1rd/libgen-cli/libgenapi"
)

func main() {
	searchFlag := flag.String("search", "", "Search query for books (required)")
	queryTypeFlag := flag.String("type", "title", "Search type: 'title', 'author', 'publisher', etc.")
	resultsFlag := flag.Int("results", 5, "Number of results to fetch")
	downloadFlag := flag.Int("download", -1, "Download the book by selecting a number from the search results")

	flag.Parse()

	if *searchFlag == "" {
		fmt.Println("Error: You must provide a search query using -search flag")
		os.Exit(1)
	}

	query := libgenapi.NewQuery(*queryTypeFlag, *searchFlag, *resultsFlag)

	if err := query.Search(); err != nil {
		log.Fatal("Error searching books:", err)
	}

	showSearchResults(query.Results)

	if *downloadFlag >= 0 && *downloadFlag < len(query.Results) {
		selectedBook := query.Results[*downloadFlag]
		err := selectedBook.AddSecondDownloadLink()
		if err != nil {
			log.Fatalf("Error adding second download link: %v", err)
		}
		downloadBook(selectedBook)
	} else {
		fmt.Println("Invalid download option. Exiting.")
	}
}

func showSearchResults(books []libgenapi.Book) {
	if len(books) == 0 {
		fmt.Println("No books found.")
		return
	}
	for i, book := range books {
		fmt.Printf("%d) Title: %s\n", i, book.Title)
		fmt.Printf("   Author: %s | Year: %s | Pages: %s\n", book.Author, book.Year, book.Pages)
		fmt.Printf("   Download: %s\n", book.DownloadLink)
		if book.AlternativeDownloadLink != "" {
			fmt.Printf("   Alternative Download: %s\n", book.AlternativeDownloadLink)
		}
	}
}

func downloadBook(book libgenapi.Book) {
	fmt.Printf("\nDownloading %s...\n", book.Title)
	fmt.Printf("Download link: %s\n", book.DownloadLink)
	fmt.Println("Download complete!")
}
