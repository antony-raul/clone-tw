package user

import (
	"github.com/antony-raul/tw/config"
	userModel "github.com/antony-raul/tw/models/user"
	userRepo "github.com/antony-raul/tw/repositories/user"
)

func ObterUsuarioPeloID(ID *int64) (user *userModel.ResUsuario, err error) {
	db, err := config.ConnectDB()
	defer db.Close()

	repo := userRepo.NewRepo(db)

	user, err = repo.ObterUsuarioPeloID(ID)
	if err != nil {
		return user, err
	}

	return
}
