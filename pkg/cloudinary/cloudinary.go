package cloudinary

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadPoster(file multipart.File, header *multipart.FileHeader) (string, error) {
	// Initialize cloudinary using environment variables
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		return "", err
	}

	// Upload file
	uploadResult, err := cld.Upload.Upload(
		context.Background(),
		file,
		uploader.UploadParams{
			Folder:         "movie-posters",
			AllowedFormats: []string{"jpg", "jpeg", "png"},
			ResourceType:   "image",
		},
	)
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
