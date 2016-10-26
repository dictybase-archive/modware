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

// Synonym is an object representing the database table.
type Synonym struct {
	SynonymID   int    `boil:"synonym_id" json:"synonym_id" toml:"synonym_id" yaml:"synonym_id"`
	Name        string `boil:"name" json:"name" toml:"name" yaml:"name"`
	TypeID      int    `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	SynonymSGML string `boil:"synonym_sgml" json:"synonym_sgml" toml:"synonym_sgml" yaml:"synonym_sgml"`

	R *synonymR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L synonymL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// synonymR is where relationships are stored.
type synonymR struct {
	Type           *Cvterm
	FeatureSynonym *FeatureSynonym
}

// synonymL is where Load methods for each relationship are stored.
type synonymL struct{}

var (
	synonymColumns               = []string{"synonym_id", "name", "type_id", "synonym_sgml"}
	synonymColumnsWithoutDefault = []string{"name", "type_id", "synonym_sgml"}
	synonymColumnsWithDefault    = []string{"synonym_id"}
	synonymPrimaryKeyColumns     = []string{"synonym_id"}
)

type (
	// SynonymSlice is an alias for a slice of pointers to Synonym.
	// This should generally be used opposed to []Synonym.
	SynonymSlice []*Synonym
	// SynonymHook is the signature for custom Synonym hook methods
	SynonymHook func(boil.Executor, *Synonym) error

	synonymQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	synonymType                 = reflect.TypeOf(&Synonym{})
	synonymMapping              = queries.MakeStructMapping(synonymType)
	synonymPrimaryKeyMapping, _ = queries.BindMapping(synonymType, synonymMapping, synonymPrimaryKeyColumns)
	synonymInsertCacheMut       sync.RWMutex
	synonymInsertCache          = make(map[string]insertCache)
	synonymUpdateCacheMut       sync.RWMutex
	synonymUpdateCache          = make(map[string]updateCache)
	synonymUpsertCacheMut       sync.RWMutex
	synonymUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var synonymBeforeInsertHooks []SynonymHook
var synonymBeforeUpdateHooks []SynonymHook
var synonymBeforeDeleteHooks []SynonymHook
var synonymBeforeUpsertHooks []SynonymHook

var synonymAfterInsertHooks []SynonymHook
var synonymAfterSelectHooks []SynonymHook
var synonymAfterUpdateHooks []SynonymHook
var synonymAfterDeleteHooks []SynonymHook
var synonymAfterUpsertHooks []SynonymHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Synonym) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range synonymBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Synonym) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range synonymBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Synonym) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range synonymBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Synonym) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range synonymBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Synonym) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range synonymAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Synonym) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range synonymAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Synonym) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range synonymAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Synonym) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range synonymAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Synonym) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range synonymAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSynonymHook registers your hook function for all future operations.
func AddSynonymHook(hookPoint boil.HookPoint, synonymHook SynonymHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		synonymBeforeInsertHooks = append(synonymBeforeInsertHooks, synonymHook)
	case boil.BeforeUpdateHook:
		synonymBeforeUpdateHooks = append(synonymBeforeUpdateHooks, synonymHook)
	case boil.BeforeDeleteHook:
		synonymBeforeDeleteHooks = append(synonymBeforeDeleteHooks, synonymHook)
	case boil.BeforeUpsertHook:
		synonymBeforeUpsertHooks = append(synonymBeforeUpsertHooks, synonymHook)
	case boil.AfterInsertHook:
		synonymAfterInsertHooks = append(synonymAfterInsertHooks, synonymHook)
	case boil.AfterSelectHook:
		synonymAfterSelectHooks = append(synonymAfterSelectHooks, synonymHook)
	case boil.AfterUpdateHook:
		synonymAfterUpdateHooks = append(synonymAfterUpdateHooks, synonymHook)
	case boil.AfterDeleteHook:
		synonymAfterDeleteHooks = append(synonymAfterDeleteHooks, synonymHook)
	case boil.AfterUpsertHook:
		synonymAfterUpsertHooks = append(synonymAfterUpsertHooks, synonymHook)
	}
}

// OneP returns a single synonym record from the query, and panics on error.
func (q synonymQuery) OneP() *Synonym {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single synonym record from the query.
func (q synonymQuery) One() (*Synonym, error) {
	o := &Synonym{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for synonym")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Synonym records from the query, and panics on error.
func (q synonymQuery) AllP() SynonymSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Synonym records from the query.
func (q synonymQuery) All() (SynonymSlice, error) {
	var o SynonymSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Synonym slice")
	}

	if len(synonymAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Synonym records in the query, and panics on error.
func (q synonymQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Synonym records in the query.
func (q synonymQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count synonym rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q synonymQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q synonymQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if synonym exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *Synonym) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Synonym) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// FeatureSynonymG pointed to by the foreign key.
func (o *Synonym) FeatureSynonymG(mods ...qm.QueryMod) featureSynonymQuery {
	return o.FeatureSynonym(boil.GetDB(), mods...)
}

// FeatureSynonym pointed to by the foreign key.
func (o *Synonym) FeatureSynonym(exec boil.Executor, mods ...qm.QueryMod) featureSynonymQuery {
	queryMods := []qm.QueryMod{
		qm.Where("synonym_id=$1", o.SynonymID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureSynonyms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_synonym\"")

	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (synonymL) LoadType(e boil.Executor, singular bool, maybeSynonym interface{}) error {
	var slice []*Synonym
	var object *Synonym

	count := 1
	if singular {
		object = maybeSynonym.(*Synonym)
	} else {
		slice = *maybeSynonym.(*SynonymSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &synonymR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &synonymR{}
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

	if len(synonymAfterSelectHooks) != 0 {
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

// LoadFeatureSynonym allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (synonymL) LoadFeatureSynonym(e boil.Executor, singular bool, maybeSynonym interface{}) error {
	var slice []*Synonym
	var object *Synonym

	count := 1
	if singular {
		object = maybeSynonym.(*Synonym)
	} else {
		slice = *maybeSynonym.(*SynonymSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &synonymR{}
		args[0] = object.SynonymID
	} else {
		for i, obj := range slice {
			obj.R = &synonymR{}
			args[i] = obj.SynonymID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_synonym\" where \"synonym_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureSynonym")
	}
	defer results.Close()

	var resultSlice []*FeatureSynonym
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureSynonym")
	}

	if len(synonymAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureSynonym = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.SynonymID == foreign.SynonymID {
				local.R.FeatureSynonym = foreign
				break
			}
		}
	}

	return nil
}

// SetType of the synonym to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeSynonym.
func (o *Synonym) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"synonym\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, synonymPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.SynonymID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &synonymR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeSynonym: o,
		}
	} else {
		related.R.TypeSynonym = o
	}

	return nil
}

// SetFeatureSynonym of the synonym to the related item.
// Sets o.R.FeatureSynonym to related.
// Adds o to related.R.Synonym.
func (o *Synonym) SetFeatureSynonym(exec boil.Executor, insert bool, related *FeatureSynonym) error {
	var err error

	if insert {
		related.SynonymID = o.SynonymID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_synonym\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"synonym_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureSynonymPrimaryKeyColumns),
		)
		values := []interface{}{o.SynonymID, related.FeatureSynonymID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.SynonymID = o.SynonymID

	}

	if o.R == nil {
		o.R = &synonymR{
			FeatureSynonym: related,
		}
	} else {
		o.R.FeatureSynonym = related
	}

	if related.R == nil {
		related.R = &featureSynonymR{
			Synonym: o,
		}
	} else {
		related.R.Synonym = o
	}
	return nil
}

// SynonymsG retrieves all records.
func SynonymsG(mods ...qm.QueryMod) synonymQuery {
	return Synonyms(boil.GetDB(), mods...)
}

// Synonyms retrieves all the records using an executor.
func Synonyms(exec boil.Executor, mods ...qm.QueryMod) synonymQuery {
	mods = append(mods, qm.From("\"synonym\""))
	return synonymQuery{NewQuery(exec, mods...)}
}

// FindSynonymG retrieves a single record by ID.
func FindSynonymG(synonymID int, selectCols ...string) (*Synonym, error) {
	return FindSynonym(boil.GetDB(), synonymID, selectCols...)
}

// FindSynonymGP retrieves a single record by ID, and panics on error.
func FindSynonymGP(synonymID int, selectCols ...string) *Synonym {
	retobj, err := FindSynonym(boil.GetDB(), synonymID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindSynonym retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSynonym(exec boil.Executor, synonymID int, selectCols ...string) (*Synonym, error) {
	synonymObj := &Synonym{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"synonym\" where \"synonym_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, synonymID)

	err := q.Bind(synonymObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from synonym")
	}

	return synonymObj, nil
}

// FindSynonymP retrieves a single record by ID with an executor, and panics on error.
func FindSynonymP(exec boil.Executor, synonymID int, selectCols ...string) *Synonym {
	retobj, err := FindSynonym(exec, synonymID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Synonym) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Synonym) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Synonym) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Synonym) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no synonym provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(synonymColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	synonymInsertCacheMut.RLock()
	cache, cached := synonymInsertCache[key]
	synonymInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			synonymColumns,
			synonymColumnsWithDefault,
			synonymColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(synonymType, synonymMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(synonymType, synonymMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"synonym\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into synonym")
	}

	if !cached {
		synonymInsertCacheMut.Lock()
		synonymInsertCache[key] = cache
		synonymInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Synonym record. See Update for
// whitelist behavior description.
func (o *Synonym) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Synonym record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Synonym) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Synonym, and panics on error.
// See Update for whitelist behavior description.
func (o *Synonym) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Synonym.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Synonym) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	synonymUpdateCacheMut.RLock()
	cache, cached := synonymUpdateCache[key]
	synonymUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(synonymColumns, synonymPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update synonym, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"synonym\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, synonymPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(synonymType, synonymMapping, append(wl, synonymPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update synonym row")
	}

	if !cached {
		synonymUpdateCacheMut.Lock()
		synonymUpdateCache[key] = cache
		synonymUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q synonymQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q synonymQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for synonym")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o SynonymSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o SynonymSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o SynonymSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SynonymSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), synonymPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"synonym\" SET %s WHERE (\"synonym_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(synonymPrimaryKeyColumns), len(colNames)+1, len(synonymPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in synonym slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Synonym) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Synonym) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Synonym) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Synonym) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no synonym provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(synonymColumnsWithDefault, o)

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

	synonymUpsertCacheMut.RLock()
	cache, cached := synonymUpsertCache[key]
	synonymUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			synonymColumns,
			synonymColumnsWithDefault,
			synonymColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			synonymColumns,
			synonymPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert synonym, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(synonymPrimaryKeyColumns))
			copy(conflict, synonymPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"synonym\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(synonymType, synonymMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(synonymType, synonymMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for synonym")
	}

	if !cached {
		synonymUpsertCacheMut.Lock()
		synonymUpsertCache[key] = cache
		synonymUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Synonym record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Synonym) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Synonym record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Synonym) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Synonym provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Synonym record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Synonym) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Synonym record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Synonym) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Synonym provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), synonymPrimaryKeyMapping)
	sql := "DELETE FROM \"synonym\" WHERE \"synonym_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from synonym")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q synonymQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q synonymQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no synonymQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from synonym")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o SynonymSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o SynonymSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Synonym slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o SynonymSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SynonymSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Synonym slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(synonymBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), synonymPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"synonym\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, synonymPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(synonymPrimaryKeyColumns), 1, len(synonymPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from synonym slice")
	}

	if len(synonymAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Synonym) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Synonym) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Synonym) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Synonym provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Synonym) Reload(exec boil.Executor) error {
	ret, err := FindSynonym(exec, o.SynonymID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *SynonymSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *SynonymSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SynonymSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty SynonymSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SynonymSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	synonyms := SynonymSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), synonymPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"synonym\".* FROM \"synonym\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, synonymPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(synonymPrimaryKeyColumns), 1, len(synonymPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&synonyms)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in SynonymSlice")
	}

	*o = synonyms

	return nil
}

// SynonymExists checks if the Synonym row exists.
func SynonymExists(exec boil.Executor, synonymID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"synonym\" where \"synonym_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, synonymID)
	}

	row := exec.QueryRow(sql, synonymID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if synonym exists")
	}

	return exists, nil
}

// SynonymExistsG checks if the Synonym row exists.
func SynonymExistsG(synonymID int) (bool, error) {
	return SynonymExists(boil.GetDB(), synonymID)
}

// SynonymExistsGP checks if the Synonym row exists. Panics on error.
func SynonymExistsGP(synonymID int) bool {
	e, err := SynonymExists(boil.GetDB(), synonymID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// SynonymExistsP checks if the Synonym row exists. Panics on error.
func SynonymExistsP(exec boil.Executor, synonymID int) bool {
	e, err := SynonymExists(exec, synonymID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
