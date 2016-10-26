package chado

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v5"
)

// Stockcollection is an object representing the database table.
type Stockcollection struct {
	StockcollectionID int         `boil:"stockcollection_id" json:"stockcollection_id" toml:"stockcollection_id" yaml:"stockcollection_id"`
	TypeID            int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	ContactID         null.Int    `boil:"contact_id" json:"contact_id,omitempty" toml:"contact_id" yaml:"contact_id,omitempty"`
	Name              null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Uniquename        string      `boil:"uniquename" json:"uniquename" toml:"uniquename" yaml:"uniquename"`

	R *stockcollectionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockcollectionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockcollectionR is where relationships are stored.
type stockcollectionR struct {
	Contact              *Contact
	Type                 *Cvterm
	StockcollectionStock *StockcollectionStock
	Stockcollectionprop  *Stockcollectionprop
}

// stockcollectionL is where Load methods for each relationship are stored.
type stockcollectionL struct{}

var (
	stockcollectionColumns               = []string{"stockcollection_id", "type_id", "contact_id", "name", "uniquename"}
	stockcollectionColumnsWithoutDefault = []string{"type_id", "contact_id", "name", "uniquename"}
	stockcollectionColumnsWithDefault    = []string{"stockcollection_id"}
	stockcollectionPrimaryKeyColumns     = []string{"stockcollection_id"}
)

type (
	// StockcollectionSlice is an alias for a slice of pointers to Stockcollection.
	// This should generally be used opposed to []Stockcollection.
	StockcollectionSlice []*Stockcollection
	// StockcollectionHook is the signature for custom Stockcollection hook methods
	StockcollectionHook func(boil.Executor, *Stockcollection) error

	stockcollectionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockcollectionType                 = reflect.TypeOf(&Stockcollection{})
	stockcollectionMapping              = queries.MakeStructMapping(stockcollectionType)
	stockcollectionPrimaryKeyMapping, _ = queries.BindMapping(stockcollectionType, stockcollectionMapping, stockcollectionPrimaryKeyColumns)
	stockcollectionInsertCacheMut       sync.RWMutex
	stockcollectionInsertCache          = make(map[string]insertCache)
	stockcollectionUpdateCacheMut       sync.RWMutex
	stockcollectionUpdateCache          = make(map[string]updateCache)
	stockcollectionUpsertCacheMut       sync.RWMutex
	stockcollectionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockcollectionBeforeInsertHooks []StockcollectionHook
var stockcollectionBeforeUpdateHooks []StockcollectionHook
var stockcollectionBeforeDeleteHooks []StockcollectionHook
var stockcollectionBeforeUpsertHooks []StockcollectionHook

var stockcollectionAfterInsertHooks []StockcollectionHook
var stockcollectionAfterSelectHooks []StockcollectionHook
var stockcollectionAfterUpdateHooks []StockcollectionHook
var stockcollectionAfterDeleteHooks []StockcollectionHook
var stockcollectionAfterUpsertHooks []StockcollectionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Stockcollection) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Stockcollection) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Stockcollection) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Stockcollection) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Stockcollection) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Stockcollection) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Stockcollection) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Stockcollection) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Stockcollection) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockcollectionHook registers your hook function for all future operations.
func AddStockcollectionHook(hookPoint boil.HookPoint, stockcollectionHook StockcollectionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockcollectionBeforeInsertHooks = append(stockcollectionBeforeInsertHooks, stockcollectionHook)
	case boil.BeforeUpdateHook:
		stockcollectionBeforeUpdateHooks = append(stockcollectionBeforeUpdateHooks, stockcollectionHook)
	case boil.BeforeDeleteHook:
		stockcollectionBeforeDeleteHooks = append(stockcollectionBeforeDeleteHooks, stockcollectionHook)
	case boil.BeforeUpsertHook:
		stockcollectionBeforeUpsertHooks = append(stockcollectionBeforeUpsertHooks, stockcollectionHook)
	case boil.AfterInsertHook:
		stockcollectionAfterInsertHooks = append(stockcollectionAfterInsertHooks, stockcollectionHook)
	case boil.AfterSelectHook:
		stockcollectionAfterSelectHooks = append(stockcollectionAfterSelectHooks, stockcollectionHook)
	case boil.AfterUpdateHook:
		stockcollectionAfterUpdateHooks = append(stockcollectionAfterUpdateHooks, stockcollectionHook)
	case boil.AfterDeleteHook:
		stockcollectionAfterDeleteHooks = append(stockcollectionAfterDeleteHooks, stockcollectionHook)
	case boil.AfterUpsertHook:
		stockcollectionAfterUpsertHooks = append(stockcollectionAfterUpsertHooks, stockcollectionHook)
	}
}

// OneP returns a single stockcollection record from the query, and panics on error.
func (q stockcollectionQuery) OneP() *Stockcollection {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockcollection record from the query.
func (q stockcollectionQuery) One() (*Stockcollection, error) {
	o := &Stockcollection{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for stockcollection")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Stockcollection records from the query, and panics on error.
func (q stockcollectionQuery) AllP() StockcollectionSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Stockcollection records from the query.
func (q stockcollectionQuery) All() (StockcollectionSlice, error) {
	var o StockcollectionSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Stockcollection slice")
	}

	if len(stockcollectionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Stockcollection records in the query, and panics on error.
func (q stockcollectionQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Stockcollection records in the query.
func (q stockcollectionQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count stockcollection rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockcollectionQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockcollectionQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if stockcollection exists")
	}

	return count > 0, nil
}

// ContactG pointed to by the foreign key.
func (o *Stockcollection) ContactG(mods ...qm.QueryMod) contactQuery {
	return o.Contact(boil.GetDB(), mods...)
}

// Contact pointed to by the foreign key.
func (o *Stockcollection) Contact(exec boil.Executor, mods ...qm.QueryMod) contactQuery {
	queryMods := []qm.QueryMod{
		qm.Where("contact_id=$1", o.ContactID),
	}

	queryMods = append(queryMods, mods...)

	query := Contacts(exec, queryMods...)
	queries.SetFrom(query.Query, "\"contact\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Stockcollection) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Stockcollection) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// StockcollectionStockG pointed to by the foreign key.
func (o *Stockcollection) StockcollectionStockG(mods ...qm.QueryMod) stockcollectionStockQuery {
	return o.StockcollectionStock(boil.GetDB(), mods...)
}

// StockcollectionStock pointed to by the foreign key.
func (o *Stockcollection) StockcollectionStock(exec boil.Executor, mods ...qm.QueryMod) stockcollectionStockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stockcollection_id=$1", o.StockcollectionID),
	}

	queryMods = append(queryMods, mods...)

	query := StockcollectionStocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stockcollection_stock\"")

	return query
}

// StockcollectionpropG pointed to by the foreign key.
func (o *Stockcollection) StockcollectionpropG(mods ...qm.QueryMod) stockcollectionpropQuery {
	return o.Stockcollectionprop(boil.GetDB(), mods...)
}

// Stockcollectionprop pointed to by the foreign key.
func (o *Stockcollection) Stockcollectionprop(exec boil.Executor, mods ...qm.QueryMod) stockcollectionpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stockcollection_id=$1", o.StockcollectionID),
	}

	queryMods = append(queryMods, mods...)

	query := Stockcollectionprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stockcollectionprop\"")

	return query
}

// LoadContact allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockcollectionL) LoadContact(e boil.Executor, singular bool, maybeStockcollection interface{}) error {
	var slice []*Stockcollection
	var object *Stockcollection

	count := 1
	if singular {
		object = maybeStockcollection.(*Stockcollection)
	} else {
		slice = *maybeStockcollection.(*StockcollectionSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockcollectionR{}
		args[0] = object.ContactID
	} else {
		for i, obj := range slice {
			obj.R = &stockcollectionR{}
			args[i] = obj.ContactID
		}
	}

	query := fmt.Sprintf(
		"select * from \"contact\" where \"contact_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Contact")
	}
	defer results.Close()

	var resultSlice []*Contact
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Contact")
	}

	if len(stockcollectionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Contact = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ContactID.Int == foreign.ContactID {
				local.R.Contact = foreign
				break
			}
		}
	}

	return nil
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockcollectionL) LoadType(e boil.Executor, singular bool, maybeStockcollection interface{}) error {
	var slice []*Stockcollection
	var object *Stockcollection

	count := 1
	if singular {
		object = maybeStockcollection.(*Stockcollection)
	} else {
		slice = *maybeStockcollection.(*StockcollectionSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockcollectionR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &stockcollectionR{}
			args[i] = obj.TypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvterm")
	}
	defer results.Close()

	var resultSlice []*Cvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvterm")
	}

	if len(stockcollectionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Type = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.TypeID == foreign.CvtermID {
				local.R.Type = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockcollectionStock allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockcollectionL) LoadStockcollectionStock(e boil.Executor, singular bool, maybeStockcollection interface{}) error {
	var slice []*Stockcollection
	var object *Stockcollection

	count := 1
	if singular {
		object = maybeStockcollection.(*Stockcollection)
	} else {
		slice = *maybeStockcollection.(*StockcollectionSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockcollectionR{}
		args[0] = object.StockcollectionID
	} else {
		for i, obj := range slice {
			obj.R = &stockcollectionR{}
			args[i] = obj.StockcollectionID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stockcollection_stock\" where \"stockcollection_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockcollectionStock")
	}
	defer results.Close()

	var resultSlice []*StockcollectionStock
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockcollectionStock")
	}

	if len(stockcollectionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockcollectionStock = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockcollectionID == foreign.StockcollectionID {
				local.R.StockcollectionStock = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockcollectionprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockcollectionL) LoadStockcollectionprop(e boil.Executor, singular bool, maybeStockcollection interface{}) error {
	var slice []*Stockcollection
	var object *Stockcollection

	count := 1
	if singular {
		object = maybeStockcollection.(*Stockcollection)
	} else {
		slice = *maybeStockcollection.(*StockcollectionSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockcollectionR{}
		args[0] = object.StockcollectionID
	} else {
		for i, obj := range slice {
			obj.R = &stockcollectionR{}
			args[i] = obj.StockcollectionID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stockcollectionprop\" where \"stockcollection_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Stockcollectionprop")
	}
	defer results.Close()

	var resultSlice []*Stockcollectionprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Stockcollectionprop")
	}

	if len(stockcollectionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Stockcollectionprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockcollectionID == foreign.StockcollectionID {
				local.R.Stockcollectionprop = foreign
				break
			}
		}
	}

	return nil
}

// SetContact of the stockcollection to the related item.
// Sets o.R.Contact to related.
// Adds o to related.R.Stockcollections.
func (o *Stockcollection) SetContact(exec boil.Executor, insert bool, related *Contact) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stockcollection\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"contact_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockcollectionPrimaryKeyColumns),
	)
	values := []interface{}{related.ContactID, o.StockcollectionID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ContactID.Int = related.ContactID
	o.ContactID.Valid = true

	if o.R == nil {
		o.R = &stockcollectionR{
			Contact: related,
		}
	} else {
		o.R.Contact = related
	}

	if related.R == nil {
		related.R = &contactR{
			Stockcollections: StockcollectionSlice{o},
		}
	} else {
		related.R.Stockcollections = append(related.R.Stockcollections, o)
	}

	return nil
}

// RemoveContact relationship.
// Sets o.R.Contact to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Stockcollection) RemoveContact(exec boil.Executor, related *Contact) error {
	var err error

	o.ContactID.Valid = false
	if err = o.Update(exec, "contact_id"); err != nil {
		o.ContactID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Contact = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Stockcollections {
		if o.ContactID.Int != ri.ContactID.Int {
			continue
		}

		ln := len(related.R.Stockcollections)
		if ln > 1 && i < ln-1 {
			related.R.Stockcollections[i] = related.R.Stockcollections[ln-1]
		}
		related.R.Stockcollections = related.R.Stockcollections[:ln-1]
		break
	}
	return nil
}

// SetType of the stockcollection to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeStockcollection.
func (o *Stockcollection) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stockcollection\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockcollectionPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.StockcollectionID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &stockcollectionR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeStockcollection: o,
		}
	} else {
		related.R.TypeStockcollection = o
	}

	return nil
}

// SetStockcollectionStock of the stockcollection to the related item.
// Sets o.R.StockcollectionStock to related.
// Adds o to related.R.Stockcollection.
func (o *Stockcollection) SetStockcollectionStock(exec boil.Executor, insert bool, related *StockcollectionStock) error {
	var err error

	if insert {
		related.StockcollectionID = o.StockcollectionID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stockcollection_stock\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"stockcollection_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockcollectionStockPrimaryKeyColumns),
		)
		values := []interface{}{o.StockcollectionID, related.StockcollectionStockID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.StockcollectionID = o.StockcollectionID

	}

	if o.R == nil {
		o.R = &stockcollectionR{
			StockcollectionStock: related,
		}
	} else {
		o.R.StockcollectionStock = related
	}

	if related.R == nil {
		related.R = &stockcollectionStockR{
			Stockcollection: o,
		}
	} else {
		related.R.Stockcollection = o
	}
	return nil
}

// SetStockcollectionprop of the stockcollection to the related item.
// Sets o.R.Stockcollectionprop to related.
// Adds o to related.R.Stockcollection.
func (o *Stockcollection) SetStockcollectionprop(exec boil.Executor, insert bool, related *Stockcollectionprop) error {
	var err error

	if insert {
		related.StockcollectionID = o.StockcollectionID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stockcollectionprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"stockcollection_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockcollectionpropPrimaryKeyColumns),
		)
		values := []interface{}{o.StockcollectionID, related.StockcollectionpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.StockcollectionID = o.StockcollectionID

	}

	if o.R == nil {
		o.R = &stockcollectionR{
			Stockcollectionprop: related,
		}
	} else {
		o.R.Stockcollectionprop = related
	}

	if related.R == nil {
		related.R = &stockcollectionpropR{
			Stockcollection: o,
		}
	} else {
		related.R.Stockcollection = o
	}
	return nil
}

// StockcollectionsG retrieves all records.
func StockcollectionsG(mods ...qm.QueryMod) stockcollectionQuery {
	return Stockcollections(boil.GetDB(), mods...)
}

// Stockcollections retrieves all the records using an executor.
func Stockcollections(exec boil.Executor, mods ...qm.QueryMod) stockcollectionQuery {
	mods = append(mods, qm.From("\"stockcollection\""))
	return stockcollectionQuery{NewQuery(exec, mods...)}
}

// FindStockcollectionG retrieves a single record by ID.
func FindStockcollectionG(stockcollectionID int, selectCols ...string) (*Stockcollection, error) {
	return FindStockcollection(boil.GetDB(), stockcollectionID, selectCols...)
}

// FindStockcollectionGP retrieves a single record by ID, and panics on error.
func FindStockcollectionGP(stockcollectionID int, selectCols ...string) *Stockcollection {
	retobj, err := FindStockcollection(boil.GetDB(), stockcollectionID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockcollection retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockcollection(exec boil.Executor, stockcollectionID int, selectCols ...string) (*Stockcollection, error) {
	stockcollectionObj := &Stockcollection{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stockcollection\" where \"stockcollection_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockcollectionID)

	err := q.Bind(stockcollectionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from stockcollection")
	}

	return stockcollectionObj, nil
}

// FindStockcollectionP retrieves a single record by ID with an executor, and panics on error.
func FindStockcollectionP(exec boil.Executor, stockcollectionID int, selectCols ...string) *Stockcollection {
	retobj, err := FindStockcollection(exec, stockcollectionID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Stockcollection) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Stockcollection) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Stockcollection) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Stockcollection) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no stockcollection provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockcollectionColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockcollectionInsertCacheMut.RLock()
	cache, cached := stockcollectionInsertCache[key]
	stockcollectionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockcollectionColumns,
			stockcollectionColumnsWithDefault,
			stockcollectionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockcollectionType, stockcollectionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockcollectionType, stockcollectionMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stockcollection\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "chado: unable to insert into stockcollection")
	}

	if !cached {
		stockcollectionInsertCacheMut.Lock()
		stockcollectionInsertCache[key] = cache
		stockcollectionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Stockcollection record. See Update for
// whitelist behavior description.
func (o *Stockcollection) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Stockcollection record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Stockcollection) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Stockcollection, and panics on error.
// See Update for whitelist behavior description.
func (o *Stockcollection) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Stockcollection.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Stockcollection) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockcollectionUpdateCacheMut.RLock()
	cache, cached := stockcollectionUpdateCache[key]
	stockcollectionUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockcollectionColumns, stockcollectionPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update stockcollection, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stockcollection\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockcollectionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockcollectionType, stockcollectionMapping, append(wl, stockcollectionPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update stockcollection row")
	}

	if !cached {
		stockcollectionUpdateCacheMut.Lock()
		stockcollectionUpdateCache[key] = cache
		stockcollectionUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockcollectionQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockcollectionQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for stockcollection")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockcollectionSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockcollectionSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockcollectionSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockcollectionSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("chado: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockcollectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stockcollection\" SET %s WHERE (\"stockcollection_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockcollectionPrimaryKeyColumns), len(colNames)+1, len(stockcollectionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in stockcollection slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Stockcollection) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Stockcollection) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Stockcollection) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Stockcollection) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no stockcollection provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockcollectionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	stockcollectionUpsertCacheMut.RLock()
	cache, cached := stockcollectionUpsertCache[key]
	stockcollectionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockcollectionColumns,
			stockcollectionColumnsWithDefault,
			stockcollectionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockcollectionColumns,
			stockcollectionPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert stockcollection, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockcollectionPrimaryKeyColumns))
			copy(conflict, stockcollectionPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stockcollection\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockcollectionType, stockcollectionMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockcollectionType, stockcollectionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "chado: unable to upsert for stockcollection")
	}

	if !cached {
		stockcollectionUpsertCacheMut.Lock()
		stockcollectionUpsertCache[key] = cache
		stockcollectionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Stockcollection record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Stockcollection) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Stockcollection record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Stockcollection) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Stockcollection provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Stockcollection record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Stockcollection) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Stockcollection record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Stockcollection) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Stockcollection provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockcollectionPrimaryKeyMapping)
	sql := "DELETE FROM \"stockcollection\" WHERE \"stockcollection_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from stockcollection")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockcollectionQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockcollectionQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no stockcollectionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from stockcollection")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockcollectionSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockcollectionSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Stockcollection slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockcollectionSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockcollectionSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Stockcollection slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockcollectionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockcollectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stockcollection\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockcollectionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockcollectionPrimaryKeyColumns), 1, len(stockcollectionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from stockcollection slice")
	}

	if len(stockcollectionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Stockcollection) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Stockcollection) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Stockcollection) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Stockcollection provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Stockcollection) Reload(exec boil.Executor) error {
	ret, err := FindStockcollection(exec, o.StockcollectionID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockcollectionSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockcollectionSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockcollectionSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty StockcollectionSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockcollectionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockcollections := StockcollectionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockcollectionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stockcollection\".* FROM \"stockcollection\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockcollectionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockcollectionPrimaryKeyColumns), 1, len(stockcollectionPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockcollections)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in StockcollectionSlice")
	}

	*o = stockcollections

	return nil
}

// StockcollectionExists checks if the Stockcollection row exists.
func StockcollectionExists(exec boil.Executor, stockcollectionID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stockcollection\" where \"stockcollection_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockcollectionID)
	}

	row := exec.QueryRow(sql, stockcollectionID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if stockcollection exists")
	}

	return exists, nil
}

// StockcollectionExistsG checks if the Stockcollection row exists.
func StockcollectionExistsG(stockcollectionID int) (bool, error) {
	return StockcollectionExists(boil.GetDB(), stockcollectionID)
}

// StockcollectionExistsGP checks if the Stockcollection row exists. Panics on error.
func StockcollectionExistsGP(stockcollectionID int) bool {
	e, err := StockcollectionExists(boil.GetDB(), stockcollectionID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockcollectionExistsP checks if the Stockcollection row exists. Panics on error.
func StockcollectionExistsP(exec boil.Executor, stockcollectionID int) bool {
	e, err := StockcollectionExists(exec, stockcollectionID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
