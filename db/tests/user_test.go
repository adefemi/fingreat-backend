package db_test

import (
	"context"
	db "github/adefemi/fingreat_backend/db/sqlc"
	"github/adefemi/fingreat_backend/utils"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func clean_up() {
	err := testQuery.DeleteAllUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func createRandomUser(t *testing.T) db.User {
	hashedPassword, err := utils.GenerateHashPassword(utils.RandomString(8))

	if err != nil {
		log.Fatal("Unable to generate hash password", err)
	}

	arg := db.CreateUserParams{
		Email:          utils.RandomEmail(),
		HashedPassword: hashedPassword,
	}

	user, err := testQuery.CreateUser(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	assert.Equal(t, user.Email, arg.Email)
	assert.Equal(t, user.HashedPassword, arg.HashedPassword)
	assert.WithinDuration(t, user.CreatedAt, time.Now(), 2*time.Second)
	assert.WithinDuration(t, user.UpdatedAt, time.Now(), 2*time.Second)

	return user
}

func TestCreateUser(t *testing.T) {
	defer clean_up()
	user1 := createRandomUser(t)

	user2, err := testQuery.CreateUser(context.Background(), db.CreateUserParams{
		Email:          user1.Email,
		HashedPassword: user1.HashedPassword,
	})
	assert.Error(t, err)
	assert.Empty(t, user2)
}

func TestUpdateUser(t *testing.T) {
	defer clean_up()
	user := createRandomUser(t)

	newPassword, err := utils.GenerateHashPassword(utils.RandomString(8))

	if err != nil {
		log.Fatal("Unable to generate hash password", err)
	}

	arg := db.UpdateUserPasswordParams{
		HashedPassword: newPassword,
		ID:             user.ID,
		UpdatedAt:      time.Now(),
	}

	newUser, err := testQuery.UpdateUserPassword(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, newUser)
	assert.Equal(t, newUser.HashedPassword, arg.HashedPassword)
	assert.Equal(t, user.Email, newUser.Email)
	assert.WithinDuration(t, user.UpdatedAt, time.Now(), 2*time.Second)
}

func TestGetUserByID(t *testing.T) {
	defer clean_up()
	user := createRandomUser(t)

	newUser, err := testQuery.GetUserByID(context.Background(), user.ID)

	assert.NoError(t, err)
	assert.NotEmpty(t, newUser)

	assert.Equal(t, newUser.HashedPassword, user.HashedPassword)
	assert.Equal(t, user.Email, newUser.Email)
}

func TestGetUserByEmail(t *testing.T) {
	defer clean_up()
	user := createRandomUser(t)

	newUser, err := testQuery.GetUserByEmail(context.Background(), user.Email)

	assert.NoError(t, err)
	assert.NotEmpty(t, newUser)

	assert.Equal(t, newUser.HashedPassword, user.HashedPassword)
	assert.Equal(t, user.Email, newUser.Email)
}

func TestDeleteUser(t *testing.T) {
	defer clean_up()
	user := createRandomUser(t)

	err := testQuery.DeleteUser(context.Background(), user.ID)

	assert.NoError(t, err)

	newUser, err := testQuery.GetUserByID(context.Background(), user.ID)

	assert.Error(t, err)
	assert.Empty(t, newUser)
}

func TestListUsers(t *testing.T) {
	defer clean_up()
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			createRandomUser(t)
		}()
	}

	wg.Wait()

	arg := db.ListUsersParams{
		Offset: 0,
		Limit:  30,
	}

	users, err := testQuery.ListUsers(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, users)
	assert.Equal(t, len(users), 30)
}
