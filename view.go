package blade

import (
	"io"
	"fmt"
)

type View struct {
	source []byte
}

func (view View)Render(writer io.Writer) {
	
}

func (view View)Strings() string {
	return fmt.Sprintf("%s", view.source)
}

