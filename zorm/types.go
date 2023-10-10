package zorm

var driverRegistry = make(map[string]SchemaDriver)

func Register(driver SchemaDriver) { driverRegistry[driver.Name()] = driver }

func Get(name string) SchemaDriver { return driverRegistry[name] }

type (
	SchemaDriver interface {
		Name() string

		Parse(dsn, schema, table string, types map[string]string) (tables []Table, err error)
	}

	Table struct {
		Name    string
		Table   string
		Schema  string
		Comment string
		Primary string
		Columns []Column
		Ext     interface{}
	}

	Column struct {
		Name          string
		Type          string
		Column        string
		Comment       string
		Nullable      bool
		MaximumLength int64
		Ext           interface{}
	}
)

func DefaultTypes() map[string]string {
	return map[string]string{
		// int
		"int":     "int",
		"tinyint": "int32",
		"bigint":  "int64",
		// float
		"double":  "float64",
		"decimal": "float64",
		"float":   "float64",
		// string
		"mediumtext": "string",
		"varchar":    "string",
		"char":       "string",
		"longtext":   "string",
		"text":       "string",
		"enum":       "string",
		// bytes
		"blob":      "[]byte",
		"binary":    "[]byte",
		"varbinary": "[]byte",
		"json":      "json.RawMessage",
		// set
		"set": "[]string",
		// time
		"timestamp": "time.Time",
		"datetime":  "time.Time",

		// nullable int
		"*int":     "sql.NullInt32",
		"*tinyint": "sql.NullInt32",
		"*bigint":  "sql.NullInt64",
		// nullable string
		"*mediumtext": "sql.NullString",
		"*varchar":    "sql.NullString",
		"*char":       "sql.NullString",
		"*longtext":   "sql.NullString",
		"*text":       "sql.NullString",
		"*enum":       "sql.NullString",
		// nullable time
		"*timestamp": "sql.NullTime",
		"*datetime":  "sql.NullTime",
	}
}
