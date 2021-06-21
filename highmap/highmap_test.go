package highmap

import (
	"fmt"
	"testing"
)

func TestNewHighMap(t *testing.T) {
	m := NewHighMap()
	m.Set(1, Tag{"name","xh"}, Tag{"age","15"},NKey("class","1"))
	m.Set(2, Tag{"name","xh"}, Tag{"home","cn"},NKey("class","1"))
	m.Set(3, Tag{"name","xm"}, Tag{"age","15"},NKey("class","1"))
	m.Set(4, Tag{"name","xm"}, Tag{"home","us"},NKey("class","1"))

	fmt.Println(m.Get(Tag{"name","xh"}, Tag{"age","15"}))
	fmt.Println(m.Get(NKey("class","1")))

}

func BenchmarkHighMap(b *testing.B) {

	m := NewHighMap()
	m.Set(1, Tag{"name","xh"}, Tag{"age","15"},NKey("class","1"))
	m.Set(2, Tag{"name","xh"}, Tag{"home","cn"},NKey("class","1"))
	m.Set(3, Tag{"name","xm"}, Tag{"age","15"},NKey("class","1"))
	m.Set(4, Tag{"name","xm"}, Tag{"home","us"},NKey("class","1"))

	for i := 0; i < b.N; i++ {
		m.Get(NKey("class","1"))
	}
}

func TestNewCounter(t *testing.T) {
	c := NewCounter("name","age","home")
	c.Inc("xh","15","1")
	c.Inc("xh","15","1")
	c.Inc("xh","15","1")
	c.Inc("xm","15","1")
	c.Inc("xm","15","1")
	c.Inc("xx","16","2")
	c.Inc("xh","15","2")
	c.Inc("xh","15","2")
	fmt.Println(c.GetCount(NKey("name","xh")))
	c.IncKeys(NKey("name","xh"))
	fmt.Println(c.GetCount(NKey("name","xh")))
	fmt.Println(c.GetCount(NKey("name","xm"),NKey("home","1")))


	c.Show()
}
