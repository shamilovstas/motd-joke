package motd

import "os"

func MotdFile() (*os.File, error) {
	motd, err := os.Create("/etc/motd")
	return motd, err
}
