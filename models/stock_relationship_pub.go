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

// StockRelationshipPub is an object representing the database table.
type StockRelationshipPub struct {
	StockRelationshipPubID int `boil:"stock_relationship_pub_id" json:"stock_relationship_pub_id" toml:"stock_relationship_pub_id" yaml:"stock_relationship_pub_id"`
	StockRelationshipID    int `boil:"stock_relationship_id" json:"stock_relationship_id" toml:"stock_relationship_id" yaml:"stock_relationship_id"`
	PubID                  int `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`

	R *stockRelationshipPubR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockRelationshipPubL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockRelationshipPubR is where relationships are stored.
type stockRelationshipPubR struct {
	Pub               *Pub
	StockRelationship *StockRelationship
}

// stockRelationshipPubL is where Load methods for each relationship are stored.
type stockRelationshipPubL struct{}

var (
	stockRelationshipPubColumns               = []string{"stock_relationship_pub_id", "stock_relationship_id", "pub_id"}
	stockRelationshipPubColumnsWithoutDefault = []string{"stock_relationship_id", "pub_id"}
	stockRelationshipPubColumnsWithDefault    = []string{"stock_relationship_pub_id"}
	stockRelationshipPubPrimaryKeyColumns     = []string{"stock_relationship_pub_id"}
)

type (
	// StockRelationshipPubSlice is an alias for a slice of pointers to StockRelationshipPub.
	// This should generally be used opposed to []StockRelationshipPub.
	StockRelationshipPubSlice []*StockRelationshipPub
	// StockRelationshipPubHook is the signature for custom StockRelationshipPub hook methods
	StockRelationshipPubHook func(boil.Executor, *StockRelationshipPub) error

	stockRelationshipPubQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockRelationshipPubType                 = reflect.TypeOf(&StockRelationshipPub{})
	stockRelationshipPubMapping              = queries.MakeStructMapping(stockRelationshipPubType)
	stockRelationshipPubPrimaryKeyMapping, _ = queries.BindMapping(stockRelationshipPubType, stockRelationshipPubMapping, stockRelationshipPubPrimaryKeyColumns)
	stockRelationshipPubInsertCacheMut       sync.RWMutex
	stockRelationshipPubInsertCache          = make(map[string]insertCache)
	stockRelationshipPubUpdateCacheMut       sync.RWMutex
	stockRelationshipPubUpdateCache          = make(map[string]updateCache)
	stockRelationshipPubUpsertCacheMut       sync.RWMutex
	stockRelationshipPubUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockRelationshipPubBeforeInsertHooks []StockRelationshipPubHook
var stockRelationshipPubBeforeUpdateHooks []StockRelationshipPubHook
var stockRelationshipPubBeforeDeleteHooks []StockRelationshipPubHook
var stockRelationshipPubBeforeUpsertHooks []StockRelationshipPubHook

var stockRelationshipPubAfterInsertHooks []StockRelationshipPubHook
var stockRelationshipPubAfterSelectHooks []StockRelationshipPubHook
var stockRelationshipPubAfterUpdateHooks []StockRelationshipPubHook
var stockRelationshipPubAfterDeleteHooks []StockRelationshipPubHook
var stockRelationshipPubAfterUpsertHooks []StockRelationshipPubHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockRelationshipPub) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipPubBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockRelationshipPub) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipPubBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockRelationshipPub) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipPubBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockRelationshipPub) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipPubBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockRelationshipPub) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipPubAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockRelationshipPub) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipPubAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockRelationshipPub) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipPubAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockRelationshipPub) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipPubAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockRelationshipPub) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockRelationshipPubAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockRelationshipPubHook registers your hook function for all future operations.
func AddStockRelationshipPubHook(hookPoint boil.HookPoint, stockRelationshipPubHook StockRelationshipPubHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockRelationshipPubBeforeInsertHooks = append(stockRelationshipPubBeforeInsertHooks, stockRelationshipPubHook)
	case boil.BeforeUpdateHook:
		stockRelationshipPubBeforeUpdateHooks = append(stockRelationshipPubBeforeUpdateHooks, stockRelationshipPubHook)
	case boil.BeforeDeleteHook:
		stockRelationshipPubBeforeDeleteHooks = append(stockRelationshipPubBeforeDeleteHooks, stockRelationshipPubHook)
	case boil.BeforeUpsertHook:
		stockRelationshipPubBeforeUpsertHooks = append(stockRelationshipPubBeforeUpsertHooks, stockRelationshipPubHook)
	case boil.AfterInsertHook:
		stockRelationshipPubAfterInsertHooks = append(stockRelationshipPubAfterInsertHooks, stockRelationshipPubHook)
	case boil.AfterSelectHook:
		stockRelationshipPubAfterSelectHooks = append(stockRelationshipPubAfterSelectHooks, stockRelationshipPubHook)
	case boil.AfterUpdateHook:
		stockRelationshipPubAfterUpdateHooks = append(stockRelationshipPubAfterUpdateHooks, stockRelationshipPubHook)
	case boil.AfterDeleteHook:
		stockRelationshipPubAfterDeleteHooks = append(stockRelationshipPubAfterDeleteHooks, stockRelationshipPubHook)
	case boil.AfterUpsertHook:
		stockRelationshipPubAfterUpsertHooks = append(stockRelationshipPubAfterUpsertHooks, stockRelationshipPubHook)
	}
}

// OneP returns a single stockRelationshipPub record from the query, and panics on error.
func (q stockRelationshipPubQuery) OneP() *StockRelationshipPub {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockRelationshipPub record from the query.
func (q stockRelationshipPubQuery) One() (*StockRelationshipPub, error) {
	o := &StockRelationshipPub{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stock_relationship_pub")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all StockRelationshipPub records from the query, and panics on error.
func (q stockRelationshipPubQuery) AllP() StockRelationshipPubSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all StockRelationshipPub records from the query.
func (q stockRelationshipPubQuery) All() (StockRelationshipPubSlice, error) {
	var o StockRelationshipPubSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StockRelationshipPub slice")
	}

	if len(stockRelationshipPubAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all StockRelationshipPub records in the query, and panics on error.
func (q stockRelationshipPubQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all StockRelationshipPub records in the query.
func (q stockRelationshipPubQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stock_relationship_pub rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockRelationshipPubQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockRelationshipPubQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stock_relationship_pub exists")
	}

	return count > 0, nil
}

// PubG pointed to by the foreign key.
func (o *StockRelationshipPub) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *StockRelationshipPub) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// StockRelationshipG pointed to by the foreign key.
func (o *StockRelationshipPub) StockRelationshipG(mods ...qm.QueryMod) stockRelationshipQuery {
	return o.StockRelationship(boil.GetDB(), mods...)
}

// StockRelationship pointed to by the foreign key.
func (o *StockRelationshipPub) StockRelationship(exec boil.Executor, mods ...qm.QueryMod) stockRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_relationship_id=$1", o.StockRelationshipID),
	}

	queryMods = append(queryMods, mods...)

	query := StockRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_relationship\"")

	return query
}

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockRelationshipPubL) LoadPub(e boil.Executor, singular bool, maybeStockRelationshipPub interface{}) error {
	var slice []*StockRelationshipPub
	var object *StockRelationshipPub

	count := 1
	if singular {
		object = maybeStockRelationshipPub.(*StockRelationshipPub)
	} else {
		slice = *maybeStockRelationshipPub.(*StockRelationshipPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockRelationshipPubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &stockRelationshipPubR{}
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

	if len(stockRelationshipPubAfterSelectHooks) != 0 {
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

// LoadStockRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockRelationshipPubL) LoadStockRelationship(e boil.Executor, singular bool, maybeStockRelationshipPub interface{}) error {
	var slice []*StockRelationshipPub
	var object *StockRelationshipPub

	count := 1
	if singular {
		object = maybeStockRelationshipPub.(*StockRelationshipPub)
	} else {
		slice = *maybeStockRelationshipPub.(*StockRelationshipPubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockRelationshipPubR{}
		args[0] = object.StockRelationshipID
	} else {
		for i, obj := range slice {
			obj.R = &stockRelationshipPubR{}
			args[i] = obj.StockRelationshipID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_relationship\" where \"stock_relationship_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockRelationship")
	}
	defer results.Close()

	var resultSlice []*StockRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockRelationship")
	}

	if len(stockRelationshipPubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.StockRelationshipID == foreign.StockRelationshipID {
				local.R.StockRelationship = foreign
				break
			}
		}
	}

	return nil
}

// SetPub of the stock_relationship_pub to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.StockRelationshipPub.
func (o *StockRelationshipPub) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_relationship_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockRelationshipPubPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.StockRelationshipPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &stockRelationshipPubR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			StockRelationshipPub: o,
		}
	} else {
		related.R.StockRelationshipPub = o
	}

	return nil
}

// SetStockRelationship of the stock_relationship_pub to the related item.
// Sets o.R.StockRelationship to related.
// Adds o to related.R.StockRelationshipPub.
func (o *StockRelationshipPub) SetStockRelationship(exec boil.Executor, insert bool, related *StockRelationship) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_relationship_pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stock_relationship_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockRelationshipPubPrimaryKeyColumns),
	)
	values := []interface{}{related.StockRelationshipID, o.StockRelationshipPubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StockRelationshipID = related.StockRelationshipID

	if o.R == nil {
		o.R = &stockRelationshipPubR{
			StockRelationship: related,
		}
	} else {
		o.R.StockRelationship = related
	}

	if related.R == nil {
		related.R = &stockRelationshipR{
			StockRelationshipPub: o,
		}
	} else {
		related.R.StockRelationshipPub = o
	}

	return nil
}

// StockRelationshipPubsG retrieves all records.
func StockRelationshipPubsG(mods ...qm.QueryMod) stockRelationshipPubQuery {
	return StockRelationshipPubs(boil.GetDB(), mods...)
}

// StockRelationshipPubs retrieves all the records using an executor.
func StockRelationshipPubs(exec boil.Executor, mods ...qm.QueryMod) stockRelationshipPubQuery {
	mods = append(mods, qm.From("\"stock_relationship_pub\""))
	return stockRelationshipPubQuery{NewQuery(exec, mods...)}
}

// FindStockRelationshipPubG retrieves a single record by ID.
func FindStockRelationshipPubG(stockRelationshipPubID int, selectCols ...string) (*StockRelationshipPub, error) {
	return FindStockRelationshipPub(boil.GetDB(), stockRelationshipPubID, selectCols...)
}

// FindStockRelationshipPubGP retrieves a single record by ID, and panics on error.
func FindStockRelationshipPubGP(stockRelationshipPubID int, selectCols ...string) *StockRelationshipPub {
	retobj, err := FindStockRelationshipPub(boil.GetDB(), stockRelationshipPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockRelationshipPub retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockRelationshipPub(exec boil.Executor, stockRelationshipPubID int, selectCols ...string) (*StockRelationshipPub, error) {
	stockRelationshipPubObj := &StockRelationshipPub{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stock_relationship_pub\" where \"stock_relationship_pub_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockRelationshipPubID)

	err := q.Bind(stockRelationshipPubObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stock_relationship_pub")
	}

	return stockRelationshipPubObj, nil
}

// FindStockRelationshipPubP retrieves a single record by ID with an executor, and panics on error.
func FindStockRelationshipPubP(exec boil.Executor, stockRelationshipPubID int, selectCols ...string) *StockRelationshipPub {
	retobj, err := FindStockRelationshipPub(exec, stockRelationshipPubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockRelationshipPub) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *StockRelationshipPub) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *StockRelationshipPub) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *StockRelationshipPub) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_relationship_pub provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockRelationshipPubColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockRelationshipPubInsertCacheMut.RLock()
	cache, cached := stockRelationshipPubInsertCache[key]
	stockRelationshipPubInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockRelationshipPubColumns,
			stockRelationshipPubColumnsWithDefault,
			stockRelationshipPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockRelationshipPubType, stockRelationshipPubMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockRelationshipPubType, stockRelationshipPubMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stock_relationship_pub\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into stock_relationship_pub")
	}

	if !cached {
		stockRelationshipPubInsertCacheMut.Lock()
		stockRelationshipPubInsertCache[key] = cache
		stockRelationshipPubInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single StockRelationshipPub record. See Update for
// whitelist behavior description.
func (o *StockRelationshipPub) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single StockRelationshipPub record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *StockRelationshipPub) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the StockRelationshipPub, and panics on error.
// See Update for whitelist behavior description.
func (o *StockRelationshipPub) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the StockRelationshipPub.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *StockRelationshipPub) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockRelationshipPubUpdateCacheMut.RLock()
	cache, cached := stockRelationshipPubUpdateCache[key]
	stockRelationshipPubUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockRelationshipPubColumns, stockRelationshipPubPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update stock_relationship_pub, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stock_relationship_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockRelationshipPubPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockRelationshipPubType, stockRelationshipPubMapping, append(wl, stockRelationshipPubPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update stock_relationship_pub row")
	}

	if !cached {
		stockRelationshipPubUpdateCacheMut.Lock()
		stockRelationshipPubUpdateCache[key] = cache
		stockRelationshipPubUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockRelationshipPubQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockRelationshipPubQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for stock_relationship_pub")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockRelationshipPubSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockRelationshipPubSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockRelationshipPubSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockRelationshipPubSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockRelationshipPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stock_relationship_pub\" SET %s WHERE (\"stock_relationship_pub_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockRelationshipPubPrimaryKeyColumns), len(colNames)+1, len(stockRelationshipPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in stockRelationshipPub slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockRelationshipPub) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *StockRelationshipPub) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *StockRelationshipPub) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *StockRelationshipPub) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stock_relationship_pub provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockRelationshipPubColumnsWithDefault, o)

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

	stockRelationshipPubUpsertCacheMut.RLock()
	cache, cached := stockRelationshipPubUpsertCache[key]
	stockRelationshipPubUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockRelationshipPubColumns,
			stockRelationshipPubColumnsWithDefault,
			stockRelationshipPubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockRelationshipPubColumns,
			stockRelationshipPubPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert stock_relationship_pub, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockRelationshipPubPrimaryKeyColumns))
			copy(conflict, stockRelationshipPubPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stock_relationship_pub\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockRelationshipPubType, stockRelationshipPubMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockRelationshipPubType, stockRelationshipPubMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for stock_relationship_pub")
	}

	if !cached {
		stockRelationshipPubUpsertCacheMut.Lock()
		stockRelationshipPubUpsertCache[key] = cache
		stockRelationshipPubUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single StockRelationshipPub record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockRelationshipPub) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single StockRelationshipPub record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockRelationshipPub) DeleteG() error {
	if o == nil {
		return errors.New("models: no StockRelationshipPub provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single StockRelationshipPub record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockRelationshipPub) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single StockRelationshipPub record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockRelationshipPub) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockRelationshipPub provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockRelationshipPubPrimaryKeyMapping)
	sql := "DELETE FROM \"stock_relationship_pub\" WHERE \"stock_relationship_pub_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from stock_relationship_pub")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockRelationshipPubQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockRelationshipPubQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no stockRelationshipPubQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stock_relationship_pub")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockRelationshipPubSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockRelationshipPubSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no StockRelationshipPub slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockRelationshipPubSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockRelationshipPubSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no StockRelationshipPub slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockRelationshipPubBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockRelationshipPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stock_relationship_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockRelationshipPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockRelationshipPubPrimaryKeyColumns), 1, len(stockRelationshipPubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockRelationshipPub slice")
	}

	if len(stockRelationshipPubAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *StockRelationshipPub) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *StockRelationshipPub) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockRelationshipPub) ReloadG() error {
	if o == nil {
		return errors.New("models: no StockRelationshipPub provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockRelationshipPub) Reload(exec boil.Executor) error {
	ret, err := FindStockRelationshipPub(exec, o.StockRelationshipPubID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockRelationshipPubSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockRelationshipPubSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockRelationshipPubSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty StockRelationshipPubSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockRelationshipPubSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockRelationshipPubs := StockRelationshipPubSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockRelationshipPubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stock_relationship_pub\".* FROM \"stock_relationship_pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockRelationshipPubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockRelationshipPubPrimaryKeyColumns), 1, len(stockRelationshipPubPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockRelationshipPubs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockRelationshipPubSlice")
	}

	*o = stockRelationshipPubs

	return nil
}

// StockRelationshipPubExists checks if the StockRelationshipPub row exists.
func StockRelationshipPubExists(exec boil.Executor, stockRelationshipPubID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stock_relationship_pub\" where \"stock_relationship_pub_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockRelationshipPubID)
	}

	row := exec.QueryRow(sql, stockRelationshipPubID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stock_relationship_pub exists")
	}

	return exists, nil
}

// StockRelationshipPubExistsG checks if the StockRelationshipPub row exists.
func StockRelationshipPubExistsG(stockRelationshipPubID int) (bool, error) {
	return StockRelationshipPubExists(boil.GetDB(), stockRelationshipPubID)
}

// StockRelationshipPubExistsGP checks if the StockRelationshipPub row exists. Panics on error.
func StockRelationshipPubExistsGP(stockRelationshipPubID int) bool {
	e, err := StockRelationshipPubExists(boil.GetDB(), stockRelationshipPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockRelationshipPubExistsP checks if the StockRelationshipPub row exists. Panics on error.
func StockRelationshipPubExistsP(exec boil.Executor, stockRelationshipPubID int) bool {
	e, err := StockRelationshipPubExists(exec, stockRelationshipPubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
