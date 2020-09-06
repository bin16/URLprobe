package sel

const (
	stClass = iota
	stID
	stAttr
	stTag

	opExists     // [a]
	opFullMatch  // =
	opContains   // *=
	opStartsWith // ^=
	opEndsWith   // $=
	opListOf     // |=
)

type selector struct {
	category int
	name     string
	value    string
	operator int
}
