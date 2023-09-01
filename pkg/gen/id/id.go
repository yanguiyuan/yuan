package id

import (
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	sf = sonyflake.NewSonyflake(sonyflake.Settings{})
}
func One() uint64 {
	id, err := sf.NextID()
	for err != nil {
		id, err = sf.NextID()
	}
	return id
}
