package usecases

import (
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/domain/repositories"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo repositories.UserRepo
	companyRepo repositories.CompanyRepo
}

func NewUserUsecase(userRepo repositories.UserRepo, companyRepo repositories.CompanyRepo) *UserUsecase {
	return &UserUsecase{userRepo: userRepo, companyRepo: companyRepo}
}

func (u *UserUsecase) RegisterUser(userRegisterInfo models.RegisterUserDTO) (*models.JwtUserDTO, error) {
	existingUser, err := u.userRepo.GetUserByUsername(userRegisterInfo.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username is taken")
	}

	company, err := u.companyRepo.GetCompanyByID(userRegisterInfo.CompanyID)
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, errors.New("company does not exist")
	}

	role, err := u.userRepo.GetRoleByID(userRegisterInfo.RoleID)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, errors.New("role does not exist")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userRegisterInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var registeredUser = models.InsertUserDTO{
		ID: uuid.New(),
		Username: userRegisterInfo.Username,
		FirstName: userRegisterInfo.FirstName,
		LastName: userRegisterInfo.LastName,
		PasswordHash: string(passwordHash),
		RoleID: userRegisterInfo.RoleID,
		CompanyID: userRegisterInfo.CompanyID,
	}

	err = u.userRepo.InsertUser(&registeredUser)
	if err != nil {
		return nil, err
	}

	return &models.JwtUserDTO{
		ID: registeredUser.ID,
		Username: registeredUser.Username,
		RoleTitle: role.Title,
	}, nil
}

func (u *UserUsecase) LoginUser(userLoginInfo models.LoginUserDTO) (*models.JwtUserDTO, error) {
	existingUser, err := u.userRepo.GetUserByUsername(userLoginInfo.Username)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("user does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(userLoginInfo.Password))
	if err != nil {
		return nil, err
	}

	role, err := u.userRepo.GetRoleByID(existingUser.RoleID)
	if err != nil {
		return nil, err
	}

	return &models.JwtUserDTO{
		ID: existingUser.ID,
		Username: existingUser.Username,
		RoleTitle: role.Title,
	}, nil
}