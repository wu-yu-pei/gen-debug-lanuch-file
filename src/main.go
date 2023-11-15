package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const GEN_DIR_NAME = ".vscode"
const GEN_FILE_NAME = "launch.json"

type LaunchFile struct {
	Configurations []ConfigItem `json:"configurations"`
}

type ConfigItem struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Request string `json:"request"`
	Command string `json:"command"`
	Cwd     string `json:"cwd"`
}

func main() {
	// 解析命令行参数 --scripts dev,pro
	var scripts string
	flag.StringVar(&scripts, "scripts", "dev,pro", "需要生成debug配置的脚本")
	flag.Parse()
	scriptArray := strings.Split(scripts, ",")
	fmt.Println(scriptArray)

	currentDir, _ := os.Getwd()
	filePath := currentDir + "/" + GEN_DIR_NAME + "/" + GEN_FILE_NAME

	fmt.Println(currentDir)

	dirEntries, _ := os.ReadDir(currentDir)

	dirNames := getDirs(dirEntries)

	if isInArray(GEN_DIR_NAME, dirNames) {
		if !fileExists(filePath) {
			fmt.Println("no launch file create...")
			file, _ := os.Create(filePath)
			fmt.Println("create launch file success!")

			// 打开文件 写内容
			defer file.Close()

			file.WriteString(getLaunchFileContent(scriptArray))

			fmt.Println("write launch file content success!")

			return
		}
		// 更新
		fmt.Println("launch.json 文件已存在！")

		//file, _ := ioutil.ReadFile(filePath)
		//var config LaunchFile
		//json.Unmarshal(file, &config)
		//
		//for i := 0; i < len(scriptArray); i++ {
		//	var configItem ConfigItem
		//	jsonString := getConfigItem(scriptArray[i])
		//	fmt.Println(jsonString)
		//	json.Unmarshal([]byte(jsonString), &configItem)
		//
		//	config.Configurations = append(config.Configurations, configItem)
		//}
		//
		//jsonString, _ := json.Marshal(config)
		//fmt.Println(string(jsonString))
		//_file, _ := os.OpenFile(filePath, os.O_WRONLY, 0644)
		//_file.WriteString(string(jsonString))
		//fmt.Println(jsonString, "--")
		//ioutil.WriteString(filePath, jsonString, 0644)
	} else {
		// 创建文件夹
		err := os.Mkdir(currentDir+"/"+GEN_DIR_NAME, 0755)
		if err != nil {
			fmt.Println("create .vscode folder fail! err: ", err)
		}

		fmt.Println("create .vscode folder success!")

		// 创建文件
		file, _ := os.Create(filePath)

		fmt.Println("create launch file success!")

		fmt.Println(file)

		// 写入文件
		defer file.Close()

		file.WriteString(getLaunchFileContent(scriptArray))

		fmt.Println("write launch file content success!")
	}
}

func isInArray(target string, array []string) bool {
	isIn := false
	for _, arrayItem := range array {
		if arrayItem == target {
			isIn = true
		}
	}
	return isIn
}

func getDirs(dirEntries []os.DirEntry) (dirs []string) {
	for _, entry := range dirEntries {
		dirs = append(dirs, entry.Name())
	}
	return
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func getLaunchFileContent(scriptArray []string) (content string) {
	content = `{
	"configurations": [
	`

	for i := 0; i < len(scriptArray); i++ {
		content += getConfigItem(scriptArray[i])
	}

	content += `]
}
	`
	return
}

func getConfigItem(script string) (content string) {
	content = `	{
			"type": "node-terminal",
			"name": "运行脚本: ` + script + `",
			"request": "launch",
			"command": "npm run ` + script + `",
			"cwd": "${workspaceFolder}"
		},
	`
	return
}
