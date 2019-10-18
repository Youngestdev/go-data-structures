package List

type List interface {
	Insert(i int, e interface{}) error
	Delete(i int) (interface{}, error)
	Get(i int) (interface{}, error)
	Put(i int, e interface{}) error
	Index(e interface{}) (int, bool)
	Slice(i, j int) (List, error)
	Equal(t []interface{}) bool
}

type Iterator interface {
	Reset()
	Done() bool
	Next() (interface{}, bool)
}
