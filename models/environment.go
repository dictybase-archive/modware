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

// Environment is an object representing the database table.
type Environment struct {
	EnvironmentID int         `boil:"environment_id" json:"environment_id" toml:"environment_id" yaml:"environment_id"`
	Uniquename    string      `boil:"uniquename" json:"uniquename" toml:"uniquename" yaml:"uniquename"`
	Description   null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`

	R *environmentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L environmentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// environmentR is where relationships are stored.
type environmentR struct {
	EnvironmentCvterm               *EnvironmentCvterm
	Phendesc                        *Phendesc
	Phenstatement                   *Phenstatement
	Environment1PhenotypeComparison *PhenotypeComparison
	Environment2PhenotypeComparison *PhenotypeComparison
}

// environmentL is where Load methods for each relationship are stored.
type environmentL struct{}

var (
	environmentColumns               = []string{"environment_id", "uniquename", "description"}
	environmentColumnsWithoutDefault = []string{"uniquename", "description"}
	environmentColumnsWithDefault    = []string{"environment_id"}
	environmentPrimaryKeyColumns     = []string{"environment_id"}
)

type (
	// EnvironmentSlice is an alias for a slice of pointers to Environment.
	// This should generally be used opposed to []Environment.
	EnvironmentSlice []*Environment
	// EnvironmentHook is the signature for custom Environment hook methods
	EnvironmentHook func(boil.Executor, *Environment) error

	environmentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	environmentType                 = reflect.TypeOf(&Environment{})
	environmentMapping              = queries.MakeStructMapping(environmentType)
	environmentPrimaryKeyMapping, _ = queries.BindMapping(environmentType, environmentMapping, environmentPrimaryKeyColumns)
	environmentInsertCacheMut       sync.RWMutex
	environmentInsertCache          = make(map[string]insertCache)
	environmentUpdateCacheMut       sync.RWMutex
	environmentUpdateCache          = make(map[string]updateCache)
	environmentUpsertCacheMut       sync.RWMutex
	environmentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var environmentBeforeInsertHooks []EnvironmentHook
var environmentBeforeUpdateHooks []EnvironmentHook
var environmentBeforeDeleteHooks []EnvironmentHook
var environmentBeforeUpsertHooks []EnvironmentHook

var environmentAfterInsertHooks []EnvironmentHook
var environmentAfterSelectHooks []EnvironmentHook
var environmentAfterUpdateHooks []EnvironmentHook
var environmentAfterDeleteHooks []EnvironmentHook
var environmentAfterUpsertHooks []EnvironmentHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Environment) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Environment) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Environment) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Environment) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Environment) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Environment) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Environment) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Environment) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Environment) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range environmentAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEnvironmentHook registers your hook function for all future operations.
func AddEnvironmentHook(hookPoint boil.HookPoint, environmentHook EnvironmentHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		environmentBeforeInsertHooks = append(environmentBeforeInsertHooks, environmentHook)
	case boil.BeforeUpdateHook:
		environmentBeforeUpdateHooks = append(environmentBeforeUpdateHooks, environmentHook)
	case boil.BeforeDeleteHook:
		environmentBeforeDeleteHooks = append(environmentBeforeDeleteHooks, environmentHook)
	case boil.BeforeUpsertHook:
		environmentBeforeUpsertHooks = append(environmentBeforeUpsertHooks, environmentHook)
	case boil.AfterInsertHook:
		environmentAfterInsertHooks = append(environmentAfterInsertHooks, environmentHook)
	case boil.AfterSelectHook:
		environmentAfterSelectHooks = append(environmentAfterSelectHooks, environmentHook)
	case boil.AfterUpdateHook:
		environmentAfterUpdateHooks = append(environmentAfterUpdateHooks, environmentHook)
	case boil.AfterDeleteHook:
		environmentAfterDeleteHooks = append(environmentAfterDeleteHooks, environmentHook)
	case boil.AfterUpsertHook:
		environmentAfterUpsertHooks = append(environmentAfterUpsertHooks, environmentHook)
	}
}

// OneP returns a single environment record from the query, and panics on error.
func (q environmentQuery) OneP() *Environment {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single environment record from the query.
func (q environmentQuery) One() (*Environment, error) {
	o := &Environment{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for environment")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Environment records from the query, and panics on error.
func (q environmentQuery) AllP() EnvironmentSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Environment records from the query.
func (q environmentQuery) All() (EnvironmentSlice, error) {
	var o EnvironmentSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Environment slice")
	}

	if len(environmentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Environment records in the query, and panics on error.
func (q environmentQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Environment records in the query.
func (q environmentQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count environment rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q environmentQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q environmentQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if environment exists")
	}

	return count > 0, nil
}

// EnvironmentCvtermG pointed to by the foreign key.
func (o *Environment) EnvironmentCvtermG(mods ...qm.QueryMod) environmentCvtermQuery {
	return o.EnvironmentCvterm(boil.GetDB(), mods...)
}

// EnvironmentCvterm pointed to by the foreign key.
func (o *Environment) EnvironmentCvterm(exec boil.Executor, mods ...qm.QueryMod) environmentCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("environment_id=$1", o.EnvironmentID),
	}

	queryMods = append(queryMods, mods...)

	query := EnvironmentCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"environment_cvterm\"")

	return query
}

// PhendescG pointed to by the foreign key.
func (o *Environment) PhendescG(mods ...qm.QueryMod) phendescQuery {
	return o.Phendesc(boil.GetDB(), mods...)
}

// Phendesc pointed to by the foreign key.
func (o *Environment) Phendesc(exec boil.Executor, mods ...qm.QueryMod) phendescQuery {
	queryMods := []qm.QueryMod{
		qm.Where("environment_id=$1", o.EnvironmentID),
	}

	queryMods = append(queryMods, mods...)

	query := Phendescs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phendesc\"")

	return query
}

// PhenstatementG pointed to by the foreign key.
func (o *Environment) PhenstatementG(mods ...qm.QueryMod) phenstatementQuery {
	return o.Phenstatement(boil.GetDB(), mods...)
}

// Phenstatement pointed to by the foreign key.
func (o *Environment) Phenstatement(exec boil.Executor, mods ...qm.QueryMod) phenstatementQuery {
	queryMods := []qm.QueryMod{
		qm.Where("environment_id=$1", o.EnvironmentID),
	}

	queryMods = append(queryMods, mods...)

	query := Phenstatements(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenstatement\"")

	return query
}

// Environment1PhenotypeComparisonG pointed to by the foreign key.
func (o *Environment) Environment1PhenotypeComparisonG(mods ...qm.QueryMod) phenotypeComparisonQuery {
	return o.Environment1PhenotypeComparison(boil.GetDB(), mods...)
}

// Environment1PhenotypeComparison pointed to by the foreign key.
func (o *Environment) Environment1PhenotypeComparison(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonQuery {
	queryMods := []qm.QueryMod{
		qm.Where("environment1_id=$1", o.EnvironmentID),
	}

	queryMods = append(queryMods, mods...)

	query := PhenotypeComparisons(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_comparison\"")

	return query
}

// Environment2PhenotypeComparisonG pointed to by the foreign key.
func (o *Environment) Environment2PhenotypeComparisonG(mods ...qm.QueryMod) phenotypeComparisonQuery {
	return o.Environment2PhenotypeComparison(boil.GetDB(), mods...)
}

// Environment2PhenotypeComparison pointed to by the foreign key.
func (o *Environment) Environment2PhenotypeComparison(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonQuery {
	queryMods := []qm.QueryMod{
		qm.Where("environment2_id=$1", o.EnvironmentID),
	}

	queryMods = append(queryMods, mods...)

	query := PhenotypeComparisons(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_comparison\"")

	return query
}

// LoadEnvironmentCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (environmentL) LoadEnvironmentCvterm(e boil.Executor, singular bool, maybeEnvironment interface{}) error {
	var slice []*Environment
	var object *Environment

	count := 1
	if singular {
		object = maybeEnvironment.(*Environment)
	} else {
		slice = *maybeEnvironment.(*EnvironmentSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &environmentR{}
		args[0] = object.EnvironmentID
	} else {
		for i, obj := range slice {
			obj.R = &environmentR{}
			args[i] = obj.EnvironmentID
		}
	}

	query := fmt.Sprintf(
		"select * from \"environment_cvterm\" where \"environment_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load EnvironmentCvterm")
	}
	defer results.Close()

	var resultSlice []*EnvironmentCvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice EnvironmentCvterm")
	}

	if len(environmentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.EnvironmentCvterm = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.EnvironmentID == foreign.EnvironmentID {
				local.R.EnvironmentCvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhendesc allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (environmentL) LoadPhendesc(e boil.Executor, singular bool, maybeEnvironment interface{}) error {
	var slice []*Environment
	var object *Environment

	count := 1
	if singular {
		object = maybeEnvironment.(*Environment)
	} else {
		slice = *maybeEnvironment.(*EnvironmentSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &environmentR{}
		args[0] = object.EnvironmentID
	} else {
		for i, obj := range slice {
			obj.R = &environmentR{}
			args[i] = obj.EnvironmentID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phendesc\" where \"environment_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Phendesc")
	}
	defer results.Close()

	var resultSlice []*Phendesc
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Phendesc")
	}

	if len(environmentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Phendesc = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.EnvironmentID == foreign.EnvironmentID {
				local.R.Phendesc = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenstatement allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (environmentL) LoadPhenstatement(e boil.Executor, singular bool, maybeEnvironment interface{}) error {
	var slice []*Environment
	var object *Environment

	count := 1
	if singular {
		object = maybeEnvironment.(*Environment)
	} else {
		slice = *maybeEnvironment.(*EnvironmentSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &environmentR{}
		args[0] = object.EnvironmentID
	} else {
		for i, obj := range slice {
			obj.R = &environmentR{}
			args[i] = obj.EnvironmentID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenstatement\" where \"environment_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Phenstatement")
	}
	defer results.Close()

	var resultSlice []*Phenstatement
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Phenstatement")
	}

	if len(environmentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Phenstatement = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.EnvironmentID == foreign.EnvironmentID {
				local.R.Phenstatement = foreign
				break
			}
		}
	}

	return nil
}

// LoadEnvironment1PhenotypeComparison allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (environmentL) LoadEnvironment1PhenotypeComparison(e boil.Executor, singular bool, maybeEnvironment interface{}) error {
	var slice []*Environment
	var object *Environment

	count := 1
	if singular {
		object = maybeEnvironment.(*Environment)
	} else {
		slice = *maybeEnvironment.(*EnvironmentSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &environmentR{}
		args[0] = object.EnvironmentID
	} else {
		for i, obj := range slice {
			obj.R = &environmentR{}
			args[i] = obj.EnvironmentID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_comparison\" where \"environment1_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PhenotypeComparison")
	}
	defer results.Close()

	var resultSlice []*PhenotypeComparison
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PhenotypeComparison")
	}

	if len(environmentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Environment1PhenotypeComparison = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.EnvironmentID == foreign.Environment1ID {
				local.R.Environment1PhenotypeComparison = foreign
				break
			}
		}
	}

	return nil
}

// LoadEnvironment2PhenotypeComparison allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (environmentL) LoadEnvironment2PhenotypeComparison(e boil.Executor, singular bool, maybeEnvironment interface{}) error {
	var slice []*Environment
	var object *Environment

	count := 1
	if singular {
		object = maybeEnvironment.(*Environment)
	} else {
		slice = *maybeEnvironment.(*EnvironmentSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &environmentR{}
		args[0] = object.EnvironmentID
	} else {
		for i, obj := range slice {
			obj.R = &environmentR{}
			args[i] = obj.EnvironmentID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_comparison\" where \"environment2_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PhenotypeComparison")
	}
	defer results.Close()

	var resultSlice []*PhenotypeComparison
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PhenotypeComparison")
	}

	if len(environmentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Environment2PhenotypeComparison = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.EnvironmentID == foreign.Environment2ID {
				local.R.Environment2PhenotypeComparison = foreign
				break
			}
		}
	}

	return nil
}

// SetEnvironmentCvterm of the environment to the related item.
// Sets o.R.EnvironmentCvterm to related.
// Adds o to related.R.Environment.
func (o *Environment) SetEnvironmentCvterm(exec boil.Executor, insert bool, related *EnvironmentCvterm) error {
	var err error

	if insert {
		related.EnvironmentID = o.EnvironmentID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"environment_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"environment_id"}),
			strmangle.WhereClause("\"", "\"", 2, environmentCvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.EnvironmentID, related.EnvironmentCvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.EnvironmentID = o.EnvironmentID

	}

	if o.R == nil {
		o.R = &environmentR{
			EnvironmentCvterm: related,
		}
	} else {
		o.R.EnvironmentCvterm = related
	}

	if related.R == nil {
		related.R = &environmentCvtermR{
			Environment: o,
		}
	} else {
		related.R.Environment = o
	}
	return nil
}

// SetPhendesc of the environment to the related item.
// Sets o.R.Phendesc to related.
// Adds o to related.R.Environment.
func (o *Environment) SetPhendesc(exec boil.Executor, insert bool, related *Phendesc) error {
	var err error

	if insert {
		related.EnvironmentID = o.EnvironmentID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phendesc\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"environment_id"}),
			strmangle.WhereClause("\"", "\"", 2, phendescPrimaryKeyColumns),
		)
		values := []interface{}{o.EnvironmentID, related.PhendescID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.EnvironmentID = o.EnvironmentID

	}

	if o.R == nil {
		o.R = &environmentR{
			Phendesc: related,
		}
	} else {
		o.R.Phendesc = related
	}

	if related.R == nil {
		related.R = &phendescR{
			Environment: o,
		}
	} else {
		related.R.Environment = o
	}
	return nil
}

// SetPhenstatement of the environment to the related item.
// Sets o.R.Phenstatement to related.
// Adds o to related.R.Environment.
func (o *Environment) SetPhenstatement(exec boil.Executor, insert bool, related *Phenstatement) error {
	var err error

	if insert {
		related.EnvironmentID = o.EnvironmentID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenstatement\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"environment_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenstatementPrimaryKeyColumns),
		)
		values := []interface{}{o.EnvironmentID, related.PhenstatementID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.EnvironmentID = o.EnvironmentID

	}

	if o.R == nil {
		o.R = &environmentR{
			Phenstatement: related,
		}
	} else {
		o.R.Phenstatement = related
	}

	if related.R == nil {
		related.R = &phenstatementR{
			Environment: o,
		}
	} else {
		related.R.Environment = o
	}
	return nil
}

// SetEnvironment1PhenotypeComparison of the environment to the related item.
// Sets o.R.Environment1PhenotypeComparison to related.
// Adds o to related.R.Environment1.
func (o *Environment) SetEnvironment1PhenotypeComparison(exec boil.Executor, insert bool, related *PhenotypeComparison) error {
	var err error

	if insert {
		related.Environment1ID = o.EnvironmentID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"environment1_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
		)
		values := []interface{}{o.EnvironmentID, related.PhenotypeComparisonID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.Environment1ID = o.EnvironmentID

	}

	if o.R == nil {
		o.R = &environmentR{
			Environment1PhenotypeComparison: related,
		}
	} else {
		o.R.Environment1PhenotypeComparison = related
	}

	if related.R == nil {
		related.R = &phenotypeComparisonR{
			Environment1: o,
		}
	} else {
		related.R.Environment1 = o
	}
	return nil
}

// SetEnvironment2PhenotypeComparison of the environment to the related item.
// Sets o.R.Environment2PhenotypeComparison to related.
// Adds o to related.R.Environment2.
func (o *Environment) SetEnvironment2PhenotypeComparison(exec boil.Executor, insert bool, related *PhenotypeComparison) error {
	var err error

	if insert {
		related.Environment2ID = o.EnvironmentID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"environment2_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
		)
		values := []interface{}{o.EnvironmentID, related.PhenotypeComparisonID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.Environment2ID = o.EnvironmentID

	}

	if o.R == nil {
		o.R = &environmentR{
			Environment2PhenotypeComparison: related,
		}
	} else {
		o.R.Environment2PhenotypeComparison = related
	}

	if related.R == nil {
		related.R = &phenotypeComparisonR{
			Environment2: o,
		}
	} else {
		related.R.Environment2 = o
	}
	return nil
}

// EnvironmentsG retrieves all records.
func EnvironmentsG(mods ...qm.QueryMod) environmentQuery {
	return Environments(boil.GetDB(), mods...)
}

// Environments retrieves all the records using an executor.
func Environments(exec boil.Executor, mods ...qm.QueryMod) environmentQuery {
	mods = append(mods, qm.From("\"environment\""))
	return environmentQuery{NewQuery(exec, mods...)}
}

// FindEnvironmentG retrieves a single record by ID.
func FindEnvironmentG(environmentID int, selectCols ...string) (*Environment, error) {
	return FindEnvironment(boil.GetDB(), environmentID, selectCols...)
}

// FindEnvironmentGP retrieves a single record by ID, and panics on error.
func FindEnvironmentGP(environmentID int, selectCols ...string) *Environment {
	retobj, err := FindEnvironment(boil.GetDB(), environmentID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindEnvironment retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEnvironment(exec boil.Executor, environmentID int, selectCols ...string) (*Environment, error) {
	environmentObj := &Environment{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"environment\" where \"environment_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, environmentID)

	err := q.Bind(environmentObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from environment")
	}

	return environmentObj, nil
}

// FindEnvironmentP retrieves a single record by ID with an executor, and panics on error.
func FindEnvironmentP(exec boil.Executor, environmentID int, selectCols ...string) *Environment {
	retobj, err := FindEnvironment(exec, environmentID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Environment) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Environment) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Environment) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Environment) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no environment provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(environmentColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	environmentInsertCacheMut.RLock()
	cache, cached := environmentInsertCache[key]
	environmentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			environmentColumns,
			environmentColumnsWithDefault,
			environmentColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(environmentType, environmentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(environmentType, environmentMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"environment\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into environment")
	}

	if !cached {
		environmentInsertCacheMut.Lock()
		environmentInsertCache[key] = cache
		environmentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Environment record. See Update for
// whitelist behavior description.
func (o *Environment) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Environment record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Environment) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Environment, and panics on error.
// See Update for whitelist behavior description.
func (o *Environment) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Environment.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Environment) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	environmentUpdateCacheMut.RLock()
	cache, cached := environmentUpdateCache[key]
	environmentUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(environmentColumns, environmentPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update environment, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"environment\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, environmentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(environmentType, environmentMapping, append(wl, environmentPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update environment row")
	}

	if !cached {
		environmentUpdateCacheMut.Lock()
		environmentUpdateCache[key] = cache
		environmentUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q environmentQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q environmentQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for environment")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o EnvironmentSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o EnvironmentSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o EnvironmentSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EnvironmentSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), environmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"environment\" SET %s WHERE (\"environment_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(environmentPrimaryKeyColumns), len(colNames)+1, len(environmentPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in environment slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Environment) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Environment) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Environment) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Environment) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no environment provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(environmentColumnsWithDefault, o)

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

	environmentUpsertCacheMut.RLock()
	cache, cached := environmentUpsertCache[key]
	environmentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			environmentColumns,
			environmentColumnsWithDefault,
			environmentColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			environmentColumns,
			environmentPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert environment, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(environmentPrimaryKeyColumns))
			copy(conflict, environmentPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"environment\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(environmentType, environmentMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(environmentType, environmentMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for environment")
	}

	if !cached {
		environmentUpsertCacheMut.Lock()
		environmentUpsertCache[key] = cache
		environmentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Environment record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Environment) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Environment record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Environment) DeleteG() error {
	if o == nil {
		return errors.New("models: no Environment provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Environment record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Environment) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Environment record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Environment) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Environment provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), environmentPrimaryKeyMapping)
	sql := "DELETE FROM \"environment\" WHERE \"environment_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from environment")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q environmentQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q environmentQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no environmentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from environment")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o EnvironmentSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o EnvironmentSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Environment slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o EnvironmentSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EnvironmentSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Environment slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(environmentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), environmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"environment\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, environmentPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(environmentPrimaryKeyColumns), 1, len(environmentPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from environment slice")
	}

	if len(environmentAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Environment) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Environment) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Environment) ReloadG() error {
	if o == nil {
		return errors.New("models: no Environment provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Environment) Reload(exec boil.Executor) error {
	ret, err := FindEnvironment(exec, o.EnvironmentID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *EnvironmentSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *EnvironmentSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EnvironmentSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty EnvironmentSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EnvironmentSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	environments := EnvironmentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), environmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"environment\".* FROM \"environment\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, environmentPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(environmentPrimaryKeyColumns), 1, len(environmentPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&environments)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EnvironmentSlice")
	}

	*o = environments

	return nil
}

// EnvironmentExists checks if the Environment row exists.
func EnvironmentExists(exec boil.Executor, environmentID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"environment\" where \"environment_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, environmentID)
	}

	row := exec.QueryRow(sql, environmentID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if environment exists")
	}

	return exists, nil
}

// EnvironmentExistsG checks if the Environment row exists.
func EnvironmentExistsG(environmentID int) (bool, error) {
	return EnvironmentExists(boil.GetDB(), environmentID)
}

// EnvironmentExistsGP checks if the Environment row exists. Panics on error.
func EnvironmentExistsGP(environmentID int) bool {
	e, err := EnvironmentExists(boil.GetDB(), environmentID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// EnvironmentExistsP checks if the Environment row exists. Panics on error.
func EnvironmentExistsP(exec boil.Executor, environmentID int) bool {
	e, err := EnvironmentExists(exec, environmentID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
