package migrations

import (
	"database/sql"

	"code.cloudfoundry.org/bbs/encryption"
	"code.cloudfoundry.org/bbs/format"
	"code.cloudfoundry.org/bbs/migration"
	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager/v3"
)

func init() {
	appendMigration(NewAddRoutableToActualLrps())
}

type AddRoutableToActualLrps struct {
	serializer format.Serializer
	clock      clock.Clock
	dbFlavor   string
}

func NewAddRoutableToActualLrps() migration.Migration {
	return new(AddRoutableToActualLrps)
}

func (e *AddRoutableToActualLrps) String() string {
	return migrationString(e)
}

func (e *AddRoutableToActualLrps) Version() int64 {
	return 1686692176
}

func (e *AddRoutableToActualLrps) SetCryptor(cryptor encryption.Cryptor) {
	e.serializer = format.NewSerializer(cryptor)
}

func (e *AddRoutableToActualLrps) SetClock(c clock.Clock)    { e.clock = c }
func (e *AddRoutableToActualLrps) SetDBFlavor(flavor string) { e.dbFlavor = flavor }

func (e *AddRoutableToActualLrps) Up(tx *sql.Tx, logger lager.Logger) error {
	var alterTablesSQL = []string{
		alterActualLRPAddRoutableSQL,
		alterActualLRPSetRoutableForRunningSQL,
	}
	for _, query := range alterTablesSQL {
		logger.Info("altering the table", lager.Data{"query": query})
		_, err := tx.Exec(query)
		if err != nil {
			logger.Error("failed-altering-tables", err)
			return err
		}
		logger.Info("altered the table", lager.Data{"query": query})
	}

	return nil
}

const alterActualLRPAddRoutableSQL = `ALTER TABLE actual_lrps
ADD COLUMN routable BOOL DEFAULT false;`

const alterActualLRPSetRoutableForRunningSQL = `UPDATE actual_lrps
SET routable = true WHERE state = 'RUNNING';`
