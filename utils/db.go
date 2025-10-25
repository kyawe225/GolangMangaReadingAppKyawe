package utils

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("pgx", "user=postgres password=kyawe host=localhost port=5432 database=manga_database sslmode=disable")

	if err != nil {
		panic(err)
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(5)

	initializeDatabase()
}

func initializeDatabase() {
	transaction, err := DB.Begin()
	if err != nil {
		panic(err)
	}

	createUserTable := `
		create table if not exists users(
			id varchar(125) primary key not null,
			name varchar(225) not null,
			email varchar(225) not null unique,
			password varchar(225) not null,
			birthdate date not null,
			role varchar(120) not null,
			created_at TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP
		)
	`

	_, err = transaction.Exec(createUserTable)

	if err != nil {
		err1 := transaction.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}

	createMangaTable := `
		create table if not exists mangas(
			id varchar(125) primary key,
			name varchar(255) not null,
			description varchar(255) not null,
			publish_date TIMESTAMP WITH TIME ZONE not null,
			is_published smallint not null default 0,
			published_url varchar(255) not null,
			user_id varchar(125) not null,
			created_at TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP ,
			updated_at TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
			foreign key (user_id) references users(id)
		)
	`

	_, err = transaction.Exec(createMangaTable)

	if err != nil {
		err1 := transaction.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}

	createCategoryTable := `
		create table if not exists category(
			id varchar(125) primary key,
			name varchar(255) not null,
			description varchar(255) not null,
			created_at TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP ,
			updated_at TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP
		)
	`
	_, err = transaction.Exec(createCategoryTable)

	if err != nil {
		err1 := transaction.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}

	createMangaCategoryTable := `
		create table if not exists manga_category(
			id varchar(125) primary key,
			manga_id varchar(125) not null,
			category_id varchar(125) not null,
			created_at TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP ,
			updated_at TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
			foreign key(manga_id) references mangas(id),
			foreign key(category_id) references category(id)
		)
	`
	_, err = transaction.Exec(createMangaCategoryTable)

	if err != nil {
		err1 := transaction.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}

	createChapterTable := `
		create table if not exists chapter(
			id varchar(125) primary key not null,
			manga_id varchar(125) not null,
			chapter_name varchar(255) not null,
			name varchar(225) not null,
			description varchar(255) not null,
			is_published smallint default 0,
			publish_url varchar(125) not null,
			user_id varchar(125) not null,
			created_at TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
			foreign key (user_id) references users(id)
		)
	`
	_, err = transaction.Exec(createChapterTable)

	if err != nil {
		err1 := transaction.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}
	createChapterPictureTable := `
		create table if not exists chapter_pictures(
			id varchar(125) primary key not null,
			picture_data text not null,
			chapter_id varchar(125) not null,
			serial int not null,
			user_id varchar(125) not null,
			created_at TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
			foreign key (user_id) references users(id)
		)
	`

	_, err = transaction.Exec(createChapterPictureTable)

	if err != nil {
		err1 := transaction.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}

	seed(transaction)

	transaction.Commit()
}

func seed(transaction *sql.Tx) {
	checkExists := `select count(*) as j from users where email = $1;`
	var j int64 = -1
	row := transaction.QueryRow(checkExists, "admin@gmail.com")
	row.Scan(&j)
	if j == -1 || j == 0 {
		id := GenerateUUIDV7()
		password, _ := EncryptPassword("password")

		birthdate, err := time.Parse("2006-01-02", "2001-01-01")

		if err != nil {
			err1 := transaction.Rollback()
			if err1 != nil {
				panic(err1)
			}
			panic(err)
		}

		addUserData := `
			insert into users(id,name,email,password,birthdate,role)
			values($1,$2,$3,$4,$5,$6);
		`
		fmt.Println(addUserData)
		_, err = transaction.Exec(addUserData, id, "admin", "admin@gmail.com", password, birthdate, "admin")

		if err != nil {
			err1 := transaction.Rollback()
			if err1 != nil {
				panic(err1)
			}
			panic(err)
		}
	}
}
