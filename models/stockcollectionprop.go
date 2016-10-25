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

// Stockcollectionprop is an object representing the database table.
type Stockcollectionprop struct {
	StockcollectionpropID int         `boil:"stockcollectionprop_id" json:"stockcollectionprop_id" toml:"stockcollectionprop_id" yaml:"stockcollectionprop_id"`
	StockcollectionID     int         `boil:"stockcollection_id" json:"stockcollection_id" toml:"stockcollection_id" yaml:"stockcollection_id"`
	TypeID                int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value                 null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank                  int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *stockcollectionpropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockcollectionpropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// stockcollectionpropR is where relationships are stored.
type stockcollectionpropR struct {
	Type            *Cvterm
	Stockcollection *Stockcollection
}

// stockcollectionpropL is where Load methods for each relationship are stored.
type stockcollectionpropL struct{}

var (
	stockcollectionpropColumns               = []string{"stockcollectionprop_id", "stockcollection_id", "type_id", "value", "rank"}
	stockcollectionpropColumnsWithoutDefault = []string{"stockcollection_id", "type_id", "value"}
	stockcollectionpropColumnsWithDefault    = []string{"stockcollectionprop_id", "rank"}
	stockcollectionpropPrimaryKeyColumns     = []string{"stockcollectionprop_id"}
)

type (
	// StockcollectionpropSlice is an alias for a slice of pointers to Stockcollectionprop.
	// This should generally be used opposed to []Stockcollectionprop.
	StockcollectionpropSlice []*Stockcollectionprop
	// StockcollectionpropHook is the signature for custom Stockcollectionprop hook methods
	StockcollectionpropHook func(boil.Executor, *Stockcollectionprop) error

	stockcollectionpropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockcollectionpropType                 = reflect.TypeOf(&Stockcollectionprop{})
	stockcollectionpropMapping              = queries.MakeStructMapping(stockcollectionpropType)
	stockcollectionpropPrimaryKeyMapping, _ = queries.BindMapping(stockcollectionpropType, stockcollectionpropMapping, stockcollectionpropPrimaryKeyColumns)
	stockcollectionpropInsertCacheMut       sync.RWMutex
	stockcollectionpropInsertCache          = make(map[string]insertCache)
	stockcollectionpropUpdateCacheMut       sync.RWMutex
	stockcollectionpropUpdateCache          = make(map[string]updateCache)
	stockcollectionpropUpsertCacheMut       sync.RWMutex
	stockcollectionpropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var stockcollectionpropBeforeInsertHooks []StockcollectionpropHook
var stockcollectionpropBeforeUpdateHooks []StockcollectionpropHook
var stockcollectionpropBeforeDeleteHooks []StockcollectionpropHook
var stockcollectionpropBeforeUpsertHooks []StockcollectionpropHook

var stockcollectionpropAfterInsertHooks []StockcollectionpropHook
var stockcollectionpropAfterSelectHooks []StockcollectionpropHook
var stockcollectionpropAfterUpdateHooks []StockcollectionpropHook
var stockcollectionpropAfterDeleteHooks []StockcollectionpropHook
var stockcollectionpropAfterUpsertHooks []StockcollectionpropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Stockcollectionprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionpropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Stockcollectionprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionpropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Stockcollectionprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionpropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Stockcollectionprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionpropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Stockcollectionprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionpropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Stockcollectionprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionpropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Stockcollectionprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionpropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Stockcollectionprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionpropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Stockcollectionprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range stockcollectionpropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockcollectionpropHook registers your hook function for all future operations.
func AddStockcollectionpropHook(hookPoint boil.HookPoint, stockcollectionpropHook StockcollectionpropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockcollectionpropBeforeInsertHooks = append(stockcollectionpropBeforeInsertHooks, stockcollectionpropHook)
	case boil.BeforeUpdateHook:
		stockcollectionpropBeforeUpdateHooks = append(stockcollectionpropBeforeUpdateHooks, stockcollectionpropHook)
	case boil.BeforeDeleteHook:
		stockcollectionpropBeforeDeleteHooks = append(stockcollectionpropBeforeDeleteHooks, stockcollectionpropHook)
	case boil.BeforeUpsertHook:
		stockcollectionpropBeforeUpsertHooks = append(stockcollectionpropBeforeUpsertHooks, stockcollectionpropHook)
	case boil.AfterInsertHook:
		stockcollectionpropAfterInsertHooks = append(stockcollectionpropAfterInsertHooks, stockcollectionpropHook)
	case boil.AfterSelectHook:
		stockcollectionpropAfterSelectHooks = append(stockcollectionpropAfterSelectHooks, stockcollectionpropHook)
	case boil.AfterUpdateHook:
		stockcollectionpropAfterUpdateHooks = append(stockcollectionpropAfterUpdateHooks, stockcollectionpropHook)
	case boil.AfterDeleteHook:
		stockcollectionpropAfterDeleteHooks = append(stockcollectionpropAfterDeleteHooks, stockcollectionpropHook)
	case boil.AfterUpsertHook:
		stockcollectionpropAfterUpsertHooks = append(stockcollectionpropAfterUpsertHooks, stockcollectionpropHook)
	}
}

// OneP returns a single stockcollectionprop record from the query, and panics on error.
func (q stockcollectionpropQuery) OneP() *Stockcollectionprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single stockcollectionprop record from the query.
func (q stockcollectionpropQuery) One() (*Stockcollectionprop, error) {
	o := &Stockcollectionprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stockcollectionprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Stockcollectionprop records from the query, and panics on error.
func (q stockcollectionpropQuery) AllP() StockcollectionpropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Stockcollectionprop records from the query.
func (q stockcollectionpropQuery) All() (StockcollectionpropSlice, error) {
	var o StockcollectionpropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Stockcollectionprop slice")
	}

	if len(stockcollectionpropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Stockcollectionprop records in the query, and panics on error.
func (q stockcollectionpropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Stockcollectionprop records in the query.
func (q stockcollectionpropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stockcollectionprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q stockcollectionpropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q stockcollectionpropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stockcollectionprop exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *Stockcollectionprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Stockcollectionprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// StockcollectionG pointed to by the foreign key.
func (o *Stockcollectionprop) StockcollectionG(mods ...qm.QueryMod) stockcollectionQuery {
	return o.Stockcollection(boil.GetDB(), mods...)
}

// Stockcollection pointed to by the foreign key.
func (o *Stockcollectionprop) Stockcollection(exec boil.Executor, mods ...qm.QueryMod) stockcollectionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("stockcollection_id=$1", o.StockcollectionID),
	}

	queryMods = append(queryMods, mods...)

	query := Stockcollections(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stockcollection\"")

	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockcollectionpropL) LoadType(e boil.Executor, singular bool, maybeStockcollectionprop interface{}) error {
	var slice []*Stockcollectionprop
	var object *Stockcollectionprop

	count := 1
	if singular {
		object = maybeStockcollectionprop.(*Stockcollectionprop)
	} else {
		slice = *maybeStockcollectionprop.(*StockcollectionpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockcollectionpropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &stockcollectionpropR{}
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

	if len(stockcollectionpropAfterSelectHooks) != 0 {
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

// LoadStockcollection allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (stockcollectionpropL) LoadStockcollection(e boil.Executor, singular bool, maybeStockcollectionprop interface{}) error {
	var slice []*Stockcollectionprop
	var object *Stockcollectionprop

	count := 1
	if singular {
		object = maybeStockcollectionprop.(*Stockcollectionprop)
	} else {
		slice = *maybeStockcollectionprop.(*StockcollectionpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &stockcollectionpropR{}
		args[0] = object.StockcollectionID
	} else {
		for i, obj := range slice {
			obj.R = &stockcollectionpropR{}
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

	if len(stockcollectionpropAfterSelectHooks) != 0 {
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

// SetType of the stockcollectionprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeStockcollectionprop.
func (o *Stockcollectionprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stockcollectionprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockcollectionpropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.StockcollectionpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &stockcollectionpropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeStockcollectionprop: o,
		}
	} else {
		related.R.TypeStockcollectionprop = o
	}

	return nil
}

// SetStockcollection of the stockcollectionprop to the related item.
// Sets o.R.Stockcollection to related.
// Adds o to related.R.Stockcollectionprop.
func (o *Stockcollectionprop) SetStockcollection(exec boil.Executor, insert bool, related *Stockcollection) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"stockcollectionprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stockcollection_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockcollectionpropPrimaryKeyColumns),
	)
	values := []interface{}{related.StockcollectionID, o.StockcollectionpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StockcollectionID = related.StockcollectionID

	if o.R == nil {
		o.R = &stockcollectionpropR{
			Stockcollection: related,
		}
	} else {
		o.R.Stockcollection = related
	}

	if related.R == nil {
		related.R = &stockcollectionR{
			Stockcollectionprop: o,
		}
	} else {
		related.R.Stockcollectionprop = o
	}

	return nil
}

// StockcollectionpropsG retrieves all records.
func StockcollectionpropsG(mods ...qm.QueryMod) stockcollectionpropQuery {
	return Stockcollectionprops(boil.GetDB(), mods...)
}

// Stockcollectionprops retrieves all the records using an executor.
func Stockcollectionprops(exec boil.Executor, mods ...qm.QueryMod) stockcollectionpropQuery {
	mods = append(mods, qm.From("\"stockcollectionprop\""))
	return stockcollectionpropQuery{NewQuery(exec, mods...)}
}

// FindStockcollectionpropG retrieves a single record by ID.
func FindStockcollectionpropG(stockcollectionpropID int, selectCols ...string) (*Stockcollectionprop, error) {
	return FindStockcollectionprop(boil.GetDB(), stockcollectionpropID, selectCols...)
}

// FindStockcollectionpropGP retrieves a single record by ID, and panics on error.
func FindStockcollectionpropGP(stockcollectionpropID int, selectCols ...string) *Stockcollectionprop {
	retobj, err := FindStockcollectionprop(boil.GetDB(), stockcollectionpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindStockcollectionprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStockcollectionprop(exec boil.Executor, stockcollectionpropID int, selectCols ...string) (*Stockcollectionprop, error) {
	stockcollectionpropObj := &Stockcollectionprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stockcollectionprop\" where \"stockcollectionprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, stockcollectionpropID)

	err := q.Bind(stockcollectionpropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stockcollectionprop")
	}

	return stockcollectionpropObj, nil
}

// FindStockcollectionpropP retrieves a single record by ID with an executor, and panics on error.
func FindStockcollectionpropP(exec boil.Executor, stockcollectionpropID int, selectCols ...string) *Stockcollectionprop {
	retobj, err := FindStockcollectionprop(exec, stockcollectionpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Stockcollectionprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Stockcollectionprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Stockcollectionprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Stockcollectionprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stockcollectionprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockcollectionpropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	stockcollectionpropInsertCacheMut.RLock()
	cache, cached := stockcollectionpropInsertCache[key]
	stockcollectionpropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			stockcollectionpropColumns,
			stockcollectionpropColumnsWithDefault,
			stockcollectionpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(stockcollectionpropType, stockcollectionpropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockcollectionpropType, stockcollectionpropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"stockcollectionprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into stockcollectionprop")
	}

	if !cached {
		stockcollectionpropInsertCacheMut.Lock()
		stockcollectionpropInsertCache[key] = cache
		stockcollectionpropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Stockcollectionprop record. See Update for
// whitelist behavior description.
func (o *Stockcollectionprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Stockcollectionprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Stockcollectionprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Stockcollectionprop, and panics on error.
// See Update for whitelist behavior description.
func (o *Stockcollectionprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Stockcollectionprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Stockcollectionprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	stockcollectionpropUpdateCacheMut.RLock()
	cache, cached := stockcollectionpropUpdateCache[key]
	stockcollectionpropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(stockcollectionpropColumns, stockcollectionpropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update stockcollectionprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stockcollectionprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockcollectionpropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockcollectionpropType, stockcollectionpropMapping, append(wl, stockcollectionpropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update stockcollectionprop row")
	}

	if !cached {
		stockcollectionpropUpdateCacheMut.Lock()
		stockcollectionpropUpdateCache[key] = cache
		stockcollectionpropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q stockcollectionpropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q stockcollectionpropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for stockcollectionprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockcollectionpropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o StockcollectionpropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o StockcollectionpropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockcollectionpropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockcollectionpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"stockcollectionprop\" SET %s WHERE (\"stockcollectionprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockcollectionpropPrimaryKeyColumns), len(colNames)+1, len(stockcollectionpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in stockcollectionprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Stockcollectionprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Stockcollectionprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Stockcollectionprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Stockcollectionprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no stockcollectionprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockcollectionpropColumnsWithDefault, o)

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

	stockcollectionpropUpsertCacheMut.RLock()
	cache, cached := stockcollectionpropUpsertCache[key]
	stockcollectionpropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			stockcollectionpropColumns,
			stockcollectionpropColumnsWithDefault,
			stockcollectionpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			stockcollectionpropColumns,
			stockcollectionpropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert stockcollectionprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockcollectionpropPrimaryKeyColumns))
			copy(conflict, stockcollectionpropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"stockcollectionprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(stockcollectionpropType, stockcollectionpropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockcollectionpropType, stockcollectionpropMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for stockcollectionprop")
	}

	if !cached {
		stockcollectionpropUpsertCacheMut.Lock()
		stockcollectionpropUpsertCache[key] = cache
		stockcollectionpropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Stockcollectionprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Stockcollectionprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Stockcollectionprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Stockcollectionprop) DeleteG() error {
	if o == nil {
		return errors.New("models: no Stockcollectionprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Stockcollectionprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Stockcollectionprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Stockcollectionprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Stockcollectionprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Stockcollectionprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockcollectionpropPrimaryKeyMapping)
	sql := "DELETE FROM \"stockcollectionprop\" WHERE \"stockcollectionprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from stockcollectionprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q stockcollectionpropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q stockcollectionpropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no stockcollectionpropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockcollectionprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o StockcollectionpropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o StockcollectionpropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Stockcollectionprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o StockcollectionpropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockcollectionpropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Stockcollectionprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(stockcollectionpropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockcollectionpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"stockcollectionprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockcollectionpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(stockcollectionpropPrimaryKeyColumns), 1, len(stockcollectionpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from stockcollectionprop slice")
	}

	if len(stockcollectionpropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Stockcollectionprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Stockcollectionprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Stockcollectionprop) ReloadG() error {
	if o == nil {
		return errors.New("models: no Stockcollectionprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Stockcollectionprop) Reload(exec boil.Executor) error {
	ret, err := FindStockcollectionprop(exec, o.StockcollectionpropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockcollectionpropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *StockcollectionpropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockcollectionpropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty StockcollectionpropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockcollectionpropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	stockcollectionprops := StockcollectionpropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockcollectionpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"stockcollectionprop\".* FROM \"stockcollectionprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, stockcollectionpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(stockcollectionpropPrimaryKeyColumns), 1, len(stockcollectionpropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&stockcollectionprops)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockcollectionpropSlice")
	}

	*o = stockcollectionprops

	return nil
}

// StockcollectionpropExists checks if the Stockcollectionprop row exists.
func StockcollectionpropExists(exec boil.Executor, stockcollectionpropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"stockcollectionprop\" where \"stockcollectionprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, stockcollectionpropID)
	}

	row := exec.QueryRow(sql, stockcollectionpropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stockcollectionprop exists")
	}

	return exists, nil
}

// StockcollectionpropExistsG checks if the Stockcollectionprop row exists.
func StockcollectionpropExistsG(stockcollectionpropID int) (bool, error) {
	return StockcollectionpropExists(boil.GetDB(), stockcollectionpropID)
}

// StockcollectionpropExistsGP checks if the Stockcollectionprop row exists. Panics on error.
func StockcollectionpropExistsGP(stockcollectionpropID int) bool {
	e, err := StockcollectionpropExists(boil.GetDB(), stockcollectionpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// StockcollectionpropExistsP checks if the Stockcollectionprop row exists. Panics on error.
func StockcollectionpropExistsP(exec boil.Executor, stockcollectionpropID int) bool {
	e, err := StockcollectionpropExists(exec, stockcollectionpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
