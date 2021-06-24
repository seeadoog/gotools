package highmap

import (
	"fmt"
	"strings"
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

func TestEs(t *testing.T) {
	es := &es{hm: NewHighMap()}
	es.Insert("hello world, chenjian is a big word")
	es.Insert("hello world, dajian is a big word")
	es.Insert("hello world, shengqi is a big word")
	es.Insert("hello world, shengqid is a big word")
	es.Insert("hello world, shengqi has a fit")
	es.Insert("hello world, shengqib is a big word")

	for i, s := range es.Search("word") {
		fmt.Println(i,":",s)
	}
}

type splitWordFunc func(b string)

type es struct {
	hm *HighMap
	sf splitWordFunc
	id string
}

func (es *es)Insert(text string){
	tags := make([]Tag,0,0)
	for _, s := range strings.Split(text, " ") {
		ss := strings.TrimSpace(s)
		if ss == ""{
			continue
		}
		tags = append(tags,Tag{
			Key: ss,
			Val: "",
		})
	}
	es.hm.Set(text,tags...)
}


func (es *es)Search(text string)(res []string){

	for _, val := range es.hm.Get(NKey(text, "")) {
		res = append(res, val.V.(string))
	}
	return res
}

func TestKeys(t *testing.T) {
	c1 := make(chan int,1)
	c2 := make(chan int,1)
	c:= make(chan string,0)
	go func() {
		for{
			<- c1
			fmt.Println("on c1")
			c<- "on c1"
			c2 <- 1
		}
	}()

	go func() {
		for{
			<- c2
			fmt.Println("on c2")
			c<- "on c2"
			c1 <- 1
		}
	}()
	c1<-1
	for{
		fmt.Println(<-c)
	}
}
