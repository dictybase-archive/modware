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

// Phendesc is an object representing the database table.
type Phendesc struct {
	PhendescID    int    `boil:"phendesc_id" json:"phendesc_id" toml:"phendesc_id" yaml:"phendesc_id"`
	GenotypeID    int    `boil:"genotype_id" json:"genotype_id" toml:"genotype_id" yaml:"genotype_id"`
	EnvironmentID int    `boil:"environment_id" json:"environment_id" toml:"environment_id" yaml:"environment_id"`
	Description   string `boil:"description" json:"description" toml:"description" yaml:"description"`
	TypeID        int    `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	PubID         int    `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`

	R *phendescR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L phendescL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// phendescR is where relationships are stored.
type phendescR struct {
	Environment *Environment
	Genotype    *Genotype
	Type        *Cvterm
	Pub         *Pub
}

// phendescL is where Load methods for each relationship are stored.
type phendescL struct{}

var (
	phendescColumns               = []string{"phendesc_id", "genotype_id", "environment_id", "description", "type_id", "pub_id"}
	phendescColumnsWithoutDefault = []string{"genotype_id", "environment_id", "description", "type_id", "pub_id"}
	phendescColumnsWithDefault    = []string{"phendesc_id"}
	phendescPrimaryKeyColumns     = []string{"phendesc_id"}
)

type (
	// PhendescSlice is an alias for a slice of pointers to Phendesc.
	// This should generally be used opposed to []Phendesc.
	PhendescSlice []*Phendesc
	// PhendescHook is the signature for custom Phendesc hook methods
	PhendescHook func(boil.Executor, *Phendesc) error

	phendescQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	phendescType                 = reflect.TypeOf(&Phendesc{})
	phendescMapping              = queries.MakeStructMapping(phendescType)
	phendescPrimaryKeyMapping, _ = queries.BindMapping(phendescType, phendescMapping, phendescPrimaryKeyColumns)
	phendescInsertCacheMut       sync.RWMutex
	phendescInsertCache          = make(map[string]insertCache)
	phendescUpdateCacheMut       sync.RWMutex
	phendescUpdateCache          = make(map[string]updateCache)
	phendescUpsertCacheMut       sync.RWMutex
	phendescUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var phendescBeforeInsertHooks []PhendescHook
var phendescBeforeUpdateHooks []PhendescHook
var phendescBeforeDeleteHooks []PhendescHook
var phendescBeforeUpsertHooks []PhendescHook

var phendescAfterInsertHooks []PhendescHook
var phendescAfterSelectHooks []PhendescHook
var phendescAfterUpdateHooks []PhendescHook
var phendescAfterDeleteHooks []PhendescHook
var phendescAfterUpsertHooks []PhendescHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Phendesc) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phendescBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Phendesc) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range phendescBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Phendesc) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range phendescBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Phendesc) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phendescBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Phendesc) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phendescAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Phendesc) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range phendescAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Phendesc) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range phendescAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Phendesc) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range phendescAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Phendesc) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phendescAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPhendescHook registers your hook function for all future operations.
func AddPhendescHook(hookPoint boil.HookPoint, phendescHook PhendescHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		phendescBeforeInsertHooks = append(phendescBeforeInsertHooks, phendescHook)
	case boil.BeforeUpdateHook:
		phendescBeforeUpdateHooks = append(phendescBeforeUpdateHooks, phendescHook)
	case boil.BeforeDeleteHook:
		phendescBeforeDeleteHooks = append(phendescBeforeDeleteHooks, phendescHook)
	case boil.BeforeUpsertHook:
		phendescBeforeUpsertHooks = append(phendescBeforeUpsertHooks, phendescHook)
	case boil.AfterInsertHook:
		phendescAfterInsertHooks = append(phendescAfterInsertHooks, phendescHook)
	case boil.AfterSelectHook:
		phendescAfterSelectHooks = append(phendescAfterSelectHooks, phendescHook)
	case boil.AfterUpdateHook:
		phendescAfterUpdateHooks = append(phendescAfterUpdateHooks, phendescHook)
	case boil.AfterDeleteHook:
		phendescAfterDeleteHooks = append(phendescAfterDeleteHooks, phendescHook)
	case boil.AfterUpsertHook:
		phendescAfterUpsertHooks = append(phendescAfterUpsertHooks, phendescHook)
	}
}

// OneP returns a single phendesc record from the query, and panics on error.
func (q phendescQuery) OneP() *Phendesc {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single phendesc record from the query.
func (q phendescQuery) One() (*Phendesc, error) {
	o := &Phendesc{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for phendesc")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Phendesc records from the query, and panics on error.
func (q phendescQuery) AllP() PhendescSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Phendesc records from the query.
func (q phendescQuery) All() (PhendescSlice, error) {
	var o PhendescSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Phendesc slice")
	}

	if len(phendescAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Phendesc records in the query, and panics on error.
func (q phendescQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Phendesc records in the query.
func (q phendescQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count phendesc rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q phendescQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q phendescQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if phendesc exists")
	}

	return count > 0, nil
}

// EnvironmentG pointed to by the foreign key.
func (o *Phendesc) EnvironmentG(mods ...qm.QueryMod) environmentQuery {
	return o.Environment(boil.GetDB(), mods...)
}

// Environment pointed to by the foreign key.
func (o *Phendesc) Environment(exec boil.Executor, mods ...qm.QueryMod) environmentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("environment_id=$1", o.EnvironmentID),
	}

	queryMods = append(queryMods, mods...)

	query := Environments(exec, queryMods...)
	queries.SetFrom(query.Query, "\"environment\"")

	return query
}

// GenotypeG pointed to by the foreign key.
func (o *Phendesc) GenotypeG(mods ...qm.QueryMod) genotypeQuery {
	return o.Genotype(boil.GetDB(), mods...)
}

// Genotype pointed to by the foreign key.
func (o *Phendesc) Genotype(exec boil.Executor, mods ...qm.QueryMod) genotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype_id=$1", o.GenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Genotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"genotype\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Phendesc) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Phendesc) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// PubG pointed to by the foreign key.
func (o *Phendesc) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *Phendesc) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// LoadEnvironment allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phendescL) LoadEnvironment(e boil.Executor, singular bool, maybePhendesc interface{}) error {
	var slice []*Phendesc
	var object *Phendesc

	count := 1
	if singular {
		object = maybePhendesc.(*Phendesc)
	} else {
		slice = *maybePhendesc.(*PhendescSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phendescR{}
		args[0] = object.EnvironmentID
	} else {
		for i, obj := range slice {
			obj.R = &phendescR{}
			args[i] = obj.EnvironmentID
		}
	}

	query := fmt.Sprintf(
		"select * from \"environment\" where \"environment_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Environment")
	}
	defer results.Close()

	var resultSlice []*Environment
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Environment")
	}

	if len(phendescAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Environment = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.EnvironmentID == foreign.EnvironmentID {
				local.R.Environment = foreign
				break
			}
		}
	}

	return nil
}

// LoadGenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phendescL) LoadGenotype(e boil.Executor, singular bool, maybePhendesc interface{}) error {
	var slice []*Phendesc
	var object *Phendesc

	count := 1
	if singular {
		object = maybePhendesc.(*Phendesc)
	} else {
		slice = *maybePhendesc.(*PhendescSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phendescR{}
		args[0] = object.GenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &phendescR{}
			args[i] = obj.GenotypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"genotype\" where \"genotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Genotype")
	}
	defer results.Close()

	var resultSlice []*Genotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Genotype")
	}

	if len(phendescAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Genotype = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GenotypeID == foreign.GenotypeID {
				local.R.Genotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phendescL) LoadType(e boil.Executor, singular bool, maybePhendesc interface{}) error {
	var slice []*Phendesc
	var object *Phendesc

	count := 1
	if singular {
		object = maybePhendesc.(*Phendesc)
	} else {
		slice = *maybePhendesc.(*PhendescSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phendescR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &phendescR{}
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

	if len(phendescAfterSelectHooks) != 0 {
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

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phendescL) LoadPub(e boil.Executor, singular bool, maybePhendesc interface{}) error {
	var slice []*Phendesc
	var object *Phendesc

	count := 1
	if singular {
		object = maybePhendesc.(*Phendesc)
	} else {
		slice = *maybePhendesc.(*PhendescSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phendescR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &phendescR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"pub\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Pub")
	}
	defer results.Close()

	var resultSlice []*Pub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Pub")
	}

	if len(phendescAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Pub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.Pub = foreign
				break
			}
		}
	}

	return nil
}

// SetEnvironment of the phendesc to the related item.
// Sets o.R.Environment to related.
// Adds o to related.R.Phendesc.
func (o *Phendesc) SetEnvironment(exec boil.Executor, insert bool, related *Environment) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phendesc\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"environment_id"}),
		strmangle.WhereClause("\"", "\"", 2, phendescPrimaryKeyColumns),
	)
	values := []interface{}{related.EnvironmentID, o.PhendescID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.EnvironmentID = related.EnvironmentID

	if o.R == nil {
		o.R = &phendescR{
			Environment: related,
		}
	} else {
		o.R.Environment = related
	}

	if related.R == nil {
		related.R = &environmentR{
			Phendesc: o,
		}
	} else {
		related.R.Phendesc = o
	}

	return nil
}

// SetGenotype of the phendesc to the related item.
// Sets o.R.Genotype to related.
// Adds o to related.R.Phendesc.
func (o *Phendesc) SetGenotype(exec boil.Executor, insert bool, related *Genotype) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phendesc\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"genotype_id"}),
		strmangle.WhereClause("\"", "\"", 2, phendescPrimaryKeyColumns),
	)
	values := []interface{}{related.GenotypeID, o.PhendescID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GenotypeID = related.GenotypeID

	if o.R == nil {
		o.R = &phendescR{
			Genotype: related,
		}
	} else {
		o.R.Genotype = related
	}

	if related.R == nil {
		related.R = &genotypeR{
			Phendesc: o,
		}
	} else {
		related.R.Phendesc = o
	}

	return nil
}

// SetType of the phendesc to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypePhendesc.
func (o *Phendesc) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phendesc\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, phendescPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.PhendescID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &phendescR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypePhendesc: o,
		}
	} else {
		related.R.TypePhendesc = o
	}

	return nil
}

// SetPub of the phendesc to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.Phendesc.
func (o *Phendesc) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phendesc\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, phendescPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.PhendescID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &phendescR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			Phendesc: o,
		}
	} else {
		related.R.Phendesc = o
	}

	return nil
}

// PhendescsG retrieves all records.
func PhendescsG(mods ...qm.QueryMod) phendescQuery {
	return Phendescs(boil.GetDB(), mods...)
}

// Phendescs retrieves all the records using an executor.
func Phendescs(exec boil.Executor, mods ...qm.QueryMod) phendescQuery {
	mods = append(mods, qm.From("\"phendesc\""))
	return phendescQuery{NewQuery(exec, mods...)}
}

// FindPhendescG retrieves a single record by ID.
func FindPhendescG(phendescID int, selectCols ...string) (*Phendesc, error) {
	return FindPhendesc(boil.GetDB(), phendescID, selectCols...)
}

// FindPhendescGP retrieves a single record by ID, and panics on error.
func FindPhendescGP(phendescID int, selectCols ...string) *Phendesc {
	retobj, err := FindPhendesc(boil.GetDB(), phendescID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindPhendesc retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPhendesc(exec boil.Executor, phendescID int, selectCols ...string) (*Phendesc, error) {
	phendescObj := &Phendesc{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"phendesc\" where \"phendesc_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, phendescID)

	err := q.Bind(phendescObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from phendesc")
	}

	return phendescObj, nil
}

// FindPhendescP retrieves a single record by ID with an executor, and panics on error.
func FindPhendescP(exec boil.Executor, phendescID int, selectCols ...string) *Phendesc {
	retobj, err := FindPhendesc(exec, phendescID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Phendesc) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Phendesc) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Phendesc) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Phendesc) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no phendesc provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(phendescColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	phendescInsertCacheMut.RLock()
	cache, cached := phendescInsertCache[key]
	phendescInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			phendescColumns,
			phendescColumnsWithDefault,
			phendescColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(phendescType, phendescMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(phendescType, phendescMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"phendesc\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into phendesc")
	}

	if !cached {
		phendescInsertCacheMut.Lock()
		phendescInsertCache[key] = cache
		phendescInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Phendesc record. See Update for
// whitelist behavior description.
func (o *Phendesc) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Phendesc record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Phendesc) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Phendesc, and panics on error.
// See Update for whitelist behavior description.
func (o *Phendesc) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Phendesc.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Phendesc) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	phendescUpdateCacheMut.RLock()
	cache, cached := phendescUpdateCache[key]
	phendescUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(phendescColumns, phendescPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update phendesc, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"phendesc\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, phendescPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(phendescType, phendescMapping, append(wl, phendescPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update phendesc row")
	}

	if !cached {
		phendescUpdateCacheMut.Lock()
		phendescUpdateCache[key] = cache
		phendescUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q phendescQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q phendescQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for phendesc")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PhendescSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o PhendescSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o PhendescSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PhendescSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phendescPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"phendesc\" SET %s WHERE (\"phendesc_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(phendescPrimaryKeyColumns), len(colNames)+1, len(phendescPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in phendesc slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Phendesc) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Phendesc) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Phendesc) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Phendesc) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no phendesc provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(phendescColumnsWithDefault, o)

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

	phendescUpsertCacheMut.RLock()
	cache, cached := phendescUpsertCache[key]
	phendescUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			phendescColumns,
			phendescColumnsWithDefault,
			phendescColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			phendescColumns,
			phendescPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert phendesc, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(phendescPrimaryKeyColumns))
			copy(conflict, phendescPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"phendesc\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(phendescType, phendescMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(phendescType, phendescMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for phendesc")
	}

	if !cached {
		phendescUpsertCacheMut.Lock()
		phendescUpsertCache[key] = cache
		phendescUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Phendesc record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Phendesc) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Phendesc record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Phendesc) DeleteG() error {
	if o == nil {
		return errors.New("models: no Phendesc provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Phendesc record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Phendesc) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Phendesc record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Phendesc) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Phendesc provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), phendescPrimaryKeyMapping)
	sql := "DELETE FROM \"phendesc\" WHERE \"phendesc_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from phendesc")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q phendescQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q phendescQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no phendescQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from phendesc")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o PhendescSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o PhendescSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Phendesc slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o PhendescSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PhendescSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Phendesc slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(phendescBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phendescPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"phendesc\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, phendescPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(phendescPrimaryKeyColumns), 1, len(phendescPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from phendesc slice")
	}

	if len(phendescAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Phendesc) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Phendesc) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Phendesc) ReloadG() error {
	if o == nil {
		return errors.New("models: no Phendesc provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Phendesc) Reload(exec boil.Executor) error {
	ret, err := FindPhendesc(exec, o.PhendescID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PhendescSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PhendescSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhendescSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty PhendescSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhendescSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	phendescs := PhendescSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phendescPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"phendesc\".* FROM \"phendesc\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, phendescPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(phendescPrimaryKeyColumns), 1, len(phendescPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&phendescs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PhendescSlice")
	}

	*o = phendescs

	return nil
}

// PhendescExists checks if the Phendesc row exists.
func PhendescExists(exec boil.Executor, phendescID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"phendesc\" where \"phendesc_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, phendescID)
	}

	row := exec.QueryRow(sql, phendescID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if phendesc exists")
	}

	return exists, nil
}

// PhendescExistsG checks if the Phendesc row exists.
func PhendescExistsG(phendescID int) (bool, error) {
	return PhendescExists(boil.GetDB(), phendescID)
}

// PhendescExistsGP checks if the Phendesc row exists. Panics on error.
func PhendescExistsGP(phendescID int) bool {
	e, err := PhendescExists(boil.GetDB(), phendescID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// PhendescExistsP checks if the Phendesc row exists. Panics on error.
func PhendescExistsP(exec boil.Executor, phendescID int) bool {
	e, err := PhendescExists(exec, phendescID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
