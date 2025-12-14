package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"tracker/internal/personaldata"
	"tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	vals := strings.Split(datastring, ",")
	if len(vals) != 3 {
		return errors.New("неверный формат данных")
	}
	steps, err := strconv.Atoi(vals[0])
	if err != nil {
		return errors.New("неверный формат данных")
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше 0")
	}
	t.Steps = steps
	t.TrainingType = vals[1]
	duration, err := time.ParseDuration(vals[2])
	if err != nil {
		return errors.New("неверный формат данных")
	}
	if duration <= time.Duration(0) {
		return errors.New("продолжительность должна быть больше 0")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.TrainingType != "Ходьба" && t.TrainingType != "Бег" {
		return "", errors.New("неизвестный тип тренировки")
	}
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	var calories float64
	var err error
	if t.TrainingType == "Ходьба" {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}
