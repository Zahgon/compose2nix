package main

type SopsConfig struct {
	FilePath string
	secrets  map[string]bool
}

func NewSopsConfig(filePath string) *SopsConfig { _ = "STUB: not implemented"; return nil }

func (s *SopsConfig) LoadSecrets() error { _ = "STUB: not implemented"; return nil }

func (s *SopsConfig) extractSecrets(data map[string]any, prefix string) {
	_ = "STUB: not implemented"
	return
}

func (s *SopsConfig) HasSecret(name string) bool { _ = "STUB: not implemented"; return false }
