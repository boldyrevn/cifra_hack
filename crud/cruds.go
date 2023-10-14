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
	row := s.conn.QueryRow(context.Background(), `SELECT id, full_name, role, balance, email
        FROM "user" WHERE email = $1`, email)
	var u model.User
	if err := row.Scan(&u.ID, &u.Name, &u.Role, &u.Balance, &u.Email); err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (s *Service) CreateUser(u model.CreateUser) (model.User, error) {
	q := `INSERT INTO "user"(FULL_NAME, EMAIL, HASHED_PASSWORD) VALUES ($1, $2, $3)
        RETURNING id, full_name, role, balance, email`
	var nu model.User
	err := s.conn.QueryRow(context.Background(), q, u.Name, u.Email,
		u.Password).Scan(&nu.ID, &nu.Name, &nu.Role, &nu.Balance, &nu.Email)
	if err != nil {
		log.Println(err)
		return model.User{}, err
	}
	return nu, nil
}

func (s *Service) BookZone(userID, zoneID int) error {
	q := `SELECT book_zone($2, $1);`
	_, err := s.conn.Exec(context.Background(), q, userID, zoneID) // userID, zoneID)
	return err
}

func (s *Service) CheckBooking(userID, zoneID int) (bool, error) {
	q := "SELECT ($1, $2) IN (SELECT * FROM user_zone)"
	var isBooked bool
	err := s.conn.QueryRow(context.Background(), q, userID, zoneID).Scan(&isBooked)
	return isBooked, err
}

func (s *Service) CancelBooking(userID, zoneID int) error {
	q := "DELETE FROM user_zone WHERE user_id = $1 AND zone_id = $2"
	_, err := s.conn.Exec(context.Background(), q, userID, zoneID)
	return err
}

func (s *Service) GetStat(email string) (model.UserStat, error) {
	q := `SELECT us.coffee_cups, us.company_days, us.today_hours
    FROM "user" JOIN public.user_stat us on "user".id = us.user_id
    WHERE "user".email = $1`
	var us model.UserStat
	err := s.conn.QueryRow(context.Background(), q, email).Scan(&us.CoffeeCups, &us.CompanyDays, &us.OfficeHours)
	if err != nil {
		return model.UserStat{}, err
	}
	return us, nil
}
