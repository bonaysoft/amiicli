package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bonaysoft/amiicli/internal/entity"
	"github.com/bonaysoft/amiicli/internal/repository"
	"github.com/samber/lo"
)

type Sim struct {
	amiibo repository.Amiibo
	device repository.Device
}

func NewSim(amiibo repository.Amiibo, device repository.Device) *Sim {
	return &Sim{amiibo: amiibo, device: device}
}

func (s *Sim) Action(ctx context.Context) error {
	var answer entity.SimulateRequest
	if err := survey.Ask(s.questions(), &answer); err != nil {
		return err
	}

	var amiiboName string
	var selectOpts []repository.AmiiboSelectOption
	if entity.Mode(answer.Mode) == entity.ModeSpecify {
		if err := survey.AskOne(s.specifyQuestion(ctx), &amiiboName); err != nil {
			return fmt.Errorf("ask amiiboName: %v", err)
		}

		selectOpts = append(selectOpts, repository.AmiiboSelectWithName(amiiboName))
	}

	for i := 0; i < answer.Times; i++ {
		amiibo, err := s.amiibo.Select(ctx, entity.Mode(answer.Mode), selectOpts...)
		if err != nil {
			return fmt.Errorf("select amiibo failed: %s", err)
		}

		fmt.Printf("starting simulate amiibo %s...\n", amiibo.Name)
		if err := s.device.Simulate(ctx, amiibo); err != nil {
			return fmt.Errorf("simulate amiibo[%s] failed: %s", amiibo.Name, err)
		}
	}

	fmt.Sprintln("simulate done.")
	return nil
}

func (s *Sim) questions() []*survey.Question {
	return []*survey.Question{
		{
			Name: "mode",
			Prompt: &survey.Select{
				Message: "请选择模式:",
				Options: []string{"random", "specify"},
				Default: "random",
			},
		},
		{
			Name:   "times",
			Prompt: &survey.Input{Message: "请设置模拟次数:"},
		},
	}
}

func (s *Sim) specifyQuestion(ctx context.Context) survey.Prompt {
	amiibos, err := s.amiibo.List(ctx)
	if err != nil {
		log.Fatalf("error listing amiibos: %v", err)
	}

	return &survey.Select{
		Message: "请选择Amiibo:",
		Options: lo.Map(amiibos, func(item entity.Amiibo, index int) string { return item.Name }),
	}
}
