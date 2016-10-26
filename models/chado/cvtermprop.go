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

// Cvtermprop is an object representing the database table.
type Cvtermprop struct {
	CvtermpropID int    `boil:"cvtermprop_id" json:"cvtermprop_id" toml:"cvtermprop_id" yaml:"cvtermprop_id"`
	CvtermID     int    `boil:"cvterm_id" json:"cvterm_id" toml:"cvterm_id" yaml:"cvterm_id"`
	TypeID       int    `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value        string `boil:"value" json:"value" toml:"value" yaml:"value"`
	Rank         int    `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *cvtermpropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cvtermpropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// cvtermpropR is where relationships are stored.
type cvtermpropR struct {
	Cvterm *Cvterm
	Type   *Cvterm
}

// cvtermpropL is where Load methods for each relationship are stored.
type cvtermpropL struct{}

var (
	cvtermpropColumns               = []string{"cvtermprop_id", "cvterm_id", "type_id", "value", "rank"}
	cvtermpropColumnsWithoutDefault = []string{"cvterm_id", "type_id"}
	cvtermpropColumnsWithDefault    = []string{"cvtermprop_id", "value", "rank"}
	cvtermpropPrimaryKeyColumns     = []string{"cvtermprop_id"}
)

type (
	// CvtermpropSlice is an alias for a slice of pointers to Cvtermprop.
	// This should generally be used opposed to []Cvtermprop.
	CvtermpropSlice []*Cvtermprop
	// CvtermpropHook is the signature for custom Cvtermprop hook methods
	CvtermpropHook func(boil.Executor, *Cvtermprop) error

	cvtermpropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cvtermpropType                 = reflect.TypeOf(&Cvtermprop{})
	cvtermpropMapping              = queries.MakeStructMapping(cvtermpropType)
	cvtermpropPrimaryKeyMapping, _ = queries.BindMapping(cvtermpropType, cvtermpropMapping, cvtermpropPrimaryKeyColumns)
	cvtermpropInsertCacheMut       sync.RWMutex
	cvtermpropInsertCache          = make(map[string]insertCache)
	cvtermpropUpdateCacheMut       sync.RWMutex
	cvtermpropUpdateCache          = make(map[string]updateCache)
	cvtermpropUpsertCacheMut       sync.RWMutex
	cvtermpropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var cvtermpropBeforeInsertHooks []CvtermpropHook
var cvtermpropBeforeUpdateHooks []CvtermpropHook
var cvtermpropBeforeDeleteHooks []CvtermpropHook
var cvtermpropBeforeUpsertHooks []CvtermpropHook

var cvtermpropAfterInsertHooks []CvtermpropHook
var cvtermpropAfterSelectHooks []CvtermpropHook
var cvtermpropAfterUpdateHooks []CvtermpropHook
var cvtermpropAfterDeleteHooks []CvtermpropHook
var cvtermpropAfterUpsertHooks []CvtermpropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Cvtermprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Cvtermprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Cvtermprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Cvtermprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Cvtermprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Cvtermprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Cvtermprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Cvtermprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Cvtermprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermpropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCvtermpropHook registers your hook function for all future operations.
func AddCvtermpropHook(hookPoint boil.HookPoint, cvtermpropHook CvtermpropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cvtermpropBeforeInsertHooks = append(cvtermpropBeforeInsertHooks, cvtermpropHook)
	case boil.BeforeUpdateHook:
		cvtermpropBeforeUpdateHooks = append(cvtermpropBeforeUpdateHooks, cvtermpropHook)
	case boil.BeforeDeleteHook:
		cvtermpropBeforeDeleteHooks = append(cvtermpropBeforeDeleteHooks, cvtermpropHook)
	case boil.BeforeUpsertHook:
		cvtermpropBeforeUpsertHooks = append(cvtermpropBeforeUpsertHooks, cvtermpropHook)
	case boil.AfterInsertHook:
		cvtermpropAfterInsertHooks = append(cvtermpropAfterInsertHooks, cvtermpropHook)
	case boil.AfterSelectHook:
		cvtermpropAfterSelectHooks = append(cvtermpropAfterSelectHooks, cvtermpropHook)
	case boil.AfterUpdateHook:
		cvtermpropAfterUpdateHooks = append(cvtermpropAfterUpdateHooks, cvtermpropHook)
	case boil.AfterDeleteHook:
		cvtermpropAfterDeleteHooks = append(cvtermpropAfterDeleteHooks, cvtermpropHook)
	case boil.AfterUpsertHook:
		cvtermpropAfterUpsertHooks = append(cvtermpropAfterUpsertHooks, cvtermpropHook)
	}
}

// OneP returns a single cvtermprop record from the query, and panics on error.
func (q cvtermpropQuery) OneP() *Cvtermprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single cvtermprop record from the query.
func (q cvtermpropQuery) One() (*Cvtermprop, error) {
	o := &Cvtermprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for cvtermprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Cvtermprop records from the query, and panics on error.
func (q cvtermpropQuery) AllP() CvtermpropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Cvtermprop records from the query.
func (q cvtermpropQuery) All() (CvtermpropSlice, error) {
	var o CvtermpropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Cvtermprop slice")
	}

	if len(cvtermpropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Cvtermprop records in the query, and panics on error.
func (q cvtermpropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Cvtermprop records in the query.
func (q cvtermpropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count cvtermprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q cvtermpropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q cvtermpropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if cvtermprop exists")
	}

	return count > 0, nil
}

// CvtermG pointed to by the foreign key.
func (o *Cvtermprop) CvtermG(mods ...qm.QueryMod) cvtermQuery {
	return o.Cvterm(boil.GetDB(), mods...)
}

// Cvterm pointed to by the foreign key.
func (o *Cvtermprop) Cvterm(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Cvtermprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Cvtermprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermpropL) LoadCvterm(e boil.Executor, singular bool, maybeCvtermprop interface{}) error {
	var slice []*Cvtermprop
	var object *Cvtermprop

	count := 1
	if singular {
		object = maybeCvtermprop.(*Cvtermprop)
	} else {
		slice = *maybeCvtermprop.(*CvtermpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermpropR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermpropR{}
			args[i] = obj.CvtermID
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

	if len(cvtermpropAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Cvterm = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.CvtermID {
				local.R.Cvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermpropL) LoadType(e boil.Executor, singular bool, maybeCvtermprop interface{}) error {
	var slice []*Cvtermprop
	var object *Cvtermprop

	count := 1
	if singular {
		object = maybeCvtermprop.(*Cvtermprop)
	} else {
		slice = *maybeCvtermprop.(*CvtermpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermpropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermpropR{}
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

	if len(cvtermpropAfterSelectHooks) != 0 {
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

// SetCvterm of the cvtermprop to the related item.
// Sets o.R.Cvterm to related.
// Adds o to related.R.Cvtermprop.
func (o *Cvtermprop) SetCvterm(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvtermprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermpropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.CvtermpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CvtermID = related.CvtermID

	if o.R == nil {
		o.R = &cvtermpropR{
			Cvterm: related,
		}
	} else {
		o.R.Cvterm = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			Cvtermprop: o,
		}
	} else {
		related.R.Cvtermprop = o
	}

	return nil
}

// SetType of the cvtermprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeCvtermprop.
func (o *Cvtermprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvtermprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermpropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.CvtermpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &cvtermpropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeCvtermprop: o,
		}
	} else {
		related.R.TypeCvtermprop = o
	}

	return nil
}

// CvtermpropsG retrieves all records.
func CvtermpropsG(mods ...qm.QueryMod) cvtermpropQuery {
	return Cvtermprops(boil.GetDB(), mods...)
}

// Cvtermprops retrieves all the records using an executor.
func Cvtermprops(exec boil.Executor, mods ...qm.QueryMod) cvtermpropQuery {
	mods = append(mods, qm.From("\"cvtermprop\""))
	return cvtermpropQuery{NewQuery(exec, mods...)}
}

// FindCvtermpropG retrieves a single record by ID.
func FindCvtermpropG(cvtermpropID int, selectCols ...string) (*Cvtermprop, error) {
	return FindCvtermprop(boil.GetDB(), cvtermpropID, selectCols...)
}

// FindCvtermpropGP retrieves a single record by ID, and panics on error.
func FindCvtermpropGP(cvtermpropID int, selectCols ...string) *Cvtermprop {
	retobj, err := FindCvtermprop(boil.GetDB(), cvtermpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCvtermprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCvtermprop(exec boil.Executor, cvtermpropID int, selectCols ...string) (*Cvtermprop, error) {
	cvtermpropObj := &Cvtermprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"cvtermprop\" where \"cvtermprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, cvtermpropID)

	err := q.Bind(cvtermpropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from cvtermprop")
	}

	return cvtermpropObj, nil
}

// FindCvtermpropP retrieves a single record by ID with an executor, and panics on error.
func FindCvtermpropP(exec boil.Executor, cvtermpropID int, selectCols ...string) *Cvtermprop {
	retobj, err := FindCvtermprop(exec, cvtermpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Cvtermprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Cvtermprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Cvtermprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Cvtermprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no cvtermprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvtermpropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	cvtermpropInsertCacheMut.RLock()
	cache, cached := cvtermpropInsertCache[key]
	cvtermpropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			cvtermpropColumns,
			cvtermpropColumnsWithDefault,
			cvtermpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(cvtermpropType, cvtermpropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cvtermpropType, cvtermpropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"cvtermprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into cvtermprop")
	}

	if !cached {
		cvtermpropInsertCacheMut.Lock()
		cvtermpropInsertCache[key] = cache
		cvtermpropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Cvtermprop record. See Update for
// whitelist behavior description.
func (o *Cvtermprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Cvtermprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Cvtermprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Cvtermprop, and panics on error.
// See Update for whitelist behavior description.
func (o *Cvtermprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Cvtermprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Cvtermprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	cvtermpropUpdateCacheMut.RLock()
	cache, cached := cvtermpropUpdateCache[key]
	cvtermpropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(cvtermpropColumns, cvtermpropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update cvtermprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"cvtermprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, cvtermpropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cvtermpropType, cvtermpropMapping, append(wl, cvtermpropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update cvtermprop row")
	}

	if !cached {
		cvtermpropUpdateCacheMut.Lock()
		cvtermpropUpdateCache[key] = cache
		cvtermpropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q cvtermpropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q cvtermpropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for cvtermprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CvtermpropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CvtermpropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CvtermpropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CvtermpropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"cvtermprop\" SET %s WHERE (\"cvtermprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvtermpropPrimaryKeyColumns), len(colNames)+1, len(cvtermpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in cvtermprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Cvtermprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Cvtermprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Cvtermprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Cvtermprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no cvtermprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvtermpropColumnsWithDefault, o)

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

	cvtermpropUpsertCacheMut.RLock()
	cache, cached := cvtermpropUpsertCache[key]
	cvtermpropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			cvtermpropColumns,
			cvtermpropColumnsWithDefault,
			cvtermpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			cvtermpropColumns,
			cvtermpropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert cvtermprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(cvtermpropPrimaryKeyColumns))
			copy(conflict, cvtermpropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"cvtermprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(cvtermpropType, cvtermpropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cvtermpropType, cvtermpropMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for cvtermprop")
	}

	if !cached {
		cvtermpropUpsertCacheMut.Lock()
		cvtermpropUpsertCache[key] = cache
		cvtermpropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Cvtermprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Cvtermprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Cvtermprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Cvtermprop) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Cvtermprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Cvtermprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Cvtermprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Cvtermprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Cvtermprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Cvtermprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cvtermpropPrimaryKeyMapping)
	sql := "DELETE FROM \"cvtermprop\" WHERE \"cvtermprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from cvtermprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q cvtermpropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q cvtermpropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no cvtermpropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from cvtermprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CvtermpropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CvtermpropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Cvtermprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CvtermpropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CvtermpropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Cvtermprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(cvtermpropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"cvtermprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvtermpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvtermpropPrimaryKeyColumns), 1, len(cvtermpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from cvtermprop slice")
	}

	if len(cvtermpropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Cvtermprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Cvtermprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Cvtermprop) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Cvtermprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Cvtermprop) Reload(exec boil.Executor) error {
	ret, err := FindCvtermprop(exec, o.CvtermpropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvtermpropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvtermpropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvtermpropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty CvtermpropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvtermpropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	cvtermprops := CvtermpropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"cvtermprop\".* FROM \"cvtermprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvtermpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(cvtermpropPrimaryKeyColumns), 1, len(cvtermpropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&cvtermprops)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in CvtermpropSlice")
	}

	*o = cvtermprops

	return nil
}

// CvtermpropExists checks if the Cvtermprop row exists.
func CvtermpropExists(exec boil.Executor, cvtermpropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"cvtermprop\" where \"cvtermprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, cvtermpropID)
	}

	row := exec.QueryRow(sql, cvtermpropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if cvtermprop exists")
	}

	return exists, nil
}

// CvtermpropExistsG checks if the Cvtermprop row exists.
func CvtermpropExistsG(cvtermpropID int) (bool, error) {
	return CvtermpropExists(boil.GetDB(), cvtermpropID)
}

// CvtermpropExistsGP checks if the Cvtermprop row exists. Panics on error.
func CvtermpropExistsGP(cvtermpropID int) bool {
	e, err := CvtermpropExists(boil.GetDB(), cvtermpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CvtermpropExistsP checks if the Cvtermprop row exists. Panics on error.
func CvtermpropExistsP(exec boil.Executor, cvtermpropID int) bool {
	e, err := CvtermpropExists(exec, cvtermpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
