// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kohTkd/experimental_learning/ent/accounts"
)

// Accounts is the model entity for the Accounts schema.
type Accounts struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt     time.Time `json:"update_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Accounts) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case accounts.FieldID:
			values[i] = new(sql.NullInt64)
		case accounts.FieldEmail, accounts.FieldPassword:
			values[i] = new(sql.NullString)
		case accounts.FieldCreatedAt, accounts.FieldUpdateAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Accounts fields.
func (a *Accounts) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case accounts.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case accounts.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				a.Email = value.String
			}
		case accounts.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				a.Password = value.String
			}
		case accounts.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case accounts.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				a.UpdateAt = value.Time
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Accounts.
// This includes values selected through modifiers, order, etc.
func (a *Accounts) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Accounts.
// Note that you need to call Accounts.Unwrap() before calling this method if this Accounts
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Accounts) Update() *AccountsUpdateOne {
	return NewAccountsClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Accounts entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Accounts) Unwrap() *Accounts {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Accounts is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Accounts) String() string {
	var builder strings.Builder
	builder.WriteString("Accounts(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("email=")
	builder.WriteString(a.Email)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(a.Password)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(a.UpdateAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AccountsSlice is a parsable slice of Accounts.
type AccountsSlice []*Accounts
