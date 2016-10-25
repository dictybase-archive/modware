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

// StockcollectionStock is an object representing the database table.
type StockcollectionStock struct {
	StockcollectionStockID int `boil:"stockcollection_stock_id" json:"stockcollection_stock_id" toml:"stockcollection_stock_id" yaml:"stockcollection_stock_id"`
	StockcollectionID      int `boil:"stockcollection_id" json:"stockcollection_id" toml:"stockcollection_id" yaml:"stockcollection_id"`
	StockID                int `boil:"stock_id" json:"stock_id" toml:"stock_id" yaml:"stock_id"`

	R *stockcollectionStockR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockcollectionStockL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockcollectionStockR is where relationships are stored.
type stockcollectionStockR struct {
	Stock           *Stock
	Stockcollection *Stockcollection
}

// stockcollectionStockL is where Load methods for each relationship are stored.
type stockcollectionStockL struct{}

var (
	stockcollectionStockColumns               = []string{"stockcollection_stock_id", "stockcollection_id", "stock_id"}
	stockcollectionStockColumnsWithoutDefault = []string{"stockcollection_id", "stock_id"}
	stockcollectionStockColumnsWithDefault    = []string{"stockcollection_stock_id"}
	stockcollectionStockPrimaryKeyColumns     = []string{"stockcollection_stock_id"}
)

type (
	// StockcollectionStockSlice is an alias for a slice of pointers to StockcollectionStock.
	// This should generally be used opposed to []StockcollectionStock.
	StockcollectionStockSlice []*StockcollectionStock
	// StockcollectionStockHook is the signature for custom StockcollectionStock hook methods
	StockcollectionStockHook func(boil.Executor, *StockcollectionStock) error

	stockcollectionStockQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockcollectionStockType                 = reflect.TypeOf(&StockcollectionStock{})
	stockcollectionStockMapping              = queries.MakeStructMapping(stockcollectionStockType)
	stockcollectionStockPrimaryKeyMapping, _ = queries.BindMapping(stockcollectionStockType, stockcollectionStockMapping, stockcollectionStockPrimaryKeyColumns)
	stockcollectionStockInsertCacheMut       sync.RWMutex
	stockcollectionStockInsertCache          = make(map[string]insertCache)
	stockcollectionStockUpdateCacheMut       sync.RWMutex
	stockcollectionStockUpdateCache          = make(map[string]updateCache)
	stockcollectionStockUpsertCacheMut       sync.RWMutex
	stockcollectionStockUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockcollectionStockBeforeInsertHooks []StockcollectionStockHook
var stockcollectionStockBeforeUpdateHooks []StockcollectionStockHook
var stockcollectionStockBeforeDeleteHooks []StockcollectionStockHook
var stockcollectionStockBeforeUpsertHooks []StockcollectionStockHook

var stockcollectionStockAfterInsertHooks []StockcollectionStockHook
var stockcollectionStockAfterSelectHooks []StockcollectionStockHook
var stockcollectionStockAfterUpdateHooks []StockcollectionStockHook
var stockcollectionStockAfterDeleteHooks []StockcollectionStockHook
var stockcollectionStockAfterUpsertHooks []StockcollectionStockHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockcollectionStock) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionStockBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockcollectionStock) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionStockBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockcollectionStock) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionStockBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockcollectionStock) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionStockBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockcollectionStock) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionStockAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockcollectionStock) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionStockAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockcollectionStock) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionStockAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockcollectionStock) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionStockAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockcollectionStock) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionStockAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockcollectionStockHook registers your hook function for all future operations.
func AddStockcollectionStockHook(hookPoint boil.HookPoint, stockcollectionStockHook StockcollectionStockHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockcollectionStockBeforeInsertHooks = append(stockcollectionStockBeforeInsertHooks, stockcollectionStockHook)
	case boil.BeforeUpdateHook:
		stockcollectionStockBeforeUpdateHooks = append(stockcollectionStockBeforeUpdateHooks, stockcollectionStockHook)
	case boil.BeforeDeleteHook:
		stockcollectionStockBeforeDeleteHooks = append(stockcollectionStockBeforeDeleteHooks, stockcollectionStockHook)
	case boil.BeforeUpsertHook:
		stockcollectionStockBeforeUpsertHooks = append(stockcollectionStockBeforeUpsertHooks, stockcollectionStockHook)
	case boil.AfterInsertHook:
		stockcollectionStockAfterInsertHooks = append(stockcollectionStockAfterInsertHooks, stockcollectionStockHook)
	case boil.AfterSelectHook:
		stockcollectionStockAfterSelectHooks = append(stockcollectionStockAfterSelectHooks, stockcollectionStockHook)
	case boil.AfterUpdateHook:
		stockcollectionStockAfterUpdateHooks = append(stockcollectionStockAfterUpdateHooks, stockcollectionStockHook)
	case boil.AfterDeleteHook:
		stockcollectionStockAfterDeleteHooks = append(stockcollectionStockAfterDeleteHooks, stockcollectionStockHook)
	case boil.AfterUpsertHook:
		stockcollectionStockAfterUpsertHooks = append(stockcollectionStockAfterUpsertHooks, stockcollectionStockHook)
	}
}

// OneP returns a single stockcollectionStock record from the query, and panics on error.
func (q stockcollectionStockQuery) OneP() *StockcollectionStock {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockcollectionStock record from the query.
func (q stockcollectionStockQuery) One() (*StockcollectionStock, error) {
	o := &StockcollectionStock{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stockcollection_stock")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all StockcollectionStock records from the query, and panics on error.
func (q stockcollectionStockQuery) AllP() StockcollectionStockSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all StockcollectionStock records from the query.
func (q stockcollectionStockQuery) All() (StockcollectionStockSlice, error) {
	var o StockcollectionStockSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StockcollectionStock slice")
	}

	if len(stockcollectionStockAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all StockcollectionStock records in the query, and panics on error.
func (q stockcollectionStockQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all StockcollectionStock records in the query.
func (q stockcollectionStockQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stockcollection_stock rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockcollectionStockQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockcollectionStockQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stockcollection_stock exists")
	}

	return count > 0, nil
}

// StockG pointed to by the foreign key.
func (o *StockcollectionStock) StockG(mods ...qm.QueryMod) stockQuery {
	return o.Stock(boil.GetDB(), mods...)
}

// Stock pointed to by the foreign key.
func (o *StockcollectionStock) Stock(exec boil.Executor, mods ...qm.QueryMod) stockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := Stocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock\"")

	return query
}

// StockcollectionG pointed to by the foreign key.
func (o *StockcollectionStock) StockcollectionG(mods ...qm.QueryMod) stockcollectionQuery {
	return o.Stockcollection(boil.GetDB(), mods...)
}

// Stockcollection pointed to by the foreign key.
func (o *StockcollectionStock) Stockcollection(exec boil.Executor, mods ...qm.QueryMod) stockcollectionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stockcollection_id=$1", o.StockcollectionID),
	}

	queryMods = append(queryMods, mods...)

	query := Stockcollections(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stockcollection\"")

	return query
}

// LoadStock allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockcollectionStockL) LoadStock(e boil.Executor, singular bool, maybeStockcollectionStock interface{}) error {
	var slice []*StockcollectionStock
	var object *StockcollectionStock

	count := 1
	if singular {
		object = maybeStockcollectionStock.(*StockcollectionStock)
	} else {
		slice = *maybeStockcollectionStock.(*StockcollectionStockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockcollectionStockR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockcollectionStockR{}
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

	if len(stockcollectionStockAfterSelectHooks) != 0 {
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

// LoadStockcollection allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockcollectionStockL) LoadStockcollection(e boil.Executor, singular bool, maybeStockcollectionStock interface{}) error {
	var slice []*StockcollectionStock
	var object *StockcollectionStock

	count := 1
	if singular {
		object = maybeStockcollectionStock.(*StockcollectionStock)
	} else {
		slice = *maybeStockcollectionStock.(*StockcollectionStockSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockcollectionStockR{}
		args[0] = object.StockcollectionID
	} else {
		for i, obj := range slice {
			obj.R = &stockcollectionStockR{}
			args[i] = obj.StockcollectionID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stockcollection\" where \"stockcollection_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Stockcollection")
	}
	defer results.Close()

	var resultSlice []*Stockcollection
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Stockcollection")
	}

	if len(stockcollectionStockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Stockcollection = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockcollectionID == foreign.StockcollectionID {
				local.R.Stockcollection = foreign
				break
			}
		}
	}

	return nil
}

// SetStock of the stockcollection_stock to the related item.
// Sets o.R.Stock to related.
// Adds o to related.R.StockcollectionStock.
func (o *StockcollectionStock) SetStock(exec boil.Executor, insert bool, related *Stock) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stockcollection_stock\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stock_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockcollectionStockPrimaryKeyColumns),
	)
	values := []interface{}{related.StockID, o.StockcollectionStockID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StockID = related.StockID

	if o.R == nil {
		o.R = &stockcollectionStockR{
			Stock: related,
		}
	} else {
		o.R.Stock = related
	}

	if related.R == nil {
		related.R = &stockR{
			StockcollectionStock: o,
		}
	} else {
		related.R.StockcollectionStock = o
	}

	return nil
}

// SetStockcollection of the stockcollection_stock to the related item.
// Sets o.R.Stockcollection to related.
// Adds o to related.R.StockcollectionStock.
func (o *StockcollectionStock) SetStockcollection(exec boil.Executor, insert bool, related *Stockcollection) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stockcollection_stock\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stockcollection_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockcollectionStockPrimaryKeyColumns),
	)
	values := []interface{}{related.StockcollectionID, o.StockcollectionStockID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StockcollectionID = related.StockcollectionID

	if o.R == nil {
		o.R = &stockcollectionStockR{
			Stockcollection: related,
		}
	} else {
		o.R.Stockcollection = related
	}

	if related.R == nil {
		related.R = &stockcollectionR{
			StockcollectionStock: o,
		}
	} else {
		related.R.StockcollectionStock = o
	}

	return nil
}

// StockcollectionStocksG retrieves all records.
func StockcollectionStocksG(mods ...qm.QueryMod) stockcollectionStockQuery {
	return StockcollectionStocks(boil.GetDB(), mods...)
}

// StockcollectionStocks retrieves all the records using an executor.
func StockcollectionStocks(exec boil.Executor, mods ...qm.QueryMod) stockcollectionStockQuery {
	mods = append(mods, qm.From("\"stockcollection_stock\""))
	return stockcollectionStockQuery{NewQuery(exec, mods...)}
}

// FindStockcollectionStockG retrieves a single record by ID.
func FindStockcollectionStockG(stockcollectionStockID int, selectCols ...string) (*StockcollectionStock, error) {
	return FindStockcollectionStock(boil.GetDB(), stockcollectionStockID, selectCols...)
}

// FindStockcollectionStockGP retrieves a single record by ID, and panics on error.
func FindStockcollectionStockGP(stockcollectionStockID int, selectCols ...string) *StockcollectionStock {
	retobj, err := FindStockcollectionStock(boil.GetDB(), stockcollectionStockID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockcollectionStock retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockcollectionStock(exec boil.Executor, stockcollectionStockID int, selectCols ...string) (*StockcollectionStock, error) {
	stockcollectionStockObj := &StockcollectionStock{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stockcollection_stock\" where \"stockcollection_stock_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockcollectionStockID)

	err := q.Bind(stockcollectionStockObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stockcollection_stock")
	}

	return stockcollectionStockObj, nil
}

// FindStockcollectionStockP retrieves a single record by ID with an executor, and panics on error.
func FindStockcollectionStockP(exec boil.Executor, stockcollectionStockID int, selectCols ...string) *StockcollectionStock {
	retobj, err := FindStockcollectionStock(exec, stockcollectionStockID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockcollectionStock) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *StockcollectionStock) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *StockcollectionStock) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *StockcollectionStock) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stockcollection_stock provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockcollectionStockColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockcollectionStockInsertCacheMut.RLock()
	cache, cached := stockcollectionStockInsertCache[key]
	stockcollectionStockInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockcollectionStockColumns,
			stockcollectionStockColumnsWithDefault,
			stockcollectionStockColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockcollectionStockType, stockcollectionStockMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockcollectionStockType, stockcollectionStockMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stockcollection_stock\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into stockcollection_stock")
	}

	if !cached {
		stockcollectionStockInsertCacheMut.Lock()
		stockcollectionStockInsertCache[key] = cache
		stockcollectionStockInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single StockcollectionStock record. See Update for
// whitelist behavior description.
func (o *StockcollectionStock) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single StockcollectionStock record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *StockcollectionStock) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the StockcollectionStock, and panics on error.
// See Update for whitelist behavior description.
func (o *StockcollectionStock) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the StockcollectionStock.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *StockcollectionStock) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockcollectionStockUpdateCacheMut.RLock()
	cache, cached := stockcollectionStockUpdateCache[key]
	stockcollectionStockUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockcollectionStockColumns, stockcollectionStockPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update stockcollection_stock, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stockcollection_stock\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockcollectionStockPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockcollectionStockType, stockcollectionStockMapping, append(wl, stockcollectionStockPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update stockcollection_stock row")
	}

	if !cached {
		stockcollectionStockUpdateCacheMut.Lock()
		stockcollectionStockUpdateCache[key] = cache
		stockcollectionStockUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockcollectionStockQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockcollectionStockQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for stockcollection_stock")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockcollectionStockSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockcollectionStockSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockcollectionStockSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockcollectionStockSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockcollectionStockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stockcollection_stock\" SET %s WHERE (\"stockcollection_stock_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockcollectionStockPrimaryKeyColumns), len(colNames)+1, len(stockcollectionStockPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in stockcollectionStock slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockcollectionStock) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *StockcollectionStock) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *StockcollectionStock) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *StockcollectionStock) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stockcollection_stock provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockcollectionStockColumnsWithDefault, o)

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

	stockcollectionStockUpsertCacheMut.RLock()
	cache, cached := stockcollectionStockUpsertCache[key]
	stockcollectionStockUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockcollectionStockColumns,
			stockcollectionStockColumnsWithDefault,
			stockcollectionStockColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockcollectionStockColumns,
			stockcollectionStockPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert stockcollection_stock, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockcollectionStockPrimaryKeyColumns))
			copy(conflict, stockcollectionStockPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stockcollection_stock\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockcollectionStockType, stockcollectionStockMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockcollectionStockType, stockcollectionStockMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for stockcollection_stock")
	}

	if !cached {
		stockcollectionStockUpsertCacheMut.Lock()
		stockcollectionStockUpsertCache[key] = cache
		stockcollectionStockUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single StockcollectionStock record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockcollectionStock) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single StockcollectionStock record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockcollectionStock) DeleteG() error {
	if o == nil {
		return errors.New("models: no StockcollectionStock provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single StockcollectionStock record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockcollectionStock) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single StockcollectionStock record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockcollectionStock) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockcollectionStock provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockcollectionStockPrimaryKeyMapping)
	sql := "DELETE FROM \"stockcollection_stock\" WHERE \"stockcollection_stock_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from stockcollection_stock")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockcollectionStockQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockcollectionStockQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no stockcollectionStockQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockcollection_stock")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockcollectionStockSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockcollectionStockSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no StockcollectionStock slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockcollectionStockSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockcollectionStockSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockcollectionStock slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockcollectionStockBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockcollectionStockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stockcollection_stock\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockcollectionStockPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockcollectionStockPrimaryKeyColumns), 1, len(stockcollectionStockPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockcollectionStock slice")
	}

	if len(stockcollectionStockAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *StockcollectionStock) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *StockcollectionStock) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockcollectionStock) ReloadG() error {
	if o == nil {
		return errors.New("models: no StockcollectionStock provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockcollectionStock) Reload(exec boil.Executor) error {
	ret, err := FindStockcollectionStock(exec, o.StockcollectionStockID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockcollectionStockSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockcollectionStockSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockcollectionStockSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty StockcollectionStockSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockcollectionStockSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockcollectionStocks := StockcollectionStockSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockcollectionStockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stockcollection_stock\".* FROM \"stockcollection_stock\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockcollectionStockPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockcollectionStockPrimaryKeyColumns), 1, len(stockcollectionStockPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockcollectionStocks)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockcollectionStockSlice")
	}

	*o = stockcollectionStocks

	return nil
}

// StockcollectionStockExists checks if the StockcollectionStock row exists.
func StockcollectionStockExists(exec boil.Executor, stockcollectionStockID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stockcollection_stock\" where \"stockcollection_stock_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockcollectionStockID)
	}

	row := exec.QueryRow(sql, stockcollectionStockID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stockcollection_stock exists")
	}

	return exists, nil
}

// StockcollectionStockExistsG checks if the StockcollectionStock row exists.
func StockcollectionStockExistsG(stockcollectionStockID int) (bool, error) {
	return StockcollectionStockExists(boil.GetDB(), stockcollectionStockID)
}

// StockcollectionStockExistsGP checks if the StockcollectionStock row exists. Panics on error.
func StockcollectionStockExistsGP(stockcollectionStockID int) bool {
	e, err := StockcollectionStockExists(boil.GetDB(), stockcollectionStockID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockcollectionStockExistsP checks if the StockcollectionStock row exists. Panics on error.
func StockcollectionStockExistsP(exec boil.Executor, stockcollectionStockID int) bool {
	e, err := StockcollectionStockExists(exec, stockcollectionStockID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
