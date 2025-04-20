package database

import "database/sql"

// scan functions could be optimised but its what ever
func ScanMirais(row *sql.Rows, m *Mirais) error {
	return row.Scan(&m.Id, &m.Host, &m.Port, &m.Username, &m.Password, &m.Tag)
}

func ScanQbots(row *sql.Rows, q *Qbots) error {
	return row.Scan(&q.Id, &q.Host, &q.Port, &q.Username, &q.Password, &q.Tag)
}

func ScanUser(row *sql.Rows, ai *User) error {
	return row.Scan(&ai.Id, &ai.Username, &ai.ApiKey, &ai.Conns, &ai.MaxTime, &ai.Rank, &ai.Expire)
}
func ScanAPIS(row *sql.Rows, a *APIS) error {
	return row.Scan(&a.Id, &a.Link, &a.Tag, &a.Active)
}

func ScanSSH(row *sql.Rows, s *SSH_SERVERS) error {
	return row.Scan(&s.Id, &s.Host, &s.Port, &s.Username, &s.Password, &s.Tag)
}
func ScanAttacks(row *sql.Rows, attack *Attacks) error {
	return row.Scan(&attack.Id, &attack.Username, &attack.Host, &attack.Port, &attack.Duration, &attack.Method, &attack.TimeSent)
}

func ScanMethods(row *sql.Rows, method *Methods) error {
	return row.Scan(&method.Id, &method.Methods, &method.Active)
}
