// +build ignore
//go:generate go run generate_mocks.go

package main

import (
	"os"
	"os/exec"
	"strings"
)

func main() {
	types := []string{
		"Tile",
		"RectangleDrawer",
		"BackgroundDrawer",
	}

	cmd := exec.Command("go", "run", "github.com/golang/mock/mockgen",
		"-destination", "mocks_test.go",
		"-package", "render_test",
		"github.com/everystreet/cartographer/internal/render",
		strings.Join(types, ","))
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
