package pkg

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryUsecase interface {
	UploadImage(ctx context.Context, file *multipart.FileHeader) (string, error)
}

type cloudinaryUsecaseImpl struct {
	cloudName string
	apiKey    string
	apiSecret string
}

func NewCloudinaryUsecase(cloudName, apiKey, apiSecret string) CloudinaryUsecase {
	return &cloudinaryUsecaseImpl{
		cloudName: cloudName,
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (u *cloudinaryUsecaseImpl) UploadImage(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {
	cld, err := cloudinary.NewFromParams(u.cloudName, u.apiKey, u.apiSecret)
	if err != nil {
		return "", fmt.Errorf("failed to init cloudinary: %w", err)
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: "go-auction",
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload to cloudinary: %w", err)
	}

	return resp.SecureURL, nil
}
