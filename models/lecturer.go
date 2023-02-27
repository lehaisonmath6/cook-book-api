package models

type Lecture struct {
	ID          int64    `json:"id,omitempty" bson:"id"`
	Title       string   `json:"title,omitempty" bson:"title"`
	Intro       string   `json:"intro,omitempty" bson:"intro"`
	VideoUrl    string   `json:"video_url,omitempty" bson:"video_url"`
	DocumentUrl string   `json:"document_url,omitempty" bson:"document_url"`
	Author      Author   `json:"author,omitempty" bson:"author"`
	ThumbUrl    string   `json:"thumb_url,omitempty" bson:"image_url"`
	Category    Category `json:"category,omitempty" bson:"category"`
	CreateTime  int64    `json:"create_time,omitempty" bson:"create_time"`
	TotalViews  int64
}

type Category struct {
	Name          string `json:"name,omitempty" bson:"name"`
	Path          string `json:"path,omitempty" bson:"path"`
	TotalLectures int    `json:"total_lectures,omitempty" bson:"total_lectures"`
}

type Author struct {
	ID     int64
	Name   string
	Avatar string
}
