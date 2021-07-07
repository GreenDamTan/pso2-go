package ice

import (
	"bytes"
	"github.com/gogf/gf/os/glog"
	"io"
)

type IceFile struct {
	buffer *bytes.Buffer
	Header struct {
		Magic    string
		Version  int
		Typical  int
		Blowfish struct {
			CompSize     int
			MagicNumbers []byte
		}
	}
}

func NewArchive(unknown1 []byte) (IceFile, error) {
	//TODO need develop
	a := IceFile{buffer: bytes.NewBuffer(unknown1)}
	a.ReadIceHeader()
	glog.Line().Debug(a)
	return a, nil
}
func (ice IceFile) GroupCount() int {
	//TODO need develop
	return 0
}
func (ice IceFile) Group(count int) IceGroup {
	//TODO need develop
	return IceGroup{}
}
func (ice IceFile) ReplaceFile(file *IceGroupFile, unkown2 interface{}, unknown3 uint32) {
	//TODO need develop
	return
}
func (ice IceFile) Write(unknown io.Writer) {
	//TODO need develop
	return
}

type IceGroup struct {
	Files []IceGroupFile
}
type IceGroupFile struct {
	Name string
	Type string
	Size int
	Data io.Reader
}
