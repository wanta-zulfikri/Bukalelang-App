package services_test

import (
	"BukaLelang/app/features/users"

	"BukaLelang/app/features/users/services"
	"errors"
	"testing"
	mock "github.com/stretchr/testify/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)
	userServices := services.New(mockRepo)

	// Test Successful registration 
	mockUser := users.Core{Username: "Test User"}
	mockRepo.EXPECT().Register(mockUser).Return(nil, errors.New("Failed to register user"))

	err := userServices.Register(mockUser) 
	assert.NoError(t, err) 

	//Test failed registration
	mockRepo.EXPECT().Register(mockUser).Return(nil, errors.New("Failed to register user"))
	err = userServices.Register(mockUser)
	assert.Error(t, err)
	assert.Equal(t, "Failed to register user", err.Error())
} 

func TestUserService_Loggin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)
	UserServices := services.New(mockRepo)

	// Test Successful login 
	email := "test@example.com"
	password := "password" 
	mockUser := users.Core{Email: email, Password: password}
	mockRepo.EXPECT().Login(email, password).Return(mockUser, nil)
	user, err := UserServices.Login(email, password)
	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)

	// Test Failed login
	mockRepo.EXPECT().Login(email, password).Return(users.Core{}, errors.New("Invalid email or password"))
	user, err = UserServices.Login(email, password) 
	assert.Error(t, err) 
	assert.Equal(t, users.Core{}, user) 
	assert.Equal(t, "Invalid email or password", err.Error())
} 

func TestUserSrvices_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)
	userServices := services.New(mockRepo)

	// Test Successful get profile 
	email := "test@example.com" 
	mockUser := users.Core{Email: email}
	mockRepo.EXPECT().GetProfile(email).Return(mockUser, nil)
	user, err := userServices.GetProfile(email)
	assert.NoError(t, err)
	assert.Equal(t, mockUser, user) 

	// Test failed get profile 
	mockRepo.EXPECT().GetProfile(email).Return(users.Core{}, errors.New("Failed to get user profile"))
	user, err = userServices.GetProfile(email)
	assert.Error(t, err)
	assert.Equal(t, users.Core{}, user)
	assert.Equal(t, "Failed to get user profile", err.Error())

}

func HashedPassword(pasword string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pasword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err 
	}
	return hash, nil
} 

func TestUserService_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)
	UserServices := services.New(mockRepo)

	// Test successful update profile
	email    := "test@example.com"
	username := "newusername"
	newEmail := "newemail@test.com"
	image    := "newimage"
	HashedPassword := []byte("hashedpassword")
	mockRepo.EXPECT(). UpdateProfile(email, users.Core{
		Username: username,
		Email: newEmail,
		Password: string(HashedPassword),
		Image: image,
	}).Return(nil)
	err := UserServices.UpdateProfile(email, username, HashedPassword, image)
	assert.NoError(t, err)


	//Test failed update profile 
	mockRepo.EXPECT(). UpdateProfile(email, users.Core{
		Username: username, 
		Email	: newEmail, 
		Password: string(HashedPassword),
		Image	: image,
	}).Return(nil)
	err = UserServices.UpdateProfile(email, username,newEmail, HashedPassword, image)
	assert.Error(t, err)
	assert.Equal(t, "Error while updating test@test.com: Failed to update user", err.Error())
} 

func TestUserService_DeleteProfile(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockRepository(ctrl)
	UserServices := services.New(mockRepo)

	// Test successful delete profile 
	email := "test@test.com"
	mockRepo.EXPECT().DeleteProfile(email).Return(nil)
	err := UserServices.DeleteProfile(email) 
	assert.Error(t, err)

	// Test successful delete profile 
	mockRepo.EXPECT().DeleteProfile(email).Return(errors.New("Failed to delete user")) 
	err = UserServices.DeleteProfile(email) 
	assert.Error(t, err)
	assert.Equal(t, "Terjadi masalah pada server", err.Error())
}