package gotwit

import (
	"strings"
)

type coordinates struct {
	Type        string `json:"type"`
	Coordinates [2]int `json:"coordinates"`
}

type entitiesT struct {
	Hashtags     []any `json:"hashtags"`
	Symbols      []any `json:"symbols"`
	UserMentions []any `json:"user_mentions"`
}

type gterror struct {
	Detail     string `json:"detail"`
	Message    string `json:"value"`
	Title      string `json:"title"`
	Error_type string `json:"type"`
}

type geo struct {
	Coordinates coordinates `json:"coordinates"`
	PlaceId     string      `json:"place_id"`
}

type media struct {
	MediaIds      []string `json:"media_ids"`
	TaggedUserIds []string `json:"tagged_user_ids"`
}

type poll struct {
	DurationMinutes uint     `json:"duration_minutes"`
	Options         []string `json:"options"`
}

type startEnd struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type getOption func() (key string, val string)

func With(key string, vals ...string) getOption {
	return func() (string, string) {
		return key, strings.Join(vals, ",")
	}
}
