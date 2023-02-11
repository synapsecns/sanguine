package logging

// Note: we can't actually exclude this module from codeanalysis since we import it
//go:generate go run github.com/synapsecns/sanguine/tools/bundle -prefix ""  -pkg logging -o logging_gen.go github.com/hashicorp/terraform-plugin-sdk/v2/internal/logging
