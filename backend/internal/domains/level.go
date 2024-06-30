package domains

import (
	"context"
)

type ILevelService interface {
	GetUserLevel(ctx context.Context, userID string) (*UserLevel, error)
	UpdateUserPoint(ctx context.Context, userID string, difficulty int32) error
}

type UserLevel struct {
	level           int
	totalPoints     int32
	levelPercentage int32
	languages       []LanguageL
}

type LanguageL struct {
	Lang        string
	Description string
}

func NewUserLevel(level int, totalPoints, levelPercentage int32, languages []LanguageL) *UserLevel {
	userLevel := &UserLevel{
		level:           level,
		totalPoints:     totalPoints,
		levelPercentage: levelPercentage,
		languages:       languages,
	}
	return userLevel
}

func (l *UserLevel) Level() int {
	return l.level
}

func (l *UserLevel) TotalPoints() int32 {
	return l.totalPoints
}

func (l *UserLevel) LevelPercentage() int32 {
	return l.levelPercentage
}

func (l *UserLevel) Languages() []LanguageL {
	return l.languages
}
