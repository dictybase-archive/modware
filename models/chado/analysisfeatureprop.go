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

// Analysisfeatureprop is an object representing the database table.
type Analysisfeatureprop struct {
	AnalysisfeaturepropID int         `boil:"analysisfeatureprop_id" json:"analysisfeatureprop_id" toml:"analysisfeatureprop_id" yaml:"analysisfeatureprop_id"`
	AnalysisfeatureID     int         `boil:"analysisfeature_id" json:"analysisfeature_id" toml:"analysisfeature_id" yaml:"analysisfeature_id"`
	TypeID                int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value                 null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank                  int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *analysisfeaturepropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L analysisfeaturepropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// analysisfeaturepropR is where relationships are stored.
type analysisfeaturepropR struct {
	Analysisfeature *Analysisfeature
	Type            *Cvterm
}

// analysisfeaturepropL is where Load methods for each relationship are stored.
type analysisfeaturepropL struct{}

var (
	analysisfeaturepropColumns               = []string{"analysisfeatureprop_id", "analysisfeature_id", "type_id", "value", "rank"}
	analysisfeaturepropColumnsWithoutDefault = []string{"analysisfeature_id", "type_id", "value", "rank"}
	analysisfeaturepropColumnsWithDefault    = []string{"analysisfeatureprop_id"}
	analysisfeaturepropPrimaryKeyColumns     = []string{"analysisfeatureprop_id"}
)

type (
	// AnalysisfeaturepropSlice is an alias for a slice of pointers to Analysisfeatureprop.
	// This should generally be used opposed to []Analysisfeatureprop.
	AnalysisfeaturepropSlice []*Analysisfeatureprop
	// AnalysisfeaturepropHook is the signature for custom Analysisfeatureprop hook methods
	AnalysisfeaturepropHook func(boil.Executor, *Analysisfeatureprop) error

	analysisfeaturepropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	analysisfeaturepropType                 = reflect.TypeOf(&Analysisfeatureprop{})
	analysisfeaturepropMapping              = queries.MakeStructMapping(analysisfeaturepropType)
	analysisfeaturepropPrimaryKeyMapping, _ = queries.BindMapping(analysisfeaturepropType, analysisfeaturepropMapping, analysisfeaturepropPrimaryKeyColumns)
	analysisfeaturepropInsertCacheMut       sync.RWMutex
	analysisfeaturepropInsertCache          = make(map[string]insertCache)
	analysisfeaturepropUpdateCacheMut       sync.RWMutex
	analysisfeaturepropUpdateCache          = make(map[string]updateCache)
	analysisfeaturepropUpsertCacheMut       sync.RWMutex
	analysisfeaturepropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var analysisfeaturepropBeforeInsertHooks []AnalysisfeaturepropHook
var analysisfeaturepropBeforeUpdateHooks []AnalysisfeaturepropHook
var analysisfeaturepropBeforeDeleteHooks []AnalysisfeaturepropHook
var analysisfeaturepropBeforeUpsertHooks []AnalysisfeaturepropHook

var analysisfeaturepropAfterInsertHooks []AnalysisfeaturepropHook
var analysisfeaturepropAfterSelectHooks []AnalysisfeaturepropHook
var analysisfeaturepropAfterUpdateHooks []AnalysisfeaturepropHook
var analysisfeaturepropAfterDeleteHooks []AnalysisfeaturepropHook
var analysisfeaturepropAfterUpsertHooks []AnalysisfeaturepropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Analysisfeatureprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeaturepropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Analysisfeatureprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeaturepropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Analysisfeatureprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeaturepropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Analysisfeatureprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeaturepropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Analysisfeatureprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeaturepropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Analysisfeatureprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeaturepropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Analysisfeatureprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeaturepropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Analysisfeatureprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeaturepropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Analysisfeatureprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeaturepropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAnalysisfeaturepropHook registers your hook function for all future operations.
func AddAnalysisfeaturepropHook(hookPoint boil.HookPoint, analysisfeaturepropHook AnalysisfeaturepropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		analysisfeaturepropBeforeInsertHooks = append(analysisfeaturepropBeforeInsertHooks, analysisfeaturepropHook)
	case boil.BeforeUpdateHook:
		analysisfeaturepropBeforeUpdateHooks = append(analysisfeaturepropBeforeUpdateHooks, analysisfeaturepropHook)
	case boil.BeforeDeleteHook:
		analysisfeaturepropBeforeDeleteHooks = append(analysisfeaturepropBeforeDeleteHooks, analysisfeaturepropHook)
	case boil.BeforeUpsertHook:
		analysisfeaturepropBeforeUpsertHooks = append(analysisfeaturepropBeforeUpsertHooks, analysisfeaturepropHook)
	case boil.AfterInsertHook:
		analysisfeaturepropAfterInsertHooks = append(analysisfeaturepropAfterInsertHooks, analysisfeaturepropHook)
	case boil.AfterSelectHook:
		analysisfeaturepropAfterSelectHooks = append(analysisfeaturepropAfterSelectHooks, analysisfeaturepropHook)
	case boil.AfterUpdateHook:
		analysisfeaturepropAfterUpdateHooks = append(analysisfeaturepropAfterUpdateHooks, analysisfeaturepropHook)
	case boil.AfterDeleteHook:
		analysisfeaturepropAfterDeleteHooks = append(analysisfeaturepropAfterDeleteHooks, analysisfeaturepropHook)
	case boil.AfterUpsertHook:
		analysisfeaturepropAfterUpsertHooks = append(analysisfeaturepropAfterUpsertHooks, analysisfeaturepropHook)
	}
}

// OneP returns a single analysisfeatureprop record from the query, and panics on error.
func (q analysisfeaturepropQuery) OneP() *Analysisfeatureprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single analysisfeatureprop record from the query.
func (q analysisfeaturepropQuery) One() (*Analysisfeatureprop, error) {
	o := &Analysisfeatureprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for analysisfeatureprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Analysisfeatureprop records from the query, and panics on error.
func (q analysisfeaturepropQuery) AllP() AnalysisfeaturepropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Analysisfeatureprop records from the query.
func (q analysisfeaturepropQuery) All() (AnalysisfeaturepropSlice, error) {
	var o AnalysisfeaturepropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Analysisfeatureprop slice")
	}

	if len(analysisfeaturepropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Analysisfeatureprop records in the query, and panics on error.
func (q analysisfeaturepropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Analysisfeatureprop records in the query.
func (q analysisfeaturepropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count analysisfeatureprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q analysisfeaturepropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q analysisfeaturepropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if analysisfeatureprop exists")
	}

	return count > 0, nil
}

// AnalysisfeatureG pointed to by the foreign key.
func (o *Analysisfeatureprop) AnalysisfeatureG(mods ...qm.QueryMod) analysisfeatureQuery {
	return o.Analysisfeature(boil.GetDB(), mods...)
}

// Analysisfeature pointed to by the foreign key.
func (o *Analysisfeatureprop) Analysisfeature(exec boil.Executor, mods ...qm.QueryMod) analysisfeatureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("analysisfeature_id=$1", o.AnalysisfeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Analysisfeatures(exec, queryMods...)
	queries.SetFrom(query.Query, "\"analysisfeature\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Analysisfeatureprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Analysisfeatureprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadAnalysisfeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (analysisfeaturepropL) LoadAnalysisfeature(e boil.Executor, singular bool, maybeAnalysisfeatureprop interface{}) error {
	var slice []*Analysisfeatureprop
	var object *Analysisfeatureprop

	count := 1
	if singular {
		object = maybeAnalysisfeatureprop.(*Analysisfeatureprop)
	} else {
		slice = *maybeAnalysisfeatureprop.(*AnalysisfeaturepropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &analysisfeaturepropR{}
		args[0] = object.AnalysisfeatureID
	} else {
		for i, obj := range slice {
			obj.R = &analysisfeaturepropR{}
			args[i] = obj.AnalysisfeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"analysisfeature\" where \"analysisfeature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Analysisfeature")
	}
	defer results.Close()

	var resultSlice []*Analysisfeature
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Analysisfeature")
	}

	if len(analysisfeaturepropAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Analysisfeature = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AnalysisfeatureID == foreign.AnalysisfeatureID {
				local.R.Analysisfeature = foreign
				break
			}
		}
	}

	return nil
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (analysisfeaturepropL) LoadType(e boil.Executor, singular bool, maybeAnalysisfeatureprop interface{}) error {
	var slice []*Analysisfeatureprop
	var object *Analysisfeatureprop

	count := 1
	if singular {
		object = maybeAnalysisfeatureprop.(*Analysisfeatureprop)
	} else {
		slice = *maybeAnalysisfeatureprop.(*AnalysisfeaturepropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &analysisfeaturepropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &analysisfeaturepropR{}
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

	if len(analysisfeaturepropAfterSelectHooks) != 0 {
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

// SetAnalysisfeature of the analysisfeatureprop to the related item.
// Sets o.R.Analysisfeature to related.
// Adds o to related.R.Analysisfeatureprop.
func (o *Analysisfeatureprop) SetAnalysisfeature(exec boil.Executor, insert bool, related *Analysisfeature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"analysisfeatureprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"analysisfeature_id"}),
		strmangle.WhereClause("\"", "\"", 2, analysisfeaturepropPrimaryKeyColumns),
	)
	values := []interface{}{related.AnalysisfeatureID, o.AnalysisfeaturepropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AnalysisfeatureID = related.AnalysisfeatureID

	if o.R == nil {
		o.R = &analysisfeaturepropR{
			Analysisfeature: related,
		}
	} else {
		o.R.Analysisfeature = related
	}

	if related.R == nil {
		related.R = &analysisfeatureR{
			Analysisfeatureprop: o,
		}
	} else {
		related.R.Analysisfeatureprop = o
	}

	return nil
}

// SetType of the analysisfeatureprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeAnalysisfeatureprop.
func (o *Analysisfeatureprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"analysisfeatureprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, analysisfeaturepropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.AnalysisfeaturepropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &analysisfeaturepropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeAnalysisfeatureprop: o,
		}
	} else {
		related.R.TypeAnalysisfeatureprop = o
	}

	return nil
}

// AnalysisfeaturepropsG retrieves all records.
func AnalysisfeaturepropsG(mods ...qm.QueryMod) analysisfeaturepropQuery {
	return Analysisfeatureprops(boil.GetDB(), mods...)
}

// Analysisfeatureprops retrieves all the records using an executor.
func Analysisfeatureprops(exec boil.Executor, mods ...qm.QueryMod) analysisfeaturepropQuery {
	mods = append(mods, qm.From("\"analysisfeatureprop\""))
	return analysisfeaturepropQuery{NewQuery(exec, mods...)}
}

// FindAnalysisfeaturepropG retrieves a single record by ID.
func FindAnalysisfeaturepropG(analysisfeaturepropID int, selectCols ...string) (*Analysisfeatureprop, error) {
	return FindAnalysisfeatureprop(boil.GetDB(), analysisfeaturepropID, selectCols...)
}

// FindAnalysisfeaturepropGP retrieves a single record by ID, and panics on error.
func FindAnalysisfeaturepropGP(analysisfeaturepropID int, selectCols ...string) *Analysisfeatureprop {
	retobj, err := FindAnalysisfeatureprop(boil.GetDB(), analysisfeaturepropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAnalysisfeatureprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAnalysisfeatureprop(exec boil.Executor, analysisfeaturepropID int, selectCols ...string) (*Analysisfeatureprop, error) {
	analysisfeaturepropObj := &Analysisfeatureprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"analysisfeatureprop\" where \"analysisfeatureprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, analysisfeaturepropID)

	err := q.Bind(analysisfeaturepropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from analysisfeatureprop")
	}

	return analysisfeaturepropObj, nil
}

// FindAnalysisfeaturepropP retrieves a single record by ID with an executor, and panics on error.
func FindAnalysisfeaturepropP(exec boil.Executor, analysisfeaturepropID int, selectCols ...string) *Analysisfeatureprop {
	retobj, err := FindAnalysisfeatureprop(exec, analysisfeaturepropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Analysisfeatureprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Analysisfeatureprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Analysisfeatureprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Analysisfeatureprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no analysisfeatureprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(analysisfeaturepropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	analysisfeaturepropInsertCacheMut.RLock()
	cache, cached := analysisfeaturepropInsertCache[key]
	analysisfeaturepropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			analysisfeaturepropColumns,
			analysisfeaturepropColumnsWithDefault,
			analysisfeaturepropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(analysisfeaturepropType, analysisfeaturepropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(analysisfeaturepropType, analysisfeaturepropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"analysisfeatureprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into analysisfeatureprop")
	}

	if !cached {
		analysisfeaturepropInsertCacheMut.Lock()
		analysisfeaturepropInsertCache[key] = cache
		analysisfeaturepropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Analysisfeatureprop record. See Update for
// whitelist behavior description.
func (o *Analysisfeatureprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Analysisfeatureprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Analysisfeatureprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Analysisfeatureprop, and panics on error.
// See Update for whitelist behavior description.
func (o *Analysisfeatureprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Analysisfeatureprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Analysisfeatureprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	analysisfeaturepropUpdateCacheMut.RLock()
	cache, cached := analysisfeaturepropUpdateCache[key]
	analysisfeaturepropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(analysisfeaturepropColumns, analysisfeaturepropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update analysisfeatureprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"analysisfeatureprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, analysisfeaturepropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(analysisfeaturepropType, analysisfeaturepropMapping, append(wl, analysisfeaturepropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update analysisfeatureprop row")
	}

	if !cached {
		analysisfeaturepropUpdateCacheMut.Lock()
		analysisfeaturepropUpdateCache[key] = cache
		analysisfeaturepropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q analysisfeaturepropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q analysisfeaturepropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for analysisfeatureprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AnalysisfeaturepropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AnalysisfeaturepropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AnalysisfeaturepropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AnalysisfeaturepropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), analysisfeaturepropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"analysisfeatureprop\" SET %s WHERE (\"analysisfeatureprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(analysisfeaturepropPrimaryKeyColumns), len(colNames)+1, len(analysisfeaturepropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in analysisfeatureprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Analysisfeatureprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Analysisfeatureprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Analysisfeatureprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Analysisfeatureprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no analysisfeatureprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(analysisfeaturepropColumnsWithDefault, o)

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

	analysisfeaturepropUpsertCacheMut.RLock()
	cache, cached := analysisfeaturepropUpsertCache[key]
	analysisfeaturepropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			analysisfeaturepropColumns,
			analysisfeaturepropColumnsWithDefault,
			analysisfeaturepropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			analysisfeaturepropColumns,
			analysisfeaturepropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert analysisfeatureprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(analysisfeaturepropPrimaryKeyColumns))
			copy(conflict, analysisfeaturepropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"analysisfeatureprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(analysisfeaturepropType, analysisfeaturepropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(analysisfeaturepropType, analysisfeaturepropMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for analysisfeatureprop")
	}

	if !cached {
		analysisfeaturepropUpsertCacheMut.Lock()
		analysisfeaturepropUpsertCache[key] = cache
		analysisfeaturepropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Analysisfeatureprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Analysisfeatureprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Analysisfeatureprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Analysisfeatureprop) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Analysisfeatureprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Analysisfeatureprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Analysisfeatureprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Analysisfeatureprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Analysisfeatureprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Analysisfeatureprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), analysisfeaturepropPrimaryKeyMapping)
	sql := "DELETE FROM \"analysisfeatureprop\" WHERE \"analysisfeatureprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from analysisfeatureprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q analysisfeaturepropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q analysisfeaturepropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no analysisfeaturepropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from analysisfeatureprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AnalysisfeaturepropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AnalysisfeaturepropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Analysisfeatureprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AnalysisfeaturepropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AnalysisfeaturepropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Analysisfeatureprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(analysisfeaturepropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), analysisfeaturepropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"analysisfeatureprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, analysisfeaturepropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(analysisfeaturepropPrimaryKeyColumns), 1, len(analysisfeaturepropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from analysisfeatureprop slice")
	}

	if len(analysisfeaturepropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Analysisfeatureprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Analysisfeatureprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Analysisfeatureprop) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Analysisfeatureprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Analysisfeatureprop) Reload(exec boil.Executor) error {
	ret, err := FindAnalysisfeatureprop(exec, o.AnalysisfeaturepropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AnalysisfeaturepropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AnalysisfeaturepropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnalysisfeaturepropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty AnalysisfeaturepropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnalysisfeaturepropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	analysisfeatureprops := AnalysisfeaturepropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), analysisfeaturepropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"analysisfeatureprop\".* FROM \"analysisfeatureprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, analysisfeaturepropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(analysisfeaturepropPrimaryKeyColumns), 1, len(analysisfeaturepropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&analysisfeatureprops)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in AnalysisfeaturepropSlice")
	}

	*o = analysisfeatureprops

	return nil
}

// AnalysisfeaturepropExists checks if the Analysisfeatureprop row exists.
func AnalysisfeaturepropExists(exec boil.Executor, analysisfeaturepropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"analysisfeatureprop\" where \"analysisfeatureprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, analysisfeaturepropID)
	}

	row := exec.QueryRow(sql, analysisfeaturepropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if analysisfeatureprop exists")
	}

	return exists, nil
}

// AnalysisfeaturepropExistsG checks if the Analysisfeatureprop row exists.
func AnalysisfeaturepropExistsG(analysisfeaturepropID int) (bool, error) {
	return AnalysisfeaturepropExists(boil.GetDB(), analysisfeaturepropID)
}

// AnalysisfeaturepropExistsGP checks if the Analysisfeatureprop row exists. Panics on error.
func AnalysisfeaturepropExistsGP(analysisfeaturepropID int) bool {
	e, err := AnalysisfeaturepropExists(boil.GetDB(), analysisfeaturepropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AnalysisfeaturepropExistsP checks if the Analysisfeatureprop row exists. Panics on error.
func AnalysisfeaturepropExistsP(exec boil.Executor, analysisfeaturepropID int) bool {
	e, err := AnalysisfeaturepropExists(exec, analysisfeaturepropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
