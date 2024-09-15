package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	UploadDate string `json:"upload_date"`
	FileType  string `json:"file_type"`
	FilePath  string `json:"file_path"`
}



