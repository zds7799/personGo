package utiljson

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strings"
)

const pg string = "utiljson"

var dir string

const confignum int = 4

type config struct {
	remark    string
	dataName  string
	datatype  string
	dataState string
}

// 获取目录下的所有excel表
func getExcels(filename string) []string {
	var allfiles []string
	files, err := os.ReadDir(filename)
	if err != nil {
		fmt.Printf("%s.gtExcels(): file not dir; file: %s\n,%v\n", pg, filename, err)
		return nil
	}
	for _, info := range files {
		if info.IsDir() {
			allfiles = append(allfiles, getExcels(filename+"/"+info.Name())...)
		}
		if strings.HasSuffix(info.Name(), ".xlsx") || strings.HasSuffix(info.Name(), ".xls") {
			f := filename + "/" + info.Name()
			allfiles = append(allfiles, f)
		}
	}
	return allfiles
}

// 解析excel
func parseExcel(filename string) {
	fmt.Printf("begin parse %s\n", filename)
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Printf("%s.parseExcel() filename not excel : %s\n %v\n", pg, filename, err)
		return
	}
	sheetName := f.GetSheetMap()[1]
	rows := f.GetRows(sheetName)
	couLine := len(rows)
	if couLine < 4 {
		fmt.Printf("%s rows is less : %d\n", filename, couLine)
		return
	}
	colNum := len(rows[1])
	confs := make([]config, 0, colNum)
	for i := 0; i < colNum; i++ {
		confs = append(confs, config{})
	}
	datalist := make([][]string, 0, couLine-4)
	for line, data := range rows {
		switch line {
		case 0:
			for i := 0; i < colNum; i++ {
				confs[i].remark = data[i]
			}
		case 1:
			for i := 0; i < colNum; i++ {
				confs[i].dataName = data[i]
			}
		case 2:
			for i := 0; i < colNum; i++ {
				confs[i].datatype = data[i]
			}
		case 3:
			for i := 0; i < colNum; i++ {
				confs[i].dataState = data[i]
			}
		default:
			dataLine := make([]string, 0)
			dataLine = append(dataLine, data...)
			datalist = append(datalist, dataLine)
		}
	}
	toJson(confs, datalist, filename)
	toCSV(confs, datalist, filename)
}

// 解析数据 转json
func toJson(confs []config, datas [][]string, filename string) {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for index, line := range datas {
		if index != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte('{')
		var c = 0
		for i, conf := range confs {
			if conf.dataState != "S" {
				continue
			}
			if c != 0 {
				buf.WriteByte(',')
			}
			c += 1
			buf.WriteByte('"')
			buf.WriteString(conf.dataName)
			buf.WriteByte('"')
			buf.WriteByte(':')
			switch conf.datatype {
			case "string":
				buf.WriteByte('"')
				buf.WriteString(line[i])
				buf.WriteByte('"')
			case "boolean":
				str := line[i]
				if str != "false" && str != "true" {
					fmt.Printf("line %d not discern data : %s\n", index+4, str)
					return
				} else {
					buf.WriteString(line[i])
				}
			case "int":
				buf.WriteString(line[i])
			case "long":
				buf.WriteString(line[i])
			default:
				fmt.Printf("line 3 not discern data type : %s\n", conf.datatype)
				return
			}
		}
		buf.WriteByte('}')
	}
	buf.WriteByte(']')
	bufWriteFile(filename, buf.String())
}
func bufWriteFile(filePath string, str string) {
	filePath = excelToDistPath(filePath, "txt")
	initDir(filePath)
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("创建文件失败！file: %s\n%v", filePath, err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(str)
	if err != nil {
		fmt.Printf("写入文件失败\n %v", err)
		return
	}
	fmt.Printf("json: %s\n", filePath)
}

// 解析数据转csv文件
func toCSV(confs []config, datas [][]string, path string) {
	csvfile := excelToDistPath(path, "csv")
	initDir(csvfile)
	f, err := os.Create(csvfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)

	w.WriteAll(configToString(confs))
	w.WriteAll(datas)
	w.Flush()
	fmt.Printf("csv: %s\n", csvfile)
}

func configToString(confs []config) [][]string {
	var strs [][]string
	for i := 0; i < confignum; i++ {
		strs = append(strs, []string{})
	}
	for _, config := range confs {
		strs[0] = append(strs[0], config.remark)
		strs[1] = append(strs[1], config.dataName)
		strs[2] = append(strs[2], config.datatype)
		strs[3] = append(strs[3], config.dataState)
	}
	return strs
}

func initDir(path string) {
	index := strings.LastIndexByte(path, '.')
	if index >= 0 {
		index = strings.LastIndexByte(path, '/')
		os.MkdirAll(path[:index], os.ModePerm)
	} else {
		os.MkdirAll(path, os.ModePerm)
	}
}

func excelToDistPath(excelPath, postfix string) string {
	index := strings.LastIndexByte(excelPath, '.')
	str := excelPath[len(dir):index]
	var buf bytes.Buffer
	buf.WriteString(dir)
	buf.WriteString("_" + postfix + "/")
	buf.WriteString(str)
	buf.WriteByte('.')
	buf.WriteString(postfix)
	return buf.String()
}

func Parse(path string) {
	//if strings.HasPrefix(path, "/") {
	//	path = path[:len(path)-1]
	//}
	dir = path
	//changecsvandjosnpath(path)
	files := getExcels(path)
	for _, f := range files {
		parseExcel(f)
	}
}
