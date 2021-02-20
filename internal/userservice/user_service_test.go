package userservice

import (
	cs "github.com/huangjiahua/coreshare"
	"github.com/huangjiahua/coreshare/internal/storage"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "temp")
	stor, err := storage.NewLevelStorage(filepath.Join(tmpDir, "test-user-service"), nil)
	if !assert.Equal(t, nil, err, "error open storage") {
		t.Fatalf("error open storage: %v", err)
	}

	defer stor.Destroy()
	defer stor.Close()

	us := NewUserService(stor)

	user1 := cs.User{
		Name: "Bill",
		Pass: "pass",
	}

	err = us.CreateUser(user1)

	if err != nil {
		t.Fatalf("error creating user: %v", err)
	}

	user1.Pass = "pass2"

	err = us.CreateUser(user1)

	if err == nil {
		t.Fatalf("create duplicate user")
	}
}

func TestUserService_GetUser(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "temp")
	stor, err := storage.NewLevelStorage(filepath.Join(tmpDir, "test-user-service"), nil)
	if !assert.Equal(t, nil, err, "error open storage") {
		t.Fatalf("error open storage: %v", err)
	}

	us := NewUserService(stor)

	user1 := cs.User{
		Name: "bill",
		Pass: "pass",
	}

	err = us.CreateUser(user1)

	user, err := us.GetUser("bill")

	if err != nil {
		t.Fatalf("error get user: %v", err)
	}

	assert.Equal(t, user1, user, "these users should be equal")
}
