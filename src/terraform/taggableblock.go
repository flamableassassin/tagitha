package terraform

import "github.com/hashicorp/hcl/v2"

type TaggableBlock struct {
	Type       string
	Name       string
	Directory  string
	FileName   string
	Location   hcl.Pos
	Tags       map[string]string
	Attributes hcl.Attributes
}

// Should there be a parent which is the file. As modifying is done across the whole file
func (self TaggableBlock) SaveTags() {

}
