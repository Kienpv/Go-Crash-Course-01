package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/Go-Crash-Course-01/entity"
	"google.golang.org/api/iterator"
)

type PostRepo interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct {}

func NewPostRepo() PostRepo {
	return &repo{}
}

const (
	projectID 		string = "First-demo"
	collectionName 	string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID) 
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{} {
		"ID": post.Id,
		"Title": post.Title,
		"Text": post.Text,
	})

	if err != nil {
		log.Fatalf("Failed to adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID) 
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	
	itor := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := itor.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to interate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post {
			Id:	doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text: doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
