package motd

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"text/template"
)

const motdFragmentFilename = "99-joke-of-the-day"

func CreateMotdfile(message string) error {
	err := tryCreateGenericMotdFragment(message)
	if errors.Is(err, fs.ErrNotExist) {
		err = tryCreateDebianMotdFragment(message)

		if errors.Is(err, fs.ErrNotExist) {
			slog.Error("debian like fragment failed", slog.Any("error", err))
			err = createDefaultMotd(message)
		}
	}
	return err
}

func tryCreateGenericMotdFragment(message string) error {
	dirPath := "/etc/motd.d"
	fileInfo, err := os.Stat(dirPath)

	if err == nil {
		if !fileInfo.IsDir() {
			return fmt.Errorf("%s not a dir: %w", dirPath, err)
		}
		filename := fmt.Sprintf("%s.motd", motdFragmentFilename)
		path := filepath.Join(dirPath, filename)
		return os.WriteFile(path, []byte(message), 0644)
	} else {
		slog.Error("cannot stat /etc/motd.d", slog.Any("error", err))
		return err
	}
}

func tryCreateDebianMotdFragment(message string) error {
	slog.Info("creating debian-like motd fragment")
	dirPath := "/etc/update-motd.d"
	fileInfo, err := os.Stat(dirPath)

	if err == nil {
		if !fileInfo.IsDir() {
			return fmt.Errorf("%s not a dir: %w", dirPath, err)
		}
		path := filepath.Join(dirPath, motdFragmentFilename)
		motd, motdErr := os.Create(path)
		if motdErr != nil {
			return motdErr
		}
		motdErr = os.Chmod(path, 0755)
		if motdErr != nil {
			return motdErr
		}
		tpl, tplErr := template.ParseFiles("/usr/share/motd-joke/debian-motd-joke.tpl")

		if tplErr != nil {
			return tplErr
		}
		return tpl.Execute(motd, message)
	} else {
		return err
	}
}

func createDefaultMotd(message string) error {
	slog.Info("writing to /etc/motd")
	return os.WriteFile("/etc/motd", []byte(message), 0644)
}
