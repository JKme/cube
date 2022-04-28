package config

import (
	"time"
)

const (
	TcpConnTimeout = 5 * time.Second
	ThreadTimeout  = 7 * time.Second
)

var CrackX = []string{"elastic", "ftp", "mongo", "mssql", "mysql", "postgres", "smb", "ssh", "redis", "oracle"}

var ProbeX = []string{"docker", "rmi", "oxid", "ms17010", "smb", "zookeeper", "dubbo", "etcd", "k8s", "smbghost", "jboss"}

var PASSWORDS = []string{" ", "123456", "admin", "admin123", "root", "5201314", "pass123", "pass@123", "password", "123123", "654321", "111111", "123", "1", "admin@123", "Admin@123", "admin123!@#", "1234qwer!@#$", "1qaz@WSX1qaz", "QAZwsxEDC", "{user}", "{user}1", "{user}12", "{user}111", "{user}123", "{user}1234", "{user}12345", "{user}123456", "{user}@123", "{user}_123", "{user}#123", "{user}@111", "{user}@2019", "P@ssw0rd!", "P@ssw0rd", "Passw0rd", "qwe123", "12345678", "test", "test123", "123qwe!@#", "123456789", "123321", "666666", "a123456.", "123456~a", "000000", "1234567890", "8888888", "!QAZ2wsx", "1qaz2wsx", "1QAZ2wsx", "1q2w3e4r", "abc123", "abc123456", "1qaz@WSX", "a11111", "a12345", "Aa1234", "Aa1234.", "Aa12345", "123456a", "123456aa", "a123456", "a123123", "Aa123123", "Aa123456", "Aa12345.", "sysadmin", "system"}
