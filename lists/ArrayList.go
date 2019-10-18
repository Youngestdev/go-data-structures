package List

import "errors"

type ArrayList struct {
	arrList []interface{}
}

func (l *ArrayList) Size() int {
	return len(l.arrList) - 1
}

func (l *ArrayList) Insert(elem interface{}) (interface{}, error) {
	l.arrList = append(l.arrList, elem)
	res := l.arrList
	return res, nil
}

// Currently doesn't resize array - I think I fixed it in line 10.
func (l *ArrayList) Delete(i int) (interface{}, error) {
	if i < 0 || l.Size() < i {
		return nil, errors.New("error: Index number is void")
	}
	a := l.arrList
	a[i] = a[len(a)-1]
	a = a[:len(a)-1]
	res := a
	return res, nil
}

func (l *ArrayList) Get(i int) (interface{}, error) {
	if i < 0 || l.Size() < i {
		return nil, errors.New("get: index is less than the size of the list")
	}
	return l.arrList[i], nil
}

func (l *ArrayList) Put(i int, e interface{}) interface{} {
	if i >= 0 || l.Size() > i {
		l.arrList[i] = e
	}
	return nil
}

func (l *ArrayList) Index(e interface{}) interface{} {
	for i := range l.arrList {
		if l.arrList[i] == e {
			return i
		}
	}
	return "Nothing was found"
}

//TODO
// Reference - https://github.com/golang/go/wiki/SliceTricks#insert
func (l *ArrayList) Slice(i, j []interface{}) []interface{} {
	return nil
}

// TODO
func (l *ArrayList) Equals(s []interface{}) bool {
	//return l.arrList == s
	return false
}
