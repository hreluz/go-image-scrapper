package serviceconfigprocessor

import (
	"github.com/hreluz/images-scrapper/pkg/models/config"
	repoconfigprocessor "github.com/hreluz/images-scrapper/pkg/repositories/config_processor"
)

type ConfigProcessorService struct {
	ConfigProcessorRepository repoconfigprocessor.ConfigFileRepository
}

func (cps *ConfigProcessorService) GetAll() ([]*config.ConfigProcessor, error) {
	configs, err := cps.ConfigProcessorRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return configs, nil
}

func (cps *ConfigProcessorService) Save(cp *config.ConfigProcessor) error {
	err := cps.ConfigProcessorRepository.Save(cp)

	if err != nil {
		return err
	}

	return nil
}

func (cps *ConfigProcessorService) SaveAll(configs []*config.ConfigProcessor) error {
	err := cps.ConfigProcessorRepository.SaveAll(configs)

	if err != nil {
		return err
	}

	return nil
}
