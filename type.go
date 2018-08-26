package script

type Type interface {
	Name() string
	Equals(interface{}) bool
	String() string
}
