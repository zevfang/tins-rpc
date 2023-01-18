package common

import (
	"encoding/json"
	"fmt"

	pp "github.com/emicklei/proto"
)

type decoder struct {
	d *Definitions
	m *pp.Message
	r map[string]interface{}
}

func NewDecoder(d *Definitions) *decoder {
	return &decoder{
		d: d,
		r: map[string]interface{}{},
	}
}

func (d *decoder) DecodeAll() map[string]string {
	data := make(map[string]string)
	for s, _ := range d.d.msgs {
		mkv, _ := d.Decode(s)
		byteList, _ := json.Marshal(&mkv)
		data[s] = FormatJSON(byteList)
		d.r = make(map[string]interface{}) //TODO 清空 待优化
	}
	return data
}

func (d *decoder) Decode(m string) (map[string]interface{}, error) {
	message, ok := d.d.Message(m)
	if !ok {
		return nil, fmt.Errorf("no definition found for message [%s]", m)
	}
	for _, each := range message.Elements {
		if f, ok := each.(*pp.NormalField); ok {
			d.decodeNormalField(f)
		}
		//if f, ok := each.(*pp.MapField); ok {
		//	d.decodeMapField(f, wire)
		//}
	}
	return d.r, nil
}

func (d *decoder) decodeNormalField(f *pp.NormalField) {
	if "string" == f.Type {
		d.add(f.Name, "Test", f.Repeated, false)
	}
	if "int64" == f.Type || "uint64" == f.Type || "int32" == f.Type || "uint32" == f.Type || "float" == f.Type {
		d.add(f.Name, 100, f.Repeated, false)
	}
	if "bool" == f.Type {
		d.add(f.Name, true, f.Repeated, false)
	}
	//if _, ok := d.d.Message(d.p, f.Type); ok {
	//	d.decodeNormalFieldMessage(f)
	//}
}

func (d *decoder) decodeMapField(f *pp.MapField) error {
	return nil
}

func (d *decoder) add(key string, value interface{}, repeated bool, isMap bool) {
	if repeated {
		if val, ok := d.r[key]; ok {
			maps := val.([]interface{})
			maps = append(maps, value)
			d.r[key] = maps
		} else {
			d.r[key] = []interface{}{value}
		}
	} else if isMap {
		if val, ok := d.r[key]; ok {
			// map exists
			outMap := val.(map[string]interface{})
			inMap := value.(map[string]interface{})
			for k, v := range inMap {
				outMap[k] = v
			}
			// needed?
			d.r[key] = outMap
		} else {
			// map did not exist
			outMap := map[string]interface{}{}
			inMap := value.(map[string]interface{})
			for k, v := range inMap {
				outMap[k] = v
			}
			d.r[key] = outMap
		}
	} else {
		d.r[key] = value
	}
}
