package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/leancodebox/fyneMiniProgram/resource"
	"github.com/leancodebox/fyneMiniProgram/tm"
	"os"
)

func main() {
	a := app.New()
	w := a.NewWindow("fyneMiniProgram")
	a.SetIcon(resource.GetLogo())
	a.Settings().SetTheme(&tm.MyTheme{})
	hello := widget.NewLabel("欢迎使用fyneMiniProgram!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("你好", func() {
			hello.SetText("Welcome :)")
		}),
	))
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	// 桌面系统设置托盘
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("bigBrother",
			fyne.NewMenuItem("关于", func() {
				w.Show()
			}),
			fyne.NewMenuItem("退出", func() {
				os.Exit(0)
			}),
		)
		desk.SetSystemTrayMenu(m)
	}
	w.Resize(fyne.NewSize(300, 80))
	w.SetFixedSize(true)
	w.ShowAndRun()
}
