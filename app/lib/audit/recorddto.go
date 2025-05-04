package audit

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app/util"
)

var (
	recordTable         = "audit_record"
	recordTableQuoted   = fmt.Sprintf("%q", recordTable)
	recordColumns       = []string{"id", "audit_id", "t", "pk", "changes", "metadata", "occurred"}
	recordColumnsQuoted = util.StringArrayQuoted(recordColumns)
	recordColumnsString = util.StringJoin(recordColumnsQuoted, ", ")
)

type recordRow struct {
	ID       uuid.UUID `db:"id"`
	AuditID  uuid.UUID `db:"audit_id"`
	T        string    `db:"t"`
	PK       string    `db:"pk"`
	Changes  string    `db:"changes"`
	Metadata string    `db:"metadata"`
	Occurred time.Time `db:"occurred"`
}

func (r *recordRow) ToRecord() *Record {
	if r == nil {
		return nil
	}
	changesArg := util.Diffs{}
	_ = util.FromJSON([]byte(r.Changes), &changesArg)
	metadataArg, _ := util.FromJSONMap([]byte(r.Metadata))
	return &Record{ID: r.ID, AuditID: r.AuditID, T: r.T, PK: r.PK, Changes: changesArg, Metadata: metadataArg, Occurred: r.Occurred}
}

type recordRows []*recordRow

func (x recordRows) ToRecords() Records {
	return lo.Map(x, func(r *recordRow, _ int) *Record {
		return r.ToRecord()
	})
}
