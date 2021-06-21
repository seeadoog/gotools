package highmap

import (
	"fmt"
	"math"
	"strings"
)

type Tag struct {
	Key string
	Val string
}


func NKey(key,val string) Tag {
	return Tag{
		Key: key,
		Val: val,
	}
}

func NKeys(kvs ...string)[]Tag{
	if len(kvs) &1 != 0{
		panic("kvs length must not be odd")
	}
	tags := make([]Tag,0, len(kvs)/2)
	for i := 0; i < len(kvs); i+=2 {
		tags = append(tags,Tag{kvs[i],kvs[i+1]})
	}
	return tags
}

type Map interface {
	Set(val interface{},keyVals ...Tag)
	Get(keys ...Tag)[]interface{}
}

type Val struct {
	Key string
	V interface{}
	Tags []Tag
}

func (v Val) String() string {
	return fmt.Sprintf("key=%s,v=%v,tags=%v",v.Key,v.V,v.Tags)
}

type HighMap struct {
	data map[string]*Val
	indexes map[string]map[string]map[string]bool//{1,2,3}  {4,5,6}, 1
}

func NewHighMap()*HighMap{
	return &HighMap{indexes: map[string]map[string]map[string]bool{},data: map[string]*Val{}}
}

func Keys(key []Tag)string{
	length := 0
	for _, k := range key {
		length += len(k.Key)
		length += len(k.Val)
		length+=2
	}
	sb := strings.Builder{}
	sb.Grow(length)
	for _, k := range key {
		sb.WriteString(k.Key)
		sb.WriteByte('_')
		sb.WriteString(k.Val)
		sb.WriteByte('.')
	}
	return sb.String()
}


func (h *HighMap) Set(val interface{}, tags ...Tag) {
	key := Keys(tags)
	for _, keyVal := range tags {
		vmap ,ok := h.indexes[keyVal.Key]
		if !ok{
			vmap = map[string]map[string]bool{}
			h.indexes[keyVal.Key] = vmap
		}
		idMap,ok := vmap[keyVal.Val]
		if !ok{
			idMap = map[string]bool{}
			vmap[keyVal.Val] = idMap
		}
		idMap[key] = true
	}
	h.data[key] = &Val{Key: key,V: val,Tags: tags}
}

func (h *HighMap) Get(keys ...Tag) (res []*Val ){
	indexes := make([]map[string]bool, len(keys))
	minIdx := 0
	minLen := math.MaxInt64
	for i, key := range keys {

		vmap ,ok := h.indexes[key.Key]
		if !ok{
			return nil
		}
		ids ,ok := vmap[key.Val]
		if !ok{
			return nil
		}

		if len(ids) < minLen{
			minLen = len(ids)
			minIdx = i
		}

		indexes[i] = ids
	}

	if len(indexes) == 0{
		return nil
	}
	indexes[0],indexes[minIdx] = indexes[minIdx],indexes[0]
	for id, _ := range indexes[0] {
		selected := true
		for _, m := range indexes[1:] {
			if !m[id]{
				selected = false
				break
			}
		}
		if selected{
			res = append(res,h.data[id])
		}
	}
	return res
}


/*

//call id   id chan msg

// appid:123456 [1,2,3]
// uid:45566 [2]

 */
