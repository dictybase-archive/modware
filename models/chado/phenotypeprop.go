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

// Phenotypeprop is an object representing the database table.
type Phenotypeprop struct {
	PhenotypepropID int         `boil:"phenotypeprop_id" json:"phenotypeprop_id" toml:"phenotypeprop_id" yaml:"phenotypeprop_id"`
	PhenotypeID     int         `boil:"phenotype_id" json:"phenotype_id" toml:"phenotype_id" yaml:"phenotype_id"`
	TypeID          int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value           null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank            int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *phenotypepropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L phenotypepropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// phenotypepropR is where relationships are stored.
type phenotypepropR struct {
	Phenotype *Phenotype
	Type      *Cvterm
}

// phenotypepropL is where Load methods for each relationship are stored.
type phenotypepropL struct{}

var (
	phenotypepropColumns               = []string{"phenotypeprop_id", "phenotype_id", "type_id", "value", "rank"}
	phenotypepropColumnsWithoutDefault = []string{"phenotype_id", "type_id", "value"}
	phenotypepropColumnsWithDefault    = []string{"phenotypeprop_id", "rank"}
	phenotypepropPrimaryKeyColumns     = []string{"phenotypeprop_id"}
)

type (
	// PhenotypepropSlice is an alias for a slice of pointers to Phenotypeprop.
	// This should generally be used opposed to []Phenotypeprop.
	PhenotypepropSlice []*Phenotypeprop
	// PhenotypepropHook is the signature for custom Phenotypeprop hook methods
	PhenotypepropHook func(boil.Executor, *Phenotypeprop) error

	phenotypepropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	phenotypepropType                 = reflect.TypeOf(&Phenotypeprop{})
	phenotypepropMapping              = queries.MakeStructMapping(phenotypepropType)
	phenotypepropPrimaryKeyMapping, _ = queries.BindMapping(phenotypepropType, phenotypepropMapping, phenotypepropPrimaryKeyColumns)
	phenotypepropInsertCacheMut       sync.RWMutex
	phenotypepropInsertCache          = make(map[string]insertCache)
	phenotypepropUpdateCacheMut       sync.RWMutex
	phenotypepropUpdateCache          = make(map[string]updateCache)
	phenotypepropUpsertCacheMut       sync.RWMutex
	phenotypepropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var phenotypepropBeforeInsertHooks []PhenotypepropHook
var phenotypepropBeforeUpdateHooks []PhenotypepropHook
var phenotypepropBeforeDeleteHooks []PhenotypepropHook
var phenotypepropBeforeUpsertHooks []PhenotypepropHook

var phenotypepropAfterInsertHooks []PhenotypepropHook
var phenotypepropAfterSelectHooks []PhenotypepropHook
var phenotypepropAfterUpdateHooks []PhenotypepropHook
var phenotypepropAfterDeleteHooks []PhenotypepropHook
var phenotypepropAfterUpsertHooks []PhenotypepropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Phenotypeprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypepropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Phenotypeprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypepropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Phenotypeprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypepropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Phenotypeprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypepropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Phenotypeprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypepropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Phenotypeprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypepropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Phenotypeprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypepropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Phenotypeprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypepropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Phenotypeprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypepropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPhenotypepropHook registers your hook function for all future operations.
func AddPhenotypepropHook(hookPoint boil.HookPoint, phenotypepropHook PhenotypepropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		phenotypepropBeforeInsertHooks = append(phenotypepropBeforeInsertHooks, phenotypepropHook)
	case boil.BeforeUpdateHook:
		phenotypepropBeforeUpdateHooks = append(phenotypepropBeforeUpdateHooks, phenotypepropHook)
	case boil.BeforeDeleteHook:
		phenotypepropBeforeDeleteHooks = append(phenotypepropBeforeDeleteHooks, phenotypepropHook)
	case boil.BeforeUpsertHook:
		phenotypepropBeforeUpsertHooks = append(phenotypepropBeforeUpsertHooks, phenotypepropHook)
	case boil.AfterInsertHook:
		phenotypepropAfterInsertHooks = append(phenotypepropAfterInsertHooks, phenotypepropHook)
	case boil.AfterSelectHook:
		phenotypepropAfterSelectHooks = append(phenotypepropAfterSelectHooks, phenotypepropHook)
	case boil.AfterUpdateHook:
		phenotypepropAfterUpdateHooks = append(phenotypepropAfterUpdateHooks, phenotypepropHook)
	case boil.AfterDeleteHook:
		phenotypepropAfterDeleteHooks = append(phenotypepropAfterDeleteHooks, phenotypepropHook)
	case boil.AfterUpsertHook:
		phenotypepropAfterUpsertHooks = append(phenotypepropAfterUpsertHooks, phenotypepropHook)
	}
}

// OneP returns a single phenotypeprop record from the query, and panics on error.
func (q phenotypepropQuery) OneP() *Phenotypeprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single phenotypeprop record from the query.
func (q phenotypepropQuery) One() (*Phenotypeprop, error) {
	o := &Phenotypeprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for phenotypeprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Phenotypeprop records from the query, and panics on error.
func (q phenotypepropQuery) AllP() PhenotypepropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Phenotypeprop records from the query.
func (q phenotypepropQuery) All() (PhenotypepropSlice, error) {
	var o PhenotypepropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Phenotypeprop slice")
	}

	if len(phenotypepropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Phenotypeprop records in the query, and panics on error.
func (q phenotypepropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Phenotypeprop records in the query.
func (q phenotypepropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count phenotypeprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q phenotypepropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q phenotypepropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if phenotypeprop exists")
	}

	return count > 0, nil
}

// PhenotypeG pointed to by the foreign key.
func (o *Phenotypeprop) PhenotypeG(mods ...qm.QueryMod) phenotypeQuery {
	return o.Phenotype(boil.GetDB(), mods...)
}

// Phenotype pointed to by the foreign key.
func (o *Phenotypeprop) Phenotype(exec boil.Executor, mods ...qm.QueryMod) phenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("phenotype_id=$1", o.PhenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Phenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Phenotypeprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Phenotypeprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadPhenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypepropL) LoadPhenotype(e boil.Executor, singular bool, maybePhenotypeprop interface{}) error {
	var slice []*Phenotypeprop
	var object *Phenotypeprop

	count := 1
	if singular {
		object = maybePhenotypeprop.(*Phenotypeprop)
	} else {
		slice = *maybePhenotypeprop.(*PhenotypepropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypepropR{}
		args[0] = object.PhenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypepropR{}
			args[i] = obj.PhenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype\" where \"phenotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Phenotype")
	}
	defer results.Close()

	var resultSlice []*Phenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Phenotype")
	}

	if len(phenotypepropAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Phenotype = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PhenotypeID == foreign.PhenotypeID {
				local.R.Phenotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypepropL) LoadType(e boil.Executor, singular bool, maybePhenotypeprop interface{}) error {
	var slice []*Phenotypeprop
	var object *Phenotypeprop

	count := 1
	if singular {
		object = maybePhenotypeprop.(*Phenotypeprop)
	} else {
		slice = *maybePhenotypeprop.(*PhenotypepropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypepropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypepropR{}
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

	if len(phenotypepropAfterSelectHooks) != 0 {
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

// SetPhenotype of the phenotypeprop to the related item.
// Sets o.R.Phenotype to related.
// Adds o to related.R.Phenotypeprop.
func (o *Phenotypeprop) SetPhenotype(exec boil.Executor, insert bool, related *Phenotype) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotypeprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"phenotype_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypepropPrimaryKeyColumns),
	)
	values := []interface{}{related.PhenotypeID, o.PhenotypepropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PhenotypeID = related.PhenotypeID

	if o.R == nil {
		o.R = &phenotypepropR{
			Phenotype: related,
		}
	} else {
		o.R.Phenotype = related
	}

	if related.R == nil {
		related.R = &phenotypeR{
			Phenotypeprop: o,
		}
	} else {
		related.R.Phenotypeprop = o
	}

	return nil
}

// SetType of the phenotypeprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypePhenotypeprop.
func (o *Phenotypeprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotypeprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypepropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.PhenotypepropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &phenotypepropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypePhenotypeprop: o,
		}
	} else {
		related.R.TypePhenotypeprop = o
	}

	return nil
}

// PhenotypepropsG retrieves all records.
func PhenotypepropsG(mods ...qm.QueryMod) phenotypepropQuery {
	return Phenotypeprops(boil.GetDB(), mods...)
}

// Phenotypeprops retrieves all the records using an executor.
func Phenotypeprops(exec boil.Executor, mods ...qm.QueryMod) phenotypepropQuery {
	mods = append(mods, qm.From("\"phenotypeprop\""))
	return phenotypepropQuery{NewQuery(exec, mods...)}
}

// FindPhenotypepropG retrieves a single record by ID.
func FindPhenotypepropG(phenotypepropID int, selectCols ...string) (*Phenotypeprop, error) {
	return FindPhenotypeprop(boil.GetDB(), phenotypepropID, selectCols...)
}

// FindPhenotypepropGP retrieves a single record by ID, and panics on error.
func FindPhenotypepropGP(phenotypepropID int, selectCols ...string) *Phenotypeprop {
	retobj, err := FindPhenotypeprop(boil.GetDB(), phenotypepropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindPhenotypeprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPhenotypeprop(exec boil.Executor, phenotypepropID int, selectCols ...string) (*Phenotypeprop, error) {
	phenotypepropObj := &Phenotypeprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"phenotypeprop\" where \"phenotypeprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, phenotypepropID)

	err := q.Bind(phenotypepropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from phenotypeprop")
	}

	return phenotypepropObj, nil
}

// FindPhenotypepropP retrieves a single record by ID with an executor, and panics on error.
func FindPhenotypepropP(exec boil.Executor, phenotypepropID int, selectCols ...string) *Phenotypeprop {
	retobj, err := FindPhenotypeprop(exec, phenotypepropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Phenotypeprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Phenotypeprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Phenotypeprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Phenotypeprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no phenotypeprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(phenotypepropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	phenotypepropInsertCacheMut.RLock()
	cache, cached := phenotypepropInsertCache[key]
	phenotypepropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			phenotypepropColumns,
			phenotypepropColumnsWithDefault,
			phenotypepropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(phenotypepropType, phenotypepropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(phenotypepropType, phenotypepropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"phenotypeprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into phenotypeprop")
	}

	if !cached {
		phenotypepropInsertCacheMut.Lock()
		phenotypepropInsertCache[key] = cache
		phenotypepropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Phenotypeprop record. See Update for
// whitelist behavior description.
func (o *Phenotypeprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Phenotypeprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Phenotypeprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Phenotypeprop, and panics on error.
// See Update for whitelist behavior description.
func (o *Phenotypeprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Phenotypeprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Phenotypeprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	phenotypepropUpdateCacheMut.RLock()
	cache, cached := phenotypepropUpdateCache[key]
	phenotypepropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(phenotypepropColumns, phenotypepropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update phenotypeprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"phenotypeprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, phenotypepropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(phenotypepropType, phenotypepropMapping, append(wl, phenotypepropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update phenotypeprop row")
	}

	if !cached {
		phenotypepropUpdateCacheMut.Lock()
		phenotypepropUpdateCache[key] = cache
		phenotypepropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q phenotypepropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q phenotypepropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for phenotypeprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PhenotypepropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o PhenotypepropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o PhenotypepropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PhenotypepropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenotypepropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"phenotypeprop\" SET %s WHERE (\"phenotypeprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(phenotypepropPrimaryKeyColumns), len(colNames)+1, len(phenotypepropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in phenotypeprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Phenotypeprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Phenotypeprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Phenotypeprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Phenotypeprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no phenotypeprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(phenotypepropColumnsWithDefault, o)

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

	phenotypepropUpsertCacheMut.RLock()
	cache, cached := phenotypepropUpsertCache[key]
	phenotypepropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			phenotypepropColumns,
			phenotypepropColumnsWithDefault,
			phenotypepropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			phenotypepropColumns,
			phenotypepropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert phenotypeprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(phenotypepropPrimaryKeyColumns))
			copy(conflict, phenotypepropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"phenotypeprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(phenotypepropType, phenotypepropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(phenotypepropType, phenotypepropMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for phenotypeprop")
	}

	if !cached {
		phenotypepropUpsertCacheMut.Lock()
		phenotypepropUpsertCache[key] = cache
		phenotypepropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Phenotypeprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Phenotypeprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Phenotypeprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Phenotypeprop) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Phenotypeprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Phenotypeprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Phenotypeprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Phenotypeprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Phenotypeprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Phenotypeprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), phenotypepropPrimaryKeyMapping)
	sql := "DELETE FROM \"phenotypeprop\" WHERE \"phenotypeprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from phenotypeprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q phenotypepropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q phenotypepropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no phenotypepropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from phenotypeprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o PhenotypepropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o PhenotypepropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Phenotypeprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o PhenotypepropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PhenotypepropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Phenotypeprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(phenotypepropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenotypepropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"phenotypeprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, phenotypepropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(phenotypepropPrimaryKeyColumns), 1, len(phenotypepropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from phenotypeprop slice")
	}

	if len(phenotypepropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Phenotypeprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Phenotypeprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Phenotypeprop) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Phenotypeprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Phenotypeprop) Reload(exec boil.Executor) error {
	ret, err := FindPhenotypeprop(exec, o.PhenotypepropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PhenotypepropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PhenotypepropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhenotypepropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty PhenotypepropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhenotypepropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	phenotypeprops := PhenotypepropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenotypepropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"phenotypeprop\".* FROM \"phenotypeprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, phenotypepropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(phenotypepropPrimaryKeyColumns), 1, len(phenotypepropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&phenotypeprops)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in PhenotypepropSlice")
	}

	*o = phenotypeprops

	return nil
}

// PhenotypepropExists checks if the Phenotypeprop row exists.
func PhenotypepropExists(exec boil.Executor, phenotypepropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"phenotypeprop\" where \"phenotypeprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, phenotypepropID)
	}

	row := exec.QueryRow(sql, phenotypepropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if phenotypeprop exists")
	}

	return exists, nil
}

// PhenotypepropExistsG checks if the Phenotypeprop row exists.
func PhenotypepropExistsG(phenotypepropID int) (bool, error) {
	return PhenotypepropExists(boil.GetDB(), phenotypepropID)
}

// PhenotypepropExistsGP checks if the Phenotypeprop row exists. Panics on error.
func PhenotypepropExistsGP(phenotypepropID int) bool {
	e, err := PhenotypepropExists(boil.GetDB(), phenotypepropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// PhenotypepropExistsP checks if the Phenotypeprop row exists. Panics on error.
func PhenotypepropExistsP(exec boil.Executor, phenotypepropID int) bool {
	e, err := PhenotypepropExists(exec, phenotypepropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
