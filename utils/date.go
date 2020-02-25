package utils

import "time"

type DateInterface struct {
  Now string
}
var Date DateInterface

// event loop で初期化する内容
func (d *DateInterface) EventInit() {
  d.Now = time.Now().Format("2006-01-02 15:04:05")
}
