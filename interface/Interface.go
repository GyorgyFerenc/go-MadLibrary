package Interface

type SignedInteger interface {
	int | int8 | int16 | int32 | int64
}

type UnsignedInteger interface {
	uint | uint16 | uint32 | uint64
}

type AnyInteger interface {
	SignedInteger | UnsignedInteger
}

type Float interface {
	float32 | float64
}

type Number interface {
	Float | AnyInteger
}
