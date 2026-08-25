package dto

import "randomNasa/internal/models"

type GetResponse struct {
	Info ResponseInfo `json:"info"`
	Results []models.User `json:"results"`
}

type ResponseInfo struct {
	Type string `json:"type"`
	TotalUsers int `json:"totalUsers"`
}