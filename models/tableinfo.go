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

// Tableinfo is an object representing the database table.
type Tableinfo struct {
	TableinfoID       int         `boil:"tableinfo_id" json:"tableinfo_id" toml:"tableinfo_id" yaml:"tableinfo_id"`
	Name              string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	PrimaryKeyColumn  null.String `boil:"primary_key_column" json:"primary_key_column,omitempty" toml:"primary_key_column" yaml:"primary_key_column,omitempty"`
	IsView            int         `boil:"is_view" json:"is_view" toml:"is_view" yaml:"is_view"`
	ViewOnTableID     null.Int    `boil:"view_on_table_id" json:"view_on_table_id,omitempty" toml:"view_on_table_id" yaml:"view_on_table_id,omitempty"`
	SuperclassTableID null.Int    `boil:"superclass_table_id" json:"superclass_table_id,omitempty" toml:"superclass_table_id" yaml:"superclass_table_id,omitempty"`
	IsUpdateable      int         `boil:"is_updateable" json:"is_updateable" toml:"is_updateable" yaml:"is_updateable"`
	ModificationDate  time.Time   `boil:"modification_date" json:"modification_date" toml:"modification_date" yaml:"modification_date"`

	R *tableinfoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tableinfoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// tableinfoR is where relationships are stored.
type tableinfoR struct {
}

// tableinfoL is where Load methods for each relationship are stored.
type tableinfoL struct{}

var (
	tableinfoColumns               = []string{"tableinfo_id", "name", "primary_key_column", "is_view", "view_on_table_id", "superclass_table_id", "is_updateable", "modification_date"}
	tableinfoColumnsWithoutDefault = []string{"name", "primary_key_column", "view_on_table_id", "superclass_table_id"}
	tableinfoColumnsWithDefault    = []string{"tableinfo_id", "is_view", "is_updateable", "modification_date"}
	tableinfoPrimaryKeyColumns     = []string{"tableinfo_id"}
)

type (
	// TableinfoSlice is an alias for a slice of pointers to Tableinfo.
	// This should generally be used opposed to []Tableinfo.
	TableinfoSlice []*Tableinfo
	// TableinfoHook is the signature for custom Tableinfo hook methods
	TableinfoHook func(boil.Executor, *Tableinfo) error

	tableinfoQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tableinfoType                 = reflect.TypeOf(&Tableinfo{})
	tableinfoMapping              = queries.MakeStructMapping(tableinfoType)
	tableinfoPrimaryKeyMapping, _ = queries.BindMapping(tableinfoType, tableinfoMapping, tableinfoPrimaryKeyColumns)
	tableinfoInsertCacheMut       sync.RWMutex
	tableinfoInsertCache          = make(map[string]insertCache)
	tableinfoUpdateCacheMut       sync.RWMutex
	tableinfoUpdateCache          = make(map[string]updateCache)
	tableinfoUpsertCacheMut       sync.RWMutex
	tableinfoUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var tableinfoBeforeInsertHooks []TableinfoHook
var tableinfoBeforeUpdateHooks []TableinfoHook
var tableinfoBeforeDeleteHooks []TableinfoHook
var tableinfoBeforeUpsertHooks []TableinfoHook

var tableinfoAfterInsertHooks []TableinfoHook
var tableinfoAfterSelectHooks []TableinfoHook
var tableinfoAfterUpdateHooks []TableinfoHook
var tableinfoAfterDeleteHooks []TableinfoHook
var tableinfoAfterUpsertHooks []TableinfoHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Tableinfo) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tableinfoBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Tableinfo) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range tableinfoBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Tableinfo) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range tableinfoBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Tableinfo) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tableinfoBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Tableinfo) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tableinfoAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Tableinfo) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range tableinfoAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Tableinfo) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range tableinfoAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Tableinfo) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range tableinfoAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Tableinfo) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range tableinfoAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTableinfoHook registers your hook function for all future operations.
func AddTableinfoHook(hookPoint boil.HookPoint, tableinfoHook TableinfoHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		tableinfoBeforeInsertHooks = append(tableinfoBeforeInsertHooks, tableinfoHook)
	case boil.BeforeUpdateHook:
		tableinfoBeforeUpdateHooks = append(tableinfoBeforeUpdateHooks, tableinfoHook)
	case boil.BeforeDeleteHook:
		tableinfoBeforeDeleteHooks = append(tableinfoBeforeDeleteHooks, tableinfoHook)
	case boil.BeforeUpsertHook:
		tableinfoBeforeUpsertHooks = append(tableinfoBeforeUpsertHooks, tableinfoHook)
	case boil.AfterInsertHook:
		tableinfoAfterInsertHooks = append(tableinfoAfterInsertHooks, tableinfoHook)
	case boil.AfterSelectHook:
		tableinfoAfterSelectHooks = append(tableinfoAfterSelectHooks, tableinfoHook)
	case boil.AfterUpdateHook:
		tableinfoAfterUpdateHooks = append(tableinfoAfterUpdateHooks, tableinfoHook)
	case boil.AfterDeleteHook:
		tableinfoAfterDeleteHooks = append(tableinfoAfterDeleteHooks, tableinfoHook)
	case boil.AfterUpsertHook:
		tableinfoAfterUpsertHooks = append(tableinfoAfterUpsertHooks, tableinfoHook)
	}
}

// OneP returns a single tableinfo record from the query, and panics on error.
func (q tableinfoQuery) OneP() *Tableinfo {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single tableinfo record from the query.
func (q tableinfoQuery) One() (*Tableinfo, error) {
	o := &Tableinfo{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for tableinfo")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Tableinfo records from the query, and panics on error.
func (q tableinfoQuery) AllP() TableinfoSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Tableinfo records from the query.
func (q tableinfoQuery) All() (TableinfoSlice, error) {
	var o TableinfoSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Tableinfo slice")
	}

	if len(tableinfoAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Tableinfo records in the query, and panics on error.
func (q tableinfoQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Tableinfo records in the query.
func (q tableinfoQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count tableinfo rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q tableinfoQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q tableinfoQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if tableinfo exists")
	}

	return count > 0, nil
}

// TableinfosG retrieves all records.
func TableinfosG(mods ...qm.QueryMod) tableinfoQuery {
	return Tableinfos(boil.GetDB(), mods...)
}

// Tableinfos retrieves all the records using an executor.
func Tableinfos(exec boil.Executor, mods ...qm.QueryMod) tableinfoQuery {
	mods = append(mods, qm.From("\"tableinfo\""))
	return tableinfoQuery{NewQuery(exec, mods...)}
}

// FindTableinfoG retrieves a single record by ID.
func FindTableinfoG(tableinfoID int, selectCols ...string) (*Tableinfo, error) {
	return FindTableinfo(boil.GetDB(), tableinfoID, selectCols...)
}

// FindTableinfoGP retrieves a single record by ID, and panics on error.
func FindTableinfoGP(tableinfoID int, selectCols ...string) *Tableinfo {
	retobj, err := FindTableinfo(boil.GetDB(), tableinfoID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindTableinfo retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTableinfo(exec boil.Executor, tableinfoID int, selectCols ...string) (*Tableinfo, error) {
	tableinfoObj := &Tableinfo{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tableinfo\" where \"tableinfo_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, tableinfoID)

	err := q.Bind(tableinfoObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from tableinfo")
	}

	return tableinfoObj, nil
}

// FindTableinfoP retrieves a single record by ID with an executor, and panics on error.
func FindTableinfoP(exec boil.Executor, tableinfoID int, selectCols ...string) *Tableinfo {
	retobj, err := FindTableinfo(exec, tableinfoID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Tableinfo) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Tableinfo) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Tableinfo) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Tableinfo) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no tableinfo provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tableinfoColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	tableinfoInsertCacheMut.RLock()
	cache, cached := tableinfoInsertCache[key]
	tableinfoInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			tableinfoColumns,
			tableinfoColumnsWithDefault,
			tableinfoColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(tableinfoType, tableinfoMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tableinfoType, tableinfoMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"tableinfo\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into tableinfo")
	}

	if !cached {
		tableinfoInsertCacheMut.Lock()
		tableinfoInsertCache[key] = cache
		tableinfoInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Tableinfo record. See Update for
// whitelist behavior description.
func (o *Tableinfo) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Tableinfo record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Tableinfo) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Tableinfo, and panics on error.
// See Update for whitelist behavior description.
func (o *Tableinfo) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Tableinfo.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Tableinfo) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	tableinfoUpdateCacheMut.RLock()
	cache, cached := tableinfoUpdateCache[key]
	tableinfoUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(tableinfoColumns, tableinfoPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update tableinfo, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tableinfo\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, tableinfoPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tableinfoType, tableinfoMapping, append(wl, tableinfoPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update tableinfo row")
	}

	if !cached {
		tableinfoUpdateCacheMut.Lock()
		tableinfoUpdateCache[key] = cache
		tableinfoUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q tableinfoQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q tableinfoQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for tableinfo")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TableinfoSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o TableinfoSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o TableinfoSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TableinfoSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tableinfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"tableinfo\" SET %s WHERE (\"tableinfo_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(tableinfoPrimaryKeyColumns), len(colNames)+1, len(tableinfoPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in tableinfo slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Tableinfo) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Tableinfo) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Tableinfo) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Tableinfo) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no tableinfo provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tableinfoColumnsWithDefault, o)

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

	tableinfoUpsertCacheMut.RLock()
	cache, cached := tableinfoUpsertCache[key]
	tableinfoUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			tableinfoColumns,
			tableinfoColumnsWithDefault,
			tableinfoColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			tableinfoColumns,
			tableinfoPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert tableinfo, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(tableinfoPrimaryKeyColumns))
			copy(conflict, tableinfoPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"tableinfo\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(tableinfoType, tableinfoMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tableinfoType, tableinfoMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for tableinfo")
	}

	if !cached {
		tableinfoUpsertCacheMut.Lock()
		tableinfoUpsertCache[key] = cache
		tableinfoUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Tableinfo record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Tableinfo) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Tableinfo record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Tableinfo) DeleteG() error {
	if o == nil {
		return errors.New("models: no Tableinfo provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Tableinfo record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Tableinfo) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Tableinfo record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Tableinfo) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Tableinfo provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tableinfoPrimaryKeyMapping)
	sql := "DELETE FROM \"tableinfo\" WHERE \"tableinfo_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from tableinfo")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q tableinfoQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q tableinfoQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no tableinfoQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from tableinfo")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o TableinfoSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o TableinfoSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Tableinfo slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o TableinfoSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TableinfoSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Tableinfo slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(tableinfoBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tableinfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"tableinfo\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, tableinfoPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(tableinfoPrimaryKeyColumns), 1, len(tableinfoPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from tableinfo slice")
	}

	if len(tableinfoAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Tableinfo) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Tableinfo) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Tableinfo) ReloadG() error {
	if o == nil {
		return errors.New("models: no Tableinfo provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Tableinfo) Reload(exec boil.Executor) error {
	ret, err := FindTableinfo(exec, o.TableinfoID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *TableinfoSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *TableinfoSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TableinfoSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty TableinfoSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TableinfoSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	tableinfos := TableinfoSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tableinfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"tableinfo\".* FROM \"tableinfo\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, tableinfoPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(tableinfoPrimaryKeyColumns), 1, len(tableinfoPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&tableinfos)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TableinfoSlice")
	}

	*o = tableinfos

	return nil
}

// TableinfoExists checks if the Tableinfo row exists.
func TableinfoExists(exec boil.Executor, tableinfoID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"tableinfo\" where \"tableinfo_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, tableinfoID)
	}

	row := exec.QueryRow(sql, tableinfoID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if tableinfo exists")
	}

	return exists, nil
}

// TableinfoExistsG checks if the Tableinfo row exists.
func TableinfoExistsG(tableinfoID int) (bool, error) {
	return TableinfoExists(boil.GetDB(), tableinfoID)
}

// TableinfoExistsGP checks if the Tableinfo row exists. Panics on error.
func TableinfoExistsGP(tableinfoID int) bool {
	e, err := TableinfoExists(boil.GetDB(), tableinfoID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// TableinfoExistsP checks if the Tableinfo row exists. Panics on error.
func TableinfoExistsP(exec boil.Executor, tableinfoID int) bool {
	e, err := TableinfoExists(exec, tableinfoID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
