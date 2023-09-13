package user

import (
	"errors"
	"go-todo-list/internal/model"
	"log"
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

func TestCreateUser(t *testing.T) {
	dbConnection, mock := ConnectionMock()
	repo := NewRepository(dbConnection)

	type args struct {
		user model.User
	}

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

				mock.ExpectExec("^INSERT INTO `users` \\(`email`,`password`\\) VALUES \\(\\?,\\?\\)$").WithArgs(testCase.args.user.Email, testCase.args.user.Password).WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
				actual, err := repo.Save(testCase.args.user)
				assert.Equal(t, testCase.args.user, actual)
				assert.Nil(t, err)
			} else {
				expectedError := errors.New("database error")
				mock.ExpectBegin()

				mock.ExpectExec("^INSERT INTO `users` \\(`email`,`password`\\) VALUES \\(\\?,\\?\\)$").WithArgs(testCase.args.user.Email, testCase.args.user.Password).WillReturnError(expectedError)

				mock.ExpectCommit()

				actual, err := repo.Save(testCase.args.user)
				assert.Equal(t, testCase.args.user, actual)
				assert.NotNil(t, err)
			}

		})
	}

}
