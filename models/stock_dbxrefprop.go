package models

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

// StockDbxrefprop is an object representing the database table.
type StockDbxrefprop struct {
	StockDbxrefpropID int         `boil:"stock_dbxrefprop_id" json:"stock_dbxrefprop_id" toml:"stock_dbxrefprop_id" yaml:"stock_dbxrefprop_id"`
	StockDbxrefID     int         `boil:"stock_dbxref_id" json:"stock_dbxref_id" toml:"stock_dbxref_id" yaml:"stock_dbxref_id"`
	TypeID            int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value             null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank              int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *stockDbxrefpropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockDbxrefpropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockDbxrefpropR is where relationships are stored.
type stockDbxrefpropR struct {
	Type        *Cvterm
	StockDbxref *StockDbxref
}

// stockDbxrefpropL is where Load methods for each relationship are stored.
type stockDbxrefpropL struct{}

var (
	stockDbxrefpropColumns               = []string{"stock_dbxrefprop_id", "stock_dbxref_id", "type_id", "value", "rank"}
	stockDbxrefpropColumnsWithoutDefault = []string{"stock_dbxref_id", "type_id", "value"}
	stockDbxrefpropColumnsWithDefault    = []string{"stock_dbxrefprop_id", "rank"}
	stockDbxrefpropPrimaryKeyColumns     = []string{"stock_dbxrefprop_id"}
)

type (
	// StockDbxrefpropSlice is an alias for a slice of pointers to StockDbxrefprop.
	// This should generally be used opposed to []StockDbxrefprop.
	StockDbxrefpropSlice []*StockDbxrefprop
	// StockDbxrefpropHook is the signature for custom StockDbxrefprop hook methods
	StockDbxrefpropHook func(boil.Executor, *StockDbxrefprop) error

	stockDbxrefpropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockDbxrefpropType                 = reflect.TypeOf(&StockDbxrefprop{})
	stockDbxrefpropMapping              = queries.MakeStructMapping(stockDbxrefpropType)
	stockDbxrefpropPrimaryKeyMapping, _ = queries.BindMapping(stockDbxrefpropType, stockDbxrefpropMapping, stockDbxrefpropPrimaryKeyColumns)
	stockDbxrefpropInsertCacheMut       sync.RWMutex
	stockDbxrefpropInsertCache          = make(map[string]insertCache)
	stockDbxrefpropUpdateCacheMut       sync.RWMutex
	stockDbxrefpropUpdateCache          = make(map[string]updateCache)
	stockDbxrefpropUpsertCacheMut       sync.RWMutex
	stockDbxrefpropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockDbxrefpropBeforeInsertHooks []StockDbxrefpropHook
var stockDbxrefpropBeforeUpdateHooks []StockDbxrefpropHook
var stockDbxrefpropBeforeDeleteHooks []StockDbxrefpropHook
var stockDbxrefpropBeforeUpsertHooks []StockDbxrefpropHook

var stockDbxrefpropAfterInsertHooks []StockDbxrefpropHook
var stockDbxrefpropAfterSelectHooks []StockDbxrefpropHook
var stockDbxrefpropAfterUpdateHooks []StockDbxrefpropHook
var stockDbxrefpropAfterDeleteHooks []StockDbxrefpropHook
var stockDbxrefpropAfterUpsertHooks []StockDbxrefpropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockDbxrefprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefpropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockDbxrefprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefpropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockDbxrefprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefpropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockDbxrefprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefpropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockDbxrefprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefpropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockDbxrefprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefpropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockDbxrefprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefpropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockDbxrefprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefpropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockDbxrefprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefpropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockDbxrefpropHook registers your hook function for all future operations.
func AddStockDbxrefpropHook(hookPoint boil.HookPoint, stockDbxrefpropHook StockDbxrefpropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockDbxrefpropBeforeInsertHooks = append(stockDbxrefpropBeforeInsertHooks, stockDbxrefpropHook)
	case boil.BeforeUpdateHook:
		stockDbxrefpropBeforeUpdateHooks = append(stockDbxrefpropBeforeUpdateHooks, stockDbxrefpropHook)
	case boil.BeforeDeleteHook:
		stockDbxrefpropBeforeDeleteHooks = append(stockDbxrefpropBeforeDeleteHooks, stockDbxrefpropHook)
	case boil.BeforeUpsertHook:
		stockDbxrefpropBeforeUpsertHooks = append(stockDbxrefpropBeforeUpsertHooks, stockDbxrefpropHook)
	case boil.AfterInsertHook:
		stockDbxrefpropAfterInsertHooks = append(stockDbxrefpropAfterInsertHooks, stockDbxrefpropHook)
	case boil.AfterSelectHook:
		stockDbxrefpropAfterSelectHooks = append(stockDbxrefpropAfterSelectHooks, stockDbxrefpropHook)
	case boil.AfterUpdateHook:
		stockDbxrefpropAfterUpdateHooks = append(stockDbxrefpropAfterUpdateHooks, stockDbxrefpropHook)
	case boil.AfterDeleteHook:
		stockDbxrefpropAfterDeleteHooks = append(stockDbxrefpropAfterDeleteHooks, stockDbxrefpropHook)
	case boil.AfterUpsertHook:
		stockDbxrefpropAfterUpsertHooks = append(stockDbxrefpropAfterUpsertHooks, stockDbxrefpropHook)
	}
}

// OneP returns a single stockDbxrefprop record from the query, and panics on error.
func (q stockDbxrefpropQuery) OneP() *StockDbxrefprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockDbxrefprop record from the query.
func (q stockDbxrefpropQuery) One() (*StockDbxrefprop, error) {
	o := &StockDbxrefprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stock_dbxrefprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all StockDbxrefprop records from the query, and panics on error.
func (q stockDbxrefpropQuery) AllP() StockDbxrefpropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all StockDbxrefprop records from the query.
func (q stockDbxrefpropQuery) All() (StockDbxrefpropSlice, error) {
	var o StockDbxrefpropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StockDbxrefprop slice")
	}

	if len(stockDbxrefpropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all StockDbxrefprop records in the query, and panics on error.
func (q stockDbxrefpropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all StockDbxrefprop records in the query.
func (q stockDbxrefpropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stock_dbxrefprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockDbxrefpropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockDbxrefpropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stock_dbxrefprop exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *StockDbxrefprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *StockDbxrefprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// StockDbxrefG pointed to by the foreign key.
func (o *StockDbxrefprop) StockDbxrefG(mods ...qm.QueryMod) stockDbxrefQuery {
	return o.StockDbxref(boil.GetDB(), mods...)
}

// StockDbxref pointed to by the foreign key.
func (o *StockDbxrefprop) StockDbxref(exec boil.Executor, mods ...qm.QueryMod) stockDbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_dbxref_id=$1", o.StockDbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := StockDbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_dbxref\"")

	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockDbxrefpropL) LoadType(e boil.Executor, singular bool, maybeStockDbxrefprop interface{}) error {
	var slice []*StockDbxrefprop
	var object *StockDbxrefprop

	count := 1
	if singular {
		object = maybeStockDbxrefprop.(*StockDbxrefprop)
	} else {
		slice = *maybeStockDbxrefprop.(*StockDbxrefpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockDbxrefpropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &stockDbxrefpropR{}
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

	if len(stockDbxrefpropAfterSelectHooks) != 0 {
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

// LoadStockDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockDbxrefpropL) LoadStockDbxref(e boil.Executor, singular bool, maybeStockDbxrefprop interface{}) error {
	var slice []*StockDbxrefprop
	var object *StockDbxrefprop

	count := 1
	if singular {
		object = maybeStockDbxrefprop.(*StockDbxrefprop)
	} else {
		slice = *maybeStockDbxrefprop.(*StockDbxrefpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockDbxrefpropR{}
		args[0] = object.StockDbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &stockDbxrefpropR{}
			args[i] = obj.StockDbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_dbxref\" where \"stock_dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockDbxref")
	}
	defer results.Close()

	var resultSlice []*StockDbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockDbxref")
	}

	if len(stockDbxrefpropAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockDbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockDbxrefID == foreign.StockDbxrefID {
				local.R.StockDbxref = foreign
				break
			}
		}
	}

	return nil
}

// SetType of the stock_dbxrefprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeStockDbxrefprop.
func (o *StockDbxrefprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_dbxrefprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockDbxrefpropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.StockDbxrefpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &stockDbxrefpropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeStockDbxrefprop: o,
		}
	} else {
		related.R.TypeStockDbxrefprop = o
	}

	return nil
}

// SetStockDbxref of the stock_dbxrefprop to the related item.
// Sets o.R.StockDbxref to related.
// Adds o to related.R.StockDbxrefprop.
func (o *StockDbxrefprop) SetStockDbxref(exec boil.Executor, insert bool, related *StockDbxref) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_dbxrefprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stock_dbxref_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockDbxrefpropPrimaryKeyColumns),
	)
	values := []interface{}{related.StockDbxrefID, o.StockDbxrefpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StockDbxrefID = related.StockDbxrefID

	if o.R == nil {
		o.R = &stockDbxrefpropR{
			StockDbxref: related,
		}
	} else {
		o.R.StockDbxref = related
	}

	if related.R == nil {
		related.R = &stockDbxrefR{
			StockDbxrefprop: o,
		}
	} else {
		related.R.StockDbxrefprop = o
	}

	return nil
}

// StockDbxrefpropsG retrieves all records.
func StockDbxrefpropsG(mods ...qm.QueryMod) stockDbxrefpropQuery {
	return StockDbxrefprops(boil.GetDB(), mods...)
}

// StockDbxrefprops retrieves all the records using an executor.
func StockDbxrefprops(exec boil.Executor, mods ...qm.QueryMod) stockDbxrefpropQuery {
	mods = append(mods, qm.From("\"stock_dbxrefprop\""))
	return stockDbxrefpropQuery{NewQuery(exec, mods...)}
}

// FindStockDbxrefpropG retrieves a single record by ID.
func FindStockDbxrefpropG(stockDbxrefpropID int, selectCols ...string) (*StockDbxrefprop, error) {
	return FindStockDbxrefprop(boil.GetDB(), stockDbxrefpropID, selectCols...)
}

// FindStockDbxrefpropGP retrieves a single record by ID, and panics on error.
func FindStockDbxrefpropGP(stockDbxrefpropID int, selectCols ...string) *StockDbxrefprop {
	retobj, err := FindStockDbxrefprop(boil.GetDB(), stockDbxrefpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockDbxrefprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockDbxrefprop(exec boil.Executor, stockDbxrefpropID int, selectCols ...string) (*StockDbxrefprop, error) {
	stockDbxrefpropObj := &StockDbxrefprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stock_dbxrefprop\" where \"stock_dbxrefprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockDbxrefpropID)

	err := q.Bind(stockDbxrefpropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stock_dbxrefprop")
	}

	return stockDbxrefpropObj, nil
}

// FindStockDbxrefpropP retrieves a single record by ID with an executor, and panics on error.
func FindStockDbxrefpropP(exec boil.Executor, stockDbxrefpropID int, selectCols ...string) *StockDbxrefprop {
	retobj, err := FindStockDbxrefprop(exec, stockDbxrefpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockDbxrefprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *StockDbxrefprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *StockDbxrefprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *StockDbxrefprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_dbxrefprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockDbxrefpropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockDbxrefpropInsertCacheMut.RLock()
	cache, cached := stockDbxrefpropInsertCache[key]
	stockDbxrefpropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockDbxrefpropColumns,
			stockDbxrefpropColumnsWithDefault,
			stockDbxrefpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockDbxrefpropType, stockDbxrefpropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockDbxrefpropType, stockDbxrefpropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stock_dbxrefprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into stock_dbxrefprop")
	}

	if !cached {
		stockDbxrefpropInsertCacheMut.Lock()
		stockDbxrefpropInsertCache[key] = cache
		stockDbxrefpropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single StockDbxrefprop record. See Update for
// whitelist behavior description.
func (o *StockDbxrefprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single StockDbxrefprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *StockDbxrefprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the StockDbxrefprop, and panics on error.
// See Update for whitelist behavior description.
func (o *StockDbxrefprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the StockDbxrefprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *StockDbxrefprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockDbxrefpropUpdateCacheMut.RLock()
	cache, cached := stockDbxrefpropUpdateCache[key]
	stockDbxrefpropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockDbxrefpropColumns, stockDbxrefpropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update stock_dbxrefprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stock_dbxrefprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockDbxrefpropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockDbxrefpropType, stockDbxrefpropMapping, append(wl, stockDbxrefpropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update stock_dbxrefprop row")
	}

	if !cached {
		stockDbxrefpropUpdateCacheMut.Lock()
		stockDbxrefpropUpdateCache[key] = cache
		stockDbxrefpropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockDbxrefpropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockDbxrefpropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for stock_dbxrefprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockDbxrefpropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockDbxrefpropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockDbxrefpropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockDbxrefpropSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockDbxrefpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stock_dbxrefprop\" SET %s WHERE (\"stock_dbxrefprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockDbxrefpropPrimaryKeyColumns), len(colNames)+1, len(stockDbxrefpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in stockDbxrefprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockDbxrefprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *StockDbxrefprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *StockDbxrefprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *StockDbxrefprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_dbxrefprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockDbxrefpropColumnsWithDefault, o)

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

	stockDbxrefpropUpsertCacheMut.RLock()
	cache, cached := stockDbxrefpropUpsertCache[key]
	stockDbxrefpropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockDbxrefpropColumns,
			stockDbxrefpropColumnsWithDefault,
			stockDbxrefpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockDbxrefpropColumns,
			stockDbxrefpropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert stock_dbxrefprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockDbxrefpropPrimaryKeyColumns))
			copy(conflict, stockDbxrefpropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stock_dbxrefprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockDbxrefpropType, stockDbxrefpropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockDbxrefpropType, stockDbxrefpropMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for stock_dbxrefprop")
	}

	if !cached {
		stockDbxrefpropUpsertCacheMut.Lock()
		stockDbxrefpropUpsertCache[key] = cache
		stockDbxrefpropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single StockDbxrefprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockDbxrefprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single StockDbxrefprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockDbxrefprop) DeleteG() error {
	if o == nil {
		return errors.New("models: no StockDbxrefprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single StockDbxrefprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockDbxrefprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single StockDbxrefprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockDbxrefprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockDbxrefprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockDbxrefpropPrimaryKeyMapping)
	sql := "DELETE FROM \"stock_dbxrefprop\" WHERE \"stock_dbxrefprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from stock_dbxrefprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockDbxrefpropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockDbxrefpropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no stockDbxrefpropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stock_dbxrefprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockDbxrefpropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockDbxrefpropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no StockDbxrefprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockDbxrefpropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockDbxrefpropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockDbxrefprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockDbxrefpropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockDbxrefpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stock_dbxrefprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockDbxrefpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockDbxrefpropPrimaryKeyColumns), 1, len(stockDbxrefpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockDbxrefprop slice")
	}

	if len(stockDbxrefpropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *StockDbxrefprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *StockDbxrefprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockDbxrefprop) ReloadG() error {
	if o == nil {
		return errors.New("models: no StockDbxrefprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockDbxrefprop) Reload(exec boil.Executor) error {
	ret, err := FindStockDbxrefprop(exec, o.StockDbxrefpropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockDbxrefpropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockDbxrefpropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockDbxrefpropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty StockDbxrefpropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockDbxrefpropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockDbxrefprops := StockDbxrefpropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockDbxrefpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stock_dbxrefprop\".* FROM \"stock_dbxrefprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockDbxrefpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockDbxrefpropPrimaryKeyColumns), 1, len(stockDbxrefpropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockDbxrefprops)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockDbxrefpropSlice")
	}

	*o = stockDbxrefprops

	return nil
}

// StockDbxrefpropExists checks if the StockDbxrefprop row exists.
func StockDbxrefpropExists(exec boil.Executor, stockDbxrefpropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stock_dbxrefprop\" where \"stock_dbxrefprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockDbxrefpropID)
	}

	row := exec.QueryRow(sql, stockDbxrefpropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stock_dbxrefprop exists")
	}

	return exists, nil
}

// StockDbxrefpropExistsG checks if the StockDbxrefprop row exists.
func StockDbxrefpropExistsG(stockDbxrefpropID int) (bool, error) {
	return StockDbxrefpropExists(boil.GetDB(), stockDbxrefpropID)
}

// StockDbxrefpropExistsGP checks if the StockDbxrefprop row exists. Panics on error.
func StockDbxrefpropExistsGP(stockDbxrefpropID int) bool {
	e, err := StockDbxrefpropExists(boil.GetDB(), stockDbxrefpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockDbxrefpropExistsP checks if the StockDbxrefprop row exists. Panics on error.
func StockDbxrefpropExistsP(exec boil.Executor, stockDbxrefpropID int) bool {
	e, err := StockDbxrefpropExists(exec, stockDbxrefpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
