package model

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"

	"github.com/cloudfstrife/ares/utils"
)

//UnitType 单位类型
type UnitType int

const (
	_ UnitType = iota
	//BUILDER 构造者
	BUILDER
	//ATTACTER 攻击者
	ATTACTER
)

//UnitCommonAttribute 单位通用属性
type UnitCommonAttribute struct {
	//lock 锁
	lock sync.Mutex
	//player 所属阵营
	Player *Player
	//类型
	Type UnitType
	//标识
	Sign string
	//Name 名称
	Name string
	//HP 生命值
	HP int
	//MP 魔法值
	MP int
	//HIT 命中率
	HIT int
	// 闪避率
	AVD int
}

//Unit 作战单位
type Unit struct {
	UnitCommonAttribute
	//ATK 物理攻击力
	ATK int
	//MGK 魔法攻击
	MGK int
	//每分钟生产的黄金数量
	Golds int
	//每分钟生产的木材数量
	Woods int
}

//BeAttacked 被攻击
func (uca *UnitCommonAttribute) BeAttacked(physicalDamage, magicDamage int) (elude, death bool) {
	uca.lock.Lock()
	defer uca.lock.Unlock()
	elude = uca.Elude()
	if !elude {
		uca.HP = uca.HP - physicalDamage - magicDamage
		if uca.HP <= 0 {
			delete(uca.Player.Units, uca.Sign)
			death = true
		}
	}
	return
}

//Elude 判断是否躲避成功
func (uca *UnitCommonAttribute) Elude() bool {
	return utils.Judge(int64(uca.AVD))
}

//HitOrNot 判断是否命中
func (u *Unit) HitOrNot() bool {
	return utils.Judge(int64(u.HIT))
}

//Attack 攻击
func (u *Unit) Attack(player *Player) (hit, elude, death bool, damage int, target string, err error) {
	//选择目标
	buf := bytes.NewBufferString("")
	player.Lock.Lock()
	defer player.Lock.Unlock()
	r, err := rand.Int(rand.Reader, big.NewInt(int64(len(player.Units))))
	if err != nil {
		buf.WriteString(fmt.Sprintf("[ %s ] 发起进攻，选择目标失败", u.Name))
		return false, false, false, 0, buf.String(), err
	}
	keys := make([]string, 0, len(player.Units))
	for k := range player.Units {
		keys = append(keys, k)
	}
	attackTarget := player.Units[keys[int(r.Int64())]]

	hit = u.HitOrNot()
	if !hit {
		buf.WriteString(fmt.Sprintf("[ %s ] 对 [ %s ] 发起进攻，未命中！", u.Name, attackTarget.Name))
		return hit, false, false, 0, buf.String(), nil
	}
	elude, death = attackTarget.BeAttacked(u.HIT, u.MGK)
	if elude {
		buf.WriteString(fmt.Sprintf("[ %s ] 对 [ %s ] 发起进攻，[ %s ] 躲避成功！", u.Name, attackTarget.Name, attackTarget.Name))
		return hit, false, false, 0, buf.String(), nil
	} else {
		buf.WriteString(fmt.Sprintf("[ %s ] 对 [ %s ] 发起进攻，造成 [ %d ] 点伤害！", u.Name, attackTarget.Name, u.ATK+u.MGK))
		if death {
			buf.WriteString(fmt.Sprintf("[ %s ] 死亡！", attackTarget.Name))
		}
		return hit, false, false, u.ATK + u.MGK, buf.String(), nil
	}
}

//Produce 生产
func (u *Unit) Produce() {
	u.Player.Lock.Lock()
	defer u.Player.Lock.Unlock()
	u.Player.Golds = u.Player.Golds + u.Golds
	u.Player.Woods = u.Player.Woods + u.Woods
}

//String String
func (u *Unit) String() string {
	return fmt.Sprintf("%s\t%s\t%d\t%d\t%d\t%d\t%10d\t%10d\t%20d\t%20d", u.Name, u.Sign, u.HP, u.MP, u.HIT, u.AVD, u.ATK, u.MGK, u.Golds, u.Woods)
}
