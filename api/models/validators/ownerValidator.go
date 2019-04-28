package validators

import (
	"errors"
	"go-minikube/api/models"
	"go-minikube/api/utils"
	"strings"
)

var (
	ErrRequiredFirstName = errors.New("required First Name")
	ErrRequiredLastName  = errors.New("required Last Name")
	ErrRequiredEmail     = errors.New("required Email")
	ErrRequiredPassword  = errors.New("required Password")
	ErrRequiredGender    = errors.New("required Gender")
	ErrInvalidGender     = errors.New("invalid Gender. Options: M or F")
	ErrInvalidEmail      = errors.New("invalid Email")
)

func ValidateOwner(owner models.Owner) (models.Owner, error) {
	if err := FieldsRequiredByOwner(owner); err != nil {
		return models.Owner{}, err
	}
	owner.Gender = strings.ToUpper(owner.Gender)
	if owner.Gender != "M" && owner.Gender != "F" {
		return models.Owner{}, ErrInvalidGender
	}
	if !utils.IsEmail(owner.Email) {
		return models.Owner{}, ErrInvalidEmail
	}
	return owner, nil
}

func FieldsRequiredByOwner(owner models.Owner) error {
	if utils.IsEmpty(utils.Trim(owner.FirstName)) {
		return ErrRequiredFirstName
	}
	if utils.IsEmpty(utils.Trim(owner.LastName)) {
		return ErrRequiredLastName
	}
	if utils.IsEmpty(utils.Trim(owner.Email)) {
		return ErrRequiredEmail
	}
	if utils.IsEmpty(strings.TrimSpace(owner.Password)) {
		return ErrRequiredPassword
	}
	if utils.IsEmpty(utils.Trim(owner.Gender)) {
		return ErrRequiredGender
	}
	return nil
}
