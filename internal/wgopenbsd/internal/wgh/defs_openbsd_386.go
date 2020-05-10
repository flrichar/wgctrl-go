//+build openbsd,386

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs.go

package wgh

const (
	SIOCGIFGMEMB = 0xc024698a

	SizeofIfgreq = 0x10
)

type Ifgroupreq struct {
	Name   [16]byte
	Len    uint32
	Pad1   [0]byte
	Groups *Ifgreq
	Pad2   [12]byte
}

type Ifgreq struct {
	Ifgrqu [16]byte
}

type Timespec struct {
	Sec  int64
	Nsec int32
}
