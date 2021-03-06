package revel

import (
	"fmt"
	"regexp"
	"time"
)

type Validator interface {
	IsSatisfied(interface{}) bool
	DefaultMessage() string
}

type Required struct{}

func (r Required) IsSatisfied(obj interface{}) bool {
	if obj == nil {
		return false
	}

	if str, ok := obj.(string); ok {
		return len(str) > 0
	}
	if list, ok := obj.([]interface{}); ok {
		return len(list) > 0
	}
	if b, ok := obj.(bool); ok {
		return b
	}
	if i, ok := obj.(int); ok {
		return i != 0
	}
	if t, ok := obj.(time.Time); ok {
		return !t.IsZero()
	}
	return true
}

func (r Required) DefaultMessage() string {
	return "Required"
}

type Min struct {
	Min int
}

func (m Min) IsSatisfied(obj interface{}) bool {
	num, ok := obj.(int)
	if ok {
		return num >= m.Min
	}
	return false
}

func (m Min) DefaultMessage() string {
	return fmt.Sprintln("Minimum is", m.Min)
}

type Max struct {
	Max int
}

func (m Max) IsSatisfied(obj interface{}) bool {
	num, ok := obj.(int)
	if ok {
		return num <= m.Max
	}
	return false
}

func (m Max) DefaultMessage() string {
	return fmt.Sprintln("Maximum is", m.Max)
}

// Requires an integer to be within Min, Max inclusive.
type Range struct {
	Min
	Max
}

func (r Range) IsSatisfied(obj interface{}) bool {
	return r.Min.IsSatisfied(obj) && r.Max.IsSatisfied(obj)
}

func (r Range) DefaultMessage() string {
	return fmt.Sprintln("Range is", r.Min.Min, "to", r.Max.Max)
}

// Requires an array or string to be at least a given length.
type MinSize struct {
	Min int
}

func (m MinSize) IsSatisfied(obj interface{}) bool {
	if arr, ok := obj.([]interface{}); ok {
		return len(arr) >= m.Min
	}
	if str, ok := obj.(string); ok {
		return len(str) >= m.Min
	}
	return false
}

func (m MinSize) DefaultMessage() string {
	return fmt.Sprintln("Minimum size is", m.Min)
}

// Requires an array or string to be at most a given length.
type MaxSize struct {
	Max int
}

func (m MaxSize) IsSatisfied(obj interface{}) bool {
	if arr, ok := obj.([]interface{}); ok {
		return len(arr) <= m.Max
	}
	if str, ok := obj.(string); ok {
		return len(str) <= m.Max
	}
	return false
}

func (m MaxSize) DefaultMessage() string {
	return fmt.Sprintln("Maximum size is", m.Max)
}

// Requires an array or string to be exactly a given length.
type Length struct {
	N int
}

func (s Length) IsSatisfied(obj interface{}) bool {
	if arr, ok := obj.([]interface{}); ok {
		return len(arr) == s.N
	}
	if str, ok := obj.(string); ok {
		return len(str) == s.N
	}
	return false
}

func (s Length) DefaultMessage() string {
	return fmt.Sprintln("Required length is", s.N)
}

// Requires a string to match a given regex.
type Match struct {
	Regexp *regexp.Regexp
}

func (m Match) IsSatisfied(obj interface{}) bool {
	str := obj.(string)
	return m.Regexp.MatchString(str)
}

func (m Match) DefaultMessage() string {
	return fmt.Sprintln("Must match", m.Regexp)
}

var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")

type Email struct {
	Match
}

func (e Email) DefaultMessage() string {
	return fmt.Sprintln("Must be a valid email address")
}
