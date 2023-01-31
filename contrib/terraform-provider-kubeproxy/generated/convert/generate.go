package convert

// Note: we can't actually exclude this module from codeanalysis since we import it
//go:generate go run github.com/synapsecns/sanguine/tools/bundle -prefix ""  -pkg convert -o convert_gen.go -import github.com/hashicorp/terraform-plugin-sdk/v2/internal/logging=github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy/generated/logging -import github.com/hashicorp/terraform-plugin-sdk/v2/internal/configs/configschema=github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy/generated/configschema github.com/hashicorp/terraform-plugin-sdk/v2/internal/plugin/convert
