package storage

import (
	"fmt"
	"spells/cmd/models"
)

type SpellStorage interface {
	GetAllSpells() []models.Spell
	GetSpell(id int) (models.Spell, bool)
	CreateSpell(newSpell models.Spell) (int, error)
	UpdateSpell(id int, newSpell models.Spell) (int, error)
	DeleteSpell(id int) (int, error)
}

type LocalSpellStorage struct {
	storage []models.Spell
}

func NewLocalStorage(intialspells []models.Spell) LocalSpellStorage {
	return LocalSpellStorage{intialspells}
}

func (l *LocalSpellStorage) GetAllSpells() []models.Spell {
	return l.storage
}

func (l *LocalSpellStorage) GetSpell(id int) (models.Spell, bool) {
	var emptySpell models.Spell

	for _, spell := range l.storage {
		if spell.ID == id {
			return spell, true
		}
	}
	return emptySpell, false
}

func (l *LocalSpellStorage) CreateSpell(newSpell models.Spell) (int, error) {
	l.storage = append(l.storage, newSpell)
	return newSpell.ID, nil
}

func (l *LocalSpellStorage) UpdateSpell(id int, newSpell models.Spell) (int, error) {
	for i, spell := range l.storage {
		if spell.ID == id {
			l.storage[i] = updateSpell(spell, newSpell)
			return l.storage[i].ID, nil
		}
	}
	return 0, fmt.Errorf("cannot find the spell %d", id)
}

func updateSpell(original, update models.Spell) (updated models.Spell) {
	updated.ID = original.ID
	updated.Name = ternaryString(update.Name != "", update.Name, original.Name)
	updated.Target = ternaryString(update.Target != "", update.Target, original.Target)
	updated.School = ternaryString(update.School != "", update.School, original.School)
	updated.CastingTime = ternaryString(update.CastingTime != "", update.CastingTime, original.CastingTime)
	updated.Range = ternaryString(update.Range != "", update.Range, original.Range)
	updated.Duration = ternaryString(update.Duration != "", update.Duration, original.Duration)
	updated.Components = ternaryString(update.Components != "", update.Components, original.Components)
	updated.Effect = ternaryString(update.Effect != "", update.Effect, original.Effect)

	if len(update.SpellList) != 0 {
		updated.SpellList = update.SpellList
	} else {
		updated.SpellList = original.SpellList
	}

	return
}

func ternaryString(expresion bool, expected, fallback string) string {
	if expresion {
		return expected
	}
	return fallback
}

func (l *LocalSpellStorage) DeleteSpell(id int) (int, error) {
	for i, spell := range l.storage {
		if spell.ID == id {
			l.storage = append(l.storage[:i], l.storage[i+1:]...)
		}
		return i, nil
	}
	return 0, fmt.Errorf("cannot find the Spell %d", id)
}
