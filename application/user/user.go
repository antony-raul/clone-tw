package user

import (
	"context"
	"github.com/antony-raul/tw/config"
	userModel "github.com/antony-raul/tw/models/user"
	userRepo "github.com/antony-raul/tw/repositories/user"
	"golang.org/x/crypto/bcrypt"
)

func ObterUsuarioPeloID(ctx context.Context, ID *int64) (user *userModel.ResUsuario, err error) {
	tx, err := config.NewTransaction(ctx, true)

	repo := userRepo.NewRepo(tx)

	user, err = repo.ObterUsuarioPeloID(ID)
	if err != nil {
		return user, err
	}

	return
}

func CriarUsuario(ctx context.Context, req *userModel.ReqUsuario) (ID int64, err error) {
	tx, err := config.NewTransaction(ctx, false)
	defer tx.Rollback()

	senhaHash, err := bcrypt.GenerateFromPassword([]byte(*req.Senha), bcrypt.DefaultCost-1)
	if err != nil {
		return ID, err
	}
	senhaGerada := string(senhaHash)

	req.Senha = &senhaGerada

	repo := userRepo.NewRepo(tx)

	if ID, err = repo.CriarUsuario(req); err != nil {
		return ID, err
	}

	tx.Commit()

	return
}
