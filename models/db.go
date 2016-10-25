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

// DB is an object representing the database table.
type DB struct {
	DBID        int         `boil:"db_id" json:"db_id" toml:"db_id" yaml:"db_id"`
	Name        string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Urlprefix   null.String `boil:"urlprefix" json:"urlprefix,omitempty" toml:"urlprefix" yaml:"urlprefix,omitempty"`
	Url         null.String `boil:"url" json:"url,omitempty" toml:"url" yaml:"url,omitempty"`

	R *dbR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dbL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dbR is where relationships are stored.
type dbR struct {
	Dbxref *Dbxref
}

// dbL is where Load methods for each relationship are stored.
type dbL struct{}

var (
	dbColumns               = []string{"db_id", "name", "description", "urlprefix", "url"}
	dbColumnsWithoutDefault = []string{"name", "description", "urlprefix", "url"}
	dbColumnsWithDefault    = []string{"db_id"}
	dbPrimaryKeyColumns     = []string{"db_id"}
)

type (
	// DBSlice is an alias for a slice of pointers to DB.
	// This should generally be used opposed to []DB.
	DBSlice []*DB
	// DBHook is the signature for custom DB hook methods
	DBHook func(boil.Executor, *DB) error

	dbQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dbType                 = reflect.TypeOf(&DB{})
	dbMapping              = queries.MakeStructMapping(dbType)
	dbPrimaryKeyMapping, _ = queries.BindMapping(dbType, dbMapping, dbPrimaryKeyColumns)
	dbInsertCacheMut       sync.RWMutex
	dbInsertCache          = make(map[string]insertCache)
	dbUpdateCacheMut       sync.RWMutex
	dbUpdateCache          = make(map[string]updateCache)
	dbUpsertCacheMut       sync.RWMutex
	dbUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var dbBeforeInsertHooks []DBHook
var dbBeforeUpdateHooks []DBHook
var dbBeforeDeleteHooks []DBHook
var dbBeforeUpsertHooks []DBHook

var dbAfterInsertHooks []DBHook
var dbAfterSelectHooks []DBHook
var dbAfterUpdateHooks []DBHook
var dbAfterDeleteHooks []DBHook
var dbAfterUpsertHooks []DBHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DB) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dbBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DB) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range dbBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DB) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range dbBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DB) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dbBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DB) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dbAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DB) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range dbAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DB) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range dbAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DB) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range dbAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DB) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dbAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDBHook registers your hook function for all future operations.
func AddDBHook(hookPoint boil.HookPoint, dbHook DBHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dbBeforeInsertHooks = append(dbBeforeInsertHooks, dbHook)
	case boil.BeforeUpdateHook:
		dbBeforeUpdateHooks = append(dbBeforeUpdateHooks, dbHook)
	case boil.BeforeDeleteHook:
		dbBeforeDeleteHooks = append(dbBeforeDeleteHooks, dbHook)
	case boil.BeforeUpsertHook:
		dbBeforeUpsertHooks = append(dbBeforeUpsertHooks, dbHook)
	case boil.AfterInsertHook:
		dbAfterInsertHooks = append(dbAfterInsertHooks, dbHook)
	case boil.AfterSelectHook:
		dbAfterSelectHooks = append(dbAfterSelectHooks, dbHook)
	case boil.AfterUpdateHook:
		dbAfterUpdateHooks = append(dbAfterUpdateHooks, dbHook)
	case boil.AfterDeleteHook:
		dbAfterDeleteHooks = append(dbAfterDeleteHooks, dbHook)
	case boil.AfterUpsertHook:
		dbAfterUpsertHooks = append(dbAfterUpsertHooks, dbHook)
	}
}

// OneP returns a single db record from the query, and panics on error.
func (q dbQuery) OneP() *DB {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single db record from the query.
func (q dbQuery) One() (*DB, error) {
	o := &DB{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for db")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all DB records from the query, and panics on error.
func (q dbQuery) AllP() DBSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DB records from the query.
func (q dbQuery) All() (DBSlice, error) {
	var o DBSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DB slice")
	}

	if len(dbAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all DB records in the query, and panics on error.
func (q dbQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DB records in the query.
func (q dbQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count db rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dbQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dbQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if db exists")
	}

	return count > 0, nil
}

// DbxrefG pointed to by the foreign key.
func (o *DB) DbxrefG(mods ...qm.QueryMod) dbxrefQuery {
	return o.Dbxref(boil.GetDB(), mods...)
}

// Dbxref pointed to by the foreign key.
func (o *DB) Dbxref(exec boil.Executor, mods ...qm.QueryMod) dbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("db_id=$1", o.DBID),
	}

	queryMods = append(queryMods, mods...)

	query := Dbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"dbxref\"")

	return query
}

// LoadDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbL) LoadDbxref(e boil.Executor, singular bool, maybeDB interface{}) error {
	var slice []*DB
	var object *DB

	count := 1
	if singular {
		object = maybeDB.(*DB)
	} else {
		slice = *maybeDB.(*DBSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbR{}
		args[0] = object.DBID
	} else {
		for i, obj := range slice {
			obj.R = &dbR{}
			args[i] = obj.DBID
		}
	}

	query := fmt.Sprintf(
		"select * from \"dbxref\" where \"db_id\" in (%s)",
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

	if len(dbAfterSelectHooks) != 0 {
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
			if local.DBID == foreign.DBID {
				local.R.Dbxref = foreign
				break
			}
		}
	}

	return nil
}

// SetDbxref of the db to the related item.
// Sets o.R.Dbxref to related.
// Adds o to related.R.DB.
func (o *DB) SetDbxref(exec boil.Executor, insert bool, related *Dbxref) error {
	var err error

	if insert {
		related.DBID = o.DBID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"db_id"}),
			strmangle.WhereClause("\"", "\"", 2, dbxrefPrimaryKeyColumns),
		)
		values := []interface{}{o.DBID, related.DbxrefID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.DBID = o.DBID

	}

	if o.R == nil {
		o.R = &dbR{
			Dbxref: related,
		}
	} else {
		o.R.Dbxref = related
	}

	if related.R == nil {
		related.R = &dbxrefR{
			DB: o,
		}
	} else {
		related.R.DB = o
	}
	return nil
}

// DBSG retrieves all records.
func DBSG(mods ...qm.QueryMod) dbQuery {
	return DBS(boil.GetDB(), mods...)
}

// DBS retrieves all the records using an executor.
func DBS(exec boil.Executor, mods ...qm.QueryMod) dbQuery {
	mods = append(mods, qm.From("\"db\""))
	return dbQuery{NewQuery(exec, mods...)}
}

// FindDBG retrieves a single record by ID.
func FindDBG(dbID int, selectCols ...string) (*DB, error) {
	return FindDB(boil.GetDB(), dbID, selectCols...)
}

// FindDBGP retrieves a single record by ID, and panics on error.
func FindDBGP(dbID int, selectCols ...string) *DB {
	retobj, err := FindDB(boil.GetDB(), dbID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDB retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDB(exec boil.Executor, dbID int, selectCols ...string) (*DB, error) {
	dbObj := &DB{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"db\" where \"db_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, dbID)

	err := q.Bind(dbObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from db")
	}

	return dbObj, nil
}

// FindDBP retrieves a single record by ID with an executor, and panics on error.
func FindDBP(exec boil.Executor, dbID int, selectCols ...string) *DB {
	retobj, err := FindDB(exec, dbID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DB) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DB) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DB) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DB) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no db provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dbColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dbInsertCacheMut.RLock()
	cache, cached := dbInsertCache[key]
	dbInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dbColumns,
			dbColumnsWithDefault,
			dbColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dbType, dbMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dbType, dbMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"db\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into db")
	}

	if !cached {
		dbInsertCacheMut.Lock()
		dbInsertCache[key] = cache
		dbInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single DB record. See Update for
// whitelist behavior description.
func (o *DB) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DB record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DB) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DB, and panics on error.
// See Update for whitelist behavior description.
func (o *DB) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DB.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DB) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	dbUpdateCacheMut.RLock()
	cache, cached := dbUpdateCache[key]
	dbUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dbColumns, dbPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update db, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"db\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dbPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dbType, dbMapping, append(wl, dbPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update db row")
	}

	if !cached {
		dbUpdateCacheMut.Lock()
		dbUpdateCache[key] = cache
		dbUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dbQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dbQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for db")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DBSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DBSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DBSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DBSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dbPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"db\" SET %s WHERE (\"db_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dbPrimaryKeyColumns), len(colNames)+1, len(dbPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in db slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DB) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DB) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DB) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DB) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no db provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dbColumnsWithDefault, o)

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

	dbUpsertCacheMut.RLock()
	cache, cached := dbUpsertCache[key]
	dbUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dbColumns,
			dbColumnsWithDefault,
			dbColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dbColumns,
			dbPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert db, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dbPrimaryKeyColumns))
			copy(conflict, dbPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"db\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dbType, dbMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dbType, dbMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for db")
	}

	if !cached {
		dbUpsertCacheMut.Lock()
		dbUpsertCache[key] = cache
		dbUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single DB record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DB) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DB record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DB) DeleteG() error {
	if o == nil {
		return errors.New("models: no DB provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DB record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DB) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DB record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DB) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DB provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dbPrimaryKeyMapping)
	sql := "DELETE FROM \"db\" WHERE \"db_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from db")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dbQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dbQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no dbQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from db")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DBSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DBSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DB slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DBSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DBSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DB slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(dbBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dbPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"db\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dbPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dbPrimaryKeyColumns), 1, len(dbPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from db slice")
	}

	if len(dbAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DB) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DB) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DB) ReloadG() error {
	if o == nil {
		return errors.New("models: no DB provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DB) Reload(exec boil.Executor) error {
	ret, err := FindDB(exec, o.DBID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DBSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DBSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DBSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DBSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DBSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dbs := DBSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dbPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"db\".* FROM \"db\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dbPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dbPrimaryKeyColumns), 1, len(dbPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dbs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DBSlice")
	}

	*o = dbs

	return nil
}

// DBExists checks if the DB row exists.
func DBExists(exec boil.Executor, dbID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"db\" where \"db_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, dbID)
	}

	row := exec.QueryRow(sql, dbID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if db exists")
	}

	return exists, nil
}

// DBExistsG checks if the DB row exists.
func DBExistsG(dbID int) (bool, error) {
	return DBExists(boil.GetDB(), dbID)
}

// DBExistsGP checks if the DB row exists. Panics on error.
func DBExistsGP(dbID int) bool {
	e, err := DBExists(boil.GetDB(), dbID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DBExistsP checks if the DB row exists. Panics on error.
func DBExistsP(exec boil.Executor, dbID int) bool {
	e, err := DBExists(exec, dbID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
