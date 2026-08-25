package models

import "time"

type RandomUserJobResult struct {
	Info ResultInfo
	Results []User
	Err error
	Duration time.Duration
}

type ResultInfo struct {
	Type string
	TotalUsers int
}