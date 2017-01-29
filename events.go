package main

import "container/list"

var (
	handlers map[string]*list.List = make(map[string]*list.List)
)

func RegisterEvent(event string, callback func(interface{})) {
	var l *list.List = handlers[event]
	if l == nil {
		l = list.New()
		handlers[event] = l
	}
	l.PushBack(callback)
}

func CallEvent(event string, value interface{}) {
	var l *list.List = handlers[event]
	if l != nil {
		for handler := l.Front(); handler != nil; handler = handler.Next() {
			handler.Value.(func(interface{}))(value)
		}
	}
}