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

func (s *Service) GetZones() []model.Zone {
	q := `SELECT * FROM zone`
	res := make([]model.Zone, 0)
	rows, _ := s.conn.Query(context.Background(), q)
	defer rows.Close()
	for rows.Next() {
		var z model.Zone
		_ = rows.Scan(&z.ID, &z.Title, &z.CurrentCount, &z.MaxCount)
		res = append(res, z)
	}
	return res
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
	q1 := `DELETE FROM user_zone WHERE user_id = $1 AND zone_id = $2 RETURNING user_id`
	q2 := `UPDATE zone SET current_count = current_count - 1
		WHERE zone.id = $1`
	tx, _ := s.conn.Begin(context.Background())
	defer tx.Rollback(context.Background())
	var uid int
	err := tx.QueryRow(context.Background(), q1, userID, zoneID).Scan(&uid)
	if err == nil {
		_, _ = tx.Exec(context.Background(), q2, zoneID)
		err := tx.Commit(context.Background())
		return err
	}
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

func (s *Service) GetInvitations(id int) []model.Event {
	q := `SELECT event.id, event.description, event.start_date, event.end_date
       FROM "user" JOIN event_invitation ON "user".id = event_invitation.user_id
	JOIN event ON event_invitation.event_id = event.id
	WHERE "user".id = $1`
	res := make([]model.Event, 0)
	rows, _ := s.conn.Query(context.Background(), q, id)
	defer rows.Close()
	for rows.Next() {
		var e model.Event
		_ = rows.Scan(&e.ID, &e.Description, &e.StartDate, &e.EndDate)
		res = append(res, e)
	}
	return res
}

func (s *Service) GetEvents(id int) []model.Event {
	q := `SELECT e.id, e.description, e.start_date, e.end_date
	FROM "user" JOIN public.event_user eu on "user".id = eu.user_id
	JOIN public.event e on eu.event_id = e.id WHERE "user".id = $1`
	res := make([]model.Event, 0)
	rows, _ := s.conn.Query(context.Background(), q, id)
	defer rows.Close()
	for rows.Next() {
		var e model.Event
		_ = rows.Scan(&e.ID, &e.Description, &e.StartDate, &e.EndDate)
		res = append(res, e)
	}
	return res
}
