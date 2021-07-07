package ice

import "github.com/gogf/gf/util/gconv"

func isFileIce() bool {
	return true
}
func (s *IceFile) ReadIceHeader() {
	s.Header.Magic = string(s.buffer.Next(3))
	s.buffer.Next(5)
	s.Header.Version = gconv.Int(s.buffer.Next(1))
	s.buffer.Next(15)
	s.Header.Typical = gconv.Int(s.buffer.Next(4))
	s.Header.Blowfish.CompSize = gconv.Int(s.buffer.Next(4))
	s.Header.Blowfish.MagicNumbers = (s.buffer.Next(256))
}
func (s *IceFile) GetBlowfishKeys() {}
func (s *IceFile) Splite()          {}
