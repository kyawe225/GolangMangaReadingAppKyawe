package repositories

import (
	"log"
	"realPj/mangaReadingApp/modules/shared/models"
	"realPj/mangaReadingApp/utils"
)

type IUserRepository interface {
	List() (*[]models.User, error)
	Save(model *models.User) error
	Update(id string, model *models.User) error
	Delete(id string) error
	FindById(id string) (*models.User, error)
}

type UserRepository struct {
}

func (user UserRepository) List() (*[]models.User, error) {
	var users []models.User
	query := `
		select * from users;
	`
	rows, err := utils.DB.Query(query)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(user.Id, user.Name, user.Email, user.Password, user.BirthDate, user.Role, user.CreatedAt, user.UpdatedAt)

		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}

func (user UserRepository) Save(model *models.User) error {
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

	_, err = utils.DB.Exec(query, id, model.Name, model.Email, password, model.BirthDate, model.Role)

	if err != nil {
		log.Println(err)
		return err
	}
	model.Id = id
	model.Password = "******"
	return nil
}
func (user UserRepository) Update(id string, model *models.User) error {

	query := `
		update users set name=$1,email=$2,birthdate=$3,role=$4
		where id=$5;
	`

	_, err := utils.DB.Exec(query, model.Name, model.Email, model.BirthDate, model.Role, id)

	if err != nil {
		log.Println(err)
		return err
	}
	model.Id = id
	model.Password = "******"
	return nil
}

func (user UserRepository) Delete(id string) error {
	query := `
		delete from users where id = $1;
	`
	_, err := utils.DB.Exec(query, id)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (repo UserRepository) FindById(id string) (*models.User, error) {
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
