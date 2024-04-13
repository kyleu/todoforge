// Package util - Content managed by Project Forge, see [projectforge.md] for details.
package util

import (
	"time"

	"github.com/araddon/dateparse"
	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
)

func TimeTruncate(t *time.Time) *time.Time {
	if t == nil {
		return nil
	}
	ret := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return &ret
}

func TimeCurrent() time.Time {
	return time.Now()
}

func TimeCurrentP() *time.Time {
	ret := TimeCurrent()
	return &ret
}

func TimeToday() *time.Time {
	return TimeTruncate(TimeCurrentP())
}

func TimeCurrentUnix() int64 {
	return TimeCurrent().Unix()
}

func TimeCurrentMillis() int64 {
	return TimeCurrent().UnixMilli()
}

func TimeCurrentNanos() int64 {
	return TimeCurrent().UnixNano()
}

func TimeRelative(t *time.Time) string {
	if t == nil {
		return "<never>"
	}
	return humanize.Time(*t)
}

func TimeToMap(t time.Time) map[string]any {
	return map[string]any{"epoch": t.UnixMilli(), "iso8601": t.Format("2006-01-02T15:04:05-0700")}
}

func TimeToString(d *time.Time, fmt string) string {
	if d == nil {
		return ""
	}
	return d.Format(fmt)
}

func TimeFromString(s string) (*time.Time, error) {
	if s == "" {
		return nil, nil
	}
	ret, err := dateparse.ParseLocal(s)
	if err != nil {
		return nil, errors.Errorf("invalid date string [%s]", s)
	}
	return &ret, nil
}

func TimeFromStringSimple(s string) *time.Time {
	ret, _ := TimeFromString(s)
	return ret
}

func TimeFromStringFmt(s string, fmt string) (*time.Time, error) {
	if s == "" {
		return nil, nil
	}
	ret, err := time.Parse(fmt, s)
	if err != nil {
		return nil, errors.Errorf("invalid date string [%s], expected [%s]", s, fmt)
	}
	return &ret, nil
}
