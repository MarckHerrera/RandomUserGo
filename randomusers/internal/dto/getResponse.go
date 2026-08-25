package dto

import (
	"randomusers/internal/models"
)

type GetResponse struct {
	Info ResponseInfo `json:"info"`
	Results []models.User `json:"results"`
}

type ResponseInfo struct {
	Type string `json:"type"`
	TotalUsers int `json:"totalUsers"`
}