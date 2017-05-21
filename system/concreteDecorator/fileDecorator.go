package concrete

import (
	. "github.com/iHelos/VinoHomework/system"
	logsmodel "github.com/iHelos/VinoHomework/model/logs"
)

func DBLogDecorator(l LogComponent) LogComponent{
	return LogStr(func (s string){
		logsmodel.NewDBLog(s).Insert()
		l.Log(s)
	})
}

