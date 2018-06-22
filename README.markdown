# go blade

 a template engine like Laravel blade
 
 go-blade can translates the syntax of blade into the syntax of Golang html/template

## Installation

```
go get -u github.com/fatrbaby/go-blade
```

## Usage
prepare the directories

`mkdir your-project/{views, caches}`

touch template file views/hello.blade.html

```HTML
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>
{{.Name}}<br>
{{.Age}}<br>
{{ .Hello }}<br>
@if(eq .Age 1)
    hello {{.Age}} years old guy.
@else
    hello world.
@endif
</body>
</html>
```

touch main.go
```go
package main

import (
	"github.com/fatrbaby/go-blade"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", Demo)
	log.Fatal(http.ListenAndServe(":9700", nil))	
}

func Demo(w http.ResponseWriter, _ *http.Request) {
	view := blade.New("views", "caches")
	
	data := blade.H{
		"Name": "fatrbaby",
		"Age": 1,
		"Hello": "what's up my bro",
		"Title": "nobody can fuck with me",
	}
	
	err := view.View("hello", data).Render(w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
```

## Todo

 - template extends
 - range
 - conditional statement compiler
 - abstract cache driver
 - funcs adds