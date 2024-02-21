// Code generated by ent, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/tier"
)

// Tier is the model entity for the Tier schema.
type Tier struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Vcpu holds the value of the "vcpu" field.
	Vcpu int64 `json:"vcpu,omitempty"`
	// RAMMB holds the value of the "ram_mb" field.
	RAMMB int64 `json:"ram_mb,omitempty"`
	// DiskMB holds the value of the "disk_mb" field.
	DiskMB int64 `json:"disk_mb,omitempty"`
	// The number of instances the team can run concurrently
	ConcurrentInstances int64 `json:"concurrent_instances,omitempty"`
	// MaxLengthHours holds the value of the "max_length_hours" field.
	MaxLengthHours int64 `json:"max_length_hours,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TierQuery when eager-loading is set.
	Edges        TierEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TierEdges holds the relations/edges for other nodes in the graph.
type TierEdges struct {
	// Teams holds the value of the teams edge.
	Teams []*Team `json:"teams,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TeamsOrErr returns the Teams value or an error if the edge
// was not loaded in eager-loading.
func (e TierEdges) TeamsOrErr() ([]*Team, error) {
	if e.loadedTypes[0] {
		return e.Teams, nil
	}
	return nil, &NotLoadedError{edge: "teams"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tier) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tier.FieldVcpu, tier.FieldRAMMB, tier.FieldDiskMB, tier.FieldConcurrentInstances, tier.FieldMaxLengthHours:
			values[i] = new(sql.NullInt64)
		case tier.FieldID, tier.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tier fields.
func (t *Tier) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tier.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				t.ID = value.String
			}
		case tier.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case tier.FieldVcpu:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vcpu", values[i])
			} else if value.Valid {
				t.Vcpu = value.Int64
			}
		case tier.FieldRAMMB:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ram_mb", values[i])
			} else if value.Valid {
				t.RAMMB = value.Int64
			}
		case tier.FieldDiskMB:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field disk_mb", values[i])
			} else if value.Valid {
				t.DiskMB = value.Int64
			}
		case tier.FieldConcurrentInstances:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field concurrent_instances", values[i])
			} else if value.Valid {
				t.ConcurrentInstances = value.Int64
			}
		case tier.FieldMaxLengthHours:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_length_hours", values[i])
			} else if value.Valid {
				t.MaxLengthHours = value.Int64
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tier.
// This includes values selected through modifiers, order, etc.
func (t *Tier) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryTeams queries the "teams" edge of the Tier entity.
func (t *Tier) QueryTeams() *TeamQuery {
	return NewTierClient(t.config).QueryTeams(t)
}

// Update returns a builder for updating this Tier.
// Note that you need to call Tier.Unwrap() before calling this method if this Tier
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tier) Update() *TierUpdateOne {
	return NewTierClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tier entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tier) Unwrap() *Tier {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("models: Tier is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tier) String() string {
	var builder strings.Builder
	builder.WriteString("Tier(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("vcpu=")
	builder.WriteString(fmt.Sprintf("%v", t.Vcpu))
	builder.WriteString(", ")
	builder.WriteString("ram_mb=")
	builder.WriteString(fmt.Sprintf("%v", t.RAMMB))
	builder.WriteString(", ")
	builder.WriteString("disk_mb=")
	builder.WriteString(fmt.Sprintf("%v", t.DiskMB))
	builder.WriteString(", ")
	builder.WriteString("concurrent_instances=")
	builder.WriteString(fmt.Sprintf("%v", t.ConcurrentInstances))
	builder.WriteString(", ")
	builder.WriteString("max_length_hours=")
	builder.WriteString(fmt.Sprintf("%v", t.MaxLengthHours))
	builder.WriteByte(')')
	return builder.String()
}

// Tiers is a parsable slice of Tier.
type Tiers []*Tier
