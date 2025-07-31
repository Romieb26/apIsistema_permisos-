package entities

type Docente struct {
	ID       int32  `json:"id" gorm:"column:id_docente;primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"column:name;not null"`
	LastName string `json:"lastname" gorm:"column:lastname;not null"`
	Email    string `json:"email" gorm:"column:email;not null"`
}

// Setters
func (d *Docente) SetID(id int32) {
	d.ID = id
}

func (d *Docente) SetName(name string) {
	d.Name = name
}

func (d *Docente) SetLastName(lastname string) {
	d.LastName = lastname
}

func (d *Docente) SetEmail(email string) {
	d.Email = email
}

// Getters
func (d *Docente) GetID() int32 {
	return d.ID
}

func (d *Docente) GetName() string {
	return d.Name
}

func (d *Docente) GetLastName() string {
	return d.LastName
}

func (d *Docente) GetEmail() string {
	return d.Email
}

// Constructor
func NewDocente(name, lastname, email string) *Docente {
	return &Docente{
		Name:     name,
		LastName: lastname,
		Email:    email,
	}
}
