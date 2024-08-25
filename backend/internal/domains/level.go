package domains

import (
	"context"
)

type ILevelService interface {
	GetUserLevel(ctx context.Context, userID string) (*UserLevel, error)
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

// LanguageLevel Getter
func (l *LanguageLevel) Lang() string {
	return l.lang
}

func (l *LanguageLevel) Description() string {
	return l.description
}

// UserLevel Getter
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

// Setters
func (l *UserLevel) SetLevel(level int) {
	l.level = level
}
func (l *UserLevel) SetTotalPoints(totalPoints int32) {
	l.totalPoints = totalPoints
}
func (l *UserLevel) SetLevelPercentage(levelPercentage int32) {
	l.levelPercentage = levelPercentage
}
func (l *UserLevel) SetLanguages(languages []LanguageLevel) {
	l.languages = languages
}
