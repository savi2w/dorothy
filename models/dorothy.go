package models

type Users struct {
	ID   string
	Role string
}

func GetStoredUser(u string) (*Users, error) {
	user := &Users{ID: u}
	if err := db.Model(user).Column("role").Where("id = ?", u).Select(); err != nil {
		return nil, err
	}

	return user, nil
}

func SetStoredUser(u string, r string) error {
	user := &Users{ID: u, Role: r}

	_, err := db.Model(user).Insert()

	return err
}
