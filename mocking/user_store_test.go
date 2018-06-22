package main

import (
	"testing"

	"github.com/jacobhaven/slides/mocking/pkg/db"
	"github.com/jacobhaven/slides/mocking/pkg/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserStore(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	var (
		userID    uint = 1
		userName       = "Malcolm Reynolds"
		userEmail      = "capn.mal@example.com"
		user           = &db.User{userID, userName, userEmail}
	)
	store := mocks.NewMockUserStore(mockCtrl)
	gomock.InOrder(
		store.EXPECT().Create(userName, userEmail).Return(user),
		store.EXPECT().Get(userID).Return(user, nil),
	)

	createdUser := store.Create(userName, userEmail)
	require.Equal(t, userName, createdUser.Name)
	require.Equal(t, userEmail, createdUser.Email)
	require.Equal(t, userID, createdUser.ID)

	gotUser, err := store.Get(createdUser.ID)
	require.NoError(t, err)
	require.Equal(t, createdUser, gotUser)
}
