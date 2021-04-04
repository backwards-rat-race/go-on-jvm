package types

const (
	void   = "V"
	int    = "I"
	float  = "F"
	double = "D"

	array    = "["
	classRef = "L"

	ConstructorName = "<init>"
)

var (
	Void        = MustParse(void)
	Int         = MustParse(int)
	Float       = MustParse(float)
	Double      = MustParse(double)
	ObjectClass = MustParse("java/lang/Object")
)
