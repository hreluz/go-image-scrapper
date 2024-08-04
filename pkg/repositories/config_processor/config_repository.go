package repoconfigprocessor

import "github.com/hreluz/images-scrapper/pkg/models/config"

type ConfigProcessorRepository interface {
	Save(cp *config.ConfigProcessor) error
	SaveAll(cps []*config.ConfigProcessor) error
	GetAll() ([]*config.ConfigProcessor, error)
	GetById(id int) (*config.ConfigProcessor, error)
}
