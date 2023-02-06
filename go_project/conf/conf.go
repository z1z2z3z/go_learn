/*
 * @Description: sql配置
 */
package conf

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func LoadMysqlData() {
	_,filePath,_,_ := runtime.Caller(0)
	fmt.Println(os.Getwd())
	file, err := ini.Load(strings.ReplaceAll(filePath,"conf.go","config.ini"))
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	GetMysqlData(file)
}

func GetMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
