package randomKit

func Bool() bool {
	i := Int(0, 2) // [0, 2)
	return i == 1
}
