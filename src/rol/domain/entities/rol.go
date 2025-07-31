package entities

type Rol struct {
	ID          int32  `json:"id_rol" gorm:"column:id_rol;primaryKey;autoIncrement"`
	Title       string `json:"title" gorm:"column:title;type:varchar(100);not null"`
	Description string `json:"description" gorm:"column:description;type:text"`
}

// Setters
func (r *Rol) SetID(id int32) {
	r.ID = id
}

func (r *Rol) SetTitle(title string) {
	r.Title = title
}

func (r *Rol) SetDescription(description string) {
	r.Description = description
}

// Getters
func (r *Rol) GetID() int32 {
	return r.ID
}

func (r *Rol) GetTitle() string {
	return r.Title
}

func (r *Rol) GetDescription() string {
	return r.Description
}

// Constructor
func NewRol(title, description string) *Rol {
	return &Rol{
		Title:       title,
		Description: description,
	}
}
