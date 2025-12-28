package terraform

import (
	"log"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Tag struct {
	Name  string
	Value string
}

type TaggableBlock struct {
	Type      string
	Name      string
	Directory string
	FileName  string
	Location  hcl.Pos
	Tags      Tag
}

// TODO: System to decide how to modify the Terraform when adding tags
// - Resources with existing tags
// - 		Making sure we update managed tags if they exist
// - Resources with tags that come from other sources like variables
// TODO: Sort out error handling
func Parse() *TaggableBlock {
	data, err := os.ReadFile("/workspaces/tagitha/main.tf")
	if err != nil {
		log.Fatal(err)
	}
	file, _ := hclsyntax.ParseConfig(data, "main.tf", hcl.Pos{Byte: 0, Line: 0, Column: 0})
	log.Printf("test %d", len(file.Bytes))
	return nil
}
