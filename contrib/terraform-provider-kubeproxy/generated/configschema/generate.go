package configschema

// Here: we copy the google module here to make some exports that we need for our module. While we could use module copier for this, the number of files
// would make the source directory illegible. Instead, we use bundler to bundle the package together and then make our exports from this generated file

// TODO: this currently breaks if ran from go generate, fix this.
// Note: we can't actually exclude this module from codeanalysis since we import it
//go:generate go run github.com/synapsecns/sanguine/tools/bundle -prefix ""  -pkg configschema -o configschema_gen.go github.com/hashicorp/terraform-plugin-sdk/v2/internal/configs/configschema
