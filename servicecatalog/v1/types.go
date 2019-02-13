package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenShiftServiceCatalogConfig holds the necessary configuration options for running the Service Catalog
type OpenShiftServiceCatalogConfig struct {
	metav1.TypeMeta `json:",inline"`

	//enabled specifies if Service Catalog should be installed and started in the cluster
	Enabled bool `json:"enabled"`

	//serverArguments are key value pairs that will override the default set of arguments.  ServerArguments may include Feature Gates & Admission Plugins.  If you specify an invalid argument the server will not start. These values may cause invalid configurations.",
	ServerArguments map[string][]string `json:"serverArguments"`
}
