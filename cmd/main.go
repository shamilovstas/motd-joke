package main

import (
	"fmt"
	"log/slog"
	"motd-joke/internal/joke"
	"motd-joke/internal/motd"
)

func main() {
	joke, err := joke.GetJoke()
	if err != nil {
		slog.Error("cannot get top joke", slog.Any("error", err))
	}

	motdFile, err := motd.MotdFile()

	if err != nil {
		slog.Error("cannot open /etc/motd", slog.Any("error", err))
	}

	if _, err := motdFile.WriteString(joke.String()); err != nil {
		slog.Error("cannot write to /etc/motd", slog.Any("error", err))
	}
	fmt.Println(joke.String())
}
