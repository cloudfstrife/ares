package model

import (
	"fmt"
	"strings"
)

//CampConfig 阵营配置
type CampConfig struct {
	Name        string
	UnitsConfig []*UnitConfig
}

func (cc *CampConfig) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("                           %s                           \n", cc.Name))
	builder.WriteString("--------------------    单位    --------------------\n")
	builder.WriteString("名称\t生命值\t魔法值\t命令率\t闪避率\t物理攻击力\t魔法攻击力\t生产黄金数量\t生产木材数量\t\n")
	for _, un := range cc.UnitsConfig {
		builder.WriteString(fmt.Sprintf("%s\n", un.String()))
	}
	return builder.String()
}

//CommonUnitConfig 公共配置字段
type CommonUnitConfig struct {
	//Name 名称
	Name string
	//Type 单位类型
	Type UnitType
	//HP 生命值
	HP int
	//MP 魔法值
	MP int
	//HIT 命中率
	HIT int
	// 闪避率
	AVD int
}

//UnitConfig 生产单位配置
type UnitConfig struct {
	CommonUnitConfig
	//ATK 物理攻击力
	ATK int
	//MGK 魔法攻击
	MGK int
	//每分钟生产的黄金数量
	Golds int
	//每分钟生产的木材数量
	Woods int
}

func (uc *UnitConfig) String() string {
	return fmt.Sprintf("%s\t%5d\t%5d\t%5d\t%5d\t%8d\t%8d\t%10d\t%10d", uc.CommonUnitConfig.Name, uc.CommonUnitConfig.HP, uc.CommonUnitConfig.MP, uc.CommonUnitConfig.HIT, uc.CommonUnitConfig.AVD, uc.ATK, uc.MGK, uc.Golds, uc.Woods)
}

//InitHumansConfig 初始化人族配置
func InitHumansConfig() *CampConfig {
	result := &CampConfig{Name: "人类", UnitsConfig: make([]*UnitConfig, 0, 4)}

	result.UnitsConfig = append(result.UnitsConfig, &UnitConfig{
		CommonUnitConfig: CommonUnitConfig{
			Name: "农民",
			Type: BUILDER,
			HP:   100,
			MP:   100,
			HIT:  5,
			AVD:  20,
		},
		Golds: 10,
		Woods: 10,
	})
	result.UnitsConfig = append(result.UnitsConfig, &UnitConfig{
		CommonUnitConfig: CommonUnitConfig{
			Name: "骑士",
			Type: ATTACTER,
			HP:   100,
			MP:   100,
			HIT:  75,
			AVD:  25,
		},
		ATK: 5,
		MGK: 2,
	})
	result.UnitsConfig = append(result.UnitsConfig, &UnitConfig{
		CommonUnitConfig: CommonUnitConfig{
			Name: "女巫",
			Type: ATTACTER,
			HP:   100,
			MP:   100,
			HIT:  25,
			AVD:  75,
		},
		ATK: 2,
		MGK: 5,
	})
	result.UnitsConfig = append(result.UnitsConfig, &UnitConfig{
		CommonUnitConfig: CommonUnitConfig{
			Name: "血法师",
			Type: ATTACTER,
			HP:   100,
			MP:   100,
			HIT:  80,
			AVD:  80,
		},
		ATK: 10,
		MGK: 10,
	})
	return result
}

//InitNightElvesConfig 初始化暗夜精灵配置
func InitNightElvesConfig() *CampConfig {
	result := &CampConfig{Name: "暗夜精灵", UnitsConfig: make([]*UnitConfig, 0, 1)}
	result.UnitsConfig = append(result.UnitsConfig, &UnitConfig{
		CommonUnitConfig: CommonUnitConfig{
			Name: "小精灵",
			Type: BUILDER,
			HP:   100,
			MP:   100,
			HIT:  5,
			AVD:  20,
		},
		Golds: 10,
		Woods: 10,
	})
	result.UnitsConfig = append(result.UnitsConfig, &UnitConfig{
		CommonUnitConfig: CommonUnitConfig{
			Name: "弓箭手",
			Type: ATTACTER,
			HP:   100,
			MP:   100,
			HIT:  75,
			AVD:  25,
		},
		ATK: 5,
		MGK: 2,
	})
	result.UnitsConfig = append(result.UnitsConfig, &UnitConfig{
		CommonUnitConfig: CommonUnitConfig{
			Name: "树妖",
			Type: ATTACTER,
			HP:   100,
			MP:   100,
			HIT:  25,
			AVD:  75,
		},
		ATK: 2,
		MGK: 5,
	})
	result.UnitsConfig = append(result.UnitsConfig, &UnitConfig{
		CommonUnitConfig: CommonUnitConfig{
			Name: "守望者",
			Type: ATTACTER,
			HP:   100,
			MP:   100,
			HIT:  80,
			AVD:  80,
		},
		ATK: 10,
		MGK: 10,
	})
	return result
}
