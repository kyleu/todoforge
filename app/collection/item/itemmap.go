// Package item - Content managed by Project Forge, see [projectforge.md] for details.
package item

import "github.com/kyleu/todoforge/app/util"

func FromMap(m util.ValueMap, setPK bool) (*Item, error) {
	ret := &Item{}
	var err error
	if setPK {
		retID, e := m.ParseUUID("id", true, true)
		if e != nil {
			return nil, e
		}
		if retID != nil {
			ret.ID = *retID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	retCollectionID, e := m.ParseUUID("collectionID", true, true)
	if e != nil {
		return nil, e
	}
	if retCollectionID != nil {
		ret.CollectionID = *retCollectionID
	}
	ret.Name, err = m.ParseString("name", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
