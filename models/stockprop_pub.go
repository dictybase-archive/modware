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

// StockpropPub is an object representing the database table.
type StockpropPub struct {
	StockpropPubID int `boil:"stockprop_pub_id" json:"stockprop_pub_id" toml:"stockprop_pub_id" yaml:"stockprop_pub_id"`
	StockpropID    int `boil:"stockprop_id" json:"stockprop_id" toml:"stockprop_id" yaml:"stockprop_id"`
	PubID          int `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`

	R *stockpropPubR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockpropPubL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockpropPubR is where relationships are stored.
type stockpropPubR struct {
	Pub       *Pub
	Stockprop *Stockprop
}

// stockpropPubL is where Load methods for each relationship are stored.
type stockpropPubL struct{}

var (
	stockpropPubColumns               = []string{"stockprop_pub_id", "stockprop_id", "pub_id"}
	stockpropPubColumnsWithoutDefault = []string{"stockprop_id", "pub_id"}
	stockpropPubColumnsWithDefault    = []string{"stockprop_pub_id"}
	stockpropPubPrimaryKeyColumns     = []string{"stockprop_pub_id"}
)

type (
	// StockpropPubSlice is an alias for a slice of pointers to StockpropPub.
	// This should generally be used opposed to []StockpropPub.
	StockpropPubSlice []*StockpropPub
	// StockpropPubHook is the signature for custom StockpropPub hook methods
	StockpropPubHook func(boil.Executor, *StockpropPub) error

	stockpropPubQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockpropPubType                 = reflect.TypeOf(&StockpropPub{})
	stockpropPubMapping              = queries.MakeStructMapping(stockpropPubType)
	stockpropPubPrimaryKeyMapping, _ = queries.BindMapping(stockpropPubType, stockpropPubMapping, stockpropPubPrimaryKeyColumns)
	stockpropPubInsertCacheMut       sync.RWMutex
	stockpropPubInsertCache          = make(map[string]insertCache)
	stockpropPubUpdateCacheMut       sync.RWMutex
	stockpropPubUpdateCache          = make(map[string]updateCache)
	stockpropPubUpsertCacheMut       sync.RWMutex
	stockpropPubUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockpropPubBeforeInsertHooks []StockpropPubHook
var stockpropPubBeforeUpdateHooks []StockpropPubHook
var stockpropPubBeforeDeleteHooks []StockpropPubHook
var stockpropPubBeforeUpsertHooks []StockpropPubHook

var stockpropPubAfterInsertHooks []StockpropPubHook
var stockpropPubAfterSelectHooks []StockpropPubHook
var stockpropPubAfterUpdateHooks []StockpropPubHook
var stockpropPubAfterDeleteHooks []StockpropPubHook
var stockpropPubAfterUpsertHooks []StockpropPubHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockpropPub) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropPubBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockpropPub) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropPubBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockpropPub) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropPubBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockpropPub) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropPubBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockpropPub) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropPubAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockpropPub) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropPubAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockpropPub) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropPubAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockpropPub) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropPubAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockpropPub) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockpropPubAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockpropPubHook registers your hook function for all future operations.
func AddStockpropPubHook(hookPoint boil.HookPoint, stockpropPubHook StockpropPubHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockpropPubBeforeInsertHooks = append(stockpropPubBeforeInsertHooks, stockpropPubHook)
	case boil.BeforeUpdateHook:
		stockpropPubBeforeUpdateHooks = append(stockpropPubBeforeUpdateHooks, stockpropPubHook)
	case boil.BeforeDeleteHook:
		stockpropPubBeforeDeleteHooks = append(stockpropPubBeforeDeleteHooks, stockpropPubHook)
	case boil.BeforeUpsertHook:
		stockpropPubBeforeUpsertHooks = append(stockpropPubBeforeUpsertHooks, stockpropPubHook)
	case boil.AfterInsertHook:
		stockpropPubAfterInsertHooks = append(stockpropPubAfterInsertHooks, stockpropPubHook)
	case boil.AfterSelectHook:
		stockpropPubAfterSelectHooks = append(stockpropPubAfterSelectHooks, stockpropPubHook)
	case boil.AfterUpdateHook:
		stockpropPubAfterUpdateHooks = append(stockpropPubAfterUpdateHooks, stockpropPubHook)
	case boil.AfterDeleteHook:
		stockpropPubAfterDeleteHooks = append(stockpropPubAfterDeleteHooks, stockpropPubHook)
	case boil.AfterUpsertHook:
		stockpropPubAfterUpsertHooks = append(stockpropPubAfterUpsertHooks, stockpropPubHook)
	}
}

// OneP returns a single stockpropPub record from the query, and panics on error.
func (q stockpropPubQuery) OneP() *StockpropPub {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockpropPub record from the query.
func (q stockpropPubQuery) One() (*StockpropPub, error) {
	o := &StockpropPub{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stockprop_pub")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all StockpropPub records from the query, and panics on error.
func (q stockpropPubQuery) AllP() StockpropPubSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all StockpropPub records from the query.
func (q stockpropPubQuery) All() (StockpropPubSlice, error) {
	var o StockpropPubSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StockpropPub slice")
	}

	if len(stockpropPubAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all StockpropPub records in the query, and panics on error.
func (q stockpropPubQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all StockpropPub records in the query.
func (q stockpropPubQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stockprop_pub rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockpropPubQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockpropPubQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stockprop_pub exists")
	}

	return count > 0, nil
}

// PubG pointed to by the foreign key.
func (o *StockpropPub) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *StockpropPub) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// StockpropG pointed to by the foreign key.
func (o *StockpropPub) StockpropG(mods ...qm.QueryMod) stockpropQuery {
	return o.Stockprop(boil.GetDB(), mods...)
}

// Stockprop pointed to by the foreign key.
func (o *StockpropPub) Stockprop(exec boil.Executor, mods ...qm.QueryMod) stockpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stockprop_id=$1", o.StockpropID),
	}

	queryMods = append(queryMods, mods...)

	query := Stockprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stockprop\"")

	return query
}

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockpropPubL) LoadPub(e boil.Executor, singular bool, maybeStockpropPub interface{}) error {
	var slice []*StockpropPub
	var object *StockpropPub

	count := 1
	if singular {
		object = maybeStockpropPub.(*StockpropPub)
	} else {
		slice = *maybeStockpropPub.(*StockpropPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockpropPubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &stockpropPubR{}
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

	if len(stockpropPubAfterSelectHooks) != 0 {
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

// LoadStockprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockpropPubL) LoadStockprop(e boil.Executor, singular bool, maybeStockpropPub interface{}) error {
	var slice []*StockpropPub
	var object *StockpropPub

	count := 1
	if singular {
		object = maybeStockpropPub.(*StockpropPub)
	} else {
		slice = *maybeStockpropPub.(*StockpropPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockpropPubR{}
		args[0] = object.StockpropID
	} else {
		for i, obj := range slice {
			obj.R = &stockpropPubR{}
			args[i] = obj.StockpropID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stockprop\" where \"stockprop_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Stockprop")
	}
	defer results.Close()

	var resultSlice []*Stockprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Stockprop")
	}

	if len(stockpropPubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Stockprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockpropID == foreign.StockpropID {
				local.R.Stockprop = foreign
				break
			}
		}
	}

	return nil
}

// SetPub of the stockprop_pub to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.StockpropPub.
func (o *StockpropPub) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stockprop_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockpropPubPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.StockpropPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &stockpropPubR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			StockpropPub: o,
		}
	} else {
		related.R.StockpropPub = o
	}

	return nil
}

// SetStockprop of the stockprop_pub to the related item.
// Sets o.R.Stockprop to related.
// Adds o to related.R.StockpropPub.
func (o *StockpropPub) SetStockprop(exec boil.Executor, insert bool, related *Stockprop) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stockprop_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stockprop_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockpropPubPrimaryKeyColumns),
	)
	values := []interface{}{related.StockpropID, o.StockpropPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StockpropID = related.StockpropID

	if o.R == nil {
		o.R = &stockpropPubR{
			Stockprop: related,
		}
	} else {
		o.R.Stockprop = related
	}

	if related.R == nil {
		related.R = &stockpropR{
			StockpropPub: o,
		}
	} else {
		related.R.StockpropPub = o
	}

	return nil
}

// StockpropPubsG retrieves all records.
func StockpropPubsG(mods ...qm.QueryMod) stockpropPubQuery {
	return StockpropPubs(boil.GetDB(), mods...)
}

// StockpropPubs retrieves all the records using an executor.
func StockpropPubs(exec boil.Executor, mods ...qm.QueryMod) stockpropPubQuery {
	mods = append(mods, qm.From("\"stockprop_pub\""))
	return stockpropPubQuery{NewQuery(exec, mods...)}
}

// FindStockpropPubG retrieves a single record by ID.
func FindStockpropPubG(stockpropPubID int, selectCols ...string) (*StockpropPub, error) {
	return FindStockpropPub(boil.GetDB(), stockpropPubID, selectCols...)
}

// FindStockpropPubGP retrieves a single record by ID, and panics on error.
func FindStockpropPubGP(stockpropPubID int, selectCols ...string) *StockpropPub {
	retobj, err := FindStockpropPub(boil.GetDB(), stockpropPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockpropPub retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockpropPub(exec boil.Executor, stockpropPubID int, selectCols ...string) (*StockpropPub, error) {
	stockpropPubObj := &StockpropPub{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stockprop_pub\" where \"stockprop_pub_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockpropPubID)

	err := q.Bind(stockpropPubObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stockprop_pub")
	}

	return stockpropPubObj, nil
}

// FindStockpropPubP retrieves a single record by ID with an executor, and panics on error.
func FindStockpropPubP(exec boil.Executor, stockpropPubID int, selectCols ...string) *StockpropPub {
	retobj, err := FindStockpropPub(exec, stockpropPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockpropPub) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *StockpropPub) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *StockpropPub) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *StockpropPub) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stockprop_pub provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockpropPubColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockpropPubInsertCacheMut.RLock()
	cache, cached := stockpropPubInsertCache[key]
	stockpropPubInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockpropPubColumns,
			stockpropPubColumnsWithDefault,
			stockpropPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockpropPubType, stockpropPubMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockpropPubType, stockpropPubMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stockprop_pub\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into stockprop_pub")
	}

	if !cached {
		stockpropPubInsertCacheMut.Lock()
		stockpropPubInsertCache[key] = cache
		stockpropPubInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single StockpropPub record. See Update for
// whitelist behavior description.
func (o *StockpropPub) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single StockpropPub record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *StockpropPub) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the StockpropPub, and panics on error.
// See Update for whitelist behavior description.
func (o *StockpropPub) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the StockpropPub.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *StockpropPub) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockpropPubUpdateCacheMut.RLock()
	cache, cached := stockpropPubUpdateCache[key]
	stockpropPubUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockpropPubColumns, stockpropPubPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update stockprop_pub, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stockprop_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockpropPubPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockpropPubType, stockpropPubMapping, append(wl, stockpropPubPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update stockprop_pub row")
	}

	if !cached {
		stockpropPubUpdateCacheMut.Lock()
		stockpropPubUpdateCache[key] = cache
		stockpropPubUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockpropPubQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockpropPubQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for stockprop_pub")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockpropPubSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockpropPubSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockpropPubSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockpropPubSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockpropPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stockprop_pub\" SET %s WHERE (\"stockprop_pub_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockpropPubPrimaryKeyColumns), len(colNames)+1, len(stockpropPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in stockpropPub slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockpropPub) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *StockpropPub) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *StockpropPub) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *StockpropPub) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stockprop_pub provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockpropPubColumnsWithDefault, o)

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

	stockpropPubUpsertCacheMut.RLock()
	cache, cached := stockpropPubUpsertCache[key]
	stockpropPubUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockpropPubColumns,
			stockpropPubColumnsWithDefault,
			stockpropPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockpropPubColumns,
			stockpropPubPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert stockprop_pub, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockpropPubPrimaryKeyColumns))
			copy(conflict, stockpropPubPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stockprop_pub\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockpropPubType, stockpropPubMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockpropPubType, stockpropPubMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for stockprop_pub")
	}

	if !cached {
		stockpropPubUpsertCacheMut.Lock()
		stockpropPubUpsertCache[key] = cache
		stockpropPubUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single StockpropPub record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockpropPub) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single StockpropPub record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockpropPub) DeleteG() error {
	if o == nil {
		return errors.New("models: no StockpropPub provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single StockpropPub record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockpropPub) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single StockpropPub record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockpropPub) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockpropPub provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockpropPubPrimaryKeyMapping)
	sql := "DELETE FROM \"stockprop_pub\" WHERE \"stockprop_pub_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from stockprop_pub")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockpropPubQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockpropPubQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no stockpropPubQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockprop_pub")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockpropPubSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockpropPubSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no StockpropPub slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockpropPubSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockpropPubSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockpropPub slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockpropPubBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockpropPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stockprop_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockpropPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockpropPubPrimaryKeyColumns), 1, len(stockpropPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockpropPub slice")
	}

	if len(stockpropPubAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *StockpropPub) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *StockpropPub) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockpropPub) ReloadG() error {
	if o == nil {
		return errors.New("models: no StockpropPub provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockpropPub) Reload(exec boil.Executor) error {
	ret, err := FindStockpropPub(exec, o.StockpropPubID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockpropPubSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockpropPubSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockpropPubSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty StockpropPubSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockpropPubSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockpropPubs := StockpropPubSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockpropPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stockprop_pub\".* FROM \"stockprop_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockpropPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockpropPubPrimaryKeyColumns), 1, len(stockpropPubPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockpropPubs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockpropPubSlice")
	}

	*o = stockpropPubs

	return nil
}

// StockpropPubExists checks if the StockpropPub row exists.
func StockpropPubExists(exec boil.Executor, stockpropPubID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stockprop_pub\" where \"stockprop_pub_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockpropPubID)
	}

	row := exec.QueryRow(sql, stockpropPubID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stockprop_pub exists")
	}

	return exists, nil
}

// StockpropPubExistsG checks if the StockpropPub row exists.
func StockpropPubExistsG(stockpropPubID int) (bool, error) {
	return StockpropPubExists(boil.GetDB(), stockpropPubID)
}

// StockpropPubExistsGP checks if the StockpropPub row exists. Panics on error.
func StockpropPubExistsGP(stockpropPubID int) bool {
	e, err := StockpropPubExists(boil.GetDB(), stockpropPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockpropPubExistsP checks if the StockpropPub row exists. Panics on error.
func StockpropPubExistsP(exec boil.Executor, stockpropPubID int) bool {
	e, err := StockpropPubExists(exec, stockpropPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
