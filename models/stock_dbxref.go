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
)

// StockDbxref is an object representing the database table.
type StockDbxref struct {
	StockDbxrefID int  `boil:"stock_dbxref_id" json:"stock_dbxref_id" toml:"stock_dbxref_id" yaml:"stock_dbxref_id"`
	StockID       int  `boil:"stock_id" json:"stock_id" toml:"stock_id" yaml:"stock_id"`
	DbxrefID      int  `boil:"dbxref_id" json:"dbxref_id" toml:"dbxref_id" yaml:"dbxref_id"`
	IsCurrent     bool `boil:"is_current" json:"is_current" toml:"is_current" yaml:"is_current"`

	R *stockDbxrefR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockDbxrefL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockDbxrefR is where relationships are stored.
type stockDbxrefR struct {
	Stock           *Stock
	Dbxref          *Dbxref
	StockDbxrefprop *StockDbxrefprop
}

// stockDbxrefL is where Load methods for each relationship are stored.
type stockDbxrefL struct{}

var (
	stockDbxrefColumns               = []string{"stock_dbxref_id", "stock_id", "dbxref_id", "is_current"}
	stockDbxrefColumnsWithoutDefault = []string{"stock_id", "dbxref_id"}
	stockDbxrefColumnsWithDefault    = []string{"stock_dbxref_id", "is_current"}
	stockDbxrefPrimaryKeyColumns     = []string{"stock_dbxref_id"}
)

type (
	// StockDbxrefSlice is an alias for a slice of pointers to StockDbxref.
	// This should generally be used opposed to []StockDbxref.
	StockDbxrefSlice []*StockDbxref
	// StockDbxrefHook is the signature for custom StockDbxref hook methods
	StockDbxrefHook func(boil.Executor, *StockDbxref) error

	stockDbxrefQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockDbxrefType                 = reflect.TypeOf(&StockDbxref{})
	stockDbxrefMapping              = queries.MakeStructMapping(stockDbxrefType)
	stockDbxrefPrimaryKeyMapping, _ = queries.BindMapping(stockDbxrefType, stockDbxrefMapping, stockDbxrefPrimaryKeyColumns)
	stockDbxrefInsertCacheMut       sync.RWMutex
	stockDbxrefInsertCache          = make(map[string]insertCache)
	stockDbxrefUpdateCacheMut       sync.RWMutex
	stockDbxrefUpdateCache          = make(map[string]updateCache)
	stockDbxrefUpsertCacheMut       sync.RWMutex
	stockDbxrefUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockDbxrefBeforeInsertHooks []StockDbxrefHook
var stockDbxrefBeforeUpdateHooks []StockDbxrefHook
var stockDbxrefBeforeDeleteHooks []StockDbxrefHook
var stockDbxrefBeforeUpsertHooks []StockDbxrefHook

var stockDbxrefAfterInsertHooks []StockDbxrefHook
var stockDbxrefAfterSelectHooks []StockDbxrefHook
var stockDbxrefAfterUpdateHooks []StockDbxrefHook
var stockDbxrefAfterDeleteHooks []StockDbxrefHook
var stockDbxrefAfterUpsertHooks []StockDbxrefHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockDbxref) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockDbxref) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockDbxref) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockDbxref) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockDbxref) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockDbxref) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockDbxref) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockDbxref) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockDbxref) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockDbxrefAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockDbxrefHook registers your hook function for all future operations.
func AddStockDbxrefHook(hookPoint boil.HookPoint, stockDbxrefHook StockDbxrefHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockDbxrefBeforeInsertHooks = append(stockDbxrefBeforeInsertHooks, stockDbxrefHook)
	case boil.BeforeUpdateHook:
		stockDbxrefBeforeUpdateHooks = append(stockDbxrefBeforeUpdateHooks, stockDbxrefHook)
	case boil.BeforeDeleteHook:
		stockDbxrefBeforeDeleteHooks = append(stockDbxrefBeforeDeleteHooks, stockDbxrefHook)
	case boil.BeforeUpsertHook:
		stockDbxrefBeforeUpsertHooks = append(stockDbxrefBeforeUpsertHooks, stockDbxrefHook)
	case boil.AfterInsertHook:
		stockDbxrefAfterInsertHooks = append(stockDbxrefAfterInsertHooks, stockDbxrefHook)
	case boil.AfterSelectHook:
		stockDbxrefAfterSelectHooks = append(stockDbxrefAfterSelectHooks, stockDbxrefHook)
	case boil.AfterUpdateHook:
		stockDbxrefAfterUpdateHooks = append(stockDbxrefAfterUpdateHooks, stockDbxrefHook)
	case boil.AfterDeleteHook:
		stockDbxrefAfterDeleteHooks = append(stockDbxrefAfterDeleteHooks, stockDbxrefHook)
	case boil.AfterUpsertHook:
		stockDbxrefAfterUpsertHooks = append(stockDbxrefAfterUpsertHooks, stockDbxrefHook)
	}
}

// OneP returns a single stockDbxref record from the query, and panics on error.
func (q stockDbxrefQuery) OneP() *StockDbxref {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockDbxref record from the query.
func (q stockDbxrefQuery) One() (*StockDbxref, error) {
	o := &StockDbxref{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stock_dbxref")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all StockDbxref records from the query, and panics on error.
func (q stockDbxrefQuery) AllP() StockDbxrefSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all StockDbxref records from the query.
func (q stockDbxrefQuery) All() (StockDbxrefSlice, error) {
	var o StockDbxrefSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StockDbxref slice")
	}

	if len(stockDbxrefAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all StockDbxref records in the query, and panics on error.
func (q stockDbxrefQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all StockDbxref records in the query.
func (q stockDbxrefQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stock_dbxref rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockDbxrefQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockDbxrefQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stock_dbxref exists")
	}

	return count > 0, nil
}

// StockG pointed to by the foreign key.
func (o *StockDbxref) StockG(mods ...qm.QueryMod) stockQuery {
	return o.Stock(boil.GetDB(), mods...)
}

// Stock pointed to by the foreign key.
func (o *StockDbxref) Stock(exec boil.Executor, mods ...qm.QueryMod) stockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := Stocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock\"")

	return query
}

// DbxrefG pointed to by the foreign key.
func (o *StockDbxref) DbxrefG(mods ...qm.QueryMod) dbxrefQuery {
	return o.Dbxref(boil.GetDB(), mods...)
}

// Dbxref pointed to by the foreign key.
func (o *StockDbxref) Dbxref(exec boil.Executor, mods ...qm.QueryMod) dbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := Dbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"dbxref\"")

	return query
}

// StockDbxrefpropG pointed to by the foreign key.
func (o *StockDbxref) StockDbxrefpropG(mods ...qm.QueryMod) stockDbxrefpropQuery {
	return o.StockDbxrefprop(boil.GetDB(), mods...)
}

// StockDbxrefprop pointed to by the foreign key.
func (o *StockDbxref) StockDbxrefprop(exec boil.Executor, mods ...qm.QueryMod) stockDbxrefpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_dbxref_id=$1", o.StockDbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := StockDbxrefprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_dbxrefprop\"")

	return query
}

// LoadStock allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockDbxrefL) LoadStock(e boil.Executor, singular bool, maybeStockDbxref interface{}) error {
	var slice []*StockDbxref
	var object *StockDbxref

	count := 1
	if singular {
		object = maybeStockDbxref.(*StockDbxref)
	} else {
		slice = *maybeStockDbxref.(*StockDbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockDbxrefR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockDbxrefR{}
			args[i] = obj.StockID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock\" where \"stock_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Stock")
	}
	defer results.Close()

	var resultSlice []*Stock
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Stock")
	}

	if len(stockDbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Stock = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockID == foreign.StockID {
				local.R.Stock = foreign
				break
			}
		}
	}

	return nil
}

// LoadDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockDbxrefL) LoadDbxref(e boil.Executor, singular bool, maybeStockDbxref interface{}) error {
	var slice []*StockDbxref
	var object *StockDbxref

	count := 1
	if singular {
		object = maybeStockDbxref.(*StockDbxref)
	} else {
		slice = *maybeStockDbxref.(*StockDbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockDbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &stockDbxrefR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"dbxref\" where \"dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Dbxref")
	}
	defer results.Close()

	var resultSlice []*Dbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Dbxref")
	}

	if len(stockDbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Dbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DbxrefID == foreign.DbxrefID {
				local.R.Dbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockDbxrefprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockDbxrefL) LoadStockDbxrefprop(e boil.Executor, singular bool, maybeStockDbxref interface{}) error {
	var slice []*StockDbxref
	var object *StockDbxref

	count := 1
	if singular {
		object = maybeStockDbxref.(*StockDbxref)
	} else {
		slice = *maybeStockDbxref.(*StockDbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockDbxrefR{}
		args[0] = object.StockDbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &stockDbxrefR{}
			args[i] = obj.StockDbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_dbxrefprop\" where \"stock_dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockDbxrefprop")
	}
	defer results.Close()

	var resultSlice []*StockDbxrefprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockDbxrefprop")
	}

	if len(stockDbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockDbxrefprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockDbxrefID == foreign.StockDbxrefID {
				local.R.StockDbxrefprop = foreign
				break
			}
		}
	}

	return nil
}

// SetStock of the stock_dbxref to the related item.
// Sets o.R.Stock to related.
// Adds o to related.R.StockDbxref.
func (o *StockDbxref) SetStock(exec boil.Executor, insert bool, related *Stock) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_dbxref\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stock_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockDbxrefPrimaryKeyColumns),
	)
	values := []interface{}{related.StockID, o.StockDbxrefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StockID = related.StockID

	if o.R == nil {
		o.R = &stockDbxrefR{
			Stock: related,
		}
	} else {
		o.R.Stock = related
	}

	if related.R == nil {
		related.R = &stockR{
			StockDbxref: o,
		}
	} else {
		related.R.StockDbxref = o
	}

	return nil
}

// SetDbxref of the stock_dbxref to the related item.
// Sets o.R.Dbxref to related.
// Adds o to related.R.StockDbxref.
func (o *StockDbxref) SetDbxref(exec boil.Executor, insert bool, related *Dbxref) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_dbxref\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockDbxrefPrimaryKeyColumns),
	)
	values := []interface{}{related.DbxrefID, o.StockDbxrefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DbxrefID = related.DbxrefID

	if o.R == nil {
		o.R = &stockDbxrefR{
			Dbxref: related,
		}
	} else {
		o.R.Dbxref = related
	}

	if related.R == nil {
		related.R = &dbxrefR{
			StockDbxref: o,
		}
	} else {
		related.R.StockDbxref = o
	}

	return nil
}

// SetStockDbxrefprop of the stock_dbxref to the related item.
// Sets o.R.StockDbxrefprop to related.
// Adds o to related.R.StockDbxref.
func (o *StockDbxref) SetStockDbxrefprop(exec boil.Executor, insert bool, related *StockDbxrefprop) error {
	var err error

	if insert {
		related.StockDbxrefID = o.StockDbxrefID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_dbxrefprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"stock_dbxref_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockDbxrefpropPrimaryKeyColumns),
		)
		values := []interface{}{o.StockDbxrefID, related.StockDbxrefpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.StockDbxrefID = o.StockDbxrefID

	}

	if o.R == nil {
		o.R = &stockDbxrefR{
			StockDbxrefprop: related,
		}
	} else {
		o.R.StockDbxrefprop = related
	}

	if related.R == nil {
		related.R = &stockDbxrefpropR{
			StockDbxref: o,
		}
	} else {
		related.R.StockDbxref = o
	}
	return nil
}

// StockDbxrefsG retrieves all records.
func StockDbxrefsG(mods ...qm.QueryMod) stockDbxrefQuery {
	return StockDbxrefs(boil.GetDB(), mods...)
}

// StockDbxrefs retrieves all the records using an executor.
func StockDbxrefs(exec boil.Executor, mods ...qm.QueryMod) stockDbxrefQuery {
	mods = append(mods, qm.From("\"stock_dbxref\""))
	return stockDbxrefQuery{NewQuery(exec, mods...)}
}

// FindStockDbxrefG retrieves a single record by ID.
func FindStockDbxrefG(stockDbxrefID int, selectCols ...string) (*StockDbxref, error) {
	return FindStockDbxref(boil.GetDB(), stockDbxrefID, selectCols...)
}

// FindStockDbxrefGP retrieves a single record by ID, and panics on error.
func FindStockDbxrefGP(stockDbxrefID int, selectCols ...string) *StockDbxref {
	retobj, err := FindStockDbxref(boil.GetDB(), stockDbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockDbxref retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockDbxref(exec boil.Executor, stockDbxrefID int, selectCols ...string) (*StockDbxref, error) {
	stockDbxrefObj := &StockDbxref{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stock_dbxref\" where \"stock_dbxref_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockDbxrefID)

	err := q.Bind(stockDbxrefObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stock_dbxref")
	}

	return stockDbxrefObj, nil
}

// FindStockDbxrefP retrieves a single record by ID with an executor, and panics on error.
func FindStockDbxrefP(exec boil.Executor, stockDbxrefID int, selectCols ...string) *StockDbxref {
	retobj, err := FindStockDbxref(exec, stockDbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockDbxref) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *StockDbxref) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *StockDbxref) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *StockDbxref) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_dbxref provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockDbxrefColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockDbxrefInsertCacheMut.RLock()
	cache, cached := stockDbxrefInsertCache[key]
	stockDbxrefInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockDbxrefColumns,
			stockDbxrefColumnsWithDefault,
			stockDbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockDbxrefType, stockDbxrefMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockDbxrefType, stockDbxrefMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stock_dbxref\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into stock_dbxref")
	}

	if !cached {
		stockDbxrefInsertCacheMut.Lock()
		stockDbxrefInsertCache[key] = cache
		stockDbxrefInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single StockDbxref record. See Update for
// whitelist behavior description.
func (o *StockDbxref) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single StockDbxref record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *StockDbxref) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the StockDbxref, and panics on error.
// See Update for whitelist behavior description.
func (o *StockDbxref) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the StockDbxref.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *StockDbxref) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockDbxrefUpdateCacheMut.RLock()
	cache, cached := stockDbxrefUpdateCache[key]
	stockDbxrefUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockDbxrefColumns, stockDbxrefPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update stock_dbxref, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stock_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockDbxrefPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockDbxrefType, stockDbxrefMapping, append(wl, stockDbxrefPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update stock_dbxref row")
	}

	if !cached {
		stockDbxrefUpdateCacheMut.Lock()
		stockDbxrefUpdateCache[key] = cache
		stockDbxrefUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockDbxrefQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockDbxrefQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for stock_dbxref")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockDbxrefSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockDbxrefSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockDbxrefSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockDbxrefSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stock_dbxref\" SET %s WHERE (\"stock_dbxref_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockDbxrefPrimaryKeyColumns), len(colNames)+1, len(stockDbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in stockDbxref slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockDbxref) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *StockDbxref) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *StockDbxref) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *StockDbxref) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_dbxref provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockDbxrefColumnsWithDefault, o)

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

	stockDbxrefUpsertCacheMut.RLock()
	cache, cached := stockDbxrefUpsertCache[key]
	stockDbxrefUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockDbxrefColumns,
			stockDbxrefColumnsWithDefault,
			stockDbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockDbxrefColumns,
			stockDbxrefPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert stock_dbxref, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockDbxrefPrimaryKeyColumns))
			copy(conflict, stockDbxrefPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stock_dbxref\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockDbxrefType, stockDbxrefMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockDbxrefType, stockDbxrefMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for stock_dbxref")
	}

	if !cached {
		stockDbxrefUpsertCacheMut.Lock()
		stockDbxrefUpsertCache[key] = cache
		stockDbxrefUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single StockDbxref record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockDbxref) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single StockDbxref record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockDbxref) DeleteG() error {
	if o == nil {
		return errors.New("models: no StockDbxref provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single StockDbxref record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockDbxref) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single StockDbxref record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockDbxref) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockDbxref provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockDbxrefPrimaryKeyMapping)
	sql := "DELETE FROM \"stock_dbxref\" WHERE \"stock_dbxref_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from stock_dbxref")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockDbxrefQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockDbxrefQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no stockDbxrefQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stock_dbxref")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockDbxrefSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockDbxrefSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no StockDbxref slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockDbxrefSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockDbxrefSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockDbxref slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockDbxrefBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stock_dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockDbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockDbxrefPrimaryKeyColumns), 1, len(stockDbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockDbxref slice")
	}

	if len(stockDbxrefAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *StockDbxref) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *StockDbxref) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockDbxref) ReloadG() error {
	if o == nil {
		return errors.New("models: no StockDbxref provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockDbxref) Reload(exec boil.Executor) error {
	ret, err := FindStockDbxref(exec, o.StockDbxrefID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockDbxrefSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockDbxrefSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockDbxrefSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty StockDbxrefSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockDbxrefSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockDbxrefs := StockDbxrefSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stock_dbxref\".* FROM \"stock_dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockDbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockDbxrefPrimaryKeyColumns), 1, len(stockDbxrefPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockDbxrefs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockDbxrefSlice")
	}

	*o = stockDbxrefs

	return nil
}

// StockDbxrefExists checks if the StockDbxref row exists.
func StockDbxrefExists(exec boil.Executor, stockDbxrefID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stock_dbxref\" where \"stock_dbxref_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockDbxrefID)
	}

	row := exec.QueryRow(sql, stockDbxrefID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stock_dbxref exists")
	}

	return exists, nil
}

// StockDbxrefExistsG checks if the StockDbxref row exists.
func StockDbxrefExistsG(stockDbxrefID int) (bool, error) {
	return StockDbxrefExists(boil.GetDB(), stockDbxrefID)
}

// StockDbxrefExistsGP checks if the StockDbxref row exists. Panics on error.
func StockDbxrefExistsGP(stockDbxrefID int) bool {
	e, err := StockDbxrefExists(boil.GetDB(), stockDbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockDbxrefExistsP checks if the StockDbxref row exists. Panics on error.
func StockDbxrefExistsP(exec boil.Executor, stockDbxrefID int) bool {
	e, err := StockDbxrefExists(exec, stockDbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
