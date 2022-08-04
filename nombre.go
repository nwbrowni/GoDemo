package main

// name struct and methods
type nombre struct {
	fname string
	lname string
}

func (n *nombre) SetNombre(fname string, lname string) {
	n.fname = fname
	n.lname = lname
}

func (n *nombre) SetFName(fname string) {
	n.fname = fname
}

func (n *nombre) SetLName(lname string) {
	n.lname = lname
}

func (n *nombre) GetNombreStrings() []string {
	return []string{n.fname, n.lname}
}
