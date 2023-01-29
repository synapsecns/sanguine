package configschema

// Note: we can't actually exclude this module from codeanalysis since we import it
//go:generate go run github.com/synapsecns/sanguine/tools/bundle -prefix ""  -pkg configschema -o configschema_gen.go github.com/hashicorp/terraform-plugin-sdk/v2/internal/configs/configschema
