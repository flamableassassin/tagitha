package config

import "github.com/invopop/jsonschema"

func GetJSONSchema() *jsonschema.Schema {
	// r := new(jsonschema.Reflector)
	// if err := r.AddGoComments("github.com/invopop/jsonschema", "./"); err != nil {
	// deal with error
	// }
	// return r.Reflect(&Config{})
	return jsonschema.Reflect(&Config{})
}
