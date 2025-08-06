package entities

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Docente struct {
	ID       int32  `json:"id" gorm:"column:id_docente;primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"column:name;not null"`
	LastName string `json:"lastname" gorm:"column:lastname;not null"`
	Email    string `json:"email" gorm:"column:email;not null"`
}

// ValidarCorreo verifica si el correo es válido para un docente
func (d *Docente) ValidarCorreo() error {
	// Patrón básico para validar formato de email
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, d.Email)
	if !matched {
		fmt.Println("Validación fallida: Formato de email inválido")
		return errors.New("formato de email inválido")
	}

	// Simulación de verificación de dominio docente
	if !strings.HasSuffix(d.Email, "@edu.pe") && !strings.HasSuffix(d.Email, "@educacion.pe") {
		fmt.Printf("Validación fallida: El email %s no es un correo docente válido\n", d.Email)
		return errors.New("el correo no pertenece a un dominio docente válido")
	}

	fmt.Printf("Validación exitosa: El email %s es válido\n", d.Email)
	return nil
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
