package user

import model "github.com/antony-raul/tw/models/user"

type Iuser interface {
	ObterUsuarioPeloID(ID *int64) (user *model.ResUsuario, err error)
}
