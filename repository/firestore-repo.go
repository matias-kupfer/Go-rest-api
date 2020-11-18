package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/matiascfgm/Go-rest-api/entity"
	"google.golang.org/api/iterator"
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
		"id":          movie.ID,
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

func (*repo) GetMovieById(id string) (*entity.Movie, error) {
	opt := option.WithCredentialsFile("/Users/matias/Development/Go/serviceAccountKey.json")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Fatal("Failed below to create firestore client: %v", err)
		return nil, err
	}
	defer client.Close()

	iter := client.Collection(collectionName).Where("id", "==", id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		return &entity.Movie{
			ID:          doc.Data()["id"].(string),
			Title:       doc.Data()["Title"].(string),
			Description: doc.Data()["Description"].(string),
			Year:        doc.Data()["Year"].(string),
		}, nil
	}
	return nil, err
}

func (*repo) DeleteMovie(id string) (*entity.Movie, error) {
	opt := option.WithCredentialsFile("/Users/matias/Development/Go/serviceAccountKey.json")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Fatal("Failed below to create firestore client: %v", err)
		return nil, err
	}
	defer client.Close()

	iter := client.Collection(collectionName).Where("id", "==", id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		_, e := client.Collection(collectionName).Doc(doc.Ref.ID).Delete(ctx)
		if e != nil {
			return nil, err
		}
		return &entity.Movie{}, nil
	}
	return nil, err
}

func (*repo) UpdateMovie(movie *entity.Movie) (*entity.Movie, error) {
	opt := option.WithCredentialsFile("/Users/matias/Development/Go/serviceAccountKey.json")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, opt)
	iter := client.Collection(collectionName).Where("id", "==", movie.ID).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		_, e := client.Collection(collectionName).Doc(doc.Ref.ID).Update(ctx, []firestore.Update{
			{
				Path:  "Title",
				Value: movie.Title,
			}, {
				Path:  "Description",
				Value: movie.Description,
			}, {
				Path:  "Year",
				Value: movie.Year,
			},
		})
		if e != nil {
			return nil, err
		}
		return &entity.Movie{
			ID:          doc.Data()["id"].(string),
			Title:       doc.Data()["Title"].(string),
			Description: doc.Data()["Description"].(string),
			Year:        doc.Data()["Year"].(string),
		}, nil
	}
	return nil, err
}
