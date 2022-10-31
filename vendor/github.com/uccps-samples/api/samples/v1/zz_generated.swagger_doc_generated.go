package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_Config = map[string]string{
	"": "Config contains the configuration and detailed condition status for the Samples Operator.",
}

func (Config) SwaggerDoc() map[string]string {
	return map_Config
}

var map_ConfigCondition = map[string]string{
	"":                   "ConfigCondition captures various conditions of the Config as entries are processed.",
	"type":               "type of condition.",
	"status":             "status of the condition, one of True, False, Unknown.",
	"lastUpdateTime":     "lastUpdateTime is the last time this condition was updated.",
	"lastTransitionTime": "lastTransitionTime is the last time the condition transitioned from one status to another.",
	"reason":             "reason is what caused the condition's last transition.",
	"message":            "message is a human readable message indicating details about the transition.",
}

func (ConfigCondition) SwaggerDoc() map[string]string {
	return map_ConfigCondition
}

var map_ConfigSpec = map[string]string{
	"":                    "ConfigSpec contains the desired configuration and state for the Samples Operator, controlling various behavior around the imagestreams and templates it creates/updates in the openshift namespace.",
	"managementState":     "managementState is top level on/off type of switch for all operators. When \"Managed\", this operator processes config and manipulates the samples accordingly. When \"Unmanaged\", this operator ignores any updates to the resources it watches. When \"Removed\", it reacts that same wasy as it does if the Config object is deleted, meaning any ImageStreams or Templates it manages (i.e. it honors the skipped lists) and the registry secret are deleted, along with the ConfigMap in the operator's namespace that represents the last config used to manipulate the samples,",
	"samplesRegistry":     "samplesRegistry allows for the specification of which registry is accessed by the ImageStreams for their image content.  Defaults on the content in https://github.com/uccps-samples/library that are pulled into this github repository, but based on our pulling only ocp content it typically defaults to registry.redhat.io.",
	"architectures":       "architectures determine which hardware architecture(s) to install, where x86_64, ppc64le, and s390x are the only supported choices currently.",
	"skippedImagestreams": "skippedImagestreams specifies names of image streams that should NOT be created/updated.  Admins can use this to allow them to delete content they don’t want.  They will still have to manually delete the content but the operator will not recreate(or update) anything listed here.",
	"skippedTemplates":    "skippedTemplates specifies names of templates that should NOT be created/updated.  Admins can use this to allow them to delete content they don’t want.  They will still have to manually delete the content but the operator will not recreate(or update) anything listed here.",
}

func (ConfigSpec) SwaggerDoc() map[string]string {
	return map_ConfigSpec
}

var map_ConfigStatus = map[string]string{
	"":                    "ConfigStatus contains the actual configuration in effect, as well as various details that describe the state of the Samples Operator.",
	"managementState":     "managementState reflects the current operational status of the on/off switch for the operator.  This operator compares the ManagementState as part of determining that we are turning the operator back on (i.e. \"Managed\") when it was previously \"Unmanaged\".",
	"conditions":          "conditions represents the available maintenance status of the sample imagestreams and templates.",
	"samplesRegistry":     "samplesRegistry allows for the specification of which registry is accessed by the ImageStreams for their image content.  Defaults on the content in https://github.com/uccps-samples/library that are pulled into this github repository, but based on our pulling only ocp content it typically defaults to registry.redhat.io.",
	"architectures":       "architectures determine which hardware architecture(s) to install, where x86_64 and ppc64le are the supported choices.",
	"skippedImagestreams": "skippedImagestreams specifies names of image streams that should NOT be created/updated.  Admins can use this to allow them to delete content they don’t want.  They will still have to manually delete the content but the operator will not recreate(or update) anything listed here.",
	"skippedTemplates":    "skippedTemplates specifies names of templates that should NOT be created/updated.  Admins can use this to allow them to delete content they don’t want.  They will still have to manually delete the content but the operator will not recreate(or update) anything listed here.",
	"version":             "version is the value of the operator's payload based version indicator when it was last successfully processed",
}

func (ConfigStatus) SwaggerDoc() map[string]string {
	return map_ConfigStatus
}

// AUTO-GENERATED FUNCTIONS END HERE