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
	languages       []LanguageLevel
}

type LanguageLevel struct {
	lang        string
	description string
}

func NewUserLevel(level int, totalPoints, levelPercentage int32, languages []LanguageLevel) *UserLevel {
	userLevel := &UserLevel{
		level:           level,
		totalPoints:     totalPoints,
		levelPercentage: levelPercentage,
		languages:       languages,
	}
	return userLevel
}

func NewLanguageLevel(lang, description string) LanguageLevel {
	return LanguageLevel{lang: lang, description: description}
}

func (l *LanguageLevel) Lang() string {
	return l.lang
}

func (l *LanguageLevel) Description() string {
	return l.description
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

func (l *UserLevel) Languages() []LanguageLevel {
	return l.languages
}
