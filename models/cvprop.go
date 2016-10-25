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

// Cvprop is an object representing the database table.
type Cvprop struct {
	CvpropID int         `boil:"cvprop_id" json:"cvprop_id" toml:"cvprop_id" yaml:"cvprop_id"`
	CVID     int         `boil:"cv_id" json:"cv_id" toml:"cv_id" yaml:"cv_id"`
	TypeID   int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value    null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank     int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *cvpropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cvpropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// cvpropR is where relationships are stored.
type cvpropR struct {
	CV   *CV
	Type *Cvterm
}

// cvpropL is where Load methods for each relationship are stored.
type cvpropL struct{}

var (
	cvpropColumns               = []string{"cvprop_id", "cv_id", "type_id", "value", "rank"}
	cvpropColumnsWithoutDefault = []string{"cv_id", "type_id", "value"}
	cvpropColumnsWithDefault    = []string{"cvprop_id", "rank"}
	cvpropPrimaryKeyColumns     = []string{"cvprop_id"}
)

type (
	// CvpropSlice is an alias for a slice of pointers to Cvprop.
	// This should generally be used opposed to []Cvprop.
	CvpropSlice []*Cvprop
	// CvpropHook is the signature for custom Cvprop hook methods
	CvpropHook func(boil.Executor, *Cvprop) error

	cvpropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cvpropType                 = reflect.TypeOf(&Cvprop{})
	cvpropMapping              = queries.MakeStructMapping(cvpropType)
	cvpropPrimaryKeyMapping, _ = queries.BindMapping(cvpropType, cvpropMapping, cvpropPrimaryKeyColumns)
	cvpropInsertCacheMut       sync.RWMutex
	cvpropInsertCache          = make(map[string]insertCache)
	cvpropUpdateCacheMut       sync.RWMutex
	cvpropUpdateCache          = make(map[string]updateCache)
	cvpropUpsertCacheMut       sync.RWMutex
	cvpropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var cvpropBeforeInsertHooks []CvpropHook
var cvpropBeforeUpdateHooks []CvpropHook
var cvpropBeforeDeleteHooks []CvpropHook
var cvpropBeforeUpsertHooks []CvpropHook

var cvpropAfterInsertHooks []CvpropHook
var cvpropAfterSelectHooks []CvpropHook
var cvpropAfterUpdateHooks []CvpropHook
var cvpropAfterDeleteHooks []CvpropHook
var cvpropAfterUpsertHooks []CvpropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Cvprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvpropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Cvprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvpropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Cvprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvpropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Cvprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvpropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Cvprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvpropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Cvprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range cvpropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Cvprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvpropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Cvprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvpropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Cvprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvpropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCvpropHook registers your hook function for all future operations.
func AddCvpropHook(hookPoint boil.HookPoint, cvpropHook CvpropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cvpropBeforeInsertHooks = append(cvpropBeforeInsertHooks, cvpropHook)
	case boil.BeforeUpdateHook:
		cvpropBeforeUpdateHooks = append(cvpropBeforeUpdateHooks, cvpropHook)
	case boil.BeforeDeleteHook:
		cvpropBeforeDeleteHooks = append(cvpropBeforeDeleteHooks, cvpropHook)
	case boil.BeforeUpsertHook:
		cvpropBeforeUpsertHooks = append(cvpropBeforeUpsertHooks, cvpropHook)
	case boil.AfterInsertHook:
		cvpropAfterInsertHooks = append(cvpropAfterInsertHooks, cvpropHook)
	case boil.AfterSelectHook:
		cvpropAfterSelectHooks = append(cvpropAfterSelectHooks, cvpropHook)
	case boil.AfterUpdateHook:
		cvpropAfterUpdateHooks = append(cvpropAfterUpdateHooks, cvpropHook)
	case boil.AfterDeleteHook:
		cvpropAfterDeleteHooks = append(cvpropAfterDeleteHooks, cvpropHook)
	case boil.AfterUpsertHook:
		cvpropAfterUpsertHooks = append(cvpropAfterUpsertHooks, cvpropHook)
	}
}

// OneP returns a single cvprop record from the query, and panics on error.
func (q cvpropQuery) OneP() *Cvprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single cvprop record from the query.
func (q cvpropQuery) One() (*Cvprop, error) {
	o := &Cvprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cvprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Cvprop records from the query, and panics on error.
func (q cvpropQuery) AllP() CvpropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Cvprop records from the query.
func (q cvpropQuery) All() (CvpropSlice, error) {
	var o CvpropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Cvprop slice")
	}

	if len(cvpropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Cvprop records in the query, and panics on error.
func (q cvpropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Cvprop records in the query.
func (q cvpropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cvprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q cvpropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q cvpropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cvprop exists")
	}

	return count > 0, nil
}

// CVG pointed to by the foreign key.
func (o *Cvprop) CVG(mods ...qm.QueryMod) cvQuery {
	return o.CV(boil.GetDB(), mods...)
}

// CV pointed to by the foreign key.
func (o *Cvprop) CV(exec boil.Executor, mods ...qm.QueryMod) cvQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cv_id=$1", o.CVID),
	}

	queryMods = append(queryMods, mods...)

	query := CVS(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cv\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Cvprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Cvprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadCV allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvpropL) LoadCV(e boil.Executor, singular bool, maybeCvprop interface{}) error {
	var slice []*Cvprop
	var object *Cvprop

	count := 1
	if singular {
		object = maybeCvprop.(*Cvprop)
	} else {
		slice = *maybeCvprop.(*CvpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvpropR{}
		args[0] = object.CVID
	} else {
		for i, obj := range slice {
			obj.R = &cvpropR{}
			args[i] = obj.CVID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cv\" where \"cv_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CV")
	}
	defer results.Close()

	var resultSlice []*CV
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CV")
	}

	if len(cvpropAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.CV = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CVID == foreign.CVID {
				local.R.CV = foreign
				break
			}
		}
	}

	return nil
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvpropL) LoadType(e boil.Executor, singular bool, maybeCvprop interface{}) error {
	var slice []*Cvprop
	var object *Cvprop

	count := 1
	if singular {
		object = maybeCvprop.(*Cvprop)
	} else {
		slice = *maybeCvprop.(*CvpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvpropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &cvpropR{}
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

	if len(cvpropAfterSelectHooks) != 0 {
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

// SetCV of the cvprop to the related item.
// Sets o.R.CV to related.
// Adds o to related.R.Cvprop.
func (o *Cvprop) SetCV(exec boil.Executor, insert bool, related *CV) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"cv_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvpropPrimaryKeyColumns),
	)
	values := []interface{}{related.CVID, o.CvpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CVID = related.CVID

	if o.R == nil {
		o.R = &cvpropR{
			CV: related,
		}
	} else {
		o.R.CV = related
	}

	if related.R == nil {
		related.R = &cvR{
			Cvprop: o,
		}
	} else {
		related.R.Cvprop = o
	}

	return nil
}

// SetType of the cvprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeCvprop.
func (o *Cvprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvpropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.CvpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &cvpropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeCvprop: o,
		}
	} else {
		related.R.TypeCvprop = o
	}

	return nil
}

// CvpropsG retrieves all records.
func CvpropsG(mods ...qm.QueryMod) cvpropQuery {
	return Cvprops(boil.GetDB(), mods...)
}

// Cvprops retrieves all the records using an executor.
func Cvprops(exec boil.Executor, mods ...qm.QueryMod) cvpropQuery {
	mods = append(mods, qm.From("\"cvprop\""))
	return cvpropQuery{NewQuery(exec, mods...)}
}

// FindCvpropG retrieves a single record by ID.
func FindCvpropG(cvpropID int, selectCols ...string) (*Cvprop, error) {
	return FindCvprop(boil.GetDB(), cvpropID, selectCols...)
}

// FindCvpropGP retrieves a single record by ID, and panics on error.
func FindCvpropGP(cvpropID int, selectCols ...string) *Cvprop {
	retobj, err := FindCvprop(boil.GetDB(), cvpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCvprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCvprop(exec boil.Executor, cvpropID int, selectCols ...string) (*Cvprop, error) {
	cvpropObj := &Cvprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"cvprop\" where \"cvprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, cvpropID)

	err := q.Bind(cvpropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cvprop")
	}

	return cvpropObj, nil
}

// FindCvpropP retrieves a single record by ID with an executor, and panics on error.
func FindCvpropP(exec boil.Executor, cvpropID int, selectCols ...string) *Cvprop {
	retobj, err := FindCvprop(exec, cvpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Cvprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Cvprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Cvprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Cvprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no cvprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvpropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	cvpropInsertCacheMut.RLock()
	cache, cached := cvpropInsertCache[key]
	cvpropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			cvpropColumns,
			cvpropColumnsWithDefault,
			cvpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(cvpropType, cvpropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cvpropType, cvpropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"cvprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into cvprop")
	}

	if !cached {
		cvpropInsertCacheMut.Lock()
		cvpropInsertCache[key] = cache
		cvpropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Cvprop record. See Update for
// whitelist behavior description.
func (o *Cvprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Cvprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Cvprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Cvprop, and panics on error.
// See Update for whitelist behavior description.
func (o *Cvprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Cvprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Cvprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	cvpropUpdateCacheMut.RLock()
	cache, cached := cvpropUpdateCache[key]
	cvpropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(cvpropColumns, cvpropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update cvprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"cvprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, cvpropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cvpropType, cvpropMapping, append(wl, cvpropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update cvprop row")
	}

	if !cached {
		cvpropUpdateCacheMut.Lock()
		cvpropUpdateCache[key] = cache
		cvpropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q cvpropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q cvpropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for cvprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CvpropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CvpropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CvpropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CvpropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"cvprop\" SET %s WHERE (\"cvprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvpropPrimaryKeyColumns), len(colNames)+1, len(cvpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in cvprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Cvprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Cvprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Cvprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Cvprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no cvprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvpropColumnsWithDefault, o)

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

	cvpropUpsertCacheMut.RLock()
	cache, cached := cvpropUpsertCache[key]
	cvpropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			cvpropColumns,
			cvpropColumnsWithDefault,
			cvpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			cvpropColumns,
			cvpropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert cvprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(cvpropPrimaryKeyColumns))
			copy(conflict, cvpropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"cvprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(cvpropType, cvpropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cvpropType, cvpropMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cvprop")
	}

	if !cached {
		cvpropUpsertCacheMut.Lock()
		cvpropUpsertCache[key] = cache
		cvpropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Cvprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Cvprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Cvprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Cvprop) DeleteG() error {
	if o == nil {
		return errors.New("models: no Cvprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Cvprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Cvprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Cvprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Cvprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Cvprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cvpropPrimaryKeyMapping)
	sql := "DELETE FROM \"cvprop\" WHERE \"cvprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from cvprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q cvpropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q cvpropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no cvpropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from cvprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CvpropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CvpropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Cvprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CvpropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CvpropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Cvprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(cvpropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"cvprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvpropPrimaryKeyColumns), 1, len(cvpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from cvprop slice")
	}

	if len(cvpropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Cvprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Cvprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Cvprop) ReloadG() error {
	if o == nil {
		return errors.New("models: no Cvprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Cvprop) Reload(exec boil.Executor) error {
	ret, err := FindCvprop(exec, o.CvpropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvpropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvpropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvpropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty CvpropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvpropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	cvprops := CvpropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"cvprop\".* FROM \"cvprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(cvpropPrimaryKeyColumns), 1, len(cvpropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&cvprops)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CvpropSlice")
	}

	*o = cvprops

	return nil
}

// CvpropExists checks if the Cvprop row exists.
func CvpropExists(exec boil.Executor, cvpropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"cvprop\" where \"cvprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, cvpropID)
	}

	row := exec.QueryRow(sql, cvpropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cvprop exists")
	}

	return exists, nil
}

// CvpropExistsG checks if the Cvprop row exists.
func CvpropExistsG(cvpropID int) (bool, error) {
	return CvpropExists(boil.GetDB(), cvpropID)
}

// CvpropExistsGP checks if the Cvprop row exists. Panics on error.
func CvpropExistsGP(cvpropID int) bool {
	e, err := CvpropExists(boil.GetDB(), cvpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CvpropExistsP checks if the Cvprop row exists. Panics on error.
func CvpropExistsP(exec boil.Executor, cvpropID int) bool {
	e, err := CvpropExists(exec, cvpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
