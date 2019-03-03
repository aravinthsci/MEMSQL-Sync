package constants

const (
	SELECT_TABLES_CMD_FORMAT = "memsql -u%s -p%s -B -N -e 'SELECT * FROM %s.%s'"
	SHOW_TABLES_CMD_FORMAT   = "memsql %s -u%s -p%s -B -N -e 'show tables'"

	CLEAN_TABLES_CMD_FORMAT                    = "memsql -u%s -p%s -B -N -e 'DELETE FROM %s.%s'"
	CLEAN_TABLES_CMD_FORMAT_WITHOUT_PASSPHRASE = "memsql -u%s -B -N -e 'DELETE FROM %s.%s'"

	DELETE_TABLE_QUERY_FORMAT = "DELETE FROM %s.%s"
	LOAD_INFILE_QUERY_FORMAT  = "LOAD DATA LOCAL INFILE '%s' INTO TABLE %s.%s"

	TMP_DIR_PATH = "/tmp/db_sync"

	//gopli sync -from production -to staging -c /root/Projects/src/gopli.toml
)
