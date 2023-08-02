package postgres

import (
	"app/api/models"
	"context"
	"testing"
)

func TestCreateUser(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CreateUser
		Output  string
		WantErr bool
	}{
		{
			Name: "User Case 1",
			Input: &models.CreateUser{
				FirstName:   "Tursinbek",
				LastName:    "Aytbaev",
				Login:       "rocket123",
				Password:    "1234567",
				PhoneNumber: "+998995430419",
			},
			WantErr: false,
		},
		{
			Name: "User Case 2",
			Input: &models.CreateUser{
				FirstName:   "Shohruh",
				LastName:    "Ramazonov",
				Login:       "shoh123",
				Password:    "11223344",
				PhoneNumber: "+123123123",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			id, err := userTestRepo.Create(context.Background(), test.Input)

			if test.WantErr || err != nil {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if id == "" {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}
		})
	}
}

func TestGetByIDUser(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.UserPrimaryKey
		Output  *models.User
		WantErr bool
	}{
		{
			Name:  "Product Case 1",
			Input: &models.UserPrimaryKey{Id: "52fb8939-e3d8-402c-9fd6-a2714052b6c9"},
			Output: &models.User{
				Id:          "52fb8939-e3d8-402c-9fd6-a2714052b6c9",
				FirstName:   "Abdurashid",
				LastName:    "Arzubaev",
				Login:       "abdu04042",
				Password:    "abdu04abdu",
				PhoneNumber: "+998975062404",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			user, err := userTestRepo.GetByID(context.Background(), test.Input)

			if test.WantErr || err != nil {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if user.FirstName != test.Output.FirstName ||
				user.LastName != test.Output.LastName ||
				user.Id != test.Output.Id ||
				user.Login != test.Output.Login ||
				user.Password != test.Output.Password ||
				user.PhoneNumber != test.Output.PhoneNumber {

				t.Errorf("%s: got: %+v, expected: %+v\n", test.Name, *user, *&test.Output)
				return
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	tests := []struct {
		Name    string
		Input   *models.UpdateUser
		Output  int64
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.UpdateUser{
				FirstName:   "Tursinbek",
				LastName:    "Aytbaev",
				Login:       "rocket123",
				Password:    "1234567",
				PhoneNumber: "+998995430419",
			},
			Output:  1,
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			rows, err := userTestRepo.Update(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if rows != test.Output {
				t.Errorf("%s: got: %v, expected: %v", test.Name, rows, test.Output)
				return
			}

		})
	}
}
