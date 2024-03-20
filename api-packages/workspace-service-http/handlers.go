package workspaceservicehttp

import "net/http"

// @tags workspace-service-rest
// @summary Uploads a file to the server
// @description This endpoint allows for the uploading of a file to the server. Upon successful upload,
// the server saves the file and returns a URL to access the saved file along with its metadata.
// This can be used to store and retrieve files within a specific workspace and folder.
// @accept multipart/form-data
// @produce application/json
// @param workspaceId formData string true "Identifier of the workspace where the file will be uploaded"
// @param folderId formData string true "Identifier of the folder within the workspace where the file will be stored"
// @param filename formData string false "The name of the file to be saved (optional)"
// @param userId formData string true "Identifier of the user uploading the file"
// @param attachment formData file true "The file to be uploaded"
// @success 200 {object} workspaceservicehttp.FileUploadResponse "On successful upload, returns the URL to the saved file and its metadata, including file name, size, type, and the upload timestamp."
// @router /workspace-microservice/rest-api/v1/file/upload [post]
func (s *RestService) UploadFile(w http.ResponseWriter, r *http.Request) {}

// ListenAndServe starts the HTTP server and returns the server instances and the ports they are listening on.
func (s *RestService) ListenAndServe() (*http.Server, *http.Server, *int32, *int32) {
	return nil, nil, nil, nil
}
