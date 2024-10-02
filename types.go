package main

type compute struct {
	Memory string `yaml:"memory"`
	CPU    string `yaml:"cpu"`
}

type resource struct {
	Requests compute `yaml:"requests"`
	Limits   compute `yaml:"limits"`
}

type ContainerItem struct {
	Name            string          `yaml:"name"`
	Resources       resource        `yaml:"resources"`
	SecurityContext securityContext `yaml:"securityContext"`
}

type spec struct {
	Container          []ContainerItem `yaml:"containers"`
	SecurityContext    securityContext `yaml:"securityContext"`
	ServiceAccountName string          `yaml:"serviceAccountName"`
}

type metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type securityContext struct {
	RunAsUser              int32 `yaml:"runAsUser"`
	RunAsGroup             int32 `yaml:"runAsGroup"`
	ReadOnlyRootFilesystem bool  `yaml:"readOnlyRootFilesystem"`
	Privileged             bool  `yaml:"privileged"`
}

type podSpec struct {
	Spec      spec     `yaml:"spec"`
	Kind      string   `yaml:"kind"`
	Immutable bool     `yaml:"immutable"`
	Metadata  metadata `yaml:"metadata"`
}
