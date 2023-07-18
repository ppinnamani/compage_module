package daos

import (
	"database/sql"
	"errors"
	"github.com/ppinnamani/compage_module/compage_module_bheem/pkg/rest/server/daos/clients/sqls"
	"github.com/ppinnamani/compage_module/compage_module_bheem/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type CompageDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateCompages(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS compages(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Password TEXT NOT NULL,
		User_name TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewCompageDao() (*CompageDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateCompages(sqlClient)
	if err != nil {
		return nil, err
	}
	return &CompageDao{
		sqlClient,
	}, nil
}

func (compageDao *CompageDao) CreateCompage(m *models.Compage) (*models.Compage, error) {
	insertQuery := "INSERT INTO compages(Password, User_name)values(?, ?)"
	res, err := compageDao.sqlClient.DB.Exec(insertQuery, m.Password, m.User_name)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("compage created")
	return m, nil
}

func (compageDao *CompageDao) UpdateCompage(id int64, m *models.Compage) (*models.Compage, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	compage, err := compageDao.GetCompage(id)
	if err != nil {
		return nil, err
	}
	if compage == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE compages SET Password = ?, User_name = ? WHERE Id = ?"
	res, err := compageDao.sqlClient.DB.Exec(updateQuery, m.Password, m.User_name, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("compage updated")
	return m, nil
}

func (compageDao *CompageDao) DeleteCompage(id int64) error {
	deleteQuery := "DELETE FROM compages WHERE Id = ?"
	res, err := compageDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("compage deleted")
	return nil
}

func (compageDao *CompageDao) ListCompages() ([]*models.Compage, error) {
	selectQuery := "SELECT * FROM compages"
	rows, err := compageDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var compages []*models.Compage
	for rows.Next() {
		m := models.Compage{}
		if err = rows.Scan(&m.Id, &m.Password, &m.User_name); err != nil {
			return nil, err
		}
		compages = append(compages, &m)
	}
	if compages == nil {
		compages = []*models.Compage{}
	}

	log.Debugf("compage listed")
	return compages, nil
}

func (compageDao *CompageDao) GetCompage(id int64) (*models.Compage, error) {
	selectQuery := "SELECT * FROM compages WHERE Id = ?"
	row := compageDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Compage{}
	if err := row.Scan(&m.Id, &m.Password, &m.User_name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("compage retrieved")
	return &m, nil
}
