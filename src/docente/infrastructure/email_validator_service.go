package infrastructure

import (
	"fmt"
	"strings"
	"time"

	"github.com/Romieb26/ApIsistema_permisos/src/docente/domain"
)

type SimulatedEmailValidator struct{}

func NewSimulatedEmailValidator() domain.IEmailValidator {
	return &SimulatedEmailValidator{}
}

func (s *SimulatedEmailValidator) Validate(email string) (bool, error) {
	fmt.Printf("Simulando verificación de email: %s\n", email)
	
	// Simulamos un retardo de red (opcional para pruebas reales)
	time.Sleep(100 * time.Millisecond)
	
	// Lógica de validación simulada
	if strings.Contains(strings.ToLower(email), "invalid") {
		fmt.Println("Resultado simulado: Email inválido (contiene 'invalid')")
		return false, nil
	}
	
	// Solo falla si el email contiene "test.fail"
	if strings.Contains(strings.ToLower(email), "test.fail") {
		fmt.Println("Resultado simulado: Fallo forzado para testing")
		return false, nil
	}
	
	fmt.Println("Resultado simulado: Email válido")
	return true, nil
}