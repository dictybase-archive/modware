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

// PhenotypeComparison is an object representing the database table.
type PhenotypeComparison struct {
	PhenotypeComparisonID int      `boil:"phenotype_comparison_id" json:"phenotype_comparison_id" toml:"phenotype_comparison_id" yaml:"phenotype_comparison_id"`
	Genotype1ID           int      `boil:"genotype1_id" json:"genotype1_id" toml:"genotype1_id" yaml:"genotype1_id"`
	Environment1ID        int      `boil:"environment1_id" json:"environment1_id" toml:"environment1_id" yaml:"environment1_id"`
	Genotype2ID           int      `boil:"genotype2_id" json:"genotype2_id" toml:"genotype2_id" yaml:"genotype2_id"`
	Environment2ID        int      `boil:"environment2_id" json:"environment2_id" toml:"environment2_id" yaml:"environment2_id"`
	Phenotype1ID          int      `boil:"phenotype1_id" json:"phenotype1_id" toml:"phenotype1_id" yaml:"phenotype1_id"`
	Phenotype2ID          null.Int `boil:"phenotype2_id" json:"phenotype2_id,omitempty" toml:"phenotype2_id" yaml:"phenotype2_id,omitempty"`
	PubID                 int      `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`
	OrganismID            int      `boil:"organism_id" json:"organism_id" toml:"organism_id" yaml:"organism_id"`

	R *phenotypeComparisonR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L phenotypeComparisonL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// phenotypeComparisonR is where relationships are stored.
type phenotypeComparisonR struct {
	Environment1              *Environment
	Environment2              *Environment
	Genotype1                 *Genotype
	Genotype2                 *Genotype
	Organism                  *Organism
	Phenotype1                *Phenotype
	Phenotype2                *Phenotype
	Pub                       *Pub
	PhenotypeComparisonCvterm *PhenotypeComparisonCvterm
}

// phenotypeComparisonL is where Load methods for each relationship are stored.
type phenotypeComparisonL struct{}

var (
	phenotypeComparisonColumns               = []string{"phenotype_comparison_id", "genotype1_id", "environment1_id", "genotype2_id", "environment2_id", "phenotype1_id", "phenotype2_id", "pub_id", "organism_id"}
	phenotypeComparisonColumnsWithoutDefault = []string{"genotype1_id", "environment1_id", "genotype2_id", "environment2_id", "phenotype1_id", "phenotype2_id", "pub_id", "organism_id"}
	phenotypeComparisonColumnsWithDefault    = []string{"phenotype_comparison_id"}
	phenotypeComparisonPrimaryKeyColumns     = []string{"phenotype_comparison_id"}
)

type (
	// PhenotypeComparisonSlice is an alias for a slice of pointers to PhenotypeComparison.
	// This should generally be used opposed to []PhenotypeComparison.
	PhenotypeComparisonSlice []*PhenotypeComparison
	// PhenotypeComparisonHook is the signature for custom PhenotypeComparison hook methods
	PhenotypeComparisonHook func(boil.Executor, *PhenotypeComparison) error

	phenotypeComparisonQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	phenotypeComparisonType                 = reflect.TypeOf(&PhenotypeComparison{})
	phenotypeComparisonMapping              = queries.MakeStructMapping(phenotypeComparisonType)
	phenotypeComparisonPrimaryKeyMapping, _ = queries.BindMapping(phenotypeComparisonType, phenotypeComparisonMapping, phenotypeComparisonPrimaryKeyColumns)
	phenotypeComparisonInsertCacheMut       sync.RWMutex
	phenotypeComparisonInsertCache          = make(map[string]insertCache)
	phenotypeComparisonUpdateCacheMut       sync.RWMutex
	phenotypeComparisonUpdateCache          = make(map[string]updateCache)
	phenotypeComparisonUpsertCacheMut       sync.RWMutex
	phenotypeComparisonUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var phenotypeComparisonBeforeInsertHooks []PhenotypeComparisonHook
var phenotypeComparisonBeforeUpdateHooks []PhenotypeComparisonHook
var phenotypeComparisonBeforeDeleteHooks []PhenotypeComparisonHook
var phenotypeComparisonBeforeUpsertHooks []PhenotypeComparisonHook

var phenotypeComparisonAfterInsertHooks []PhenotypeComparisonHook
var phenotypeComparisonAfterSelectHooks []PhenotypeComparisonHook
var phenotypeComparisonAfterUpdateHooks []PhenotypeComparisonHook
var phenotypeComparisonAfterDeleteHooks []PhenotypeComparisonHook
var phenotypeComparisonAfterUpsertHooks []PhenotypeComparisonHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PhenotypeComparison) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PhenotypeComparison) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PhenotypeComparison) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PhenotypeComparison) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PhenotypeComparison) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PhenotypeComparison) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PhenotypeComparison) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PhenotypeComparison) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PhenotypeComparison) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range phenotypeComparisonAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPhenotypeComparisonHook registers your hook function for all future operations.
func AddPhenotypeComparisonHook(hookPoint boil.HookPoint, phenotypeComparisonHook PhenotypeComparisonHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		phenotypeComparisonBeforeInsertHooks = append(phenotypeComparisonBeforeInsertHooks, phenotypeComparisonHook)
	case boil.BeforeUpdateHook:
		phenotypeComparisonBeforeUpdateHooks = append(phenotypeComparisonBeforeUpdateHooks, phenotypeComparisonHook)
	case boil.BeforeDeleteHook:
		phenotypeComparisonBeforeDeleteHooks = append(phenotypeComparisonBeforeDeleteHooks, phenotypeComparisonHook)
	case boil.BeforeUpsertHook:
		phenotypeComparisonBeforeUpsertHooks = append(phenotypeComparisonBeforeUpsertHooks, phenotypeComparisonHook)
	case boil.AfterInsertHook:
		phenotypeComparisonAfterInsertHooks = append(phenotypeComparisonAfterInsertHooks, phenotypeComparisonHook)
	case boil.AfterSelectHook:
		phenotypeComparisonAfterSelectHooks = append(phenotypeComparisonAfterSelectHooks, phenotypeComparisonHook)
	case boil.AfterUpdateHook:
		phenotypeComparisonAfterUpdateHooks = append(phenotypeComparisonAfterUpdateHooks, phenotypeComparisonHook)
	case boil.AfterDeleteHook:
		phenotypeComparisonAfterDeleteHooks = append(phenotypeComparisonAfterDeleteHooks, phenotypeComparisonHook)
	case boil.AfterUpsertHook:
		phenotypeComparisonAfterUpsertHooks = append(phenotypeComparisonAfterUpsertHooks, phenotypeComparisonHook)
	}
}

// OneP returns a single phenotypeComparison record from the query, and panics on error.
func (q phenotypeComparisonQuery) OneP() *PhenotypeComparison {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single phenotypeComparison record from the query.
func (q phenotypeComparisonQuery) One() (*PhenotypeComparison, error) {
	o := &PhenotypeComparison{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for phenotype_comparison")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all PhenotypeComparison records from the query, and panics on error.
func (q phenotypeComparisonQuery) AllP() PhenotypeComparisonSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all PhenotypeComparison records from the query.
func (q phenotypeComparisonQuery) All() (PhenotypeComparisonSlice, error) {
	var o PhenotypeComparisonSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to PhenotypeComparison slice")
	}

	if len(phenotypeComparisonAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all PhenotypeComparison records in the query, and panics on error.
func (q phenotypeComparisonQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all PhenotypeComparison records in the query.
func (q phenotypeComparisonQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count phenotype_comparison rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q phenotypeComparisonQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q phenotypeComparisonQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if phenotype_comparison exists")
	}

	return count > 0, nil
}

// Environment1G pointed to by the foreign key.
func (o *PhenotypeComparison) Environment1G(mods ...qm.QueryMod) environmentQuery {
	return o.Environment1(boil.GetDB(), mods...)
}

// Environment1 pointed to by the foreign key.
func (o *PhenotypeComparison) Environment1(exec boil.Executor, mods ...qm.QueryMod) environmentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("environment_id=$1", o.Environment1ID),
	}

	queryMods = append(queryMods, mods...)

	query := Environments(exec, queryMods...)
	queries.SetFrom(query.Query, "\"environment\"")

	return query
}

// Environment2G pointed to by the foreign key.
func (o *PhenotypeComparison) Environment2G(mods ...qm.QueryMod) environmentQuery {
	return o.Environment2(boil.GetDB(), mods...)
}

// Environment2 pointed to by the foreign key.
func (o *PhenotypeComparison) Environment2(exec boil.Executor, mods ...qm.QueryMod) environmentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("environment_id=$1", o.Environment2ID),
	}

	queryMods = append(queryMods, mods...)

	query := Environments(exec, queryMods...)
	queries.SetFrom(query.Query, "\"environment\"")

	return query
}

// Genotype1G pointed to by the foreign key.
func (o *PhenotypeComparison) Genotype1G(mods ...qm.QueryMod) genotypeQuery {
	return o.Genotype1(boil.GetDB(), mods...)
}

// Genotype1 pointed to by the foreign key.
func (o *PhenotypeComparison) Genotype1(exec boil.Executor, mods ...qm.QueryMod) genotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype_id=$1", o.Genotype1ID),
	}

	queryMods = append(queryMods, mods...)

	query := Genotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"genotype\"")

	return query
}

// Genotype2G pointed to by the foreign key.
func (o *PhenotypeComparison) Genotype2G(mods ...qm.QueryMod) genotypeQuery {
	return o.Genotype2(boil.GetDB(), mods...)
}

// Genotype2 pointed to by the foreign key.
func (o *PhenotypeComparison) Genotype2(exec boil.Executor, mods ...qm.QueryMod) genotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype_id=$1", o.Genotype2ID),
	}

	queryMods = append(queryMods, mods...)

	query := Genotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"genotype\"")

	return query
}

// OrganismG pointed to by the foreign key.
func (o *PhenotypeComparison) OrganismG(mods ...qm.QueryMod) organismQuery {
	return o.Organism(boil.GetDB(), mods...)
}

// Organism pointed to by the foreign key.
func (o *PhenotypeComparison) Organism(exec boil.Executor, mods ...qm.QueryMod) organismQuery {
	queryMods := []qm.QueryMod{
		qm.Where("organism_id=$1", o.OrganismID),
	}

	queryMods = append(queryMods, mods...)

	query := Organisms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"organism\"")

	return query
}

// Phenotype1G pointed to by the foreign key.
func (o *PhenotypeComparison) Phenotype1G(mods ...qm.QueryMod) phenotypeQuery {
	return o.Phenotype1(boil.GetDB(), mods...)
}

// Phenotype1 pointed to by the foreign key.
func (o *PhenotypeComparison) Phenotype1(exec boil.Executor, mods ...qm.QueryMod) phenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("phenotype_id=$1", o.Phenotype1ID),
	}

	queryMods = append(queryMods, mods...)

	query := Phenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype\"")

	return query
}

// Phenotype2G pointed to by the foreign key.
func (o *PhenotypeComparison) Phenotype2G(mods ...qm.QueryMod) phenotypeQuery {
	return o.Phenotype2(boil.GetDB(), mods...)
}

// Phenotype2 pointed to by the foreign key.
func (o *PhenotypeComparison) Phenotype2(exec boil.Executor, mods ...qm.QueryMod) phenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("phenotype_id=$1", o.Phenotype2ID),
	}

	queryMods = append(queryMods, mods...)

	query := Phenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype\"")

	return query
}

// PubG pointed to by the foreign key.
func (o *PhenotypeComparison) PubG(mods ...qm.QueryMod) pubQuery {
	return o.Pub(boil.GetDB(), mods...)
}

// Pub pointed to by the foreign key.
func (o *PhenotypeComparison) Pub(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\"")

	return query
}

// PhenotypeComparisonCvtermG pointed to by the foreign key.
func (o *PhenotypeComparison) PhenotypeComparisonCvtermG(mods ...qm.QueryMod) phenotypeComparisonCvtermQuery {
	return o.PhenotypeComparisonCvterm(boil.GetDB(), mods...)
}

// PhenotypeComparisonCvterm pointed to by the foreign key.
func (o *PhenotypeComparison) PhenotypeComparisonCvterm(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("phenotype_comparison_id=$1", o.PhenotypeComparisonID),
	}

	queryMods = append(queryMods, mods...)

	query := PhenotypeComparisonCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_comparison_cvterm\"")

	return query
}

// LoadEnvironment1 allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeComparisonL) LoadEnvironment1(e boil.Executor, singular bool, maybePhenotypeComparison interface{}) error {
	var slice []*PhenotypeComparison
	var object *PhenotypeComparison

	count := 1
	if singular {
		object = maybePhenotypeComparison.(*PhenotypeComparison)
	} else {
		slice = *maybePhenotypeComparison.(*PhenotypeComparisonSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeComparisonR{}
		args[0] = object.Environment1ID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeComparisonR{}
			args[i] = obj.Environment1ID
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

	if len(phenotypeComparisonAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Environment1 = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.Environment1ID == foreign.EnvironmentID {
				local.R.Environment1 = foreign
				break
			}
		}
	}

	return nil
}

// LoadEnvironment2 allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeComparisonL) LoadEnvironment2(e boil.Executor, singular bool, maybePhenotypeComparison interface{}) error {
	var slice []*PhenotypeComparison
	var object *PhenotypeComparison

	count := 1
	if singular {
		object = maybePhenotypeComparison.(*PhenotypeComparison)
	} else {
		slice = *maybePhenotypeComparison.(*PhenotypeComparisonSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeComparisonR{}
		args[0] = object.Environment2ID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeComparisonR{}
			args[i] = obj.Environment2ID
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

	if len(phenotypeComparisonAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Environment2 = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.Environment2ID == foreign.EnvironmentID {
				local.R.Environment2 = foreign
				break
			}
		}
	}

	return nil
}

// LoadGenotype1 allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeComparisonL) LoadGenotype1(e boil.Executor, singular bool, maybePhenotypeComparison interface{}) error {
	var slice []*PhenotypeComparison
	var object *PhenotypeComparison

	count := 1
	if singular {
		object = maybePhenotypeComparison.(*PhenotypeComparison)
	} else {
		slice = *maybePhenotypeComparison.(*PhenotypeComparisonSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeComparisonR{}
		args[0] = object.Genotype1ID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeComparisonR{}
			args[i] = obj.Genotype1ID
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

	if len(phenotypeComparisonAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Genotype1 = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.Genotype1ID == foreign.GenotypeID {
				local.R.Genotype1 = foreign
				break
			}
		}
	}

	return nil
}

// LoadGenotype2 allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeComparisonL) LoadGenotype2(e boil.Executor, singular bool, maybePhenotypeComparison interface{}) error {
	var slice []*PhenotypeComparison
	var object *PhenotypeComparison

	count := 1
	if singular {
		object = maybePhenotypeComparison.(*PhenotypeComparison)
	} else {
		slice = *maybePhenotypeComparison.(*PhenotypeComparisonSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeComparisonR{}
		args[0] = object.Genotype2ID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeComparisonR{}
			args[i] = obj.Genotype2ID
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

	if len(phenotypeComparisonAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Genotype2 = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.Genotype2ID == foreign.GenotypeID {
				local.R.Genotype2 = foreign
				break
			}
		}
	}

	return nil
}

// LoadOrganism allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeComparisonL) LoadOrganism(e boil.Executor, singular bool, maybePhenotypeComparison interface{}) error {
	var slice []*PhenotypeComparison
	var object *PhenotypeComparison

	count := 1
	if singular {
		object = maybePhenotypeComparison.(*PhenotypeComparison)
	} else {
		slice = *maybePhenotypeComparison.(*PhenotypeComparisonSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeComparisonR{}
		args[0] = object.OrganismID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeComparisonR{}
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

	if len(phenotypeComparisonAfterSelectHooks) != 0 {
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

// LoadPhenotype1 allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeComparisonL) LoadPhenotype1(e boil.Executor, singular bool, maybePhenotypeComparison interface{}) error {
	var slice []*PhenotypeComparison
	var object *PhenotypeComparison

	count := 1
	if singular {
		object = maybePhenotypeComparison.(*PhenotypeComparison)
	} else {
		slice = *maybePhenotypeComparison.(*PhenotypeComparisonSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeComparisonR{}
		args[0] = object.Phenotype1ID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeComparisonR{}
			args[i] = obj.Phenotype1ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype\" where \"phenotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Phenotype")
	}
	defer results.Close()

	var resultSlice []*Phenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Phenotype")
	}

	if len(phenotypeComparisonAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Phenotype1 = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.Phenotype1ID == foreign.PhenotypeID {
				local.R.Phenotype1 = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenotype2 allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeComparisonL) LoadPhenotype2(e boil.Executor, singular bool, maybePhenotypeComparison interface{}) error {
	var slice []*PhenotypeComparison
	var object *PhenotypeComparison

	count := 1
	if singular {
		object = maybePhenotypeComparison.(*PhenotypeComparison)
	} else {
		slice = *maybePhenotypeComparison.(*PhenotypeComparisonSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeComparisonR{}
		args[0] = object.Phenotype2ID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeComparisonR{}
			args[i] = obj.Phenotype2ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype\" where \"phenotype_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Phenotype")
	}
	defer results.Close()

	var resultSlice []*Phenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Phenotype")
	}

	if len(phenotypeComparisonAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Phenotype2 = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.Phenotype2ID.Int == foreign.PhenotypeID {
				local.R.Phenotype2 = foreign
				break
			}
		}
	}

	return nil
}

// LoadPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeComparisonL) LoadPub(e boil.Executor, singular bool, maybePhenotypeComparison interface{}) error {
	var slice []*PhenotypeComparison
	var object *PhenotypeComparison

	count := 1
	if singular {
		object = maybePhenotypeComparison.(*PhenotypeComparison)
	} else {
		slice = *maybePhenotypeComparison.(*PhenotypeComparisonSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeComparisonR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeComparisonR{}
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

	if len(phenotypeComparisonAfterSelectHooks) != 0 {
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

// LoadPhenotypeComparisonCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (phenotypeComparisonL) LoadPhenotypeComparisonCvterm(e boil.Executor, singular bool, maybePhenotypeComparison interface{}) error {
	var slice []*PhenotypeComparison
	var object *PhenotypeComparison

	count := 1
	if singular {
		object = maybePhenotypeComparison.(*PhenotypeComparison)
	} else {
		slice = *maybePhenotypeComparison.(*PhenotypeComparisonSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &phenotypeComparisonR{}
		args[0] = object.PhenotypeComparisonID
	} else {
		for i, obj := range slice {
			obj.R = &phenotypeComparisonR{}
			args[i] = obj.PhenotypeComparisonID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_comparison_cvterm\" where \"phenotype_comparison_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PhenotypeComparisonCvterm")
	}
	defer results.Close()

	var resultSlice []*PhenotypeComparisonCvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PhenotypeComparisonCvterm")
	}

	if len(phenotypeComparisonAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.PhenotypeComparisonCvterm = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PhenotypeComparisonID == foreign.PhenotypeComparisonID {
				local.R.PhenotypeComparisonCvterm = foreign
				break
			}
		}
	}

	return nil
}

// SetEnvironment1 of the phenotype_comparison to the related item.
// Sets o.R.Environment1 to related.
// Adds o to related.R.Environment1PhenotypeComparison.
func (o *PhenotypeComparison) SetEnvironment1(exec boil.Executor, insert bool, related *Environment) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"environment1_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
	)
	values := []interface{}{related.EnvironmentID, o.PhenotypeComparisonID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Environment1ID = related.EnvironmentID

	if o.R == nil {
		o.R = &phenotypeComparisonR{
			Environment1: related,
		}
	} else {
		o.R.Environment1 = related
	}

	if related.R == nil {
		related.R = &environmentR{
			Environment1PhenotypeComparison: o,
		}
	} else {
		related.R.Environment1PhenotypeComparison = o
	}

	return nil
}

// SetEnvironment2 of the phenotype_comparison to the related item.
// Sets o.R.Environment2 to related.
// Adds o to related.R.Environment2PhenotypeComparison.
func (o *PhenotypeComparison) SetEnvironment2(exec boil.Executor, insert bool, related *Environment) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"environment2_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
	)
	values := []interface{}{related.EnvironmentID, o.PhenotypeComparisonID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Environment2ID = related.EnvironmentID

	if o.R == nil {
		o.R = &phenotypeComparisonR{
			Environment2: related,
		}
	} else {
		o.R.Environment2 = related
	}

	if related.R == nil {
		related.R = &environmentR{
			Environment2PhenotypeComparison: o,
		}
	} else {
		related.R.Environment2PhenotypeComparison = o
	}

	return nil
}

// SetGenotype1 of the phenotype_comparison to the related item.
// Sets o.R.Genotype1 to related.
// Adds o to related.R.Genotype1PhenotypeComparison.
func (o *PhenotypeComparison) SetGenotype1(exec boil.Executor, insert bool, related *Genotype) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"genotype1_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
	)
	values := []interface{}{related.GenotypeID, o.PhenotypeComparisonID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Genotype1ID = related.GenotypeID

	if o.R == nil {
		o.R = &phenotypeComparisonR{
			Genotype1: related,
		}
	} else {
		o.R.Genotype1 = related
	}

	if related.R == nil {
		related.R = &genotypeR{
			Genotype1PhenotypeComparison: o,
		}
	} else {
		related.R.Genotype1PhenotypeComparison = o
	}

	return nil
}

// SetGenotype2 of the phenotype_comparison to the related item.
// Sets o.R.Genotype2 to related.
// Adds o to related.R.Genotype2PhenotypeComparison.
func (o *PhenotypeComparison) SetGenotype2(exec boil.Executor, insert bool, related *Genotype) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"genotype2_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
	)
	values := []interface{}{related.GenotypeID, o.PhenotypeComparisonID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Genotype2ID = related.GenotypeID

	if o.R == nil {
		o.R = &phenotypeComparisonR{
			Genotype2: related,
		}
	} else {
		o.R.Genotype2 = related
	}

	if related.R == nil {
		related.R = &genotypeR{
			Genotype2PhenotypeComparison: o,
		}
	} else {
		related.R.Genotype2PhenotypeComparison = o
	}

	return nil
}

// SetOrganism of the phenotype_comparison to the related item.
// Sets o.R.Organism to related.
// Adds o to related.R.PhenotypeComparisons.
func (o *PhenotypeComparison) SetOrganism(exec boil.Executor, insert bool, related *Organism) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"organism_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
	)
	values := []interface{}{related.OrganismID, o.PhenotypeComparisonID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrganismID = related.OrganismID

	if o.R == nil {
		o.R = &phenotypeComparisonR{
			Organism: related,
		}
	} else {
		o.R.Organism = related
	}

	if related.R == nil {
		related.R = &organismR{
			PhenotypeComparisons: PhenotypeComparisonSlice{o},
		}
	} else {
		related.R.PhenotypeComparisons = append(related.R.PhenotypeComparisons, o)
	}

	return nil
}

// SetPhenotype1 of the phenotype_comparison to the related item.
// Sets o.R.Phenotype1 to related.
// Adds o to related.R.Phenotype1PhenotypeComparison.
func (o *PhenotypeComparison) SetPhenotype1(exec boil.Executor, insert bool, related *Phenotype) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"phenotype1_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
	)
	values := []interface{}{related.PhenotypeID, o.PhenotypeComparisonID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Phenotype1ID = related.PhenotypeID

	if o.R == nil {
		o.R = &phenotypeComparisonR{
			Phenotype1: related,
		}
	} else {
		o.R.Phenotype1 = related
	}

	if related.R == nil {
		related.R = &phenotypeR{
			Phenotype1PhenotypeComparison: o,
		}
	} else {
		related.R.Phenotype1PhenotypeComparison = o
	}

	return nil
}

// SetPhenotype2 of the phenotype_comparison to the related item.
// Sets o.R.Phenotype2 to related.
// Adds o to related.R.Phenotype2PhenotypeComparisons.
func (o *PhenotypeComparison) SetPhenotype2(exec boil.Executor, insert bool, related *Phenotype) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"phenotype2_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
	)
	values := []interface{}{related.PhenotypeID, o.PhenotypeComparisonID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Phenotype2ID.Int = related.PhenotypeID
	o.Phenotype2ID.Valid = true

	if o.R == nil {
		o.R = &phenotypeComparisonR{
			Phenotype2: related,
		}
	} else {
		o.R.Phenotype2 = related
	}

	if related.R == nil {
		related.R = &phenotypeR{
			Phenotype2PhenotypeComparisons: PhenotypeComparisonSlice{o},
		}
	} else {
		related.R.Phenotype2PhenotypeComparisons = append(related.R.Phenotype2PhenotypeComparisons, o)
	}

	return nil
}

// RemovePhenotype2 relationship.
// Sets o.R.Phenotype2 to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *PhenotypeComparison) RemovePhenotype2(exec boil.Executor, related *Phenotype) error {
	var err error

	o.Phenotype2ID.Valid = false
	if err = o.Update(exec, "phenotype2_id"); err != nil {
		o.Phenotype2ID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Phenotype2 = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Phenotype2PhenotypeComparisons {
		if o.Phenotype2ID.Int != ri.Phenotype2ID.Int {
			continue
		}

		ln := len(related.R.Phenotype2PhenotypeComparisons)
		if ln > 1 && i < ln-1 {
			related.R.Phenotype2PhenotypeComparisons[i] = related.R.Phenotype2PhenotypeComparisons[ln-1]
		}
		related.R.Phenotype2PhenotypeComparisons = related.R.Phenotype2PhenotypeComparisons[:ln-1]
		break
	}
	return nil
}

// SetPub of the phenotype_comparison to the related item.
// Sets o.R.Pub to related.
// Adds o to related.R.PhenotypeComparison.
func (o *PhenotypeComparison) SetPub(exec boil.Executor, insert bool, related *Pub) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
		strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
	)
	values := []interface{}{related.PubID, o.PhenotypeComparisonID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PubID = related.PubID

	if o.R == nil {
		o.R = &phenotypeComparisonR{
			Pub: related,
		}
	} else {
		o.R.Pub = related
	}

	if related.R == nil {
		related.R = &pubR{
			PhenotypeComparison: o,
		}
	} else {
		related.R.PhenotypeComparison = o
	}

	return nil
}

// SetPhenotypeComparisonCvterm of the phenotype_comparison to the related item.
// Sets o.R.PhenotypeComparisonCvterm to related.
// Adds o to related.R.PhenotypeComparison.
func (o *PhenotypeComparison) SetPhenotypeComparisonCvterm(exec boil.Executor, insert bool, related *PhenotypeComparisonCvterm) error {
	var err error

	if insert {
		related.PhenotypeComparisonID = o.PhenotypeComparisonID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenotype_comparison_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"phenotype_comparison_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonCvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.PhenotypeComparisonID, related.PhenotypeComparisonCvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PhenotypeComparisonID = o.PhenotypeComparisonID

	}

	if o.R == nil {
		o.R = &phenotypeComparisonR{
			PhenotypeComparisonCvterm: related,
		}
	} else {
		o.R.PhenotypeComparisonCvterm = related
	}

	if related.R == nil {
		related.R = &phenotypeComparisonCvtermR{
			PhenotypeComparison: o,
		}
	} else {
		related.R.PhenotypeComparison = o
	}
	return nil
}

// PhenotypeComparisonsG retrieves all records.
func PhenotypeComparisonsG(mods ...qm.QueryMod) phenotypeComparisonQuery {
	return PhenotypeComparisons(boil.GetDB(), mods...)
}

// PhenotypeComparisons retrieves all the records using an executor.
func PhenotypeComparisons(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonQuery {
	mods = append(mods, qm.From("\"phenotype_comparison\""))
	return phenotypeComparisonQuery{NewQuery(exec, mods...)}
}

// FindPhenotypeComparisonG retrieves a single record by ID.
func FindPhenotypeComparisonG(phenotypeComparisonID int, selectCols ...string) (*PhenotypeComparison, error) {
	return FindPhenotypeComparison(boil.GetDB(), phenotypeComparisonID, selectCols...)
}

// FindPhenotypeComparisonGP retrieves a single record by ID, and panics on error.
func FindPhenotypeComparisonGP(phenotypeComparisonID int, selectCols ...string) *PhenotypeComparison {
	retobj, err := FindPhenotypeComparison(boil.GetDB(), phenotypeComparisonID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindPhenotypeComparison retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPhenotypeComparison(exec boil.Executor, phenotypeComparisonID int, selectCols ...string) (*PhenotypeComparison, error) {
	phenotypeComparisonObj := &PhenotypeComparison{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"phenotype_comparison\" where \"phenotype_comparison_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, phenotypeComparisonID)

	err := q.Bind(phenotypeComparisonObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from phenotype_comparison")
	}

	return phenotypeComparisonObj, nil
}

// FindPhenotypeComparisonP retrieves a single record by ID with an executor, and panics on error.
func FindPhenotypeComparisonP(exec boil.Executor, phenotypeComparisonID int, selectCols ...string) *PhenotypeComparison {
	retobj, err := FindPhenotypeComparison(exec, phenotypeComparisonID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *PhenotypeComparison) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *PhenotypeComparison) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *PhenotypeComparison) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *PhenotypeComparison) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no phenotype_comparison provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(phenotypeComparisonColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	phenotypeComparisonInsertCacheMut.RLock()
	cache, cached := phenotypeComparisonInsertCache[key]
	phenotypeComparisonInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			phenotypeComparisonColumns,
			phenotypeComparisonColumnsWithDefault,
			phenotypeComparisonColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(phenotypeComparisonType, phenotypeComparisonMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(phenotypeComparisonType, phenotypeComparisonMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"phenotype_comparison\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into phenotype_comparison")
	}

	if !cached {
		phenotypeComparisonInsertCacheMut.Lock()
		phenotypeComparisonInsertCache[key] = cache
		phenotypeComparisonInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single PhenotypeComparison record. See Update for
// whitelist behavior description.
func (o *PhenotypeComparison) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single PhenotypeComparison record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *PhenotypeComparison) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the PhenotypeComparison, and panics on error.
// See Update for whitelist behavior description.
func (o *PhenotypeComparison) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the PhenotypeComparison.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *PhenotypeComparison) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	phenotypeComparisonUpdateCacheMut.RLock()
	cache, cached := phenotypeComparisonUpdateCache[key]
	phenotypeComparisonUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(phenotypeComparisonColumns, phenotypeComparisonPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update phenotype_comparison, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"phenotype_comparison\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, phenotypeComparisonPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(phenotypeComparisonType, phenotypeComparisonMapping, append(wl, phenotypeComparisonPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update phenotype_comparison row")
	}

	if !cached {
		phenotypeComparisonUpdateCacheMut.Lock()
		phenotypeComparisonUpdateCache[key] = cache
		phenotypeComparisonUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q phenotypeComparisonQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q phenotypeComparisonQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for phenotype_comparison")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PhenotypeComparisonSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o PhenotypeComparisonSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o PhenotypeComparisonSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PhenotypeComparisonSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenotypeComparisonPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"phenotype_comparison\" SET %s WHERE (\"phenotype_comparison_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(phenotypeComparisonPrimaryKeyColumns), len(colNames)+1, len(phenotypeComparisonPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in phenotypeComparison slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *PhenotypeComparison) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *PhenotypeComparison) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *PhenotypeComparison) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *PhenotypeComparison) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no phenotype_comparison provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(phenotypeComparisonColumnsWithDefault, o)

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

	phenotypeComparisonUpsertCacheMut.RLock()
	cache, cached := phenotypeComparisonUpsertCache[key]
	phenotypeComparisonUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			phenotypeComparisonColumns,
			phenotypeComparisonColumnsWithDefault,
			phenotypeComparisonColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			phenotypeComparisonColumns,
			phenotypeComparisonPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert phenotype_comparison, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(phenotypeComparisonPrimaryKeyColumns))
			copy(conflict, phenotypeComparisonPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"phenotype_comparison\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(phenotypeComparisonType, phenotypeComparisonMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(phenotypeComparisonType, phenotypeComparisonMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for phenotype_comparison")
	}

	if !cached {
		phenotypeComparisonUpsertCacheMut.Lock()
		phenotypeComparisonUpsertCache[key] = cache
		phenotypeComparisonUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single PhenotypeComparison record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *PhenotypeComparison) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single PhenotypeComparison record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *PhenotypeComparison) DeleteG() error {
	if o == nil {
		return errors.New("chado: no PhenotypeComparison provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single PhenotypeComparison record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *PhenotypeComparison) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single PhenotypeComparison record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PhenotypeComparison) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no PhenotypeComparison provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), phenotypeComparisonPrimaryKeyMapping)
	sql := "DELETE FROM \"phenotype_comparison\" WHERE \"phenotype_comparison_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from phenotype_comparison")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q phenotypeComparisonQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q phenotypeComparisonQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no phenotypeComparisonQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from phenotype_comparison")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o PhenotypeComparisonSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o PhenotypeComparisonSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no PhenotypeComparison slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o PhenotypeComparisonSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PhenotypeComparisonSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no PhenotypeComparison slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(phenotypeComparisonBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenotypeComparisonPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"phenotype_comparison\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, phenotypeComparisonPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(phenotypeComparisonPrimaryKeyColumns), 1, len(phenotypeComparisonPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from phenotypeComparison slice")
	}

	if len(phenotypeComparisonAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *PhenotypeComparison) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *PhenotypeComparison) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *PhenotypeComparison) ReloadG() error {
	if o == nil {
		return errors.New("chado: no PhenotypeComparison provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PhenotypeComparison) Reload(exec boil.Executor) error {
	ret, err := FindPhenotypeComparison(exec, o.PhenotypeComparisonID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PhenotypeComparisonSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PhenotypeComparisonSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhenotypeComparisonSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty PhenotypeComparisonSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PhenotypeComparisonSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	phenotypeComparisons := PhenotypeComparisonSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), phenotypeComparisonPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"phenotype_comparison\".* FROM \"phenotype_comparison\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, phenotypeComparisonPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(phenotypeComparisonPrimaryKeyColumns), 1, len(phenotypeComparisonPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&phenotypeComparisons)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in PhenotypeComparisonSlice")
	}

	*o = phenotypeComparisons

	return nil
}

// PhenotypeComparisonExists checks if the PhenotypeComparison row exists.
func PhenotypeComparisonExists(exec boil.Executor, phenotypeComparisonID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"phenotype_comparison\" where \"phenotype_comparison_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, phenotypeComparisonID)
	}

	row := exec.QueryRow(sql, phenotypeComparisonID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if phenotype_comparison exists")
	}

	return exists, nil
}

// PhenotypeComparisonExistsG checks if the PhenotypeComparison row exists.
func PhenotypeComparisonExistsG(phenotypeComparisonID int) (bool, error) {
	return PhenotypeComparisonExists(boil.GetDB(), phenotypeComparisonID)
}

// PhenotypeComparisonExistsGP checks if the PhenotypeComparison row exists. Panics on error.
func PhenotypeComparisonExistsGP(phenotypeComparisonID int) bool {
	e, err := PhenotypeComparisonExists(boil.GetDB(), phenotypeComparisonID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// PhenotypeComparisonExistsP checks if the PhenotypeComparison row exists. Panics on error.
func PhenotypeComparisonExistsP(exec boil.Executor, phenotypeComparisonID int) bool {
	e, err := PhenotypeComparisonExists(exec, phenotypeComparisonID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
