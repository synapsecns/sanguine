package manifest

import (
	apply_context "context"
	clients_context "context"
	configure_context "context"
	datasource_context "context"
	getproviderschema_context "context"
	import_context "context"
	plan_context "context"
	plugin_context "context"
	read_context "context"
	resource_context "context"
	server_context "context"
	upgrade_state_context "context"
	validate_context "context"
	waiter_context "context"
	plugin_json "encoding/json"
	resource_json "encoding/json"
	configure_pem "encoding/pem"
	apply_errors "errors"
	configure_errors "errors"
	resource_errors "errors"
	apply_fmt "fmt"
	clients_fmt "fmt"
	configure_fmt "fmt"
	datasource_fmt "fmt"
	diagnostics_fmt "fmt"
	import_fmt "fmt"
	plan_fmt "fmt"
	plugin_fmt "fmt"
	provider_fmt "fmt"
	read_fmt "fmt"
	resource_fmt "fmt"
	upgrade_state_fmt "fmt"
	validate_fmt "fmt"
	waiter_fmt "fmt"
	waiter_big "math/big"
	clients_http "net/http"
	configure_url "net/url"
	configure_os "os"
	plugin_os "os"
	configure_filepath "path/filepath"
	waiter_regexp "regexp"
	configure_strconv "strconv"
	configure_strings "strings"
	validate_strings "strings"
	plugin_testing "testing"
	apply_time "time"
	clients_time "time"
	plugin_time "time"
	validate_time "time"
	waiter_time "time"

	plugin_hclog "github.com/hashicorp/go-hclog"
	server_hclog "github.com/hashicorp/go-hclog"
	waiter_hclog "github.com/hashicorp/go-hclog"
	plugin_plugin "github.com/hashicorp/go-plugin"
	waiter_hclhcl "github.com/hashicorp/hcl/v2"
	waiter_hclsyntax "github.com/hashicorp/hcl/v2/hclsyntax"
	plugin_tfexec "github.com/hashicorp/terraform-exec/tfexec"
	apply_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	clients_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	configure_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	datasource_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	diagnostics_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	getproviderschema_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	import_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	plan_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	plugin_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	provider_config_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	provider_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	read_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	server_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	upgrade_state_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	validate_tfprotov5 "github.com/hashicorp/terraform-plugin-go/tfprotov5"
	plugin_tf5servertf5server "github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	apply_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	configure_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	datasource_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	import_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	plan_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	provider_config_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	provider_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	read_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	resource_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	upgrade_state_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	validate_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	waiter_tftypes "github.com/hashicorp/terraform-plugin-go/tftypes"
	clients_logging "github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	plan_manifest "github.com/hashicorp/terraform-provider-kubernetes/manifest"
	apply_morph "github.com/hashicorp/terraform-provider-kubernetes/manifest/morph"
	datasource_morph "github.com/hashicorp/terraform-provider-kubernetes/manifest/morph"
	import_morph "github.com/hashicorp/terraform-provider-kubernetes/manifest/morph"
	plan_morph "github.com/hashicorp/terraform-provider-kubernetes/manifest/morph"
	read_morph "github.com/hashicorp/terraform-provider-kubernetes/manifest/morph"
	upgrade_state_morph "github.com/hashicorp/terraform-provider-kubernetes/manifest/morph"
	clients_openapi "github.com/hashicorp/terraform-provider-kubernetes/manifest/openapi"
	resource_openapi "github.com/hashicorp/terraform-provider-kubernetes/manifest/openapi"
	server_openapi "github.com/hashicorp/terraform-provider-kubernetes/manifest/openapi"
	apply_payload "github.com/hashicorp/terraform-provider-kubernetes/manifest/payload"
	datasource_payload "github.com/hashicorp/terraform-provider-kubernetes/manifest/payload"
	import_payload "github.com/hashicorp/terraform-provider-kubernetes/manifest/payload"
	plan_payload "github.com/hashicorp/terraform-provider-kubernetes/manifest/payload"
	read_payload "github.com/hashicorp/terraform-provider-kubernetes/manifest/payload"
	waiter_payload "github.com/hashicorp/terraform-provider-kubernetes/manifest/payload"
	import_util "github.com/hashicorp/terraform-provider-kubernetes/util"
	configure_homedir "github.com/mitchellh/go-homedir"
	waiter_cty "github.com/zclconf/go-cty/cty"
	configure_semver "golang.org/x/mod/semver"
	server_codes "google.golang.org/grpc/codes"
	server_status "google.golang.org/grpc/status"
	server_install "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/install"
	apply_errorsapierrors "k8s.io/apimachinery/pkg/api/errors"
	clients_errorsapierrors "k8s.io/apimachinery/pkg/api/errors"
	datasource_errorsapierrors "k8s.io/apimachinery/pkg/api/errors"
	read_errorsapierrors "k8s.io/apimachinery/pkg/api/errors"
	waiter_errors "k8s.io/apimachinery/pkg/api/errors"
	clients_meta "k8s.io/apimachinery/pkg/api/meta"
	datasource_meta "k8s.io/apimachinery/pkg/api/meta"
	resource_meta "k8s.io/apimachinery/pkg/api/meta"
	server_meta "k8s.io/apimachinery/pkg/api/meta"
	apply_v1metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	datasource_v1metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	diagnostics_v1metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	import_v1metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	plan_v1metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	read_v1metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	resource_v1v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	waiter_v1v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apply_unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	datasource_unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	import_unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	plan_unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	read_unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	resource_unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	configure_runtime "k8s.io/apimachinery/pkg/runtime"
	configure_schemaapimachineryschema "k8s.io/apimachinery/pkg/runtime/schema"
	datasource_schema "k8s.io/apimachinery/pkg/runtime/schema"
	resource_schema "k8s.io/apimachinery/pkg/runtime/schema"
	configure_serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	apply_types "k8s.io/apimachinery/pkg/types"
	plan_types "k8s.io/apimachinery/pkg/types"
	clients_discovery "k8s.io/client-go/discovery"
	server_discovery "k8s.io/client-go/discovery"
	clients_memory "k8s.io/client-go/discovery/cached/memory"
	apply_dynamic "k8s.io/client-go/dynamic"
	clients_dynamic "k8s.io/client-go/dynamic"
	plan_dynamic "k8s.io/client-go/dynamic"
	server_dynamic "k8s.io/client-go/dynamic"
	waiter_dynamic "k8s.io/client-go/dynamic"
	configure_scheme "k8s.io/client-go/kubernetes/scheme"
	server_scheme "k8s.io/client-go/kubernetes/scheme"
	clients_rest "k8s.io/client-go/rest"
	configure_rest "k8s.io/client-go/rest"
	server_rest "k8s.io/client-go/rest"
	clients_restmapper "k8s.io/client-go/restmapper"
	configure_clientcmd "k8s.io/client-go/tools/clientcmd"
	configure_apiclientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	waiter_polymorphichelpers "k8s.io/kubectl/pkg/polymorphichelpers"
)

var defaultCreateTimeout = "10m"

var defaultUpdateTimeout = "10m"

var defaultDeleteTimeout = "10m"

func (s *RawProviderServer) ApplyResourceChange(ctx apply_context.Context, req *apply_tfprotov5.ApplyResourceChangeRequest) (*apply_tfprotov5.ApplyResourceChangeResponse, error) {
	resp := &apply_tfprotov5.ApplyResourceChangeResponse{}

	execDiag := s.canExecute()
	if len(execDiag) > 0 {
		resp.Diagnostics = append(resp.Diagnostics, execDiag...)
		return resp, nil
	}

	rt, err := GetResourceType(req.TypeName)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
			Severity: apply_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to determine planned resource type",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	applyPlannedState, err := req.PlannedState.Unmarshal(rt)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
			Severity: apply_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to unmarshal planned resource state",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	s.logger.Trace("[ApplyResourceChange][PlannedState] %#v", applyPlannedState)

	applyPriorState, err := req.PriorState.Unmarshal(rt)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
			Severity: apply_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to unmarshal prior resource state",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	s.logger.Trace("[ApplyResourceChange]", "[PriorState]", dump(applyPriorState))

	config, err := req.Config.Unmarshal(rt)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
			Severity: apply_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to unmarshal manifest configuration",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	confVals := make(map[string]apply_tftypes.Value)
	err = config.As(&confVals)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
			Severity: apply_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to extract attributes from resource configuration",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	var plannedStateVal map[string]apply_tftypes.Value = make(map[string]apply_tftypes.Value)
	err = applyPlannedState.As(&plannedStateVal)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
			Severity: apply_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to extract planned resource state values",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	computedFields := make(map[string]*apply_tftypes.AttributePath)
	var atp *apply_tftypes.AttributePath
	cfVal, ok := plannedStateVal["computed_fields"]
	if ok && !cfVal.IsNull() && cfVal.IsKnown() {
		var cf []apply_tftypes.Value
		cfVal.As(&cf)
		for _, v := range cf {
			var vs string
			err := v.As(&vs)
			if err != nil {
				s.logger.Error("[computed_fields] cannot extract element from list")
				continue
			}
			atp, err := FieldPathToTftypesPath(vs)
			if err != nil {
				s.logger.Error("[Configure]", "[computed_fields] cannot parse field path element", err)
				resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
					Severity: apply_tfprotov5.DiagnosticSeverityError,
					Summary:  "[computed_fields] cannot parse filed path element: " + vs,
					Detail:   err.Error(),
				})
				continue
			}
			computedFields[atp.String()] = atp
		}
	} else {

		atp = apply_tftypes.NewAttributePath().WithAttributeName("metadata").WithAttributeName("annotations")
		computedFields[atp.String()] = atp

		atp = apply_tftypes.NewAttributePath().WithAttributeName("metadata").WithAttributeName("labels")
		computedFields[atp.String()] = atp
	}

	c, err := s.getDynamicClient()
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics,
			&apply_tfprotov5.Diagnostic{
				Severity: apply_tfprotov5.DiagnosticSeverityError,
				Summary:  "Failed to retrieve Kubernetes dynamic client during apply",
				Detail:   err.Error(),
			})
		return resp, nil
	}
	m, err := s.getRestMapper()
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics,
			&apply_tfprotov5.Diagnostic{
				Severity: apply_tfprotov5.DiagnosticSeverityError,
				Summary:  "Failed to retrieve Kubernetes RESTMapper client during apply",
				Detail:   err.Error(),
			})
		return resp, nil
	}
	var rs apply_dynamic.ResourceInterface

	switch {
	case applyPriorState.IsNull() || (!applyPlannedState.IsNull() && !applyPriorState.IsNull()):

		obj, ok := plannedStateVal["object"]
		if !ok {
			resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
				Severity: apply_tfprotov5.DiagnosticSeverityError,
				Summary:  "Failed to find object value in planned resource state",
			})
			return resp, nil
		}

		gvk, err := GVKFromTftypesObject(&obj, m)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
				Severity: apply_tfprotov5.DiagnosticSeverityError,
				Summary:  "Failed to determine the type of the resource",
				Detail:   apply_fmt.Sprintf(`This can happen when the "apiVersion" or "kind" fields are not present in the manifest, or when the corresponding "kind" or "apiVersion" could not be found on the Kubernetes cluster.\nError: %s`, err),
			})
			return resp, nil
		}

		tsch, th, err := s.TFTypeFromOpenAPI(ctx, gvk, false)
		if err != nil {
			return resp, apply_fmt.Errorf("failed to determine resource type ID: %s", err)
		}

		obj, err = apply_tftypes.Transform(obj, func(ap *apply_tftypes.AttributePath, v apply_tftypes.Value) (apply_tftypes.Value, error) {
			_, isComputed := computedFields[ap.String()]
			if !isComputed {
				return v, nil
			}
			if v.IsKnown() {
				return v, nil
			}
			ppMan, restPath, err := apply_tftypes.WalkAttributePath(plannedStateVal["manifest"], ap)
			if err != nil {
				if len(restPath.Steps()) > 0 {

					return v, nil
				}
				return v, ap.NewError(err)
			}
			nv, d := apply_morph.ValueToType(ppMan.(apply_tftypes.Value), v.Type(), apply_tftypes.NewAttributePath())
			if len(d) > 0 {
				resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
					Severity: apply_tfprotov5.DiagnosticSeverityError,
					Summary:  "Manifest configuration is incompatible with resource schema",
					Detail:   "Detailed descriptions of errors will follow below.",
				})
				resp.Diagnostics = append(resp.Diagnostics, d...)
				return v, nil
			}
			return nv, nil
		})
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
				Severity: apply_tfprotov5.DiagnosticSeverityError,
				Summary:  "Failed to backfill computed values in proposed value",
				Detail:   err.Error(),
			})
			return resp, nil
		}

		nullObj := apply_morph.UnknownToNull(obj)
		s.logger.Trace("[ApplyResourceChange][Apply]", "[UnknownToNull]", dump(nullObj))

		minObj, err := apply_tftypes.Transform(nullObj, func(ap *apply_tftypes.AttributePath, v apply_tftypes.Value) (apply_tftypes.Value, error) {
			if v.IsNull() {
				return apply_tftypes.NewValue(v.Type(), nil), nil
			}
			switch {
			case v.Type().Is(apply_tftypes.Object{}) || v.Type().Is(apply_tftypes.Map{}):
				atts := make(map[string]apply_tftypes.Value)
				err := v.As(&atts)
				if err != nil {
					return v, err
				}
				var isEmpty bool = true
				for _, atv := range atts {
					if !atv.IsNull() {
						isEmpty = false
						break
					}
				}

				_, restPath, err := apply_tftypes.WalkAttributePath(confVals["manifest"], ap)
				if (err == nil && len(restPath.Steps()) == 0) || !isEmpty {

					return apply_tftypes.NewValue(v.Type(), atts), nil
				}
				return apply_tftypes.NewValue(v.Type(), nil), nil
			case v.Type().Is(apply_tftypes.List{}) || v.Type().Is(apply_tftypes.Set{}) || v.Type().Is(apply_tftypes.Tuple{}):
				atts := make([]apply_tftypes.Value, 0)
				err := v.As(&atts)
				if err != nil {
					return v, err
				}
				return apply_tftypes.NewValue(v.Type(), atts), nil
			default:
				return v, nil
			}
		})
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics,
				&apply_tfprotov5.Diagnostic{
					Severity: apply_tfprotov5.DiagnosticSeverityError,
					Detail:   err.Error(),
					Summary:  "Failed to sanitize empty block ahead of payload preparation",
				})
			return resp, nil
		}

		pu, err := apply_payload.FromTFValue(minObj, th, apply_tftypes.NewAttributePath())
		if err != nil {
			return resp, err
		}
		s.logger.Trace("[ApplyResourceChange][Apply]", "[payload.FromTFValue]", dump(pu))

		rqObj := mapRemoveNulls(pu.(map[string]interface{}))

		uo := apply_unstructured.Unstructured{}
		uo.SetUnstructuredContent(rqObj)
		rnamespace := uo.GetNamespace()
		rname := uo.GetName()
		rnn := apply_types.NamespacedName{Namespace: rnamespace, Name: rname}.String()

		gvr, err := GVRFromUnstructured(&uo, m)
		if err != nil {
			return resp, apply_fmt.Errorf("failed to determine resource GVR: %s", err)
		}

		ns, err := IsResourceNamespaced(gvk, m)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics,
				&apply_tfprotov5.Diagnostic{
					Severity: apply_tfprotov5.DiagnosticSeverityError,
					Detail:   err.Error(),
					Summary:  apply_fmt.Sprintf("Failed to discover scope of resource '%s'", rnn),
				})
			return resp, nil
		}

		if ns {
			rs = c.Resource(gvr).Namespace(rnamespace)
		} else {
			rs = c.Resource(gvr)
		}

		if applyPriorState.IsNull() {
			_, err := rs.Get(ctx, rname, apply_v1metav1.GetOptions{})
			if err == nil {
				resp.Diagnostics = append(resp.Diagnostics,
					&apply_tfprotov5.Diagnostic{
						Severity: apply_tfprotov5.DiagnosticSeverityError,
						Summary:  "Cannot create resource that already exists",
						Detail:   apply_fmt.Sprintf("resource %q already exists", rnn),
					})
				return resp, nil
			} else if !apply_errorsapierrors.IsNotFound(err) {
				resp.Diagnostics = append(resp.Diagnostics,
					&apply_tfprotov5.Diagnostic{
						Severity: apply_tfprotov5.DiagnosticSeverityError,
						Summary:  apply_fmt.Sprintf("Failed to determine if resource %q exists", rnn),
						Detail:   err.Error(),
					})
				return resp, nil
			}
		}

		jsonManifest, err := uo.MarshalJSON()
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics,
				&apply_tfprotov5.Diagnostic{
					Severity: apply_tfprotov5.DiagnosticSeverityError,
					Detail:   err.Error(),
					Summary:  apply_fmt.Sprintf("Failed to marshall resource '%s' to JSON", rnn),
				})
			return resp, nil
		}

		fieldManagerName, forceConflicts, err := s.getFieldManagerConfig(plannedStateVal)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
				Severity: apply_tfprotov5.DiagnosticSeverityError,
				Summary:  "Could not extract field_manager config",
				Detail:   err.Error(),
			})
			return resp, nil
		}

		timeouts := s.getTimeouts(plannedStateVal)
		var timeout apply_time.Duration
		if applyPriorState.IsNull() {
			timeout, _ = apply_time.ParseDuration(timeouts["create"])
		} else {
			timeout, _ = apply_time.ParseDuration(timeouts["update"])
		}
		deadline := apply_time.Now().Add(timeout)
		ctxDeadline, cancel := apply_context.WithDeadline(ctx, deadline)
		defer cancel()

		s.logger.Trace("[ApplyResourceChange][API Payload]: %s", jsonManifest)
		result, err := rs.Patch(ctxDeadline, rname, apply_types.ApplyPatchType, jsonManifest,
			apply_v1metav1.PatchOptions{
				FieldManager: fieldManagerName,
				Force:        &forceConflicts,
			},
		)
		if err != nil {
			s.logger.Error("[ApplyResourceChange][Apply]", "API error", dump(err), "API response", dump(result))
			if apply_errorsapierrors.IsConflict(err) {
				resp.Diagnostics = append(resp.Diagnostics,
					&apply_tfprotov5.Diagnostic{
						Severity: apply_tfprotov5.DiagnosticSeverityError,
						Summary:  apply_fmt.Sprintf(`There was a field manager conflict when trying to apply the manifest for %q`, rnn),
						Detail: apply_fmt.Sprintf(
							"The API returned the following conflict: %q\n\n"+
								"You can override this conflict by setting \"force_conflicts\" to true in the \"field_manager\" block.",
							err.Error(),
						),
					},
				)
			} else if status := apply_errorsapierrors.APIStatus(nil); apply_errors.As(err, &status) {
				resp.Diagnostics = append(resp.Diagnostics, APIStatusErrorToDiagnostics(status.Status())...)
			} else {
				resp.Diagnostics = append(resp.Diagnostics,
					&apply_tfprotov5.Diagnostic{
						Severity: apply_tfprotov5.DiagnosticSeverityError,
						Detail:   err.Error(),
						Summary:  apply_fmt.Sprintf(`PATCH for resource "%s" failed to apply`, rnn),
					})
			}
			return resp, nil
		}

		newResObject, err := apply_payload.ToTFValue(RemoveServerSideFields(result.Object), tsch, th, apply_tftypes.NewAttributePath())
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics,
				&apply_tfprotov5.Diagnostic{
					Severity: apply_tfprotov5.DiagnosticSeverityError,
					Summary:  "Conversion from Unstructured to tftypes.Value failed",
					Detail:   err.Error(),
				})
			return resp, nil
		}
		s.logger.Trace("[ApplyResourceChange][Apply]", "[payload.ToTFValue]", dump(newResObject))

		wt, _, err := s.TFTypeFromOpenAPI(ctx, gvk, true)
		if err != nil {
			return resp, apply_fmt.Errorf("failed to determine resource type ID: %s", err)
		}

		var waitConfig apply_tftypes.Value
		if w, ok := plannedStateVal["wait"]; ok && !w.IsNull() {
			s.logger.Trace("[ApplyResourceChange][Wait] Using waiter config from `wait` block")
			var waitBlocks []apply_tftypes.Value
			w.As(&waitBlocks)
			if len(waitBlocks) > 0 {
				waitConfig = waitBlocks[0]
			}
		}
		if wf, ok := plannedStateVal["wait_for"]; ok && !wf.IsNull() {
			s.logger.Trace("[ApplyResourceChange][Wait] Using waiter config from deprecated `wait_for` attribute")
			waitConfig = wf
		}
		if !waitConfig.IsNull() {
			err = s.waitForCompletion(ctxDeadline, waitConfig, rs, rname, wt, th)
			if err != nil {
				if err == apply_context.DeadlineExceeded {
					resp.Diagnostics = append(resp.Diagnostics,
						&apply_tfprotov5.Diagnostic{
							Severity: apply_tfprotov5.DiagnosticSeverityError,
							Summary:  "Operation timed out",
							Detail:   "Terraform timed out waiting on the operation to complete",
						})
				} else {
					resp.Diagnostics = append(resp.Diagnostics,
						&apply_tfprotov5.Diagnostic{
							Severity: apply_tfprotov5.DiagnosticSeverityError,
							Summary:  "Error waiting for operation to complete",
							Detail:   err.Error(),
						})
				}
				return resp, nil
			}
		}

		compObj, err := apply_morph.DeepUnknown(tsch, newResObject, apply_tftypes.NewAttributePath())
		if err != nil {
			return resp, err
		}
		plannedStateVal["object"] = apply_morph.UnknownToNull(compObj)

		newStateVal := apply_tftypes.NewValue(applyPlannedState.Type(), plannedStateVal)
		s.logger.Trace("[ApplyResourceChange][Apply]", "new state value", dump(newStateVal))

		newResState, err := apply_tfprotov5.NewDynamicValue(newStateVal.Type(), newStateVal)
		if err != nil {
			return resp, err
		}
		resp.NewState = &newResState
	case applyPlannedState.IsNull():

		priorStateVal := make(map[string]apply_tftypes.Value)
		err = applyPriorState.As(&priorStateVal)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
				Severity: apply_tfprotov5.DiagnosticSeverityError,
				Summary:  "Failed to extract prior resource state values",
				Detail:   err.Error(),
			})
			return resp, nil
		}
		pco, ok := priorStateVal["object"]
		if !ok {
			resp.Diagnostics = append(resp.Diagnostics, &apply_tfprotov5.Diagnostic{
				Severity: apply_tfprotov5.DiagnosticSeverityError,
				Summary:  "Failed to find object value in prior resource state",
			})
			return resp, nil
		}

		pu, err := apply_payload.FromTFValue(pco, nil, apply_tftypes.NewAttributePath())
		if err != nil {
			return resp, err
		}

		uo := apply_unstructured.Unstructured{Object: pu.(map[string]interface{})}
		gvr, err := GVRFromUnstructured(&uo, m)
		if err != nil {
			return resp, err
		}

		gvk, err := GVKFromTftypesObject(&pco, m)
		if err != nil {
			return resp, apply_fmt.Errorf("failed to determine resource GVK: %s", err)
		}

		ns, err := IsResourceNamespaced(gvk, m)
		if err != nil {
			return resp, err
		}
		rnamespace := uo.GetNamespace()
		rname := uo.GetName()
		if ns {
			rs = c.Resource(gvr).Namespace(rnamespace)
		} else {
			rs = c.Resource(gvr)
		}

		timeouts := s.getTimeouts(priorStateVal)
		timeout, _ := apply_time.ParseDuration(timeouts["delete"])
		deadline := apply_time.Now().Add(timeout)
		ctxDeadline, cancel := apply_context.WithDeadline(ctx, deadline)
		defer cancel()

		err = rs.Delete(ctxDeadline, rname, apply_v1metav1.DeleteOptions{})
		if err != nil {
			rn := apply_types.NamespacedName{Namespace: rnamespace, Name: rname}.String()
			resp.Diagnostics = append(resp.Diagnostics,
				&apply_tfprotov5.Diagnostic{
					Severity: apply_tfprotov5.DiagnosticSeverityError,
					Summary:  apply_fmt.Sprintf("Error deleting resource %s: %s", rn, err),
					Detail:   err.Error(),
				})
			return resp, nil
		}

		for {
			if apply_time.Now().After(deadline) {
				resp.Diagnostics = append(resp.Diagnostics,
					&apply_tfprotov5.Diagnostic{
						Severity: apply_tfprotov5.DiagnosticSeverityError,
						Summary:  apply_fmt.Sprintf("Timed out when waiting for resource %q to be deleted", rname),
						Detail:   "Deletion timed out. This can happen when there is a finalizer on a resource. You may need to delete this resource manually with kubectl.",
					})
				return resp, nil
			}
			_, err := rs.Get(ctxDeadline, rname, apply_v1metav1.GetOptions{})
			if err != nil {
				if apply_errorsapierrors.IsNotFound(err) {
					s.logger.Trace("[ApplyResourceChange][Delete]", "Resource is deleted")
					break
				}
				resp.Diagnostics = append(resp.Diagnostics,
					&apply_tfprotov5.Diagnostic{
						Severity: apply_tfprotov5.DiagnosticSeverityError,
						Summary:  "Error waiting for deletion.",
						Detail:   apply_fmt.Sprintf("Error when waiting for resource %q to be deleted: %v", rname, err),
					})
				return resp, nil
			}
			apply_time.Sleep(1 * apply_time.Second)
		}

		resp.NewState = req.PlannedState
	}

	return resp, nil
}

func (s *RawProviderServer) getTimeouts(v map[string]apply_tftypes.Value) map[string]string {
	timeouts := map[string]string{
		"create": defaultCreateTimeout,
		"update": defaultUpdateTimeout,
		"delete": defaultDeleteTimeout,
	}
	if !v["timeouts"].IsNull() && v["timeouts"].IsKnown() {
		var timeoutsBlock []apply_tftypes.Value
		v["timeouts"].As(&timeoutsBlock)
		if len(timeoutsBlock) > 0 {
			var t map[string]apply_tftypes.Value
			timeoutsBlock[0].As(&t)
			var s string
			for _, k := range []string{"create", "update", "delete"} {
				if vv, ok := t[k]; ok && !vv.IsNull() {
					vv.As(&s)
					if s != "" {
						timeouts[k] = s
					}
				}
			}
		}
	}
	return timeouts
}

const (
	OAPIFoundry string = "OPENAPIFOUNDRY"
)

func (ps *RawProviderServer) getDynamicClient() (clients_dynamic.Interface, error) {
	if ps.dynamicClient != nil {
		return ps.dynamicClient, nil
	}
	if ps.clientConfig == nil {
		return nil, clients_fmt.Errorf("cannot create dynamic client: no client config")
	}
	dynClient, err := clients_dynamic.NewForConfig(ps.clientConfig)
	if err != nil {
		return nil, err
	}
	ps.dynamicClient = dynClient
	return dynClient, nil
}

func (ps *RawProviderServer) getDiscoveryClient() (clients_discovery.DiscoveryInterface, error) {
	if ps.discoveryClient != nil {
		return ps.discoveryClient, nil
	}
	if ps.clientConfig == nil {
		return nil, clients_fmt.Errorf("cannot create discovery client: no client config")
	}
	discoClient, err := clients_discovery.NewDiscoveryClientForConfig(ps.clientConfig)
	if err != nil {
		return nil, err
	}
	ps.discoveryClient = discoClient
	return discoClient, nil
}

func (ps *RawProviderServer) getRestMapper() (clients_meta.RESTMapper, error) {
	if ps.restMapper != nil {
		return ps.restMapper, nil
	}
	dc, err := ps.getDiscoveryClient()
	if err != nil {
		return nil, err
	}

	cache := clients_memory.NewMemCacheClient(dc)
	ps.restMapper = clients_restmapper.NewDeferredDiscoveryRESTMapper(cache)
	return ps.restMapper, nil
}

func (ps *RawProviderServer) getRestClient() (clients_rest.Interface, error) {
	if ps.restClient != nil {
		return ps.restClient, nil
	}
	if ps.clientConfig == nil {
		return nil, clients_fmt.Errorf("cannot create REST client: no client config")
	}
	restClient, err := clients_rest.UnversionedRESTClientFor(ps.clientConfig)
	if err != nil {
		return nil, err
	}
	ps.restClient = restClient
	return restClient, nil
}

func (ps *RawProviderServer) getOAPIv2Foundry() (clients_openapi.Foundry, error) {
	if ps.OAPIFoundry != nil {
		return ps.OAPIFoundry, nil
	}

	rc, err := ps.getRestClient()
	if err != nil {
		return nil, clients_fmt.Errorf("failed get OpenAPI spec: %s", err)
	}

	rq := rc.Verb("GET").Timeout(30*clients_time.Second).AbsPath("openapi", "v2")
	rs, err := rq.DoRaw(clients_context.TODO())
	if err != nil {
		return nil, clients_fmt.Errorf("failed get OpenAPI spec: %s", err)
	}

	oapif, err := clients_openapi.NewFoundryFromSpecV2(rs)
	if err != nil {
		return nil, clients_fmt.Errorf("failed construct OpenAPI foundry: %s", err)
	}

	ps.OAPIFoundry = oapif

	return oapif, nil
}

func loggingTransport(rt clients_http.RoundTripper) clients_http.RoundTripper {
	return &loggingRountTripper{
		ot: rt,
		lt: clients_logging.NewTransport("Kubernetes API", rt),
	}
}

type loggingRountTripper struct {
	ot clients_http.RoundTripper
	lt clients_http.RoundTripper
}

func (t *loggingRountTripper) RoundTrip(req *clients_http.Request) (*clients_http.Response, error) {
	if req.URL.Path == "/openapi/v2" {

		return t.ot.RoundTrip(req)
	}
	return t.lt.RoundTrip(req)
}

func (ps *RawProviderServer) checkValidCredentials(ctx clients_context.Context) (diags []*clients_tfprotov5.Diagnostic) {
	rc, err := ps.getRestClient()
	if err != nil {
		diags = append(diags, &clients_tfprotov5.Diagnostic{
			Severity: clients_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to construct REST client",
			Detail:   err.Error(),
		})
		return
	}
	vpath := []string{"/apis"}
	rs := rc.Get().AbsPath(vpath...).Do(ctx)
	if rs.Error() != nil {
		switch {
		case clients_errorsapierrors.IsUnauthorized(rs.Error()):
			diags = append(diags, &clients_tfprotov5.Diagnostic{
				Severity: clients_tfprotov5.DiagnosticSeverityError,
				Summary:  "Invalid credentials",
				Detail:   clients_fmt.Sprintf("The credentials configured in the provider block are not accepted by the API server. Error: %s\n\nSet TF_LOG=debug and look for '[InvalidClientConfiguration]' in the log to see actual configuration.", rs.Error().Error()),
			})
		default:
			diags = append(diags, &clients_tfprotov5.Diagnostic{
				Severity: clients_tfprotov5.DiagnosticSeverityError,
				Summary:  "Invalid configuration for API client",
				Detail:   rs.Error().Error(),
			})
		}
		ps.logger.Debug("[InvalidClientConfiguration]", "Config", dump(ps.clientConfig))
	}
	return
}

const minTFVersion string = "v0.14.8"

func (s *RawProviderServer) ConfigureProvider(ctx configure_context.Context, req *configure_tfprotov5.ConfigureProviderRequest) (*configure_tfprotov5.ConfigureProviderResponse, error) {
	response := &configure_tfprotov5.ConfigureProviderResponse{}
	diags := []*configure_tfprotov5.Diagnostic{}
	var providerConfig map[string]configure_tftypes.Value
	var err error

	s.hostTFVersion = "v" + req.TerraformVersion

	cfgType := GetObjectTypeFromSchema(GetProviderConfigSchema())
	cfgVal, err := req.Config.Unmarshal(cfgType)
	if err != nil {
		response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
			Severity: configure_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to decode ConfigureProvider request parameter",
			Detail:   err.Error(),
		})
		return response, nil
	}
	err = cfgVal.As(&providerConfig)
	if err != nil {

		response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
			Severity: configure_tfprotov5.DiagnosticSeverityError,
			Summary:  "Provider configuration: failed to extract 'config_path' value",
			Detail:   err.Error(),
		})
		return response, nil
	}

	providerEnabled := true
	if !providerConfig["experiments"].IsNull() && providerConfig["experiments"].IsKnown() {
		var experimentsBlock []configure_tftypes.Value
		err = providerConfig["experiments"].As(&experimentsBlock)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to extract 'experiments' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
		if len(experimentsBlock) > 0 {
			var experimentsObj map[string]configure_tftypes.Value
			err := experimentsBlock[0].As(&experimentsObj)
			if err != nil {

				response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
					Severity: configure_tfprotov5.DiagnosticSeverityError,
					Summary:  "Provider configuration: failed to extract 'experiments' value",
					Detail:   err.Error(),
				})
				return response, nil
			}
			if !experimentsObj["manifest_resource"].IsNull() && experimentsObj["manifest_resource"].IsKnown() {
				err = experimentsObj["manifest_resource"].As(&providerEnabled)
				if err != nil {

					response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
						Severity: configure_tfprotov5.DiagnosticSeverityError,
						Summary:  "Provider configuration: failed to extract 'manifest_resource' value",
						Detail:   err.Error(),
					})
					return response, nil
				}
			}
		}
	}
	if v := configure_os.Getenv("TF_X_KUBERNETES_MANIFEST_RESOURCE"); v != "" {
		providerEnabled, err = configure_strconv.ParseBool(v)
		if err != nil {
			if err != nil {

				response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
					Severity: configure_tfprotov5.DiagnosticSeverityError,
					Summary:  "Provider configuration: failed to parse boolean from `TF_X_KUBERNETES_MANIFEST_RESOURCE` env var",
					Detail:   err.Error(),
				})
				return response, nil
			}
		}
	}
	s.providerEnabled = providerEnabled

	if !providerEnabled {

		return response, nil
	}

	overrides := &configure_clientcmd.ConfigOverrides{}
	loader := &configure_clientcmd.ClientConfigLoadingRules{}

	var configPath string
	if !providerConfig["config_path"].IsNull() && providerConfig["config_path"].IsKnown() {
		err = providerConfig["config_path"].As(&configPath)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to extract 'config_path' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
	}

	if configPathEnv, ok := configure_os.LookupEnv("KUBE_CONFIG_PATH"); ok && configPathEnv != "" {
		configPath = configPathEnv
	}
	if len(configPath) > 0 {
		configPathAbs, err := configure_homedir.Expand(configPath)
		if err == nil {
			_, err = configure_os.Stat(configPathAbs)
		}
		if err != nil {
			diags = append(diags, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityInvalid,
				Summary:  "Invalid attribute in provider configuration",
				Detail:   configure_fmt.Sprintf("'config_path' refers to an invalid path: %q: %v", configPathAbs, err),
			})
		}
		loader.ExplicitPath = configPathAbs
	}

	var precedence []string
	if !providerConfig["config_paths"].IsNull() && providerConfig["config_paths"].IsKnown() {
		var configPaths []configure_tftypes.Value
		err = providerConfig["config_paths"].As(&configPaths)
		if err != nil {
			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to extract 'config_paths' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
		for _, p := range configPaths {
			var pp string
			p.As(&pp)
			precedence = append(precedence, pp)
		}
	}

	if configPathsEnv, ok := configure_os.LookupEnv("KUBE_CONFIG_PATHS"); ok && configPathsEnv != "" {
		precedence = configure_filepath.SplitList(configPathsEnv)
	}
	if len(precedence) > 0 {
		for i, p := range precedence {
			absPath, err := configure_homedir.Expand(p)
			if err == nil {
				_, err = configure_os.Stat(absPath)
			}
			if err != nil {
				diags = append(diags, &configure_tfprotov5.Diagnostic{
					Severity: configure_tfprotov5.DiagnosticSeverityInvalid,
					Summary:  "Invalid attribute in provider configuration",
					Detail:   configure_fmt.Sprintf("'config_paths' refers to an invalid path: %q: %v", absPath, err),
				})
			}
			precedence[i] = absPath
		}
		loader.Precedence = precedence
	}

	var clientCertificate string
	if !providerConfig["client_certificate"].IsNull() && providerConfig["client_certificate"].IsKnown() {
		err = providerConfig["client_certificate"].As(&clientCertificate)
		if err != nil {
			diags = append(diags, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityInvalid,
				Summary:  "Invalid attribute in provider configuration",
				Detail:   "'client_certificate' type cannot be asserted: " + err.Error(),
			})
			return response, nil
		}
	}
	if clientCrtEnv, ok := configure_os.LookupEnv("KUBE_CLIENT_CERT_DATA"); ok && clientCrtEnv != "" {
		clientCertificate = clientCrtEnv
	}
	if len(clientCertificate) > 0 {
		ccPEM, _ := configure_pem.Decode([]byte(clientCertificate))
		if ccPEM == nil || ccPEM.Type != "CERTIFICATE" {
			diags = append(diags, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityInvalid,
				Summary:  "Invalid attribute in provider configuration",
				Detail:   "'client_certificate' is not a valid PEM encoded certificate",
			})
		}
		overrides.AuthInfo.ClientCertificateData = []byte(clientCertificate)
	}

	var clusterCaCertificate string
	if !providerConfig["cluster_ca_certificate"].IsNull() && providerConfig["cluster_ca_certificate"].IsKnown() {
		err = providerConfig["cluster_ca_certificate"].As(&clusterCaCertificate)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to extract 'cluster_ca_certificate' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
	}
	if clusterCAEnv, ok := configure_os.LookupEnv("KUBE_CLUSTER_CA_CERT_DATA"); ok && clusterCAEnv != "" {
		clusterCaCertificate = clusterCAEnv
	}
	if len(clusterCaCertificate) > 0 {
		ca, _ := configure_pem.Decode([]byte(clusterCaCertificate))
		if ca == nil || ca.Type != "CERTIFICATE" {
			diags = append(diags, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityInvalid,
				Summary:  "Invalid attribute in provider configuration",
				Detail:   "'cluster_ca_certificate' is not a valid PEM encoded certificate",
			})
		}
		overrides.ClusterInfo.CertificateAuthorityData = []byte(clusterCaCertificate)
	}

	var insecure bool
	if !providerConfig["insecure"].IsNull() && providerConfig["insecure"].IsKnown() {
		err = providerConfig["insecure"].As(&insecure)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to assert type of 'insecure' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
	}
	if insecureEnv, ok := configure_os.LookupEnv("KUBE_INSECURE"); ok && insecureEnv != "" {
		iv, err := configure_strconv.ParseBool(insecureEnv)
		if err != nil {
			diags = append(diags, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityInvalid,
				Summary:  "Invalid provider configuration",
				Detail:   "Environment variable KUBE_INSECURE contains invalid value: " + err.Error(),
			})
		} else {
			insecure = iv
		}
	}
	overrides.ClusterInfo.InsecureSkipTLSVerify = insecure

	hasCA := len(overrides.ClusterInfo.CertificateAuthorityData) != 0
	hasCert := len(overrides.AuthInfo.ClientCertificateData) != 0
	defaultTLS := hasCA || hasCert || overrides.ClusterInfo.InsecureSkipTLSVerify

	var host string
	if !providerConfig["host"].IsNull() && providerConfig["host"].IsKnown() {
		err = providerConfig["host"].As(&host)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to extract 'host' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
	}

	if hostEnv, ok := configure_os.LookupEnv("KUBE_HOST"); ok && hostEnv != "" {
		host = hostEnv
	}
	if len(host) > 0 {
		_, err = configure_url.ParseRequestURI(host)
		if err != nil {
			diags = append(diags, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityInvalid,
				Summary:  "Invalid attribute in provider configuration",
				Detail:   "'host' is not a valid URL",
			})
		}
		hostURL, _, err := configure_rest.DefaultServerURL(host, "", configure_schemaapimachineryschema.GroupVersion{}, defaultTLS)
		if err != nil {
			diags = append(diags, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityInvalid,
				Summary:  "Invalid attribute in provider configuration",
				Detail:   "Invalid value for 'host': " + err.Error(),
			})
			return response, nil
		}

		overrides.ClusterInfo.Server = hostURL.String()
	}

	var clientKey string
	if !providerConfig["client_key"].IsNull() && providerConfig["client_key"].IsKnown() {
		err = providerConfig["client_key"].As(&clientKey)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: ",
				Detail:   "Failed to extract 'client_key' value" + err.Error(),
			})
			return response, nil
		}
	}

	if clientKeyEnv, ok := configure_os.LookupEnv("KUBE_CLIENT_KEY_DATA"); ok && clientKeyEnv != "" {
		clientKey = clientKeyEnv
	}
	if len(clientKey) > 0 {
		ck, _ := configure_pem.Decode([]byte(clientKey))
		if ck == nil || !configure_strings.Contains(ck.Type, "PRIVATE KEY") {
			diags = append(diags, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityInvalid,
				Summary:  "Invalid attribute in provider configuration",
				Detail:   "'client_key' is not a valid PEM encoded private key",
			})
		}
		overrides.AuthInfo.ClientKeyData = []byte(clientKey)
	}

	if len(diags) > 0 {
		response.Diagnostics = diags
		return response, nil
	}

	var cfgContext string
	if !providerConfig["config_context"].IsNull() && providerConfig["config_context"].IsKnown() {
		err = providerConfig["config_context"].As(&cfgContext)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to assert type of 'config_context' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
		overrides.CurrentContext = cfgContext
	}
	if cfgContext, ok := configure_os.LookupEnv("KUBE_CTX"); ok && cfgContext != "" {
		overrides.CurrentContext = cfgContext
	}

	overrides.Context = configure_apiclientcmdapi.Context{}

	var cfgCtxCluster string
	if !providerConfig["config_context_cluster"].IsNull() && providerConfig["config_context_cluster"].IsKnown() {
		err = providerConfig["config_context_cluster"].As(&cfgCtxCluster)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to assert type of 'config_context_cluster' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
		overrides.Context.Cluster = cfgCtxCluster
	}
	if cfgCtxCluster, ok := configure_os.LookupEnv("KUBE_CTX_CLUSTER"); ok && cfgCtxCluster != "" {
		overrides.Context.Cluster = cfgCtxCluster
	}

	var cfgContextAuthInfo *string
	if !providerConfig["config_context_user"].IsNull() && providerConfig["config_context_user"].IsKnown() {
		err = providerConfig["config_context_user"].As(&cfgContextAuthInfo)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to assert type of 'config_context_user' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
		if cfgContextAuthInfo != nil {
			overrides.Context.AuthInfo = *cfgContextAuthInfo
		}
	}
	if cfgContextAuthInfoEnv, ok := configure_os.LookupEnv("KUBE_CTX_AUTH_INFO"); ok && cfgContextAuthInfoEnv != "" {
		overrides.Context.AuthInfo = cfgContextAuthInfoEnv
	}

	var username string
	if !providerConfig["username"].IsNull() && providerConfig["username"].IsKnown() {
		err = providerConfig["username"].As(&username)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to assert type of 'username' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
		overrides.AuthInfo.Username = username
	}
	if username, ok := configure_os.LookupEnv("KUBE_USERNAME"); ok && username != "" {
		overrides.AuthInfo.Username = username
	}

	var password string
	if !providerConfig["password"].IsNull() && providerConfig["password"].IsKnown() {
		err = providerConfig["password"].As(&password)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to assert type of 'password' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
		overrides.AuthInfo.Password = password
	}
	if password, ok := configure_os.LookupEnv("KUBE_PASSWORD"); ok && password != "" {
		overrides.AuthInfo.Password = password
	}

	var token string
	if !providerConfig["token"].IsNull() && providerConfig["token"].IsKnown() {
		err = providerConfig["token"].As(&token)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to assert type of 'token' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
		overrides.AuthInfo.Token = token
	}
	if token, ok := configure_os.LookupEnv("KUBE_TOKEN"); ok && token != "" {
		overrides.AuthInfo.Token = token
	}

	var proxyURL string
	if !providerConfig["proxy_url"].IsNull() && providerConfig["proxy_url"].IsKnown() {
		err = providerConfig["proxy_url"].As(&proxyURL)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to assert type of 'proxy_url' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
		overrides.ClusterDefaults.ProxyURL = proxyURL
	}
	if proxyUrl, ok := configure_os.LookupEnv("KUBE_PROXY_URL"); ok && proxyUrl != "" {
		overrides.ClusterDefaults.ProxyURL = proxyURL
	}

	if !providerConfig["exec"].IsNull() && providerConfig["exec"].IsKnown() {
		var execBlock []configure_tftypes.Value
		err = providerConfig["exec"].As(&execBlock)
		if err != nil {

			response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
				Severity: configure_tfprotov5.DiagnosticSeverityError,
				Summary:  "Provider configuration: failed to assert type of 'exec' value",
				Detail:   err.Error(),
			})
			return response, nil
		}
		execCfg := configure_apiclientcmdapi.ExecConfig{}
		execCfg.InteractiveMode = configure_apiclientcmdapi.IfAvailableExecInteractiveMode
		if len(execBlock) > 0 {
			var execObj map[string]configure_tftypes.Value
			err := execBlock[0].As(&execObj)
			if err != nil {
				response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
					Severity: configure_tfprotov5.DiagnosticSeverityError,
					Summary:  `Provider configuration: failed to assert type of "exec" block`,
					Detail:   err.Error(),
				})
				return response, nil
			}
			if !execObj["api_version"].IsNull() && execObj["api_version"].IsKnown() {
				var apiv string
				err = execObj["api_version"].As(&apiv)
				if err != nil {

					response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
						Severity: configure_tfprotov5.DiagnosticSeverityError,
						Summary:  "Provider configuration: failed to assert type of 'api_version' value",
						Detail:   err.Error(),
					})
					return response, nil
				}
				execCfg.APIVersion = apiv
			}
			if !execObj["command"].IsNull() && execObj["command"].IsKnown() {
				var cmd string
				err = execObj["command"].As(&cmd)
				if err != nil {

					response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
						Severity: configure_tfprotov5.DiagnosticSeverityError,
						Summary:  "Provider configuration: failed to assert type of 'command' value",
						Detail:   err.Error(),
					})
					return response, nil
				}
				execCfg.Command = cmd
			}
			if !execObj["args"].IsNull() && execObj["args"].IsFullyKnown() {
				var xcmdArgs []configure_tftypes.Value
				err = execObj["args"].As(&xcmdArgs)
				if err != nil {

					response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
						Severity: configure_tfprotov5.DiagnosticSeverityError,
						Summary:  "Provider configuration: failed to assert type of 'args' value",
						Detail:   err.Error(),
					})
					return response, nil
				}
				execCfg.Args = make([]string, 0, len(xcmdArgs))
				for _, arg := range xcmdArgs {
					var v string
					err := arg.As(&v)
					if err != nil {

						response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
							Severity: configure_tfprotov5.DiagnosticSeverityError,
							Summary:  "Provider configuration: failed to assert type of element in 'args' value",
							Detail:   err.Error(),
						})
						return response, nil
					}
					execCfg.Args = append(execCfg.Args, v)
				}
			}
			if !execObj["env"].IsNull() && execObj["env"].IsFullyKnown() {
				var xcmdEnvs map[string]configure_tftypes.Value
				err = execObj["env"].As(&xcmdEnvs)
				if err != nil {

					response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
						Severity: configure_tfprotov5.DiagnosticSeverityError,
						Summary:  "Provider configuration: failed to assert type of element in 'env' value",
						Detail:   err.Error(),
					})
					return response, nil
				}
				execCfg.Env = make([]configure_apiclientcmdapi.ExecEnvVar, 0, len(xcmdEnvs))
				for k, v := range xcmdEnvs {
					var vs string
					err = v.As(&vs)
					if err != nil {

						response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
							Severity: configure_tfprotov5.DiagnosticSeverityError,
							Summary:  "Provider configuration: failed to assert type of element in 'env' value",
							Detail:   err.Error(),
						})
						return response, nil
					}
					execCfg.Env = append(execCfg.Env, configure_apiclientcmdapi.ExecEnvVar{
						Name:  k,
						Value: vs,
					})
				}
			}
			overrides.AuthInfo.Exec = &execCfg
		}
	}

	cc := configure_clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, overrides)
	clientConfig, err := cc.ClientConfig()
	if err != nil {
		s.logger.Error("[Configure]", "Failed to load config:", dump(cc))
		if configure_errors.Is(err, configure_clientcmd.ErrEmptyConfig) {

			return response, nil
		}
		response.Diagnostics = append(response.Diagnostics, &configure_tfprotov5.Diagnostic{
			Severity: configure_tfprotov5.DiagnosticSeverityError,
			Summary:  "Provider configuration: cannot load Kubernetes client config",
			Detail:   err.Error(),
		})
		return response, nil
	}

	if s.logger.IsTrace() {
		clientConfig.WrapTransport = loggingTransport
	}

	codec := configure_runtime.NoopEncoder{Decoder: configure_scheme.Codecs.UniversalDecoder()}
	clientConfig.NegotiatedSerializer = configure_serializer.NegotiatedSerializerWrapper(configure_runtime.SerializerInfo{Serializer: codec})

	s.logger.Trace("[Configure]", "[ClientConfig]", dump(*clientConfig))
	s.clientConfig = clientConfig

	return response, nil
}

func (s *RawProviderServer) canExecute() (resp []*configure_tfprotov5.Diagnostic) {
	if !s.providerEnabled {
		resp = append(resp, &configure_tfprotov5.Diagnostic{
			Severity: configure_tfprotov5.DiagnosticSeverityError,
			Summary:  "Experimental feature not enabled.",
			Detail:   "The `kubernetes_manifest` resource is an experimental feature and must be explicitly enabled in the provider configuration block.",
		})
	}
	if configure_semver.IsValid(s.hostTFVersion) && configure_semver.Compare(s.hostTFVersion, minTFVersion) < 0 {
		resp = append(resp, &configure_tfprotov5.Diagnostic{
			Severity: configure_tfprotov5.DiagnosticSeverityError,
			Summary:  "Incompatible terraform version",
			Detail:   configure_fmt.Sprintf("The `kubernetes_manifest` resource requires Terraform %s or above", minTFVersion),
		})
	}
	return
}

func (s *RawProviderServer) ReadDataSource(ctx datasource_context.Context, req *datasource_tfprotov5.ReadDataSourceRequest) (*datasource_tfprotov5.ReadDataSourceResponse, error) {
	s.logger.Trace("[ReadDataSource][Request]\n%s\n", dump(*req))

	resp := &datasource_tfprotov5.ReadDataSourceResponse{}

	execDiag := s.canExecute()
	if len(execDiag) > 0 {
		resp.Diagnostics = append(resp.Diagnostics, execDiag...)
		return resp, nil
	}

	rt, err := GetDataSourceType(req.TypeName)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &datasource_tfprotov5.Diagnostic{
			Severity: datasource_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to determine data source type",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	config, err := req.Config.Unmarshal(rt)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &datasource_tfprotov5.Diagnostic{
			Severity: datasource_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to unmarshal data source configuration",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	var dsConfig map[string]datasource_tftypes.Value
	err = config.As(&dsConfig)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &datasource_tfprotov5.Diagnostic{
			Severity: datasource_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to extract attributes from data source configuration",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	rm, err := s.getRestMapper()
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &datasource_tfprotov5.Diagnostic{
			Severity: datasource_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to get RESTMapper client",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	client, err := s.getDynamicClient()
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &datasource_tfprotov5.Diagnostic{
			Severity: datasource_tfprotov5.DiagnosticSeverityError,
			Summary:  "failed to get Dynamic client",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	var apiVersion, kind string
	dsConfig["api_version"].As(&apiVersion)
	dsConfig["kind"].As(&kind)

	gvr, err := getGVR(apiVersion, kind, rm)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &datasource_tfprotov5.Diagnostic{
			Severity: datasource_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to determine resource GroupVersion",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	gvk := gvr.GroupVersion().WithKind(kind)
	ns, err := IsResourceNamespaced(gvk, rm)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &datasource_tfprotov5.Diagnostic{
			Severity: datasource_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed determine if resource is namespaced",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	rcl := client.Resource(gvr)

	objectType, th, err := s.TFTypeFromOpenAPI(ctx, gvk, false)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &datasource_tfprotov5.Diagnostic{
			Severity: datasource_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to save resource state",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	var metadataBlock []datasource_tftypes.Value
	dsConfig["metadata"].As(&metadataBlock)

	var metadata map[string]datasource_tftypes.Value
	metadataBlock[0].As(&metadata)

	var name string
	metadata["name"].As(&name)

	var res *datasource_unstructured.Unstructured
	if ns {
		var namespace string
		metadata["namespace"].As(&namespace)
		if namespace == "" {
			namespace = "default"
		}
		res, err = rcl.Namespace(namespace).Get(ctx, name, datasource_v1metav1.GetOptions{})
	} else {
		res, err = rcl.Get(ctx, name, datasource_v1metav1.GetOptions{})
	}
	if err != nil {
		if datasource_errorsapierrors.IsNotFound(err) {
			return resp, nil
		}
		d := datasource_tfprotov5.Diagnostic{
			Severity: datasource_tfprotov5.DiagnosticSeverityError,
			Summary:  datasource_fmt.Sprintf("Failed to get data source"),
			Detail:   err.Error(),
		}
		resp.Diagnostics = append(resp.Diagnostics, &d)
		return resp, nil
	}

	fo := RemoveServerSideFields(res.Object)
	nobj, err := datasource_payload.ToTFValue(fo, objectType, th, datasource_tftypes.NewAttributePath())
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &datasource_tfprotov5.Diagnostic{
			Severity: datasource_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to convert API response to Terraform value type",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	nobj, err = datasource_morph.DeepUnknown(objectType, nobj, datasource_tftypes.NewAttributePath())
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &datasource_tfprotov5.Diagnostic{
			Severity: datasource_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to save resource state",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	rawState := make(map[string]datasource_tftypes.Value)
	err = config.As(&rawState)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &datasource_tfprotov5.Diagnostic{
			Severity: datasource_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to save resource state",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	rawState["object"] = datasource_morph.UnknownToNull(nobj)

	v := datasource_tftypes.NewValue(rt, rawState)
	state, err := datasource_tfprotov5.NewDynamicValue(v.Type(), v)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &datasource_tfprotov5.Diagnostic{
			Severity: datasource_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to save resource state",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	resp.State = &state
	return resp, nil
}

func getGVR(apiVersion, kind string, m datasource_meta.RESTMapper) (datasource_schema.GroupVersionResource, error) {
	gv, err := datasource_schema.ParseGroupVersion(apiVersion)
	if err != nil {
		return datasource_schema.GroupVersionResource{}, err
	}
	mapping, err := m.RESTMapping(gv.WithKind(kind).GroupKind(), gv.Version)
	if err != nil {
		return datasource_schema.GroupVersionResource{}, err
	}
	return mapping.Resource, err
}

func APIStatusErrorToDiagnostics(s diagnostics_v1metav1.Status) []*diagnostics_tfprotov5.Diagnostic {
	var diags []*diagnostics_tfprotov5.Diagnostic
	diags = append(diags, &diagnostics_tfprotov5.Diagnostic{
		Severity: diagnostics_tfprotov5.DiagnosticSeverityError,
		Summary:  "API response status: " + s.Status,
		Detail:   s.Message,
	})
	if s.Details == nil {
		return diags
	}
	gk := diagnostics_v1metav1.GroupKind{Group: s.Details.Group, Kind: s.Details.Kind}
	diags = append(diags, &diagnostics_tfprotov5.Diagnostic{
		Severity: diagnostics_tfprotov5.DiagnosticSeverityError,
		Summary:  diagnostics_fmt.Sprintf("Kubernetes API Error: %s %s [%s]", string(s.Reason), gk.String(), s.Details.Name),
	})
	for _, c := range s.Details.Causes {
		diags = append(diags, &diagnostics_tfprotov5.Diagnostic{
			Severity: diagnostics_tfprotov5.DiagnosticSeverityError,
			Detail:   c.Message,
			Summary:  c.Field,
		})
	}
	return diags
}

func (s *RawProviderServer) GetProviderSchema(ctx getproviderschema_context.Context, req *getproviderschema_tfprotov5.GetProviderSchemaRequest) (*getproviderschema_tfprotov5.GetProviderSchemaResponse, error) {
	cfgSchema := GetProviderConfigSchema()
	resSchema := GetProviderResourceSchema()
	dsSchema := GetProviderDataSourceSchema()

	return &getproviderschema_tfprotov5.GetProviderSchemaResponse{
		Provider:          cfgSchema,
		ResourceSchemas:   resSchema,
		DataSourceSchemas: dsSchema,
	}, nil
}

func (s *RawProviderServer) ImportResourceState(ctx import_context.Context, req *import_tfprotov5.ImportResourceStateRequest) (*import_tfprotov5.ImportResourceStateResponse, error) {

	resp := &import_tfprotov5.ImportResourceStateResponse{}

	execDiag := s.canExecute()
	if len(execDiag) > 0 {
		resp.Diagnostics = append(resp.Diagnostics, execDiag...)
		return resp, nil
	}

	gvk, name, namespace, err := import_util.ParseResourceID(req.ID)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &import_tfprotov5.Diagnostic{
			Severity: import_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to parse import ID",
			Detail:   err.Error(),
		})
	}
	s.logger.Trace("[ImportResourceState]", "[ID]", gvk, name, namespace)
	rt, err := GetResourceType(req.TypeName)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &import_tfprotov5.Diagnostic{
			Severity: import_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to determine resource type",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	rm, err := s.getRestMapper()
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &import_tfprotov5.Diagnostic{
			Severity: import_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to get RESTMapper client",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	client, err := s.getDynamicClient()
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &import_tfprotov5.Diagnostic{
			Severity: import_tfprotov5.DiagnosticSeverityError,
			Summary:  "failed to get Dynamic client",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	ns, err := IsResourceNamespaced(gvk, rm)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &import_tfprotov5.Diagnostic{
			Severity: import_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to get namespacing requirement from RESTMapper",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	io := import_unstructured.Unstructured{}
	io.SetKind(gvk.Kind)
	io.SetAPIVersion(gvk.GroupVersion().String())
	io.SetName(name)
	io.SetNamespace(namespace)

	gvr, err := GVRFromUnstructured(&io, rm)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &import_tfprotov5.Diagnostic{
			Severity: import_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to get GVR from GVK via RESTMapper",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	rcl := client.Resource(gvr)

	var ro *import_unstructured.Unstructured
	if ns {
		ro, err = rcl.Namespace(namespace).Get(ctx, name, import_v1metav1.GetOptions{})
	} else {
		ro, err = rcl.Get(ctx, name, import_v1metav1.GetOptions{})
	}
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &import_tfprotov5.Diagnostic{
			Severity: import_tfprotov5.DiagnosticSeverityError,
			Summary:  import_fmt.Sprintf("Failed to get resource %+v from API", io),
			Detail:   err.Error(),
		})
		return resp, nil
	}
	s.logger.Trace("[ImportResourceState]", "[API Resource]", ro)

	objectType, th, err := s.TFTypeFromOpenAPI(ctx, gvk, false)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &import_tfprotov5.Diagnostic{
			Severity: import_tfprotov5.DiagnosticSeverityError,
			Summary:  import_fmt.Sprintf("Failed to determine resource type from GVK: %s", gvk),
			Detail:   err.Error(),
		})
		return resp, nil
	}

	fo := RemoveServerSideFields(ro.UnstructuredContent())
	nobj, err := import_payload.ToTFValue(fo, objectType, th, import_tftypes.NewAttributePath())
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &import_tfprotov5.Diagnostic{
			Severity: import_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to convert unstructured to tftypes.Value",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	nobj, err = import_morph.DeepUnknown(objectType, nobj, import_tftypes.NewAttributePath())
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &import_tfprotov5.Diagnostic{
			Severity: import_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to backfill unknown values during import",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	s.logger.Trace("[ImportResourceState]", "[tftypes.Value]", nobj)

	newState := make(map[string]import_tftypes.Value)
	wftype := rt.(import_tftypes.Object).AttributeTypes["wait_for"]
	wtype := rt.(import_tftypes.Object).AttributeTypes["wait"]
	timeoutsType := rt.(import_tftypes.Object).AttributeTypes["timeouts"]
	fmType := rt.(import_tftypes.Object).AttributeTypes["field_manager"]
	cmpType := rt.(import_tftypes.Object).AttributeTypes["computed_fields"]

	newState["manifest"] = import_tftypes.NewValue(import_tftypes.Object{AttributeTypes: map[string]import_tftypes.Type{}}, nil)
	newState["object"] = import_morph.UnknownToNull(nobj)
	newState["wait_for"] = import_tftypes.NewValue(wftype, nil)
	newState["wait"] = import_tftypes.NewValue(wtype, nil)
	newState["timeouts"] = import_tftypes.NewValue(timeoutsType, nil)
	newState["field_manager"] = import_tftypes.NewValue(fmType, nil)
	newState["computed_fields"] = import_tftypes.NewValue(cmpType, nil)

	nsVal := import_tftypes.NewValue(rt, newState)

	impState, err := import_tfprotov5.NewDynamicValue(nsVal.Type(), nsVal)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &import_tfprotov5.Diagnostic{
			Severity: import_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to construct dynamic value for imported state",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	impf := import_tftypes.NewValue(privateStateSchema,
		map[string]import_tftypes.Value{"IsImported": import_tftypes.NewValue(import_tftypes.Bool, true)},
	)
	fb, err := impf.MarshalMsgPack(privateStateSchema)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &import_tfprotov5.Diagnostic{
			Severity: import_tfprotov5.DiagnosticSeverityWarning,
			Summary:  "Failed to earmark imported resource",
			Detail:   err.Error(),
		})
	}
	nr := &import_tfprotov5.ImportedResource{
		TypeName: req.TypeName,
		State:    &impState,
		Private:  fb,
	}
	resp.ImportedResources = append(resp.ImportedResources, nr)
	resp.Diagnostics = append(resp.Diagnostics, &import_tfprotov5.Diagnostic{
		Severity: import_tfprotov5.DiagnosticSeverityWarning,
		Summary:  "Apply needed after 'import'",
		Detail:   "Please run apply after a successful import to realign the resource state to the configuration in Terraform.",
	})
	return resp, nil
}

func (s *RawProviderServer) dryRun(ctx plan_context.Context, obj plan_tftypes.Value, fieldManager string, forceConflicts bool, isNamespaced bool) error {
	c, err := s.getDynamicClient()
	if err != nil {
		return plan_fmt.Errorf("failed to retrieve Kubernetes dynamic client during apply: %v", err)
	}
	m, err := s.getRestMapper()
	if err != nil {
		return plan_fmt.Errorf("failed to retrieve Kubernetes RESTMapper client during apply: %v", err)
	}

	minObj := plan_morph.UnknownToNull(obj)
	pu, err := plan_payload.FromTFValue(minObj, nil, plan_tftypes.NewAttributePath())
	if err != nil {
		return err
	}

	rqObj := mapRemoveNulls(pu.(map[string]interface{}))
	uo := plan_unstructured.Unstructured{}
	uo.SetUnstructuredContent(rqObj)
	rnamespace := uo.GetNamespace()
	rname := uo.GetName()
	rnn := plan_types.NamespacedName{Namespace: rnamespace, Name: rname}.String()

	gvr, err := GVRFromUnstructured(&uo, m)
	if err != nil {
		return plan_fmt.Errorf("failed to determine resource GVR: %s", err)
	}

	var rs plan_dynamic.ResourceInterface
	if isNamespaced {
		rs = c.Resource(gvr).Namespace(rnamespace)
	} else {
		rs = c.Resource(gvr)
	}

	jsonManifest, err := uo.MarshalJSON()
	if err != nil {
		return plan_fmt.Errorf("failed to marshall resource %q to JSON: %v", rnn, err)
	}
	_, err = rs.Patch(ctx, rname, plan_types.ApplyPatchType, jsonManifest,
		plan_v1metav1.PatchOptions{
			FieldManager: fieldManager,
			Force:        &forceConflicts,
			DryRun:       []string{"All"},
		},
	)

	return err
}

const defaultFieldManagerName = "Terraform"

func (s *RawProviderServer) getFieldManagerConfig(v map[string]plan_tftypes.Value) (string, bool, error) {
	fieldManagerName := defaultFieldManagerName
	forceConflicts := false
	if !v["field_manager"].IsNull() && v["field_manager"].IsKnown() {
		var fieldManagerBlock []plan_tftypes.Value
		err := v["field_manager"].As(&fieldManagerBlock)
		if err != nil {
			return "", false, err
		}
		if len(fieldManagerBlock) > 0 {
			var fieldManagerObj map[string]plan_tftypes.Value
			err := fieldManagerBlock[0].As(&fieldManagerObj)
			if err != nil {
				return "", false, err
			}
			if !fieldManagerObj["name"].IsNull() && fieldManagerObj["name"].IsKnown() {
				err = fieldManagerObj["name"].As(&fieldManagerName)
				if err != nil {
					return "", false, err
				}
			}
			if !fieldManagerObj["force_conflicts"].IsNull() && fieldManagerObj["force_conflicts"].IsKnown() {
				err = fieldManagerObj["force_conflicts"].As(&forceConflicts)
				if err != nil {
					return "", false, err
				}
			}
		}
	}
	return fieldManagerName, forceConflicts, nil
}

func isImportedFlagFromPrivate(p []byte) (f bool, d []*plan_tfprotov5.Diagnostic) {
	if p == nil || len(p) == 0 {
		return
	}
	ps, err := getPrivateStateValue(p)
	if err != nil {
		d = append(d, &plan_tfprotov5.Diagnostic{
			Severity: plan_tfprotov5.DiagnosticSeverityError,
			Summary:  "Unexpected format for private state",
			Detail:   err.Error(),
		})
	}
	err = ps["IsImported"].As(&f)
	if err != nil {
		d = append(d, &plan_tfprotov5.Diagnostic{
			Severity: plan_tfprotov5.DiagnosticSeverityError,
			Summary:  "Unexpected format for import flag in private state",
			Detail:   err.Error(),
		})
	}
	return
}

func (s *RawProviderServer) PlanResourceChange(ctx plan_context.Context, req *plan_tfprotov5.PlanResourceChangeRequest) (*plan_tfprotov5.PlanResourceChangeResponse, error) {
	resp := &plan_tfprotov5.PlanResourceChangeResponse{}

	isImported, d := isImportedFlagFromPrivate(req.PriorPrivate)
	resp.Diagnostics = append(resp.Diagnostics, d...)
	if !isImported {
		resp.RequiresReplace = append(resp.RequiresReplace,
			plan_tftypes.NewAttributePath().WithAttributeName("manifest").WithAttributeName("apiVersion"),
			plan_tftypes.NewAttributePath().WithAttributeName("manifest").WithAttributeName("kind"),
			plan_tftypes.NewAttributePath().WithAttributeName("manifest").WithAttributeName("metadata").WithAttributeName("name"),
		)
	} else {
		resp.PlannedPrivate = req.PriorPrivate
	}

	execDiag := s.canExecute()
	if len(execDiag) > 0 {
		resp.Diagnostics = append(resp.Diagnostics, execDiag...)
		return resp, nil
	}

	resp.Diagnostics = append(resp.Diagnostics, s.checkValidCredentials(ctx)...)
	if len(resp.Diagnostics) > 0 {
		return resp, nil
	}

	rt, err := GetResourceType(req.TypeName)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
			Severity: plan_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to determine planned resource type",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	proposedState, err := req.ProposedNewState.Unmarshal(rt)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
			Severity: plan_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to unmarshal planned resource state",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	s.logger.Trace("[PlanResourceChange]", "[ProposedState]", dump(proposedState))

	proposedVal := make(map[string]plan_tftypes.Value)
	err = proposedState.As(&proposedVal)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
			Severity: plan_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to extract planned resource state from tftypes.Value",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	computedFields := make(map[string]*plan_tftypes.AttributePath)
	var atp *plan_tftypes.AttributePath
	cfVal, ok := proposedVal["computed_fields"]
	if ok && !cfVal.IsNull() && cfVal.IsKnown() {
		var cf []plan_tftypes.Value
		cfVal.As(&cf)
		for _, v := range cf {
			var vs string
			err := v.As(&vs)
			if err != nil {
				s.logger.Error("[computed_fields] cannot extract element from list")
				continue
			}
			atp, err := FieldPathToTftypesPath(vs)
			if err != nil {
				s.logger.Error("[Configure]", "[computed_fields] cannot parse filed path element", err)
				resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
					Severity: plan_tfprotov5.DiagnosticSeverityError,
					Summary:  "[computed_fields] cannot parse field path element: " + vs,
					Detail:   err.Error(),
				})
				continue
			}
			computedFields[atp.String()] = atp
		}
	} else {

		atp = plan_tftypes.NewAttributePath().WithAttributeName("metadata").WithAttributeName("annotations")
		computedFields[atp.String()] = atp

		atp = plan_tftypes.NewAttributePath().WithAttributeName("metadata").WithAttributeName("labels")
		computedFields[atp.String()] = atp
	}

	priorState, err := req.PriorState.Unmarshal(rt)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
			Severity: plan_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to unmarshal prior resource state",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	s.logger.Trace("[PlanResourceChange]", "[PriorState]", dump(priorState))

	priorVal := make(map[string]plan_tftypes.Value)
	err = priorState.As(&priorVal)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
			Severity: plan_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to extract prior resource state from tftypes.Value",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	if proposedState.IsNull() {

		if _, ok := priorVal["object"]; ok {
			resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
				Severity: plan_tfprotov5.DiagnosticSeverityError,
				Summary:  "Invalid prior state while planning for destroy",
				Detail:   plan_fmt.Sprintf("'object' attribute missing from state: %s", err),
			})
			return resp, nil
		}
		resp.PlannedState = req.ProposedNewState
		return resp, nil
	}

	ppMan, ok := proposedVal["manifest"]
	if !ok {
		resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
			Severity:  plan_tfprotov5.DiagnosticSeverityError,
			Summary:   "Invalid proposed state during planning",
			Detail:    "Missing 'manifest' attribute",
			Attribute: plan_tftypes.NewAttributePath().WithAttributeName("manifest"),
		})
		return resp, nil
	}

	rm, err := s.getRestMapper()
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
			Severity: plan_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to create K8s RESTMapper client",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	gvk, err := GVKFromTftypesObject(&ppMan, rm)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
			Severity: plan_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to determine GroupVersionResource for manifest",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	vdiags := s.validateResourceOnline(&ppMan)
	if len(vdiags) > 0 {
		resp.Diagnostics = append(resp.Diagnostics, vdiags...)
		return resp, nil
	}

	ns, err := IsResourceNamespaced(gvk, rm)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
			Severity: plan_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to discover scope of resource",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	if ns && !isImported {
		resp.RequiresReplace = append(resp.RequiresReplace,
			plan_tftypes.NewAttributePath().WithAttributeName("manifest").WithAttributeName("metadata").WithAttributeName("namespace"),
		)
	}

	objectType, hints, err := s.TFTypeFromOpenAPI(ctx, gvk, false)
	if err != nil {
		return resp, plan_fmt.Errorf("failed to determine resource type ID: %s", err)
	}

	if !objectType.Is(plan_tftypes.Object{}) {

		objectType = ppMan.Type()

		resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
			Severity: plan_tfprotov5.DiagnosticSeverityWarning,
			Summary:  "This custom resource does not have an associated OpenAPI schema.",
			Detail:   "We could not find an OpenAPI schema for this custom resource. Updates to this resource will cause a forced replacement.",
		})

		fieldManagerName, forceConflicts, err := s.getFieldManagerConfig(proposedVal)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
				Severity: plan_tfprotov5.DiagnosticSeverityError,
				Summary:  "Could not extract field_manager config",
				Detail:   err.Error(),
			})
			return resp, nil
		}

		err = s.dryRun(ctx, ppMan, fieldManagerName, forceConflicts, ns)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
				Severity: plan_tfprotov5.DiagnosticSeverityError,
				Summary:  "Dry-run failed for non-structured resource",
				Detail:   plan_fmt.Sprintf("A dry-run apply was performed for this resource but was unsuccessful: %v", err),
			})
			return resp, nil
		}

		resp.RequiresReplace = []*plan_tftypes.AttributePath{
			plan_tftypes.NewAttributePath().WithAttributeName("manifest"),
			plan_tftypes.NewAttributePath().WithAttributeName("object"),
		}
	}

	so := objectType.(plan_tftypes.Object)
	s.logger.Debug("[PlanUpdateResource]", "OAPI type", dump(so))

	morphedManifest, d := plan_morph.ValueToType(ppMan, objectType, plan_tftypes.NewAttributePath().WithAttributeName("object"))
	if len(d) > 0 {
		resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
			Severity: plan_tfprotov5.DiagnosticSeverityError,
			Summary:  "Manifest configuration incompatible with resource schema",
			Detail:   "Detailed descriptions of errors will follow below.",
		})
		resp.Diagnostics = append(resp.Diagnostics, d...)
		return resp, nil
	}
	s.logger.Debug("[PlanResourceChange]", "morphed manifest", dump(morphedManifest))

	completePropMan, err := plan_morph.DeepUnknown(objectType, morphedManifest, plan_tftypes.NewAttributePath().WithAttributeName("object"))
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
			Severity:  plan_tfprotov5.DiagnosticSeverityError,
			Summary:   "Failed to backfill manifest from OpenAPI type",
			Detail:    plan_fmt.Sprintf("This usually happens when the provider cannot fully process the schema retrieved from cluster. Please report this to the provider maintainers.\nError: %s", err.Error()),
			Attribute: plan_tftypes.NewAttributePath().WithAttributeName("object"),
		})
		return resp, nil
	}
	s.logger.Debug("[PlanResourceChange]", "backfilled manifest", dump(completePropMan))

	if proposedVal["object"].IsNull() {

		s.logger.Debug("[PlanResourceChange]", "creating object", dump(completePropMan))
		newObj, err := plan_tftypes.Transform(completePropMan, func(ap *plan_tftypes.AttributePath, v plan_tftypes.Value) (plan_tftypes.Value, error) {
			_, ok := computedFields[ap.String()]
			if ok {
				return plan_tftypes.NewValue(v.Type(), plan_tftypes.UnknownValue), nil
			}
			return v, nil
		})
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
				Severity:  plan_tfprotov5.DiagnosticSeverityError,
				Summary:   "Failed to set computed attributes in new resource state",
				Detail:    err.Error(),
				Attribute: plan_tftypes.NewAttributePath().WithAttributeName("object"),
			})
			return resp, nil
		}
		proposedVal["object"] = newObj
	} else {

		priorObj, ok := priorVal["object"]
		if !ok {
			resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
				Severity:  plan_tfprotov5.DiagnosticSeverityError,
				Summary:   "Invalid prior state during planning",
				Detail:    "Missing 'object' attribute",
				Attribute: plan_tftypes.NewAttributePath().WithAttributeName("object"),
			})
			return resp, nil
		}
		priorMan, ok := priorVal["manifest"]
		if !ok {
			resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
				Severity:  plan_tfprotov5.DiagnosticSeverityError,
				Summary:   "Invalid prior state during planning",
				Detail:    "Missing 'manifest' attribute",
				Attribute: plan_tftypes.NewAttributePath().WithAttributeName("manifest"),
			})
			return resp, nil
		}
		updatedObj, err := plan_tftypes.Transform(completePropMan, func(ap *plan_tftypes.AttributePath, v plan_tftypes.Value) (plan_tftypes.Value, error) {
			_, isComputed := computedFields[ap.String()]
			if v.IsKnown() {
				hasChanged := false
				wasCfg, restPath, err := plan_tftypes.WalkAttributePath(priorMan, ap)
				if err != nil && len(restPath.Steps()) != 0 {
					hasChanged = true
				}
				nowCfg, restPath, err := plan_tftypes.WalkAttributePath(ppMan, ap)
				hasChanged = err == nil && len(restPath.Steps()) == 0 && wasCfg.(plan_tftypes.Value).IsKnown() && !wasCfg.(plan_tftypes.Value).Equal(nowCfg.(plan_tftypes.Value))
				if hasChanged {
					h, ok := hints[plan_morph.ValueToTypePath(ap).String()]
					if ok && h == plan_manifest.PreserveUnknownFieldsLabel {
						apm := append(plan_tftypes.NewAttributePath().WithAttributeName("manifest").Steps(), ap.Steps()...)
						resp.RequiresReplace = append(resp.RequiresReplace, plan_tftypes.NewAttributePathWithSteps(apm))
					}
				}
				if isComputed {
					if hasChanged {
						return plan_tftypes.NewValue(v.Type(), plan_tftypes.UnknownValue), nil
					}
					nowVal, restPath, err := plan_tftypes.WalkAttributePath(proposedVal["object"], ap)
					if err == nil && len(restPath.Steps()) == 0 {
						return nowVal.(plan_tftypes.Value), nil
					}
				}
				return v, nil
			}

			wasVal, restPath, err := plan_tftypes.WalkAttributePath(priorMan, ap)
			if err == nil && len(restPath.Steps()) == 0 && wasVal.(plan_tftypes.Value).IsKnown() {

				return v, nil
			}

			priorAtrVal, restPath, err := plan_tftypes.WalkAttributePath(priorObj, ap)
			if err != nil {
				if len(restPath.Steps()) > 0 {

					return v, nil
				}

				return v, ap.NewError(err)
			}
			if len(restPath.Steps()) > 0 {
				s.logger.Warn("[PlanResourceChange]", "Unexpected missing attribute from state at", ap.String(), " + ", restPath.String())
			}
			return priorAtrVal.(plan_tftypes.Value), nil
		})
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
				Severity:  plan_tfprotov5.DiagnosticSeverityError,
				Summary:   "Failed to update proposed state from prior state",
				Detail:    err.Error(),
				Attribute: plan_tftypes.NewAttributePath().WithAttributeName("object"),
			})
			return resp, nil
		}

		proposedVal["object"] = updatedObj
	}

	propStateVal := plan_tftypes.NewValue(proposedState.Type(), proposedVal)
	s.logger.Trace("[PlanResourceChange]", "new planned state", dump(propStateVal))

	plannedState, err := plan_tfprotov5.NewDynamicValue(propStateVal.Type(), propStateVal)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &plan_tfprotov5.Diagnostic{
			Severity: plan_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to assemble proposed state during plan",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	resp.PlannedState = &plannedState
	return resp, nil
}

func getAttributeValue(v plan_tftypes.Value, path string) (plan_tftypes.Value, error) {
	p, err := FieldPathToTftypesPath(path)
	if err != nil {
		return plan_tftypes.Value{}, err
	}
	vv, _, err := plan_tftypes.WalkAttributePath(v, p)
	if err != nil {
		return plan_tftypes.Value{}, err
	}
	return vv.(plan_tftypes.Value), nil
}

var providerName = "registry.terraform.io/hashicorp/kubernetes"

func Serve(ctx plugin_context.Context, logger plugin_hclog.Logger) error {
	return plugin_tf5servertf5server.Serve(providerName, func() plugin_tfprotov5.ProviderServer { return &(RawProviderServer{logger: logger}) })
}

func Provider() func() plugin_tfprotov5.ProviderServer {
	var logLevel string
	var ok bool = false
	for _, ev := range []string{"TF_LOG_PROVIDER_KUBERNETES", "TF_LOG_PROVIDER", "TF_LOG"} {
		logLevel, ok = plugin_os.LookupEnv(ev)
		if ok {
			break
		}
	}
	if !ok {
		logLevel = "off"
	}

	return func() plugin_tfprotov5.ProviderServer {
		return &(RawProviderServer{logger: plugin_hclog.New(&plugin_hclog.LoggerOptions{
			Level:  plugin_hclog.LevelFromString(logLevel),
			Output: plugin_os.Stderr,
		})})
	}
}

func ServeTest(ctx plugin_context.Context, logger plugin_hclog.Logger, t *plugin_testing.T) (plugin_tfexec.ReattachInfo, error) {
	reattachConfigCh := make(chan *plugin_plugin.ReattachConfig)

	go plugin_tf5servertf5server.Serve(providerName,
		func() plugin_tfprotov5.ProviderServer { return &(RawProviderServer{logger: logger}) },
		plugin_tf5servertf5server.WithDebug(ctx, reattachConfigCh, nil),
		plugin_tf5servertf5server.WithLoggingSink(t),
		plugin_tf5servertf5server.WithGoPluginLogger(logger),
	)

	reattachConfig, err := waitForReattachConfig(reattachConfigCh)
	if err != nil {
		return nil, plugin_fmt.Errorf("Error getting reattach config: %s", err)
	}

	return map[string]plugin_tfexec.ReattachConfig{
		providerName: convertReattachConfig(reattachConfig),
	}, nil
}

func convertReattachConfig(reattachConfig *plugin_plugin.ReattachConfig) plugin_tfexec.ReattachConfig {
	return plugin_tfexec.ReattachConfig{
		Protocol: string(reattachConfig.Protocol),
		Pid:      reattachConfig.Pid,
		Test:     true,
		Addr: plugin_tfexec.ReattachConfigAddr{
			Network: reattachConfig.Addr.Network(),
			String:  reattachConfig.Addr.String(),
		},
	}
}

func printReattachConfig(config *plugin_plugin.ReattachConfig) {
	reattachStr, err := plugin_json.Marshal(map[string]plugin_tfexec.ReattachConfig{
		providerName: convertReattachConfig(config),
	})
	if err != nil {
		plugin_fmt.Printf("Error building reattach string: %s", err)
		return
	}
	plugin_fmt.Printf("# Provider server started\nexport TF_REATTACH_PROVIDERS='%s'\n", string(reattachStr))
}

func waitForReattachConfig(ch chan *plugin_plugin.ReattachConfig) (*plugin_plugin.ReattachConfig, error) {
	select {
	case config := <-ch:
		return config, nil
	case <-plugin_time.After(2 * plugin_time.Second):
		return nil, plugin_fmt.Errorf("timeout while waiting for reattach configuration")
	}
}

func GetObjectTypeFromSchema(schema *provider_tfprotov5.Schema) provider_tftypes.Type {
	bm := map[string]provider_tftypes.Type{}

	for _, att := range schema.Block.Attributes {
		bm[att.Name] = att.Type
	}

	for _, b := range schema.Block.BlockTypes {
		a := map[string]provider_tftypes.Type{}
		for _, att := range b.Block.Attributes {
			a[att.Name] = att.Type
		}
		bm[b.TypeName] = provider_tftypes.List{
			ElementType: provider_tftypes.Object{AttributeTypes: a},
		}

		for _, bb := range b.Block.BlockTypes {
			aa := map[string]provider_tftypes.Type{}
			for _, att := range bb.Block.Attributes {
				aa[att.Name] = att.Type
			}
			a[bb.TypeName] = provider_tftypes.List{
				ElementType: provider_tftypes.Object{AttributeTypes: aa},
			}
		}
	}

	return provider_tftypes.Object{AttributeTypes: bm}
}

func GetResourceType(name string) (provider_tftypes.Type, error) {
	sch := GetProviderResourceSchema()
	rsch, ok := sch[name]
	if !ok {
		return provider_tftypes.DynamicPseudoType, provider_fmt.Errorf("unknown resource %s - cannot find schema", name)
	}
	return GetObjectTypeFromSchema(rsch), nil
}

func GetDataSourceType(name string) (provider_tftypes.Type, error) {
	sch := GetProviderDataSourceSchema()
	rsch, ok := sch[name]
	if !ok {
		return provider_tftypes.DynamicPseudoType, provider_fmt.Errorf("unknown data source %q: cannot find schema", name)
	}
	return GetObjectTypeFromSchema(rsch), nil
}

func GetProviderResourceSchema() map[string]*provider_tfprotov5.Schema {
	return map[string]*provider_tfprotov5.Schema{
		"kubernetes_manifest": {
			Version: 1,
			Block: &provider_tfprotov5.SchemaBlock{
				BlockTypes: []*provider_tfprotov5.SchemaNestedBlock{
					{
						TypeName: "timeouts",
						Nesting:  provider_tfprotov5.SchemaNestedBlockNestingModeList,
						MinItems: 0,
						MaxItems: 1,
						Block: &provider_tfprotov5.SchemaBlock{
							Attributes: []*provider_tfprotov5.SchemaAttribute{
								{
									Name:        "create",
									Type:        provider_tftypes.String,
									Description: "Timeout for the create operation.",
									Optional:    true,
								},
								{
									Name:        "update",
									Type:        provider_tftypes.String,
									Description: "Timeout for the update operation.",
									Optional:    true,
								},
								{
									Name:        "delete",
									Type:        provider_tftypes.String,
									Description: "Timeout for the delete operation.",
									Optional:    true,
								},
							},
						},
					},
					{
						TypeName: "field_manager",
						Nesting:  provider_tfprotov5.SchemaNestedBlockNestingModeList,
						MinItems: 0,
						MaxItems: 1,
						Block: &provider_tfprotov5.SchemaBlock{
							Description: "Configure field manager options.",
							Attributes: []*provider_tfprotov5.SchemaAttribute{
								{
									Name:            "name",
									Type:            provider_tftypes.String,
									Required:        false,
									Optional:        true,
									Computed:        false,
									Sensitive:       false,
									Description:     "The name to use for the field manager when creating and updating the resource.",
									DescriptionKind: 0,
									Deprecated:      false,
								},
								{
									Name:            "force_conflicts",
									Type:            provider_tftypes.Bool,
									Required:        false,
									Optional:        true,
									Computed:        false,
									Sensitive:       false,
									Description:     "Force changes against conflicts.",
									DescriptionKind: 0,
									Deprecated:      false,
								},
							},
						},
					},
					{
						TypeName: "wait",
						Nesting:  provider_tfprotov5.SchemaNestedBlockNestingModeList,
						MinItems: 0,
						MaxItems: 1,
						Block: &provider_tfprotov5.SchemaBlock{
							Description: "Configure waiter options.",
							BlockTypes: []*provider_tfprotov5.SchemaNestedBlock{
								{
									TypeName: "condition",
									Nesting:  provider_tfprotov5.SchemaNestedBlockNestingModeList,
									MinItems: 0,
									Block: &provider_tfprotov5.SchemaBlock{
										Attributes: []*provider_tfprotov5.SchemaAttribute{
											{
												Name:        "status",
												Type:        provider_tftypes.String,
												Optional:    true,
												Description: "The condition status.",
											}, {
												Name:        "type",
												Type:        provider_tftypes.String,
												Optional:    true,
												Description: "The type of condition.",
											},
										},
									},
								},
							},
							Attributes: []*provider_tfprotov5.SchemaAttribute{
								{
									Name:        "rollout",
									Type:        provider_tftypes.Bool,
									Optional:    true,
									Description: "Wait for rollout to complete on resources that support `kubectl rollout status`.",
								},
								{
									Name:        "fields",
									Type:        provider_tftypes.Map{ElementType: provider_tftypes.String},
									Optional:    true,
									Description: "A map of paths to fields to wait for a specific field value.",
								},
							},
						},
					},
				},
				Attributes: []*provider_tfprotov5.SchemaAttribute{
					{
						Name:        "manifest",
						Type:        provider_tftypes.DynamicPseudoType,
						Required:    true,
						Description: "A Kubernetes manifest describing the desired state of the resource in HCL format.",
					},
					{
						Name:        "object",
						Type:        provider_tftypes.DynamicPseudoType,
						Optional:    true,
						Computed:    true,
						Description: "The resulting resource state, as returned by the API server after applying the desired state from `manifest`.",
					},
					{
						Name: "wait_for",
						Type: provider_tftypes.Object{
							AttributeTypes: map[string]provider_tftypes.Type{
								"fields": provider_tftypes.Map{
									ElementType: provider_tftypes.String,
								},
							},
						},
						Optional:    true,
						Deprecated:  true,
						Description: "A map of attribute paths and desired patterns to be matched. After each apply the provider will wait for all attributes listed here to reach a value that matches the desired pattern.",
					},
					{
						Name:        "computed_fields",
						Type:        provider_tftypes.List{ElementType: provider_tftypes.String},
						Description: "List of manifest fields whose values can be altered by the API server during 'apply'. Defaults to: [\"metadata.annotations\", \"metadata.labels\"]",
						Optional:    true,
					},
				},
			},
		},
	}
}

func GetProviderDataSourceSchema() map[string]*provider_tfprotov5.Schema {
	return map[string]*provider_tfprotov5.Schema{
		"kubernetes_resource": {
			Version: 1,
			Block: &provider_tfprotov5.SchemaBlock{
				Attributes: []*provider_tfprotov5.SchemaAttribute{
					{
						Name:        "api_version",
						Type:        provider_tftypes.String,
						Required:    true,
						Description: "The resource apiVersion.",
					},
					{
						Name:        "kind",
						Type:        provider_tftypes.String,
						Required:    true,
						Description: "The resource kind.",
					},
					{
						Name:        "object",
						Type:        provider_tftypes.DynamicPseudoType,
						Optional:    true,
						Computed:    true,
						Description: "The response from the API server.",
					},
				},
				BlockTypes: []*provider_tfprotov5.SchemaNestedBlock{
					{
						TypeName: "metadata",
						Nesting:  provider_tfprotov5.SchemaNestedBlockNestingModeList,
						MinItems: 1,
						MaxItems: 1,
						Block: &provider_tfprotov5.SchemaBlock{
							Description: "Metadata for the resource",
							Attributes: []*provider_tfprotov5.SchemaAttribute{
								{
									Name:        "name",
									Type:        provider_tftypes.String,
									Required:    true,
									Description: "The resource name.",
								},
								{
									Name:        "namespace",
									Type:        provider_tftypes.String,
									Optional:    true,
									Description: "The resource namespace.",
								},
							},
						},
					},
				},
			},
		},
	}
}

func GetProviderConfigSchema() *provider_config_tfprotov5.Schema {
	b := provider_config_tfprotov5.SchemaBlock{

		Attributes: []*provider_config_tfprotov5.SchemaAttribute{
			{
				Name:            "host",
				Type:            provider_config_tftypes.String,
				Description:     "The hostname (in form of URI) of Kubernetes master.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "username",
				Type:            provider_config_tftypes.String,
				Description:     "The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "password",
				Type:            provider_config_tftypes.String,
				Description:     "The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "insecure",
				Type:            provider_config_tftypes.Bool,
				Description:     "Whether server should be accessed without verifying the TLS certificate.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "client_certificate",
				Type:            provider_config_tftypes.String,
				Description:     "PEM-encoded client certificate for TLS authentication.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "client_key",
				Type:            provider_config_tftypes.String,
				Description:     "PEM-encoded client certificate key for TLS authentication.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "cluster_ca_certificate",
				Type:            provider_config_tftypes.String,
				Description:     "PEM-encoded root certificates bundle for TLS authentication.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "config_paths",
				Type:            provider_config_tftypes.List{ElementType: provider_config_tftypes.String},
				Description:     "A list of paths to kube config files. Can be set with KUBE_CONFIG_PATHS environment variable.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "config_path",
				Type:            provider_config_tftypes.String,
				Description:     "Path to the kube config file. Can be set with KUBE_CONFIG_PATH.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "config_context",
				Type:            provider_config_tftypes.String,
				Description:     "",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "config_context_auth_info",
				Type:            provider_config_tftypes.String,
				Description:     "",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "config_context_cluster",
				Type:            provider_config_tftypes.String,
				Description:     "",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "token",
				Type:            provider_config_tftypes.String,
				Description:     "Token to authenticate an service account",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "proxy_url",
				Type:            provider_config_tftypes.String,
				Description:     "URL to the proxy to be used for all API requests",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "ignore_annotations",
				Type:            provider_config_tftypes.List{ElementType: provider_config_tftypes.String},
				Description:     "List of Kubernetes metadata annotations to ignore across all resources handled by this provider for situations where external systems are managing certain resource annotations. Each item is a regular expression.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "ignore_labels",
				Type:            provider_config_tftypes.List{ElementType: provider_config_tftypes.String},
				Description:     "List of Kubernetes metadata labels to ignore across all resources handled by this provider for situations where external systems are managing certain resource labels. Each item is a regular expression.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
		},
		BlockTypes: []*provider_config_tfprotov5.SchemaNestedBlock{
			{
				TypeName: "exec",
				Nesting:  provider_config_tfprotov5.SchemaNestedBlockNestingModeList,
				MinItems: 0,
				MaxItems: 1,
				Block: &provider_config_tfprotov5.SchemaBlock{
					Attributes: []*provider_config_tfprotov5.SchemaAttribute{
						{
							Name:            "api_version",
							Type:            provider_config_tftypes.String,
							Required:        true,
							Optional:        false,
							Computed:        false,
							Sensitive:       false,
							DescriptionKind: 0,
							Deprecated:      false,
						},
						{
							Name:            "command",
							Type:            provider_config_tftypes.String,
							Required:        true,
							Optional:        false,
							Computed:        false,
							Sensitive:       false,
							DescriptionKind: 0,
							Deprecated:      false,
						},
						{
							Name:            "env",
							Type:            provider_config_tftypes.Map{ElementType: provider_config_tftypes.String},
							Required:        false,
							Optional:        true,
							Computed:        false,
							Sensitive:       false,
							DescriptionKind: 0,
							Deprecated:      false,
						},
						{
							Name:            "args",
							Type:            provider_config_tftypes.List{ElementType: provider_config_tftypes.String},
							Required:        false,
							Optional:        true,
							Computed:        false,
							Sensitive:       false,
							DescriptionKind: 0,
							Deprecated:      false,
						},
					},
				},
			},
			{
				TypeName: "experiments",
				Nesting:  provider_config_tfprotov5.SchemaNestedBlockNestingModeList,
				MinItems: 0,
				MaxItems: 1,
				Block: &provider_config_tfprotov5.SchemaBlock{
					Description: "Enable and disable experimental features.",
					Attributes: []*provider_config_tfprotov5.SchemaAttribute{
						{
							Name:            "manifest_resource",
							Type:            provider_config_tftypes.Bool,
							Required:        false,
							Optional:        true,
							Computed:        false,
							Sensitive:       false,
							Description:     "Enable the `kubernetes_manifest` resource.",
							DescriptionKind: 0,
							Deprecated:      false,
						},
					},
				},
			},
		},
	}

	return &provider_config_tfprotov5.Schema{
		Version: 0,
		Block:   &b,
	}
}

func (s *RawProviderServer) ReadResource(ctx read_context.Context, req *read_tfprotov5.ReadResourceRequest) (*read_tfprotov5.ReadResourceResponse, error) {
	resp := &read_tfprotov5.ReadResourceResponse{}

	resp.Private = req.Private

	execDiag := s.canExecute()
	if len(execDiag) > 0 {
		resp.Diagnostics = append(resp.Diagnostics, execDiag...)
		return resp, nil
	}

	var resState map[string]read_tftypes.Value
	var err error
	rt, err := GetResourceType(req.TypeName)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &read_tfprotov5.Diagnostic{
			Severity: read_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to determine resource type",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	currentState, err := req.CurrentState.Unmarshal(rt)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &read_tfprotov5.Diagnostic{
			Severity: read_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to decode current state",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	if currentState.IsNull() {
		resp.Diagnostics = append(resp.Diagnostics, &read_tfprotov5.Diagnostic{
			Severity: read_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to read resource",
			Detail:   "Incomplete of missing state",
		})
		return resp, nil
	}
	err = currentState.As(&resState)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &read_tfprotov5.Diagnostic{
			Severity: read_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to extract resource from current state",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	co, hasOb := resState["object"]
	if !hasOb || co.IsNull() {
		resp.Diagnostics = append(resp.Diagnostics, &read_tfprotov5.Diagnostic{
			Severity: read_tfprotov5.DiagnosticSeverityError,
			Summary:  "Current state of resource has no 'object' attribute",
			Detail:   "This should not happen. The state may be incomplete or corrupted.\nIf this error is reproducible, plese report issue to provider maintainers.",
		})
		return resp, nil
	}
	rm, err := s.getRestMapper()
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &read_tfprotov5.Diagnostic{
			Severity: read_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to get RESTMapper client",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	gvk, err := GVKFromTftypesObject(&co, rm)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &read_tfprotov5.Diagnostic{
			Severity: read_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to determine GroupVersionResource for manifest",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	objectType, th, err := s.TFTypeFromOpenAPI(ctx, gvk, false)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &read_tfprotov5.Diagnostic{
			Severity: read_tfprotov5.DiagnosticSeverityError,
			Summary:  read_fmt.Sprintf("Failed to determine resource type from GVK: %s", gvk),
			Detail:   err.Error(),
		})
		return resp, nil
	}

	cu, err := read_payload.FromTFValue(co, th, read_tftypes.NewAttributePath())
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &read_tfprotov5.Diagnostic{
			Severity: read_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed encode 'object' attribute to Unstructured",
			Detail:   err.Error(),
		})
		return resp, nil
	}
	s.logger.Trace("[ReadResource]", "[unstructured.FromTFValue]", dump(cu))

	client, err := s.getDynamicClient()
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &read_tfprotov5.Diagnostic{
			Severity: read_tfprotov5.DiagnosticSeverityError,
			Summary:  "failed to get Dynamic client",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	uo := read_unstructured.Unstructured{Object: cu.(map[string]interface{})}
	cGVR, err := GVRFromUnstructured(&uo, rm)
	if err != nil {
		return resp, err
	}
	ns, err := IsResourceNamespaced(uo.GroupVersionKind(), rm)
	if err != nil {
		return resp, err
	}
	rcl := client.Resource(cGVR)

	rnamespace := uo.GetNamespace()
	rname := uo.GetName()

	var ro *read_unstructured.Unstructured
	if ns {
		ro, err = rcl.Namespace(rnamespace).Get(ctx, rname, read_v1metav1.GetOptions{})
	} else {
		ro, err = rcl.Get(ctx, rname, read_v1metav1.GetOptions{})
	}
	if err != nil {
		if read_errorsapierrors.IsNotFound(err) {
			return resp, nil
		}
		d := read_tfprotov5.Diagnostic{
			Severity: read_tfprotov5.DiagnosticSeverityError,
			Summary:  read_fmt.Sprintf("Cannot GET resource %s", dump(co)),
			Detail:   err.Error(),
		}
		resp.Diagnostics = append(resp.Diagnostics, &d)
		return resp, nil
	}

	fo := RemoveServerSideFields(ro.Object)
	nobj, err := read_payload.ToTFValue(fo, objectType, th, read_tftypes.NewAttributePath())
	if err != nil {
		return resp, err
	}

	nobj, err = read_morph.DeepUnknown(objectType, nobj, read_tftypes.NewAttributePath())
	if err != nil {
		return resp, err
	}

	rawState := make(map[string]read_tftypes.Value)
	err = currentState.As(&rawState)
	if err != nil {
		return resp, err
	}
	rawState["object"] = read_morph.UnknownToNull(nobj)

	nsVal := read_tftypes.NewValue(currentState.Type(), rawState)
	newState, err := read_tfprotov5.NewDynamicValue(nsVal.Type(), nsVal)
	if err != nil {
		return resp, err
	}
	resp.NewState = &newState
	return resp, nil
}

func GVRFromUnstructured(o *resource_unstructured.Unstructured, m resource_meta.RESTMapper) (resource_schema.GroupVersionResource, error) {
	apv := o.GetAPIVersion()
	kind := o.GetKind()
	gv, err := resource_schema.ParseGroupVersion(apv)
	if err != nil {
		return resource_schema.GroupVersionResource{}, err
	}
	mapping, err := m.RESTMapping(gv.WithKind(kind).GroupKind(), gv.Version)
	if err != nil {
		return resource_schema.GroupVersionResource{}, err
	}
	return mapping.Resource, err
}

func GVKFromTftypesObject(in *resource_tftypes.Value, m resource_meta.RESTMapper) (resource_schema.GroupVersionKind, error) {
	var obj map[string]resource_tftypes.Value
	err := in.As(&obj)
	if err != nil {
		return resource_schema.GroupVersionKind{}, err
	}
	var apv string
	var kind string
	err = obj["apiVersion"].As(&apv)
	if err != nil {
		return resource_schema.GroupVersionKind{}, err
	}
	err = obj["kind"].As(&kind)
	if err != nil {
		return resource_schema.GroupVersionKind{}, err
	}
	gv, err := resource_schema.ParseGroupVersion(apv)
	if err != nil {
		return resource_schema.GroupVersionKind{}, err
	}
	mappings, err := m.RESTMappings(gv.WithKind(kind).GroupKind())
	if err != nil {
		return resource_schema.GroupVersionKind{}, err
	}
	for _, m := range mappings {
		if m.GroupVersionKind.GroupVersion().String() == apv {
			return m.GroupVersionKind, nil
		}
	}
	return resource_schema.GroupVersionKind{}, resource_errors.New("cannot select exact GV from REST mapper")
}

func IsResourceNamespaced(gvk resource_schema.GroupVersionKind, m resource_meta.RESTMapper) (bool, error) {
	rm, err := m.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return false, err
	}
	if rm.Scope.Name() == resource_meta.RESTScopeNameNamespace {
		return true, nil
	}
	return false, nil
}

func (ps *RawProviderServer) TFTypeFromOpenAPI(ctx resource_context.Context, gvk resource_schema.GroupVersionKind, status bool) (resource_tftypes.Type, map[string]string, error) {
	var tsch resource_tftypes.Type
	var hints map[string]string

	oapi, err := ps.getOAPIv2Foundry()
	if err != nil {
		return nil, hints, resource_fmt.Errorf("cannot get OpenAPI foundry: %s", err)
	}

	crdSchema, err := ps.lookUpGVKinCRDs(ctx, gvk)
	if err != nil {
		return nil, hints, resource_fmt.Errorf("failed to look up GVK [%s] among available CRDs: %s", gvk.String(), err)
	}
	if crdSchema != nil {
		js, err := resource_json.Marshal(resource_openapi.SchemaToSpec("", crdSchema.(map[string]interface{})))
		if err != nil {
			return nil, hints, resource_fmt.Errorf("CRD schema fails to marshal into JSON: %s", err)
		}
		oapiv3, err := resource_openapi.NewFoundryFromSpecV3(js)
		if err != nil {
			return nil, hints, err
		}
		tsch, hints, err = oapiv3.GetTypeByGVK(gvk)
		if err != nil {
			return nil, hints, resource_fmt.Errorf("failed to generate tftypes for GVK [%s] from CRD schema: %s", gvk.String(), err)
		}
	}
	if tsch == nil {

		tsch, hints, err = oapi.GetTypeByGVK(gvk)
		if err != nil {
			return nil, hints, resource_fmt.Errorf("cannot get resource type from OpenAPI (%s): %s", gvk.String(), err)
		}
	}

	if tsch.Is(resource_tftypes.Object{}) && !status {
		ot := tsch.(resource_tftypes.Object)
		atts := make(map[string]resource_tftypes.Type)
		for k, t := range ot.AttributeTypes {
			if k != "status" {
				atts[k] = t
			}
		}

		if _, ok := atts["apiVersion"]; !ok {
			atts["apiVersion"] = resource_tftypes.String
		}
		if _, ok := atts["kind"]; !ok {
			atts["kind"] = resource_tftypes.String
		}
		metaType, _, err := oapi.GetTypeByGVK(resource_openapi.ObjectMetaGVK)
		if err != nil {
			return nil, hints, resource_fmt.Errorf("failed to generate tftypes for v1.ObjectMeta: %s", err)
		}
		atts["metadata"] = metaType.(resource_tftypes.Object)

		tsch = resource_tftypes.Object{AttributeTypes: atts}
	}

	return tsch, hints, nil
}

func mapRemoveNulls(in map[string]interface{}) map[string]interface{} {
	for k, v := range in {
		switch tv := v.(type) {
		case []interface{}:
			in[k] = sliceRemoveNulls(tv)
		case map[string]interface{}:
			in[k] = mapRemoveNulls(tv)
		default:
			if v == nil {
				delete(in, k)
			}
		}
	}
	return in
}

func sliceRemoveNulls(in []interface{}) []interface{} {
	s := []interface{}{}
	for _, v := range in {
		switch tv := v.(type) {
		case []interface{}:
			s = append(s, sliceRemoveNulls(tv))
		case map[string]interface{}:
			s = append(s, mapRemoveNulls(tv))
		default:
			if v != nil {
				s = append(s, v)
			}
		}
	}
	return s
}

func RemoveServerSideFields(in map[string]interface{}) map[string]interface{} {

	delete(in, "status")

	resource_meta := in["metadata"].(map[string]interface{})

	delete(resource_meta, "uid")
	delete(resource_meta, "creationTimestamp")
	delete(resource_meta, "resourceVersion")
	delete(resource_meta, "generation")
	delete(resource_meta, "selfLink")

	delete(resource_meta, "managedFields")

	return in
}

func (ps *RawProviderServer) lookUpGVKinCRDs(ctx resource_context.Context, gvk resource_schema.GroupVersionKind) (interface{}, error) {
	c, err := ps.getDynamicClient()
	if err != nil {
		return nil, err
	}
	m, err := ps.getRestMapper()
	if err != nil {
		return nil, err
	}

	crd := resource_schema.GroupKind{Group: "apiextensions.k8s.io", Kind: "CustomResourceDefinition"}
	crms, err := m.RESTMappings(crd)
	if err != nil {
		return nil, resource_fmt.Errorf("could not extract resource version mappings for apiextensions.k8s.io.CustomResourceDefinition: %s", err)
	}

	for _, crm := range crms {
		crdRes, err := c.Resource(crm.Resource).List(ctx, resource_v1v1.ListOptions{})
		if err != nil {
			return nil, err
		}

		for _, r := range crdRes.Items {
			spec := r.Object["spec"].(map[string]interface{})
			if spec == nil {
				continue
			}
			grp := spec["group"].(string)
			if grp != gvk.Group {
				continue
			}
			names := spec["names"]
			if names == nil {
				continue
			}
			kind := names.(map[string]interface{})["kind"]
			if kind != gvk.Kind {
				continue
			}
			ver := spec["versions"]
			if ver == nil {
				ver = spec["version"]
				if ver == nil {
					continue
				}
			}
			for _, rv := range ver.([]interface{}) {
				if rv == nil {
					continue
				}
				v := rv.(map[string]interface{})
				if v["name"] == gvk.Version {
					s, ok := v["schema"].(map[string]interface{})
					if !ok {
						return nil, nil
					}
					return s["openAPIV3Schema"], nil
				}
			}
		}
	}
	return nil, nil
}

var privateStateSchema resource_tftypes.Object = resource_tftypes.Object{AttributeTypes: map[string]resource_tftypes.Type{
	"IsImported": resource_tftypes.Bool,
}}

func getPrivateStateValue(p []byte) (ps map[string]resource_tftypes.Value, err error) {
	if p == nil {
		err = resource_errors.New("private state value is nil")
		return
	}
	pv, err := resource_tftypes.ValueFromMsgPack(p, privateStateSchema)
	err = pv.As(&ps)
	return
}

func init() {
	server_install.Install(server_scheme.Scheme)
}

type RawProviderServer struct {
	logger          server_hclog.Logger
	clientConfig    *server_rest.Config
	dynamicClient   server_dynamic.Interface
	discoveryClient server_discovery.DiscoveryInterface
	restMapper      server_meta.RESTMapper
	restClient      server_rest.Interface
	OAPIFoundry     server_openapi.Foundry

	providerEnabled bool
	hostTFVersion   string
}

func dump(v interface{}) server_hclog.Format {
	return server_hclog.Fmt("%v", v)
}

func (s *RawProviderServer) PrepareProviderConfig(ctx server_context.Context, req *server_tfprotov5.PrepareProviderConfigRequest) (*server_tfprotov5.PrepareProviderConfigResponse, error) {
	s.logger.Trace("[PrepareProviderConfig][Request]\n%s\n", dump(*req))
	resp := &server_tfprotov5.PrepareProviderConfigResponse{}
	return resp, nil
}

func (s *RawProviderServer) ValidateDataSourceConfig(ctx server_context.Context, req *server_tfprotov5.ValidateDataSourceConfigRequest) (*server_tfprotov5.ValidateDataSourceConfigResponse, error) {
	s.logger.Trace("[ValidateDataSourceConfig][Request]\n%s\n", dump(*req))
	resp := &server_tfprotov5.ValidateDataSourceConfigResponse{}
	return resp, nil
}

func (s *RawProviderServer) StopProvider(ctx server_context.Context, req *server_tfprotov5.StopProviderRequest) (*server_tfprotov5.StopProviderResponse, error) {
	s.logger.Trace("[StopProvider][Request]\n%s\n", dump(*req))

	return nil, server_status.Errorf(server_codes.Unimplemented, "method Stop not implemented")
}

func (s *RawProviderServer) UpgradeResourceState(ctx upgrade_state_context.Context, req *upgrade_state_tfprotov5.UpgradeResourceStateRequest) (*upgrade_state_tfprotov5.UpgradeResourceStateResponse, error) {
	resp := &upgrade_state_tfprotov5.UpgradeResourceStateResponse{}
	resp.Diagnostics = []*upgrade_state_tfprotov5.Diagnostic{}

	sch := GetProviderResourceSchema()
	rt := GetObjectTypeFromSchema(sch[req.TypeName])

	rv, err := req.RawState.Unmarshal(rt)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &upgrade_state_tfprotov5.Diagnostic{
			Severity: upgrade_state_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to unmarshal old state during upgrade",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	cd := s.checkValidCredentials(ctx)
	if len(cd) > 0 {
		us, err := upgrade_state_tfprotov5.NewDynamicValue(rt, rv)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &upgrade_state_tfprotov5.Diagnostic{
				Severity: upgrade_state_tfprotov5.DiagnosticSeverityError,
				Summary:  "Failed to encode new state during upgrade",
				Detail:   err.Error(),
			})
		}
		resp.UpgradedState = &us

		return resp, nil
	}

	var cs map[string]upgrade_state_tftypes.Value
	err = rv.As(&cs)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &upgrade_state_tfprotov5.Diagnostic{
			Severity: upgrade_state_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to extract values from old state during upgrade",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	obj, ok := cs["object"]
	if !ok {
		resp.Diagnostics = append(resp.Diagnostics, &upgrade_state_tfprotov5.Diagnostic{
			Severity: upgrade_state_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to find object value in existing resource state",
		})
		return resp, nil
	}

	m, err := s.getRestMapper()
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics,
			&upgrade_state_tfprotov5.Diagnostic{
				Severity: upgrade_state_tfprotov5.DiagnosticSeverityError,
				Summary:  "Failed to retrieve Kubernetes RESTMapper client during state upgrade",
				Detail:   err.Error(),
			})
		return resp, nil
	}

	gvk, err := GVKFromTftypesObject(&obj, m)
	if err != nil {
		return resp, upgrade_state_fmt.Errorf("failed to determine resource GVK: %s", err)
	}

	tsch, _, err := s.TFTypeFromOpenAPI(ctx, gvk, false)
	if err != nil {
		return resp, upgrade_state_fmt.Errorf("failed to determine resource type ID: %s", err)
	}

	morphedObject, d := upgrade_state_morph.ValueToType(obj, tsch, upgrade_state_tftypes.NewAttributePath())
	if len(d) > 0 {
		resp.Diagnostics = append(resp.Diagnostics, d...)
		for i := range d {
			if d[i].Severity == upgrade_state_tfprotov5.DiagnosticSeverityError {
				return resp, nil
			}
		}
	}
	s.logger.Debug("[UpgradeResourceState]", "morphed object", dump(morphedObject))

	cs["object"] = obj

	newStateVal := upgrade_state_tftypes.NewValue(rv.Type(), cs)

	us, err := upgrade_state_tfprotov5.NewDynamicValue(rt, newStateVal)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &upgrade_state_tfprotov5.Diagnostic{
			Severity: upgrade_state_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to encode new state during upgrade",
			Detail:   err.Error(),
		})
	}
	resp.UpgradedState = &us

	return resp, nil
}

func (s *RawProviderServer) ValidateResourceTypeConfig(ctx validate_context.Context, req *validate_tfprotov5.ValidateResourceTypeConfigRequest) (*validate_tfprotov5.ValidateResourceTypeConfigResponse, error) {
	resp := &validate_tfprotov5.ValidateResourceTypeConfigResponse{}
	requiredKeys := []string{"apiVersion", "kind", "metadata"}
	forbiddenKeys := []string{"status"}

	rt, err := GetResourceType(req.TypeName)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &validate_tfprotov5.Diagnostic{
			Severity: validate_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to determine resource type",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	config, err := req.Config.Unmarshal(rt)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &validate_tfprotov5.Diagnostic{
			Severity: validate_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to unmarshal resource state",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	att := validate_tftypes.NewAttributePath()
	att = att.WithAttributeName("manifest")

	configVal := make(map[string]validate_tftypes.Value)
	err = config.As(&configVal)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &validate_tfprotov5.Diagnostic{
			Severity: validate_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to extract resource state from SDK value",
			Detail:   err.Error(),
		})
		return resp, nil
	}

	manifest, ok := configVal["manifest"]
	if !ok {
		resp.Diagnostics = append(resp.Diagnostics, &validate_tfprotov5.Diagnostic{
			Severity:  validate_tfprotov5.DiagnosticSeverityError,
			Summary:   "Manifest missing from resource configuration",
			Detail:    "A manifest attribute containing a valid Kubernetes resource configuration is required.",
			Attribute: att,
		})
		return resp, nil
	}

	rawManifest := make(map[string]validate_tftypes.Value)
	err = manifest.As(&rawManifest)
	if err != nil {
		if err.Error() == "unmarshaling unknown values is not supported" {

			return resp, nil
		}
		resp.Diagnostics = append(resp.Diagnostics, &validate_tfprotov5.Diagnostic{
			Severity:  validate_tfprotov5.DiagnosticSeverityError,
			Summary:   `Failed to extract "manifest" attribute value from resource configuration`,
			Detail:    err.Error(),
			Attribute: att,
		})
		return resp, nil
	}

	for _, key := range requiredKeys {
		if _, present := rawManifest[key]; !present {
			kp := att.WithAttributeName(key)
			resp.Diagnostics = append(resp.Diagnostics, &validate_tfprotov5.Diagnostic{
				Severity:  validate_tfprotov5.DiagnosticSeverityError,
				Summary:   `Attribute key missing from "manifest" value`,
				Detail:    validate_fmt.Sprintf("'%s' attribute key is missing from manifest configuration", key),
				Attribute: kp,
			})
		}
	}

	for _, key := range forbiddenKeys {
		if _, present := rawManifest[key]; present {
			kp := att.WithAttributeName(key)
			resp.Diagnostics = append(resp.Diagnostics, &validate_tfprotov5.Diagnostic{
				Severity:  validate_tfprotov5.DiagnosticSeverityError,
				Summary:   `Forbidden attribute key in "manifest" value`,
				Detail:    validate_fmt.Sprintf("'%s' attribute key is not allowed in manifest configuration", key),
				Attribute: kp,
			})
		}
	}

	timeouts := s.getTimeouts(configVal)
	path := validate_tftypes.NewAttributePath().WithAttributeName("timeouts")
	for k, v := range timeouts {
		_, err := validate_time.ParseDuration(v)
		if err != nil {
			resp.Diagnostics = append(resp.Diagnostics, &validate_tfprotov5.Diagnostic{
				Severity:  validate_tfprotov5.DiagnosticSeverityError,
				Summary:   validate_fmt.Sprintf("Error parsing timeout for %q", k),
				Detail:    err.Error(),
				Attribute: path.WithAttributeName(k),
			})
		}
	}

	if wait, ok := configVal["wait"]; ok && !wait.IsNull() {
		var waitBlock []validate_tftypes.Value
		wait.As(&waitBlock)
		if len(waitBlock) > 0 {
			var w map[string]validate_tftypes.Value
			waitBlock[0].As(&w)
			waiters := []string{}
			for k, ww := range w {
				if !ww.IsNull() {
					if k == "condition" {
						var cb []validate_tftypes.Value
						ww.As(&cb)
						if len(cb) == 0 {
							continue
						}
					}
					waiters = append(waiters, k)
				}
			}
			if len(waiters) > 1 {
				resp.Diagnostics = append(resp.Diagnostics, &validate_tfprotov5.Diagnostic{
					Severity:  validate_tfprotov5.DiagnosticSeverityError,
					Summary:   "Invalid wait configuration",
					Detail:    validate_fmt.Sprintf(`You may only set one of "%s".`, validate_strings.Join(waiters, "\", \"")),
					Attribute: validate_tftypes.NewAttributePath().WithAttributeName("wait"),
				})
			}
		}
	}
	if waitFor, ok := configVal["wait_for"]; ok && !waitFor.IsNull() {
		resp.Diagnostics = append(resp.Diagnostics, &validate_tfprotov5.Diagnostic{
			Severity:  validate_tfprotov5.DiagnosticSeverityWarning,
			Summary:   "Deprecated Attribute",
			Detail:    `The "wait_for" attribute has been deprecated. Please use the "wait" block instead.`,
			Attribute: validate_tftypes.NewAttributePath().WithAttributeName("wait_for"),
		})
	}

	return resp, nil
}

func (s *RawProviderServer) validateResourceOnline(manifest *validate_tftypes.Value) (diags []*validate_tfprotov5.Diagnostic) {
	rm, err := s.getRestMapper()
	if err != nil {
		diags = append(diags, &validate_tfprotov5.Diagnostic{
			Severity: validate_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to create K8s RESTMapper client",
			Detail:   err.Error(),
		})
		return
	}
	gvk, err := GVKFromTftypesObject(manifest, rm)
	if err != nil {
		diags = append(diags, &validate_tfprotov5.Diagnostic{
			Severity: validate_tfprotov5.DiagnosticSeverityError,
			Summary:  "Failed to determine GroupVersionResource for manifest",
			Detail:   err.Error(),
		})
		return
	}

	ns, err := IsResourceNamespaced(gvk, rm)
	if err != nil {
		diags = append(diags,
			&validate_tfprotov5.Diagnostic{
				Severity: validate_tfprotov5.DiagnosticSeverityError,
				Detail:   err.Error(),
				Summary:  validate_fmt.Sprintf("Failed to discover scope of resource '%s'", gvk.String()),
			})
		return
	}
	nsPath := validate_tftypes.NewAttributePath()
	nsPath = nsPath.WithAttributeName("metadata").WithAttributeName("namespace")
	nsVal, restPath, err := validate_tftypes.WalkAttributePath(*manifest, nsPath)
	if ns {
		if err != nil || len(restPath.Steps()) > 0 {
			diags = append(diags,
				&validate_tfprotov5.Diagnostic{
					Severity: validate_tfprotov5.DiagnosticSeverityError,
					Detail:   validate_fmt.Sprintf("Resources of type '%s' require a namespace", gvk.String()),
					Summary:  "Namespace required",
				})
			return
		}
		if nsVal.(validate_tftypes.Value).IsNull() {
			diags = append(diags,
				&validate_tfprotov5.Diagnostic{
					Severity: validate_tfprotov5.DiagnosticSeverityError,
					Detail:   validate_fmt.Sprintf("Namespace for resource '%s' cannot be nil", gvk.String()),
					Summary:  "Namespace required",
				})
		}
		var nsStr string
		err := nsVal.(validate_tftypes.Value).As(&nsStr)
		if nsStr == "" && err == nil {
			diags = append(diags,
				&validate_tfprotov5.Diagnostic{
					Severity: validate_tfprotov5.DiagnosticSeverityError,
					Detail:   validate_fmt.Sprintf("Namespace for resource '%s' cannot be empty", gvk.String()),
					Summary:  "Namespace required",
				})
		}
	} else {
		if err == nil && len(restPath.Steps()) == 0 && !nsVal.(validate_tftypes.Value).IsNull() {
			diags = append(diags,
				&validate_tfprotov5.Diagnostic{
					Severity: validate_tfprotov5.DiagnosticSeverityError,
					Detail:   validate_fmt.Sprintf("Resources of type '%s' cannot have a namespace", gvk.String()),
					Summary:  "Cluster level resource cannot take namespace",
				})
		}
	}
	return
}

const waiterSleepTime = 1 * waiter_time.Second

func (s *RawProviderServer) waitForCompletion(ctx waiter_context.Context, waitForBlock waiter_tftypes.Value, rs waiter_dynamic.ResourceInterface, rname string, rtype waiter_tftypes.Type, th map[string]string) error {
	if waitForBlock.IsNull() || !waitForBlock.IsKnown() {
		return nil
	}

	waiter, err := NewResourceWaiter(rs, rname, rtype, th, waitForBlock, s.logger)
	if err != nil {
		return err
	}
	return waiter.Wait(ctx)
}

type Waiter interface {
	Wait(waiter_context.Context) error
}

func NewResourceWaiter(resource waiter_dynamic.ResourceInterface, resourceName string, resourceType waiter_tftypes.Type, th map[string]string, waitForBlock waiter_tftypes.Value, hl waiter_hclog.Logger) (Waiter, error) {
	var waitForBlockVal map[string]waiter_tftypes.Value
	err := waitForBlock.As(&waitForBlockVal)
	if err != nil {
		return nil, err
	}

	if v, ok := waitForBlockVal["rollout"]; ok {
		var rollout bool
		v.As(&rollout)
		if rollout {
			return &RolloutWaiter{
				resource,
				resourceName,
				hl,
			}, nil
		}
	}

	if v, ok := waitForBlockVal["condition"]; ok {
		var conditionsBlocks []waiter_tftypes.Value
		v.As(&conditionsBlocks)
		if len(conditionsBlocks) > 0 {
			return &ConditionsWaiter{
				resource,
				resourceName,
				conditionsBlocks,
				hl,
			}, nil
		}
	}

	fields, ok := waitForBlockVal["fields"]
	if !ok || fields.IsNull() || !fields.IsKnown() {
		return &NoopWaiter{}, nil
	}

	if !fields.Type().Is(waiter_tftypes.Map{}) {
		return nil, waiter_fmt.Errorf(`"fields" should be a map of strings`)
	}

	var vm map[string]waiter_tftypes.Value
	fields.As(&vm)
	var matchers []FieldMatcher

	for k, v := range vm {
		var expr string
		v.As(&expr)
		var re *waiter_regexp.Regexp
		if expr == "*" {

			re = waiter_regexp.MustCompile("(.*)?")
		} else {
			var err error
			re, err = waiter_regexp.Compile(expr)
			if err != nil {
				return nil, waiter_fmt.Errorf("invalid regular expression: %q", expr)
			}
		}

		p, err := FieldPathToTftypesPath(k)
		if err != nil {
			return nil, err
		}
		matchers = append(matchers, FieldMatcher{p, re})
	}

	return &FieldWaiter{
		resource,
		resourceName,
		resourceType,
		th,
		matchers,
		hl,
	}, nil

}

type FieldMatcher struct {
	path         *waiter_tftypes.AttributePath
	valueMatcher *waiter_regexp.Regexp
}

type FieldWaiter struct {
	resource      waiter_dynamic.ResourceInterface
	resourceName  string
	resourceType  waiter_tftypes.Type
	typeHints     map[string]string
	fieldMatchers []FieldMatcher
	logger        waiter_hclog.Logger
}

func (w *FieldWaiter) Wait(ctx waiter_context.Context) error {
	w.logger.Info("[ApplyResourceChange][Wait] Waiting until ready...\n")
	for {
		if deadline, ok := ctx.Deadline(); ok {
			if waiter_time.Now().After(deadline) {
				return waiter_context.DeadlineExceeded
			}
		}

		res, err := w.resource.Get(ctx, w.resourceName, waiter_v1v1.GetOptions{})
		if err != nil {
			return err
		}
		if waiter_errors.IsGone(err) {
			return waiter_fmt.Errorf("resource was deleted")
		}
		resObj := res.Object
		meta := resObj["metadata"].(map[string]interface{})
		delete(meta, "managedFields")

		w.logger.Trace("[ApplyResourceChange][Wait]", "API Response", resObj)

		obj, err := waiter_payload.ToTFValue(resObj, w.resourceType, w.typeHints, waiter_tftypes.NewAttributePath())
		if err != nil {
			return err
		}

		done, err := func(obj waiter_tftypes.Value) (bool, error) {
			for _, m := range w.fieldMatchers {
				vi, rp, err := waiter_tftypes.WalkAttributePath(obj, m.path)
				if err != nil {
					return false, err
				}
				if len(rp.Steps()) > 0 {
					return false, waiter_fmt.Errorf("attribute not present at path '%s'", m.path.String())
				}

				var s string
				v := vi.(waiter_tftypes.Value)
				switch {
				case v.Type().Is(waiter_tftypes.String):
					v.As(&s)
				case v.Type().Is(waiter_tftypes.Bool):
					var vb bool
					v.As(&vb)
					s = waiter_fmt.Sprintf("%t", vb)
				case v.Type().Is(waiter_tftypes.Number):
					var f waiter_big.Float
					v.As(&f)
					if f.IsInt() {
						i, _ := f.Int64()
						s = waiter_fmt.Sprintf("%d", i)
					} else {
						i, _ := f.Float64()
						s = waiter_fmt.Sprintf("%f", i)
					}
				default:
					return true, waiter_fmt.Errorf("wait_for: cannot match on type %q", v.Type().String())
				}

				if !m.valueMatcher.Match([]byte(s)) {
					return false, nil
				}
			}

			return true, nil
		}(obj)

		if done {
			w.logger.Info("[ApplyResourceChange][Wait] Done waiting.\n")
			return err
		}

		waiter_time.Sleep(waiterSleepTime)
	}
}

type NoopWaiter struct{}

func (w *NoopWaiter) Wait(_ waiter_context.Context) error {
	return nil
}

func FieldPathToTftypesPath(fieldPath string) (*waiter_tftypes.AttributePath, error) {
	t, d := waiter_hclsyntax.ParseTraversalAbs([]byte(fieldPath), "", waiter_hclhcl.Pos{Line: 1, Column: 1})
	if d.HasErrors() {
		return waiter_tftypes.NewAttributePath(), waiter_fmt.Errorf("invalid field path %q: %s", fieldPath, d.Error())
	}

	path := waiter_tftypes.NewAttributePath()
	for _, p := range t {
		switch p.(type) {
		case waiter_hclhcl.TraverseRoot:
			path = path.WithAttributeName(p.(waiter_hclhcl.TraverseRoot).Name)
		case waiter_hclhcl.TraverseIndex:
			indexKey := p.(waiter_hclhcl.TraverseIndex).Key
			indexKeyType := indexKey.Type()
			if indexKeyType.Equals(waiter_cty.String) {
				path = path.WithElementKeyString(indexKey.AsString())
			} else if indexKeyType.Equals(waiter_cty.Number) {
				f := indexKey.AsBigFloat()
				if f.IsInt() {
					i, _ := f.Int64()
					path = path.WithElementKeyInt(int(i))
				} else {
					return waiter_tftypes.NewAttributePath(), waiter_fmt.Errorf("index in field path must be an integer")
				}
			} else {
				return waiter_tftypes.NewAttributePath(), waiter_fmt.Errorf("unsupported type in field path: %s", indexKeyType.FriendlyName())
			}
		case waiter_hclhcl.TraverseAttr:
			path = path.WithAttributeName(p.(waiter_hclhcl.TraverseAttr).Name)
		case waiter_hclhcl.TraverseSplat:
			return waiter_tftypes.NewAttributePath(), waiter_fmt.Errorf("splat is not supported")
		}
	}

	return path, nil
}

type RolloutWaiter struct {
	resource     waiter_dynamic.ResourceInterface
	resourceName string
	logger       waiter_hclog.Logger
}

func (w *RolloutWaiter) Wait(ctx waiter_context.Context) error {
	w.logger.Info("[ApplyResourceChange][Wait] Waiting until rollout complete...\n")
	for {
		if deadline, ok := ctx.Deadline(); ok {
			if waiter_time.Now().After(deadline) {
				return waiter_context.DeadlineExceeded
			}
		}

		res, err := w.resource.Get(ctx, w.resourceName, waiter_v1v1.GetOptions{})
		if err != nil {
			return err
		}
		if waiter_errors.IsGone(err) {
			return waiter_fmt.Errorf("resource was deleted")
		}

		gk := res.GetObjectKind().GroupVersionKind().GroupKind()
		statusViewer, err := waiter_polymorphichelpers.StatusViewerFor(gk)
		if err != nil {
			return waiter_fmt.Errorf("error getting resource status: %v", err)
		}

		_, done, err := statusViewer.Status(res, 0)
		if err != nil {
			return waiter_fmt.Errorf("error getting resource status: %v", err)
		}

		if done {
			break
		}

		waiter_time.Sleep(waiterSleepTime)
	}

	w.logger.Info("[ApplyResourceChange][Wait] Rollout complete\n")
	return nil
}

type ConditionsWaiter struct {
	resource     waiter_dynamic.ResourceInterface
	resourceName string
	conditions   []waiter_tftypes.Value
	logger       waiter_hclog.Logger
}

func (w *ConditionsWaiter) Wait(ctx waiter_context.Context) error {
	w.logger.Info("[ApplyResourceChange][Wait] Waiting for conditions...\n")

	for {
		if deadline, ok := ctx.Deadline(); ok {
			if waiter_time.Now().After(deadline) {
				return waiter_context.DeadlineExceeded
			}
		}

		res, err := w.resource.Get(ctx, w.resourceName, waiter_v1v1.GetOptions{})
		if err != nil {
			return err
		}
		if waiter_errors.IsGone(err) {
			return waiter_fmt.Errorf("resource was deleted")
		}

		status := res.Object["status"].(map[string]interface{})
		conditions := status["conditions"].([]interface{})
		conditionsMet := true
		for _, c := range w.conditions {
			var condition map[string]waiter_tftypes.Value
			c.As(&condition)
			var conditionType, conditionStatus string
			condition["type"].As(&conditionType)
			condition["status"].As(&conditionStatus)
			conditionMet := false
			for _, cc := range conditions {
				ccc := cc.(map[string]interface{})
				if ccc["type"].(string) == conditionType {
					conditionMet = ccc["status"].(string) == conditionStatus
					break
				}
			}
			conditionsMet = conditionsMet && conditionMet
		}

		if conditionsMet {
			break
		}

		waiter_time.Sleep(waiterSleepTime)
	}

	w.logger.Info("[ApplyResourceChange][Wait] All conditions met.\n")
	return nil
}
