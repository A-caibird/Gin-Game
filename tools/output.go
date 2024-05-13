package tools

import (
	"github.com/fatih/color"
	"io"
	"os"
	"time"
)

// Info output info log with color to log file writer and os.stdout
func Info(format string, a ...interface{}) {
	file, err := os.OpenFile(Conf.RootPath.Path+"/log/gin.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		color.Red("file file" + Conf.RootPath.Path)
		return
	}
	multiWriter := io.MultiWriter(file, os.Stdout)
	red := color.New(color.BgRed)
	red.Fprintf(multiWriter, "%s\t", time.Now().String())
	red.Fprintf(multiWriter, format, a)
}

// Red output error log with color to log file writer and os.stdout
func Red(format string, a ...interface{}) {
	file, err := os.OpenFile(Conf.RootPath.Path+"/log/gin.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		color.Red("file file" + Conf.RootPath.Path)
		return
	}
	multiWriter := io.MultiWriter(file, os.Stdout)
	red := color.New(color.BgRed)
	red.Fprintf(multiWriter, "%s\t", time.Now().String())
	red.Fprintf(multiWriter, format+"\n", a)
}
