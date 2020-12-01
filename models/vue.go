package models

import "github.com/jinzhu/gorm"

type VueData struct {
	gorm.Model
	Timestamp int `json:"timestamp"`
	Author string `json:"author"`
	Reviewer string `json:"reviewer"`
	Title string `json:"title"`
	Content_short string `json:"content_short"`
	Content string `json:"content"`
	Forecast float64 `json:"forecast"`
	Importance int `json:"importance"`
	Type string `json:"type"`
	Status string `json:"status"`
	Display_time string `json:"display_time"`
	Comment_disabled bool `json:"comment_disabled"`
	Pageviews int `json:"pageviews"`
	Image_uri string `json:"image_uri"`
	Platforms []Platforms `json:"platforms"`
}

type Platforms struct {
	gorm.Model
	Name string
	VueDataID uint
}