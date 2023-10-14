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

func (s *Service) GetUser(conn pgx.Conn, email string) (model.User, error) {
    row := s.conn.QueryRow(context.Background(), "SELECT full_name, hashed_password, role, balance "+
        "FROM users WHERE email = $1", email)
    var u model.User
    if err := row.Scan(&u.Name, &u.HashedPassword, &u.Role, &u.Balance); err != nil {
        return model.User{}, err
    }
    return u, nil
}
