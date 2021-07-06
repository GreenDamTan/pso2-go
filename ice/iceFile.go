package ice

import "io"

type IceFile struct {
}

func NewArchive(io.ReadSeeker) (IceFile, error) {
	//TODO need develop
	return IceFile{}, nil
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
