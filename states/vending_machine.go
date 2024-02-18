package states

import (
	"golearnpatternstate/consumer"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type VendingMachine struct {
	HasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	currentState  State
	CurS          string
	itemCount     int
	itemPrice     int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}

	// default state
	v.SetState(hasItemState, "hasItemState")

	v.HasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *VendingMachine) AddItem(count int, bot *consumer.Consumer, update tgbotapi.Update) {
	v.currentState.addItem(count, bot, update)
}

func (v *VendingMachine) RequestItem(bot *consumer.Consumer, update tgbotapi.Update) {
	v.currentState.requestItem(bot, update)
}

func (v *VendingMachine) InsertMoney(money int, bot *consumer.Consumer, update tgbotapi.Update) {
	v.currentState.insertMoney(money, bot, update)
}

func (v *VendingMachine) DispenseItem(bot *consumer.Consumer, update tgbotapi.Update) {
	v.currentState.dispenseItem(bot, update)
}

func (v *VendingMachine) Exit(bot *consumer.Consumer, update tgbotapi.Update) {
	v.currentState.exit(bot, update)
}

func (v *VendingMachine) GetItemCount() int {
	return v.itemCount
}

func (v *VendingMachine) SetState(s State, n string) {
	v.currentState = s
	v.CurS = n
}

func (v *VendingMachine) incrementItemCount(count int) {
	log.Printf("machine: Adding %d items\n", count)
	v.itemCount += count
}
