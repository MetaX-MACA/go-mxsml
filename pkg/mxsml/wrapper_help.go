package mxsml

/*
#cgo CFLAGS: -I.
#include "wrapper_helper.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

type MachineUUIDResult struct {
	LocalUUID   []byte   // local UUID
	RemoteUUIDs [][]byte // remote UUID list
}

func GetMachineUUIDs(maxRemotes uint32, uuidSize uint32) (*MachineUUIDResult, MxSmlReturn) {
	// malloc: local * 1 + remote * maxRemotes
	totalSize := int(1+maxRemotes) * int(uuidSize)
	buffer := make([]byte, totalSize)

	cUuidSize := C.uint(uuidSize)
	cRemoteSize := C.uint(maxRemotes)

	ret := C.mxSmlGetLocalAndMultipleRemoteUuidHelper(
		(*C.uchar)(unsafe.Pointer(&buffer[0])),
		&cRemoteSize,
		&cUuidSize,
	)

	if MxSmlReturn(ret) != MXSML_Success {
		return nil, MxSmlReturn(ret)
	}

	actualUuidSize := int(cUuidSize)
	// actual size larger than the allocated size, this should not happen
	if actualUuidSize > int(uuidSize) {
		return nil, MXSML_Failure
	}

	// first actualUuidSize bytes are local UUID, the rest are remote UUIDs
	result := &MachineUUIDResult{
		LocalUUID: buffer[0:actualUuidSize],
	}

	actualRemotes := int(cRemoteSize)
	result.RemoteUUIDs = make([][]byte, 0, actualRemotes)

	// split remote UUIDs with initial size
	for i := range actualRemotes {
		start := int(uuidSize) + i*int(uuidSize)
		end := start + actualUuidSize
		result.RemoteUUIDs = append(result.RemoteUUIDs, buffer[start:end])
	}

	return result, MXSML_Success
}

func MakeStructVersion(version uint32, structSize uintptr) uint32 {
	return uint32(structSize) | (version << 24)
}
