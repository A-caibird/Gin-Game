package tools

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"
)

type Config struct {
	MySQLServer struct {
		Host      string `toml:"host"`
		Port      string `toml:"port"`
		User      string `toml:"user"`
		Password  string `toml:"password"`
		Database  string `toml:"database"`
		Charset   string `toml:"charset"`
		ParseTime string `toml:"parseTime"`
		Loc       string `toml:"loc"`
		Dev       struct {
			Host      string `toml:"host"`
			Port      string `toml:"port"`
			User      string `toml:"user"`
			Password  string `toml:"password"`
			Database  string `toml:"database"`
			Charset   string `toml:"charset"`
			ParseTime string `toml:"parseTime"`
			Loc       string `toml:"loc"`
		} `toml:"dev"`
	} `toml:"mysql-server"`
	RedisServer struct {
		Host           string `toml:"host"`
		Port           string `toml:"port"`
		Password       string `toml:"password"`
		Database       int    `toml:"database"`
		MaxActiveConns int    `toml:"maxActiveConns"`
		MaxIdleConns   int    `toml:"maxIdleConns"`
	} `toml:"redis"`
	StaticFilePath struct {
		Path string `toml:"path"`
	} `toml:"static-files-path"`
	LogPath struct {
		Path string `toml:"path"`
	} `toml:"log-path"`
	Email struct {
		Host         string `toml:"host"`
		Port         int    `toml:"port"`
		User         string `toml:"user"`
		AuthPassword string `toml:"AuthPassword"`
	} `toml:"email"`
	Aliyun struct {
		AccessKeyId     string `toml:"AccessKeyId"`
		AccessKeySecret string `toml:"AccessKeySecret"`
		SMS             struct {
			SignName     string `toml:"SignName"`
			TemplateCode string `toml:"TemplateCode"`
			Domain       string `toml:"Domain"`
		} `toml:"sms"`
	} `toml:"aliyun"`
	RabbitMQ struct {
		User     string `toml:"user"`
		Password string `toml:"password"`
		Host     string `toml:"host"`
		Port     string `toml:"port"`
	} `toml:"RabbitMQ"`
	RootPath struct {
		Path string `toml:"path"`
	} `toml:"RootPath"`
}

var (
	once sync.Once
	Conf *Config
)

func init() {
	once.Do(ParseToml)
}
func ParseToml() {
	// 获取当前运行堆栈的帧信息
	//Retrieve stack frame information for the current execution stack
	//var configPath string
	var configPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		index := strings.Index(filename, "Game/")
		if index != -1 {
			configPath = filename[:index+5] + "config.toml"
		}
	}
	//Retrieve file status information
	if _, err := os.Stat(configPath); err != nil {
		panic(err)
	}
	// decode toml file
	var config Config
	meta, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%#v", err)
		os.Exit(1)
	} else {
		Conf = &config
	}
	// display toml file's structure
	indent := strings.Repeat(" ", 14)
	fmt.Print("Decoded")
	typ, val := reflect.TypeOf(config), reflect.ValueOf(config)
	for i := 0; i < typ.NumField(); i++ {
		indent := indent
		if i == 0 {
			indent = strings.Repeat(" ", 7)
		}
		fmt.Printf("%s%-11s → %v\n", indent, typ.Field(i).Name, val.Field(i))
	}
	fmt.Print("\nKeys")
	keys := meta.Keys()
	sort.Slice(keys, func(i, j int) bool { return keys[i].String() < keys[j].String() })
	for i, k := range keys {
		indent := indent
		if i == 0 {
			indent = strings.Repeat(" ", 10)
		}
		log.Printf("%s%-10s %s\n", indent, meta.Type(k...), k)
	}
	fmt.Print("\nUndecoded\n")
	keys = meta.Undecoded()
	sort.Slice(keys, func(i, j int) bool { return keys[i].String() < keys[j].String() })
	for i, k := range keys {
		indent := indent
		if i == 0 {
			indent = strings.Repeat(" ", 5)
		}
		log.Printf("%s%-10s %s\n", indent, meta.Type(k...), k)
	}
}
