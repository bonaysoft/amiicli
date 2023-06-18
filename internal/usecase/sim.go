package usecase

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/AlecAivazis/survey/v2"
	"github.com/abusomani/go-palette/palette"
	"github.com/bonaysoft/amiicli/internal/entity"
	"github.com/bonaysoft/amiicli/internal/repository"
	"github.com/samber/lo"
)

type Sim struct {
	device repository.Device
}

func NewSim(device repository.Device) *Sim {
	return &Sim{device: device}
}

func (s *Sim) Action(ctx context.Context, amiibos []entity.Amiibo) error {
	var answer entity.SimulateRequest
	if err := survey.Ask(s.questions(), &answer); err != nil {
		return err
	}

	amiiboSelect := func() (*entity.Amiibo, error) { return &amiibos[rand.Intn(len(amiibos))], nil }
	if entity.Mode(answer.Mode) == entity.ModeSpecify {
		var amiiboName string
		if err := survey.AskOne(s.specifyQuestion(amiibos), &amiiboName); err != nil {
			return fmt.Errorf("ask amiiboName: %v", err)
		}

		amiibo, ok := lo.Find(amiibos, func(item entity.Amiibo) bool { return item.Name == amiiboName })
		if !ok {
			return fmt.Errorf("not found the specified amiibo: %s", amiiboName)
		}

		amiiboSelect = func() (*entity.Amiibo, error) {
			return &amiibo, nil
		}
	}

	p := palette.New()
	p.SetOptions(palette.WithForeground(palette.Blue))
	for i := 0; i < answer.Times; i++ {
		amiibo, err := amiiboSelect()
		if err != nil {
			return fmt.Errorf("select amiibo failed: %s", err)
		}

		if err := amiibo.Build(); err != nil {
			return fmt.Errorf("build amiibo failed: %s", err)
		}

		p.Printf("starting simulate amiibo[%d] %s...\n", i+1, amiibo.Name)
		if err := s.device.Simulate(ctx, amiibo); err != nil {
			return fmt.Errorf("simulate amiibo[%s] failed: %s", amiibo.Name, err)
		}
	}

	p.Println("simulate done.")
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

func (s *Sim) specifyQuestion(amiibos []entity.Amiibo) survey.Prompt {
	return &survey.Select{
		Message: "请选择Amiibo:",
		Options: lo.Map(amiibos, func(item entity.Amiibo, index int) string { return item.Name }),
	}
}
