package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/compose-spec/compose-go/v2/types"
)

const (
	// https://www.freedesktop.org/software/systemd/man/latest/systemd-system.conf.html#DefaultTimeoutStartSec=
	defaultSystemdStopTimeout = 90 * time.Second
)

var (
	// We purposefully do not support no/yes or 0/1 for false/true to avoid
	// ambiguity. Nix supports strings and bools for boolean keys anyways.
	systemdTrue  = []string{"true", "on"}
	systemdFalse = []string{"false", "off"}

	// Examples:
	// compose2nix.systemd.service.RuntimeMaxSec=100
	// compose2nix.systemd.unit.StartLimitBurst=10
	systemdLabelRegexp = regexp.MustCompile(fmt.Sprintf(`%s\.systemd\.(service|unit)\.(\w+)`, composeLabelPrefix))
)

// https://www.freedesktop.org/software/systemd/man/latest/systemd.syntax.html
func parseSystemdValue(v string) any {
	_ = "STUB: not implemented"
	return *

	// Number
	new(any)
}

// Boolean

// String
// Remove all quotes from the string.

// TODO(aksiksi): Add support for repeated keys.
type ServiceConfig struct {
	// Map for generic options.
	Options map[string]any
}

func (s *ServiceConfig) Set(key string, value any) { _ = "STUB: not implemented"; return }

// TODO(aksiksi): Add support for repeated keys.
type UnitConfig struct {
	After             []string
	Requires          []string
	PartOf            []string
	UpheldBy          []string
	WantedBy          []string
	RequiresMountsFor []string
	// Map for generic options.
	Options map[string]any
}

func (u *UnitConfig) Set(key string, value any) { _ = "STUB: not implemented"; return }

func (c *NixContainerSystemdConfig) ParseRestartPolicy(service *types.ServiceConfig, runtime ContainerRuntime) error {
	_ = "STUB: not implemented"
	return nil

	// https://docs.docker.com/compose/compose-file/compose-file-v2/#restart
}

// Both of these match the systemd restart options.

// We don't have an equivalent in systemd. Podman does the same thing.

// The newer "deploy" config will always override the legacy "restart" config.
// https://docs.docker.com/compose/compose-file/compose-file-v3/#restart_policy

// If unset, defaults to "any".

// TODO(aksiksi): Investigate if StartLimitIntervalSec lines up with Compose's "window".

// This simulates the default behavior of Docker. Basically, Docker will restart
// the container with a sleep period of 100ms. This sleep period is doubled until a
// maximum of 1 minute.
// See: https://docs.docker.com/reference/cli/docker/container/run/#restart

// 2^(9 attempts) = 512 (* 100ms) ~= 1 minute

func (c *NixContainerSystemdConfig) ParseSystemdLabels(service *types.ServiceConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *NixContainerSystemdConfig) Sort() { _ = "STUB: not implemented"; return }
