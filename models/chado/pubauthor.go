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
	"gopkg.in/nullbio/null.v5"
)

// Pubauthor is an object representing the database table.
type Pubauthor struct {
	PubauthorID int         `boil:"pubauthor_id" json:"pubauthor_id" toml:"pubauthor_id" yaml:"pubauthor_id"`
	PubID       int         `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`
	Rank        int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`
	Editor      null.Bool   `boil:"editor" json:"editor,omitempty" toml:"editor" yaml:"editor,omitempty"`
	Surname     string      `boil:"surname" json:"surname" toml:"surname" yaml:"surname"`
	Givennames  null.String `boil:"givennames" json:"givennames,omitempty" toml:"givennames" yaml:"givennames,omitempty"`
	Suffix      null.String `boil:"suffix" json:"suffix,omitempty" toml:"suffix" yaml:"suffix,omitempty"`

	R *pubauthorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pubauthorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// pubauthorR is where relationships are stored.
type pubauthorR struct {
	Pub *Pub
}

// pubauthorL is where Load methods for each relationship are stored.
type pubauthorL struct{}

var (
	pubauthorColumns               = []string{"pubauthor_id", "pub_id", "rank", "editor", "surname", "givennames", "suffix"}
	pubauthorColumnsWithoutDefault = []string{"pub_id", "rank", "surname", "givennames", "suffix"}
	pubauthorColumnsWithDefault    = []string{"pubauthor_id", "editor"}
	pubauthorPrimaryKeyColumns     = []string{"pubauthor_id"}
)

type (
	// PubauthorSlice is an alias for a slice of pointers to Pubauthor.
	// This should generally be used opposed to []Pubauthor.
	PubauthorSlice []*Pubauthor
	// PubauthorHook is the signature for custom Pubauthor hook methods
	PubauthorHook func(boil.Executor, *Pubauthor) error

	pubauthorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pubauthorType                 = reflect.TypeOf(&Pubauthor{})
	pubauthorMapping              = queries.MakeStructMapping(pubauthorType)
	pubauthorPrimaryKeyMapping, _ = queries.BindMapping(pubauthorType, pubauthorMapping, pubauthorPrimaryKeyColumns)
	pubauthorInsertCacheMut       sync.RWMutex
	pubauthorInsertCache          = make(map[string]insertCache)
	pubauthorUpdateCacheMut       sync.RWMutex
	pubauthorUpdateCache          = make(map[string]updateCache)
	pubauthorUpsertCacheMut       sync.RWMutex
	pubauthorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var pubauthorBeforeInsertHooks []PubauthorHook
var pubauthorBeforeUpdateHooks []PubauthorHook
var pubauthorBeforeDeleteHooks []PubauthorHook
var pubauthorBeforeUpsertHooks []PubauthorHook

var pubauthorAfterInsertHooks []PubauthorHook
var pubauthorAfterSelectHooks []PubauthorHook
var pubauthorAfterUpdateHooks []PubauthorHook
var pubauthorAfterDeleteHooks []PubauthorHook
var pubauthorAfterUpsertHooks []PubauthorHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Pubauthor) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range pubauthorBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Pubauthor) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range pubauthorBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Pubauthor) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range pubauthorBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Pubauthor) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range pubauthorBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Pubauthor) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range pubauthorAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Pubauthor) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range pubauthorAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Pubauthor) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range pubauthorAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Pubauthor) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range pubauthorAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Pubauthor) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range pubauthorAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPubauthorHook registers your hook function for all future operations.
func AddPubauthorHook(hookPoint boil.HookPoint, pubauthorHook PubauthorHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		pubauthorBeforeInsertHooks = append(pubauthorBeforeInsertHooks, pubauthorHook)
	case boil.BeforeUpdateHook:
		pubauthorBeforeUpdateHooks = append(pubauthorBeforeUpdateHooks, pubauthorHook)
	case boil.BeforeDeleteHook:
		pubauthorBeforeDeleteHooks = append(pubauthorBeforeDeleteHooks, pubauthorHook)
	case boil.BeforeUpsertHook:
		pubauthorBeforeUpsertHooks = append(pubauthorBeforeUpsertHooks, pubauthorHook)
	case boil.AfterInsertHook:
		pubauthorAfterInsertHooks = append(pubauthorAfterInsertHooks, pubauthorHook)
	case boil.AfterSelectHook:
		pubauthorAfterSelectHooks = append(pubauthorAfterSelectHooks, pubauthorHook)
	case boil.AfterUpdateHook:
		pubauthorAfterUpdateHooks = append(pubauthorAfterUpdateHooks, pubauthorHook)
	case boil.AfterDeleteHook:
		pubauthorAfterDeleteHooks = append(pubauthorAfterDeleteHooks, pubauthorHook)
	case boil.AfterUpsertHook:
		pubauthorAfterUpsertHooks = append(pubauthorAfterUpsertHooks, pubauthorHook)
	}
}

// OneP returns a single pubauthor record from the query, and panics on error.
func (q pubauthorQuery) OneP() *Pubauthor {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single pubauthor record from the query.
func (q pubauthorQuery) One() (*Pubauthor, error) {
	o := &Pubauthor{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for pubauthor")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Pubauthor records from the query, and panics on error.
func (q pubauthorQuery) AllP() PubauthorSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Pubauthor records from the query.
func (q pubauthorQuery) All() (PubauthorSlice, error) {
	var o PubauthorSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Pubauthor slice")
	}

	if len(pubauthorAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Pubauthor records in the query, and panics on error.
func (q pubauthorQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Pubauthor records in the query.
func (q pubauthorQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count pubauthor rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q pubauthorQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q pubauthorQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if pubauthor exists")
	}

	return count > 0, nil
}

// PubG pointed to by the foreign key.
func (o *Pubauthor) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *Pubauthor) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubauthorL) LoadPub(e boil.Executor, singular bool, maybePubauthor interface{}) error {
	var slice []*Pubauthor
	var object *Pubauthor

	count := 1
	if singular {
		object = maybePubauthor.(*Pubauthor)
	} else {
		slice = *maybePubauthor.(*PubauthorSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubauthorR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubauthorR{}
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

	if len(pubauthorAfterSelectHooks) != 0 {
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

// SetPub of the pubauthor to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.Pubauthor.
func (o *Pubauthor) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"pubauthor\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, pubauthorPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.PubauthorID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &pubauthorR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			Pubauthor: o,
		}
	} else {
		related.R.Pubauthor = o
	}

	return nil
}

// PubauthorsG retrieves all records.
func PubauthorsG(mods ...qm.QueryMod) pubauthorQuery {
	return Pubauthors(boil.GetDB(), mods...)
}

// Pubauthors retrieves all the records using an executor.
func Pubauthors(exec boil.Executor, mods ...qm.QueryMod) pubauthorQuery {
	mods = append(mods, qm.From("\"pubauthor\""))
	return pubauthorQuery{NewQuery(exec, mods...)}
}

// FindPubauthorG retrieves a single record by ID.
func FindPubauthorG(pubauthorID int, selectCols ...string) (*Pubauthor, error) {
	return FindPubauthor(boil.GetDB(), pubauthorID, selectCols...)
}

// FindPubauthorGP retrieves a single record by ID, and panics on error.
func FindPubauthorGP(pubauthorID int, selectCols ...string) *Pubauthor {
	retobj, err := FindPubauthor(boil.GetDB(), pubauthorID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindPubauthor retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPubauthor(exec boil.Executor, pubauthorID int, selectCols ...string) (*Pubauthor, error) {
	pubauthorObj := &Pubauthor{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pubauthor\" where \"pubauthor_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, pubauthorID)

	err := q.Bind(pubauthorObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from pubauthor")
	}

	return pubauthorObj, nil
}

// FindPubauthorP retrieves a single record by ID with an executor, and panics on error.
func FindPubauthorP(exec boil.Executor, pubauthorID int, selectCols ...string) *Pubauthor {
	retobj, err := FindPubauthor(exec, pubauthorID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Pubauthor) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Pubauthor) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Pubauthor) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Pubauthor) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no pubauthor provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pubauthorColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	pubauthorInsertCacheMut.RLock()
	cache, cached := pubauthorInsertCache[key]
	pubauthorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			pubauthorColumns,
			pubauthorColumnsWithDefault,
			pubauthorColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(pubauthorType, pubauthorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pubauthorType, pubauthorMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"pubauthor\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into pubauthor")
	}

	if !cached {
		pubauthorInsertCacheMut.Lock()
		pubauthorInsertCache[key] = cache
		pubauthorInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Pubauthor record. See Update for
// whitelist behavior description.
func (o *Pubauthor) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Pubauthor record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Pubauthor) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Pubauthor, and panics on error.
// See Update for whitelist behavior description.
func (o *Pubauthor) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Pubauthor.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Pubauthor) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	pubauthorUpdateCacheMut.RLock()
	cache, cached := pubauthorUpdateCache[key]
	pubauthorUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(pubauthorColumns, pubauthorPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update pubauthor, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pubauthor\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pubauthorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pubauthorType, pubauthorMapping, append(wl, pubauthorPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update pubauthor row")
	}

	if !cached {
		pubauthorUpdateCacheMut.Lock()
		pubauthorUpdateCache[key] = cache
		pubauthorUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q pubauthorQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q pubauthorQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for pubauthor")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PubauthorSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o PubauthorSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o PubauthorSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PubauthorSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pubauthorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"pubauthor\" SET %s WHERE (\"pubauthor_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(pubauthorPrimaryKeyColumns), len(colNames)+1, len(pubauthorPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in pubauthor slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Pubauthor) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Pubauthor) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Pubauthor) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Pubauthor) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no pubauthor provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pubauthorColumnsWithDefault, o)

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

	pubauthorUpsertCacheMut.RLock()
	cache, cached := pubauthorUpsertCache[key]
	pubauthorUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			pubauthorColumns,
			pubauthorColumnsWithDefault,
			pubauthorColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			pubauthorColumns,
			pubauthorPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert pubauthor, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pubauthorPrimaryKeyColumns))
			copy(conflict, pubauthorPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"pubauthor\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(pubauthorType, pubauthorMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pubauthorType, pubauthorMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for pubauthor")
	}

	if !cached {
		pubauthorUpsertCacheMut.Lock()
		pubauthorUpsertCache[key] = cache
		pubauthorUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Pubauthor record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Pubauthor) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Pubauthor record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Pubauthor) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Pubauthor provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Pubauthor record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Pubauthor) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Pubauthor record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Pubauthor) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Pubauthor provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pubauthorPrimaryKeyMapping)
	sql := "DELETE FROM \"pubauthor\" WHERE \"pubauthor_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from pubauthor")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q pubauthorQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q pubauthorQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no pubauthorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from pubauthor")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o PubauthorSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o PubauthorSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Pubauthor slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o PubauthorSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PubauthorSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Pubauthor slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(pubauthorBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pubauthorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"pubauthor\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, pubauthorPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(pubauthorPrimaryKeyColumns), 1, len(pubauthorPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from pubauthor slice")
	}

	if len(pubauthorAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Pubauthor) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Pubauthor) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Pubauthor) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Pubauthor provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Pubauthor) Reload(exec boil.Executor) error {
	ret, err := FindPubauthor(exec, o.PubauthorID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PubauthorSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PubauthorSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PubauthorSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty PubauthorSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PubauthorSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	pubauthors := PubauthorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pubauthorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"pubauthor\".* FROM \"pubauthor\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, pubauthorPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(pubauthorPrimaryKeyColumns), 1, len(pubauthorPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&pubauthors)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in PubauthorSlice")
	}

	*o = pubauthors

	return nil
}

// PubauthorExists checks if the Pubauthor row exists.
func PubauthorExists(exec boil.Executor, pubauthorID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"pubauthor\" where \"pubauthor_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, pubauthorID)
	}

	row := exec.QueryRow(sql, pubauthorID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if pubauthor exists")
	}

	return exists, nil
}

// PubauthorExistsG checks if the Pubauthor row exists.
func PubauthorExistsG(pubauthorID int) (bool, error) {
	return PubauthorExists(boil.GetDB(), pubauthorID)
}

// PubauthorExistsGP checks if the Pubauthor row exists. Panics on error.
func PubauthorExistsGP(pubauthorID int) bool {
	e, err := PubauthorExists(boil.GetDB(), pubauthorID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// PubauthorExistsP checks if the Pubauthor row exists. Panics on error.
func PubauthorExistsP(exec boil.Executor, pubauthorID int) bool {
	e, err := PubauthorExists(exec, pubauthorID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
