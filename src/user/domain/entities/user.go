package entities

type User struct {
	ID       int32  `json:"id_user" gorm:"column:id_user;primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"column:name;type:varchar(100);not null"`
	LastName string `json:"lastname" gorm:"column:lastname;type:varchar(100);not null"`
	Email    string `json:"email" gorm:"column:email;type:varchar(150);unique;not null"`
	Password string `json:"password" gorm:"column:password;type:varchar(255);not null"`
	Username string `json:"username" gorm:"column:username;type:varchar(100);unique;not null"`
	IdRolFK  int32  `json:"id_rol_fk" gorm:"column:id_rol_fk"`
}

// Setters
func (u *User) SetID(id int32) {
	u.ID = id
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetLastName(lastname string) {
	u.LastName = lastname
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetUsername(username string) {
	u.Username = username
}

func (u *User) SetIdRolFK(idRol int32) {
	u.IdRolFK = idRol
}

// Getters
func (u *User) GetID() int32 {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetIdRolFK() int32 {
	return u.IdRolFK
}

// Constructor
func NewUser(name, lastname, email, password, username string, idRolFK int32) *User {
	return &User{
		Name:     name,
		LastName: lastname,
		Email:    email,
		Password: password,
		Username: username,
		IdRolFK:  idRolFK,
	}
}
