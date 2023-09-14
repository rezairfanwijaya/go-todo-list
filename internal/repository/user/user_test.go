package user

import (
	"errors"
	"go-todo-list/internal/model"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionMock() (*gorm.DB, sqlmock.Sqlmock) {
	sqlMock, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	dialector := mysql.New(mysql.Config{
		DriverName:                "mysql",
		Conn:                      sqlMock,
		SkipInitializeWithVersion: true,
	})

	dbConnection, _ := gorm.Open(dialector, &gorm.Config{})

	return dbConnection, mock
}

type args struct {
	user model.User
}

func TestCreateUser(t *testing.T) {
	dbConnection, mock := ConnectionMock()
	repo := NewRepository(dbConnection)

	testCases := []struct {
		name        string
		args        args
		expectation args
		wantError   bool
	}{
		{
			name: "success",
			args: args{
				user: model.User{
					Email:    "johntest@gmail.com",
					Password: "1234567",
				},
			},
			expectation: args{
				user: model.User{
					Email:    "johntest@gmail.com",
					Password: "1234567",
				},
			},
			wantError: false,
		}, {
			name: "failed",
			args: args{
				user: model.User{
					Email:    "",
					Password: "",
				},
			},
			expectation: args{
				user: model.User{
					Email:    "",
					Password: "",
				},
			},
			wantError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if !testCase.wantError {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`email`,`password`) VALUES (?,?)")).WithArgs(testCase.args.user.Email, testCase.args.user.Password).WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()

				actual, err := repo.Save(testCase.args.user)
				assert.Equal(t, testCase.args.user, actual)
				assert.Nil(t, err)
			} else {
				expectedError := errors.New("database error")
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`email`,`password`) VALUES (?,?)")).WithArgs(testCase.args.user.Email, testCase.args.user.Password).WillReturnError(expectedError)

				mock.ExpectCommit()

				actual, err := repo.Save(testCase.args.user)
				assert.Equal(t, testCase.args.user, actual)
				assert.NotNil(t, err)
			}

		})
	}

}

func TestGetUserByEmail(t *testing.T) {
	dbConnection, mock := ConnectionMock()
	repo := NewRepository(dbConnection)

	testCases := []struct {
		name        string
		email       string
		expectation args
		wantError   bool
	}{
		{
			name:  "sucess",
			email: "john@gmail.com",
			expectation: args{
				user: model.User{
					ID:    1,
					Email: "john@gmail.com",
				},
			},
			wantError: false,
		}, {
			name:  "user not found",
			email: "doe@gmail.com",
			expectation: args{user: model.User{
				ID:    0,
				Email: "",
			}},
			wantError: false,
		}, {
			name:        "failed",
			email:       "",
			expectation: args{user: model.User{}},
			wantError:   true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if !testCase.wantError {

				rows := sqlmock.NewRows([]string{"id", "email"}).
					AddRow(
						testCase.expectation.user.ID,
						testCase.expectation.user.Email,
					)

				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ?")).WithArgs(testCase.email).WillReturnRows(rows)

				actual, err := repo.GetUserByEmail(testCase.email)
				assert.Equal(t, testCase.expectation.user, actual)
				assert.Nil(t, err)
			} else {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ?")).WithArgs(testCase.email).WillReturnError(errors.New("failed"))

				actual, err := repo.GetUserByEmail(testCase.email)
				assert.Equal(t, testCase.expectation.user, actual)
				assert.NotNil(t, err)
			}
		})
	}
}

func TestUserByID(t *testing.T) {
	dbConnection, mock := ConnectionMock()
	repo := NewRepository(dbConnection)

	testCase := []struct {
		name        string
		ID          uint
		expectation args
		wantError   bool
	}{
		{
			name: "success",
			ID:   1,
			expectation: args{
				user: model.User{
					ID:       1,
					Email:    "john@gmail.com",
					Password: "12345678",
				},
			},
			wantError: false,
		},
		{
			name: "failed id not found",
			ID:   90,
			expectation: args{
				user: model.User{
					ID:       0,
					Email:    "",
					Password: "",
				},
			},
			wantError: true,
		},
	}

	for _, testCase := range testCase {
		t.Run(testCase.name, func(t *testing.T) {
			if !testCase.wantError {
				rows := sqlmock.NewRows([]string{"id", "email", "password"}).
					AddRow(
						testCase.expectation.user.ID,
						testCase.expectation.user.Email,
						testCase.expectation.user.Password,
					)

				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ?")).WillReturnRows(rows)

				actual, err := repo.GetUserByID(testCase.ID)
				assert.Equal(t, testCase.expectation.user, actual)
				assert.Nil(t, err)
			} else {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ?")).WillReturnError(errors.New("failed get id"))

				actual, err := repo.GetUserByID(testCase.ID)
				assert.Equal(t, testCase.expectation.user, actual)
				assert.NotNil(t, err)
			}
		})
	}
}
