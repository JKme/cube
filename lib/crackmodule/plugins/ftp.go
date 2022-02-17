package plugins

type FtpCrack struct {
	*Crack
}

func (ftpCrack *FtpCrack) SetName() (s string) {
	return "ftp"
}

func (ftpCrack *FtpCrack) IsLoad() (b bool) {
	return true
}
func (ftpCrack *FtpCrack) SetPort() (s string) {
	return "21"
}

func (ftpCrack *FtpCrack) Exec() (result CrackResult) {
	result = CrackResult{Crack: *ftpCrack.Crack, Result: "", Err: nil}

	return result
}

func init() {
	AddKeys("ftp")
}
