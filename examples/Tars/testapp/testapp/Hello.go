// Package testapp comment
// This file was generated by tars2go 1.1.4
// Generated from Hello.tars
package testapp

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// Message struct implement
type Message struct {
	Nodename  string `json:"Nodename"`
	Podname   string `json:"Podname"`
	Podip     string `json:"Podip"`
	Set       string `json:"Set"`
	Gridkey   string `json:"Gridkey"`
	Gridvalue string `json:"Gridvalue"`
	Content   string `json:"Content"`
}

func (st *Message) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *Message) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_string(&st.Nodename, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Podname, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Podip, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Set, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Gridkey, 5, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Gridvalue, 6, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Content, 7, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *Message) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require Message, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *Message) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Nodename, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Podname, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Podip, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Set, 4)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Gridkey, 5)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Gridvalue, 6)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Content, 7)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *Message) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}