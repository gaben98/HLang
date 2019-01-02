package variables

import "HLang/typesystem"

//HVar represents a variable instance
type HVar struct {
	htype      *typesystem.IHType
	identifier string
	data       []byte
}

//GlobalVars is the map of all global variables defined
var GlobalVars map[string]*HVar

//DefGlobalVar returns a pointer to a variable instance, or nil if this variable has already been defined
func DefGlobalVar(name string, htype *typesystem.IHType) *HVar {
	if GlobalVars[name] != nil {
		return nil
	}
	v := &HVar{htype: htype, identifier: name, data: nil}
	GlobalVars[name] = v
	return v
}
