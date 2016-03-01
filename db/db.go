package db

import "database/sql"

type Member struct {
	ID             int
	Name           string
	UUID           string
	IP             string
	RequestedIndex int
	Heartbeat      int
	Ports          string
	Index          int
}

type DB struct {
	db *sql.DB
}

func New(driverName, dsn string) (*DB, error) {
	db, err := sql.Open(driverName, dsn)
	return &DB{
		db: db,
	}, err
}

type Members []Member

func (a Members) Len() int           { return len(a) }
func (a Members) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Members) Less(i, j int) bool { return a[i].ID < a[j].ID }

func (d *DB) Members() ([]Member, error) {
	rows, err := d.db.Query(`SELECT id, name, heartbeat, uuid, assigned_index, requested_index, ports, ip_address
		FROM cluster ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []Member{}

	for rows.Next() {
		member := Member{}
		if err := rows.Scan(&member.ID, &NullStringWrapper{String: &member.Name}, &member.Heartbeat, &member.UUID, &member.Index, &member.RequestedIndex, &member.Ports,
			&member.IP); err != nil {
			return nil, err
		}
		result = append(result, member)
	}

	return result, rows.Err()
}

func (d *DB) Delete(uuid string) error {
	_, err := d.execCount(`DELETE FROM cluster WHERE uuid = ?`, uuid)
	return err
}

func (d *DB) Checkin(member Member, i int) error {
	count, err := d.execCount(`UPDATE cluster SET heartbeat = ? WHERE uuid = ?`, i, member.UUID)
	if err != nil {
		return err
	}

	if count == 0 {
		_, err := d.execCount(`INSERT INTO cluster(name,uuid,ip_address,requested_index) values(?, ?, ?, ?)`,
			member.Name, member.UUID, member.IP, member.RequestedIndex)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DB) SaveIndex(indexes map[int]Member) error {
	for index, member := range indexes {
		_, err := d.execCount(`UPDATE cluster SET  assigned_index = ?, requested_index = ? WHERE ID = ?`,
			index, 0, member.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DB) execCount(sql string, args ...interface{}) (int64, error) {
	res, err := d.db.Exec(sql, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

type NullStringWrapper struct {
	sql.NullString
	String *string
}

func (n *NullStringWrapper) Scan(value interface{}) error {
	if err := n.NullString.Scan(value); err != nil {
		return err
	}
	n.String = &n.NullString.String
	return nil
}
