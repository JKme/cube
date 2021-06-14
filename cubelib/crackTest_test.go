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
	dict := []string{"111", "admin123", "admin", "12345678", "1234567", "p@$$w0rd", "passw0rd", "zhan1234", "19900307", "19850517"}

	//dict := []string{"admin123", "admin", "12345678", "1234567", "p@$$w0rd", "passw0rd", "Password1", "pass#123", "p@ssw0rd", "111"}
	for _, d := range dict {
		authList = append(authList, model.Auth{
			User:     "root",
			Password: d,
		})
	}
	runCrack(plugins, ips, authList)
}
