package webview

import (
	"fmt"
	"log"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const (
	channelName = "webview"
)

// WebviewPlugin implements flutter.Plugin and handles method.
type WebviewPlugin struct {
	AppName           string
	BaseDirectoryPath string
}

var _ flutter.Plugin = &WebviewPlugin{
	AppName:           "ION",
	BaseDirectoryPath: "build",
} // compile-time type check

// InitPlugin initializes the plugin.
func (p *WebviewPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getWebview", p.handleWebview)
	return nil
}

func (p *WebviewPlugin) handleWebview(arguments interface{}) (reply interface{}, err error) {

	// Set logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Create astilectron
	a, err := astilectron.New(l, astilectron.Options{
		AppName:           p.AppName,
		BaseDirectoryPath: p.BaseDirectoryPath,
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start
	if err = a.Start(); err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}
	// New window
	var w *astilectron.Window
	url := arguments.(string)

	if w, err = a.NewWindow(url, &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(700),
		Width:  astikit.IntPtr(700),
		WebPreferences: &astilectron.WebPreferences{
			Webaudio: astikit.BoolPtr(true),
			Webgl:    astikit.BoolPtr(true),
		},
	}); err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	// Create windows
	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	// Blocking pattern
	a.Wait()

	return "go-flutter " + flutter.PlatformVersion, nil
}
