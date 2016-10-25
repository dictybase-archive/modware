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

// Chadoprop is an object representing the database table.
type Chadoprop struct {
	ChadopropID int         `boil:"chadoprop_id" json:"chadoprop_id" toml:"chadoprop_id" yaml:"chadoprop_id"`
	TypeID      int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value       null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank        int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *chadopropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L chadopropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// chadopropR is where relationships are stored.
type chadopropR struct {
	Type *Cvterm
}

// chadopropL is where Load methods for each relationship are stored.
type chadopropL struct{}

var (
	chadopropColumns               = []string{"chadoprop_id", "type_id", "value", "rank"}
	chadopropColumnsWithoutDefault = []string{"type_id", "value"}
	chadopropColumnsWithDefault    = []string{"chadoprop_id", "rank"}
	chadopropPrimaryKeyColumns     = []string{"chadoprop_id"}
)

type (
	// ChadopropSlice is an alias for a slice of pointers to Chadoprop.
	// This should generally be used opposed to []Chadoprop.
	ChadopropSlice []*Chadoprop
	// ChadopropHook is the signature for custom Chadoprop hook methods
	ChadopropHook func(boil.Executor, *Chadoprop) error

	chadopropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	chadopropType                 = reflect.TypeOf(&Chadoprop{})
	chadopropMapping              = queries.MakeStructMapping(chadopropType)
	chadopropPrimaryKeyMapping, _ = queries.BindMapping(chadopropType, chadopropMapping, chadopropPrimaryKeyColumns)
	chadopropInsertCacheMut       sync.RWMutex
	chadopropInsertCache          = make(map[string]insertCache)
	chadopropUpdateCacheMut       sync.RWMutex
	chadopropUpdateCache          = make(map[string]updateCache)
	chadopropUpsertCacheMut       sync.RWMutex
	chadopropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var chadopropBeforeInsertHooks []ChadopropHook
var chadopropBeforeUpdateHooks []ChadopropHook
var chadopropBeforeDeleteHooks []ChadopropHook
var chadopropBeforeUpsertHooks []ChadopropHook

var chadopropAfterInsertHooks []ChadopropHook
var chadopropAfterSelectHooks []ChadopropHook
var chadopropAfterUpdateHooks []ChadopropHook
var chadopropAfterDeleteHooks []ChadopropHook
var chadopropAfterUpsertHooks []ChadopropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Chadoprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range chadopropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Chadoprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range chadopropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Chadoprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range chadopropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Chadoprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range chadopropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Chadoprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range chadopropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Chadoprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range chadopropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Chadoprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range chadopropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Chadoprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range chadopropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Chadoprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range chadopropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChadopropHook registers your hook function for all future operations.
func AddChadopropHook(hookPoint boil.HookPoint, chadopropHook ChadopropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		chadopropBeforeInsertHooks = append(chadopropBeforeInsertHooks, chadopropHook)
	case boil.BeforeUpdateHook:
		chadopropBeforeUpdateHooks = append(chadopropBeforeUpdateHooks, chadopropHook)
	case boil.BeforeDeleteHook:
		chadopropBeforeDeleteHooks = append(chadopropBeforeDeleteHooks, chadopropHook)
	case boil.BeforeUpsertHook:
		chadopropBeforeUpsertHooks = append(chadopropBeforeUpsertHooks, chadopropHook)
	case boil.AfterInsertHook:
		chadopropAfterInsertHooks = append(chadopropAfterInsertHooks, chadopropHook)
	case boil.AfterSelectHook:
		chadopropAfterSelectHooks = append(chadopropAfterSelectHooks, chadopropHook)
	case boil.AfterUpdateHook:
		chadopropAfterUpdateHooks = append(chadopropAfterUpdateHooks, chadopropHook)
	case boil.AfterDeleteHook:
		chadopropAfterDeleteHooks = append(chadopropAfterDeleteHooks, chadopropHook)
	case boil.AfterUpsertHook:
		chadopropAfterUpsertHooks = append(chadopropAfterUpsertHooks, chadopropHook)
	}
}

// OneP returns a single chadoprop record from the query, and panics on error.
func (q chadopropQuery) OneP() *Chadoprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single chadoprop record from the query.
func (q chadopropQuery) One() (*Chadoprop, error) {
	o := &Chadoprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for chadoprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Chadoprop records from the query, and panics on error.
func (q chadopropQuery) AllP() ChadopropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Chadoprop records from the query.
func (q chadopropQuery) All() (ChadopropSlice, error) {
	var o ChadopropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Chadoprop slice")
	}

	if len(chadopropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Chadoprop records in the query, and panics on error.
func (q chadopropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Chadoprop records in the query.
func (q chadopropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count chadoprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q chadopropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q chadopropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if chadoprop exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *Chadoprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Chadoprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (chadopropL) LoadType(e boil.Executor, singular bool, maybeChadoprop interface{}) error {
	var slice []*Chadoprop
	var object *Chadoprop

	count := 1
	if singular {
		object = maybeChadoprop.(*Chadoprop)
	} else {
		slice = *maybeChadoprop.(*ChadopropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &chadopropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &chadopropR{}
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

	if len(chadopropAfterSelectHooks) != 0 {
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

// SetType of the chadoprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeChadoprop.
func (o *Chadoprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"chadoprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, chadopropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.ChadopropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &chadopropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeChadoprop: o,
		}
	} else {
		related.R.TypeChadoprop = o
	}

	return nil
}

// ChadopropsG retrieves all records.
func ChadopropsG(mods ...qm.QueryMod) chadopropQuery {
	return Chadoprops(boil.GetDB(), mods...)
}

// Chadoprops retrieves all the records using an executor.
func Chadoprops(exec boil.Executor, mods ...qm.QueryMod) chadopropQuery {
	mods = append(mods, qm.From("\"chadoprop\""))
	return chadopropQuery{NewQuery(exec, mods...)}
}

// FindChadopropG retrieves a single record by ID.
func FindChadopropG(chadopropID int, selectCols ...string) (*Chadoprop, error) {
	return FindChadoprop(boil.GetDB(), chadopropID, selectCols...)
}

// FindChadopropGP retrieves a single record by ID, and panics on error.
func FindChadopropGP(chadopropID int, selectCols ...string) *Chadoprop {
	retobj, err := FindChadoprop(boil.GetDB(), chadopropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindChadoprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChadoprop(exec boil.Executor, chadopropID int, selectCols ...string) (*Chadoprop, error) {
	chadopropObj := &Chadoprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"chadoprop\" where \"chadoprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, chadopropID)

	err := q.Bind(chadopropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from chadoprop")
	}

	return chadopropObj, nil
}

// FindChadopropP retrieves a single record by ID with an executor, and panics on error.
func FindChadopropP(exec boil.Executor, chadopropID int, selectCols ...string) *Chadoprop {
	retobj, err := FindChadoprop(exec, chadopropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Chadoprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Chadoprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Chadoprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Chadoprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no chadoprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(chadopropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	chadopropInsertCacheMut.RLock()
	cache, cached := chadopropInsertCache[key]
	chadopropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			chadopropColumns,
			chadopropColumnsWithDefault,
			chadopropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(chadopropType, chadopropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(chadopropType, chadopropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"chadoprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into chadoprop")
	}

	if !cached {
		chadopropInsertCacheMut.Lock()
		chadopropInsertCache[key] = cache
		chadopropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Chadoprop record. See Update for
// whitelist behavior description.
func (o *Chadoprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Chadoprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Chadoprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Chadoprop, and panics on error.
// See Update for whitelist behavior description.
func (o *Chadoprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Chadoprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Chadoprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	chadopropUpdateCacheMut.RLock()
	cache, cached := chadopropUpdateCache[key]
	chadopropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(chadopropColumns, chadopropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update chadoprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"chadoprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, chadopropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(chadopropType, chadopropMapping, append(wl, chadopropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update chadoprop row")
	}

	if !cached {
		chadopropUpdateCacheMut.Lock()
		chadopropUpdateCache[key] = cache
		chadopropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q chadopropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q chadopropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for chadoprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ChadopropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ChadopropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ChadopropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChadopropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chadopropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"chadoprop\" SET %s WHERE (\"chadoprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(chadopropPrimaryKeyColumns), len(colNames)+1, len(chadopropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in chadoprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Chadoprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Chadoprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Chadoprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Chadoprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no chadoprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(chadopropColumnsWithDefault, o)

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

	chadopropUpsertCacheMut.RLock()
	cache, cached := chadopropUpsertCache[key]
	chadopropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			chadopropColumns,
			chadopropColumnsWithDefault,
			chadopropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			chadopropColumns,
			chadopropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert chadoprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(chadopropPrimaryKeyColumns))
			copy(conflict, chadopropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"chadoprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(chadopropType, chadopropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(chadopropType, chadopropMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for chadoprop")
	}

	if !cached {
		chadopropUpsertCacheMut.Lock()
		chadopropUpsertCache[key] = cache
		chadopropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Chadoprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Chadoprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Chadoprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Chadoprop) DeleteG() error {
	if o == nil {
		return errors.New("models: no Chadoprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Chadoprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Chadoprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Chadoprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Chadoprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Chadoprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), chadopropPrimaryKeyMapping)
	sql := "DELETE FROM \"chadoprop\" WHERE \"chadoprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from chadoprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q chadopropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q chadopropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no chadopropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from chadoprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ChadopropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ChadopropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Chadoprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ChadopropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChadopropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Chadoprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(chadopropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chadopropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"chadoprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, chadopropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(chadopropPrimaryKeyColumns), 1, len(chadopropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from chadoprop slice")
	}

	if len(chadopropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Chadoprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Chadoprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Chadoprop) ReloadG() error {
	if o == nil {
		return errors.New("models: no Chadoprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Chadoprop) Reload(exec boil.Executor) error {
	ret, err := FindChadoprop(exec, o.ChadopropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ChadopropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ChadopropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChadopropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ChadopropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChadopropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	chadoprops := ChadopropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chadopropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"chadoprop\".* FROM \"chadoprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, chadopropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(chadopropPrimaryKeyColumns), 1, len(chadopropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&chadoprops)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChadopropSlice")
	}

	*o = chadoprops

	return nil
}

// ChadopropExists checks if the Chadoprop row exists.
func ChadopropExists(exec boil.Executor, chadopropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"chadoprop\" where \"chadoprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, chadopropID)
	}

	row := exec.QueryRow(sql, chadopropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if chadoprop exists")
	}

	return exists, nil
}

// ChadopropExistsG checks if the Chadoprop row exists.
func ChadopropExistsG(chadopropID int) (bool, error) {
	return ChadopropExists(boil.GetDB(), chadopropID)
}

// ChadopropExistsGP checks if the Chadoprop row exists. Panics on error.
func ChadopropExistsGP(chadopropID int) bool {
	e, err := ChadopropExists(boil.GetDB(), chadopropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ChadopropExistsP checks if the Chadoprop row exists. Panics on error.
func ChadopropExistsP(exec boil.Executor, chadopropID int) bool {
	e, err := ChadopropExists(exec, chadopropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
