package user

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	userInterface "github.com/antony-raul/tw/interfaces/user"
	userModel "github.com/antony-raul/tw/models/user"
)

type repository struct {
	builder sq.StatementBuilderType
}

func NewRepo(db *sql.DB) userInterface.Iuser {
	return &repository{
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db),
	}
}

func (r *repository) ObterUsuarioPeloID(ID *int64) (user *userModel.ResUsuario, err error) {
	user = new(userModel.ResUsuario)

	if err = r.builder.Select("id",
		"nome",
		"email",
		"data_criacao").
		From("public.t_usuario").
		Where(
			sq.Eq{
				"id": ID,
			}).QueryRow().
		Scan(&user.ID, &user.Nome, &user.Email, &user.DataCriacao); err != nil {
		return user, err
	}
	return
}
