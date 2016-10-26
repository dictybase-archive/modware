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
	"github.com/vattle/sqlboiler/types"
	"gopkg.in/nullbio/null.v5"
)

// JbrowseTrack is an object representing the database table.
type JbrowseTrack struct {
	JbrowseTrackID    int        `boil:"jbrowse_track_id" json:"jbrowse_track_id" toml:"jbrowse_track_id" yaml:"jbrowse_track_id"`
	Configuration     types.JSON `boil:"configuration" json:"configuration" toml:"configuration" yaml:"configuration"`
	TypeID            null.Int   `boil:"type_id" json:"type_id,omitempty" toml:"type_id" yaml:"type_id,omitempty"`
	JbrowseOrganismID int        `boil:"jbrowse_organism_id" json:"jbrowse_organism_id" toml:"jbrowse_organism_id" yaml:"jbrowse_organism_id"`

	R *jbrowseTrackR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L jbrowseTrackL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// jbrowseTrackR is where relationships are stored.
type jbrowseTrackR struct {
	Type            *Cvterm
	JbrowseOrganism *JbrowseOrganism
}

// jbrowseTrackL is where Load methods for each relationship are stored.
type jbrowseTrackL struct{}

var (
	jbrowseTrackColumns               = []string{"jbrowse_track_id", "configuration", "type_id", "jbrowse_organism_id"}
	jbrowseTrackColumnsWithoutDefault = []string{"configuration", "type_id", "jbrowse_organism_id"}
	jbrowseTrackColumnsWithDefault    = []string{"jbrowse_track_id"}
	jbrowseTrackPrimaryKeyColumns     = []string{"jbrowse_track_id"}
)

type (
	// JbrowseTrackSlice is an alias for a slice of pointers to JbrowseTrack.
	// This should generally be used opposed to []JbrowseTrack.
	JbrowseTrackSlice []*JbrowseTrack
	// JbrowseTrackHook is the signature for custom JbrowseTrack hook methods
	JbrowseTrackHook func(boil.Executor, *JbrowseTrack) error

	jbrowseTrackQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	jbrowseTrackType                 = reflect.TypeOf(&JbrowseTrack{})
	jbrowseTrackMapping              = queries.MakeStructMapping(jbrowseTrackType)
	jbrowseTrackPrimaryKeyMapping, _ = queries.BindMapping(jbrowseTrackType, jbrowseTrackMapping, jbrowseTrackPrimaryKeyColumns)
	jbrowseTrackInsertCacheMut       sync.RWMutex
	jbrowseTrackInsertCache          = make(map[string]insertCache)
	jbrowseTrackUpdateCacheMut       sync.RWMutex
	jbrowseTrackUpdateCache          = make(map[string]updateCache)
	jbrowseTrackUpsertCacheMut       sync.RWMutex
	jbrowseTrackUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var jbrowseTrackBeforeInsertHooks []JbrowseTrackHook
var jbrowseTrackBeforeUpdateHooks []JbrowseTrackHook
var jbrowseTrackBeforeDeleteHooks []JbrowseTrackHook
var jbrowseTrackBeforeUpsertHooks []JbrowseTrackHook

var jbrowseTrackAfterInsertHooks []JbrowseTrackHook
var jbrowseTrackAfterSelectHooks []JbrowseTrackHook
var jbrowseTrackAfterUpdateHooks []JbrowseTrackHook
var jbrowseTrackAfterDeleteHooks []JbrowseTrackHook
var jbrowseTrackAfterUpsertHooks []JbrowseTrackHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *JbrowseTrack) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseTrackBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *JbrowseTrack) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseTrackBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *JbrowseTrack) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseTrackBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *JbrowseTrack) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseTrackBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *JbrowseTrack) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseTrackAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *JbrowseTrack) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseTrackAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *JbrowseTrack) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseTrackAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *JbrowseTrack) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseTrackAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *JbrowseTrack) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseTrackAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddJbrowseTrackHook registers your hook function for all future operations.
func AddJbrowseTrackHook(hookPoint boil.HookPoint, jbrowseTrackHook JbrowseTrackHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		jbrowseTrackBeforeInsertHooks = append(jbrowseTrackBeforeInsertHooks, jbrowseTrackHook)
	case boil.BeforeUpdateHook:
		jbrowseTrackBeforeUpdateHooks = append(jbrowseTrackBeforeUpdateHooks, jbrowseTrackHook)
	case boil.BeforeDeleteHook:
		jbrowseTrackBeforeDeleteHooks = append(jbrowseTrackBeforeDeleteHooks, jbrowseTrackHook)
	case boil.BeforeUpsertHook:
		jbrowseTrackBeforeUpsertHooks = append(jbrowseTrackBeforeUpsertHooks, jbrowseTrackHook)
	case boil.AfterInsertHook:
		jbrowseTrackAfterInsertHooks = append(jbrowseTrackAfterInsertHooks, jbrowseTrackHook)
	case boil.AfterSelectHook:
		jbrowseTrackAfterSelectHooks = append(jbrowseTrackAfterSelectHooks, jbrowseTrackHook)
	case boil.AfterUpdateHook:
		jbrowseTrackAfterUpdateHooks = append(jbrowseTrackAfterUpdateHooks, jbrowseTrackHook)
	case boil.AfterDeleteHook:
		jbrowseTrackAfterDeleteHooks = append(jbrowseTrackAfterDeleteHooks, jbrowseTrackHook)
	case boil.AfterUpsertHook:
		jbrowseTrackAfterUpsertHooks = append(jbrowseTrackAfterUpsertHooks, jbrowseTrackHook)
	}
}

// OneP returns a single jbrowseTrack record from the query, and panics on error.
func (q jbrowseTrackQuery) OneP() *JbrowseTrack {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single jbrowseTrack record from the query.
func (q jbrowseTrackQuery) One() (*JbrowseTrack, error) {
	o := &JbrowseTrack{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for jbrowse_track")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all JbrowseTrack records from the query, and panics on error.
func (q jbrowseTrackQuery) AllP() JbrowseTrackSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all JbrowseTrack records from the query.
func (q jbrowseTrackQuery) All() (JbrowseTrackSlice, error) {
	var o JbrowseTrackSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to JbrowseTrack slice")
	}

	if len(jbrowseTrackAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all JbrowseTrack records in the query, and panics on error.
func (q jbrowseTrackQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all JbrowseTrack records in the query.
func (q jbrowseTrackQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count jbrowse_track rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q jbrowseTrackQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q jbrowseTrackQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if jbrowse_track exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *JbrowseTrack) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *JbrowseTrack) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// JbrowseOrganismG pointed to by the foreign key.
func (o *JbrowseTrack) JbrowseOrganismG(mods ...qm.QueryMod) jbrowseOrganismQuery {
	return o.JbrowseOrganism(boil.GetDB(), mods...)
}

// JbrowseOrganism pointed to by the foreign key.
func (o *JbrowseTrack) JbrowseOrganism(exec boil.Executor, mods ...qm.QueryMod) jbrowseOrganismQuery {
	queryMods := []qm.QueryMod{
		qm.Where("jbrowse_organism_id=$1", o.JbrowseOrganismID),
	}

	queryMods = append(queryMods, mods...)

	query := JbrowseOrganisms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"jbrowse_organism\"")

	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (jbrowseTrackL) LoadType(e boil.Executor, singular bool, maybeJbrowseTrack interface{}) error {
	var slice []*JbrowseTrack
	var object *JbrowseTrack

	count := 1
	if singular {
		object = maybeJbrowseTrack.(*JbrowseTrack)
	} else {
		slice = *maybeJbrowseTrack.(*JbrowseTrackSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &jbrowseTrackR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &jbrowseTrackR{}
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

	if len(jbrowseTrackAfterSelectHooks) != 0 {
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

// LoadJbrowseOrganism allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (jbrowseTrackL) LoadJbrowseOrganism(e boil.Executor, singular bool, maybeJbrowseTrack interface{}) error {
	var slice []*JbrowseTrack
	var object *JbrowseTrack

	count := 1
	if singular {
		object = maybeJbrowseTrack.(*JbrowseTrack)
	} else {
		slice = *maybeJbrowseTrack.(*JbrowseTrackSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &jbrowseTrackR{}
		args[0] = object.JbrowseOrganismID
	} else {
		for i, obj := range slice {
			obj.R = &jbrowseTrackR{}
			args[i] = obj.JbrowseOrganismID
		}
	}

	query := fmt.Sprintf(
		"select * from \"jbrowse_organism\" where \"jbrowse_organism_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load JbrowseOrganism")
	}
	defer results.Close()

	var resultSlice []*JbrowseOrganism
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice JbrowseOrganism")
	}

	if len(jbrowseTrackAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.JbrowseOrganism = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.JbrowseOrganismID == foreign.JbrowseOrganismID {
				local.R.JbrowseOrganism = foreign
				break
			}
		}
	}

	return nil
}

// SetType of the jbrowse_track to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeJbrowseTracks.
func (o *JbrowseTrack) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"jbrowse_track\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, jbrowseTrackPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.JbrowseTrackID}

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
		o.R = &jbrowseTrackR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeJbrowseTracks: JbrowseTrackSlice{o},
		}
	} else {
		related.R.TypeJbrowseTracks = append(related.R.TypeJbrowseTracks, o)
	}

	return nil
}

// RemoveType relationship.
// Sets o.R.Type to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *JbrowseTrack) RemoveType(exec boil.Executor, related *Cvterm) error {
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

	for i, ri := range related.R.TypeJbrowseTracks {
		if o.TypeID.Int != ri.TypeID.Int {
			continue
		}

		ln := len(related.R.TypeJbrowseTracks)
		if ln > 1 && i < ln-1 {
			related.R.TypeJbrowseTracks[i] = related.R.TypeJbrowseTracks[ln-1]
		}
		related.R.TypeJbrowseTracks = related.R.TypeJbrowseTracks[:ln-1]
		break
	}
	return nil
}

// SetJbrowseOrganism of the jbrowse_track to the related item.
// Sets o.R.JbrowseOrganism to related.
// Adds o to related.R.JbrowseTrack.
func (o *JbrowseTrack) SetJbrowseOrganism(exec boil.Executor, insert bool, related *JbrowseOrganism) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"jbrowse_track\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"jbrowse_organism_id"}),
		strmangle.WhereClause("\"", "\"", 2, jbrowseTrackPrimaryKeyColumns),
	)
	values := []interface{}{related.JbrowseOrganismID, o.JbrowseTrackID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.JbrowseOrganismID = related.JbrowseOrganismID

	if o.R == nil {
		o.R = &jbrowseTrackR{
			JbrowseOrganism: related,
		}
	} else {
		o.R.JbrowseOrganism = related
	}

	if related.R == nil {
		related.R = &jbrowseOrganismR{
			JbrowseTrack: o,
		}
	} else {
		related.R.JbrowseTrack = o
	}

	return nil
}

// JbrowseTracksG retrieves all records.
func JbrowseTracksG(mods ...qm.QueryMod) jbrowseTrackQuery {
	return JbrowseTracks(boil.GetDB(), mods...)
}

// JbrowseTracks retrieves all the records using an executor.
func JbrowseTracks(exec boil.Executor, mods ...qm.QueryMod) jbrowseTrackQuery {
	mods = append(mods, qm.From("\"jbrowse_track\""))
	return jbrowseTrackQuery{NewQuery(exec, mods...)}
}

// FindJbrowseTrackG retrieves a single record by ID.
func FindJbrowseTrackG(jbrowseTrackID int, selectCols ...string) (*JbrowseTrack, error) {
	return FindJbrowseTrack(boil.GetDB(), jbrowseTrackID, selectCols...)
}

// FindJbrowseTrackGP retrieves a single record by ID, and panics on error.
func FindJbrowseTrackGP(jbrowseTrackID int, selectCols ...string) *JbrowseTrack {
	retobj, err := FindJbrowseTrack(boil.GetDB(), jbrowseTrackID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindJbrowseTrack retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindJbrowseTrack(exec boil.Executor, jbrowseTrackID int, selectCols ...string) (*JbrowseTrack, error) {
	jbrowseTrackObj := &JbrowseTrack{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"jbrowse_track\" where \"jbrowse_track_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, jbrowseTrackID)

	err := q.Bind(jbrowseTrackObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from jbrowse_track")
	}

	return jbrowseTrackObj, nil
}

// FindJbrowseTrackP retrieves a single record by ID with an executor, and panics on error.
func FindJbrowseTrackP(exec boil.Executor, jbrowseTrackID int, selectCols ...string) *JbrowseTrack {
	retobj, err := FindJbrowseTrack(exec, jbrowseTrackID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *JbrowseTrack) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *JbrowseTrack) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *JbrowseTrack) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *JbrowseTrack) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no jbrowse_track provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(jbrowseTrackColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	jbrowseTrackInsertCacheMut.RLock()
	cache, cached := jbrowseTrackInsertCache[key]
	jbrowseTrackInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			jbrowseTrackColumns,
			jbrowseTrackColumnsWithDefault,
			jbrowseTrackColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(jbrowseTrackType, jbrowseTrackMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(jbrowseTrackType, jbrowseTrackMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"jbrowse_track\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into jbrowse_track")
	}

	if !cached {
		jbrowseTrackInsertCacheMut.Lock()
		jbrowseTrackInsertCache[key] = cache
		jbrowseTrackInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single JbrowseTrack record. See Update for
// whitelist behavior description.
func (o *JbrowseTrack) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single JbrowseTrack record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *JbrowseTrack) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the JbrowseTrack, and panics on error.
// See Update for whitelist behavior description.
func (o *JbrowseTrack) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the JbrowseTrack.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *JbrowseTrack) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	jbrowseTrackUpdateCacheMut.RLock()
	cache, cached := jbrowseTrackUpdateCache[key]
	jbrowseTrackUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(jbrowseTrackColumns, jbrowseTrackPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update jbrowse_track, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"jbrowse_track\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, jbrowseTrackPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(jbrowseTrackType, jbrowseTrackMapping, append(wl, jbrowseTrackPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update jbrowse_track row")
	}

	if !cached {
		jbrowseTrackUpdateCacheMut.Lock()
		jbrowseTrackUpdateCache[key] = cache
		jbrowseTrackUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q jbrowseTrackQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q jbrowseTrackQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for jbrowse_track")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o JbrowseTrackSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o JbrowseTrackSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o JbrowseTrackSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o JbrowseTrackSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jbrowseTrackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"jbrowse_track\" SET %s WHERE (\"jbrowse_track_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(jbrowseTrackPrimaryKeyColumns), len(colNames)+1, len(jbrowseTrackPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in jbrowseTrack slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *JbrowseTrack) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *JbrowseTrack) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *JbrowseTrack) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *JbrowseTrack) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no jbrowse_track provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(jbrowseTrackColumnsWithDefault, o)

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

	jbrowseTrackUpsertCacheMut.RLock()
	cache, cached := jbrowseTrackUpsertCache[key]
	jbrowseTrackUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			jbrowseTrackColumns,
			jbrowseTrackColumnsWithDefault,
			jbrowseTrackColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			jbrowseTrackColumns,
			jbrowseTrackPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert jbrowse_track, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(jbrowseTrackPrimaryKeyColumns))
			copy(conflict, jbrowseTrackPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"jbrowse_track\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(jbrowseTrackType, jbrowseTrackMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(jbrowseTrackType, jbrowseTrackMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for jbrowse_track")
	}

	if !cached {
		jbrowseTrackUpsertCacheMut.Lock()
		jbrowseTrackUpsertCache[key] = cache
		jbrowseTrackUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single JbrowseTrack record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *JbrowseTrack) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single JbrowseTrack record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *JbrowseTrack) DeleteG() error {
	if o == nil {
		return errors.New("chado: no JbrowseTrack provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single JbrowseTrack record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *JbrowseTrack) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single JbrowseTrack record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *JbrowseTrack) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no JbrowseTrack provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), jbrowseTrackPrimaryKeyMapping)
	sql := "DELETE FROM \"jbrowse_track\" WHERE \"jbrowse_track_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from jbrowse_track")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q jbrowseTrackQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q jbrowseTrackQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no jbrowseTrackQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from jbrowse_track")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o JbrowseTrackSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o JbrowseTrackSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no JbrowseTrack slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o JbrowseTrackSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o JbrowseTrackSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no JbrowseTrack slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(jbrowseTrackBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jbrowseTrackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"jbrowse_track\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, jbrowseTrackPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(jbrowseTrackPrimaryKeyColumns), 1, len(jbrowseTrackPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from jbrowseTrack slice")
	}

	if len(jbrowseTrackAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *JbrowseTrack) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *JbrowseTrack) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *JbrowseTrack) ReloadG() error {
	if o == nil {
		return errors.New("chado: no JbrowseTrack provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *JbrowseTrack) Reload(exec boil.Executor) error {
	ret, err := FindJbrowseTrack(exec, o.JbrowseTrackID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *JbrowseTrackSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *JbrowseTrackSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JbrowseTrackSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty JbrowseTrackSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JbrowseTrackSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	jbrowseTracks := JbrowseTrackSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jbrowseTrackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"jbrowse_track\".* FROM \"jbrowse_track\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, jbrowseTrackPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(jbrowseTrackPrimaryKeyColumns), 1, len(jbrowseTrackPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&jbrowseTracks)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in JbrowseTrackSlice")
	}

	*o = jbrowseTracks

	return nil
}

// JbrowseTrackExists checks if the JbrowseTrack row exists.
func JbrowseTrackExists(exec boil.Executor, jbrowseTrackID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"jbrowse_track\" where \"jbrowse_track_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, jbrowseTrackID)
	}

	row := exec.QueryRow(sql, jbrowseTrackID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if jbrowse_track exists")
	}

	return exists, nil
}

// JbrowseTrackExistsG checks if the JbrowseTrack row exists.
func JbrowseTrackExistsG(jbrowseTrackID int) (bool, error) {
	return JbrowseTrackExists(boil.GetDB(), jbrowseTrackID)
}

// JbrowseTrackExistsGP checks if the JbrowseTrack row exists. Panics on error.
func JbrowseTrackExistsGP(jbrowseTrackID int) bool {
	e, err := JbrowseTrackExists(boil.GetDB(), jbrowseTrackID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// JbrowseTrackExistsP checks if the JbrowseTrack row exists. Panics on error.
func JbrowseTrackExistsP(exec boil.Executor, jbrowseTrackID int) bool {
	e, err := JbrowseTrackExists(exec, jbrowseTrackID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
