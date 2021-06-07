package cubelib

import (
	"cube/model"
	"testing"
)

func TestName(t *testing.T) {
	plugins := []string{"SSH"}
	ips := []string{"1.1.1.1", "2.2.2.2", "3.3.3.3"}
	auth := model.Auth{
		User:     "root",
		Password: "root",
	}
	authList := []model.Auth{auth}
	runCrack(plugins, ips, authList)
}
