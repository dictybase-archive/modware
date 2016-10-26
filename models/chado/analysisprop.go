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

// Analysisprop is an object representing the database table.
type Analysisprop struct {
	AnalysispropID int         `boil:"analysisprop_id" json:"analysisprop_id" toml:"analysisprop_id" yaml:"analysisprop_id"`
	AnalysisID     int         `boil:"analysis_id" json:"analysis_id" toml:"analysis_id" yaml:"analysis_id"`
	TypeID         int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value          null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank           int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *analysispropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L analysispropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// analysispropR is where relationships are stored.
type analysispropR struct {
	Analysi *Analysi
	Type    *Cvterm
}

// analysispropL is where Load methods for each relationship are stored.
type analysispropL struct{}

var (
	analysispropColumns               = []string{"analysisprop_id", "analysis_id", "type_id", "value", "rank"}
	analysispropColumnsWithoutDefault = []string{"analysis_id", "type_id", "value"}
	analysispropColumnsWithDefault    = []string{"analysisprop_id", "rank"}
	analysispropPrimaryKeyColumns     = []string{"analysisprop_id"}
)

type (
	// AnalysispropSlice is an alias for a slice of pointers to Analysisprop.
	// This should generally be used opposed to []Analysisprop.
	AnalysispropSlice []*Analysisprop
	// AnalysispropHook is the signature for custom Analysisprop hook methods
	AnalysispropHook func(boil.Executor, *Analysisprop) error

	analysispropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	analysispropType                 = reflect.TypeOf(&Analysisprop{})
	analysispropMapping              = queries.MakeStructMapping(analysispropType)
	analysispropPrimaryKeyMapping, _ = queries.BindMapping(analysispropType, analysispropMapping, analysispropPrimaryKeyColumns)
	analysispropInsertCacheMut       sync.RWMutex
	analysispropInsertCache          = make(map[string]insertCache)
	analysispropUpdateCacheMut       sync.RWMutex
	analysispropUpdateCache          = make(map[string]updateCache)
	analysispropUpsertCacheMut       sync.RWMutex
	analysispropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var analysispropBeforeInsertHooks []AnalysispropHook
var analysispropBeforeUpdateHooks []AnalysispropHook
var analysispropBeforeDeleteHooks []AnalysispropHook
var analysispropBeforeUpsertHooks []AnalysispropHook

var analysispropAfterInsertHooks []AnalysispropHook
var analysispropAfterSelectHooks []AnalysispropHook
var analysispropAfterUpdateHooks []AnalysispropHook
var analysispropAfterDeleteHooks []AnalysispropHook
var analysispropAfterUpsertHooks []AnalysispropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Analysisprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysispropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Analysisprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range analysispropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Analysisprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range analysispropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Analysisprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysispropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Analysisprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysispropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Analysisprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range analysispropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Analysisprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range analysispropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Analysisprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range analysispropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Analysisprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysispropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAnalysispropHook registers your hook function for all future operations.
func AddAnalysispropHook(hookPoint boil.HookPoint, analysispropHook AnalysispropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		analysispropBeforeInsertHooks = append(analysispropBeforeInsertHooks, analysispropHook)
	case boil.BeforeUpdateHook:
		analysispropBeforeUpdateHooks = append(analysispropBeforeUpdateHooks, analysispropHook)
	case boil.BeforeDeleteHook:
		analysispropBeforeDeleteHooks = append(analysispropBeforeDeleteHooks, analysispropHook)
	case boil.BeforeUpsertHook:
		analysispropBeforeUpsertHooks = append(analysispropBeforeUpsertHooks, analysispropHook)
	case boil.AfterInsertHook:
		analysispropAfterInsertHooks = append(analysispropAfterInsertHooks, analysispropHook)
	case boil.AfterSelectHook:
		analysispropAfterSelectHooks = append(analysispropAfterSelectHooks, analysispropHook)
	case boil.AfterUpdateHook:
		analysispropAfterUpdateHooks = append(analysispropAfterUpdateHooks, analysispropHook)
	case boil.AfterDeleteHook:
		analysispropAfterDeleteHooks = append(analysispropAfterDeleteHooks, analysispropHook)
	case boil.AfterUpsertHook:
		analysispropAfterUpsertHooks = append(analysispropAfterUpsertHooks, analysispropHook)
	}
}

// OneP returns a single analysisprop record from the query, and panics on error.
func (q analysispropQuery) OneP() *Analysisprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single analysisprop record from the query.
func (q analysispropQuery) One() (*Analysisprop, error) {
	o := &Analysisprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for analysisprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Analysisprop records from the query, and panics on error.
func (q analysispropQuery) AllP() AnalysispropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Analysisprop records from the query.
func (q analysispropQuery) All() (AnalysispropSlice, error) {
	var o AnalysispropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Analysisprop slice")
	}

	if len(analysispropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Analysisprop records in the query, and panics on error.
func (q analysispropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Analysisprop records in the query.
func (q analysispropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count analysisprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q analysispropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q analysispropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if analysisprop exists")
	}

	return count > 0, nil
}

// AnalysiG pointed to by the foreign key.
func (o *Analysisprop) AnalysiG(mods ...qm.QueryMod) analysiQuery {
	return o.Analysi(boil.GetDB(), mods...)
}

// Analysi pointed to by the foreign key.
func (o *Analysisprop) Analysi(exec boil.Executor, mods ...qm.QueryMod) analysiQuery {
	queryMods := []qm.QueryMod{
		qm.Where("analysis_id=$1", o.AnalysisID),
	}

	queryMods = append(queryMods, mods...)

	query := Analyses(exec, queryMods...)
	queries.SetFrom(query.Query, "\"analysis\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Analysisprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Analysisprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadAnalysi allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (analysispropL) LoadAnalysi(e boil.Executor, singular bool, maybeAnalysisprop interface{}) error {
	var slice []*Analysisprop
	var object *Analysisprop

	count := 1
	if singular {
		object = maybeAnalysisprop.(*Analysisprop)
	} else {
		slice = *maybeAnalysisprop.(*AnalysispropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &analysispropR{}
		args[0] = object.AnalysisID
	} else {
		for i, obj := range slice {
			obj.R = &analysispropR{}
			args[i] = obj.AnalysisID
		}
	}

	query := fmt.Sprintf(
		"select * from \"analysis\" where \"analysis_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Analysi")
	}
	defer results.Close()

	var resultSlice []*Analysi
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Analysi")
	}

	if len(analysispropAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Analysi = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AnalysisID == foreign.AnalysisID {
				local.R.Analysi = foreign
				break
			}
		}
	}

	return nil
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (analysispropL) LoadType(e boil.Executor, singular bool, maybeAnalysisprop interface{}) error {
	var slice []*Analysisprop
	var object *Analysisprop

	count := 1
	if singular {
		object = maybeAnalysisprop.(*Analysisprop)
	} else {
		slice = *maybeAnalysisprop.(*AnalysispropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &analysispropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &analysispropR{}
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

	if len(analysispropAfterSelectHooks) != 0 {
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

// SetAnalysi of the analysisprop to the related item.
// Sets o.R.Analysi to related.
// Adds o to related.R.Analysisprop.
func (o *Analysisprop) SetAnalysi(exec boil.Executor, insert bool, related *Analysi) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"analysisprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"analysis_id"}),
		strmangle.WhereClause("\"", "\"", 2, analysispropPrimaryKeyColumns),
	)
	values := []interface{}{related.AnalysisID, o.AnalysispropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AnalysisID = related.AnalysisID

	if o.R == nil {
		o.R = &analysispropR{
			Analysi: related,
		}
	} else {
		o.R.Analysi = related
	}

	if related.R == nil {
		related.R = &analysiR{
			Analysisprop: o,
		}
	} else {
		related.R.Analysisprop = o
	}

	return nil
}

// SetType of the analysisprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeAnalysisprop.
func (o *Analysisprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"analysisprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, analysispropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.AnalysispropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &analysispropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeAnalysisprop: o,
		}
	} else {
		related.R.TypeAnalysisprop = o
	}

	return nil
}

// AnalysispropsG retrieves all records.
func AnalysispropsG(mods ...qm.QueryMod) analysispropQuery {
	return Analysisprops(boil.GetDB(), mods...)
}

// Analysisprops retrieves all the records using an executor.
func Analysisprops(exec boil.Executor, mods ...qm.QueryMod) analysispropQuery {
	mods = append(mods, qm.From("\"analysisprop\""))
	return analysispropQuery{NewQuery(exec, mods...)}
}

// FindAnalysispropG retrieves a single record by ID.
func FindAnalysispropG(analysispropID int, selectCols ...string) (*Analysisprop, error) {
	return FindAnalysisprop(boil.GetDB(), analysispropID, selectCols...)
}

// FindAnalysispropGP retrieves a single record by ID, and panics on error.
func FindAnalysispropGP(analysispropID int, selectCols ...string) *Analysisprop {
	retobj, err := FindAnalysisprop(boil.GetDB(), analysispropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAnalysisprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAnalysisprop(exec boil.Executor, analysispropID int, selectCols ...string) (*Analysisprop, error) {
	analysispropObj := &Analysisprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"analysisprop\" where \"analysisprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, analysispropID)

	err := q.Bind(analysispropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from analysisprop")
	}

	return analysispropObj, nil
}

// FindAnalysispropP retrieves a single record by ID with an executor, and panics on error.
func FindAnalysispropP(exec boil.Executor, analysispropID int, selectCols ...string) *Analysisprop {
	retobj, err := FindAnalysisprop(exec, analysispropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Analysisprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Analysisprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Analysisprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Analysisprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no analysisprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(analysispropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	analysispropInsertCacheMut.RLock()
	cache, cached := analysispropInsertCache[key]
	analysispropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			analysispropColumns,
			analysispropColumnsWithDefault,
			analysispropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(analysispropType, analysispropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(analysispropType, analysispropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"analysisprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into analysisprop")
	}

	if !cached {
		analysispropInsertCacheMut.Lock()
		analysispropInsertCache[key] = cache
		analysispropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Analysisprop record. See Update for
// whitelist behavior description.
func (o *Analysisprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Analysisprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Analysisprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Analysisprop, and panics on error.
// See Update for whitelist behavior description.
func (o *Analysisprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Analysisprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Analysisprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	analysispropUpdateCacheMut.RLock()
	cache, cached := analysispropUpdateCache[key]
	analysispropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(analysispropColumns, analysispropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update analysisprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"analysisprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, analysispropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(analysispropType, analysispropMapping, append(wl, analysispropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update analysisprop row")
	}

	if !cached {
		analysispropUpdateCacheMut.Lock()
		analysispropUpdateCache[key] = cache
		analysispropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q analysispropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q analysispropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for analysisprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AnalysispropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AnalysispropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AnalysispropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AnalysispropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), analysispropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"analysisprop\" SET %s WHERE (\"analysisprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(analysispropPrimaryKeyColumns), len(colNames)+1, len(analysispropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in analysisprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Analysisprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Analysisprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Analysisprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Analysisprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no analysisprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(analysispropColumnsWithDefault, o)

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

	analysispropUpsertCacheMut.RLock()
	cache, cached := analysispropUpsertCache[key]
	analysispropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			analysispropColumns,
			analysispropColumnsWithDefault,
			analysispropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			analysispropColumns,
			analysispropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert analysisprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(analysispropPrimaryKeyColumns))
			copy(conflict, analysispropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"analysisprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(analysispropType, analysispropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(analysispropType, analysispropMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for analysisprop")
	}

	if !cached {
		analysispropUpsertCacheMut.Lock()
		analysispropUpsertCache[key] = cache
		analysispropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Analysisprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Analysisprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Analysisprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Analysisprop) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Analysisprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Analysisprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Analysisprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Analysisprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Analysisprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Analysisprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), analysispropPrimaryKeyMapping)
	sql := "DELETE FROM \"analysisprop\" WHERE \"analysisprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from analysisprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q analysispropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q analysispropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no analysispropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from analysisprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AnalysispropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AnalysispropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Analysisprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AnalysispropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AnalysispropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Analysisprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(analysispropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), analysispropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"analysisprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, analysispropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(analysispropPrimaryKeyColumns), 1, len(analysispropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from analysisprop slice")
	}

	if len(analysispropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Analysisprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Analysisprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Analysisprop) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Analysisprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Analysisprop) Reload(exec boil.Executor) error {
	ret, err := FindAnalysisprop(exec, o.AnalysispropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AnalysispropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AnalysispropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnalysispropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty AnalysispropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnalysispropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	analysisprops := AnalysispropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), analysispropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"analysisprop\".* FROM \"analysisprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, analysispropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(analysispropPrimaryKeyColumns), 1, len(analysispropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&analysisprops)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in AnalysispropSlice")
	}

	*o = analysisprops

	return nil
}

// AnalysispropExists checks if the Analysisprop row exists.
func AnalysispropExists(exec boil.Executor, analysispropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"analysisprop\" where \"analysisprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, analysispropID)
	}

	row := exec.QueryRow(sql, analysispropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if analysisprop exists")
	}

	return exists, nil
}

// AnalysispropExistsG checks if the Analysisprop row exists.
func AnalysispropExistsG(analysispropID int) (bool, error) {
	return AnalysispropExists(boil.GetDB(), analysispropID)
}

// AnalysispropExistsGP checks if the Analysisprop row exists. Panics on error.
func AnalysispropExistsGP(analysispropID int) bool {
	e, err := AnalysispropExists(boil.GetDB(), analysispropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AnalysispropExistsP checks if the Analysisprop row exists. Panics on error.
func AnalysispropExistsP(exec boil.Executor, analysispropID int) bool {
	e, err := AnalysispropExists(exec, analysispropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
