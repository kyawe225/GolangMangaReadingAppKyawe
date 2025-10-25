package repositories

import (
	"log"
	"realPj/mangaReadingApp/modules/shared/dtos"
	"realPj/mangaReadingApp/modules/shared/models"
	"realPj/mangaReadingApp/utils"
)

type IAuthRepository interface {
	Register(model *dtos.RegisterDto) error
	Login(model *dtos.LoginDto) (*string, error)
	Profile(id string) (*models.User, error)
}

type AuthRepository struct {
}

func (auth AuthRepository) Register(model *dtos.RegisterDto) error {
	id := utils.GenerateUUIDV7()

	password, err := utils.EncryptPassword(model.Password)
	if err != nil {
		log.Println(err)
		return err
	}
	query := `
		insert into users(id,name,email,password,birthdate,role)
		values($1,$2,$3,$4,$5,$6);
	`
	_, err = utils.DB.Exec(query, id, model.Name, model.Email, password, model.BirthDate, "user")

	if err != nil {
		log.Println(err)
		return err
	}

	model.Role = "user"
	model.Password = "*******"

	return nil
}

func (auth AuthRepository) Login(model *dtos.LoginDto) (*string, error) {
	query := `
		select * from users where email = $1;
	`
	row := utils.DB.QueryRow(query, model.Email)

	var user models.User
	row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.BirthDate, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	err := utils.CheckPassword(model.Password, user.Password)

	if err != nil {
		return nil, err
	}

	tokenString, err := utils.GenerateToken(&user)

	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (repo AuthRepository) Profile(id string) (*models.User, error) {
	var user models.User
	var n string
	query := `select *
	from users
	where id = $1
	limit 1;`
	resultRow := utils.DB.QueryRow(query, id)
	err := resultRow.Scan(&user.Id, &user.Name, &user.Email, n, &user.BirthDate, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
