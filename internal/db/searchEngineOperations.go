package db

import (
	"fmt"

	"github.com/meilisearch/meilisearch-go"
	"github.com/worldwidepaniel/ria-course-crud/internal/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func SearchPhrase(phrase string, uid primitive.ObjectID) []interface{} {
	client := connectToSearchEngine()
	searchRes, err := client.Index("notes").Search(phrase, &meilisearch.SearchRequest{
		Filter:                fmt.Sprintf("UID = %s", uid.Hex()),
		AttributesToHighlight: []string{"*"},
	})
	fmt.Println(err)
	if err != nil {
		return nil
	}
	return searchRes.Hits
}
