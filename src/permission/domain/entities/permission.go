package entities

type Permission struct {
	ID           int32  `json:"id" gorm:"column:id_permission;primaryKey;autoIncrement"`
	TutoradoID   int32  `json:"id_tutorado_fk" gorm:"column:id_tutorado_fk"`
	DocenteID    int32  `json:"id_docente_fk" gorm:"column:id_docente_fk"`
	Evidencia    string `json:"evidencia" gorm:"column:evidencia"`
	Date         string `json:"date" gorm:"column:date;not null"` // formato YYYY-MM-DD
	Motivo       string `json:"motivo" gorm:"column:motivo"`
	Estatus      string `json:"estatus" gorm:"column:estatus"`
}

// Setters
func (p *Permission) SetID(id int32) {
	p.ID = id
}

func (p *Permission) SetTutoradoID(id int32) {
	p.TutoradoID = id
}

func (p *Permission) SetDocenteID(id int32) {
	p.DocenteID = id
}

func (p *Permission) SetEvidencia(evidencia string) {
	p.Evidencia = evidencia
}

func (p *Permission) SetDate(date string) {
	p.Date = date
}

func (p *Permission) SetMotivo(motivo string) {
	p.Motivo = motivo
}

func (p *Permission) SetEstatus(estatus string) {
	p.Estatus = estatus
}

// Getters
func (p *Permission) GetID() int32 {
	return p.ID
}

func (p *Permission) GetTutoradoID() int32 {
	return p.TutoradoID
}

func (p *Permission) GetDocenteID() int32 {
	return p.DocenteID
}

func (p *Permission) GetEvidencia() string {
	return p.Evidencia
}

func (p *Permission) GetDate() string {
	return p.Date
}

func (p *Permission) GetMotivo() string {
	return p.Motivo
}

func (p *Permission) GetEstatus() string {
	return p.Estatus
}

// Constructor
func NewPermission(tutoradoID, docenteID int32, evidencia, date, motivo, estatus string) *Permission {
	return &Permission{
		TutoradoID: tutoradoID,
		DocenteID:  docenteID,
		Evidencia:  evidencia,
		Date:       date,
		Motivo:     motivo,
		Estatus:    estatus,
	}
}
