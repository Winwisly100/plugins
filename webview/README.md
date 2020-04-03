# webview

This Go package implements the host-side of the Flutter [webview](https://github.com/Winwisely100/hwv) plugin.

## Usage

Import as:

```go
import webview "github.com/Winwisely100/hwv/go"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&webview.WebviewPlugin{}),
```
