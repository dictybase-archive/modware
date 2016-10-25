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

// Organismprop is an object representing the database table.
type Organismprop struct {
	OrganismpropID int         `boil:"organismprop_id" json:"organismprop_id" toml:"organismprop_id" yaml:"organismprop_id"`
	OrganismID     int         `boil:"organism_id" json:"organism_id" toml:"organism_id" yaml:"organism_id"`
	TypeID         int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Value          null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Rank           int         `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *organismpropR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L organismpropL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// organismpropR is where relationships are stored.
type organismpropR struct {
	Organism *Organism
	Type     *Cvterm
}

// organismpropL is where Load methods for each relationship are stored.
type organismpropL struct{}

var (
	organismpropColumns               = []string{"organismprop_id", "organism_id", "type_id", "value", "rank"}
	organismpropColumnsWithoutDefault = []string{"organism_id", "type_id", "value"}
	organismpropColumnsWithDefault    = []string{"organismprop_id", "rank"}
	organismpropPrimaryKeyColumns     = []string{"organismprop_id"}
)

type (
	// OrganismpropSlice is an alias for a slice of pointers to Organismprop.
	// This should generally be used opposed to []Organismprop.
	OrganismpropSlice []*Organismprop
	// OrganismpropHook is the signature for custom Organismprop hook methods
	OrganismpropHook func(boil.Executor, *Organismprop) error

	organismpropQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	organismpropType                 = reflect.TypeOf(&Organismprop{})
	organismpropMapping              = queries.MakeStructMapping(organismpropType)
	organismpropPrimaryKeyMapping, _ = queries.BindMapping(organismpropType, organismpropMapping, organismpropPrimaryKeyColumns)
	organismpropInsertCacheMut       sync.RWMutex
	organismpropInsertCache          = make(map[string]insertCache)
	organismpropUpdateCacheMut       sync.RWMutex
	organismpropUpdateCache          = make(map[string]updateCache)
	organismpropUpsertCacheMut       sync.RWMutex
	organismpropUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var organismpropBeforeInsertHooks []OrganismpropHook
var organismpropBeforeUpdateHooks []OrganismpropHook
var organismpropBeforeDeleteHooks []OrganismpropHook
var organismpropBeforeUpsertHooks []OrganismpropHook

var organismpropAfterInsertHooks []OrganismpropHook
var organismpropAfterSelectHooks []OrganismpropHook
var organismpropAfterUpdateHooks []OrganismpropHook
var organismpropAfterDeleteHooks []OrganismpropHook
var organismpropAfterUpsertHooks []OrganismpropHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Organismprop) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organismpropBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Organismprop) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range organismpropBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Organismprop) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range organismpropBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Organismprop) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organismpropBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Organismprop) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organismpropAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Organismprop) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range organismpropAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Organismprop) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range organismpropAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Organismprop) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range organismpropAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Organismprop) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range organismpropAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrganismpropHook registers your hook function for all future operations.
func AddOrganismpropHook(hookPoint boil.HookPoint, organismpropHook OrganismpropHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		organismpropBeforeInsertHooks = append(organismpropBeforeInsertHooks, organismpropHook)
	case boil.BeforeUpdateHook:
		organismpropBeforeUpdateHooks = append(organismpropBeforeUpdateHooks, organismpropHook)
	case boil.BeforeDeleteHook:
		organismpropBeforeDeleteHooks = append(organismpropBeforeDeleteHooks, organismpropHook)
	case boil.BeforeUpsertHook:
		organismpropBeforeUpsertHooks = append(organismpropBeforeUpsertHooks, organismpropHook)
	case boil.AfterInsertHook:
		organismpropAfterInsertHooks = append(organismpropAfterInsertHooks, organismpropHook)
	case boil.AfterSelectHook:
		organismpropAfterSelectHooks = append(organismpropAfterSelectHooks, organismpropHook)
	case boil.AfterUpdateHook:
		organismpropAfterUpdateHooks = append(organismpropAfterUpdateHooks, organismpropHook)
	case boil.AfterDeleteHook:
		organismpropAfterDeleteHooks = append(organismpropAfterDeleteHooks, organismpropHook)
	case boil.AfterUpsertHook:
		organismpropAfterUpsertHooks = append(organismpropAfterUpsertHooks, organismpropHook)
	}
}

// OneP returns a single organismprop record from the query, and panics on error.
func (q organismpropQuery) OneP() *Organismprop {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single organismprop record from the query.
func (q organismpropQuery) One() (*Organismprop, error) {
	o := &Organismprop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for organismprop")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Organismprop records from the query, and panics on error.
func (q organismpropQuery) AllP() OrganismpropSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Organismprop records from the query.
func (q organismpropQuery) All() (OrganismpropSlice, error) {
	var o OrganismpropSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Organismprop slice")
	}

	if len(organismpropAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Organismprop records in the query, and panics on error.
func (q organismpropQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Organismprop records in the query.
func (q organismpropQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count organismprop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q organismpropQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q organismpropQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if organismprop exists")
	}

	return count > 0, nil
}

// OrganismG pointed to by the foreign key.
func (o *Organismprop) OrganismG(mods ...qm.QueryMod) organismQuery {
	return o.Organism(boil.GetDB(), mods...)
}

// Organism pointed to by the foreign key.
func (o *Organismprop) Organism(exec boil.Executor, mods ...qm.QueryMod) organismQuery {
	queryMods := []qm.QueryMod{
		qm.Where("organism_id=$1", o.OrganismID),
	}

	queryMods = append(queryMods, mods...)

	query := Organisms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"organism\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Organismprop) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Organismprop) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadOrganism allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (organismpropL) LoadOrganism(e boil.Executor, singular bool, maybeOrganismprop interface{}) error {
	var slice []*Organismprop
	var object *Organismprop

	count := 1
	if singular {
		object = maybeOrganismprop.(*Organismprop)
	} else {
		slice = *maybeOrganismprop.(*OrganismpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &organismpropR{}
		args[0] = object.OrganismID
	} else {
		for i, obj := range slice {
			obj.R = &organismpropR{}
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

	if len(organismpropAfterSelectHooks) != 0 {
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

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (organismpropL) LoadType(e boil.Executor, singular bool, maybeOrganismprop interface{}) error {
	var slice []*Organismprop
	var object *Organismprop

	count := 1
	if singular {
		object = maybeOrganismprop.(*Organismprop)
	} else {
		slice = *maybeOrganismprop.(*OrganismpropSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &organismpropR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &organismpropR{}
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

	if len(organismpropAfterSelectHooks) != 0 {
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

// SetOrganism of the organismprop to the related item.
// Sets o.R.Organism to related.
// Adds o to related.R.Organismprop.
func (o *Organismprop) SetOrganism(exec boil.Executor, insert bool, related *Organism) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"organismprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"organism_id"}),
		strmangle.WhereClause("\"", "\"", 2, organismpropPrimaryKeyColumns),
	)
	values := []interface{}{related.OrganismID, o.OrganismpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrganismID = related.OrganismID

	if o.R == nil {
		o.R = &organismpropR{
			Organism: related,
		}
	} else {
		o.R.Organism = related
	}

	if related.R == nil {
		related.R = &organismR{
			Organismprop: o,
		}
	} else {
		related.R.Organismprop = o
	}

	return nil
}

// SetType of the organismprop to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeOrganismprop.
func (o *Organismprop) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"organismprop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, organismpropPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.OrganismpropID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &organismpropR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeOrganismprop: o,
		}
	} else {
		related.R.TypeOrganismprop = o
	}

	return nil
}

// OrganismpropsG retrieves all records.
func OrganismpropsG(mods ...qm.QueryMod) organismpropQuery {
	return Organismprops(boil.GetDB(), mods...)
}

// Organismprops retrieves all the records using an executor.
func Organismprops(exec boil.Executor, mods ...qm.QueryMod) organismpropQuery {
	mods = append(mods, qm.From("\"organismprop\""))
	return organismpropQuery{NewQuery(exec, mods...)}
}

// FindOrganismpropG retrieves a single record by ID.
func FindOrganismpropG(organismpropID int, selectCols ...string) (*Organismprop, error) {
	return FindOrganismprop(boil.GetDB(), organismpropID, selectCols...)
}

// FindOrganismpropGP retrieves a single record by ID, and panics on error.
func FindOrganismpropGP(organismpropID int, selectCols ...string) *Organismprop {
	retobj, err := FindOrganismprop(boil.GetDB(), organismpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindOrganismprop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrganismprop(exec boil.Executor, organismpropID int, selectCols ...string) (*Organismprop, error) {
	organismpropObj := &Organismprop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"organismprop\" where \"organismprop_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, organismpropID)

	err := q.Bind(organismpropObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from organismprop")
	}

	return organismpropObj, nil
}

// FindOrganismpropP retrieves a single record by ID with an executor, and panics on error.
func FindOrganismpropP(exec boil.Executor, organismpropID int, selectCols ...string) *Organismprop {
	retobj, err := FindOrganismprop(exec, organismpropID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Organismprop) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Organismprop) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Organismprop) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Organismprop) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no organismprop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(organismpropColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	organismpropInsertCacheMut.RLock()
	cache, cached := organismpropInsertCache[key]
	organismpropInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			organismpropColumns,
			organismpropColumnsWithDefault,
			organismpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(organismpropType, organismpropMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(organismpropType, organismpropMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"organismprop\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into organismprop")
	}

	if !cached {
		organismpropInsertCacheMut.Lock()
		organismpropInsertCache[key] = cache
		organismpropInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Organismprop record. See Update for
// whitelist behavior description.
func (o *Organismprop) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Organismprop record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Organismprop) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Organismprop, and panics on error.
// See Update for whitelist behavior description.
func (o *Organismprop) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Organismprop.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Organismprop) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	organismpropUpdateCacheMut.RLock()
	cache, cached := organismpropUpdateCache[key]
	organismpropUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(organismpropColumns, organismpropPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update organismprop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"organismprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, organismpropPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(organismpropType, organismpropMapping, append(wl, organismpropPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update organismprop row")
	}

	if !cached {
		organismpropUpdateCacheMut.Lock()
		organismpropUpdateCache[key] = cache
		organismpropUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q organismpropQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q organismpropQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for organismprop")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OrganismpropSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o OrganismpropSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o OrganismpropSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrganismpropSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organismpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"organismprop\" SET %s WHERE (\"organismprop_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(organismpropPrimaryKeyColumns), len(colNames)+1, len(organismpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in organismprop slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Organismprop) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Organismprop) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Organismprop) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Organismprop) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no organismprop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(organismpropColumnsWithDefault, o)

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

	organismpropUpsertCacheMut.RLock()
	cache, cached := organismpropUpsertCache[key]
	organismpropUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			organismpropColumns,
			organismpropColumnsWithDefault,
			organismpropColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			organismpropColumns,
			organismpropPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert organismprop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(organismpropPrimaryKeyColumns))
			copy(conflict, organismpropPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"organismprop\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(organismpropType, organismpropMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(organismpropType, organismpropMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for organismprop")
	}

	if !cached {
		organismpropUpsertCacheMut.Lock()
		organismpropUpsertCache[key] = cache
		organismpropUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Organismprop record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Organismprop) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Organismprop record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Organismprop) DeleteG() error {
	if o == nil {
		return errors.New("models: no Organismprop provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Organismprop record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Organismprop) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Organismprop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Organismprop) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Organismprop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), organismpropPrimaryKeyMapping)
	sql := "DELETE FROM \"organismprop\" WHERE \"organismprop_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from organismprop")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q organismpropQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q organismpropQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no organismpropQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from organismprop")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o OrganismpropSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o OrganismpropSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Organismprop slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o OrganismpropSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrganismpropSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Organismprop slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(organismpropBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organismpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"organismprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, organismpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(organismpropPrimaryKeyColumns), 1, len(organismpropPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from organismprop slice")
	}

	if len(organismpropAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Organismprop) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Organismprop) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Organismprop) ReloadG() error {
	if o == nil {
		return errors.New("models: no Organismprop provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Organismprop) Reload(exec boil.Executor) error {
	ret, err := FindOrganismprop(exec, o.OrganismpropID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OrganismpropSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OrganismpropSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrganismpropSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty OrganismpropSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrganismpropSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	organismprops := OrganismpropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organismpropPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"organismprop\".* FROM \"organismprop\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, organismpropPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(organismpropPrimaryKeyColumns), 1, len(organismpropPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&organismprops)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrganismpropSlice")
	}

	*o = organismprops

	return nil
}

// OrganismpropExists checks if the Organismprop row exists.
func OrganismpropExists(exec boil.Executor, organismpropID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"organismprop\" where \"organismprop_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, organismpropID)
	}

	row := exec.QueryRow(sql, organismpropID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if organismprop exists")
	}

	return exists, nil
}

// OrganismpropExistsG checks if the Organismprop row exists.
func OrganismpropExistsG(organismpropID int) (bool, error) {
	return OrganismpropExists(boil.GetDB(), organismpropID)
}

// OrganismpropExistsGP checks if the Organismprop row exists. Panics on error.
func OrganismpropExistsGP(organismpropID int) bool {
	e, err := OrganismpropExists(boil.GetDB(), organismpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// OrganismpropExistsP checks if the Organismprop row exists. Panics on error.
func OrganismpropExistsP(exec boil.Executor, organismpropID int) bool {
	e, err := OrganismpropExists(exec, organismpropID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
