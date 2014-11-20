// Copyright 2014 The Go-NetCDF Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// These files are autogenerated from nc_double.go using generate.go
// DO NOT EDIT (except nc_double.go).

package netcdf

import (
	"fmt"
	"unsafe"
)

// #include <stdlib.h>
// #include <netcdf.h>
import "C"

// WriteUint64s writes data as the entire data for variable v.
func (v Var) WriteUint64s(data []uint64) error {
	if err := okData(v, UINT64, len(data)); err != nil {
		return err
	}
	return newError(C.nc_put_var_ulonglong(C.int(v.ds), C.int(v.id), (*C.ulonglong)(unsafe.Pointer(&data[0]))))
}

// ReadUint64s reads the entire variable v into data, which must have enough
// space for all the values (i.e. len(data) must be at least v.Len()).
func (v Var) ReadUint64s(data []uint64) error {
	if err := okData(v, UINT64, len(data)); err != nil {
		return err
	}
	return newError(C.nc_get_var_ulonglong(C.int(v.ds), C.int(v.id), (*C.ulonglong)(unsafe.Pointer(&data[0]))))
}

// WriteUint64s sets the value of attribute a to val.
func (a Attr) WriteUint64s(val []uint64) error {
	// We don't need okData here because netcdf library doesn't know
	// the length or type of the attribute yet.
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	return newError(C.nc_put_att_ulonglong(C.int(a.v.ds), C.int(a.v.id), cname,
		C.nc_type(UINT64), C.size_t(len(val)), (*C.ulonglong)(unsafe.Pointer(&val[0]))))
}

// ReadUint64s reads the entire attribute value into val.
func (a Attr) ReadUint64s(val []uint64) (err error) {
	if err := okData(a, UINT64, len(val)); err != nil {
		return err
	}
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	err = newError(C.nc_get_att_ulonglong(C.int(a.v.ds), C.int(a.v.id), cname,
		(*C.ulonglong)(unsafe.Pointer(&val[0]))))
	return
}

// Uint64sReader is a interface that allows reading a sequence of values of fixed length.
type Uint64sReader interface {
	Len() (n uint64, err error)
	ReadUint64s(val []uint64) (err error)
}

// GetUint64s reads the entire data in r and returns it.
func GetUint64s(r Uint64sReader) (data []uint64, err error) {
	n, err := r.Len()
	if err != nil {
		return
	}
	data = make([]uint64, n)
	err = r.ReadUint64s(data)
	return
}

// testReadUint64s writes somes data to v. N is v.Len().
// This function is only used for testing.
func testWriteUint64s(v Var, n uint64) error {
	data := make([]uint64, n)
	for i := 0; i < int(n); i++ {
		data[i] = uint64(i + 10)
	}
	return v.WriteUint64s(data)
}

// testReadUint64s reads data from v and checks that it's the same as what
// was written by testWriteDouble. N is v.Len().
// This function is only used for testing.
func testReadUint64s(v Var, n uint64) error {
	data := make([]uint64, n)
	if err := v.ReadUint64s(data); err != nil {
		return err
	}
	for i := 0; i < int(n); i++ {
		if val := uint64(i + 10); data[i] != val {
			return fmt.Errorf("data at position %d is %v; expected %v\n", i, data[i], val)
		}
	}
	return nil
}
