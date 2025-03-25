package v1alpha1

// TODO
//
// - Underlying operations (installing bundles) should all be doable without the operator.
//   Aim to use plain `oc apply` with -k for variables?
// - Install bundle format: directory, configmap - other?
//

type ClusterObservability struct {
	Spec               ClusterObservabilitySpec
	Status             ClusterObservabilityStatus
	InstallDefinitions []InstallDefinitionSpec
}

type ClusterObservabilitySpec struct {
	// DefaultInstall is the default install type for signals that are not listed or
	// are listed without an `install` field.
	// For example, `{ defaultinstall: Default }` with no `signals` field installs all signal
	// types with default settings.
	DefaultInstall Install `json:"defaultInstall"`
	// Signals specifies an install type for each signal.
	Signals []SignalSpec `json:"signals,omitempty"`
}

type SignalSpec struct {
	// Name of this signal type.
	Name Signal `json:"signal"`
	// Install type for this signal. Optional, if absent use ..defaultInstall
	Install Install `json:"installType,omitempty"`
	// Namespace to install. Optional, each signal type has a default namespace.
	// A signal can be listed multiple times with different `namespace` values,
	// to install in multiple namespaces.
	Namespace string `json:"namespace,omitempty"`
}

// Signal is the name of a signal type
type Signal string

const (
	SignalLog     = "Log"
	SignalTrace   = "Trace"
	SignalMetric  = "Metric"
	SignalNetflow = "Netflow"
	// TODO ... fill the list
)

// Install is the name of a installation type, the following names are always recognized:
//   - Default: Install operators, custom resource definitions and default resources
//     suitable for most use cases in a single cluster.
//     Reconcile the resources to keep them in the default state.
//   - Custom: Install operators and custom resource definitions, but do not create resources.
//     User can create customized resources, they will not be owned or modified by this operator.
//   - None: Do not install anything, no operators or resource definitions.
//   - Other name: Other named installation types may be available.
type Install string

const (
	InstallDefault         = "Default"
	InstallCustom          = "Custom"
	InstallNone    Install = "None"
)

// InstallDefinitionSpec defines a new "Install" type.
type InstallDefinitionSpec struct {
	Name Install
	// ConfigMap contains deployment bundles for the install type,
	// with key=signal type.
	// TODO: details. Are there other places to get definitions besides configmaps?
	ConfigMap *NamespacedName
}
