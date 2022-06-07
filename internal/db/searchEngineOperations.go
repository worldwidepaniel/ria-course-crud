package db

import (
	"fmt"

	"github.com/meilisearch/meilisearch-go"
	"github.com/worldwidepaniel/ria-course-crud/internal/config"
)

func connectToSearchEngine() *meilisearch.Client {
	return meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://ria-meilisearch:7700",
		APIKey: config.AppConfig.SearchEngine.Key,
	})
}

func AddToSearchEngine(Notes []Note) error {
	client := connectToSearchEngine()
	_, err := client.Index("notes").AddDocuments(Notes)
	if err != nil {
		fmt.Errorf("error while adding documents to search engine")
	}
	settings := meilisearch.Settings{
		FilterableAttributes: []string{"UID"},
	}
	client.Index("notes").UpdateSettings(&settings)
	return nil

}
