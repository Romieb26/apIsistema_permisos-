package entities

type Tutorado struct {
	ID        int32  `json:"id_tutorado" gorm:"column:id_tutorado;primaryKey;autoIncrement"`
	IdUserFK  int32  `json:"id_user_fk" gorm:"column:id_user_fk"`
	Name      string `json:"name" gorm:"column:name;not null"`
	LastName  string `json:"lastname" gorm:"column:lastname;not null"`
	Matricula string `json:"matricula" gorm:"column:matricula;not null;unique"`
	Email     string `json:"email" gorm:"column:email;not null"`
	Estatus   string `json:"estatus" gorm:"column:estatus"`
}

// Setters
func (t *Tutorado) SetID(id int32) {
	t.ID = id
}

func (t *Tutorado) SetIdUserFK(idUser int32) {
	t.IdUserFK = idUser
}

func (t *Tutorado) SetName(name string) {
	t.Name = name
}

func (t *Tutorado) SetLastName(lastname string) {
	t.LastName = lastname
}

func (t *Tutorado) SetMatricula(matricula string) {
	t.Matricula = matricula
}

func (t *Tutorado) SetEmail(email string) {
	t.Email = email
}

func (t *Tutorado) SetEstatus(estatus string) {
	t.Estatus = estatus
}

// Getters
func (t *Tutorado) GetID() int32 {
	return t.ID
}

func (t *Tutorado) GetIdUserFK() int32 {
	return t.IdUserFK
}

func (t *Tutorado) GetName() string {
	return t.Name
}

func (t *Tutorado) GetLastName() string {
	return t.LastName
}

func (t *Tutorado) GetMatricula() string {
	return t.Matricula
}

func (t *Tutorado) GetEmail() string {
	return t.Email
}

func (t *Tutorado) GetEstatus() string {
	return t.Estatus
}

func NewTutorado(idUser int32, name, lastname, matricula, email, estatus string) *Tutorado {
	return &Tutorado{
		IdUserFK:  idUser,
		Name:      name,
		LastName:  lastname,
		Matricula: matricula,
		Email:     email,
		Estatus:   estatus,
	}
}
