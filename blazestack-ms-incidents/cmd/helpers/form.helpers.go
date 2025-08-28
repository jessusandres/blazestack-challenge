package helpers

import (
	"encoding/base64"
	"fmt"
	"io"

	"blazestack.com/ms-incidents/cmd/apperrors"
	"github.com/gin-gonic/gin"
)

func ExtractImageForm(key string, c *gin.Context) (string, error) {
	fmt.Println("Extracting image from form:")

	imageEncoded := ""
	imageFile, err := c.FormFile(key)

	if err != nil {
		fmt.Printf("err: %+v", err)

		return imageEncoded, nil
	}

	file, err := imageFile.Open()

	defer file.Close()

	if err != nil {
		return imageEncoded, apperrors.NewBadRequestError("invalid image file")
	}

	contentType := imageFile.Header.Get("Content-Type")

	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/jpg" {

		return imageEncoded, apperrors.NewBadRequestError("only jpeg and png images are allowed")

	}

	data, err := io.ReadAll(file)

	if err != nil {

		return imageEncoded, apperrors.NewInternalServerError("error reading image file")

	}

	imageEncoded = base64.StdEncoding.EncodeToString(data)

	return imageEncoded, nil
}
