package main

// mapToKeyValArray converts a map into a _sorted_ list of KEY=VAL entries.
func mapToKeyValArray(m map[string]string) []string { _ = "STUB: not implemented"; return nil }

func mapToRepeatedKeyValFlag(flagName string, m map[string]string) []string {
	_ = "STUB: not implemented"
	return nil
}

func sliceToStringArray(s []string) string { _ = "STUB: not implemented"; return "" }

// We purposefully do not use %q to avoid Go's built-in string escaping.
// Otherwise, we'd escape " characters a 2nd time for Nix.

// ReadEnvFiles reads the given set of env files into a list of KEY=VAL entries.
//
// If mergeWithEnv is set, the running env is merged with the provided env files. Any
// duplicate variables will be overridden by the running env.
//
// If ignoreMissing is set, any missing env files will be ignored. This is useful for cases
// where an env file is not available during conversion to Nix.
func ReadEnvFiles(envFiles []string, mergeWithEnv, ignoreMissing bool) (env []string, _ error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// formatNixCode will format Nix code by calling 'nixfmt' and passing in the
// given code via stdin.
func formatNixCode(contents []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// Check for existence of 'nixfmt' in $PATH.
	return nil, nil
}

// Overwrite contents with formatted output.
