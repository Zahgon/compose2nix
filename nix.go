package main

import (
	"io"
)

// Compose V2 uses "-" for container names: https://docs.docker.com/compose/migrate/#service-container-names
//
// Volumes and networks still use "_", but we'll ignore that here: https://github.com/docker/compose/issues/9618
const DefaultProjectSeparator = "-"

type ContainerRuntime int

const (
	ContainerRuntimeInvalid ContainerRuntime = iota
	ContainerRuntimeDocker
	ContainerRuntimePodman
)

func (c ContainerRuntime) String() string { _ = "STUB: not implemented"; return "" }

type Project struct {
	Name      string
	separator string
}

func NewProject(name string) *Project { _ = "STUB: not implemented"; return nil }

func (p *Project) With(name string) string { _ = "STUB: not implemented"; return "" }

type IpamConfig struct {
	Subnet       string
	IPRange      string
	Gateway      string
	AuxAddresses []string
}

type NixNetwork struct {
	Runtime      ContainerRuntime
	Name         string
	OriginalName string
	Driver       string
	DriverOpts   map[string]string
	External     bool
	Labels       map[string]string
	IpamDriver   string
	IpamConfigs  []IpamConfig
	ExtraOptions []string
}

func (n *NixNetwork) Unit() string { _ = "STUB: not implemented"; return "" }

func (n *NixNetwork) Command() string { _ = "STUB: not implemented"; return "" }

type NixVolume struct {
	Runtime           ContainerRuntime
	Name              string
	Driver            string
	DriverOpts        map[string]string
	External          bool
	Labels            map[string]string
	RemoveOnStop      bool
	RequiresMountsFor []string
}

func (v *NixVolume) Path() string { _ = "STUB: not implemented"; return "" }

func (v *NixVolume) Unit() string { _ = "STUB: not implemented"; return "" }

func (v *NixVolume) Command() string { _ = "STUB: not implemented"; return "" }

// NixContainerSystemdConfig configures the container's systemd config.
// In particular, this allows control of the container restart policy through systemd
// service and unit configs.
//
// Each key-value pair in a map represents a systemd key and its value (e.g., Restart=always).
// Users can provide custom config keys by setting the compose2nix.systemd.* label on the service.
type NixContainerSystemdConfig struct {
	Service ServiceConfig
	Unit    UnitConfig
	// NixOS treats these differently, probably to fix the rename issue in
	// earlier systemd versions.
	// See: https://unix.stackexchange.com/a/464098
	StartLimitBurst *int
}

func NewNixContainerSystemdConfig() *NixContainerSystemdConfig {
	_ = "STUB: not implemented"
	return nil
}

// https://search.nixos.org/options?channel=unstable&from=0&size=50&sort=relevance&type=packages&query=oci-container
type NixContainer struct {
	Runtime       ContainerRuntime
	Name          string
	Image         string
	Environment   map[string]string
	EnvFiles      []string
	Volumes       map[string]string
	Ports         []string
	Labels        map[string]string
	Networks      []string
	DependsOn     []string
	LogDriver     string
	ExtraOptions  []string
	SystemdConfig *NixContainerSystemdConfig
	User          string
	Command       []string
	AutoStart     bool
	SopsSecrets   []string
}

func (c *NixContainer) Unit() string { _ = "STUB: not implemented"; return "" }

// https://docs.docker.com/reference/compose-file/services/#pull_policy
// https://docs.podman.io/en/latest/markdown/podman-build.1.html#pull-policy
type ServicePullPolicy int

const (
	ServicePullPolicyInvalid ServicePullPolicy = iota
	ServicePullPolicyAlways
	ServicePullPolicyNever
	ServicePullPolicyMissing
	ServicePullPolicyBuild
	ServicePullPolicyUnset
)

func NewServicePullPolicy(s string) ServicePullPolicy {
	_ = "STUB: not implemented"
	return *new(ServicePullPolicy)
}

// https://docs.docker.com/reference/compose-file/build/
// https://docs.docker.com/reference/cli/docker/buildx/build/
type NixBuild struct {
	Runtime       ContainerRuntime
	Context       string
	PullPolicy    ServicePullPolicy
	IsGitRepo     bool
	Args          map[string]*string
	Tags          []string
	Dockerfile    string // Relative to context path.
	ContainerName string // Name of the resolved Nix container.
}

func (b *NixBuild) UnitName() string { _ = "STUB: not implemented"; return "" }

func (b *NixBuild) Unit() string { _ = "STUB: not implemented"; return "" }

func (b *NixBuild) Command() string { _ = "STUB: not implemented"; return "" }

type NixContainerConfig struct {
	Version          string
	Project          *Project
	Runtime          ContainerRuntime
	Containers       []*NixContainer
	Builds           []*NixBuild
	Networks         []*NixNetwork
	Volumes          []*NixVolume
	CreateRootTarget bool
	WriteNixSetup    bool
	AutoFormat       bool
	AutoStart        bool
	IncludeBuild     bool
	Option           string
	EnableOption     bool
	SopsConfig       *SopsConfig
}

func (c *NixContainerConfig) HasSopsSecrets() bool { _ = "STUB: not implemented"; return false }

func (c *NixContainerConfig) String() string { _ = "STUB: not implemented"; return "" }

// This should never be hit under normal operation.

// Write writes out the Nix config to the provided Writer.
//
// If the AutoFormat option on this struct is set to "true", this method will
// attempt to format the Nix config by calling "nixfmt" and passing in the
// fully built config via stdin.
func (c *NixContainerConfig) Write(out io.Writer) error { _ = "STUB: not implemented"; return nil }

func rootTarget(runtime ContainerRuntime, project *Project) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *NixContainerConfig) rootTargetTemplateFunc() string { _ = "STUB: not implemented"; return "" }

func (c *NixContainerConfig) configTemplateFunc() *NixContainerConfig {
	_ = "STUB: not implemented"
	return nil
}
