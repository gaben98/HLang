package typesystem

import (
	"errors"
	"fmt"
	"strings"
)

//HType is a struct denoting an atomic type within the type tree
type HType struct {
	children   []*HType
	parent     *HType
	checkValue int64
	name       string
	defaultVal []byte
}

//IHType is the interface for all types within the typesystem
type IHType interface {
	Is(IHType) bool
}

var nextPrime = erastosthenes()
var curPrime = nextPrime()

//Object is the root type of the type tree
var Object = HType{children: nil, parent: nil, checkValue: 2, name: "Object"}

var typeMap = map[string]*HType{"Object": &Object}

//Is checks if this type inherits from or directly is another type
func (htype HType) Is(other IHType) bool {
	ht, ok := other.(HType)
	if !ok {
		return false
	}
	return ht.checkValue%htype.checkValue == 0
}

//DefinePrimitive defines a primitive type inheriting from object with a default value
func DefinePrimitive(name string, defaultVal interface{}) *HType {
	htype := Object.Define(name)
	htype.MakeDefault(defaultVal)
	return htype
}

//MakeDefault makes a default value for a type when created
func (htype *HType) MakeDefault(data interface{}) error {
	btptrs, ok := data.([]*byte)
	if !ok {
		return errors.New("data could not be interpreted")
	}
	htype.defaultVal = make([]byte, 0)
	for _, bptr := range btptrs {
		htype.defaultVal = append(htype.defaultVal, *bptr)
	}
	return nil
}

//Define creates a type that inherits from a parent type
func (htype *HType) Define(name string) *HType {
	if typeMap[name] != nil {
		return nil
	}
	nType := &HType{children: nil, checkValue: curPrime, name: name, parent: htype, defaultVal: nil}
	if htype.children == nil {
		nType.checkValue = htype.checkValue
		for node := htype; node != nil; node = node.parent {
			node.checkValue *= nType.checkValue
		}
	} else if len(htype.children) == 1 { //if have 1 child, the parent has a prime factor to a power, so I need the new child to get a unique prime and to bubble up
		nPrime := nextPrime()
		oPrime := htype.children[0].checkValue
		nType.checkValue = nPrime
		for node := htype; node != nil; node = node.parent {
			node.checkValue /= oPrime
			node.checkValue *= nPrime
		}
		curPrime = nPrime
	} else { //I have multiple children, I can add a node with a unique prime and bubble up normally
		curPrime = nextPrime()
		nType.checkValue = curPrime
		for node := htype; node != nil; node = node.parent {
			node.checkValue *= curPrime
		}
	}
	htype.children = append(htype.children, nType)
	typeMap[nType.name] = nType
	return nType
}

//PrintObjectTree returns a string representation of the subtree of the htype specified
func PrintObjectTree(htype *HType) string {
	if htype.children == nil {
		return fmt.Sprintf("%d: %s", htype.checkValue, htype.name)
	}
	return fmt.Sprintf("%d: %s(%s)", htype.checkValue, htype.name, strings.Join(forall(htype.children, PrintObjectTree), ", "))
}
