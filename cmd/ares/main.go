package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/cloudfstrife/ares/utils"

	"github.com/cloudfstrife/ares/model"
)

func main() {
	HumanPlayer := &model.Player{
		Lock:  sync.Mutex{},
		Name:  "人类",
		Color: utils.YELLOW,
		Golds: 100,
		Woods: 100,
		Units: make(map[string]*model.Unit),
	}
	HumanPlayer.LoadUnits(model.InitHumansConfig().UnitsConfig)

	NPlayer := &model.Player{
		Lock:  sync.Mutex{},
		Name:  "暗夜精灵",
		Color: utils.RED,
		Golds: 100,
		Woods: 100,
		Units: make(map[string]*model.Unit),
	}
	NPlayer.LoadUnits(model.InitNightElvesConfig().UnitsConfig)
	Fright(HumanPlayer, NPlayer)
}

//Fright 战斗
func Fright(a, b *model.Player) {
	fmt.Println(a)
	fmt.Println(b)
	i := 0
done:
	for {

		if i%2 == 0 {
			for _, u := range a.Units {
				if u.Type == model.BUILDER {
					u.Produce()
				}
				if u.Type == model.ATTACTER {
					_, _, _, _, msg, _ := u.Attack(b)
					utils.PrintlnWithColor(u.Player.Color, msg)
				}
				if IsDone(a, b) {
					break done
				}
				time.Sleep(time.Second)
			}
		} else {
			for _, u := range b.Units {
				if u.Type == model.BUILDER {
					u.Produce()
				}
				if u.Type == model.ATTACTER {
					_, _, _, _, msg, _ := u.Attack(a)
					utils.PrintlnWithColor(u.Player.Color, msg)
				}
				if IsDone(a, b) {
					break done
				}
				time.Sleep(time.Second)
			}
		}
		i++
	}
	fmt.Println(a)
	fmt.Println(b)
}

//IsDone 检查是否结束游戏
func IsDone(a, b *model.Player) bool {
	return len(a.Units) <= 0 || len(b.Units) == 0
}
