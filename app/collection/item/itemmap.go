package item

import "github.com/kyleu/todoforge/app/util"

func (i *Item) ToMap() util.ValueMap {
	return util.ValueMap{"id": i.ID, "collectionID": i.CollectionID, "name": i.Name, "created": i.Created}
}

func FromMap(m util.ValueMap, setPK bool) (*Item, util.ValueMap, error) {
	ret := &Item{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "id":
			if setPK {
				retID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retID != nil {
					ret.ID = *retID
				}
			}
		case "collectionID":
			retCollectionID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retCollectionID != nil {
				ret.CollectionID = *retCollectionID
			}
		case "name":
			ret.Name, err = m.ParseString(k, true, true)
		default:
			extra[k] = v
		}
		if err != nil {
			return nil, nil, err
		}
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, extra, nil
}
