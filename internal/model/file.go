package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type FileUploadInput struct {
	File *ghttp.UploadFile
}

type FileUploadOutput struct {
	Id   int
	Name string
	URL  string
}

type FileGetInput struct {
	Id int `json:"id"`
}

type FileGetOutput struct {
	URL        string      `json:"url"`
	Name       string      `json:"name"`
	HashString string      `json:"hash"`
	UploaderId int         `json:"uploader_id" orm:"uploader_id"`
	CreatedAt  *gtime.Time `json:"created_at" orm:"created_at"`
}