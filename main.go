package main

import (
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

const (
	wndWidth  int = 1024
	wndHeight int = 768
)

func main() {
	opts := &options.App{
		Title:  "assistool",
		Width:  wndWidth,
		Height: wndHeight,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []any{
			NewJSONTab(),
			NewIDTab(),
			NewPercentTab(),
			NewGoStructTab(),
		},
	}
	if err := wails.Run(opts); err != nil {
		fmt.Printf("Error starting app: %v\n", err)
	}
}
