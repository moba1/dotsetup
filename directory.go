package docsetup

type Directory struct {
	Path string
	Mode string
}

func (d *Directory) Command() []RawCommand {
	opts := []string{}
	if len(d.Mode) != 0 {
		opts = append(opts, "-m", d.Mode)
	}
	command := append([]string{"mkdir", "-p"}, opts...)
	command = append(command, d.Path)
	return []RawCommand{
		command,
	}
}
