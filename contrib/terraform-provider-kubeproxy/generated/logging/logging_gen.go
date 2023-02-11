package logging

import (
	context_context "context"
	helper_resource_context "context"
	helper_schema_context "context"

	context_tfsdklog "github.com/hashicorp/terraform-plugin-log/tfsdklog"
	helper_resource_tfsdklog "github.com/hashicorp/terraform-plugin-log/tfsdklog"
	helper_schema_tfsdklog "github.com/hashicorp/terraform-plugin-log/tfsdklog"
	context_logginghelperlogging "github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	context_testingtesting "github.com/mitchellh/go-testing-interface"
)

func InitContext(ctx context_context.Context) context_context.Context {
	ctx = context_tfsdklog.NewSubsystem(ctx, SubsystemHelperSchema,

		context_tfsdklog.WithAdditionalLocationOffset(1),
		context_tfsdklog.WithLevelFromEnv(EnvTfLogSdkHelperSchema),

		context_tfsdklog.WithRootFields(),
	)

	return ctx
}

func InitTestContext(ctx context_context.Context, t context_testingtesting.T) context_context.Context {
	context_logginghelperlogging.SetOutput(t)

	ctx = context_tfsdklog.RegisterTestSink(ctx, t)
	ctx = context_tfsdklog.NewRootSDKLogger(ctx, context_tfsdklog.WithLevelFromEnv(EnvTfLogSdk))
	ctx = context_tfsdklog.NewSubsystem(ctx, SubsystemHelperResource,

		context_tfsdklog.WithAdditionalLocationOffset(1),
		context_tfsdklog.WithLevelFromEnv(EnvTfLogSdkHelperResource),
	)
	ctx = TestNameContext(ctx, t.Name())

	return ctx
}

func TestNameContext(ctx context_context.Context, testName string) context_context.Context {
	ctx = context_tfsdklog.SubsystemSetField(ctx, SubsystemHelperResource, KeyTestName, testName)

	return ctx
}

func TestStepNumberContext(ctx context_context.Context, stepNumber int) context_context.Context {
	ctx = context_tfsdklog.SubsystemSetField(ctx, SubsystemHelperResource, KeyTestStepNumber, stepNumber)

	return ctx
}

func TestTerraformPathContext(ctx context_context.Context, terraformPath string) context_context.Context {
	ctx = context_tfsdklog.SubsystemSetField(ctx, SubsystemHelperResource, KeyTestTerraformPath, terraformPath)

	return ctx
}

func TestWorkingDirectoryContext(ctx context_context.Context, workingDirectory string) context_context.Context {
	ctx = context_tfsdklog.SubsystemSetField(ctx, SubsystemHelperResource, KeyTestWorkingDirectory, workingDirectory)

	return ctx
}

const (
	EnvTfLogSdk = "TF_LOG_SDK"

	EnvTfLogSdkHelperResource = "TF_LOG_SDK_HELPER_RESOURCE"

	EnvTfLogSdkHelperSchema = "TF_LOG_SDK_HELPER_SCHEMA"
)

const (
	SubsystemHelperResource = "helper_resource"
)

func HelperResourceTrace(ctx helper_resource_context.Context, msg string, additionalFields ...map[string]interface{}) {
	helper_resource_tfsdklog.SubsystemTrace(ctx, SubsystemHelperResource, msg, additionalFields...)
}

func HelperResourceDebug(ctx helper_resource_context.Context, msg string, additionalFields ...map[string]interface{}) {
	helper_resource_tfsdklog.SubsystemDebug(ctx, SubsystemHelperResource, msg, additionalFields...)
}

func HelperResourceWarn(ctx helper_resource_context.Context, msg string, additionalFields ...map[string]interface{}) {
	helper_resource_tfsdklog.SubsystemWarn(ctx, SubsystemHelperResource, msg, additionalFields...)
}

func HelperResourceError(ctx helper_resource_context.Context, msg string, additionalFields ...map[string]interface{}) {
	helper_resource_tfsdklog.SubsystemError(ctx, SubsystemHelperResource, msg, additionalFields...)
}

const (
	SubsystemHelperSchema = "helper_schema"
)

func HelperSchemaDebug(ctx helper_schema_context.Context, msg string, additionalFields ...map[string]interface{}) {
	helper_schema_tfsdklog.SubsystemDebug(ctx, SubsystemHelperSchema, msg, additionalFields...)
}

func HelperSchemaError(ctx helper_schema_context.Context, msg string, additionalFields ...map[string]interface{}) {
	helper_schema_tfsdklog.SubsystemError(ctx, SubsystemHelperSchema, msg, additionalFields...)
}

func HelperSchemaTrace(ctx helper_schema_context.Context, msg string, additionalFields ...map[string]interface{}) {
	helper_schema_tfsdklog.SubsystemTrace(ctx, SubsystemHelperSchema, msg, additionalFields...)
}

func HelperSchemaWarn(ctx helper_schema_context.Context, msg string, additionalFields ...map[string]interface{}) {
	helper_schema_tfsdklog.SubsystemWarn(ctx, SubsystemHelperSchema, msg, additionalFields...)
}

const (
	KeyAttributePath = "tf_attribute_path"

	KeyDataSourceType = "tf_data_source_type"

	KeyError = "error"

	KeyProviderAddress = "tf_provider_addr"

	KeyResourceType = "tf_resource_type"

	KeyTestName = "test_name"

	KeyTestStepNumber = "test_step_number"

	KeyTestTerraformConfiguration = "test_terraform_configuration"

	KeyTestTerraformLogLevel = "test_terraform_log_level"

	KeyTestTerraformLogCoreLevel = "test_terraform_log_core_level"

	KeyTestTerraformLogProviderLevel = "test_terraform_log_provider_level"

	KeyTestTerraformLogPath = "test_terraform_log_path"

	KeyTestTerraformPath = "test_terraform_path"

	KeyTestTerraformPlan = "test_terraform_plan"

	KeyTestWorkingDirectory = "test_working_directory"
)
