package model

import "time"

const (
	DefaultTime = 3
	TIMEOUT     = time.Duration(DefaultTime) * time.Second
)
