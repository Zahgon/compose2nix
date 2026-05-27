package main

import (
	"context"
	"regexp"
	"time"

	"github.com/compose-spec/compose-go/v2/types"
)

const (
	composeLabelPrefix = "compose2nix"
)

func parseNixContainerLabels(c *NixContainer, sopsConfig *SopsConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// This will be handled later.

func composeEnvironmentToMap(env types.MappingWithEquals) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Skip empty env variables.

func portConfigsToPortStrings(portConfigs []types.ServicePortConfig) []string {
	_ = "STUB: not implemented"
	return nil
}

// Dummy interface that allows patching os.Getwd() in tests.
type getWorkingDir interface {
	GetWd() (string, error)
}

func (g *Generator) GetRootPath() (string, error) { _ = "STUB: not implemented"; return "", nil }

type Generator struct {
	Project                 *Project
	Runtime                 ContainerRuntime
	Inputs                  []string
	EnvFiles                []string
	RootPath                string
	IncludeEnvFiles         bool
	EnvFilesOnly            bool
	IgnoreMissingEnvFiles   bool
	ServiceInclude          *regexp.Regexp
	AutoStart               bool
	UseComposeLogDriver     bool
	GenerateUnusedResources bool
	CheckSystemdMounts      bool
	CheckBindMounts         bool
	UseUpheldBy             bool
	RemoveVolumes           bool
	NoCreateRootTarget      bool
	AutoFormat              bool
	WriteHeader             bool
	NoWriteNixSetup         bool
	DefaultStopTimeout      time.Duration
	IncludeBuild            bool
	GetWorkingDir           getWorkingDir
	OptionPrefix            string
	EnableOption            bool
	SopsConfig              *SopsConfig
	WarningsAsErrors        bool

	serviceToContainerName map[string]string
	rootPath               string
}

func (g *Generator) Run(ctx context.Context) (*NixContainerConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transform env files into absolute paths. This ensures that we can compare
// them to Compose env files when building Nix containers.

// Always override the project we have set. It could have been normalized or, more commonly,
// pulled from one of the Compose files' top-level "name" setting.

// Construct a map of service to container name.

// Post-process any Compose settings that require the full state.

func (g *Generator) postProcess(containers []*NixContainer, networks []*NixNetwork, volumes []*NixVolume) ([]*NixNetwork, []*NixVolume) {
	_ = "STUB: not implemented"
	// Drop any networks that are unused or external.
	return nil, nil
}

// Drop any volumes that are unused or external.

func healthCheckCommandToString(cmd []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Health check.
// https://docs.docker.com/compose/compose-file/05-services/#healthcheck
func parseHealthCheck(c *NixContainer, service types.ServiceConfig, runtime ContainerRuntime) error {
	_ = "STUB: not implemented"
	return nil
}

// Figure out if the Dockerfile health check is disabled.

// https://docs.podman.io/en/latest/markdown/podman-run.1.html#health-startup-interval-interval

func (g *Generator) handleVolumesForService(service types.ServiceConfig, volumeMap map[string]*NixVolume, c *NixContainer) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle tmpfs volumes first.

// compose-go does not differentiate between short and long syntax, so we'll use long-only fields to try to tell the difference.

// Handle long syntax by passing --mount flag to the backend
// Docker: https://docs.docker.com/reference/cli/docker/container/run/#mount
// Podman: https://docs.podman.io/en/latest/markdown/podman-run.1.html#mount-type-type-type-specific-option

// Used to force generation of volume creation.
// Empty strings are filtered from the volumes attribute.

// Replace the Compose volume name with the actual Docker volume
// name (i.e., potentially prefixed with project).
//
// This is what we'll use to refer to the volume in the generated
// container config.

// Add systemd dependencies on volume(s).

// This is a bind mount.

// Replace the source path in the volume string.

// Handle tmpfs short syntax.
// https://docs.docker.com/reference/compose-file/services/#tmpfs

func (g *Generator) checkOrWarn(format string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) buildNixContainer(service types.ServiceConfig, networkMap map[string]*NixNetwork, volumeMap map[string]*NixVolume) (*NixContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is the NixOS default

// Env files provided via CLI.

// Env files set on the Compose service.

// It's possible that an env file that was passed in via CLI is
// _also_ present in the Compose service (as an env_file).

// Figure out explicit dependencies for this container.
//
// TODO(aksiksi): Support the long syntax.
// https://docs.docker.com/compose/compose-file/05-services/#long-syntax-1

// If the container is connected to a network, it's counted as being in a bridge network.
// We need to know this to be able to determine if we can configure a network alias.
//
// NOTE(aksiksi): Is this even correct?

// https://docs.docker.com/compose/compose-file/05-services/#network_mode
// https://docs.podman.io/en/latest/markdown/podman-run.1.html#network-mode-net

// https://docs.podman.io/en/latest/markdown/podman-run.1.html#network-mode-net

// TODO(aksiksi): Can we even do anything for Docker?

// Convert the Compose "service" network mode to a "container" network mode.

// container:[name] mode is supported by both Docker and Podman.
// This container could be external, so we can't fail if it doesn't exist in this Compose
// project.

// TODO(aksiksi): Should we even be doing this?

// Add systemd dependencies on network.

// If we don't have any additional config set on this network, stop here.

// Aliases are scoped to all networks - I think?

// Aliases are scoped to the current network.
// https://docs.podman.io/en/latest/markdown/podman-run.1.html#network-mode-net

// Below, we fallback to using --ip/--ip6 if a single network is
// specified. This aligns with Docker behavior.

// NOTE(aksiksi): Docker might actually support network-scoped IPs.
//
// But, we need to think about this carefully because the flags seem to be
// order-depdendent. That is, the order of --network and --ip/--ip6 needs
// to be maintained. Since we sort the ExtraOptions array, the ordering
// would break without some changes.

// Allow other containers to use service name as an alias.
//
// In the case of Podman, this alias applies to all networks the container is a part of.
// Network-scoped aliases are handled below.
//
// See: https://docs.podman.io/en/latest/markdown/podman-run.1.html#network-alias-alias

// https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities

// Special case: Nvidia CDI devices use short syntax.

// Otherwise, we default to using the long syntax.
// https://docs.docker.com/reference/cli/docker/container/run/#device

// https://docs.docker.com/compose/compose-file/05-services/#ulimits
// https://docs.docker.com/reference/cli/docker/container/run/#ulimit
// https://docs.podman.io/en/latest/markdown/podman-run.1.html#ulimit-option
//
// NOTE: Single, Soft, and Hard can all be 0 due to Go zero-values, so
// the single vs. soft/hard form is ambiguous when all are 0 (upstream
// compose-go limitation). We emit the short form when soft == hard
// since --ulimit=X=V is equivalent to --ulimit=X=V:V.

// https://docs.docker.com/compose/compose-file/05-services/#extra_hosts
// https://github.com/compose-spec/compose-spec/blob/master/spec.md#extra_hosts
// https://docs.docker.com/engine/reference/commandline/run/#add-host
// https://docs.podman.io/en/latest/markdown/podman-run.1.html#add-host-host-ip

// We can get upto two IPs per hostname: one v4 and one v6.
// See: https://github.com/compose-spec/compose-go/pull/563

// https://docs.docker.com/compose/compose-file/05-services/#group_add
// https://docs.docker.com/engine/reference/commandline/run/#group-add
// https://docs.podman.io/en/latest/markdown/podman-run.1.html#group-add-group

// https://docs.docker.com/compose/compose-file/05-services/#ipc
// https://docs.docker.com/engine/reference/run/#ipc-settings---ipc
// https://docs.podman.io/en/latest/markdown/podman-run.1.html#ipc-ipc

// Convert the Compose "service" IPC mode to a "container" IPC mode.
// Note: compose-go automatically validates the service exists and adds it as a dependency.

// https://docs.docker.com/compose/compose-file/05-services/#sysctls
// https://docs.docker.com/engine/reference/commandline/run/#sysctl
// https://docs.podman.io/en/latest/markdown/podman-run.1.html#sysctl-name-value

// Compose defaults to "json-file", so we'll treat _any_ "json-file" setting as a default.
// Users can override this behavior via CLI.
//
// https://docs.docker.com/config/containers/logging/configure/
// https://docs.podman.io/en/latest/markdown/podman-run.1.html#log-driver-driver

// New logging setting always overrides the legacy setting.
// https://docs.docker.com/compose/compose-file/compose-file-v3/#logging

// Deploy resources configuration.
// https://docs.docker.com/compose/compose-file/deploy/#resources

// Name is misleading - this actually is the exact number passed in with "cpus".

// CPU reservation is a Docker Swarm option.

// CDI GPU support.

// Pass in all GPUs in CDI format.
//
// TODO(aksiksi): Maybe we can do something better here?

// Restart policy.

// Override systemd stop timeout to match Docker/Podman default of 10 seconds.
// https://docs.podman.io/en/latest/markdown/podman-stop.1.html
//
// Users can always override this by setting per-service Compose labels, or by passing in a CLI
// flag.

// We only set a timeout if it's not the same as the systemd default.

// Sort slices now that we're done processing the container.

// Add systemd dependency on root target.
//
// NOTE(aksiksi): We must check auto-start here because the root target
// could have auto-start set, which would implicitly bring up the container.

// Unfortunately, UpheldBy does not work as expected, so we're keeping it
// behind a flag.
//
// Refer to the "Known Issues" section in the README for details.

// Set UpheldBy for this service's dependencies. This ensures that, when
// the dependency comes up, this container will also be started - and
// continuously restarted with backoff - until it comes up.
//
// See: https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#Upholds=
//
// Why do we need to do this? Because, by default, systemd does not
// attempt to start failed dependent units when the parent (dependency)
// comes up. See: https://github.com/systemd/systemd/issues/1312.
//
// For further discussion, see: https://github.com/aksiksi/compose2nix/issues/19

// systemd configs provided via labels always override everything else.

func (g *Generator) parseServiceBuild(service types.ServiceConfig, c *NixContainer) (*NixBuild, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Process this as a Git repo.

// If no image is set on the service, we'll define an image name based
// on the container name.

// Always use the image name as a tag.

// Apply additional tags on top, prefixed with the image name.

// Set the image on the container.

// Podman automatically prepends a registry name of "localhost" to any
// tag we set.
//
// See: https://docs.podman.io/en/latest/markdown/podman-build.1.html#tag-t-imagename

// Add dependency on build systemd service.

func (g *Generator) buildNixContainers(composeProject *types.Project, networkMap map[string]*NixNetwork, volumeMap map[string]*NixVolume) (containers []*NixContainer, builds []*NixBuild, _ error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (g *Generator) containerNameToService(name string) string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Generator) networkNameToService(name string) string { _ = "STUB: not implemented"; return "" }

func (g *Generator) volumeNameToService(name string) string { _ = "STUB: not implemented"; return "" }

func (g *Generator) buildNixNetworks(composeProject *types.Project) ([]*NixNetwork, map[string]*NixNetwork) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IPAM configuration.
// https://docs.docker.com/compose/compose-file/06-networks/#ipam
// https://docs.docker.com/reference/cli/docker/network/create/
// https://docs.podman.io/en/latest/markdown/podman-network-create.1.html

// If driver is set to "default", we'll omit it and fallback to the
// runtime.

func (g *Generator) buildNixVolumes(composeProject *types.Project) ([]*NixVolume, map[string]*NixVolume) {
	_ = "STUB: not implemented"
	return nil, nil
}
