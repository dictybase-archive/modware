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

// StockPub is an object representing the database table.
type StockPub struct {
	StockPubID int `boil:"stock_pub_id" json:"stock_pub_id" toml:"stock_pub_id" yaml:"stock_pub_id"`
	StockID    int `boil:"stock_id" json:"stock_id" toml:"stock_id" yaml:"stock_id"`
	PubID      int `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`

	R *stockPubR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockPubL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockPubR is where relationships are stored.
type stockPubR struct {
	Pub   *Pub
	Stock *Stock
}

// stockPubL is where Load methods for each relationship are stored.
type stockPubL struct{}

var (
	stockPubColumns               = []string{"stock_pub_id", "stock_id", "pub_id"}
	stockPubColumnsWithoutDefault = []string{"stock_id", "pub_id"}
	stockPubColumnsWithDefault    = []string{"stock_pub_id"}
	stockPubPrimaryKeyColumns     = []string{"stock_pub_id"}
)

type (
	// StockPubSlice is an alias for a slice of pointers to StockPub.
	// This should generally be used opposed to []StockPub.
	StockPubSlice []*StockPub
	// StockPubHook is the signature for custom StockPub hook methods
	StockPubHook func(boil.Executor, *StockPub) error

	stockPubQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockPubType                 = reflect.TypeOf(&StockPub{})
	stockPubMapping              = queries.MakeStructMapping(stockPubType)
	stockPubPrimaryKeyMapping, _ = queries.BindMapping(stockPubType, stockPubMapping, stockPubPrimaryKeyColumns)
	stockPubInsertCacheMut       sync.RWMutex
	stockPubInsertCache          = make(map[string]insertCache)
	stockPubUpdateCacheMut       sync.RWMutex
	stockPubUpdateCache          = make(map[string]updateCache)
	stockPubUpsertCacheMut       sync.RWMutex
	stockPubUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockPubBeforeInsertHooks []StockPubHook
var stockPubBeforeUpdateHooks []StockPubHook
var stockPubBeforeDeleteHooks []StockPubHook
var stockPubBeforeUpsertHooks []StockPubHook

var stockPubAfterInsertHooks []StockPubHook
var stockPubAfterSelectHooks []StockPubHook
var stockPubAfterUpdateHooks []StockPubHook
var stockPubAfterDeleteHooks []StockPubHook
var stockPubAfterUpsertHooks []StockPubHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockPub) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockPubBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockPub) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockPubBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockPub) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockPubBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockPub) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockPubBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockPub) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockPubAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockPub) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockPubAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockPub) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockPubAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockPub) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockPubAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockPub) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockPubAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockPubHook registers your hook function for all future operations.
func AddStockPubHook(hookPoint boil.HookPoint, stockPubHook StockPubHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockPubBeforeInsertHooks = append(stockPubBeforeInsertHooks, stockPubHook)
	case boil.BeforeUpdateHook:
		stockPubBeforeUpdateHooks = append(stockPubBeforeUpdateHooks, stockPubHook)
	case boil.BeforeDeleteHook:
		stockPubBeforeDeleteHooks = append(stockPubBeforeDeleteHooks, stockPubHook)
	case boil.BeforeUpsertHook:
		stockPubBeforeUpsertHooks = append(stockPubBeforeUpsertHooks, stockPubHook)
	case boil.AfterInsertHook:
		stockPubAfterInsertHooks = append(stockPubAfterInsertHooks, stockPubHook)
	case boil.AfterSelectHook:
		stockPubAfterSelectHooks = append(stockPubAfterSelectHooks, stockPubHook)
	case boil.AfterUpdateHook:
		stockPubAfterUpdateHooks = append(stockPubAfterUpdateHooks, stockPubHook)
	case boil.AfterDeleteHook:
		stockPubAfterDeleteHooks = append(stockPubAfterDeleteHooks, stockPubHook)
	case boil.AfterUpsertHook:
		stockPubAfterUpsertHooks = append(stockPubAfterUpsertHooks, stockPubHook)
	}
}

// OneP returns a single stockPub record from the query, and panics on error.
func (q stockPubQuery) OneP() *StockPub {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockPub record from the query.
func (q stockPubQuery) One() (*StockPub, error) {
	o := &StockPub{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stock_pub")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all StockPub records from the query, and panics on error.
func (q stockPubQuery) AllP() StockPubSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all StockPub records from the query.
func (q stockPubQuery) All() (StockPubSlice, error) {
	var o StockPubSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StockPub slice")
	}

	if len(stockPubAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all StockPub records in the query, and panics on error.
func (q stockPubQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all StockPub records in the query.
func (q stockPubQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stock_pub rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockPubQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockPubQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stock_pub exists")
	}

	return count > 0, nil
}

// PubG pointed to by the foreign key.
func (o *StockPub) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *StockPub) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// StockG pointed to by the foreign key.
func (o *StockPub) StockG(mods ...qm.QueryMod) stockQuery {
	return o.Stock(boil.GetDB(), mods...)
}

// Stock pointed to by the foreign key.
func (o *StockPub) Stock(exec boil.Executor, mods ...qm.QueryMod) stockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := Stocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock\"")

	return query
}

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockPubL) LoadPub(e boil.Executor, singular bool, maybeStockPub interface{}) error {
	var slice []*StockPub
	var object *StockPub

	count := 1
	if singular {
		object = maybeStockPub.(*StockPub)
	} else {
		slice = *maybeStockPub.(*StockPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockPubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &stockPubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"pub\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Pub")
	}
	defer results.Close()

	var resultSlice []*Pub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Pub")
	}

	if len(stockPubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Pub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.Pub = foreign
				break
			}
		}
	}

	return nil
}

// LoadStock allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockPubL) LoadStock(e boil.Executor, singular bool, maybeStockPub interface{}) error {
	var slice []*StockPub
	var object *StockPub

	count := 1
	if singular {
		object = maybeStockPub.(*StockPub)
	} else {
		slice = *maybeStockPub.(*StockPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockPubR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockPubR{}
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

	if len(stockPubAfterSelectHooks) != 0 {
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

// SetPub of the stock_pub to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.StockPub.
func (o *StockPub) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockPubPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.StockPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &stockPubR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			StockPub: o,
		}
	} else {
		related.R.StockPub = o
	}

	return nil
}

// SetStock of the stock_pub to the related item.
// Sets o.R.Stock to related.
// Adds o to related.R.StockPub.
func (o *StockPub) SetStock(exec boil.Executor, insert bool, related *Stock) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stock_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockPubPrimaryKeyColumns),
	)
	values := []interface{}{related.StockID, o.StockPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StockID = related.StockID

	if o.R == nil {
		o.R = &stockPubR{
			Stock: related,
		}
	} else {
		o.R.Stock = related
	}

	if related.R == nil {
		related.R = &stockR{
			StockPub: o,
		}
	} else {
		related.R.StockPub = o
	}

	return nil
}

// StockPubsG retrieves all records.
func StockPubsG(mods ...qm.QueryMod) stockPubQuery {
	return StockPubs(boil.GetDB(), mods...)
}

// StockPubs retrieves all the records using an executor.
func StockPubs(exec boil.Executor, mods ...qm.QueryMod) stockPubQuery {
	mods = append(mods, qm.From("\"stock_pub\""))
	return stockPubQuery{NewQuery(exec, mods...)}
}

// FindStockPubG retrieves a single record by ID.
func FindStockPubG(stockPubID int, selectCols ...string) (*StockPub, error) {
	return FindStockPub(boil.GetDB(), stockPubID, selectCols...)
}

// FindStockPubGP retrieves a single record by ID, and panics on error.
func FindStockPubGP(stockPubID int, selectCols ...string) *StockPub {
	retobj, err := FindStockPub(boil.GetDB(), stockPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockPub retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockPub(exec boil.Executor, stockPubID int, selectCols ...string) (*StockPub, error) {
	stockPubObj := &StockPub{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stock_pub\" where \"stock_pub_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockPubID)

	err := q.Bind(stockPubObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stock_pub")
	}

	return stockPubObj, nil
}

// FindStockPubP retrieves a single record by ID with an executor, and panics on error.
func FindStockPubP(exec boil.Executor, stockPubID int, selectCols ...string) *StockPub {
	retobj, err := FindStockPub(exec, stockPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockPub) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *StockPub) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *StockPub) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *StockPub) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_pub provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockPubColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockPubInsertCacheMut.RLock()
	cache, cached := stockPubInsertCache[key]
	stockPubInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockPubColumns,
			stockPubColumnsWithDefault,
			stockPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockPubType, stockPubMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockPubType, stockPubMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stock_pub\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into stock_pub")
	}

	if !cached {
		stockPubInsertCacheMut.Lock()
		stockPubInsertCache[key] = cache
		stockPubInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single StockPub record. See Update for
// whitelist behavior description.
func (o *StockPub) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single StockPub record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *StockPub) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the StockPub, and panics on error.
// See Update for whitelist behavior description.
func (o *StockPub) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the StockPub.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *StockPub) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockPubUpdateCacheMut.RLock()
	cache, cached := stockPubUpdateCache[key]
	stockPubUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockPubColumns, stockPubPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update stock_pub, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stock_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockPubPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockPubType, stockPubMapping, append(wl, stockPubPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update stock_pub row")
	}

	if !cached {
		stockPubUpdateCacheMut.Lock()
		stockPubUpdateCache[key] = cache
		stockPubUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockPubQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockPubQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for stock_pub")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockPubSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockPubSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockPubSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockPubSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stock_pub\" SET %s WHERE (\"stock_pub_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockPubPrimaryKeyColumns), len(colNames)+1, len(stockPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in stockPub slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockPub) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *StockPub) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *StockPub) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *StockPub) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_pub provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockPubColumnsWithDefault, o)

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

	stockPubUpsertCacheMut.RLock()
	cache, cached := stockPubUpsertCache[key]
	stockPubUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockPubColumns,
			stockPubColumnsWithDefault,
			stockPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockPubColumns,
			stockPubPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert stock_pub, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockPubPrimaryKeyColumns))
			copy(conflict, stockPubPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stock_pub\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockPubType, stockPubMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockPubType, stockPubMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for stock_pub")
	}

	if !cached {
		stockPubUpsertCacheMut.Lock()
		stockPubUpsertCache[key] = cache
		stockPubUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single StockPub record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockPub) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single StockPub record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockPub) DeleteG() error {
	if o == nil {
		return errors.New("models: no StockPub provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single StockPub record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockPub) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single StockPub record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockPub) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockPub provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockPubPrimaryKeyMapping)
	sql := "DELETE FROM \"stock_pub\" WHERE \"stock_pub_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from stock_pub")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockPubQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockPubQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no stockPubQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stock_pub")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockPubSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockPubSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no StockPub slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockPubSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockPubSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockPub slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockPubBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stock_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockPubPrimaryKeyColumns), 1, len(stockPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockPub slice")
	}

	if len(stockPubAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *StockPub) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *StockPub) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockPub) ReloadG() error {
	if o == nil {
		return errors.New("models: no StockPub provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockPub) Reload(exec boil.Executor) error {
	ret, err := FindStockPub(exec, o.StockPubID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockPubSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockPubSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockPubSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty StockPubSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockPubSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockPubs := StockPubSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stock_pub\".* FROM \"stock_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockPubPrimaryKeyColumns), 1, len(stockPubPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockPubs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockPubSlice")
	}

	*o = stockPubs

	return nil
}

// StockPubExists checks if the StockPub row exists.
func StockPubExists(exec boil.Executor, stockPubID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stock_pub\" where \"stock_pub_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockPubID)
	}

	row := exec.QueryRow(sql, stockPubID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stock_pub exists")
	}

	return exists, nil
}

// StockPubExistsG checks if the StockPub row exists.
func StockPubExistsG(stockPubID int) (bool, error) {
	return StockPubExists(boil.GetDB(), stockPubID)
}

// StockPubExistsGP checks if the StockPub row exists. Panics on error.
func StockPubExistsGP(stockPubID int) bool {
	e, err := StockPubExists(boil.GetDB(), stockPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockPubExistsP checks if the StockPub row exists. Panics on error.
func StockPubExistsP(exec boil.Executor, stockPubID int) bool {
	e, err := StockPubExists(exec, stockPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
