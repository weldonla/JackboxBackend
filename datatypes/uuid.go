package datatypes

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

//MYTYPE -> new datatype
type MYTYPE uuid.UUID

// StringToMYTYPE -> parse string to MYTYPE
func StringToMYTYPE(s string) (MYTYPE, error) {
	id, err := uuid.Parse(s)
	return MYTYPE(id), err
}

//String -> String Representation of Binary16
func (my MYTYPE) String() string {
	return uuid.UUID(my).String()
}

//GormDataType -> sets type to binary(16)
func (my MYTYPE) GormDataType() string {
	return "binary(16)"
}

func (my MYTYPE) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(my)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

func (my *MYTYPE) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*my = MYTYPE(s)
	return err
}

// Scan --> tells GORM how to receive from the database
func (my *MYTYPE) Scan(value interface{}) error {

	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*my = MYTYPE(parseByte)
	return err
}

// Value -> tells GORM how to save into the database
func (my MYTYPE) Value() (driver.Value, error) {
	return uuid.UUID(my).MarshalBinary()
}
