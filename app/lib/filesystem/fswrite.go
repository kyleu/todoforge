package filesystem

import (
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/kyleu/todoforge/app/util"
)

type Writer interface {
	io.Writer
	io.WriterAt
	io.Closer
}

func (f *FileSystem) WriteJSONFile(path string, x any, mode FileMode, overwrite bool) error {
	return f.WriteFile(path, append(util.ToJSONBytes(x, true), '\n'), mode, overwrite)
}

func (f *FileSystem) WriteFile(path string, content []byte, mode FileMode, overwrite bool) error {
	p := f.getPath(path)
	s, _ := f.f.Stat(p)
	if s != nil && s.Size() > 0 && !overwrite {
		return errors.Errorf("file [%s] exists, will not overwrite", p)
	}
	if mode == 0 {
		if s == nil {
			mode = DefaultMode
		} else {
			mode = FileMode(s.Mode())
		}
	}
	dd := filepath.Dir(path)
	err := f.CreateDirectory(dd)
	if err != nil {
		return errors.Wrapf(err, "unable to create data directory [%s]", dd)
	}
	file, err := f.f.Create(p)
	if err != nil {
		return errors.Wrapf(err, "unable to create file [%s]", p)
	}
	err = f.SetMode(p, mode)
	if err != nil {
		return errors.Wrapf(err, "unable to set mode [%d] for file [%s]", mode, p)
	}
	defer func() { _ = file.Close() }()
	_, err = file.Write(content)
	if err != nil {
		return errors.Wrapf(err, "unable to write content to file [%s]", p)
	}
	return nil
}

func (f *FileSystem) FileWriter(fn string, createIfNeeded bool, appendMode bool) (Writer, error) {
	p := f.getPath(fn)
	if createIfNeeded && !f.Exists(p) {
		_, err := f.f.Create(p)
		if err != nil {
			return nil, err
		}
	}
	mode := os.O_WRONLY
	if appendMode {
		mode = os.O_APPEND | os.O_WRONLY
	}
	return f.f.OpenFile(p, mode, os.ModeAppend)
}
