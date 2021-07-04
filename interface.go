package docsetup

type CommandConverter interface {
	ToCommand() []string
}
