package abstractfactory

import (
	"fmt"
	"testing"
)

// 抽象工厂模式
func TestAbstractFactory(t *testing.T) {
	fmt.Printf("breakfast: %v\n", HaveMeal(NewBreakfastCook()))
	fmt.Printf("lunch: %v\n", HaveMeal(NewLunchCook()))
	fmt.Printf("dinner: %v\n", HaveMeal(NewDinnerCook()))
}

// HaveMeal 吃饭
func HaveMeal(cook Cook) string {
	return fmt.Sprintf("%s %s", cook.MakeFood().Eaten(), cook.MakeDrink().Drunk())
}
