// Content managed by Project Forge, see [projectforge.md] for details.
package filesystem

import "github.com/kyleu/todoforge/app/util"

type FileLoader interface {
	Root() string
	Clone() FileLoader
	PeekFile(path string, maxSize int) ([]byte, error)
	Size(path string) int
	ReadFile(path string) ([]byte, error)
	CreateDirectory(path string) error
	WriteFile(path string, content []byte, mode FileMode, overwrite bool) error
	CopyFile(src string, tgt string) error
	CopyRecursive(src string, tgt string, ignore []string, logger util.Logger) error
	Move(src string, tgt string) error
	ListFiles(path string, ignore []string, logger util.Logger) FileInfos
	ListFilesRecursive(path string, ignore []string, logger util.Logger) ([]string, error)
	ListTree(cfg util.ValueMap, path string, ignore []string, logger util.Logger, tags ...string) (*Tree, error)
	ListJSON(path string, ignore []string, trimExtension bool, logger util.Logger) []string
	ListExtension(path string, ext string, ignore []string, trimExtension bool, logger util.Logger) []string
	ListDirectories(path string, ignore []string, logger util.Logger) []string
	Walk(path string, ign []string, fn func(fp string, info *FileInfo, err error) error) error
	Stat(path string) (*FileInfo, error)
	SetMode(path string, mode FileMode) error
	Exists(path string) bool
	IsDir(path string) bool
	Remove(path string, logger util.Logger) error
	RemoveRecursive(pt string, logger util.Logger) error
	String() string
}
