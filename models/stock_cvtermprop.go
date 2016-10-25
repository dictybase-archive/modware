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

// StockCvtermprop is an object representing the database table.
type StockCvtermprop struct {
	StockCvtermpropID int         `boil:"stock_cvtermprop_id" json:"stock_cvtermprop_id" toml:"stock_cvtermprop_id" yaml:"stock_cvtermprop_id"`
	StockCvtermID     int         `boil:"stock_cvterm_id" json:"stock_cvterm_id" toml:"stock_cvterm_id" yaml:"stock_cvterm_id"`
	TypeID            int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value             null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank              int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *stockCvtermpropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockCvtermpropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockCvtermpropR is where relationships are stored.
type stockCvtermpropR struct {
	Type        *Cvterm
	StockCvterm *StockCvterm
}

// stockCvtermpropL is where Load methods for each relationship are stored.
type stockCvtermpropL struct{}

var (
	stockCvtermpropColumns               = []string{"stock_cvtermprop_id", "stock_cvterm_id", "type_id", "value", "rank"}
	stockCvtermpropColumnsWithoutDefault = []string{"stock_cvterm_id", "type_id", "value"}
	stockCvtermpropColumnsWithDefault    = []string{"stock_cvtermprop_id", "rank"}
	stockCvtermpropPrimaryKeyColumns     = []string{"stock_cvtermprop_id"}
)

type (
	// StockCvtermpropSlice is an alias for a slice of pointers to StockCvtermprop.
	// This should generally be used opposed to []StockCvtermprop.
	StockCvtermpropSlice []*StockCvtermprop
	// StockCvtermpropHook is the signature for custom StockCvtermprop hook methods
	StockCvtermpropHook func(boil.Executor, *StockCvtermprop) error

	stockCvtermpropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockCvtermpropType                 = reflect.TypeOf(&StockCvtermprop{})
	stockCvtermpropMapping              = queries.MakeStructMapping(stockCvtermpropType)
	stockCvtermpropPrimaryKeyMapping, _ = queries.BindMapping(stockCvtermpropType, stockCvtermpropMapping, stockCvtermpropPrimaryKeyColumns)
	stockCvtermpropInsertCacheMut       sync.RWMutex
	stockCvtermpropInsertCache          = make(map[string]insertCache)
	stockCvtermpropUpdateCacheMut       sync.RWMutex
	stockCvtermpropUpdateCache          = make(map[string]updateCache)
	stockCvtermpropUpsertCacheMut       sync.RWMutex
	stockCvtermpropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockCvtermpropBeforeInsertHooks []StockCvtermpropHook
var stockCvtermpropBeforeUpdateHooks []StockCvtermpropHook
var stockCvtermpropBeforeDeleteHooks []StockCvtermpropHook
var stockCvtermpropBeforeUpsertHooks []StockCvtermpropHook

var stockCvtermpropAfterInsertHooks []StockCvtermpropHook
var stockCvtermpropAfterSelectHooks []StockCvtermpropHook
var stockCvtermpropAfterUpdateHooks []StockCvtermpropHook
var stockCvtermpropAfterDeleteHooks []StockCvtermpropHook
var stockCvtermpropAfterUpsertHooks []StockCvtermpropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockCvtermprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermpropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockCvtermprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermpropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockCvtermprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermpropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockCvtermprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermpropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockCvtermprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermpropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockCvtermprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermpropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockCvtermprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermpropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockCvtermprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermpropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockCvtermprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockCvtermpropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockCvtermpropHook registers your hook function for all future operations.
func AddStockCvtermpropHook(hookPoint boil.HookPoint, stockCvtermpropHook StockCvtermpropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockCvtermpropBeforeInsertHooks = append(stockCvtermpropBeforeInsertHooks, stockCvtermpropHook)
	case boil.BeforeUpdateHook:
		stockCvtermpropBeforeUpdateHooks = append(stockCvtermpropBeforeUpdateHooks, stockCvtermpropHook)
	case boil.BeforeDeleteHook:
		stockCvtermpropBeforeDeleteHooks = append(stockCvtermpropBeforeDeleteHooks, stockCvtermpropHook)
	case boil.BeforeUpsertHook:
		stockCvtermpropBeforeUpsertHooks = append(stockCvtermpropBeforeUpsertHooks, stockCvtermpropHook)
	case boil.AfterInsertHook:
		stockCvtermpropAfterInsertHooks = append(stockCvtermpropAfterInsertHooks, stockCvtermpropHook)
	case boil.AfterSelectHook:
		stockCvtermpropAfterSelectHooks = append(stockCvtermpropAfterSelectHooks, stockCvtermpropHook)
	case boil.AfterUpdateHook:
		stockCvtermpropAfterUpdateHooks = append(stockCvtermpropAfterUpdateHooks, stockCvtermpropHook)
	case boil.AfterDeleteHook:
		stockCvtermpropAfterDeleteHooks = append(stockCvtermpropAfterDeleteHooks, stockCvtermpropHook)
	case boil.AfterUpsertHook:
		stockCvtermpropAfterUpsertHooks = append(stockCvtermpropAfterUpsertHooks, stockCvtermpropHook)
	}
}

// OneP returns a single stockCvtermprop record from the query, and panics on error.
func (q stockCvtermpropQuery) OneP() *StockCvtermprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockCvtermprop record from the query.
func (q stockCvtermpropQuery) One() (*StockCvtermprop, error) {
	o := &StockCvtermprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stock_cvtermprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all StockCvtermprop records from the query, and panics on error.
func (q stockCvtermpropQuery) AllP() StockCvtermpropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all StockCvtermprop records from the query.
func (q stockCvtermpropQuery) All() (StockCvtermpropSlice, error) {
	var o StockCvtermpropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StockCvtermprop slice")
	}

	if len(stockCvtermpropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all StockCvtermprop records in the query, and panics on error.
func (q stockCvtermpropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all StockCvtermprop records in the query.
func (q stockCvtermpropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stock_cvtermprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockCvtermpropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockCvtermpropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stock_cvtermprop exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *StockCvtermprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *StockCvtermprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// StockCvtermG pointed to by the foreign key.
func (o *StockCvtermprop) StockCvtermG(mods ...qm.QueryMod) stockCvtermQuery {
	return o.StockCvterm(boil.GetDB(), mods...)
}

// StockCvterm pointed to by the foreign key.
func (o *StockCvtermprop) StockCvterm(exec boil.Executor, mods ...qm.QueryMod) stockCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_cvterm_id=$1", o.StockCvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := StockCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_cvterm\"")

	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockCvtermpropL) LoadType(e boil.Executor, singular bool, maybeStockCvtermprop interface{}) error {
	var slice []*StockCvtermprop
	var object *StockCvtermprop

	count := 1
	if singular {
		object = maybeStockCvtermprop.(*StockCvtermprop)
	} else {
		slice = *maybeStockCvtermprop.(*StockCvtermpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockCvtermpropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &stockCvtermpropR{}
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

	if len(stockCvtermpropAfterSelectHooks) != 0 {
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

// LoadStockCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockCvtermpropL) LoadStockCvterm(e boil.Executor, singular bool, maybeStockCvtermprop interface{}) error {
	var slice []*StockCvtermprop
	var object *StockCvtermprop

	count := 1
	if singular {
		object = maybeStockCvtermprop.(*StockCvtermprop)
	} else {
		slice = *maybeStockCvtermprop.(*StockCvtermpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockCvtermpropR{}
		args[0] = object.StockCvtermID
	} else {
		for i, obj := range slice {
			obj.R = &stockCvtermpropR{}
			args[i] = obj.StockCvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_cvterm\" where \"stock_cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockCvterm")
	}
	defer results.Close()

	var resultSlice []*StockCvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockCvterm")
	}

	if len(stockCvtermpropAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockCvterm = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockCvtermID == foreign.StockCvtermID {
				local.R.StockCvterm = foreign
				break
			}
		}
	}

	return nil
}

// SetType of the stock_cvtermprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeStockCvtermprop.
func (o *StockCvtermprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_cvtermprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockCvtermpropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.StockCvtermpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &stockCvtermpropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeStockCvtermprop: o,
		}
	} else {
		related.R.TypeStockCvtermprop = o
	}

	return nil
}

// SetStockCvterm of the stock_cvtermprop to the related item.
// Sets o.R.StockCvterm to related.
// Adds o to related.R.StockCvtermprop.
func (o *StockCvtermprop) SetStockCvterm(exec boil.Executor, insert bool, related *StockCvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_cvtermprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stock_cvterm_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockCvtermpropPrimaryKeyColumns),
	)
	values := []interface{}{related.StockCvtermID, o.StockCvtermpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StockCvtermID = related.StockCvtermID

	if o.R == nil {
		o.R = &stockCvtermpropR{
			StockCvterm: related,
		}
	} else {
		o.R.StockCvterm = related
	}

	if related.R == nil {
		related.R = &stockCvtermR{
			StockCvtermprop: o,
		}
	} else {
		related.R.StockCvtermprop = o
	}

	return nil
}

// StockCvtermpropsG retrieves all records.
func StockCvtermpropsG(mods ...qm.QueryMod) stockCvtermpropQuery {
	return StockCvtermprops(boil.GetDB(), mods...)
}

// StockCvtermprops retrieves all the records using an executor.
func StockCvtermprops(exec boil.Executor, mods ...qm.QueryMod) stockCvtermpropQuery {
	mods = append(mods, qm.From("\"stock_cvtermprop\""))
	return stockCvtermpropQuery{NewQuery(exec, mods...)}
}

// FindStockCvtermpropG retrieves a single record by ID.
func FindStockCvtermpropG(stockCvtermpropID int, selectCols ...string) (*StockCvtermprop, error) {
	return FindStockCvtermprop(boil.GetDB(), stockCvtermpropID, selectCols...)
}

// FindStockCvtermpropGP retrieves a single record by ID, and panics on error.
func FindStockCvtermpropGP(stockCvtermpropID int, selectCols ...string) *StockCvtermprop {
	retobj, err := FindStockCvtermprop(boil.GetDB(), stockCvtermpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockCvtermprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockCvtermprop(exec boil.Executor, stockCvtermpropID int, selectCols ...string) (*StockCvtermprop, error) {
	stockCvtermpropObj := &StockCvtermprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stock_cvtermprop\" where \"stock_cvtermprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockCvtermpropID)

	err := q.Bind(stockCvtermpropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stock_cvtermprop")
	}

	return stockCvtermpropObj, nil
}

// FindStockCvtermpropP retrieves a single record by ID with an executor, and panics on error.
func FindStockCvtermpropP(exec boil.Executor, stockCvtermpropID int, selectCols ...string) *StockCvtermprop {
	retobj, err := FindStockCvtermprop(exec, stockCvtermpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockCvtermprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *StockCvtermprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *StockCvtermprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *StockCvtermprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_cvtermprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockCvtermpropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockCvtermpropInsertCacheMut.RLock()
	cache, cached := stockCvtermpropInsertCache[key]
	stockCvtermpropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockCvtermpropColumns,
			stockCvtermpropColumnsWithDefault,
			stockCvtermpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockCvtermpropType, stockCvtermpropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockCvtermpropType, stockCvtermpropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stock_cvtermprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into stock_cvtermprop")
	}

	if !cached {
		stockCvtermpropInsertCacheMut.Lock()
		stockCvtermpropInsertCache[key] = cache
		stockCvtermpropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single StockCvtermprop record. See Update for
// whitelist behavior description.
func (o *StockCvtermprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single StockCvtermprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *StockCvtermprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the StockCvtermprop, and panics on error.
// See Update for whitelist behavior description.
func (o *StockCvtermprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the StockCvtermprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *StockCvtermprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockCvtermpropUpdateCacheMut.RLock()
	cache, cached := stockCvtermpropUpdateCache[key]
	stockCvtermpropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockCvtermpropColumns, stockCvtermpropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update stock_cvtermprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stock_cvtermprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockCvtermpropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockCvtermpropType, stockCvtermpropMapping, append(wl, stockCvtermpropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update stock_cvtermprop row")
	}

	if !cached {
		stockCvtermpropUpdateCacheMut.Lock()
		stockCvtermpropUpdateCache[key] = cache
		stockCvtermpropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockCvtermpropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockCvtermpropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for stock_cvtermprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockCvtermpropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockCvtermpropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockCvtermpropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockCvtermpropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockCvtermpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stock_cvtermprop\" SET %s WHERE (\"stock_cvtermprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockCvtermpropPrimaryKeyColumns), len(colNames)+1, len(stockCvtermpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in stockCvtermprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockCvtermprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *StockCvtermprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *StockCvtermprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *StockCvtermprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_cvtermprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockCvtermpropColumnsWithDefault, o)

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

	stockCvtermpropUpsertCacheMut.RLock()
	cache, cached := stockCvtermpropUpsertCache[key]
	stockCvtermpropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockCvtermpropColumns,
			stockCvtermpropColumnsWithDefault,
			stockCvtermpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockCvtermpropColumns,
			stockCvtermpropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert stock_cvtermprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockCvtermpropPrimaryKeyColumns))
			copy(conflict, stockCvtermpropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stock_cvtermprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockCvtermpropType, stockCvtermpropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockCvtermpropType, stockCvtermpropMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for stock_cvtermprop")
	}

	if !cached {
		stockCvtermpropUpsertCacheMut.Lock()
		stockCvtermpropUpsertCache[key] = cache
		stockCvtermpropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single StockCvtermprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockCvtermprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single StockCvtermprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockCvtermprop) DeleteG() error {
	if o == nil {
		return errors.New("models: no StockCvtermprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single StockCvtermprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockCvtermprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single StockCvtermprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockCvtermprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockCvtermprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockCvtermpropPrimaryKeyMapping)
	sql := "DELETE FROM \"stock_cvtermprop\" WHERE \"stock_cvtermprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from stock_cvtermprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockCvtermpropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockCvtermpropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no stockCvtermpropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stock_cvtermprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockCvtermpropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockCvtermpropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no StockCvtermprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockCvtermpropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockCvtermpropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockCvtermprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockCvtermpropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockCvtermpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stock_cvtermprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockCvtermpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockCvtermpropPrimaryKeyColumns), 1, len(stockCvtermpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockCvtermprop slice")
	}

	if len(stockCvtermpropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *StockCvtermprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *StockCvtermprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockCvtermprop) ReloadG() error {
	if o == nil {
		return errors.New("models: no StockCvtermprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockCvtermprop) Reload(exec boil.Executor) error {
	ret, err := FindStockCvtermprop(exec, o.StockCvtermpropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockCvtermpropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockCvtermpropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockCvtermpropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty StockCvtermpropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockCvtermpropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockCvtermprops := StockCvtermpropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockCvtermpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stock_cvtermprop\".* FROM \"stock_cvtermprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockCvtermpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockCvtermpropPrimaryKeyColumns), 1, len(stockCvtermpropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockCvtermprops)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockCvtermpropSlice")
	}

	*o = stockCvtermprops

	return nil
}

// StockCvtermpropExists checks if the StockCvtermprop row exists.
func StockCvtermpropExists(exec boil.Executor, stockCvtermpropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stock_cvtermprop\" where \"stock_cvtermprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockCvtermpropID)
	}

	row := exec.QueryRow(sql, stockCvtermpropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stock_cvtermprop exists")
	}

	return exists, nil
}

// StockCvtermpropExistsG checks if the StockCvtermprop row exists.
func StockCvtermpropExistsG(stockCvtermpropID int) (bool, error) {
	return StockCvtermpropExists(boil.GetDB(), stockCvtermpropID)
}

// StockCvtermpropExistsGP checks if the StockCvtermprop row exists. Panics on error.
func StockCvtermpropExistsGP(stockCvtermpropID int) bool {
	e, err := StockCvtermpropExists(boil.GetDB(), stockCvtermpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockCvtermpropExistsP checks if the StockCvtermprop row exists. Panics on error.
func StockCvtermpropExistsP(exec boil.Executor, stockCvtermpropID int) bool {
	e, err := StockCvtermpropExists(exec, stockCvtermpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
