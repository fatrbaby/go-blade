package blade

import (
	"fmt"
	"io"
)

type View struct {
	source []byte
}

func (view View) Render(writer io.Writer) {

}

func (view View) Strings() string {
	return fmt.Sprintf("%s", view.source)
}
