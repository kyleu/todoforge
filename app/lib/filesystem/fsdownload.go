package filesystem

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/todoforge/app/util"
)

func (f *FileSystem) Download(ctx context.Context, url string, path string, overwrite bool, _ util.Logger) (int, error) {
	if !strings.HasPrefix(url, "https://") {
		return 0, errors.New("only [https] requests are supported")
	}
	if f.Exists(path) && !overwrite {
		return 0, errors.Errorf("file [%s] exists", path)
	}
	rsp, err := util.NewHTTPRequest(ctx, http.MethodGet, url).Run()
	if err != nil {
		return 0, errors.Wrapf(err, "unable to load url [%s]", url)
	}
	defer func() {
		_ = rsp.Body.Close()
	}()
	if rsp.StatusCode != http.StatusOK {
		return 0, errors.Errorf("response from url [%s] has status [%d], expected [200]", url, rsp.StatusCode)
	}
	w, err := f.FileWriter(path, true, false)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to open file at path [%s]", path)
	}
	x, err := io.Copy(w, rsp.Body)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to save response from url [%s] to file [%s]", url, path)
	}
	return int(x), nil
}
