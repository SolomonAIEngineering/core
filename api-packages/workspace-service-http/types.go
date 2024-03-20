package workspaceservicehttp

import (
	"mime"
	"path/filepath"
)

type (
	ArrayResponse []string
	MapResponse   map[string]string
)

// Define a struct to hold the additional fields from the request body
type FileUploadRequest struct {
	WorkspaceID uint64 `json:"workspaceId"`
	FolderID    uint64 `json:"folderId"`
	Filename    string `json:"filename"` // This can be used instead of the filename from the header
	UserID      string `json:"userId"`
}

type FileUploadResponse struct {
	FileURL string `json:"fileUrl"`
	FileId  uint64 `json:"file_id"`
}

func getFileType(filename string) string {
	ext := filepath.Ext(filename)
	mimeType := mime.TypeByExtension(ext)
	return mimeType
}
