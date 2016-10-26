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

// Cvtermsynonym is an object representing the database table.
type Cvtermsynonym struct {
	CvtermsynonymID int      `boil:"cvtermsynonym_id" json:"cvtermsynonym_id" toml:"cvtermsynonym_id" yaml:"cvtermsynonym_id"`
	CvtermID        int      `boil:"cvterm_id" json:"cvterm_id" toml:"cvterm_id" yaml:"cvterm_id"`
	Synonym         string   `boil:"synonym" json:"synonym" toml:"synonym" yaml:"synonym"`
	TypeID          null.Int `boil:"type_id" json:"type_id,omitempty" toml:"type_id" yaml:"type_id,omitempty"`

	R *cvtermsynonymR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cvtermsynonymL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// cvtermsynonymR is where relationships are stored.
type cvtermsynonymR struct {
	Cvterm *Cvterm
	Type   *Cvterm
}

// cvtermsynonymL is where Load methods for each relationship are stored.
type cvtermsynonymL struct{}

var (
	cvtermsynonymColumns               = []string{"cvtermsynonym_id", "cvterm_id", "synonym", "type_id"}
	cvtermsynonymColumnsWithoutDefault = []string{"cvterm_id", "synonym", "type_id"}
	cvtermsynonymColumnsWithDefault    = []string{"cvtermsynonym_id"}
	cvtermsynonymPrimaryKeyColumns     = []string{"cvtermsynonym_id"}
)

type (
	// CvtermsynonymSlice is an alias for a slice of pointers to Cvtermsynonym.
	// This should generally be used opposed to []Cvtermsynonym.
	CvtermsynonymSlice []*Cvtermsynonym
	// CvtermsynonymHook is the signature for custom Cvtermsynonym hook methods
	CvtermsynonymHook func(boil.Executor, *Cvtermsynonym) error

	cvtermsynonymQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cvtermsynonymType                 = reflect.TypeOf(&Cvtermsynonym{})
	cvtermsynonymMapping              = queries.MakeStructMapping(cvtermsynonymType)
	cvtermsynonymPrimaryKeyMapping, _ = queries.BindMapping(cvtermsynonymType, cvtermsynonymMapping, cvtermsynonymPrimaryKeyColumns)
	cvtermsynonymInsertCacheMut       sync.RWMutex
	cvtermsynonymInsertCache          = make(map[string]insertCache)
	cvtermsynonymUpdateCacheMut       sync.RWMutex
	cvtermsynonymUpdateCache          = make(map[string]updateCache)
	cvtermsynonymUpsertCacheMut       sync.RWMutex
	cvtermsynonymUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var cvtermsynonymBeforeInsertHooks []CvtermsynonymHook
var cvtermsynonymBeforeUpdateHooks []CvtermsynonymHook
var cvtermsynonymBeforeDeleteHooks []CvtermsynonymHook
var cvtermsynonymBeforeUpsertHooks []CvtermsynonymHook

var cvtermsynonymAfterInsertHooks []CvtermsynonymHook
var cvtermsynonymAfterSelectHooks []CvtermsynonymHook
var cvtermsynonymAfterUpdateHooks []CvtermsynonymHook
var cvtermsynonymAfterDeleteHooks []CvtermsynonymHook
var cvtermsynonymAfterUpsertHooks []CvtermsynonymHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Cvtermsynonym) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermsynonymBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Cvtermsynonym) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermsynonymBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Cvtermsynonym) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermsynonymBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Cvtermsynonym) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermsynonymBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Cvtermsynonym) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermsynonymAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Cvtermsynonym) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermsynonymAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Cvtermsynonym) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermsynonymAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Cvtermsynonym) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermsynonymAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Cvtermsynonym) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermsynonymAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCvtermsynonymHook registers your hook function for all future operations.
func AddCvtermsynonymHook(hookPoint boil.HookPoint, cvtermsynonymHook CvtermsynonymHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cvtermsynonymBeforeInsertHooks = append(cvtermsynonymBeforeInsertHooks, cvtermsynonymHook)
	case boil.BeforeUpdateHook:
		cvtermsynonymBeforeUpdateHooks = append(cvtermsynonymBeforeUpdateHooks, cvtermsynonymHook)
	case boil.BeforeDeleteHook:
		cvtermsynonymBeforeDeleteHooks = append(cvtermsynonymBeforeDeleteHooks, cvtermsynonymHook)
	case boil.BeforeUpsertHook:
		cvtermsynonymBeforeUpsertHooks = append(cvtermsynonymBeforeUpsertHooks, cvtermsynonymHook)
	case boil.AfterInsertHook:
		cvtermsynonymAfterInsertHooks = append(cvtermsynonymAfterInsertHooks, cvtermsynonymHook)
	case boil.AfterSelectHook:
		cvtermsynonymAfterSelectHooks = append(cvtermsynonymAfterSelectHooks, cvtermsynonymHook)
	case boil.AfterUpdateHook:
		cvtermsynonymAfterUpdateHooks = append(cvtermsynonymAfterUpdateHooks, cvtermsynonymHook)
	case boil.AfterDeleteHook:
		cvtermsynonymAfterDeleteHooks = append(cvtermsynonymAfterDeleteHooks, cvtermsynonymHook)
	case boil.AfterUpsertHook:
		cvtermsynonymAfterUpsertHooks = append(cvtermsynonymAfterUpsertHooks, cvtermsynonymHook)
	}
}

// OneP returns a single cvtermsynonym record from the query, and panics on error.
func (q cvtermsynonymQuery) OneP() *Cvtermsynonym {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single cvtermsynonym record from the query.
func (q cvtermsynonymQuery) One() (*Cvtermsynonym, error) {
	o := &Cvtermsynonym{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for cvtermsynonym")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Cvtermsynonym records from the query, and panics on error.
func (q cvtermsynonymQuery) AllP() CvtermsynonymSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Cvtermsynonym records from the query.
func (q cvtermsynonymQuery) All() (CvtermsynonymSlice, error) {
	var o CvtermsynonymSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Cvtermsynonym slice")
	}

	if len(cvtermsynonymAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Cvtermsynonym records in the query, and panics on error.
func (q cvtermsynonymQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Cvtermsynonym records in the query.
func (q cvtermsynonymQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count cvtermsynonym rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q cvtermsynonymQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q cvtermsynonymQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if cvtermsynonym exists")
	}

	return count > 0, nil
}

// CvtermG pointed to by the foreign key.
func (o *Cvtermsynonym) CvtermG(mods ...qm.QueryMod) cvtermQuery {
	return o.Cvterm(boil.GetDB(), mods...)
}

// Cvterm pointed to by the foreign key.
func (o *Cvtermsynonym) Cvterm(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Cvtermsynonym) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Cvtermsynonym) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
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
func (cvtermsynonymL) LoadCvterm(e boil.Executor, singular bool, maybeCvtermsynonym interface{}) error {
	var slice []*Cvtermsynonym
	var object *Cvtermsynonym

	count := 1
	if singular {
		object = maybeCvtermsynonym.(*Cvtermsynonym)
	} else {
		slice = *maybeCvtermsynonym.(*CvtermsynonymSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermsynonymR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermsynonymR{}
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

	if len(cvtermsynonymAfterSelectHooks) != 0 {
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
func (cvtermsynonymL) LoadType(e boil.Executor, singular bool, maybeCvtermsynonym interface{}) error {
	var slice []*Cvtermsynonym
	var object *Cvtermsynonym

	count := 1
	if singular {
		object = maybeCvtermsynonym.(*Cvtermsynonym)
	} else {
		slice = *maybeCvtermsynonym.(*CvtermsynonymSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermsynonymR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermsynonymR{}
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

	if len(cvtermsynonymAfterSelectHooks) != 0 {
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
			if local.TypeID.Int == foreign.CvtermID {
				local.R.Type = foreign
				break
			}
		}
	}

	return nil
}

// SetCvterm of the cvtermsynonym to the related item.
// Sets o.R.Cvterm to related.
// Adds o to related.R.Cvtermsynonym.
func (o *Cvtermsynonym) SetCvterm(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvtermsynonym\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermsynonymPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.CvtermsynonymID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CvtermID = related.CvtermID

	if o.R == nil {
		o.R = &cvtermsynonymR{
			Cvterm: related,
		}
	} else {
		o.R.Cvterm = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			Cvtermsynonym: o,
		}
	} else {
		related.R.Cvtermsynonym = o
	}

	return nil
}

// SetType of the cvtermsynonym to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeCvtermsynonyms.
func (o *Cvtermsynonym) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvtermsynonym\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermsynonymPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.CvtermsynonymID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID.Int = related.CvtermID
	o.TypeID.Valid = true

	if o.R == nil {
		o.R = &cvtermsynonymR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeCvtermsynonyms: CvtermsynonymSlice{o},
		}
	} else {
		related.R.TypeCvtermsynonyms = append(related.R.TypeCvtermsynonyms, o)
	}

	return nil
}

// RemoveType relationship.
// Sets o.R.Type to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Cvtermsynonym) RemoveType(exec boil.Executor, related *Cvterm) error {
	var err error

	o.TypeID.Valid = false
	if err = o.Update(exec, "type_id"); err != nil {
		o.TypeID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Type = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.TypeCvtermsynonyms {
		if o.TypeID.Int != ri.TypeID.Int {
			continue
		}

		ln := len(related.R.TypeCvtermsynonyms)
		if ln > 1 && i < ln-1 {
			related.R.TypeCvtermsynonyms[i] = related.R.TypeCvtermsynonyms[ln-1]
		}
		related.R.TypeCvtermsynonyms = related.R.TypeCvtermsynonyms[:ln-1]
		break
	}
	return nil
}

// CvtermsynonymsG retrieves all records.
func CvtermsynonymsG(mods ...qm.QueryMod) cvtermsynonymQuery {
	return Cvtermsynonyms(boil.GetDB(), mods...)
}

// Cvtermsynonyms retrieves all the records using an executor.
func Cvtermsynonyms(exec boil.Executor, mods ...qm.QueryMod) cvtermsynonymQuery {
	mods = append(mods, qm.From("\"cvtermsynonym\""))
	return cvtermsynonymQuery{NewQuery(exec, mods...)}
}

// FindCvtermsynonymG retrieves a single record by ID.
func FindCvtermsynonymG(cvtermsynonymID int, selectCols ...string) (*Cvtermsynonym, error) {
	return FindCvtermsynonym(boil.GetDB(), cvtermsynonymID, selectCols...)
}

// FindCvtermsynonymGP retrieves a single record by ID, and panics on error.
func FindCvtermsynonymGP(cvtermsynonymID int, selectCols ...string) *Cvtermsynonym {
	retobj, err := FindCvtermsynonym(boil.GetDB(), cvtermsynonymID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCvtermsynonym retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCvtermsynonym(exec boil.Executor, cvtermsynonymID int, selectCols ...string) (*Cvtermsynonym, error) {
	cvtermsynonymObj := &Cvtermsynonym{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"cvtermsynonym\" where \"cvtermsynonym_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, cvtermsynonymID)

	err := q.Bind(cvtermsynonymObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from cvtermsynonym")
	}

	return cvtermsynonymObj, nil
}

// FindCvtermsynonymP retrieves a single record by ID with an executor, and panics on error.
func FindCvtermsynonymP(exec boil.Executor, cvtermsynonymID int, selectCols ...string) *Cvtermsynonym {
	retobj, err := FindCvtermsynonym(exec, cvtermsynonymID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Cvtermsynonym) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Cvtermsynonym) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Cvtermsynonym) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Cvtermsynonym) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no cvtermsynonym provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvtermsynonymColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	cvtermsynonymInsertCacheMut.RLock()
	cache, cached := cvtermsynonymInsertCache[key]
	cvtermsynonymInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			cvtermsynonymColumns,
			cvtermsynonymColumnsWithDefault,
			cvtermsynonymColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(cvtermsynonymType, cvtermsynonymMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cvtermsynonymType, cvtermsynonymMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"cvtermsynonym\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into cvtermsynonym")
	}

	if !cached {
		cvtermsynonymInsertCacheMut.Lock()
		cvtermsynonymInsertCache[key] = cache
		cvtermsynonymInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Cvtermsynonym record. See Update for
// whitelist behavior description.
func (o *Cvtermsynonym) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Cvtermsynonym record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Cvtermsynonym) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Cvtermsynonym, and panics on error.
// See Update for whitelist behavior description.
func (o *Cvtermsynonym) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Cvtermsynonym.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Cvtermsynonym) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	cvtermsynonymUpdateCacheMut.RLock()
	cache, cached := cvtermsynonymUpdateCache[key]
	cvtermsynonymUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(cvtermsynonymColumns, cvtermsynonymPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update cvtermsynonym, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"cvtermsynonym\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, cvtermsynonymPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cvtermsynonymType, cvtermsynonymMapping, append(wl, cvtermsynonymPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update cvtermsynonym row")
	}

	if !cached {
		cvtermsynonymUpdateCacheMut.Lock()
		cvtermsynonymUpdateCache[key] = cache
		cvtermsynonymUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q cvtermsynonymQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q cvtermsynonymQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for cvtermsynonym")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CvtermsynonymSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CvtermsynonymSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CvtermsynonymSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CvtermsynonymSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermsynonymPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"cvtermsynonym\" SET %s WHERE (\"cvtermsynonym_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvtermsynonymPrimaryKeyColumns), len(colNames)+1, len(cvtermsynonymPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in cvtermsynonym slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Cvtermsynonym) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Cvtermsynonym) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Cvtermsynonym) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Cvtermsynonym) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no cvtermsynonym provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvtermsynonymColumnsWithDefault, o)

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

	cvtermsynonymUpsertCacheMut.RLock()
	cache, cached := cvtermsynonymUpsertCache[key]
	cvtermsynonymUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			cvtermsynonymColumns,
			cvtermsynonymColumnsWithDefault,
			cvtermsynonymColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			cvtermsynonymColumns,
			cvtermsynonymPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert cvtermsynonym, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(cvtermsynonymPrimaryKeyColumns))
			copy(conflict, cvtermsynonymPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"cvtermsynonym\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(cvtermsynonymType, cvtermsynonymMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cvtermsynonymType, cvtermsynonymMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for cvtermsynonym")
	}

	if !cached {
		cvtermsynonymUpsertCacheMut.Lock()
		cvtermsynonymUpsertCache[key] = cache
		cvtermsynonymUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Cvtermsynonym record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Cvtermsynonym) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Cvtermsynonym record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Cvtermsynonym) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Cvtermsynonym provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Cvtermsynonym record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Cvtermsynonym) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Cvtermsynonym record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Cvtermsynonym) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Cvtermsynonym provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cvtermsynonymPrimaryKeyMapping)
	sql := "DELETE FROM \"cvtermsynonym\" WHERE \"cvtermsynonym_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from cvtermsynonym")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q cvtermsynonymQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q cvtermsynonymQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no cvtermsynonymQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from cvtermsynonym")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CvtermsynonymSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CvtermsynonymSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Cvtermsynonym slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CvtermsynonymSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CvtermsynonymSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Cvtermsynonym slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(cvtermsynonymBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermsynonymPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"cvtermsynonym\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvtermsynonymPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvtermsynonymPrimaryKeyColumns), 1, len(cvtermsynonymPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from cvtermsynonym slice")
	}

	if len(cvtermsynonymAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Cvtermsynonym) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Cvtermsynonym) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Cvtermsynonym) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Cvtermsynonym provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Cvtermsynonym) Reload(exec boil.Executor) error {
	ret, err := FindCvtermsynonym(exec, o.CvtermsynonymID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvtermsynonymSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvtermsynonymSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvtermsynonymSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty CvtermsynonymSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvtermsynonymSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	cvtermsynonyms := CvtermsynonymSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermsynonymPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"cvtermsynonym\".* FROM \"cvtermsynonym\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvtermsynonymPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(cvtermsynonymPrimaryKeyColumns), 1, len(cvtermsynonymPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&cvtermsynonyms)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in CvtermsynonymSlice")
	}

	*o = cvtermsynonyms

	return nil
}

// CvtermsynonymExists checks if the Cvtermsynonym row exists.
func CvtermsynonymExists(exec boil.Executor, cvtermsynonymID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"cvtermsynonym\" where \"cvtermsynonym_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, cvtermsynonymID)
	}

	row := exec.QueryRow(sql, cvtermsynonymID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if cvtermsynonym exists")
	}

	return exists, nil
}

// CvtermsynonymExistsG checks if the Cvtermsynonym row exists.
func CvtermsynonymExistsG(cvtermsynonymID int) (bool, error) {
	return CvtermsynonymExists(boil.GetDB(), cvtermsynonymID)
}

// CvtermsynonymExistsGP checks if the Cvtermsynonym row exists. Panics on error.
func CvtermsynonymExistsGP(cvtermsynonymID int) bool {
	e, err := CvtermsynonymExists(boil.GetDB(), cvtermsynonymID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CvtermsynonymExistsP checks if the Cvtermsynonym row exists. Panics on error.
func CvtermsynonymExistsP(exec boil.Executor, cvtermsynonymID int) bool {
	e, err := CvtermsynonymExists(exec, cvtermsynonymID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
