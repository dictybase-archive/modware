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

// FeatureGenotype is an object representing the database table.
type FeatureGenotype struct {
	FeatureGenotypeID int      `boil:"feature_genotype_id" json:"feature_genotype_id" toml:"feature_genotype_id" yaml:"feature_genotype_id"`
	FeatureID         int      `boil:"feature_id" json:"feature_id" toml:"feature_id" yaml:"feature_id"`
	GenotypeID        int      `boil:"genotype_id" json:"genotype_id" toml:"genotype_id" yaml:"genotype_id"`
	ChromosomeID      null.Int `boil:"chromosome_id" json:"chromosome_id,omitempty" toml:"chromosome_id" yaml:"chromosome_id,omitempty"`
	Rank              int      `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`
	Cgroup            int      `boil:"cgroup" json:"cgroup" toml:"cgroup" yaml:"cgroup"`
	CvtermID          int      `boil:"cvterm_id" json:"cvterm_id" toml:"cvterm_id" yaml:"cvterm_id"`

	R *featureGenotypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featureGenotypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featureGenotypeR is where relationships are stored.
type featureGenotypeR struct {
	Genotype   *Genotype
	Cvterm     *Cvterm
	Chromosome *Feature
	Feature    *Feature
}

// featureGenotypeL is where Load methods for each relationship are stored.
type featureGenotypeL struct{}

var (
	featureGenotypeColumns               = []string{"feature_genotype_id", "feature_id", "genotype_id", "chromosome_id", "rank", "cgroup", "cvterm_id"}
	featureGenotypeColumnsWithoutDefault = []string{"feature_id", "genotype_id", "chromosome_id", "rank", "cgroup", "cvterm_id"}
	featureGenotypeColumnsWithDefault    = []string{"feature_genotype_id"}
	featureGenotypePrimaryKeyColumns     = []string{"feature_genotype_id"}
)

type (
	// FeatureGenotypeSlice is an alias for a slice of pointers to FeatureGenotype.
	// This should generally be used opposed to []FeatureGenotype.
	FeatureGenotypeSlice []*FeatureGenotype
	// FeatureGenotypeHook is the signature for custom FeatureGenotype hook methods
	FeatureGenotypeHook func(boil.Executor, *FeatureGenotype) error

	featureGenotypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featureGenotypeType                 = reflect.TypeOf(&FeatureGenotype{})
	featureGenotypeMapping              = queries.MakeStructMapping(featureGenotypeType)
	featureGenotypePrimaryKeyMapping, _ = queries.BindMapping(featureGenotypeType, featureGenotypeMapping, featureGenotypePrimaryKeyColumns)
	featureGenotypeInsertCacheMut       sync.RWMutex
	featureGenotypeInsertCache          = make(map[string]insertCache)
	featureGenotypeUpdateCacheMut       sync.RWMutex
	featureGenotypeUpdateCache          = make(map[string]updateCache)
	featureGenotypeUpsertCacheMut       sync.RWMutex
	featureGenotypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featureGenotypeBeforeInsertHooks []FeatureGenotypeHook
var featureGenotypeBeforeUpdateHooks []FeatureGenotypeHook
var featureGenotypeBeforeDeleteHooks []FeatureGenotypeHook
var featureGenotypeBeforeUpsertHooks []FeatureGenotypeHook

var featureGenotypeAfterInsertHooks []FeatureGenotypeHook
var featureGenotypeAfterSelectHooks []FeatureGenotypeHook
var featureGenotypeAfterUpdateHooks []FeatureGenotypeHook
var featureGenotypeAfterDeleteHooks []FeatureGenotypeHook
var featureGenotypeAfterUpsertHooks []FeatureGenotypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FeatureGenotype) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureGenotypeBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FeatureGenotype) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureGenotypeBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FeatureGenotype) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureGenotypeBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FeatureGenotype) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureGenotypeBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FeatureGenotype) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureGenotypeAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FeatureGenotype) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featureGenotypeAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FeatureGenotype) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureGenotypeAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FeatureGenotype) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureGenotypeAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FeatureGenotype) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureGenotypeAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeatureGenotypeHook registers your hook function for all future operations.
func AddFeatureGenotypeHook(hookPoint boil.HookPoint, featureGenotypeHook FeatureGenotypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featureGenotypeBeforeInsertHooks = append(featureGenotypeBeforeInsertHooks, featureGenotypeHook)
	case boil.BeforeUpdateHook:
		featureGenotypeBeforeUpdateHooks = append(featureGenotypeBeforeUpdateHooks, featureGenotypeHook)
	case boil.BeforeDeleteHook:
		featureGenotypeBeforeDeleteHooks = append(featureGenotypeBeforeDeleteHooks, featureGenotypeHook)
	case boil.BeforeUpsertHook:
		featureGenotypeBeforeUpsertHooks = append(featureGenotypeBeforeUpsertHooks, featureGenotypeHook)
	case boil.AfterInsertHook:
		featureGenotypeAfterInsertHooks = append(featureGenotypeAfterInsertHooks, featureGenotypeHook)
	case boil.AfterSelectHook:
		featureGenotypeAfterSelectHooks = append(featureGenotypeAfterSelectHooks, featureGenotypeHook)
	case boil.AfterUpdateHook:
		featureGenotypeAfterUpdateHooks = append(featureGenotypeAfterUpdateHooks, featureGenotypeHook)
	case boil.AfterDeleteHook:
		featureGenotypeAfterDeleteHooks = append(featureGenotypeAfterDeleteHooks, featureGenotypeHook)
	case boil.AfterUpsertHook:
		featureGenotypeAfterUpsertHooks = append(featureGenotypeAfterUpsertHooks, featureGenotypeHook)
	}
}

// OneP returns a single featureGenotype record from the query, and panics on error.
func (q featureGenotypeQuery) OneP() *FeatureGenotype {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single featureGenotype record from the query.
func (q featureGenotypeQuery) One() (*FeatureGenotype, error) {
	o := &FeatureGenotype{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for feature_genotype")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all FeatureGenotype records from the query, and panics on error.
func (q featureGenotypeQuery) AllP() FeatureGenotypeSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all FeatureGenotype records from the query.
func (q featureGenotypeQuery) All() (FeatureGenotypeSlice, error) {
	var o FeatureGenotypeSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FeatureGenotype slice")
	}

	if len(featureGenotypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all FeatureGenotype records in the query, and panics on error.
func (q featureGenotypeQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all FeatureGenotype records in the query.
func (q featureGenotypeQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count feature_genotype rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featureGenotypeQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featureGenotypeQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if feature_genotype exists")
	}

	return count > 0, nil
}

// GenotypeG pointed to by the foreign key.
func (o *FeatureGenotype) GenotypeG(mods ...qm.QueryMod) genotypeQuery {
	return o.Genotype(boil.GetDB(), mods...)
}

// Genotype pointed to by the foreign key.
func (o *FeatureGenotype) Genotype(exec boil.Executor, mods ...qm.QueryMod) genotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("genotype_id=$1", o.GenotypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Genotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"genotype\"")

	return query
}

// CvtermG pointed to by the foreign key.
func (o *FeatureGenotype) CvtermG(mods ...qm.QueryMod) cvtermQuery {
	return o.Cvterm(boil.GetDB(), mods...)
}

// Cvterm pointed to by the foreign key.
func (o *FeatureGenotype) Cvterm(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// ChromosomeG pointed to by the foreign key.
func (o *FeatureGenotype) ChromosomeG(mods ...qm.QueryMod) featureQuery {
	return o.Chromosome(boil.GetDB(), mods...)
}

// Chromosome pointed to by the foreign key.
func (o *FeatureGenotype) Chromosome(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.ChromosomeID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// FeatureG pointed to by the foreign key.
func (o *FeatureGenotype) FeatureG(mods ...qm.QueryMod) featureQuery {
	return o.Feature(boil.GetDB(), mods...)
}

// Feature pointed to by the foreign key.
func (o *FeatureGenotype) Feature(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// LoadGenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureGenotypeL) LoadGenotype(e boil.Executor, singular bool, maybeFeatureGenotype interface{}) error {
	var slice []*FeatureGenotype
	var object *FeatureGenotype

	count := 1
	if singular {
		object = maybeFeatureGenotype.(*FeatureGenotype)
	} else {
		slice = *maybeFeatureGenotype.(*FeatureGenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureGenotypeR{}
		args[0] = object.GenotypeID
	} else {
		for i, obj := range slice {
			obj.R = &featureGenotypeR{}
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

	if len(featureGenotypeAfterSelectHooks) != 0 {
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

// LoadCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureGenotypeL) LoadCvterm(e boil.Executor, singular bool, maybeFeatureGenotype interface{}) error {
	var slice []*FeatureGenotype
	var object *FeatureGenotype

	count := 1
	if singular {
		object = maybeFeatureGenotype.(*FeatureGenotype)
	} else {
		slice = *maybeFeatureGenotype.(*FeatureGenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureGenotypeR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &featureGenotypeR{}
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

	if len(featureGenotypeAfterSelectHooks) != 0 {
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

// LoadChromosome allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureGenotypeL) LoadChromosome(e boil.Executor, singular bool, maybeFeatureGenotype interface{}) error {
	var slice []*FeatureGenotype
	var object *FeatureGenotype

	count := 1
	if singular {
		object = maybeFeatureGenotype.(*FeatureGenotype)
	} else {
		slice = *maybeFeatureGenotype.(*FeatureGenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureGenotypeR{}
		args[0] = object.ChromosomeID
	} else {
		for i, obj := range slice {
			obj.R = &featureGenotypeR{}
			args[i] = obj.ChromosomeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Feature")
	}
	defer results.Close()

	var resultSlice []*Feature
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Feature")
	}

	if len(featureGenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Chromosome = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ChromosomeID.Int == foreign.FeatureID {
				local.R.Chromosome = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureGenotypeL) LoadFeature(e boil.Executor, singular bool, maybeFeatureGenotype interface{}) error {
	var slice []*FeatureGenotype
	var object *FeatureGenotype

	count := 1
	if singular {
		object = maybeFeatureGenotype.(*FeatureGenotype)
	} else {
		slice = *maybeFeatureGenotype.(*FeatureGenotypeSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureGenotypeR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureGenotypeR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Feature")
	}
	defer results.Close()

	var resultSlice []*Feature
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Feature")
	}

	if len(featureGenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Feature = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.FeatureID {
				local.R.Feature = foreign
				break
			}
		}
	}

	return nil
}

// SetGenotype of the feature_genotype to the related item.
// Sets o.R.Genotype to related.
// Adds o to related.R.FeatureGenotype.
func (o *FeatureGenotype) SetGenotype(exec boil.Executor, insert bool, related *Genotype) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_genotype\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"genotype_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureGenotypePrimaryKeyColumns),
	)
	values := []interface{}{related.GenotypeID, o.FeatureGenotypeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GenotypeID = related.GenotypeID

	if o.R == nil {
		o.R = &featureGenotypeR{
			Genotype: related,
		}
	} else {
		o.R.Genotype = related
	}

	if related.R == nil {
		related.R = &genotypeR{
			FeatureGenotype: o,
		}
	} else {
		related.R.FeatureGenotype = o
	}

	return nil
}

// SetCvterm of the feature_genotype to the related item.
// Sets o.R.Cvterm to related.
// Adds o to related.R.FeatureGenotype.
func (o *FeatureGenotype) SetCvterm(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_genotype\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureGenotypePrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.FeatureGenotypeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CvtermID = related.CvtermID

	if o.R == nil {
		o.R = &featureGenotypeR{
			Cvterm: related,
		}
	} else {
		o.R.Cvterm = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			FeatureGenotype: o,
		}
	} else {
		related.R.FeatureGenotype = o
	}

	return nil
}

// SetChromosome of the feature_genotype to the related item.
// Sets o.R.Chromosome to related.
// Adds o to related.R.ChromosomeFeatureGenotype.
func (o *FeatureGenotype) SetChromosome(exec boil.Executor, insert bool, related *Feature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_genotype\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"chromosome_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureGenotypePrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureID, o.FeatureGenotypeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ChromosomeID.Int = related.FeatureID
	o.ChromosomeID.Valid = true

	if o.R == nil {
		o.R = &featureGenotypeR{
			Chromosome: related,
		}
	} else {
		o.R.Chromosome = related
	}

	if related.R == nil {
		related.R = &featureR{
			ChromosomeFeatureGenotype: o,
		}
	} else {
		related.R.ChromosomeFeatureGenotype = o
	}

	return nil
}

// RemoveChromosome relationship.
// Sets o.R.Chromosome to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *FeatureGenotype) RemoveChromosome(exec boil.Executor, related *Feature) error {
	var err error

	o.ChromosomeID.Valid = false
	if err = o.Update(exec, "chromosome_id"); err != nil {
		o.ChromosomeID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Chromosome = nil
	if related == nil || related.R == nil {
		return nil
	}

	related.R.ChromosomeFeatureGenotype = nil
	return nil
}

// SetFeature of the feature_genotype to the related item.
// Sets o.R.Feature to related.
// Adds o to related.R.FeatureGenotype.
func (o *FeatureGenotype) SetFeature(exec boil.Executor, insert bool, related *Feature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature_genotype\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
		strmangle.WhereClause("\"", "\"", 2, featureGenotypePrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureID, o.FeatureGenotypeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureID = related.FeatureID

	if o.R == nil {
		o.R = &featureGenotypeR{
			Feature: related,
		}
	} else {
		o.R.Feature = related
	}

	if related.R == nil {
		related.R = &featureR{
			FeatureGenotype: o,
		}
	} else {
		related.R.FeatureGenotype = o
	}

	return nil
}

// FeatureGenotypesG retrieves all records.
func FeatureGenotypesG(mods ...qm.QueryMod) featureGenotypeQuery {
	return FeatureGenotypes(boil.GetDB(), mods...)
}

// FeatureGenotypes retrieves all the records using an executor.
func FeatureGenotypes(exec boil.Executor, mods ...qm.QueryMod) featureGenotypeQuery {
	mods = append(mods, qm.From("\"feature_genotype\""))
	return featureGenotypeQuery{NewQuery(exec, mods...)}
}

// FindFeatureGenotypeG retrieves a single record by ID.
func FindFeatureGenotypeG(featureGenotypeID int, selectCols ...string) (*FeatureGenotype, error) {
	return FindFeatureGenotype(boil.GetDB(), featureGenotypeID, selectCols...)
}

// FindFeatureGenotypeGP retrieves a single record by ID, and panics on error.
func FindFeatureGenotypeGP(featureGenotypeID int, selectCols ...string) *FeatureGenotype {
	retobj, err := FindFeatureGenotype(boil.GetDB(), featureGenotypeID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeatureGenotype retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeatureGenotype(exec boil.Executor, featureGenotypeID int, selectCols ...string) (*FeatureGenotype, error) {
	featureGenotypeObj := &FeatureGenotype{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature_genotype\" where \"feature_genotype_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featureGenotypeID)

	err := q.Bind(featureGenotypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from feature_genotype")
	}

	return featureGenotypeObj, nil
}

// FindFeatureGenotypeP retrieves a single record by ID with an executor, and panics on error.
func FindFeatureGenotypeP(exec boil.Executor, featureGenotypeID int, selectCols ...string) *FeatureGenotype {
	retobj, err := FindFeatureGenotype(exec, featureGenotypeID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *FeatureGenotype) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *FeatureGenotype) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *FeatureGenotype) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *FeatureGenotype) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature_genotype provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureGenotypeColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featureGenotypeInsertCacheMut.RLock()
	cache, cached := featureGenotypeInsertCache[key]
	featureGenotypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featureGenotypeColumns,
			featureGenotypeColumnsWithDefault,
			featureGenotypeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featureGenotypeType, featureGenotypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featureGenotypeType, featureGenotypeMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature_genotype\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into feature_genotype")
	}

	if !cached {
		featureGenotypeInsertCacheMut.Lock()
		featureGenotypeInsertCache[key] = cache
		featureGenotypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single FeatureGenotype record. See Update for
// whitelist behavior description.
func (o *FeatureGenotype) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single FeatureGenotype record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *FeatureGenotype) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the FeatureGenotype, and panics on error.
// See Update for whitelist behavior description.
func (o *FeatureGenotype) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the FeatureGenotype.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *FeatureGenotype) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featureGenotypeUpdateCacheMut.RLock()
	cache, cached := featureGenotypeUpdateCache[key]
	featureGenotypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featureGenotypeColumns, featureGenotypePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update feature_genotype, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature_genotype\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featureGenotypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featureGenotypeType, featureGenotypeMapping, append(wl, featureGenotypePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update feature_genotype row")
	}

	if !cached {
		featureGenotypeUpdateCacheMut.Lock()
		featureGenotypeUpdateCache[key] = cache
		featureGenotypeUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featureGenotypeQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featureGenotypeQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for feature_genotype")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeatureGenotypeSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeatureGenotypeSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeatureGenotypeSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeatureGenotypeSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureGenotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature_genotype\" SET %s WHERE (\"feature_genotype_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureGenotypePrimaryKeyColumns), len(colNames)+1, len(featureGenotypePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in featureGenotype slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *FeatureGenotype) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *FeatureGenotype) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *FeatureGenotype) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *FeatureGenotype) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature_genotype provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureGenotypeColumnsWithDefault, o)

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

	featureGenotypeUpsertCacheMut.RLock()
	cache, cached := featureGenotypeUpsertCache[key]
	featureGenotypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featureGenotypeColumns,
			featureGenotypeColumnsWithDefault,
			featureGenotypeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featureGenotypeColumns,
			featureGenotypePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert feature_genotype, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featureGenotypePrimaryKeyColumns))
			copy(conflict, featureGenotypePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature_genotype\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featureGenotypeType, featureGenotypeMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featureGenotypeType, featureGenotypeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for feature_genotype")
	}

	if !cached {
		featureGenotypeUpsertCacheMut.Lock()
		featureGenotypeUpsertCache[key] = cache
		featureGenotypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single FeatureGenotype record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureGenotype) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single FeatureGenotype record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *FeatureGenotype) DeleteG() error {
	if o == nil {
		return errors.New("models: no FeatureGenotype provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single FeatureGenotype record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *FeatureGenotype) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single FeatureGenotype record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FeatureGenotype) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeatureGenotype provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featureGenotypePrimaryKeyMapping)
	sql := "DELETE FROM \"feature_genotype\" WHERE \"feature_genotype_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from feature_genotype")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featureGenotypeQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featureGenotypeQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no featureGenotypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from feature_genotype")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeatureGenotypeSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeatureGenotypeSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no FeatureGenotype slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeatureGenotypeSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeatureGenotypeSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no FeatureGenotype slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featureGenotypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureGenotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature_genotype\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureGenotypePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featureGenotypePrimaryKeyColumns), 1, len(featureGenotypePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from featureGenotype slice")
	}

	if len(featureGenotypeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *FeatureGenotype) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *FeatureGenotype) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *FeatureGenotype) ReloadG() error {
	if o == nil {
		return errors.New("models: no FeatureGenotype provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FeatureGenotype) Reload(exec boil.Executor) error {
	ret, err := FindFeatureGenotype(exec, o.FeatureGenotypeID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureGenotypeSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureGenotypeSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureGenotypeSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty FeatureGenotypeSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureGenotypeSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	featureGenotypes := FeatureGenotypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureGenotypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature_genotype\".* FROM \"feature_genotype\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featureGenotypePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featureGenotypePrimaryKeyColumns), 1, len(featureGenotypePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&featureGenotypes)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FeatureGenotypeSlice")
	}

	*o = featureGenotypes

	return nil
}

// FeatureGenotypeExists checks if the FeatureGenotype row exists.
func FeatureGenotypeExists(exec boil.Executor, featureGenotypeID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature_genotype\" where \"feature_genotype_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featureGenotypeID)
	}

	row := exec.QueryRow(sql, featureGenotypeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if feature_genotype exists")
	}

	return exists, nil
}

// FeatureGenotypeExistsG checks if the FeatureGenotype row exists.
func FeatureGenotypeExistsG(featureGenotypeID int) (bool, error) {
	return FeatureGenotypeExists(boil.GetDB(), featureGenotypeID)
}

// FeatureGenotypeExistsGP checks if the FeatureGenotype row exists. Panics on error.
func FeatureGenotypeExistsGP(featureGenotypeID int) bool {
	e, err := FeatureGenotypeExists(boil.GetDB(), featureGenotypeID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeatureGenotypeExistsP checks if the FeatureGenotype row exists. Panics on error.
func FeatureGenotypeExistsP(exec boil.Executor, featureGenotypeID int) bool {
	e, err := FeatureGenotypeExists(exec, featureGenotypeID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
