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

// OrganismDbxref is an object representing the database table.
type OrganismDbxref struct {
	OrganismDbxrefID int `boil:"organism_dbxref_id" json:"organism_dbxref_id" toml:"organism_dbxref_id" yaml:"organism_dbxref_id"`
	OrganismID       int `boil:"organism_id" json:"organism_id" toml:"organism_id" yaml:"organism_id"`
	DbxrefID         int `boil:"dbxref_id" json:"dbxref_id" toml:"dbxref_id" yaml:"dbxref_id"`

	R *organismDbxrefR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L organismDbxrefL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// organismDbxrefR is where relationships are stored.
type organismDbxrefR struct {
	Organism *Organism
	Dbxref   *Dbxref
}

// organismDbxrefL is where Load methods for each relationship are stored.
type organismDbxrefL struct{}

var (
	organismDbxrefColumns               = []string{"organism_dbxref_id", "organism_id", "dbxref_id"}
	organismDbxrefColumnsWithoutDefault = []string{"organism_id", "dbxref_id"}
	organismDbxrefColumnsWithDefault    = []string{"organism_dbxref_id"}
	organismDbxrefPrimaryKeyColumns     = []string{"organism_dbxref_id"}
)

type (
	// OrganismDbxrefSlice is an alias for a slice of pointers to OrganismDbxref.
	// This should generally be used opposed to []OrganismDbxref.
	OrganismDbxrefSlice []*OrganismDbxref
	// OrganismDbxrefHook is the signature for custom OrganismDbxref hook methods
	OrganismDbxrefHook func(boil.Executor, *OrganismDbxref) error

	organismDbxrefQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	organismDbxrefType                 = reflect.TypeOf(&OrganismDbxref{})
	organismDbxrefMapping              = queries.MakeStructMapping(organismDbxrefType)
	organismDbxrefPrimaryKeyMapping, _ = queries.BindMapping(organismDbxrefType, organismDbxrefMapping, organismDbxrefPrimaryKeyColumns)
	organismDbxrefInsertCacheMut       sync.RWMutex
	organismDbxrefInsertCache          = make(map[string]insertCache)
	organismDbxrefUpdateCacheMut       sync.RWMutex
	organismDbxrefUpdateCache          = make(map[string]updateCache)
	organismDbxrefUpsertCacheMut       sync.RWMutex
	organismDbxrefUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var organismDbxrefBeforeInsertHooks []OrganismDbxrefHook
var organismDbxrefBeforeUpdateHooks []OrganismDbxrefHook
var organismDbxrefBeforeDeleteHooks []OrganismDbxrefHook
var organismDbxrefBeforeUpsertHooks []OrganismDbxrefHook

var organismDbxrefAfterInsertHooks []OrganismDbxrefHook
var organismDbxrefAfterSelectHooks []OrganismDbxrefHook
var organismDbxrefAfterUpdateHooks []OrganismDbxrefHook
var organismDbxrefAfterDeleteHooks []OrganismDbxrefHook
var organismDbxrefAfterUpsertHooks []OrganismDbxrefHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OrganismDbxref) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organismDbxrefBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OrganismDbxref) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range organismDbxrefBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OrganismDbxref) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range organismDbxrefBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OrganismDbxref) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organismDbxrefBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OrganismDbxref) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organismDbxrefAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OrganismDbxref) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range organismDbxrefAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OrganismDbxref) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range organismDbxrefAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OrganismDbxref) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range organismDbxrefAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OrganismDbxref) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organismDbxrefAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrganismDbxrefHook registers your hook function for all future operations.
func AddOrganismDbxrefHook(hookPoint boil.HookPoint, organismDbxrefHook OrganismDbxrefHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		organismDbxrefBeforeInsertHooks = append(organismDbxrefBeforeInsertHooks, organismDbxrefHook)
	case boil.BeforeUpdateHook:
		organismDbxrefBeforeUpdateHooks = append(organismDbxrefBeforeUpdateHooks, organismDbxrefHook)
	case boil.BeforeDeleteHook:
		organismDbxrefBeforeDeleteHooks = append(organismDbxrefBeforeDeleteHooks, organismDbxrefHook)
	case boil.BeforeUpsertHook:
		organismDbxrefBeforeUpsertHooks = append(organismDbxrefBeforeUpsertHooks, organismDbxrefHook)
	case boil.AfterInsertHook:
		organismDbxrefAfterInsertHooks = append(organismDbxrefAfterInsertHooks, organismDbxrefHook)
	case boil.AfterSelectHook:
		organismDbxrefAfterSelectHooks = append(organismDbxrefAfterSelectHooks, organismDbxrefHook)
	case boil.AfterUpdateHook:
		organismDbxrefAfterUpdateHooks = append(organismDbxrefAfterUpdateHooks, organismDbxrefHook)
	case boil.AfterDeleteHook:
		organismDbxrefAfterDeleteHooks = append(organismDbxrefAfterDeleteHooks, organismDbxrefHook)
	case boil.AfterUpsertHook:
		organismDbxrefAfterUpsertHooks = append(organismDbxrefAfterUpsertHooks, organismDbxrefHook)
	}
}

// OneP returns a single organismDbxref record from the query, and panics on error.
func (q organismDbxrefQuery) OneP() *OrganismDbxref {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single organismDbxref record from the query.
func (q organismDbxrefQuery) One() (*OrganismDbxref, error) {
	o := &OrganismDbxref{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for organism_dbxref")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all OrganismDbxref records from the query, and panics on error.
func (q organismDbxrefQuery) AllP() OrganismDbxrefSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all OrganismDbxref records from the query.
func (q organismDbxrefQuery) All() (OrganismDbxrefSlice, error) {
	var o OrganismDbxrefSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OrganismDbxref slice")
	}

	if len(organismDbxrefAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all OrganismDbxref records in the query, and panics on error.
func (q organismDbxrefQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all OrganismDbxref records in the query.
func (q organismDbxrefQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count organism_dbxref rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q organismDbxrefQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q organismDbxrefQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if organism_dbxref exists")
	}

	return count > 0, nil
}

// OrganismG pointed to by the foreign key.
func (o *OrganismDbxref) OrganismG(mods ...qm.QueryMod) organismQuery {
	return o.Organism(boil.GetDB(), mods...)
}

// Organism pointed to by the foreign key.
func (o *OrganismDbxref) Organism(exec boil.Executor, mods ...qm.QueryMod) organismQuery {
	queryMods := []qm.QueryMod{
		qm.Where("organism_id=$1", o.OrganismID),
	}

	queryMods = append(queryMods, mods...)

	query := Organisms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"organism\"")

	return query
}

// DbxrefG pointed to by the foreign key.
func (o *OrganismDbxref) DbxrefG(mods ...qm.QueryMod) dbxrefQuery {
	return o.Dbxref(boil.GetDB(), mods...)
}

// Dbxref pointed to by the foreign key.
func (o *OrganismDbxref) Dbxref(exec boil.Executor, mods ...qm.QueryMod) dbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := Dbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"dbxref\"")

	return query
}

// LoadOrganism allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (organismDbxrefL) LoadOrganism(e boil.Executor, singular bool, maybeOrganismDbxref interface{}) error {
	var slice []*OrganismDbxref
	var object *OrganismDbxref

	count := 1
	if singular {
		object = maybeOrganismDbxref.(*OrganismDbxref)
	} else {
		slice = *maybeOrganismDbxref.(*OrganismDbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &organismDbxrefR{}
		args[0] = object.OrganismID
	} else {
		for i, obj := range slice {
			obj.R = &organismDbxrefR{}
			args[i] = obj.OrganismID
		}
	}

	query := fmt.Sprintf(
		"select * from \"organism\" where \"organism_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Organism")
	}
	defer results.Close()

	var resultSlice []*Organism
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Organism")
	}

	if len(organismDbxrefAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Organism = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.OrganismID == foreign.OrganismID {
				local.R.Organism = foreign
				break
			}
		}
	}

	return nil
}

// LoadDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (organismDbxrefL) LoadDbxref(e boil.Executor, singular bool, maybeOrganismDbxref interface{}) error {
	var slice []*OrganismDbxref
	var object *OrganismDbxref

	count := 1
	if singular {
		object = maybeOrganismDbxref.(*OrganismDbxref)
	} else {
		slice = *maybeOrganismDbxref.(*OrganismDbxrefSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &organismDbxrefR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &organismDbxrefR{}
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

	if len(organismDbxrefAfterSelectHooks) != 0 {
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

// SetOrganism of the organism_dbxref to the related item.
// Sets o.R.Organism to related.
// Adds o to related.R.OrganismDbxref.
func (o *OrganismDbxref) SetOrganism(exec boil.Executor, insert bool, related *Organism) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"organism_dbxref\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"organism_id"}),
		strmangle.WhereClause("\"", "\"", 2, organismDbxrefPrimaryKeyColumns),
	)
	values := []interface{}{related.OrganismID, o.OrganismDbxrefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrganismID = related.OrganismID

	if o.R == nil {
		o.R = &organismDbxrefR{
			Organism: related,
		}
	} else {
		o.R.Organism = related
	}

	if related.R == nil {
		related.R = &organismR{
			OrganismDbxref: o,
		}
	} else {
		related.R.OrganismDbxref = o
	}

	return nil
}

// SetDbxref of the organism_dbxref to the related item.
// Sets o.R.Dbxref to related.
// Adds o to related.R.OrganismDbxref.
func (o *OrganismDbxref) SetDbxref(exec boil.Executor, insert bool, related *Dbxref) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"organism_dbxref\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
		strmangle.WhereClause("\"", "\"", 2, organismDbxrefPrimaryKeyColumns),
	)
	values := []interface{}{related.DbxrefID, o.OrganismDbxrefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DbxrefID = related.DbxrefID

	if o.R == nil {
		o.R = &organismDbxrefR{
			Dbxref: related,
		}
	} else {
		o.R.Dbxref = related
	}

	if related.R == nil {
		related.R = &dbxrefR{
			OrganismDbxref: o,
		}
	} else {
		related.R.OrganismDbxref = o
	}

	return nil
}

// OrganismDbxrefsG retrieves all records.
func OrganismDbxrefsG(mods ...qm.QueryMod) organismDbxrefQuery {
	return OrganismDbxrefs(boil.GetDB(), mods...)
}

// OrganismDbxrefs retrieves all the records using an executor.
func OrganismDbxrefs(exec boil.Executor, mods ...qm.QueryMod) organismDbxrefQuery {
	mods = append(mods, qm.From("\"organism_dbxref\""))
	return organismDbxrefQuery{NewQuery(exec, mods...)}
}

// FindOrganismDbxrefG retrieves a single record by ID.
func FindOrganismDbxrefG(organismDbxrefID int, selectCols ...string) (*OrganismDbxref, error) {
	return FindOrganismDbxref(boil.GetDB(), organismDbxrefID, selectCols...)
}

// FindOrganismDbxrefGP retrieves a single record by ID, and panics on error.
func FindOrganismDbxrefGP(organismDbxrefID int, selectCols ...string) *OrganismDbxref {
	retobj, err := FindOrganismDbxref(boil.GetDB(), organismDbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindOrganismDbxref retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrganismDbxref(exec boil.Executor, organismDbxrefID int, selectCols ...string) (*OrganismDbxref, error) {
	organismDbxrefObj := &OrganismDbxref{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"organism_dbxref\" where \"organism_dbxref_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, organismDbxrefID)

	err := q.Bind(organismDbxrefObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from organism_dbxref")
	}

	return organismDbxrefObj, nil
}

// FindOrganismDbxrefP retrieves a single record by ID with an executor, and panics on error.
func FindOrganismDbxrefP(exec boil.Executor, organismDbxrefID int, selectCols ...string) *OrganismDbxref {
	retobj, err := FindOrganismDbxref(exec, organismDbxrefID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *OrganismDbxref) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *OrganismDbxref) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *OrganismDbxref) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *OrganismDbxref) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no organism_dbxref provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(organismDbxrefColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	organismDbxrefInsertCacheMut.RLock()
	cache, cached := organismDbxrefInsertCache[key]
	organismDbxrefInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			organismDbxrefColumns,
			organismDbxrefColumnsWithDefault,
			organismDbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(organismDbxrefType, organismDbxrefMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(organismDbxrefType, organismDbxrefMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"organism_dbxref\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into organism_dbxref")
	}

	if !cached {
		organismDbxrefInsertCacheMut.Lock()
		organismDbxrefInsertCache[key] = cache
		organismDbxrefInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single OrganismDbxref record. See Update for
// whitelist behavior description.
func (o *OrganismDbxref) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single OrganismDbxref record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *OrganismDbxref) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the OrganismDbxref, and panics on error.
// See Update for whitelist behavior description.
func (o *OrganismDbxref) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the OrganismDbxref.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *OrganismDbxref) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	organismDbxrefUpdateCacheMut.RLock()
	cache, cached := organismDbxrefUpdateCache[key]
	organismDbxrefUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(organismDbxrefColumns, organismDbxrefPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update organism_dbxref, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"organism_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, organismDbxrefPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(organismDbxrefType, organismDbxrefMapping, append(wl, organismDbxrefPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update organism_dbxref row")
	}

	if !cached {
		organismDbxrefUpdateCacheMut.Lock()
		organismDbxrefUpdateCache[key] = cache
		organismDbxrefUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q organismDbxrefQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q organismDbxrefQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for organism_dbxref")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OrganismDbxrefSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o OrganismDbxrefSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o OrganismDbxrefSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrganismDbxrefSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organismDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"organism_dbxref\" SET %s WHERE (\"organism_dbxref_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(organismDbxrefPrimaryKeyColumns), len(colNames)+1, len(organismDbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in organismDbxref slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *OrganismDbxref) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *OrganismDbxref) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *OrganismDbxref) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *OrganismDbxref) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no organism_dbxref provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(organismDbxrefColumnsWithDefault, o)

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

	organismDbxrefUpsertCacheMut.RLock()
	cache, cached := organismDbxrefUpsertCache[key]
	organismDbxrefUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			organismDbxrefColumns,
			organismDbxrefColumnsWithDefault,
			organismDbxrefColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			organismDbxrefColumns,
			organismDbxrefPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert organism_dbxref, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(organismDbxrefPrimaryKeyColumns))
			copy(conflict, organismDbxrefPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"organism_dbxref\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(organismDbxrefType, organismDbxrefMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(organismDbxrefType, organismDbxrefMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for organism_dbxref")
	}

	if !cached {
		organismDbxrefUpsertCacheMut.Lock()
		organismDbxrefUpsertCache[key] = cache
		organismDbxrefUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single OrganismDbxref record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OrganismDbxref) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single OrganismDbxref record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *OrganismDbxref) DeleteG() error {
	if o == nil {
		return errors.New("models: no OrganismDbxref provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single OrganismDbxref record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OrganismDbxref) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single OrganismDbxref record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OrganismDbxref) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no OrganismDbxref provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), organismDbxrefPrimaryKeyMapping)
	sql := "DELETE FROM \"organism_dbxref\" WHERE \"organism_dbxref_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from organism_dbxref")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q organismDbxrefQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q organismDbxrefQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no organismDbxrefQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from organism_dbxref")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o OrganismDbxrefSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o OrganismDbxrefSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no OrganismDbxref slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o OrganismDbxrefSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrganismDbxrefSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no OrganismDbxref slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(organismDbxrefBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organismDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"organism_dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, organismDbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(organismDbxrefPrimaryKeyColumns), 1, len(organismDbxrefPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from organismDbxref slice")
	}

	if len(organismDbxrefAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *OrganismDbxref) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *OrganismDbxref) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *OrganismDbxref) ReloadG() error {
	if o == nil {
		return errors.New("models: no OrganismDbxref provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OrganismDbxref) Reload(exec boil.Executor) error {
	ret, err := FindOrganismDbxref(exec, o.OrganismDbxrefID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OrganismDbxrefSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OrganismDbxrefSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrganismDbxrefSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty OrganismDbxrefSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrganismDbxrefSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	organismDbxrefs := OrganismDbxrefSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organismDbxrefPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"organism_dbxref\".* FROM \"organism_dbxref\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, organismDbxrefPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(organismDbxrefPrimaryKeyColumns), 1, len(organismDbxrefPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&organismDbxrefs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrganismDbxrefSlice")
	}

	*o = organismDbxrefs

	return nil
}

// OrganismDbxrefExists checks if the OrganismDbxref row exists.
func OrganismDbxrefExists(exec boil.Executor, organismDbxrefID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"organism_dbxref\" where \"organism_dbxref_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, organismDbxrefID)
	}

	row := exec.QueryRow(sql, organismDbxrefID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if organism_dbxref exists")
	}

	return exists, nil
}

// OrganismDbxrefExistsG checks if the OrganismDbxref row exists.
func OrganismDbxrefExistsG(organismDbxrefID int) (bool, error) {
	return OrganismDbxrefExists(boil.GetDB(), organismDbxrefID)
}

// OrganismDbxrefExistsGP checks if the OrganismDbxref row exists. Panics on error.
func OrganismDbxrefExistsGP(organismDbxrefID int) bool {
	e, err := OrganismDbxrefExists(boil.GetDB(), organismDbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// OrganismDbxrefExistsP checks if the OrganismDbxref row exists. Panics on error.
func OrganismDbxrefExistsP(exec boil.Executor, organismDbxrefID int) bool {
	e, err := OrganismDbxrefExists(exec, organismDbxrefID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
