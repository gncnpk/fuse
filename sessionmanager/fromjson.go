package sessionmanager

import (
	"context"
	"log/slog"
	"os"
	"path"
)

func FromJSON(ctx context.Context, name string) []byte {
	if ctx.Value("fromjson") == nil {
		return nil
	}

	pp := path.Join(".", "json", name)
	data, err := os.ReadFile(pp)
	if err != nil {
		slog.Error("from json", "error", err.Error(), "name", pp)
		return nil
	}
	slog.Warn("from json", "msgid", name)

	return data
}
