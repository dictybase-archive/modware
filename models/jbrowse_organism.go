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
)

// JbrowseOrganism is an object representing the database table.
type JbrowseOrganism struct {
	JbrowseOrganismID int    `boil:"jbrowse_organism_id" json:"jbrowse_organism_id" toml:"jbrowse_organism_id" yaml:"jbrowse_organism_id"`
	OrganismID        int    `boil:"organism_id" json:"organism_id" toml:"organism_id" yaml:"organism_id"`
	JbrowseID         int    `boil:"jbrowse_id" json:"jbrowse_id" toml:"jbrowse_id" yaml:"jbrowse_id"`
	Dataset           string `boil:"dataset" json:"dataset" toml:"dataset" yaml:"dataset"`

	R *jbrowseOrganismR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L jbrowseOrganismL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// jbrowseOrganismR is where relationships are stored.
type jbrowseOrganismR struct {
	Organism     *Organism
	Jbrowse      *Jbrowse
	JbrowseTrack *JbrowseTrack
}

// jbrowseOrganismL is where Load methods for each relationship are stored.
type jbrowseOrganismL struct{}

var (
	jbrowseOrganismColumns               = []string{"jbrowse_organism_id", "organism_id", "jbrowse_id", "dataset"}
	jbrowseOrganismColumnsWithoutDefault = []string{"organism_id", "jbrowse_id", "dataset"}
	jbrowseOrganismColumnsWithDefault    = []string{"jbrowse_organism_id"}
	jbrowseOrganismPrimaryKeyColumns     = []string{"jbrowse_organism_id"}
)

type (
	// JbrowseOrganismSlice is an alias for a slice of pointers to JbrowseOrganism.
	// This should generally be used opposed to []JbrowseOrganism.
	JbrowseOrganismSlice []*JbrowseOrganism
	// JbrowseOrganismHook is the signature for custom JbrowseOrganism hook methods
	JbrowseOrganismHook func(boil.Executor, *JbrowseOrganism) error

	jbrowseOrganismQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	jbrowseOrganismType                 = reflect.TypeOf(&JbrowseOrganism{})
	jbrowseOrganismMapping              = queries.MakeStructMapping(jbrowseOrganismType)
	jbrowseOrganismPrimaryKeyMapping, _ = queries.BindMapping(jbrowseOrganismType, jbrowseOrganismMapping, jbrowseOrganismPrimaryKeyColumns)
	jbrowseOrganismInsertCacheMut       sync.RWMutex
	jbrowseOrganismInsertCache          = make(map[string]insertCache)
	jbrowseOrganismUpdateCacheMut       sync.RWMutex
	jbrowseOrganismUpdateCache          = make(map[string]updateCache)
	jbrowseOrganismUpsertCacheMut       sync.RWMutex
	jbrowseOrganismUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var jbrowseOrganismBeforeInsertHooks []JbrowseOrganismHook
var jbrowseOrganismBeforeUpdateHooks []JbrowseOrganismHook
var jbrowseOrganismBeforeDeleteHooks []JbrowseOrganismHook
var jbrowseOrganismBeforeUpsertHooks []JbrowseOrganismHook

var jbrowseOrganismAfterInsertHooks []JbrowseOrganismHook
var jbrowseOrganismAfterSelectHooks []JbrowseOrganismHook
var jbrowseOrganismAfterUpdateHooks []JbrowseOrganismHook
var jbrowseOrganismAfterDeleteHooks []JbrowseOrganismHook
var jbrowseOrganismAfterUpsertHooks []JbrowseOrganismHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *JbrowseOrganism) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseOrganismBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *JbrowseOrganism) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseOrganismBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *JbrowseOrganism) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseOrganismBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *JbrowseOrganism) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseOrganismBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *JbrowseOrganism) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseOrganismAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *JbrowseOrganism) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseOrganismAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *JbrowseOrganism) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseOrganismAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *JbrowseOrganism) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseOrganismAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *JbrowseOrganism) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range jbrowseOrganismAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddJbrowseOrganismHook registers your hook function for all future operations.
func AddJbrowseOrganismHook(hookPoint boil.HookPoint, jbrowseOrganismHook JbrowseOrganismHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		jbrowseOrganismBeforeInsertHooks = append(jbrowseOrganismBeforeInsertHooks, jbrowseOrganismHook)
	case boil.BeforeUpdateHook:
		jbrowseOrganismBeforeUpdateHooks = append(jbrowseOrganismBeforeUpdateHooks, jbrowseOrganismHook)
	case boil.BeforeDeleteHook:
		jbrowseOrganismBeforeDeleteHooks = append(jbrowseOrganismBeforeDeleteHooks, jbrowseOrganismHook)
	case boil.BeforeUpsertHook:
		jbrowseOrganismBeforeUpsertHooks = append(jbrowseOrganismBeforeUpsertHooks, jbrowseOrganismHook)
	case boil.AfterInsertHook:
		jbrowseOrganismAfterInsertHooks = append(jbrowseOrganismAfterInsertHooks, jbrowseOrganismHook)
	case boil.AfterSelectHook:
		jbrowseOrganismAfterSelectHooks = append(jbrowseOrganismAfterSelectHooks, jbrowseOrganismHook)
	case boil.AfterUpdateHook:
		jbrowseOrganismAfterUpdateHooks = append(jbrowseOrganismAfterUpdateHooks, jbrowseOrganismHook)
	case boil.AfterDeleteHook:
		jbrowseOrganismAfterDeleteHooks = append(jbrowseOrganismAfterDeleteHooks, jbrowseOrganismHook)
	case boil.AfterUpsertHook:
		jbrowseOrganismAfterUpsertHooks = append(jbrowseOrganismAfterUpsertHooks, jbrowseOrganismHook)
	}
}

// OneP returns a single jbrowseOrganism record from the query, and panics on error.
func (q jbrowseOrganismQuery) OneP() *JbrowseOrganism {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single jbrowseOrganism record from the query.
func (q jbrowseOrganismQuery) One() (*JbrowseOrganism, error) {
	o := &JbrowseOrganism{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for jbrowse_organism")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all JbrowseOrganism records from the query, and panics on error.
func (q jbrowseOrganismQuery) AllP() JbrowseOrganismSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all JbrowseOrganism records from the query.
func (q jbrowseOrganismQuery) All() (JbrowseOrganismSlice, error) {
	var o JbrowseOrganismSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to JbrowseOrganism slice")
	}

	if len(jbrowseOrganismAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all JbrowseOrganism records in the query, and panics on error.
func (q jbrowseOrganismQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all JbrowseOrganism records in the query.
func (q jbrowseOrganismQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count jbrowse_organism rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q jbrowseOrganismQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q jbrowseOrganismQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if jbrowse_organism exists")
	}

	return count > 0, nil
}

// OrganismG pointed to by the foreign key.
func (o *JbrowseOrganism) OrganismG(mods ...qm.QueryMod) organismQuery {
	return o.Organism(boil.GetDB(), mods...)
}

// Organism pointed to by the foreign key.
func (o *JbrowseOrganism) Organism(exec boil.Executor, mods ...qm.QueryMod) organismQuery {
	queryMods := []qm.QueryMod{
		qm.Where("organism_id=$1", o.OrganismID),
	}

	queryMods = append(queryMods, mods...)

	query := Organisms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"organism\"")

	return query
}

// JbrowseG pointed to by the foreign key.
func (o *JbrowseOrganism) JbrowseG(mods ...qm.QueryMod) jbrowseQuery {
	return o.Jbrowse(boil.GetDB(), mods...)
}

// Jbrowse pointed to by the foreign key.
func (o *JbrowseOrganism) Jbrowse(exec boil.Executor, mods ...qm.QueryMod) jbrowseQuery {
	queryMods := []qm.QueryMod{
		qm.Where("jbrowse_id=$1", o.JbrowseID),
	}

	queryMods = append(queryMods, mods...)

	query := Jbrowses(exec, queryMods...)
	queries.SetFrom(query.Query, "\"jbrowse\"")

	return query
}

// JbrowseTrackG pointed to by the foreign key.
func (o *JbrowseOrganism) JbrowseTrackG(mods ...qm.QueryMod) jbrowseTrackQuery {
	return o.JbrowseTrack(boil.GetDB(), mods...)
}

// JbrowseTrack pointed to by the foreign key.
func (o *JbrowseOrganism) JbrowseTrack(exec boil.Executor, mods ...qm.QueryMod) jbrowseTrackQuery {
	queryMods := []qm.QueryMod{
		qm.Where("jbrowse_organism_id=$1", o.JbrowseOrganismID),
	}

	queryMods = append(queryMods, mods...)

	query := JbrowseTracks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"jbrowse_track\"")

	return query
}

// LoadOrganism allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (jbrowseOrganismL) LoadOrganism(e boil.Executor, singular bool, maybeJbrowseOrganism interface{}) error {
	var slice []*JbrowseOrganism
	var object *JbrowseOrganism

	count := 1
	if singular {
		object = maybeJbrowseOrganism.(*JbrowseOrganism)
	} else {
		slice = *maybeJbrowseOrganism.(*JbrowseOrganismSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &jbrowseOrganismR{}
		args[0] = object.OrganismID
	} else {
		for i, obj := range slice {
			obj.R = &jbrowseOrganismR{}
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

	if len(jbrowseOrganismAfterSelectHooks) != 0 {
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

// LoadJbrowse allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (jbrowseOrganismL) LoadJbrowse(e boil.Executor, singular bool, maybeJbrowseOrganism interface{}) error {
	var slice []*JbrowseOrganism
	var object *JbrowseOrganism

	count := 1
	if singular {
		object = maybeJbrowseOrganism.(*JbrowseOrganism)
	} else {
		slice = *maybeJbrowseOrganism.(*JbrowseOrganismSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &jbrowseOrganismR{}
		args[0] = object.JbrowseID
	} else {
		for i, obj := range slice {
			obj.R = &jbrowseOrganismR{}
			args[i] = obj.JbrowseID
		}
	}

	query := fmt.Sprintf(
		"select * from \"jbrowse\" where \"jbrowse_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Jbrowse")
	}
	defer results.Close()

	var resultSlice []*Jbrowse
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Jbrowse")
	}

	if len(jbrowseOrganismAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Jbrowse = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.JbrowseID == foreign.JbrowseID {
				local.R.Jbrowse = foreign
				break
			}
		}
	}

	return nil
}

// LoadJbrowseTrack allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (jbrowseOrganismL) LoadJbrowseTrack(e boil.Executor, singular bool, maybeJbrowseOrganism interface{}) error {
	var slice []*JbrowseOrganism
	var object *JbrowseOrganism

	count := 1
	if singular {
		object = maybeJbrowseOrganism.(*JbrowseOrganism)
	} else {
		slice = *maybeJbrowseOrganism.(*JbrowseOrganismSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &jbrowseOrganismR{}
		args[0] = object.JbrowseOrganismID
	} else {
		for i, obj := range slice {
			obj.R = &jbrowseOrganismR{}
			args[i] = obj.JbrowseOrganismID
		}
	}

	query := fmt.Sprintf(
		"select * from \"jbrowse_track\" where \"jbrowse_organism_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load JbrowseTrack")
	}
	defer results.Close()

	var resultSlice []*JbrowseTrack
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice JbrowseTrack")
	}

	if len(jbrowseOrganismAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.JbrowseTrack = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.JbrowseOrganismID == foreign.JbrowseOrganismID {
				local.R.JbrowseTrack = foreign
				break
			}
		}
	}

	return nil
}

// SetOrganism of the jbrowse_organism to the related item.
// Sets o.R.Organism to related.
// Adds o to related.R.JbrowseOrganism.
func (o *JbrowseOrganism) SetOrganism(exec boil.Executor, insert bool, related *Organism) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"jbrowse_organism\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"organism_id"}),
		strmangle.WhereClause("\"", "\"", 2, jbrowseOrganismPrimaryKeyColumns),
	)
	values := []interface{}{related.OrganismID, o.JbrowseOrganismID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrganismID = related.OrganismID

	if o.R == nil {
		o.R = &jbrowseOrganismR{
			Organism: related,
		}
	} else {
		o.R.Organism = related
	}

	if related.R == nil {
		related.R = &organismR{
			JbrowseOrganism: o,
		}
	} else {
		related.R.JbrowseOrganism = o
	}

	return nil
}

// SetJbrowse of the jbrowse_organism to the related item.
// Sets o.R.Jbrowse to related.
// Adds o to related.R.JbrowseOrganism.
func (o *JbrowseOrganism) SetJbrowse(exec boil.Executor, insert bool, related *Jbrowse) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"jbrowse_organism\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"jbrowse_id"}),
		strmangle.WhereClause("\"", "\"", 2, jbrowseOrganismPrimaryKeyColumns),
	)
	values := []interface{}{related.JbrowseID, o.JbrowseOrganismID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.JbrowseID = related.JbrowseID

	if o.R == nil {
		o.R = &jbrowseOrganismR{
			Jbrowse: related,
		}
	} else {
		o.R.Jbrowse = related
	}

	if related.R == nil {
		related.R = &jbrowseR{
			JbrowseOrganism: o,
		}
	} else {
		related.R.JbrowseOrganism = o
	}

	return nil
}

// SetJbrowseTrack of the jbrowse_organism to the related item.
// Sets o.R.JbrowseTrack to related.
// Adds o to related.R.JbrowseOrganism.
func (o *JbrowseOrganism) SetJbrowseTrack(exec boil.Executor, insert bool, related *JbrowseTrack) error {
	var err error

	if insert {
		related.JbrowseOrganismID = o.JbrowseOrganismID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"jbrowse_track\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"jbrowse_organism_id"}),
			strmangle.WhereClause("\"", "\"", 2, jbrowseTrackPrimaryKeyColumns),
		)
		values := []interface{}{o.JbrowseOrganismID, related.JbrowseTrackID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.JbrowseOrganismID = o.JbrowseOrganismID

	}

	if o.R == nil {
		o.R = &jbrowseOrganismR{
			JbrowseTrack: related,
		}
	} else {
		o.R.JbrowseTrack = related
	}

	if related.R == nil {
		related.R = &jbrowseTrackR{
			JbrowseOrganism: o,
		}
	} else {
		related.R.JbrowseOrganism = o
	}
	return nil
}

// JbrowseOrganismsG retrieves all records.
func JbrowseOrganismsG(mods ...qm.QueryMod) jbrowseOrganismQuery {
	return JbrowseOrganisms(boil.GetDB(), mods...)
}

// JbrowseOrganisms retrieves all the records using an executor.
func JbrowseOrganisms(exec boil.Executor, mods ...qm.QueryMod) jbrowseOrganismQuery {
	mods = append(mods, qm.From("\"jbrowse_organism\""))
	return jbrowseOrganismQuery{NewQuery(exec, mods...)}
}

// FindJbrowseOrganismG retrieves a single record by ID.
func FindJbrowseOrganismG(jbrowseOrganismID int, selectCols ...string) (*JbrowseOrganism, error) {
	return FindJbrowseOrganism(boil.GetDB(), jbrowseOrganismID, selectCols...)
}

// FindJbrowseOrganismGP retrieves a single record by ID, and panics on error.
func FindJbrowseOrganismGP(jbrowseOrganismID int, selectCols ...string) *JbrowseOrganism {
	retobj, err := FindJbrowseOrganism(boil.GetDB(), jbrowseOrganismID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindJbrowseOrganism retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindJbrowseOrganism(exec boil.Executor, jbrowseOrganismID int, selectCols ...string) (*JbrowseOrganism, error) {
	jbrowseOrganismObj := &JbrowseOrganism{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"jbrowse_organism\" where \"jbrowse_organism_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, jbrowseOrganismID)

	err := q.Bind(jbrowseOrganismObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from jbrowse_organism")
	}

	return jbrowseOrganismObj, nil
}

// FindJbrowseOrganismP retrieves a single record by ID with an executor, and panics on error.
func FindJbrowseOrganismP(exec boil.Executor, jbrowseOrganismID int, selectCols ...string) *JbrowseOrganism {
	retobj, err := FindJbrowseOrganism(exec, jbrowseOrganismID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *JbrowseOrganism) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *JbrowseOrganism) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *JbrowseOrganism) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *JbrowseOrganism) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no jbrowse_organism provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(jbrowseOrganismColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	jbrowseOrganismInsertCacheMut.RLock()
	cache, cached := jbrowseOrganismInsertCache[key]
	jbrowseOrganismInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			jbrowseOrganismColumns,
			jbrowseOrganismColumnsWithDefault,
			jbrowseOrganismColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(jbrowseOrganismType, jbrowseOrganismMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(jbrowseOrganismType, jbrowseOrganismMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"jbrowse_organism\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into jbrowse_organism")
	}

	if !cached {
		jbrowseOrganismInsertCacheMut.Lock()
		jbrowseOrganismInsertCache[key] = cache
		jbrowseOrganismInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single JbrowseOrganism record. See Update for
// whitelist behavior description.
func (o *JbrowseOrganism) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single JbrowseOrganism record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *JbrowseOrganism) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the JbrowseOrganism, and panics on error.
// See Update for whitelist behavior description.
func (o *JbrowseOrganism) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the JbrowseOrganism.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *JbrowseOrganism) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	jbrowseOrganismUpdateCacheMut.RLock()
	cache, cached := jbrowseOrganismUpdateCache[key]
	jbrowseOrganismUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(jbrowseOrganismColumns, jbrowseOrganismPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update jbrowse_organism, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"jbrowse_organism\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, jbrowseOrganismPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(jbrowseOrganismType, jbrowseOrganismMapping, append(wl, jbrowseOrganismPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update jbrowse_organism row")
	}

	if !cached {
		jbrowseOrganismUpdateCacheMut.Lock()
		jbrowseOrganismUpdateCache[key] = cache
		jbrowseOrganismUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q jbrowseOrganismQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q jbrowseOrganismQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for jbrowse_organism")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o JbrowseOrganismSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o JbrowseOrganismSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o JbrowseOrganismSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o JbrowseOrganismSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jbrowseOrganismPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"jbrowse_organism\" SET %s WHERE (\"jbrowse_organism_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(jbrowseOrganismPrimaryKeyColumns), len(colNames)+1, len(jbrowseOrganismPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in jbrowseOrganism slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *JbrowseOrganism) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *JbrowseOrganism) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *JbrowseOrganism) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *JbrowseOrganism) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no jbrowse_organism provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(jbrowseOrganismColumnsWithDefault, o)

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

	jbrowseOrganismUpsertCacheMut.RLock()
	cache, cached := jbrowseOrganismUpsertCache[key]
	jbrowseOrganismUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			jbrowseOrganismColumns,
			jbrowseOrganismColumnsWithDefault,
			jbrowseOrganismColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			jbrowseOrganismColumns,
			jbrowseOrganismPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert jbrowse_organism, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(jbrowseOrganismPrimaryKeyColumns))
			copy(conflict, jbrowseOrganismPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"jbrowse_organism\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(jbrowseOrganismType, jbrowseOrganismMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(jbrowseOrganismType, jbrowseOrganismMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for jbrowse_organism")
	}

	if !cached {
		jbrowseOrganismUpsertCacheMut.Lock()
		jbrowseOrganismUpsertCache[key] = cache
		jbrowseOrganismUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single JbrowseOrganism record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *JbrowseOrganism) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single JbrowseOrganism record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *JbrowseOrganism) DeleteG() error {
	if o == nil {
		return errors.New("models: no JbrowseOrganism provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single JbrowseOrganism record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *JbrowseOrganism) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single JbrowseOrganism record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *JbrowseOrganism) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no JbrowseOrganism provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), jbrowseOrganismPrimaryKeyMapping)
	sql := "DELETE FROM \"jbrowse_organism\" WHERE \"jbrowse_organism_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from jbrowse_organism")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q jbrowseOrganismQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q jbrowseOrganismQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no jbrowseOrganismQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from jbrowse_organism")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o JbrowseOrganismSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o JbrowseOrganismSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no JbrowseOrganism slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o JbrowseOrganismSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o JbrowseOrganismSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no JbrowseOrganism slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(jbrowseOrganismBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jbrowseOrganismPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"jbrowse_organism\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, jbrowseOrganismPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(jbrowseOrganismPrimaryKeyColumns), 1, len(jbrowseOrganismPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from jbrowseOrganism slice")
	}

	if len(jbrowseOrganismAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *JbrowseOrganism) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *JbrowseOrganism) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *JbrowseOrganism) ReloadG() error {
	if o == nil {
		return errors.New("models: no JbrowseOrganism provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *JbrowseOrganism) Reload(exec boil.Executor) error {
	ret, err := FindJbrowseOrganism(exec, o.JbrowseOrganismID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *JbrowseOrganismSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *JbrowseOrganismSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JbrowseOrganismSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty JbrowseOrganismSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JbrowseOrganismSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	jbrowseOrganisms := JbrowseOrganismSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jbrowseOrganismPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"jbrowse_organism\".* FROM \"jbrowse_organism\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, jbrowseOrganismPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(jbrowseOrganismPrimaryKeyColumns), 1, len(jbrowseOrganismPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&jbrowseOrganisms)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in JbrowseOrganismSlice")
	}

	*o = jbrowseOrganisms

	return nil
}

// JbrowseOrganismExists checks if the JbrowseOrganism row exists.
func JbrowseOrganismExists(exec boil.Executor, jbrowseOrganismID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"jbrowse_organism\" where \"jbrowse_organism_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, jbrowseOrganismID)
	}

	row := exec.QueryRow(sql, jbrowseOrganismID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if jbrowse_organism exists")
	}

	return exists, nil
}

// JbrowseOrganismExistsG checks if the JbrowseOrganism row exists.
func JbrowseOrganismExistsG(jbrowseOrganismID int) (bool, error) {
	return JbrowseOrganismExists(boil.GetDB(), jbrowseOrganismID)
}

// JbrowseOrganismExistsGP checks if the JbrowseOrganism row exists. Panics on error.
func JbrowseOrganismExistsGP(jbrowseOrganismID int) bool {
	e, err := JbrowseOrganismExists(boil.GetDB(), jbrowseOrganismID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// JbrowseOrganismExistsP checks if the JbrowseOrganism row exists. Panics on error.
func JbrowseOrganismExistsP(exec boil.Executor, jbrowseOrganismID int) bool {
	e, err := JbrowseOrganismExists(exec, jbrowseOrganismID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
