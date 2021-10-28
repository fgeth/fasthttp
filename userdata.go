package fasthttp

import (
	"io"
)
//update to upper case
type userDataKV struct {
	Key   []byte
	Value interface{}
}

type userData []userDataKV

func (d *userData) Set(key string, value interface{}) {
	args := *d
	n := len(args)
	for i := 0; i < n; i++ {
		kv := &args[i]
		if string(kv.Key) == key {
			kv.Value = value
			return
		}
	}

	if value == nil {
		return
	}

	c := cap(args)
	if c > n {
		args = args[:n+1]
		kv := &args[n]
		kv.Key = append(kv.Key[:0], key...)
		kv.Value = value
		*d = args
		return
	}

	kv := userDataKV{}
	kv.Key = append(kv.Key[:0], key...)
	kv.Value = value
	*d = append(args, kv)
}

func (d *userData) SetBytes(key []byte, value interface{}) {
	d.Set(b2s(key), value)
}

func (d *userData) Get(key string) interface{} {
	args := *d
	n := len(args)
	for i := 0; i < n; i++ {
		kv := &args[i]
		if string(kv.Key) == key {
			return kv.Value
		}
	}
	return nil
}

func (d *userData) GetBytes(key []byte) interface{} {
	return d.Get(b2s(key))
}

func (d *userData) Reset() {
	args := *d
	n := len(args)
	for i := 0; i < n; i++ {
		v := args[i].Value
		if vc, ok := v.(io.Closer); ok {
			vc.Close()
		}
	}
	*d = (*d)[:0]
}

func (d *userData) Remove(key string) {
	args := *d
	n := len(args)
	for i := 0; i < n; i++ {
		kv := &args[i]
		if string(kv.Key) == key {
			n--
			args[i] = args[n]
			args[n].Value = nil
			args = args[:n]
			*d = args
			return
		}
	}
}

func (d *userData) RemoveBytes(key []byte) {
	d.Remove(b2s(key))
}
