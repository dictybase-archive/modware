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
	"github.com/vattle/sqlboiler/types"
)

// Jbrowse is an object representing the database table.
type Jbrowse struct {
	JbrowseID     int        `boil:"jbrowse_id" json:"jbrowse_id" toml:"jbrowse_id" yaml:"jbrowse_id"`
	Name          string     `boil:"name" json:"name" toml:"name" yaml:"name"`
	Configuration types.JSON `boil:"configuration" json:"configuration" toml:"configuration" yaml:"configuration"`

	R *jbrowseR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L jbrowseL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// jbrowseR is where relationships are stored.
type jbrowseR struct {
	JbrowseOrganism *JbrowseOrganism
}

// jbrowseL is where Load methods for each relationship are stored.
type jbrowseL struct{}

var (
	jbrowseColumns               = []string{"jbrowse_id", "name", "configuration"}
	jbrowseColumnsWithoutDefault = []string{"name", "configuration"}
	jbrowseColumnsWithDefault    = []string{"jbrowse_id"}
	jbrowsePrimaryKeyColumns     = []string{"jbrowse_id"}
)

type (
	// JbrowseSlice is an alias for a slice of pointers to Jbrowse.
	// This should generally be used opposed to []Jbrowse.
	JbrowseSlice []*Jbrowse
	// JbrowseHook is the signature for custom Jbrowse hook methods
	JbrowseHook func(boil.Executor, *Jbrowse) error

	jbrowseQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	jbrowseType                 = reflect.TypeOf(&Jbrowse{})
	jbrowseMapping              = queries.MakeStructMapping(jbrowseType)
	jbrowsePrimaryKeyMapping, _ = queries.BindMapping(jbrowseType, jbrowseMapping, jbrowsePrimaryKeyColumns)
	jbrowseInsertCacheMut       sync.RWMutex
	jbrowseInsertCache          = make(map[string]insertCache)
	jbrowseUpdateCacheMut       sync.RWMutex
	jbrowseUpdateCache          = make(map[string]updateCache)
	jbrowseUpsertCacheMut       sync.RWMutex
	jbrowseUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var jbrowseBeforeInsertHooks []JbrowseHook
var jbrowseBeforeUpdateHooks []JbrowseHook
var jbrowseBeforeDeleteHooks []JbrowseHook
var jbrowseBeforeUpsertHooks []JbrowseHook

var jbrowseAfterInsertHooks []JbrowseHook
var jbrowseAfterSelectHooks []JbrowseHook
var jbrowseAfterUpdateHooks []JbrowseHook
var jbrowseAfterDeleteHooks []JbrowseHook
var jbrowseAfterUpsertHooks []JbrowseHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Jbrowse) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Jbrowse) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Jbrowse) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Jbrowse) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Jbrowse) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Jbrowse) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Jbrowse) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Jbrowse) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Jbrowse) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddJbrowseHook registers your hook function for all future operations.
func AddJbrowseHook(hookPoint boil.HookPoint, jbrowseHook JbrowseHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		jbrowseBeforeInsertHooks = append(jbrowseBeforeInsertHooks, jbrowseHook)
	case boil.BeforeUpdateHook:
		jbrowseBeforeUpdateHooks = append(jbrowseBeforeUpdateHooks, jbrowseHook)
	case boil.BeforeDeleteHook:
		jbrowseBeforeDeleteHooks = append(jbrowseBeforeDeleteHooks, jbrowseHook)
	case boil.BeforeUpsertHook:
		jbrowseBeforeUpsertHooks = append(jbrowseBeforeUpsertHooks, jbrowseHook)
	case boil.AfterInsertHook:
		jbrowseAfterInsertHooks = append(jbrowseAfterInsertHooks, jbrowseHook)
	case boil.AfterSelectHook:
		jbrowseAfterSelectHooks = append(jbrowseAfterSelectHooks, jbrowseHook)
	case boil.AfterUpdateHook:
		jbrowseAfterUpdateHooks = append(jbrowseAfterUpdateHooks, jbrowseHook)
	case boil.AfterDeleteHook:
		jbrowseAfterDeleteHooks = append(jbrowseAfterDeleteHooks, jbrowseHook)
	case boil.AfterUpsertHook:
		jbrowseAfterUpsertHooks = append(jbrowseAfterUpsertHooks, jbrowseHook)
	}
}

// OneP returns a single jbrowse record from the query, and panics on error.
func (q jbrowseQuery) OneP() *Jbrowse {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single jbrowse record from the query.
func (q jbrowseQuery) One() (*Jbrowse, error) {
	o := &Jbrowse{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for jbrowse")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Jbrowse records from the query, and panics on error.
func (q jbrowseQuery) AllP() JbrowseSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Jbrowse records from the query.
func (q jbrowseQuery) All() (JbrowseSlice, error) {
	var o JbrowseSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Jbrowse slice")
	}

	if len(jbrowseAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Jbrowse records in the query, and panics on error.
func (q jbrowseQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Jbrowse records in the query.
func (q jbrowseQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count jbrowse rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q jbrowseQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q jbrowseQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if jbrowse exists")
	}

	return count > 0, nil
}

// JbrowseOrganismG pointed to by the foreign key.
func (o *Jbrowse) JbrowseOrganismG(mods ...qm.QueryMod) jbrowseOrganismQuery {
	return o.JbrowseOrganism(boil.GetDB(), mods...)
}

// JbrowseOrganism pointed to by the foreign key.
func (o *Jbrowse) JbrowseOrganism(exec boil.Executor, mods ...qm.QueryMod) jbrowseOrganismQuery {
	queryMods := []qm.QueryMod{
		qm.Where("jbrowse_id=$1", o.JbrowseID),
	}

	queryMods = append(queryMods, mods...)

	query := JbrowseOrganisms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"jbrowse_organism\"")

	return query
}

// LoadJbrowseOrganism allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (jbrowseL) LoadJbrowseOrganism(e boil.Executor, singular bool, maybeJbrowse interface{}) error {
	var slice []*Jbrowse
	var object *Jbrowse

	count := 1
	if singular {
		object = maybeJbrowse.(*Jbrowse)
	} else {
		slice = *maybeJbrowse.(*JbrowseSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &jbrowseR{}
		args[0] = object.JbrowseID
	} else {
		for i, obj := range slice {
			obj.R = &jbrowseR{}
			args[i] = obj.JbrowseID
		}
	}

	query := fmt.Sprintf(
		"select * from \"jbrowse_organism\" where \"jbrowse_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load JbrowseOrganism")
	}
	defer results.Close()

	var resultSlice []*JbrowseOrganism
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice JbrowseOrganism")
	}

	if len(jbrowseAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.JbrowseOrganism = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.JbrowseID == foreign.JbrowseID {
				local.R.JbrowseOrganism = foreign
				break
			}
		}
	}

	return nil
}

// SetJbrowseOrganism of the jbrowse to the related item.
// Sets o.R.JbrowseOrganism to related.
// Adds o to related.R.Jbrowse.
func (o *Jbrowse) SetJbrowseOrganism(exec boil.Executor, insert bool, related *JbrowseOrganism) error {
	var err error

	if insert {
		related.JbrowseID = o.JbrowseID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"jbrowse_organism\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"jbrowse_id"}),
			strmangle.WhereClause("\"", "\"", 2, jbrowseOrganismPrimaryKeyColumns),
		)
		values := []interface{}{o.JbrowseID, related.JbrowseOrganismID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.JbrowseID = o.JbrowseID

	}

	if o.R == nil {
		o.R = &jbrowseR{
			JbrowseOrganism: related,
		}
	} else {
		o.R.JbrowseOrganism = related
	}

	if related.R == nil {
		related.R = &jbrowseOrganismR{
			Jbrowse: o,
		}
	} else {
		related.R.Jbrowse = o
	}
	return nil
}

// JbrowsesG retrieves all records.
func JbrowsesG(mods ...qm.QueryMod) jbrowseQuery {
	return Jbrowses(boil.GetDB(), mods...)
}

// Jbrowses retrieves all the records using an executor.
func Jbrowses(exec boil.Executor, mods ...qm.QueryMod) jbrowseQuery {
	mods = append(mods, qm.From("\"jbrowse\""))
	return jbrowseQuery{NewQuery(exec, mods...)}
}

// FindJbrowseG retrieves a single record by ID.
func FindJbrowseG(jbrowseID int, selectCols ...string) (*Jbrowse, error) {
	return FindJbrowse(boil.GetDB(), jbrowseID, selectCols...)
}

// FindJbrowseGP retrieves a single record by ID, and panics on error.
func FindJbrowseGP(jbrowseID int, selectCols ...string) *Jbrowse {
	retobj, err := FindJbrowse(boil.GetDB(), jbrowseID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindJbrowse retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindJbrowse(exec boil.Executor, jbrowseID int, selectCols ...string) (*Jbrowse, error) {
	jbrowseObj := &Jbrowse{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"jbrowse\" where \"jbrowse_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, jbrowseID)

	err := q.Bind(jbrowseObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from jbrowse")
	}

	return jbrowseObj, nil
}

// FindJbrowseP retrieves a single record by ID with an executor, and panics on error.
func FindJbrowseP(exec boil.Executor, jbrowseID int, selectCols ...string) *Jbrowse {
	retobj, err := FindJbrowse(exec, jbrowseID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Jbrowse) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Jbrowse) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Jbrowse) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Jbrowse) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no jbrowse provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(jbrowseColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	jbrowseInsertCacheMut.RLock()
	cache, cached := jbrowseInsertCache[key]
	jbrowseInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			jbrowseColumns,
			jbrowseColumnsWithDefault,
			jbrowseColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(jbrowseType, jbrowseMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(jbrowseType, jbrowseMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"jbrowse\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into jbrowse")
	}

	if !cached {
		jbrowseInsertCacheMut.Lock()
		jbrowseInsertCache[key] = cache
		jbrowseInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Jbrowse record. See Update for
// whitelist behavior description.
func (o *Jbrowse) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Jbrowse record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Jbrowse) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Jbrowse, and panics on error.
// See Update for whitelist behavior description.
func (o *Jbrowse) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Jbrowse.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Jbrowse) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	jbrowseUpdateCacheMut.RLock()
	cache, cached := jbrowseUpdateCache[key]
	jbrowseUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(jbrowseColumns, jbrowsePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update jbrowse, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"jbrowse\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, jbrowsePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(jbrowseType, jbrowseMapping, append(wl, jbrowsePrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update jbrowse row")
	}

	if !cached {
		jbrowseUpdateCacheMut.Lock()
		jbrowseUpdateCache[key] = cache
		jbrowseUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q jbrowseQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q jbrowseQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for jbrowse")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o JbrowseSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o JbrowseSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o JbrowseSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o JbrowseSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jbrowsePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"jbrowse\" SET %s WHERE (\"jbrowse_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(jbrowsePrimaryKeyColumns), len(colNames)+1, len(jbrowsePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in jbrowse slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Jbrowse) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Jbrowse) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Jbrowse) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Jbrowse) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no jbrowse provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(jbrowseColumnsWithDefault, o)

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

	jbrowseUpsertCacheMut.RLock()
	cache, cached := jbrowseUpsertCache[key]
	jbrowseUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			jbrowseColumns,
			jbrowseColumnsWithDefault,
			jbrowseColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			jbrowseColumns,
			jbrowsePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert jbrowse, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(jbrowsePrimaryKeyColumns))
			copy(conflict, jbrowsePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"jbrowse\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(jbrowseType, jbrowseMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(jbrowseType, jbrowseMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for jbrowse")
	}

	if !cached {
		jbrowseUpsertCacheMut.Lock()
		jbrowseUpsertCache[key] = cache
		jbrowseUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Jbrowse record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Jbrowse) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Jbrowse record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Jbrowse) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Jbrowse provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Jbrowse record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Jbrowse) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Jbrowse record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Jbrowse) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Jbrowse provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), jbrowsePrimaryKeyMapping)
	sql := "DELETE FROM \"jbrowse\" WHERE \"jbrowse_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from jbrowse")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q jbrowseQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q jbrowseQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no jbrowseQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from jbrowse")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o JbrowseSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o JbrowseSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Jbrowse slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o JbrowseSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o JbrowseSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Jbrowse slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(jbrowseBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jbrowsePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"jbrowse\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, jbrowsePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(jbrowsePrimaryKeyColumns), 1, len(jbrowsePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from jbrowse slice")
	}

	if len(jbrowseAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Jbrowse) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Jbrowse) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Jbrowse) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Jbrowse provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Jbrowse) Reload(exec boil.Executor) error {
	ret, err := FindJbrowse(exec, o.JbrowseID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *JbrowseSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *JbrowseSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JbrowseSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty JbrowseSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JbrowseSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	jbrowses := JbrowseSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jbrowsePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"jbrowse\".* FROM \"jbrowse\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, jbrowsePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(jbrowsePrimaryKeyColumns), 1, len(jbrowsePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&jbrowses)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in JbrowseSlice")
	}

	*o = jbrowses

	return nil
}

// JbrowseExists checks if the Jbrowse row exists.
func JbrowseExists(exec boil.Executor, jbrowseID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"jbrowse\" where \"jbrowse_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, jbrowseID)
	}

	row := exec.QueryRow(sql, jbrowseID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if jbrowse exists")
	}

	return exists, nil
}

// JbrowseExistsG checks if the Jbrowse row exists.
func JbrowseExistsG(jbrowseID int) (bool, error) {
	return JbrowseExists(boil.GetDB(), jbrowseID)
}

// JbrowseExistsGP checks if the Jbrowse row exists. Panics on error.
func JbrowseExistsGP(jbrowseID int) bool {
	e, err := JbrowseExists(boil.GetDB(), jbrowseID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// JbrowseExistsP checks if the Jbrowse row exists. Panics on error.
func JbrowseExistsP(exec boil.Executor, jbrowseID int) bool {
	e, err := JbrowseExists(exec, jbrowseID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
