package userservice

import (
	"errors"
	"fmt"
	cs "github.com/huangjiahua/coreshare"
	"github.com/huangjiahua/coreshare/internal/storage"
	"sync"
)

type UserService struct {
	stor storage.KVStorage
	mut  sync.RWMutex
}

func (us *UserService) CreateUser(user cs.User) error {
	us.mut.Lock()
	defer us.mut.Unlock()

	res, err := us.stor.Has(user.Name)

	if err != nil {
		return fmt.Errorf("error accessing storage: %w", err)
	}

	if res {
		return fmt.Errorf("name already exists")
	}

	err = us.stor.Put(user.Name, &user)

	if err != nil {
		return fmt.Errorf("error storing user: %w", err)
	}

	return nil
}

func (us *UserService) UpdateUser(user cs.User) (err error) {
	us.mut.Lock()
	defer us.mut.Unlock()

	ret, err := us.stor.Has(user.Name)
	if err != nil {
		return fmt.Errorf("error accessing storage: %w", err)
	}

	if !ret {
		return fmt.Errorf("name not found")
	}

	err = us.stor.Put(user.Name, &user)

	if err != nil {
		return fmt.Errorf("error storing user: %w", err)
	}

	return nil
}

func (us *UserService) GetUser(name string) (user cs.User, err error) {
	//us.mut.RLock()
	//defer us.mut.RUnlock()

	err = us.stor.Get(name, &user)

	if errors.Is(err, storage.ErrNotFound) {
		return user, fmt.Errorf("name not found")
	}

	if err != nil {
		return user, fmt.Errorf("error accessing storage")
	}

	return
}

func (us *UserService) DeleteUser(name string) (err error) {
	us.mut.Lock()
	defer us.mut.Unlock()

	err = us.stor.Delete(name)

	if errors.Is(err, storage.ErrNotFound) {
		return fmt.Errorf("name not found")
	}

	if err != nil {
		return fmt.Errorf("error accessing storage")
	}

	return nil
}
