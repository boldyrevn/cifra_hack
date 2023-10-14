package crud

import (
	"context"
	"first_goland_project/model"
	"github.com/jackc/pgx/v5"
	"log"
	"net/url"
	"os"
)

type Service struct {
	conn *pgx.Conn
}

func GetService() Service {
	ds := Service{
		conn: connectDb(),
	}
	return ds
}

func connectDb() *pgx.Conn {
	dsn := url.URL{
		Scheme: "postgresql",
		User:   url.UserPassword("postgres", os.Getenv("PG_PASSWORD")),
		Host:   "localhost",
		Path:   "/cifra_db",
	}
	conn, err := pgx.Connect(context.Background(), dsn.String())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connection established")
	return conn
}

func (s *Service) CloseConn() {
	s.conn.Close(context.Background())
}

func (s *Service) GetUser(email string) (model.User, error) {
	row := s.conn.QueryRow(context.Background(), "SELECT full_name, role, balance, email "+
		"FROM users WHERE email = $1", email)
	var u model.User
	if err := row.Scan(&u.Name, &u.Role, &u.Balance, &u.Email); err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (s *Service) CreateUser(u model.CreateUser) (model.User, error) {
	q := "INSERT INTO users(FULL_NAME, EMAIL, HASHED_PASSWORD) VALUES ($1, $2, $3)" +
		"RETURNING full_name, role, balance, email"
	var nu model.User
	err := s.conn.QueryRow(context.Background(), q, u.Name, u.Email,
		u.Password).Scan(&nu.Name, &nu.Role, &nu.Balance, &nu.Email)
	if err != nil {
		log.Println(err)
		return model.User{}, err
	}
	return nu, nil
}

//func (s *Service)
