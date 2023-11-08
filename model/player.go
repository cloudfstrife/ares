package model

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

//Player 游戏玩家
type Player struct {
	Lock  sync.Mutex
	Name  string
	Color string
	Golds int
	Woods int
	Units map[string]*Unit
}

//String 输出
func (player *Player) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("昵称：%-10s\t黄金：%-10d\t木材：%-10d\t单位：%-10d\n", player.Name, player.Golds, player.Woods, len(player.Units)))
	builder.WriteString("-------------------------------------------------------------------------------------------------------------------------------\n")
	builder.WriteString("名称\t标识\t生命值\t魔法值\t命中率\t闪避率\t物理攻击力\t魔法攻击力\t每分钟生产的黄金数量\t每分钟生产的木材数量\n")
	for _, u := range player.Units {
		builder.WriteString(fmt.Sprintf("%s\n", u.String()))
	}
	builder.WriteString("-------------------------------------------------------------------------------------------------------------------------------\n")
	return builder.String()
}

//LoadUnits 初始化作战单位
func (player *Player) LoadUnits(ucs []*UnitConfig) {
	for i, uc := range ucs {
		player.Units[uc.Name+strconv.Itoa(i+1)] = &Unit{
			UnitCommonAttribute: UnitCommonAttribute{
				lock:   sync.Mutex{},
				Player: player,
				Type:   uc.Type,
				Sign:   uc.Name + strconv.Itoa(i+1),
				Name:   uc.Name,
				HP:     uc.HP,
				MP:     uc.MP,
				HIT:    uc.HIT,
				AVD:    uc.AVD,
			},
			ATK:   uc.ATK,
			MGK:   uc.MGK,
			Golds: uc.Golds,
			Woods: uc.Woods,
		}
	}
}
