package plugins

type FtpCrack struct {
	Crack
}
func (ftpCrack *FtpCrack) Name() (s string){
	return "ftp"
}

func (ftpCrack *FtpCrack) Desc() (s string){
	return "crack ftp service password"
}
func (ftpCrack *FtpCrack) Load() (b bool){
	return true
}
func (ftpCrack *FtpCrack) Port() (s string){
	return "21"
}

func (ftpCrack *FtpCrack) Exec() (result CrackResult){
	result = CrackResult{Crack: ftpCrack.Crack, Result: "", Err: nil}

	return result
}