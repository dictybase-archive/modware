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

// PubDbxref is an object representing the database table.
type PubDbxref struct {
	PubDbxrefID int  `boil:"pub_dbxref_id" json:"pub_dbxref_id" toml:"pub_dbxref_id" yaml:"pub_dbxref_id"`
	PubID       int  `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`
	DbxrefID    int  `boil:"dbxref_id" json:"dbxref_id" toml:"dbxref_id" yaml:"dbxref_id"`
	IsCurrent   bool `boil:"is_current" json:"is_current" toml:"is_current" yaml:"is_current"`

	R *pubDbxrefR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pubDbxrefL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// pubDbxrefR is where relationships are stored.
type pubDbxrefR struct {
	Pub    *Pub
	Dbxref *Dbxref
}

// pubDbxrefL is where Load methods for each relationship are stored.
type pubDbxrefL struct{}

var (
	pubDbxrefColumns               = []string{"pub_dbxref_id", "pub_id", "dbxref_id", "is_current"}
	pubDbxrefColumnsWithoutDefault = []string{"pub_id", "dbxref_id"}
	pubDbxrefColumnsWithDefault    = []string{"pub_dbxref_id", "is_current"}
	pubDbxrefPrimaryKeyColumns     = []string{"pub_dbxref_id"}
)

type (
	// PubDbxrefSlice is an alias for a slice of pointers to PubDbxref.
	// This should generally be used opposed to []PubDbxref.
	PubDbxrefSlice []*PubDbxref
	// PubDbxrefHook is the signature for custom PubDbxref hook methods
	PubDbxrefHook func(boil.Executor, *PubDbxref) error

	pubDbxrefQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pubDbxrefType                 = reflect.TypeOf(&PubDbxref{})
	pubDbxrefMapping              = queries.MakeStructMapping(pubDbxrefType)
	pubDbxrefPrimaryKeyMapping, _ = queries.BindMapping(pubDbxrefType, pubDbxrefMapping, pubDbxrefPrimaryKeyColumns)
	pubDbxrefInsertCacheMut       sync.RWMutex
	pubDbxrefInsertCache          = make(map[string]insertCache)
	pubDbxrefUpdateCacheMut       sync.RWMutex
	pubDbxrefUpdateCache          = make(map[string]updateCache)
	pubDbxrefUpsertCacheMut       sync.RWMutex
	pubDbxrefUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var pubDbxrefBeforeInsertHooks []PubDbxrefHook
var pubDbxrefBeforeUpdateHooks []PubDbxrefHook
var pubDbxrefBeforeDeleteHooks []PubDbxrefHook
var pubDbxrefBeforeUpsertHooks []PubDbxrefHook

var pubDbxrefAfterInsertHooks []PubDbxrefHook
var pubDbxrefAfterSelectHooks []PubDbxrefHook
var pubDbxrefAfterUpdateHooks []PubDbxrefHook
var pubDbxrefAfterDeleteHooks []PubDbxrefHook
var pubDbxrefAfterUpsertHooks []PubDbxrefHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PubDbxref) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range pubDbxrefBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PubDbxref) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range pubDbxrefBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PubDbxref) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range pubDbxrefBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PubDbxref) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range pubDbxrefBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PubDbxref) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range pubDbxrefAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PubDbxref) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range pubDbxrefAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PubDbxref) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range pubDbxrefAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PubDbxref) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range pubDbxrefAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PubDbxref) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range pubDbxrefAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPubDbxrefHook registers your hook function for all future operations.
func AddPubDbxrefHook(hookPoint boil.HookPoint, pubDbxrefHook PubDbxrefHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		pubDbxrefBeforeInsertHooks = append(pubDbxrefBeforeInsertHooks, pubDbxrefHook)
	case boil.BeforeUpdateHook:
		pubDbxrefBeforeUpdateHooks = append(pubDbxrefBeforeUpdateHooks, pubDbxrefHook)
	case boil.BeforeDeleteHook:
		pubDbxrefBeforeDeleteHooks = append(pubDbxrefBeforeDeleteHooks, pubDbxrefHook)
	case boil.BeforeUpsertHook:
		pubDbxrefBeforeUpsertHooks = append(pubDbxrefBeforeUpsertHooks, pubDbxrefHook)
	case boil.AfterInsertHook:
		pubDbxrefAfterInsertHooks = append(pubDbxrefAfterInsertHooks, pubDbxrefHook)
	case boil.AfterSelectHook:
		pubDbxrefAfterSelectHooks = append(pubDbxrefAfterSelectHooks, pubDbxrefHook)
	case boil.AfterUpdateHook:
		pubDbxrefAfterUpdateHooks = append(pubDbxrefAfterUpdateHooks, pubDbxrefHook)
	case boil.AfterDeleteHook:
		pubDbxrefAfterDeleteHooks = append(pubDbxrefAfterDeleteHooks, pubDbxrefHook)
	case boil.AfterUpsertHook:
		pubDbxrefAfterUpsertHooks = append(pubDbxrefAfterUpsertHooks, pubDbxrefHook)
	}
}

// OneP returns a single pubDbxref record from the query, and panics on error.
func (q pubDbxrefQuery) OneP() *PubDbxref {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single pubDbxref record from the query.
func (q pubDbxrefQuery) One() (*PubDbxref, error) {
	o := &PubDbxref{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for pub_dbxref")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all PubDbxref records from the query, and panics on error.
func (q pubDbxrefQuery) AllP() PubDbxrefSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all PubDbxref records from the query.
func (q pubDbxrefQuery) All() (PubDbxrefSlice, error) {
	var o PubDbxrefSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to PubDbxref slice")
	}

	if len(pubDbxrefAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all PubDbxref records in the query, and panics on error.
func (q pubDbxrefQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all PubDbxref records in the query.
func (q pubDbxrefQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count pub_dbxref rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q pubDbxrefQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q pubDbxrefQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if pub_dbxref exists")
	}

	return count > 0, nil
}

// PubG pointed to by the foreign key.
func (o *PubDbxref) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *PubDbxref) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// DbxrefG pointed to by the foreign key.
func (o *PubDbxref) DbxrefG(mods ...qm.QueryMod) dbxrefQuery {
	return o.Dbxref(boil.GetDB(), mods...)
}

// Dbxref pointed to by the foreign key.
func (o *PubDbxref) Dbxref(exec boil.Executor, mods ...qm.QueryMod) dbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := Dbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"dbxref\"")

	return query
}

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubDbxrefL) LoadPub(e boil.Executor, singular bool, maybePubDbxref interface{}) error {
	var slice []*PubDbxref
	var object *PubDbxref

	count := 1
	if singular {
		object = maybePubDbxref.(*PubDbxref)
	} else {
		slice = *maybePubDbxref.(*PubDbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubDbxrefR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubDbxrefR{}
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

	if len(pubDbxrefAfterSelectHooks) != 0 {
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

// LoadDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubDbxrefL) LoadDbxref(e boil.Executor, singular bool, maybePubDbxref interface{}) error {
	var slice []*PubDbxref
	var object *PubDbxref

	count := 1
	if singular {
		object = maybePubDbxref.(*PubDbxref)
	} else {
		slice = *maybePubDbxref.(*PubDbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubDbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &pubDbxrefR{}
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

	if len(pubDbxrefAfterSelectHooks) != 0 {
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

// SetPub of the pub_dbxref to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.PubDbxref.
func (o *PubDbxref) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"pub_dbxref\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, pubDbxrefPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.PubDbxrefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &pubDbxrefR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			PubDbxref: o,
		}
	} else {
		related.R.PubDbxref = o
	}

	return nil
}

// SetDbxref of the pub_dbxref to the related item.
// Sets o.R.Dbxref to related.
// Adds o to related.R.PubDbxref.
func (o *PubDbxref) SetDbxref(exec boil.Executor, insert bool, related *Dbxref) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"pub_dbxref\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
		strmangle.WhereClause("\"", "\"", 2, pubDbxrefPrimaryKeyColumns),
	)
	values := []interface{}{related.DbxrefID, o.PubDbxrefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DbxrefID = related.DbxrefID

	if o.R == nil {
		o.R = &pubDbxrefR{
			Dbxref: related,
		}
	} else {
		o.R.Dbxref = related
	}

	if related.R == nil {
		related.R = &dbxrefR{
			PubDbxref: o,
		}
	} else {
		related.R.PubDbxref = o
	}

	return nil
}

// PubDbxrefsG retrieves all records.
func PubDbxrefsG(mods ...qm.QueryMod) pubDbxrefQuery {
	return PubDbxrefs(boil.GetDB(), mods...)
}

// PubDbxrefs retrieves all the records using an executor.
func PubDbxrefs(exec boil.Executor, mods ...qm.QueryMod) pubDbxrefQuery {
	mods = append(mods, qm.From("\"pub_dbxref\""))
	return pubDbxrefQuery{NewQuery(exec, mods...)}
}

// FindPubDbxrefG retrieves a single record by ID.
func FindPubDbxrefG(pubDbxrefID int, selectCols ...string) (*PubDbxref, error) {
	return FindPubDbxref(boil.GetDB(), pubDbxrefID, selectCols...)
}

// FindPubDbxrefGP retrieves a single record by ID, and panics on error.
func FindPubDbxrefGP(pubDbxrefID int, selectCols ...string) *PubDbxref {
	retobj, err := FindPubDbxref(boil.GetDB(), pubDbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindPubDbxref retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPubDbxref(exec boil.Executor, pubDbxrefID int, selectCols ...string) (*PubDbxref, error) {
	pubDbxrefObj := &PubDbxref{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pub_dbxref\" where \"pub_dbxref_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, pubDbxrefID)

	err := q.Bind(pubDbxrefObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from pub_dbxref")
	}

	return pubDbxrefObj, nil
}

// FindPubDbxrefP retrieves a single record by ID with an executor, and panics on error.
func FindPubDbxrefP(exec boil.Executor, pubDbxrefID int, selectCols ...string) *PubDbxref {
	retobj, err := FindPubDbxref(exec, pubDbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *PubDbxref) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *PubDbxref) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *PubDbxref) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *PubDbxref) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no pub_dbxref provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pubDbxrefColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	pubDbxrefInsertCacheMut.RLock()
	cache, cached := pubDbxrefInsertCache[key]
	pubDbxrefInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			pubDbxrefColumns,
			pubDbxrefColumnsWithDefault,
			pubDbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(pubDbxrefType, pubDbxrefMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pubDbxrefType, pubDbxrefMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"pub_dbxref\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into pub_dbxref")
	}

	if !cached {
		pubDbxrefInsertCacheMut.Lock()
		pubDbxrefInsertCache[key] = cache
		pubDbxrefInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single PubDbxref record. See Update for
// whitelist behavior description.
func (o *PubDbxref) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single PubDbxref record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *PubDbxref) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the PubDbxref, and panics on error.
// See Update for whitelist behavior description.
func (o *PubDbxref) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the PubDbxref.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *PubDbxref) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	pubDbxrefUpdateCacheMut.RLock()
	cache, cached := pubDbxrefUpdateCache[key]
	pubDbxrefUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(pubDbxrefColumns, pubDbxrefPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update pub_dbxref, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pub_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pubDbxrefPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pubDbxrefType, pubDbxrefMapping, append(wl, pubDbxrefPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update pub_dbxref row")
	}

	if !cached {
		pubDbxrefUpdateCacheMut.Lock()
		pubDbxrefUpdateCache[key] = cache
		pubDbxrefUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q pubDbxrefQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q pubDbxrefQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for pub_dbxref")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PubDbxrefSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o PubDbxrefSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o PubDbxrefSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PubDbxrefSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pubDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"pub_dbxref\" SET %s WHERE (\"pub_dbxref_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(pubDbxrefPrimaryKeyColumns), len(colNames)+1, len(pubDbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in pubDbxref slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *PubDbxref) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *PubDbxref) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *PubDbxref) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *PubDbxref) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no pub_dbxref provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pubDbxrefColumnsWithDefault, o)

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

	pubDbxrefUpsertCacheMut.RLock()
	cache, cached := pubDbxrefUpsertCache[key]
	pubDbxrefUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			pubDbxrefColumns,
			pubDbxrefColumnsWithDefault,
			pubDbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			pubDbxrefColumns,
			pubDbxrefPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert pub_dbxref, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pubDbxrefPrimaryKeyColumns))
			copy(conflict, pubDbxrefPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"pub_dbxref\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(pubDbxrefType, pubDbxrefMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pubDbxrefType, pubDbxrefMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for pub_dbxref")
	}

	if !cached {
		pubDbxrefUpsertCacheMut.Lock()
		pubDbxrefUpsertCache[key] = cache
		pubDbxrefUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single PubDbxref record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *PubDbxref) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single PubDbxref record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *PubDbxref) DeleteG() error {
	if o == nil {
		return errors.New("chado: no PubDbxref provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single PubDbxref record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *PubDbxref) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single PubDbxref record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PubDbxref) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no PubDbxref provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pubDbxrefPrimaryKeyMapping)
	sql := "DELETE FROM \"pub_dbxref\" WHERE \"pub_dbxref_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from pub_dbxref")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q pubDbxrefQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q pubDbxrefQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no pubDbxrefQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from pub_dbxref")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o PubDbxrefSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o PubDbxrefSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no PubDbxref slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o PubDbxrefSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PubDbxrefSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no PubDbxref slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(pubDbxrefBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pubDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"pub_dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, pubDbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(pubDbxrefPrimaryKeyColumns), 1, len(pubDbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from pubDbxref slice")
	}

	if len(pubDbxrefAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *PubDbxref) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *PubDbxref) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *PubDbxref) ReloadG() error {
	if o == nil {
		return errors.New("chado: no PubDbxref provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PubDbxref) Reload(exec boil.Executor) error {
	ret, err := FindPubDbxref(exec, o.PubDbxrefID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PubDbxrefSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PubDbxrefSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PubDbxrefSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty PubDbxrefSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PubDbxrefSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	pubDbxrefs := PubDbxrefSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pubDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"pub_dbxref\".* FROM \"pub_dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, pubDbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(pubDbxrefPrimaryKeyColumns), 1, len(pubDbxrefPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&pubDbxrefs)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in PubDbxrefSlice")
	}

	*o = pubDbxrefs

	return nil
}

// PubDbxrefExists checks if the PubDbxref row exists.
func PubDbxrefExists(exec boil.Executor, pubDbxrefID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"pub_dbxref\" where \"pub_dbxref_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, pubDbxrefID)
	}

	row := exec.QueryRow(sql, pubDbxrefID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if pub_dbxref exists")
	}

	return exists, nil
}

// PubDbxrefExistsG checks if the PubDbxref row exists.
func PubDbxrefExistsG(pubDbxrefID int) (bool, error) {
	return PubDbxrefExists(boil.GetDB(), pubDbxrefID)
}

// PubDbxrefExistsGP checks if the PubDbxref row exists. Panics on error.
func PubDbxrefExistsGP(pubDbxrefID int) bool {
	e, err := PubDbxrefExists(boil.GetDB(), pubDbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// PubDbxrefExistsP checks if the PubDbxref row exists. Panics on error.
func PubDbxrefExistsP(exec boil.Executor, pubDbxrefID int) bool {
	e, err := PubDbxrefExists(exec, pubDbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
