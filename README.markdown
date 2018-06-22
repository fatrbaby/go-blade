# go blade

 a template engine like Laravel blade
 
 go-blade can translates the syntax of blade into the syntax of Golang html/template

## Installation

```
go get -u github.com/fatrbaby/go-blade
```

## Usage
```$xslt
import "github.com/fatrbaby/go-blade"
view := blade.New("view-path", "cache-path")
parsed := view.View("hello", nil).Strings()
fmt.Print(parsed)
```