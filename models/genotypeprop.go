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

// Genotypeprop is an object representing the database table.
type Genotypeprop struct {
	GenotypepropID int         `boil:"genotypeprop_id" json:"genotypeprop_id" toml:"genotypeprop_id" yaml:"genotypeprop_id"`
	GenotypeID     int         `boil:"genotype_id" json:"genotype_id" toml:"genotype_id" yaml:"genotype_id"`
	TypeID         int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value          null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank           int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *genotypepropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L genotypepropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// genotypepropR is where relationships are stored.
type genotypepropR struct {
	Genotype *Genotype
	Type     *Cvterm
}

// genotypepropL is where Load methods for each relationship are stored.
type genotypepropL struct{}

var (
	genotypepropColumns               = []string{"genotypeprop_id", "genotype_id", "type_id", "value", "rank"}
	genotypepropColumnsWithoutDefault = []string{"genotype_id", "type_id", "value"}
	genotypepropColumnsWithDefault    = []string{"genotypeprop_id", "rank"}
	genotypepropPrimaryKeyColumns     = []string{"genotypeprop_id"}
)

type (
	// GenotypepropSlice is an alias for a slice of pointers to Genotypeprop.
	// This should generally be used opposed to []Genotypeprop.
	GenotypepropSlice []*Genotypeprop
	// GenotypepropHook is the signature for custom Genotypeprop hook methods
	GenotypepropHook func(boil.Executor, *Genotypeprop) error

	genotypepropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	genotypepropType                 = reflect.TypeOf(&Genotypeprop{})
	genotypepropMapping              = queries.MakeStructMapping(genotypepropType)
	genotypepropPrimaryKeyMapping, _ = queries.BindMapping(genotypepropType, genotypepropMapping, genotypepropPrimaryKeyColumns)
	genotypepropInsertCacheMut       sync.RWMutex
	genotypepropInsertCache          = make(map[string]insertCache)
	genotypepropUpdateCacheMut       sync.RWMutex
	genotypepropUpdateCache          = make(map[string]updateCache)
	genotypepropUpsertCacheMut       sync.RWMutex
	genotypepropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var genotypepropBeforeInsertHooks []GenotypepropHook
var genotypepropBeforeUpdateHooks []GenotypepropHook
var genotypepropBeforeDeleteHooks []GenotypepropHook
var genotypepropBeforeUpsertHooks []GenotypepropHook

var genotypepropAfterInsertHooks []GenotypepropHook
var genotypepropAfterSelectHooks []GenotypepropHook
var genotypepropAfterUpdateHooks []GenotypepropHook
var genotypepropAfterDeleteHooks []GenotypepropHook
var genotypepropAfterUpsertHooks []GenotypepropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Genotypeprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypepropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Genotypeprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypepropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Genotypeprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypepropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Genotypeprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypepropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Genotypeprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypepropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Genotypeprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypepropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Genotypeprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypepropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Genotypeprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypepropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Genotypeprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range genotypepropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGenotypepropHook registers your hook function for all future operations.
func AddGenotypepropHook(hookPoint boil.HookPoint, genotypepropHook GenotypepropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		genotypepropBeforeInsertHooks = append(genotypepropBeforeInsertHooks, genotypepropHook)
	case boil.BeforeUpdateHook:
		genotypepropBeforeUpdateHooks = append(genotypepropBeforeUpdateHooks, genotypepropHook)
	case boil.BeforeDeleteHook:
		genotypepropBeforeDeleteHooks = append(genotypepropBeforeDeleteHooks, genotypepropHook)
	case boil.BeforeUpsertHook:
		genotypepropBeforeUpsertHooks = append(genotypepropBeforeUpsertHooks, genotypepropHook)
	case boil.AfterInsertHook:
		genotypepropAfterInsertHooks = append(genotypepropAfterInsertHooks, genotypepropHook)
	case boil.AfterSelectHook:
		genotypepropAfterSelectHooks = append(genotypepropAfterSelectHooks, genotypepropHook)
	case boil.AfterUpdateHook:
		genotypepropAfterUpdateHooks = append(genotypepropAfterUpdateHooks, genotypepropHook)
	case boil.AfterDeleteHook:
		genotypepropAfterDeleteHooks = append(genotypepropAfterDeleteHooks, genotypepropHook)
	case boil.AfterUpsertHook:
		genotypepropAfterUpsertHooks = append(genotypepropAfterUpsertHooks, genotypepropHook)
	}
}

// OneP returns a single genotypeprop record from the query, and panics on error.
func (q genotypepropQuery) OneP() *Genotypeprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single genotypeprop record from the query.
func (q genotypepropQuery) One() (*Genotypeprop, error) {
	o := &Genotypeprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for genotypeprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Genotypeprop records from the query, and panics on error.
func (q genotypepropQuery) AllP() GenotypepropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Genotypeprop records from the query.
func (q genotypepropQuery) All() (GenotypepropSlice, error) {
	var o GenotypepropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Genotypeprop slice")
	}

	if len(genotypepropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Genotypeprop records in the query, and panics on error.
func (q genotypepropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Genotypeprop records in the query.
func (q genotypepropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count genotypeprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q genotypepropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q genotypepropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if genotypeprop exists")
	}

	return count > 0, nil
}

// GenotypeG pointed to by the foreign key.
func (o *Genotypeprop) GenotypeG(mods ...qm.QueryMod) genotypeQuery {
	return o.Genotype(boil.GetDB(), mods...)
}

// Genotype pointed to by the foreign key.
func (o *Genotypeprop) Genotype(exec boil.Executor, mods ...qm.QueryMod) genotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype_id=$1", o.GenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Genotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"genotype\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Genotypeprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Genotypeprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadGenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (genotypepropL) LoadGenotype(e boil.Executor, singular bool, maybeGenotypeprop interface{}) error {
	var slice []*Genotypeprop
	var object *Genotypeprop

	count := 1
	if singular {
		object = maybeGenotypeprop.(*Genotypeprop)
	} else {
		slice = *maybeGenotypeprop.(*GenotypepropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &genotypepropR{}
		args[0] = object.GenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &genotypepropR{}
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

	if len(genotypepropAfterSelectHooks) != 0 {
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

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (genotypepropL) LoadType(e boil.Executor, singular bool, maybeGenotypeprop interface{}) error {
	var slice []*Genotypeprop
	var object *Genotypeprop

	count := 1
	if singular {
		object = maybeGenotypeprop.(*Genotypeprop)
	} else {
		slice = *maybeGenotypeprop.(*GenotypepropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &genotypepropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &genotypepropR{}
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

	if len(genotypepropAfterSelectHooks) != 0 {
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

// SetGenotype of the genotypeprop to the related item.
// Sets o.R.Genotype to related.
// Adds o to related.R.Genotypeprop.
func (o *Genotypeprop) SetGenotype(exec boil.Executor, insert bool, related *Genotype) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"genotypeprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"genotype_id"}),
		strmangle.WhereClause("\"", "\"", 2, genotypepropPrimaryKeyColumns),
	)
	values := []interface{}{related.GenotypeID, o.GenotypepropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GenotypeID = related.GenotypeID

	if o.R == nil {
		o.R = &genotypepropR{
			Genotype: related,
		}
	} else {
		o.R.Genotype = related
	}

	if related.R == nil {
		related.R = &genotypeR{
			Genotypeprop: o,
		}
	} else {
		related.R.Genotypeprop = o
	}

	return nil
}

// SetType of the genotypeprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeGenotypeprop.
func (o *Genotypeprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"genotypeprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, genotypepropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.GenotypepropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &genotypepropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeGenotypeprop: o,
		}
	} else {
		related.R.TypeGenotypeprop = o
	}

	return nil
}

// GenotypepropsG retrieves all records.
func GenotypepropsG(mods ...qm.QueryMod) genotypepropQuery {
	return Genotypeprops(boil.GetDB(), mods...)
}

// Genotypeprops retrieves all the records using an executor.
func Genotypeprops(exec boil.Executor, mods ...qm.QueryMod) genotypepropQuery {
	mods = append(mods, qm.From("\"genotypeprop\""))
	return genotypepropQuery{NewQuery(exec, mods...)}
}

// FindGenotypepropG retrieves a single record by ID.
func FindGenotypepropG(genotypepropID int, selectCols ...string) (*Genotypeprop, error) {
	return FindGenotypeprop(boil.GetDB(), genotypepropID, selectCols...)
}

// FindGenotypepropGP retrieves a single record by ID, and panics on error.
func FindGenotypepropGP(genotypepropID int, selectCols ...string) *Genotypeprop {
	retobj, err := FindGenotypeprop(boil.GetDB(), genotypepropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindGenotypeprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGenotypeprop(exec boil.Executor, genotypepropID int, selectCols ...string) (*Genotypeprop, error) {
	genotypepropObj := &Genotypeprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"genotypeprop\" where \"genotypeprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, genotypepropID)

	err := q.Bind(genotypepropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from genotypeprop")
	}

	return genotypepropObj, nil
}

// FindGenotypepropP retrieves a single record by ID with an executor, and panics on error.
func FindGenotypepropP(exec boil.Executor, genotypepropID int, selectCols ...string) *Genotypeprop {
	retobj, err := FindGenotypeprop(exec, genotypepropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Genotypeprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Genotypeprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Genotypeprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Genotypeprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no genotypeprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(genotypepropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	genotypepropInsertCacheMut.RLock()
	cache, cached := genotypepropInsertCache[key]
	genotypepropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			genotypepropColumns,
			genotypepropColumnsWithDefault,
			genotypepropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(genotypepropType, genotypepropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(genotypepropType, genotypepropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"genotypeprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into genotypeprop")
	}

	if !cached {
		genotypepropInsertCacheMut.Lock()
		genotypepropInsertCache[key] = cache
		genotypepropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Genotypeprop record. See Update for
// whitelist behavior description.
func (o *Genotypeprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Genotypeprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Genotypeprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Genotypeprop, and panics on error.
// See Update for whitelist behavior description.
func (o *Genotypeprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Genotypeprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Genotypeprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	genotypepropUpdateCacheMut.RLock()
	cache, cached := genotypepropUpdateCache[key]
	genotypepropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(genotypepropColumns, genotypepropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update genotypeprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"genotypeprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, genotypepropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(genotypepropType, genotypepropMapping, append(wl, genotypepropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update genotypeprop row")
	}

	if !cached {
		genotypepropUpdateCacheMut.Lock()
		genotypepropUpdateCache[key] = cache
		genotypepropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q genotypepropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q genotypepropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for genotypeprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o GenotypepropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o GenotypepropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o GenotypepropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GenotypepropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), genotypepropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"genotypeprop\" SET %s WHERE (\"genotypeprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(genotypepropPrimaryKeyColumns), len(colNames)+1, len(genotypepropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in genotypeprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Genotypeprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Genotypeprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Genotypeprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Genotypeprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no genotypeprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(genotypepropColumnsWithDefault, o)

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

	genotypepropUpsertCacheMut.RLock()
	cache, cached := genotypepropUpsertCache[key]
	genotypepropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			genotypepropColumns,
			genotypepropColumnsWithDefault,
			genotypepropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			genotypepropColumns,
			genotypepropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert genotypeprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(genotypepropPrimaryKeyColumns))
			copy(conflict, genotypepropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"genotypeprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(genotypepropType, genotypepropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(genotypepropType, genotypepropMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for genotypeprop")
	}

	if !cached {
		genotypepropUpsertCacheMut.Lock()
		genotypepropUpsertCache[key] = cache
		genotypepropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Genotypeprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Genotypeprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Genotypeprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Genotypeprop) DeleteG() error {
	if o == nil {
		return errors.New("models: no Genotypeprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Genotypeprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Genotypeprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Genotypeprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Genotypeprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Genotypeprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), genotypepropPrimaryKeyMapping)
	sql := "DELETE FROM \"genotypeprop\" WHERE \"genotypeprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from genotypeprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q genotypepropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q genotypepropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no genotypepropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from genotypeprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o GenotypepropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o GenotypepropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Genotypeprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o GenotypepropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GenotypepropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Genotypeprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(genotypepropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), genotypepropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"genotypeprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, genotypepropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(genotypepropPrimaryKeyColumns), 1, len(genotypepropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from genotypeprop slice")
	}

	if len(genotypepropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Genotypeprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Genotypeprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Genotypeprop) ReloadG() error {
	if o == nil {
		return errors.New("models: no Genotypeprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Genotypeprop) Reload(exec boil.Executor) error {
	ret, err := FindGenotypeprop(exec, o.GenotypepropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *GenotypepropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *GenotypepropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GenotypepropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty GenotypepropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GenotypepropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	genotypeprops := GenotypepropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), genotypepropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"genotypeprop\".* FROM \"genotypeprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, genotypepropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(genotypepropPrimaryKeyColumns), 1, len(genotypepropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&genotypeprops)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GenotypepropSlice")
	}

	*o = genotypeprops

	return nil
}

// GenotypepropExists checks if the Genotypeprop row exists.
func GenotypepropExists(exec boil.Executor, genotypepropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"genotypeprop\" where \"genotypeprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, genotypepropID)
	}

	row := exec.QueryRow(sql, genotypepropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if genotypeprop exists")
	}

	return exists, nil
}

// GenotypepropExistsG checks if the Genotypeprop row exists.
func GenotypepropExistsG(genotypepropID int) (bool, error) {
	return GenotypepropExists(boil.GetDB(), genotypepropID)
}

// GenotypepropExistsGP checks if the Genotypeprop row exists. Panics on error.
func GenotypepropExistsGP(genotypepropID int) bool {
	e, err := GenotypepropExists(boil.GetDB(), genotypepropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// GenotypepropExistsP checks if the Genotypeprop row exists. Panics on error.
func GenotypepropExistsP(exec boil.Executor, genotypepropID int) bool {
	e, err := GenotypepropExists(exec, genotypepropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
