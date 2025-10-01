package db

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"database/sql"
)

func RunMigrations(db *sql.DB, dir string) error {
	var files []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(strings.ToLower(d.Name()), ".sql") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	sort.Strings(files)
	for _, f := range files {
		sqlBytes, err := os.ReadFile(f)
		if err != nil {
			return fmt.Errorf("leyendo %s: %w", f, err)
		}
		if _, err := db.Exec(string(sqlBytes)); err != nil {
			return fmt.Errorf("ejecutando %s: %w", f, err)
		}
	}
	return nil
}
