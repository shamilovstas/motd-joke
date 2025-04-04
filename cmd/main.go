package main

import (
	"log/slog"
	"motd-joke/internal/joke"
	"motd-joke/internal/motd"
	"motd-joke/pkg/table"
	"motd-joke/pkg/textwrap"
)

func main() {
	joke, err := joke.GetJoke()
	if err != nil {
		slog.Error("cannot get top joke", slog.Any("error", err))
	}
	wrapped := textwrap.NewMinLines().Wrap(joke.String(), 40)
	t := table.Draw(wrapped, table.Border{HorizontalSymbol: '-', VerticalSymbol: '|'}, 1)
	if err := motd.CreateMotdfile(t); err != nil {
		slog.Error("cannot write to motd file", slog.Any("error", err))
	}
}
