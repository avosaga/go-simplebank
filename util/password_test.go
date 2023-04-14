package util

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCorrectPassword(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)
}

func TestIncorrectPassword(t *testing.T) {
	originalPassword := RandomString(6)

	hashedPassword, err := HashPassword(originalPassword)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}

func TestHashedPasswordReturnDifferentValueWithSamePassword(t *testing.T) {
	password := RandomString(6)

	hashedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)

	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
