package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"spells/cmd/models"
	"spells/cmd/storage"
	"strconv"

	"github.com/labstack/echo/v4"
)

var initialValue = []models.Spell{
	{
		ID: 1, Name: "Bola de Fuego",
		School:      "Evoation",
		CastingTime: "1 action",
		Target:      "A point you choose within range",
		Range:       "120 feets",
		Duration:    "instantaneus",
		Components:  "V, S, M",
		Effect:      "A bright streak flashes from your pointing finger to a point you choose within range and then blossoms with a low roar into an explosion of flame. Each creature in a 20-foot-radius sphere centered on that point must make a Dexterity saving throw. A target takes 8d6 fire damage on a failed save, or half as much damage on a successful one. The fire spreads around corners. It ignites flammable objects in the area that aren`t being worn or carried.At Higher Levels: When you cast this spell using a spell slot of 4th level or higher, the damage increases by 1d6 for each slot level above 3rd.",
		SpellList:   []string{"Wizzard", "Sorcerer"},
	},
}

var Storage = storage.NewLocalStorage(initialValue)

func GetAllSpells(c echo.Context) error {
	return c.JSON(http.StatusOK, Storage.GetAllSpells())
}

func GetSpell(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("spellId"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid Id")
	}

	spell, ok := Storage.GetSpell(id)

	if ok {
		return c.JSON(http.StatusOK, spell)
	}

	return c.String(http.StatusNotFound, fmt.Sprintf("The spell with ID %d was not found", id))
}

func CreateSpell(c echo.Context) error {
	newSpell := models.Spell{ID: rand.Intn(10000)}
	err := c.Bind(&newSpell)

	Storage.CreateSpell(newSpell)

	if err != nil {
		log.Printf("The was an error in the body %v", err)
		return c.String(http.StatusBadRequest, "Bad spell")
	}
	return c.String(http.StatusCreated, fmt.Sprintf("the Spell %d was created", newSpell.ID))
}

func UpdateSpell(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("spellId"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid id")
	}
	updatededSpell := models.Spell{ID: rand.Intn(10000)}
	err = c.Bind(&updatededSpell)

	if err != nil {
		log.Printf("The was an error in the body %v", err)
		return c.String(http.StatusBadRequest, "Bad spell")
	}

	_, err = Storage.UpdateSpell(id, updatededSpell)

	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot updated spell")
	}
	return c.String(http.StatusOK, fmt.Sprintf("THe Spell %d was modified succefully", id))
}

func DeleteSpell(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("spellId"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid id")
	}
	_, err = Storage.DeleteSpell(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot delete spell")
	}

	return c.String(http.StatusOK, fmt.Sprintf("The spell %d was deleted", id))
}
