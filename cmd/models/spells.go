package models

type Spell struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	School      string   `json:"school"`
	Target      string   `json:"target"`
	CastingTime string   `json:"castingTime"`
	Range       string   `json:"range"`
	Duration    string   `json:"duration"`
	Components  string   `json:"components"`
	Effect      string   `json:"effect"`
	SpellList   []string `json:"spellList"`
}
