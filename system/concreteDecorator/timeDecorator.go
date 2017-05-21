package concrete

import (
	"fmt"
	. "github.com/iHelos/VinoHomework/system"
	"time"
)

func TimeLogDecorator(l LogComponent) LogComponent{
	return LogStr(func (s string){
		s = fmt.Sprintf("%s : %s", time.Now().String(), s)
		l.Log(s)
	})
}
