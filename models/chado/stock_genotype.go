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
)

// StockGenotype is an object representing the database table.
type StockGenotype struct {
	StockGenotypeID int `boil:"stock_genotype_id" json:"stock_genotype_id" toml:"stock_genotype_id" yaml:"stock_genotype_id"`
	StockID         int `boil:"stock_id" json:"stock_id" toml:"stock_id" yaml:"stock_id"`
	GenotypeID      int `boil:"genotype_id" json:"genotype_id" toml:"genotype_id" yaml:"genotype_id"`

	R *stockGenotypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockGenotypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockGenotypeR is where relationships are stored.
type stockGenotypeR struct {
	Genotype *Genotype
	Stock    *Stock
}

// stockGenotypeL is where Load methods for each relationship are stored.
type stockGenotypeL struct{}

var (
	stockGenotypeColumns               = []string{"stock_genotype_id", "stock_id", "genotype_id"}
	stockGenotypeColumnsWithoutDefault = []string{"stock_id", "genotype_id"}
	stockGenotypeColumnsWithDefault    = []string{"stock_genotype_id"}
	stockGenotypePrimaryKeyColumns     = []string{"stock_genotype_id"}
)

type (
	// StockGenotypeSlice is an alias for a slice of pointers to StockGenotype.
	// This should generally be used opposed to []StockGenotype.
	StockGenotypeSlice []*StockGenotype
	// StockGenotypeHook is the signature for custom StockGenotype hook methods
	StockGenotypeHook func(boil.Executor, *StockGenotype) error

	stockGenotypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockGenotypeType                 = reflect.TypeOf(&StockGenotype{})
	stockGenotypeMapping              = queries.MakeStructMapping(stockGenotypeType)
	stockGenotypePrimaryKeyMapping, _ = queries.BindMapping(stockGenotypeType, stockGenotypeMapping, stockGenotypePrimaryKeyColumns)
	stockGenotypeInsertCacheMut       sync.RWMutex
	stockGenotypeInsertCache          = make(map[string]insertCache)
	stockGenotypeUpdateCacheMut       sync.RWMutex
	stockGenotypeUpdateCache          = make(map[string]updateCache)
	stockGenotypeUpsertCacheMut       sync.RWMutex
	stockGenotypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockGenotypeBeforeInsertHooks []StockGenotypeHook
var stockGenotypeBeforeUpdateHooks []StockGenotypeHook
var stockGenotypeBeforeDeleteHooks []StockGenotypeHook
var stockGenotypeBeforeUpsertHooks []StockGenotypeHook

var stockGenotypeAfterInsertHooks []StockGenotypeHook
var stockGenotypeAfterSelectHooks []StockGenotypeHook
var stockGenotypeAfterUpdateHooks []StockGenotypeHook
var stockGenotypeAfterDeleteHooks []StockGenotypeHook
var stockGenotypeAfterUpsertHooks []StockGenotypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StockGenotype) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockGenotypeBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StockGenotype) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockGenotypeBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StockGenotype) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockGenotypeBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StockGenotype) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockGenotypeBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StockGenotype) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockGenotypeAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StockGenotype) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockGenotypeAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StockGenotype) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockGenotypeAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StockGenotype) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockGenotypeAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StockGenotype) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockGenotypeAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockGenotypeHook registers your hook function for all future operations.
func AddStockGenotypeHook(hookPoint boil.HookPoint, stockGenotypeHook StockGenotypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockGenotypeBeforeInsertHooks = append(stockGenotypeBeforeInsertHooks, stockGenotypeHook)
	case boil.BeforeUpdateHook:
		stockGenotypeBeforeUpdateHooks = append(stockGenotypeBeforeUpdateHooks, stockGenotypeHook)
	case boil.BeforeDeleteHook:
		stockGenotypeBeforeDeleteHooks = append(stockGenotypeBeforeDeleteHooks, stockGenotypeHook)
	case boil.BeforeUpsertHook:
		stockGenotypeBeforeUpsertHooks = append(stockGenotypeBeforeUpsertHooks, stockGenotypeHook)
	case boil.AfterInsertHook:
		stockGenotypeAfterInsertHooks = append(stockGenotypeAfterInsertHooks, stockGenotypeHook)
	case boil.AfterSelectHook:
		stockGenotypeAfterSelectHooks = append(stockGenotypeAfterSelectHooks, stockGenotypeHook)
	case boil.AfterUpdateHook:
		stockGenotypeAfterUpdateHooks = append(stockGenotypeAfterUpdateHooks, stockGenotypeHook)
	case boil.AfterDeleteHook:
		stockGenotypeAfterDeleteHooks = append(stockGenotypeAfterDeleteHooks, stockGenotypeHook)
	case boil.AfterUpsertHook:
		stockGenotypeAfterUpsertHooks = append(stockGenotypeAfterUpsertHooks, stockGenotypeHook)
	}
}

// OneP returns a single stockGenotype record from the query, and panics on error.
func (q stockGenotypeQuery) OneP() *StockGenotype {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockGenotype record from the query.
func (q stockGenotypeQuery) One() (*StockGenotype, error) {
	o := &StockGenotype{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for stock_genotype")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all StockGenotype records from the query, and panics on error.
func (q stockGenotypeQuery) AllP() StockGenotypeSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all StockGenotype records from the query.
func (q stockGenotypeQuery) All() (StockGenotypeSlice, error) {
	var o StockGenotypeSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to StockGenotype slice")
	}

	if len(stockGenotypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all StockGenotype records in the query, and panics on error.
func (q stockGenotypeQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all StockGenotype records in the query.
func (q stockGenotypeQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count stock_genotype rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockGenotypeQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockGenotypeQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if stock_genotype exists")
	}

	return count > 0, nil
}

// GenotypeG pointed to by the foreign key.
func (o *StockGenotype) GenotypeG(mods ...qm.QueryMod) genotypeQuery {
	return o.Genotype(boil.GetDB(), mods...)
}

// Genotype pointed to by the foreign key.
func (o *StockGenotype) Genotype(exec boil.Executor, mods ...qm.QueryMod) genotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype_id=$1", o.GenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Genotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"genotype\"")

	return query
}

// StockG pointed to by the foreign key.
func (o *StockGenotype) StockG(mods ...qm.QueryMod) stockQuery {
	return o.Stock(boil.GetDB(), mods...)
}

// Stock pointed to by the foreign key.
func (o *StockGenotype) Stock(exec boil.Executor, mods ...qm.QueryMod) stockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stock_id=$1", o.StockID),
	}

	queryMods = append(queryMods, mods...)

	query := Stocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock\"")

	return query
}

// LoadGenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockGenotypeL) LoadGenotype(e boil.Executor, singular bool, maybeStockGenotype interface{}) error {
	var slice []*StockGenotype
	var object *StockGenotype

	count := 1
	if singular {
		object = maybeStockGenotype.(*StockGenotype)
	} else {
		slice = *maybeStockGenotype.(*StockGenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockGenotypeR{}
		args[0] = object.GenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &stockGenotypeR{}
			args[i] = obj.GenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"genotype\" where \"genotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Genotype")
	}
	defer results.Close()

	var resultSlice []*Genotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Genotype")
	}

	if len(stockGenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Genotype = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GenotypeID == foreign.GenotypeID {
				local.R.Genotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadStock allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockGenotypeL) LoadStock(e boil.Executor, singular bool, maybeStockGenotype interface{}) error {
	var slice []*StockGenotype
	var object *StockGenotype

	count := 1
	if singular {
		object = maybeStockGenotype.(*StockGenotype)
	} else {
		slice = *maybeStockGenotype.(*StockGenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockGenotypeR{}
		args[0] = object.StockID
	} else {
		for i, obj := range slice {
			obj.R = &stockGenotypeR{}
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

	if len(stockGenotypeAfterSelectHooks) != 0 {
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

// SetGenotype of the stock_genotype to the related item.
// Sets o.R.Genotype to related.
// Adds o to related.R.StockGenotype.
func (o *StockGenotype) SetGenotype(exec boil.Executor, insert bool, related *Genotype) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_genotype\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"genotype_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockGenotypePrimaryKeyColumns),
	)
	values := []interface{}{related.GenotypeID, o.StockGenotypeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GenotypeID = related.GenotypeID

	if o.R == nil {
		o.R = &stockGenotypeR{
			Genotype: related,
		}
	} else {
		o.R.Genotype = related
	}

	if related.R == nil {
		related.R = &genotypeR{
			StockGenotype: o,
		}
	} else {
		related.R.StockGenotype = o
	}

	return nil
}

// SetStock of the stock_genotype to the related item.
// Sets o.R.Stock to related.
// Adds o to related.R.StockGenotype.
func (o *StockGenotype) SetStock(exec boil.Executor, insert bool, related *Stock) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stock_genotype\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stock_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockGenotypePrimaryKeyColumns),
	)
	values := []interface{}{related.StockID, o.StockGenotypeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StockID = related.StockID

	if o.R == nil {
		o.R = &stockGenotypeR{
			Stock: related,
		}
	} else {
		o.R.Stock = related
	}

	if related.R == nil {
		related.R = &stockR{
			StockGenotype: o,
		}
	} else {
		related.R.StockGenotype = o
	}

	return nil
}

// StockGenotypesG retrieves all records.
func StockGenotypesG(mods ...qm.QueryMod) stockGenotypeQuery {
	return StockGenotypes(boil.GetDB(), mods...)
}

// StockGenotypes retrieves all the records using an executor.
func StockGenotypes(exec boil.Executor, mods ...qm.QueryMod) stockGenotypeQuery {
	mods = append(mods, qm.From("\"stock_genotype\""))
	return stockGenotypeQuery{NewQuery(exec, mods...)}
}

// FindStockGenotypeG retrieves a single record by ID.
func FindStockGenotypeG(stockGenotypeID int, selectCols ...string) (*StockGenotype, error) {
	return FindStockGenotype(boil.GetDB(), stockGenotypeID, selectCols...)
}

// FindStockGenotypeGP retrieves a single record by ID, and panics on error.
func FindStockGenotypeGP(stockGenotypeID int, selectCols ...string) *StockGenotype {
	retobj, err := FindStockGenotype(boil.GetDB(), stockGenotypeID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockGenotype retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockGenotype(exec boil.Executor, stockGenotypeID int, selectCols ...string) (*StockGenotype, error) {
	stockGenotypeObj := &StockGenotype{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stock_genotype\" where \"stock_genotype_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockGenotypeID)

	err := q.Bind(stockGenotypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from stock_genotype")
	}

	return stockGenotypeObj, nil
}

// FindStockGenotypeP retrieves a single record by ID with an executor, and panics on error.
func FindStockGenotypeP(exec boil.Executor, stockGenotypeID int, selectCols ...string) *StockGenotype {
	retobj, err := FindStockGenotype(exec, stockGenotypeID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *StockGenotype) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *StockGenotype) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *StockGenotype) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *StockGenotype) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no stock_genotype provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockGenotypeColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockGenotypeInsertCacheMut.RLock()
	cache, cached := stockGenotypeInsertCache[key]
	stockGenotypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockGenotypeColumns,
			stockGenotypeColumnsWithDefault,
			stockGenotypeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockGenotypeType, stockGenotypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockGenotypeType, stockGenotypeMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stock_genotype\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into stock_genotype")
	}

	if !cached {
		stockGenotypeInsertCacheMut.Lock()
		stockGenotypeInsertCache[key] = cache
		stockGenotypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single StockGenotype record. See Update for
// whitelist behavior description.
func (o *StockGenotype) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single StockGenotype record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *StockGenotype) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the StockGenotype, and panics on error.
// See Update for whitelist behavior description.
func (o *StockGenotype) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the StockGenotype.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *StockGenotype) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockGenotypeUpdateCacheMut.RLock()
	cache, cached := stockGenotypeUpdateCache[key]
	stockGenotypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockGenotypeColumns, stockGenotypePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update stock_genotype, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stock_genotype\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockGenotypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockGenotypeType, stockGenotypeMapping, append(wl, stockGenotypePrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update stock_genotype row")
	}

	if !cached {
		stockGenotypeUpdateCacheMut.Lock()
		stockGenotypeUpdateCache[key] = cache
		stockGenotypeUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockGenotypeQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockGenotypeQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for stock_genotype")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockGenotypeSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockGenotypeSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockGenotypeSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockGenotypeSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockGenotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stock_genotype\" SET %s WHERE (\"stock_genotype_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockGenotypePrimaryKeyColumns), len(colNames)+1, len(stockGenotypePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in stockGenotype slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *StockGenotype) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *StockGenotype) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *StockGenotype) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *StockGenotype) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no stock_genotype provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockGenotypeColumnsWithDefault, o)

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

	stockGenotypeUpsertCacheMut.RLock()
	cache, cached := stockGenotypeUpsertCache[key]
	stockGenotypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockGenotypeColumns,
			stockGenotypeColumnsWithDefault,
			stockGenotypeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockGenotypeColumns,
			stockGenotypePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert stock_genotype, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockGenotypePrimaryKeyColumns))
			copy(conflict, stockGenotypePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stock_genotype\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockGenotypeType, stockGenotypeMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockGenotypeType, stockGenotypeMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for stock_genotype")
	}

	if !cached {
		stockGenotypeUpsertCacheMut.Lock()
		stockGenotypeUpsertCache[key] = cache
		stockGenotypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single StockGenotype record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockGenotype) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single StockGenotype record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *StockGenotype) DeleteG() error {
	if o == nil {
		return errors.New("chado: no StockGenotype provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single StockGenotype record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *StockGenotype) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single StockGenotype record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StockGenotype) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no StockGenotype provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockGenotypePrimaryKeyMapping)
	sql := "DELETE FROM \"stock_genotype\" WHERE \"stock_genotype_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from stock_genotype")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockGenotypeQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockGenotypeQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no stockGenotypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from stock_genotype")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockGenotypeSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockGenotypeSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no StockGenotype slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockGenotypeSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockGenotypeSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no StockGenotype slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockGenotypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockGenotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stock_genotype\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockGenotypePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockGenotypePrimaryKeyColumns), 1, len(stockGenotypePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from stockGenotype slice")
	}

	if len(stockGenotypeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *StockGenotype) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *StockGenotype) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *StockGenotype) ReloadG() error {
	if o == nil {
		return errors.New("chado: no StockGenotype provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *StockGenotype) Reload(exec boil.Executor) error {
	ret, err := FindStockGenotype(exec, o.StockGenotypeID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockGenotypeSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockGenotypeSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockGenotypeSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty StockGenotypeSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockGenotypeSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockGenotypes := StockGenotypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockGenotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stock_genotype\".* FROM \"stock_genotype\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockGenotypePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockGenotypePrimaryKeyColumns), 1, len(stockGenotypePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockGenotypes)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in StockGenotypeSlice")
	}

	*o = stockGenotypes

	return nil
}

// StockGenotypeExists checks if the StockGenotype row exists.
func StockGenotypeExists(exec boil.Executor, stockGenotypeID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stock_genotype\" where \"stock_genotype_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockGenotypeID)
	}

	row := exec.QueryRow(sql, stockGenotypeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if stock_genotype exists")
	}

	return exists, nil
}

// StockGenotypeExistsG checks if the StockGenotype row exists.
func StockGenotypeExistsG(stockGenotypeID int) (bool, error) {
	return StockGenotypeExists(boil.GetDB(), stockGenotypeID)
}

// StockGenotypeExistsGP checks if the StockGenotype row exists. Panics on error.
func StockGenotypeExistsGP(stockGenotypeID int) bool {
	e, err := StockGenotypeExists(boil.GetDB(), stockGenotypeID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockGenotypeExistsP checks if the StockGenotype row exists. Panics on error.
func StockGenotypeExistsP(exec boil.Executor, stockGenotypeID int) bool {
	e, err := StockGenotypeExists(exec, stockGenotypeID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
