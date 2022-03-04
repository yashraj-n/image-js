package utils

import "github.com/fatih/color"

func Error(err string) {
	color.Red(err)
}

func Info(info string) {
	color.Green(info)
}

func Warn(warn string) {
	color.Yellow(warn)
}

func Debug(debug string) {
	color.Blue(debug)
}

func Success(success string) {
	color.Green(success)
}
