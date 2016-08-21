package models

import "gopkg.in/mgo.v2/bson"

type Admin struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Title        string
	Content      string
	Summary      string
	ViewCount    string `bson:"view_count"`
	CommentCount string `bson:"comment_count"`
	//Classify         string
	//Tags         string
	Status       int
}