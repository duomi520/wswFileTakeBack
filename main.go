package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io/ioutil"
	"os"
)

func main() {
	var outTE *walk.TextEdit
	mw := new(MyMainWindow)
	MainWindow{
		Title:    "简易加密文件",
		AssignTo: &mw.MainWindow,
		MinSize:  Size{300, 140},
		Layout:   VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "选择文件求反",
				OnClicked: func() {
					dlg := new(walk.FileDialog)
					dlg.Title = "选择文件"
					if ok, err := dlg.ShowOpen(mw); err != nil {
						outTE.SetText(err.Error())
					} else if !ok {
						outTE.SetText("未选择文件")
					} else {
						fileName := dlg.FilePath
						file, err := os.Open(fileName)
						defer file.Close()
						if err != nil {
							outTE.SetText("未找到待处理文件")
						}
						//读取文件内容
						plain, _ := ioutil.ReadAll(file)
						//求反
						for i, b := range plain {
							plain[i] = ^b
						}
						//写入文件
						if fileName[len(fileName)-1:] != "+" {
							fileName += "+"
						} else {
							fileName = fileName[:len(fileName)-1]
						}
						err = ioutil.WriteFile(fileName, plain, 0777)
						if err != nil {
							outTE.SetText(dlg.FilePath + " 保存转换后文件失败!")
						} else {
							outTE.SetText(dlg.FilePath + " 文件已转换!")
						}
					}
				},
			},
		},
	}.Run()
}

type MyMainWindow struct {
	*walk.MainWindow
}
