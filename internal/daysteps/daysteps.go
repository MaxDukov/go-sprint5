package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"tracker/internal/personaldata"
	"tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	vals := strings.Split(datastring, ",")
	if len(vals) != 2 {
		return errors.New("неверный формат данных")
	}
	steps, err := strconv.Atoi(vals[0])
	if err != nil {
		return errors.New("неверный формат данных")
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше 0")
	}
	duration, err := time.ParseDuration(vals[1])
	if err != nil {
		return errors.New("неверный формат данных")
	}
	if duration <= time.Duration(0) {
		return errors.New("продолжительность должна быть больше 0")
	}
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	calories, err := spentenergy.RunningSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d\nДистанция: %.2f км.\nСожгли калорий: %.2f\n", ds.Steps, distance, calories), nil

}
