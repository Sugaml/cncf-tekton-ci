package utils

import (
	"encoding/json"

	"github.com/babulalt/tekton/models"
	"github.com/sirupsen/logrus"
)

func GetCIRequest(data interface{}) *models.CIRequest {
	ciRequest := models.CIRequest{}
	dataByte := data.([]byte)
	err := json.Unmarshal(dataByte, &ciRequest)
	if err != nil {
		logrus.Error(err)
	}
	return &ciRequest
}
