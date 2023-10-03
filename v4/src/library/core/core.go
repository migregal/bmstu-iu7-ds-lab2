package core

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/migregal/bmstu-iu7-ds-lab2/library/core/ports/libraries"
	"github.com/migregal/bmstu-iu7-ds-lab2/pkg/readiness"
)

type Core struct {
	libraries libraries.Client
}

func New(lg *slog.Logger, probe *readiness.Probe, library libraries.Client) (*Core, error) {
	probe.Mark("core", true)
	lg.Warn("[startup] core ready")

	return &Core{libraries: library}, nil
}

func (c *Core) GetLibraries(
	ctx context.Context, city string, page uint64, size uint64,
) (libraries.Libraries, error) {
	books, err := c.libraries.GetLibraries(ctx, city, page, size)
	if err != nil {
		return libraries.Libraries{}, fmt.Errorf("failed to get books: %w", err)
	}

	return books, nil
}

func (c *Core) GetLibraryBooks(
	ctx context.Context, libraryID string, showAll bool, page uint64, size uint64,
) (libraries.LibraryBooks, error) {
	books, err := c.libraries.GetLibraryBooks(ctx, libraryID, showAll, page, size)
	if err != nil {
		return libraries.LibraryBooks{}, fmt.Errorf("failed to get books: %w", err)
	}

	return books, nil
}

func (c *Core) TakeBook(ctx context.Context, libraryID, bookID string) (libraries.ReservedBook, error) {
	data, err := c.libraries.TakeBookFromLibrary(ctx, libraryID, bookID)
	if err != nil {
		return libraries.ReservedBook{}, fmt.Errorf("failed to take book from db: %w", err)
	}

	return data, nil
}
