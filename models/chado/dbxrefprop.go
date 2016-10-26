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

// Dbxrefprop is an object representing the database table.
type Dbxrefprop struct {
	DbxrefpropID int    `boil:"dbxrefprop_id" json:"dbxrefprop_id" toml:"dbxrefprop_id" yaml:"dbxrefprop_id"`
	DbxrefID     int    `boil:"dbxref_id" json:"dbxref_id" toml:"dbxref_id" yaml:"dbxref_id"`
	TypeID       int    `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value        string `boil:"value" json:"value" toml:"value" yaml:"value"`
	Rank         int    `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *dbxrefpropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dbxrefpropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// dbxrefpropR is where relationships are stored.
type dbxrefpropR struct {
	Type   *Cvterm
	Dbxref *Dbxref
}

// dbxrefpropL is where Load methods for each relationship are stored.
type dbxrefpropL struct{}

var (
	dbxrefpropColumns               = []string{"dbxrefprop_id", "dbxref_id", "type_id", "value", "rank"}
	dbxrefpropColumnsWithoutDefault = []string{"dbxref_id", "type_id"}
	dbxrefpropColumnsWithDefault    = []string{"dbxrefprop_id", "value", "rank"}
	dbxrefpropPrimaryKeyColumns     = []string{"dbxrefprop_id"}
)

type (
	// DbxrefpropSlice is an alias for a slice of pointers to Dbxrefprop.
	// This should generally be used opposed to []Dbxrefprop.
	DbxrefpropSlice []*Dbxrefprop
	// DbxrefpropHook is the signature for custom Dbxrefprop hook methods
	DbxrefpropHook func(boil.Executor, *Dbxrefprop) error

	dbxrefpropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dbxrefpropType                 = reflect.TypeOf(&Dbxrefprop{})
	dbxrefpropMapping              = queries.MakeStructMapping(dbxrefpropType)
	dbxrefpropPrimaryKeyMapping, _ = queries.BindMapping(dbxrefpropType, dbxrefpropMapping, dbxrefpropPrimaryKeyColumns)
	dbxrefpropInsertCacheMut       sync.RWMutex
	dbxrefpropInsertCache          = make(map[string]insertCache)
	dbxrefpropUpdateCacheMut       sync.RWMutex
	dbxrefpropUpdateCache          = make(map[string]updateCache)
	dbxrefpropUpsertCacheMut       sync.RWMutex
	dbxrefpropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var dbxrefpropBeforeInsertHooks []DbxrefpropHook
var dbxrefpropBeforeUpdateHooks []DbxrefpropHook
var dbxrefpropBeforeDeleteHooks []DbxrefpropHook
var dbxrefpropBeforeUpsertHooks []DbxrefpropHook

var dbxrefpropAfterInsertHooks []DbxrefpropHook
var dbxrefpropAfterSelectHooks []DbxrefpropHook
var dbxrefpropAfterUpdateHooks []DbxrefpropHook
var dbxrefpropAfterDeleteHooks []DbxrefpropHook
var dbxrefpropAfterUpsertHooks []DbxrefpropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Dbxrefprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefpropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Dbxrefprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefpropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Dbxrefprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefpropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Dbxrefprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefpropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Dbxrefprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefpropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Dbxrefprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefpropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Dbxrefprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefpropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Dbxrefprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefpropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Dbxrefprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range dbxrefpropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDbxrefpropHook registers your hook function for all future operations.
func AddDbxrefpropHook(hookPoint boil.HookPoint, dbxrefpropHook DbxrefpropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		dbxrefpropBeforeInsertHooks = append(dbxrefpropBeforeInsertHooks, dbxrefpropHook)
	case boil.BeforeUpdateHook:
		dbxrefpropBeforeUpdateHooks = append(dbxrefpropBeforeUpdateHooks, dbxrefpropHook)
	case boil.BeforeDeleteHook:
		dbxrefpropBeforeDeleteHooks = append(dbxrefpropBeforeDeleteHooks, dbxrefpropHook)
	case boil.BeforeUpsertHook:
		dbxrefpropBeforeUpsertHooks = append(dbxrefpropBeforeUpsertHooks, dbxrefpropHook)
	case boil.AfterInsertHook:
		dbxrefpropAfterInsertHooks = append(dbxrefpropAfterInsertHooks, dbxrefpropHook)
	case boil.AfterSelectHook:
		dbxrefpropAfterSelectHooks = append(dbxrefpropAfterSelectHooks, dbxrefpropHook)
	case boil.AfterUpdateHook:
		dbxrefpropAfterUpdateHooks = append(dbxrefpropAfterUpdateHooks, dbxrefpropHook)
	case boil.AfterDeleteHook:
		dbxrefpropAfterDeleteHooks = append(dbxrefpropAfterDeleteHooks, dbxrefpropHook)
	case boil.AfterUpsertHook:
		dbxrefpropAfterUpsertHooks = append(dbxrefpropAfterUpsertHooks, dbxrefpropHook)
	}
}

// OneP returns a single dbxrefprop record from the query, and panics on error.
func (q dbxrefpropQuery) OneP() *Dbxrefprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single dbxrefprop record from the query.
func (q dbxrefpropQuery) One() (*Dbxrefprop, error) {
	o := &Dbxrefprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for dbxrefprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Dbxrefprop records from the query, and panics on error.
func (q dbxrefpropQuery) AllP() DbxrefpropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Dbxrefprop records from the query.
func (q dbxrefpropQuery) All() (DbxrefpropSlice, error) {
	var o DbxrefpropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Dbxrefprop slice")
	}

	if len(dbxrefpropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Dbxrefprop records in the query, and panics on error.
func (q dbxrefpropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Dbxrefprop records in the query.
func (q dbxrefpropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count dbxrefprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q dbxrefpropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q dbxrefpropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if dbxrefprop exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *Dbxrefprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Dbxrefprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// DbxrefG pointed to by the foreign key.
func (o *Dbxrefprop) DbxrefG(mods ...qm.QueryMod) dbxrefQuery {
	return o.Dbxref(boil.GetDB(), mods...)
}

// Dbxref pointed to by the foreign key.
func (o *Dbxrefprop) Dbxref(exec boil.Executor, mods ...qm.QueryMod) dbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := Dbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"dbxref\"")

	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbxrefpropL) LoadType(e boil.Executor, singular bool, maybeDbxrefprop interface{}) error {
	var slice []*Dbxrefprop
	var object *Dbxrefprop

	count := 1
	if singular {
		object = maybeDbxrefprop.(*Dbxrefprop)
	} else {
		slice = *maybeDbxrefprop.(*DbxrefpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbxrefpropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &dbxrefpropR{}
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

	if len(dbxrefpropAfterSelectHooks) != 0 {
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

// LoadDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (dbxrefpropL) LoadDbxref(e boil.Executor, singular bool, maybeDbxrefprop interface{}) error {
	var slice []*Dbxrefprop
	var object *Dbxrefprop

	count := 1
	if singular {
		object = maybeDbxrefprop.(*Dbxrefprop)
	} else {
		slice = *maybeDbxrefprop.(*DbxrefpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &dbxrefpropR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &dbxrefpropR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"dbxref\" where \"dbxref_id\" in (%s)",
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

	if len(dbxrefpropAfterSelectHooks) != 0 {
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
			if local.DbxrefID == foreign.DbxrefID {
				local.R.Dbxref = foreign
				break
			}
		}
	}

	return nil
}

// SetType of the dbxrefprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeDbxrefprop.
func (o *Dbxrefprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"dbxrefprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, dbxrefpropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.DbxrefpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &dbxrefpropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeDbxrefprop: o,
		}
	} else {
		related.R.TypeDbxrefprop = o
	}

	return nil
}

// SetDbxref of the dbxrefprop to the related item.
// Sets o.R.Dbxref to related.
// Adds o to related.R.Dbxrefprop.
func (o *Dbxrefprop) SetDbxref(exec boil.Executor, insert bool, related *Dbxref) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"dbxrefprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
		strmangle.WhereClause("\"", "\"", 2, dbxrefpropPrimaryKeyColumns),
	)
	values := []interface{}{related.DbxrefID, o.DbxrefpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DbxrefID = related.DbxrefID

	if o.R == nil {
		o.R = &dbxrefpropR{
			Dbxref: related,
		}
	} else {
		o.R.Dbxref = related
	}

	if related.R == nil {
		related.R = &dbxrefR{
			Dbxrefprop: o,
		}
	} else {
		related.R.Dbxrefprop = o
	}

	return nil
}

// DbxrefpropsG retrieves all records.
func DbxrefpropsG(mods ...qm.QueryMod) dbxrefpropQuery {
	return Dbxrefprops(boil.GetDB(), mods...)
}

// Dbxrefprops retrieves all the records using an executor.
func Dbxrefprops(exec boil.Executor, mods ...qm.QueryMod) dbxrefpropQuery {
	mods = append(mods, qm.From("\"dbxrefprop\""))
	return dbxrefpropQuery{NewQuery(exec, mods...)}
}

// FindDbxrefpropG retrieves a single record by ID.
func FindDbxrefpropG(dbxrefpropID int, selectCols ...string) (*Dbxrefprop, error) {
	return FindDbxrefprop(boil.GetDB(), dbxrefpropID, selectCols...)
}

// FindDbxrefpropGP retrieves a single record by ID, and panics on error.
func FindDbxrefpropGP(dbxrefpropID int, selectCols ...string) *Dbxrefprop {
	retobj, err := FindDbxrefprop(boil.GetDB(), dbxrefpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDbxrefprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDbxrefprop(exec boil.Executor, dbxrefpropID int, selectCols ...string) (*Dbxrefprop, error) {
	dbxrefpropObj := &Dbxrefprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"dbxrefprop\" where \"dbxrefprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, dbxrefpropID)

	err := q.Bind(dbxrefpropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from dbxrefprop")
	}

	return dbxrefpropObj, nil
}

// FindDbxrefpropP retrieves a single record by ID with an executor, and panics on error.
func FindDbxrefpropP(exec boil.Executor, dbxrefpropID int, selectCols ...string) *Dbxrefprop {
	retobj, err := FindDbxrefprop(exec, dbxrefpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Dbxrefprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Dbxrefprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Dbxrefprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Dbxrefprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no dbxrefprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dbxrefpropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	dbxrefpropInsertCacheMut.RLock()
	cache, cached := dbxrefpropInsertCache[key]
	dbxrefpropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			dbxrefpropColumns,
			dbxrefpropColumnsWithDefault,
			dbxrefpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(dbxrefpropType, dbxrefpropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dbxrefpropType, dbxrefpropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"dbxrefprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into dbxrefprop")
	}

	if !cached {
		dbxrefpropInsertCacheMut.Lock()
		dbxrefpropInsertCache[key] = cache
		dbxrefpropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Dbxrefprop record. See Update for
// whitelist behavior description.
func (o *Dbxrefprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Dbxrefprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Dbxrefprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Dbxrefprop, and panics on error.
// See Update for whitelist behavior description.
func (o *Dbxrefprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Dbxrefprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Dbxrefprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	dbxrefpropUpdateCacheMut.RLock()
	cache, cached := dbxrefpropUpdateCache[key]
	dbxrefpropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(dbxrefpropColumns, dbxrefpropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update dbxrefprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"dbxrefprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dbxrefpropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dbxrefpropType, dbxrefpropMapping, append(wl, dbxrefpropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update dbxrefprop row")
	}

	if !cached {
		dbxrefpropUpdateCacheMut.Lock()
		dbxrefpropUpdateCache[key] = cache
		dbxrefpropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q dbxrefpropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q dbxrefpropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for dbxrefprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DbxrefpropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DbxrefpropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DbxrefpropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DbxrefpropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dbxrefpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"dbxrefprop\" SET %s WHERE (\"dbxrefprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dbxrefpropPrimaryKeyColumns), len(colNames)+1, len(dbxrefpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in dbxrefprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Dbxrefprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Dbxrefprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Dbxrefprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Dbxrefprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no dbxrefprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dbxrefpropColumnsWithDefault, o)

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

	dbxrefpropUpsertCacheMut.RLock()
	cache, cached := dbxrefpropUpsertCache[key]
	dbxrefpropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			dbxrefpropColumns,
			dbxrefpropColumnsWithDefault,
			dbxrefpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			dbxrefpropColumns,
			dbxrefpropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert dbxrefprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dbxrefpropPrimaryKeyColumns))
			copy(conflict, dbxrefpropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"dbxrefprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(dbxrefpropType, dbxrefpropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dbxrefpropType, dbxrefpropMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for dbxrefprop")
	}

	if !cached {
		dbxrefpropUpsertCacheMut.Lock()
		dbxrefpropUpsertCache[key] = cache
		dbxrefpropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Dbxrefprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Dbxrefprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Dbxrefprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Dbxrefprop) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Dbxrefprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Dbxrefprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Dbxrefprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Dbxrefprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Dbxrefprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Dbxrefprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dbxrefpropPrimaryKeyMapping)
	sql := "DELETE FROM \"dbxrefprop\" WHERE \"dbxrefprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from dbxrefprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q dbxrefpropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q dbxrefpropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no dbxrefpropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from dbxrefprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DbxrefpropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DbxrefpropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Dbxrefprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DbxrefpropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DbxrefpropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Dbxrefprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(dbxrefpropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dbxrefpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"dbxrefprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dbxrefpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(dbxrefpropPrimaryKeyColumns), 1, len(dbxrefpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from dbxrefprop slice")
	}

	if len(dbxrefpropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Dbxrefprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Dbxrefprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Dbxrefprop) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Dbxrefprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Dbxrefprop) Reload(exec boil.Executor) error {
	ret, err := FindDbxrefprop(exec, o.DbxrefpropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DbxrefpropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DbxrefpropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DbxrefpropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty DbxrefpropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DbxrefpropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	dbxrefprops := DbxrefpropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dbxrefpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"dbxrefprop\".* FROM \"dbxrefprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, dbxrefpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(dbxrefpropPrimaryKeyColumns), 1, len(dbxrefpropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&dbxrefprops)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in DbxrefpropSlice")
	}

	*o = dbxrefprops

	return nil
}

// DbxrefpropExists checks if the Dbxrefprop row exists.
func DbxrefpropExists(exec boil.Executor, dbxrefpropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"dbxrefprop\" where \"dbxrefprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, dbxrefpropID)
	}

	row := exec.QueryRow(sql, dbxrefpropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if dbxrefprop exists")
	}

	return exists, nil
}

// DbxrefpropExistsG checks if the Dbxrefprop row exists.
func DbxrefpropExistsG(dbxrefpropID int) (bool, error) {
	return DbxrefpropExists(boil.GetDB(), dbxrefpropID)
}

// DbxrefpropExistsGP checks if the Dbxrefprop row exists. Panics on error.
func DbxrefpropExistsGP(dbxrefpropID int) bool {
	e, err := DbxrefpropExists(boil.GetDB(), dbxrefpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DbxrefpropExistsP checks if the Dbxrefprop row exists. Panics on error.
func DbxrefpropExistsP(exec boil.Executor, dbxrefpropID int) bool {
	e, err := DbxrefpropExists(exec, dbxrefpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
