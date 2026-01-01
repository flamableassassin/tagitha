package terraform

import (
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

var ignoredDirectories = []string{".terraform"}

// TODO: System to decide how to modify the Terraform when adding tags
// - Resources with existing tags
// - 		Making sure we update managed tags if they exist
// - Resources with tags that come from other sources like variables
// TODO: Sort out error handling
func parseFile(path string, fileInfo os.DirEntry, _ error) error {
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

	file, _ := hclsyntax.ParseConfig(data, fileInfo., hcl.InitialPos)
	log.Printf("test %d", len(file.Bytes))
	return nil
}

func ParseDirectory(directory string) ([]TaggableBlock, error) {
	//? Does cleaning convert between unix and windows separators
	directory = filepath.Clean(directory)
	
	filepath.WalkDir(directory, parseFile)

	return nil, nil
}
