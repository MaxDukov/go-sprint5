package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// RunningSpentCalories рассчитывает количество сожженных калорий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 || steps <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("неверные данные")
	}

	speed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * speed * durationInMinutes) / minInH
	return calories, nil
}

// WalkingSpentCalories рассчитывает количество сожженных калорий при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 || steps <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("неверные данные")
	}
	durationInMinutes := duration.Minutes()
	speed := MeanSpeed(steps, height, duration)
	calories := walkingCaloriesCoefficient * (weight * speed * durationInMinutes) / minInH
	return calories, nil
}

// meanSpeed рассчитывает среднюю скорость, с которой была пройдена дистанция
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	if steps <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// Distance рассчитывает дистанцию, пройденную за тренировку
func Distance(steps int, height float64) float64 {
	stepsLength := float64(steps) * stepLengthCoefficient * height
	return stepsLength / mInKm
}
