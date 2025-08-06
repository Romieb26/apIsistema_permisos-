package application

import (
	"errors"
	"fmt"

	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain"
	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain/entities"
)

type SaveDocenteUseCase struct {
	Repository domain.IDocente
	EmailService domain.IEmailValidator
}

func NewSaveDocenteUseCase(repo domain.IDocente, emailService domain.IEmailValidator) *SaveDocenteUseCase {
	return &SaveDocenteUseCase{
		Repository: repo,
		EmailService: emailService,
	}
}

func (uc *SaveDocenteUseCase) Execute(docente *entities.Docente) error {
	fmt.Println("Iniciando validación de correo docente...")
	
	if err := docente.ValidarCorreo(); err != nil {
		fmt.Printf("Error en validación básica: %v\n", err)
		return fmt.Errorf("validación fallida: %w", err)
	}

	isValid, err := uc.EmailService.Validate(docente.Email)
	if err != nil {
		fmt.Printf("Error en servicio de validación: %v\n", err)
		return fmt.Errorf("error en servicio de validación: %w", err)
	}

	if !isValid {
		fmt.Printf("Email rechazado por el servicio: %s\n", docente.Email)
		return errors.New("el correo fue rechazado por el servicio de validación")
	}
	
	fmt.Printf("Email %s aprobado, procediendo a guardar\n", docente.Email)
	return uc.Repository.Save(docente)
}