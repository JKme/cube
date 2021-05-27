package model

import "time"

const (
	TIMEUNIT = 3
	TIMEOUT  = time.Duration(TIMEUNIT) * time.Second
)
