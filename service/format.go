package service

import (
	"github.com/Shumodan/go-bookshelf/model"
	"github.com/Shumodan/go-bookshelf/mycontext"
)

// FormatService is a service for managing master data such as format and category.
type FormatService struct {
	context mycontext.Context
}

// NewFormatService is constructor.
func NewFormatService(context mycontext.Context) *FormatService {
	return &FormatService{context: context}
}

// FindAllFormats returns the list of all formats.
func (m *FormatService) FindAllFormats() *[]model.Format {
	rep := m.context.GetRepository()
	format := model.Format{}
	result, err := format.FindAll(rep)
	if err != nil {
		m.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}
