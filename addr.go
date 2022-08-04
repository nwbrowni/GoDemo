package main

// address struct and methods
type addr struct {
	street string
	apt    string
	city   string
	state  string
	zip    string
}

func (a *addr) SetAddress(street string, apt string, city string, state string, zip string) {
	a.street = street
	a.apt = apt
	a.city = city
	a.state = state
	a.zip = zip
}

func (a *addr) SetStreet(street string) {
	a.street = street
}

func (a *addr) SetApt(apt string) {
	a.apt = apt
}

func (a *addr) SetCity(city string) {
	a.city = city
}

func (a *addr) SetState(state string) {
	a.state = state
}

func (a *addr) SetZip(zip string) {
	a.zip = zip
}

func (a *addr) GetAddressStrings() []string {
	return []string{a.street, a.apt, a.city, a.state, a.zip}
}
