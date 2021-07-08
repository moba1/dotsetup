package dotsetup

type Directory struct {
	Path string
	Mode string
}

func (d *Directory) Commands() []Command {
	opts := []string{}
	if len(d.Mode) != 0 {
		opts = append(opts, "-m", d.Mode)
	}
	rawCommand := append([]string{"mkdir", "-p"}, opts...)
	rawCommand = append(rawCommand, d.Path)
	return []Command{
		{
			rawCommand: rawCommand,
			doRoot:     false,
		},
	}
}
