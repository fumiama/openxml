package openxml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unsafe"
)

var (
	selfclosere = regexp.MustCompile(`></[\w:\s]+>`)
)

// Marshal from v but use tags of t
func Marshal(v, t any) ([]byte, error) {
	data, err := MarshalIndent(v, t, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.ReplaceAll(bytes.ReplaceAll(data, []byte("\t"), nil), []byte("\n"), nil), nil
}

// MarshalIndent from v but use tags of t
func MarshalIndent(v, t any, prefix, indent string) ([]byte, error) {
	tmp := t
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(&tmp), unsafe.Sizeof(uintptr(0)))) = *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(&v), unsafe.Sizeof(uintptr(0))))
	tn := t.(fmt.Stringer).String()
	data, err := xml.MarshalIndent(tmp, prefix, indent)
	if err != nil {
		return nil, err
	}
	nm := reflect.ValueOf(t).Type().String()
	nm = nm[strings.LastIndex(nm, ".")+1:]
	return selfclosere.ReplaceAll(bytes.ReplaceAll(data, []byte(nm), []byte(tn)), []byte("/>")), nil
}
