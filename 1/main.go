package main

import (
	"fmt"
)

type Human struct {
	Age    int
	Name   string
	Height int
	Weight float32
}

type Action struct {
	Human
	ActionType string
}

/////// Создадим методы для обоих структур

func (h *Human) getAge() int {
	return h.Age
}

func (h *Human) getName() string {
	return h.Name
}

func (h *Human) getHeight() int {
	return h.Height
}

func (h *Human) getWeight() float32 {
	return h.Weight
}

func (h *Human) updateAge(newAge int) {
	h.Age = newAge
}

func (h *Human) updateName(newName string) {
	h.Name = newName
}

func (h *Human) updateHeight(newHeight int) {
	h.Height = newHeight
}

func (h *Human) updateWeight(newWeight float32) {
	h.Weight = newWeight
}

func (a *Action) getActionType() string {
	return a.ActionType
}

func (a *Action) updateActionType(newActionType string) {
	a.ActionType = newActionType
}

////////

func main() {

	a := Action{} // Создадим структуру Action

	// Через ранее созданные методы укажем значения полей
	a.updateName("Nikita")
	a.updateAge(20)
	a.updateHeight(185)
	a.updateWeight(78)
	a.updateActionType("Writing code")

	// Испольузем созданные методы для получения значений полей
	myName := a.getName()
	myAge := a.getAge()
	myHeight := a.getHeight()
	myWeight := a.getWeight()
	myActionType := a.getActionType()

	// Выведем значения полей
	fmt.Printf("Имя: %s\nВозраст(годы): %d\nРост(см): %d\nВес(кг): %.2f\nТекущее действие: %s\n", myName, myAge, myHeight, myWeight, myActionType)

}
