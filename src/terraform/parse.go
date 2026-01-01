package terraform

import (
	"log"
	"os"
	"path/filepath"
	"slices"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

var ignoredDirectories = []string{".terraform"}

type TaggableBlock struct {
	Type      string
	Name      string
	Directory string
	FileName  string
	Location  hcl.Pos
	Tags      map[string]string
}

// TODO: System to decide how to modify the Terraform when adding tags
// - Resources with existing tags
// - 		Making sure we update managed tags if they exist
// - Resources with tags that come from other sources like variables
// TODO: Sort out error handling
func ParseFile(path string, fileInfo os.FileInfo, _ error) error {
	if fileInfo.IsDir() {
		return nil
	}

	// Checking if ignored paths exists in the file path
	fileDirectories := strings.Split(path, os.PathSeparator)
	for _, ignoredDirectory := range ignoredDirectories {
		if slices.Contains(fileDirectories, ignoredDirectory) {
			return filepath.SkipDir
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	file, _ := hclsyntax.ParseConfig(data, "main.tf", hcl.Pos{Byte: 0, Line: 0, Column: 0})
	log.Printf("test %d", len(file.Bytes))
	return nil
}
