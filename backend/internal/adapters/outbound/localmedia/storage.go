// Package localmedia contains filesystem media storage adapters.
package localmedia

// Storage is the local filesystem adapter boundary for media bytes.
//
// Phase 9 implements application.MediaStorage on this type.
type Storage struct {
	rootDir string
}

// NewStorage records the root directory without creating files during composition.
func NewStorage(rootDir string) *Storage {
	return &Storage{
		rootDir: rootDir,
	}
}

// RootDir returns the configured storage root for composition smoke checks.
func (s *Storage) RootDir() string {
	if s == nil {
		return ""
	}

	return s.rootDir
}
