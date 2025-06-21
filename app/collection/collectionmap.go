package collection

import "github.com/kyleu/todoforge/app/util"

func (c *Collection) ToMap() util.ValueMap {
	return util.ValueMap{"id": c.ID, "name": c.Name, "created": c.Created}
}

func CollectionFromMap(m util.ValueMap, setPK bool) (*Collection, util.ValueMap, error) {
	ret := &Collection{}
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

func (c *Collection) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "id", V: c.ID}, {K: "name", V: c.Name}, {K: "created", V: c.Created}}
	return util.NewOrderedMap(false, 4, pairs...)
}
