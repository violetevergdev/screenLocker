package main

import (
	"embed"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"golang.org/x/sys/windows/registry"
	"log"
)

var iconFile embed.FS

func main() {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	if _, _, err := key.GetIntegerValue("InactivityTimeoutSecs"); err != nil {
		log.Fatal(err)
	}

	var mw *walk.MainWindow

	MainWindow{
		Title:  "Screen Locker",
		Layout: VBox{},
		Size:   Size{Width: 250, Height: 150},

		Children: []Widget{
			TextLabel{
				Font:    Font{PointSize: 14},
				MaxSize: Size{200, 100},
				Text:    "Screen Locker",
			},
			HSplitter{
				MaxSize: Size{120, 100},
				Children: []Widget{
					PushButton{
						Text:    "Отключить",
						MaxSize: Size{100, 50},
						Font:    Font{PointSize: 10},

						OnClicked: func() {
							err := key.SetDWordValue("InactivityTimeoutSecs", 0)
							if err != nil {
								log.Fatal(err)
							}
						},
					},
					PushButton{
						Text:    "Включить",
						MaxSize: Size{100, 50},
						Font:    Font{PointSize: 10},
						OnClicked: func() {
							err := key.SetDWordValue("InactivityTimeoutSecs", 900)
							if err != nil {
								log.Fatal(err)
							}
						},
					},
				},
			},
		},

		AssignTo: &mw,
	}.Run()

	mw.Run()
}
