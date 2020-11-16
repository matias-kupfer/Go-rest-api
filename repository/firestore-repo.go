package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/matiascfgm/restAPI/entity"
	"google.golang.org/api/option"
	"log"
)

type repo struct{}

func NewFirestoreMovieRepository() MovieRepository {
	return &repo{}
}

const (
	projectId      string = "golang-b4034"
	collectionName string = "movies"
)

func (*repo) Save(movie *entity.Movie) (*entity.Movie, error) {
	opt := option.WithCredentialsFile("/Users/matias/Development/Go/serviceAccountKey.json")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Fatal("Failed above to create firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":          movie.ID,
		"Title":       movie.Title,
		"Description": movie.Description,
		"Year":        movie.Year,
	})
	if err != nil {
		log.Fatalf("Failed adding new movie: %v", err)
		return nil, err
	}
	return movie, nil
}

func (*repo) FindAll() ([]entity.Movie, error) {
	/*ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
		return nil, err
	}*/
	opt := option.WithCredentialsFile("/Users/matias/Development/Go/serviceAccountKey.json")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Fatal("Failed below to create firestore client: %v", err)
		return nil, err
	}

	defer client.Close()
	var movies []entity.Movie
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iterator.Next()
		if err != nil {
			break
			log.Fatalf("Failed to iterate movies: %v", err)
			return nil, err
		}
		movie := entity.Movie{
			ID:          doc.Data()["ID"].(string),
			Title:       doc.Data()["Title"].(string),
			Description: doc.Data()["Description"].(string),
			Year:        doc.Data()["Year"].(string),
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
