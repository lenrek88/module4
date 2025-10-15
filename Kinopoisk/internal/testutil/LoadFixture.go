package testutil

import (
	"fmt"
	"os"
	"path/filepath"
)

func LoadFixture(dir, filename string) []byte {
	data, err := os.ReadFile(filepath.Join("..", dir, "testdata", filename))
	if err != nil {
		panic(fmt.Sprintf("failed to load fixture: %v", err))
	}

	return data
}
