package docsetup

type Command interface {
	Script() []string
}
