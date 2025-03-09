package adapters

import (
	"goAPI/src/db/handlers"
)

type IBaseAdapter interface {
	GetAll() any
	Get(id uint) any
	DeleteAll() any
	Delete(id uint) any
}

type BaseAdapter struct {
	IBaseAdapter
	handler handlers.IBaseHandler
}

func (b BaseAdapter) GetAll() any {
	return b.handler.GetAll()
}

func (b BaseAdapter) Get(id uint) any {
	return b.handler.Get(id)
}

func (b BaseAdapter) DeleteAll() any {
	return b.handler.DeleteAll()
}

func (b BaseAdapter) Delete(id uint) any {
	return b.handler.Delete(id)
}

func NewBaseAdapter(handler handlers.IBaseHandler) BaseAdapter {
	return BaseAdapter{handler: handler}
}
