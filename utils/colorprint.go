package utils

import (
	"io"
	"os"
)

const (

	//BLACK 黑色
	BLACK = "\u001b[30m"
	//RED 红色
	RED = "\u001b[31m"
	//GREEN 绿色
	GREEN = "\u001b[32m"
	//YELLOW 黄色
	YELLOW = "\u001b[33m"
	//BLUE 蓝色
	BLUE = "\u001b[34m"
	//CARMINE 洋红色
	CARMINE = "\u001b[35m"
	//CYAN 青色
	CYAN = "\u001b[36m"
	//WHITE 白色
	WHITE = "\u001b[37m"
	//GLOSSBLACK 亮黑色
	GLOSSBLACK = "\u001b[30;1m"
	//GLOSSRED 亮红色
	GLOSSRED = "\u001b[31;1m"
	//GLOSSGREEN 亮绿色
	GLOSSGREEN = "\u001b[32;1m"
	//GLOSSYELLOW 亮黄色
	GLOSSYELLOW = "\u001b[33;1m"
	//GLOSSBLUE 亮蓝色
	GLOSSBLUE = "\u001b[34;1m"
	//GLOSSCARMINE 亮洋红色
	GLOSSCARMINE = "\u001b[35;1m"
	//GLOSSCYAN 亮青色
	GLOSSCYAN = "\u001b[36;1m"
	//GLOSSWHITE 亮白色
	GLOSSWHITE = "\u001b[37;1m"
	//RESET 重置
	RESET = "\u001b[0m"
)

//FprintWithColor 以指定颜色，将指定内容输出到流
func FprintWithColor(w io.Writer, color, v string) {
	w.Write([]byte(color))
	w.Write([]byte(v))
	w.Write([]byte(RESET))
}

//PrintWithColor 以指定颜色，将指定内容输出到标准输出设备
func PrintWithColor(color, v string) {
	os.Stdout.Write([]byte(color))
	os.Stdout.Write([]byte(v))
	os.Stdout.Write([]byte(RESET))
}

//PrintlnWithColor 以指定颜色，将指定内容输出到标准输出设备，并换行
func PrintlnWithColor(color, v string) {
	os.Stdout.Write([]byte(color))
	os.Stdout.Write([]byte(v))
	os.Stdout.Write([]byte(RESET))
	os.Stdout.Write([]byte("\n"))
}
