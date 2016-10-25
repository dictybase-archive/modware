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

// Feature is an object representing the database table.
type Feature struct {
	FeatureID        int         `boil:"feature_id" json:"feature_id" toml:"feature_id" yaml:"feature_id"`
	DbxrefID         null.Int    `boil:"dbxref_id" json:"dbxref_id,omitempty" toml:"dbxref_id" yaml:"dbxref_id,omitempty"`
	OrganismID       int         `boil:"organism_id" json:"organism_id" toml:"organism_id" yaml:"organism_id"`
	Name             null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Uniquename       string      `boil:"uniquename" json:"uniquename" toml:"uniquename" yaml:"uniquename"`
	Residues         null.String `boil:"residues" json:"residues,omitempty" toml:"residues" yaml:"residues,omitempty"`
	Seqlen           null.Int    `boil:"seqlen" json:"seqlen,omitempty" toml:"seqlen" yaml:"seqlen,omitempty"`
	Md5checksum      null.String `boil:"md5checksum" json:"md5checksum,omitempty" toml:"md5checksum" yaml:"md5checksum,omitempty"`
	TypeID           int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	IsAnalysis       bool        `boil:"is_analysis" json:"is_analysis" toml:"is_analysis" yaml:"is_analysis"`
	IsObsolete       bool        `boil:"is_obsolete" json:"is_obsolete" toml:"is_obsolete" yaml:"is_obsolete"`
	Timeaccessioned  time.Time   `boil:"timeaccessioned" json:"timeaccessioned" toml:"timeaccessioned" yaml:"timeaccessioned"`
	Timelastmodified time.Time   `boil:"timelastmodified" json:"timelastmodified" toml:"timelastmodified" yaml:"timelastmodified"`

	R *featureR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featureL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// featureR is where relationships are stored.
type featureR struct {
	Organism                   *Organism
	Type                       *Cvterm
	Dbxref                     *Dbxref
	FeatureDbxref              *FeatureDbxref
	FeaturePhenotype           *FeaturePhenotype
	Analysisfeature            *Analysisfeature
	Featureprop                *Featureprop
	FeatureCvterm              *FeatureCvterm
	ChromosomeFeatureGenotype  *FeatureGenotype
	FeatureGenotype            *FeatureGenotype
	FeaturePub                 *FeaturePub
	FeatureSynonym             *FeatureSynonym
	ObjectFeatureRelationship  *FeatureRelationship
	SubjectFeatureRelationship *FeatureRelationship
	Featureloc                 *Featureloc
	SrcfeatureFeaturelocs      FeaturelocSlice
}

// featureL is where Load methods for each relationship are stored.
type featureL struct{}

var (
	featureColumns               = []string{"feature_id", "dbxref_id", "organism_id", "name", "uniquename", "residues", "seqlen", "md5checksum", "type_id", "is_analysis", "is_obsolete", "timeaccessioned", "timelastmodified"}
	featureColumnsWithoutDefault = []string{"dbxref_id", "organism_id", "name", "uniquename", "residues", "seqlen", "md5checksum", "type_id"}
	featureColumnsWithDefault    = []string{"feature_id", "is_analysis", "is_obsolete", "timeaccessioned", "timelastmodified"}
	featurePrimaryKeyColumns     = []string{"feature_id"}
)

type (
	// FeatureSlice is an alias for a slice of pointers to Feature.
	// This should generally be used opposed to []Feature.
	FeatureSlice []*Feature
	// FeatureHook is the signature for custom Feature hook methods
	FeatureHook func(boil.Executor, *Feature) error

	featureQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featureType                 = reflect.TypeOf(&Feature{})
	featureMapping              = queries.MakeStructMapping(featureType)
	featurePrimaryKeyMapping, _ = queries.BindMapping(featureType, featureMapping, featurePrimaryKeyColumns)
	featureInsertCacheMut       sync.RWMutex
	featureInsertCache          = make(map[string]insertCache)
	featureUpdateCacheMut       sync.RWMutex
	featureUpdateCache          = make(map[string]updateCache)
	featureUpsertCacheMut       sync.RWMutex
	featureUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var featureBeforeInsertHooks []FeatureHook
var featureBeforeUpdateHooks []FeatureHook
var featureBeforeDeleteHooks []FeatureHook
var featureBeforeUpsertHooks []FeatureHook

var featureAfterInsertHooks []FeatureHook
var featureAfterSelectHooks []FeatureHook
var featureAfterUpdateHooks []FeatureHook
var featureAfterDeleteHooks []FeatureHook
var featureAfterUpsertHooks []FeatureHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Feature) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Feature) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Feature) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Feature) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Feature) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Feature) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range featureAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Feature) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range featureAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Feature) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range featureAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Feature) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range featureAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFeatureHook registers your hook function for all future operations.
func AddFeatureHook(hookPoint boil.HookPoint, featureHook FeatureHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		featureBeforeInsertHooks = append(featureBeforeInsertHooks, featureHook)
	case boil.BeforeUpdateHook:
		featureBeforeUpdateHooks = append(featureBeforeUpdateHooks, featureHook)
	case boil.BeforeDeleteHook:
		featureBeforeDeleteHooks = append(featureBeforeDeleteHooks, featureHook)
	case boil.BeforeUpsertHook:
		featureBeforeUpsertHooks = append(featureBeforeUpsertHooks, featureHook)
	case boil.AfterInsertHook:
		featureAfterInsertHooks = append(featureAfterInsertHooks, featureHook)
	case boil.AfterSelectHook:
		featureAfterSelectHooks = append(featureAfterSelectHooks, featureHook)
	case boil.AfterUpdateHook:
		featureAfterUpdateHooks = append(featureAfterUpdateHooks, featureHook)
	case boil.AfterDeleteHook:
		featureAfterDeleteHooks = append(featureAfterDeleteHooks, featureHook)
	case boil.AfterUpsertHook:
		featureAfterUpsertHooks = append(featureAfterUpsertHooks, featureHook)
	}
}

// OneP returns a single feature record from the query, and panics on error.
func (q featureQuery) OneP() *Feature {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single feature record from the query.
func (q featureQuery) One() (*Feature, error) {
	o := &Feature{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for feature")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Feature records from the query, and panics on error.
func (q featureQuery) AllP() FeatureSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Feature records from the query.
func (q featureQuery) All() (FeatureSlice, error) {
	var o FeatureSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Feature slice")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Feature records in the query, and panics on error.
func (q featureQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Feature records in the query.
func (q featureQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count feature rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q featureQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q featureQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if feature exists")
	}

	return count > 0, nil
}

// OrganismG pointed to by the foreign key.
func (o *Feature) OrganismG(mods ...qm.QueryMod) organismQuery {
	return o.Organism(boil.GetDB(), mods...)
}

// Organism pointed to by the foreign key.
func (o *Feature) Organism(exec boil.Executor, mods ...qm.QueryMod) organismQuery {
	queryMods := []qm.QueryMod{
		qm.Where("organism_id=$1", o.OrganismID),
	}

	queryMods = append(queryMods, mods...)

	query := Organisms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"organism\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *Feature) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Feature) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// DbxrefG pointed to by the foreign key.
func (o *Feature) DbxrefG(mods ...qm.QueryMod) dbxrefQuery {
	return o.Dbxref(boil.GetDB(), mods...)
}

// Dbxref pointed to by the foreign key.
func (o *Feature) Dbxref(exec boil.Executor, mods ...qm.QueryMod) dbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := Dbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"dbxref\"")

	return query
}

// FeatureDbxrefG pointed to by the foreign key.
func (o *Feature) FeatureDbxrefG(mods ...qm.QueryMod) featureDbxrefQuery {
	return o.FeatureDbxref(boil.GetDB(), mods...)
}

// FeatureDbxref pointed to by the foreign key.
func (o *Feature) FeatureDbxref(exec boil.Executor, mods ...qm.QueryMod) featureDbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureDbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_dbxref\"")

	return query
}

// FeaturePhenotypeG pointed to by the foreign key.
func (o *Feature) FeaturePhenotypeG(mods ...qm.QueryMod) featurePhenotypeQuery {
	return o.FeaturePhenotype(boil.GetDB(), mods...)
}

// FeaturePhenotype pointed to by the foreign key.
func (o *Feature) FeaturePhenotype(exec boil.Executor, mods ...qm.QueryMod) featurePhenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := FeaturePhenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_phenotype\"")

	return query
}

// AnalysisfeatureG pointed to by the foreign key.
func (o *Feature) AnalysisfeatureG(mods ...qm.QueryMod) analysisfeatureQuery {
	return o.Analysisfeature(boil.GetDB(), mods...)
}

// Analysisfeature pointed to by the foreign key.
func (o *Feature) Analysisfeature(exec boil.Executor, mods ...qm.QueryMod) analysisfeatureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Analysisfeatures(exec, queryMods...)
	queries.SetFrom(query.Query, "\"analysisfeature\"")

	return query
}

// FeaturepropG pointed to by the foreign key.
func (o *Feature) FeaturepropG(mods ...qm.QueryMod) featurepropQuery {
	return o.Featureprop(boil.GetDB(), mods...)
}

// Featureprop pointed to by the foreign key.
func (o *Feature) Featureprop(exec boil.Executor, mods ...qm.QueryMod) featurepropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Featureprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"featureprop\"")

	return query
}

// FeatureCvtermG pointed to by the foreign key.
func (o *Feature) FeatureCvtermG(mods ...qm.QueryMod) featureCvtermQuery {
	return o.FeatureCvterm(boil.GetDB(), mods...)
}

// FeatureCvterm pointed to by the foreign key.
func (o *Feature) FeatureCvterm(exec boil.Executor, mods ...qm.QueryMod) featureCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_cvterm\"")

	return query
}

// ChromosomeFeatureGenotypeG pointed to by the foreign key.
func (o *Feature) ChromosomeFeatureGenotypeG(mods ...qm.QueryMod) featureGenotypeQuery {
	return o.ChromosomeFeatureGenotype(boil.GetDB(), mods...)
}

// ChromosomeFeatureGenotype pointed to by the foreign key.
func (o *Feature) ChromosomeFeatureGenotype(exec boil.Executor, mods ...qm.QueryMod) featureGenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("chromosome_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureGenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_genotype\"")

	return query
}

// FeatureGenotypeG pointed to by the foreign key.
func (o *Feature) FeatureGenotypeG(mods ...qm.QueryMod) featureGenotypeQuery {
	return o.FeatureGenotype(boil.GetDB(), mods...)
}

// FeatureGenotype pointed to by the foreign key.
func (o *Feature) FeatureGenotype(exec boil.Executor, mods ...qm.QueryMod) featureGenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureGenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_genotype\"")

	return query
}

// FeaturePubG pointed to by the foreign key.
func (o *Feature) FeaturePubG(mods ...qm.QueryMod) featurePubQuery {
	return o.FeaturePub(boil.GetDB(), mods...)
}

// FeaturePub pointed to by the foreign key.
func (o *Feature) FeaturePub(exec boil.Executor, mods ...qm.QueryMod) featurePubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := FeaturePubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_pub\"")

	return query
}

// FeatureSynonymG pointed to by the foreign key.
func (o *Feature) FeatureSynonymG(mods ...qm.QueryMod) featureSynonymQuery {
	return o.FeatureSynonym(boil.GetDB(), mods...)
}

// FeatureSynonym pointed to by the foreign key.
func (o *Feature) FeatureSynonym(exec boil.Executor, mods ...qm.QueryMod) featureSynonymQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureSynonyms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_synonym\"")

	return query
}

// ObjectFeatureRelationshipG pointed to by the foreign key.
func (o *Feature) ObjectFeatureRelationshipG(mods ...qm.QueryMod) featureRelationshipQuery {
	return o.ObjectFeatureRelationship(boil.GetDB(), mods...)
}

// ObjectFeatureRelationship pointed to by the foreign key.
func (o *Feature) ObjectFeatureRelationship(exec boil.Executor, mods ...qm.QueryMod) featureRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("object_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_relationship\"")

	return query
}

// SubjectFeatureRelationshipG pointed to by the foreign key.
func (o *Feature) SubjectFeatureRelationshipG(mods ...qm.QueryMod) featureRelationshipQuery {
	return o.SubjectFeatureRelationship(boil.GetDB(), mods...)
}

// SubjectFeatureRelationship pointed to by the foreign key.
func (o *Feature) SubjectFeatureRelationship(exec boil.Executor, mods ...qm.QueryMod) featureRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("subject_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_relationship\"")

	return query
}

// FeaturelocG pointed to by the foreign key.
func (o *Feature) FeaturelocG(mods ...qm.QueryMod) featurelocQuery {
	return o.Featureloc(boil.GetDB(), mods...)
}

// Featureloc pointed to by the foreign key.
func (o *Feature) Featureloc(exec boil.Executor, mods ...qm.QueryMod) featurelocQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Featurelocs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"featureloc\"")

	return query
}

// SrcfeatureFeaturelocsG retrieves all the featureloc's featureloc via srcfeature_id column.
func (o *Feature) SrcfeatureFeaturelocsG(mods ...qm.QueryMod) featurelocQuery {
	return o.SrcfeatureFeaturelocs(boil.GetDB(), mods...)
}

// SrcfeatureFeaturelocs retrieves all the featureloc's featureloc with an executor via srcfeature_id column.
func (o *Feature) SrcfeatureFeaturelocs(exec boil.Executor, mods ...qm.QueryMod) featurelocQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"srcfeature_id\"=$1", o.FeatureID),
	)

	query := Featurelocs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"featureloc\" as \"a\"")
	return query
}

// LoadOrganism allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadOrganism(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.OrganismID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
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

	if len(featureAfterSelectHooks) != 0 {
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
func (featureL) LoadType(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
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

	if len(featureAfterSelectHooks) != 0 {
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

// LoadDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadDbxref(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.DbxrefID
		}
	}

	query := fmt.Sprintf(
		"select * from \"dbxref\" where \"dbxref_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Dbxref")
	}
	defer results.Close()

	var resultSlice []*Dbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Dbxref")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Dbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.DbxrefID.Int == foreign.DbxrefID {
				local.R.Dbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadFeatureDbxref(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_dbxref\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureDbxref")
	}
	defer results.Close()

	var resultSlice []*FeatureDbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureDbxref")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureDbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.FeatureID {
				local.R.FeatureDbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeaturePhenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadFeaturePhenotype(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_phenotype\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeaturePhenotype")
	}
	defer results.Close()

	var resultSlice []*FeaturePhenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeaturePhenotype")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeaturePhenotype = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.FeatureID {
				local.R.FeaturePhenotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadAnalysisfeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadAnalysisfeature(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"analysisfeature\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Analysisfeature")
	}
	defer results.Close()

	var resultSlice []*Analysisfeature
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Analysisfeature")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Analysisfeature = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.FeatureID {
				local.R.Analysisfeature = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadFeatureprop(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"featureprop\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Featureprop")
	}
	defer results.Close()

	var resultSlice []*Featureprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Featureprop")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Featureprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.FeatureID {
				local.R.Featureprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadFeatureCvterm(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_cvterm\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureCvterm")
	}
	defer results.Close()

	var resultSlice []*FeatureCvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureCvterm")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureCvterm = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.FeatureID {
				local.R.FeatureCvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadChromosomeFeatureGenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadChromosomeFeatureGenotype(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_genotype\" where \"chromosome_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureGenotype")
	}
	defer results.Close()

	var resultSlice []*FeatureGenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureGenotype")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.ChromosomeFeatureGenotype = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.ChromosomeID.Int {
				local.R.ChromosomeFeatureGenotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureGenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadFeatureGenotype(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_genotype\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureGenotype")
	}
	defer results.Close()

	var resultSlice []*FeatureGenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureGenotype")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureGenotype = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.FeatureID {
				local.R.FeatureGenotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeaturePub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadFeaturePub(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_pub\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeaturePub")
	}
	defer results.Close()

	var resultSlice []*FeaturePub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeaturePub")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeaturePub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.FeatureID {
				local.R.FeaturePub = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureSynonym allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadFeatureSynonym(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_synonym\" where \"feature_id\" in (%s)",
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

	if len(featureAfterSelectHooks) != 0 {
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
			if local.FeatureID == foreign.FeatureID {
				local.R.FeatureSynonym = foreign
				break
			}
		}
	}

	return nil
}

// LoadObjectFeatureRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadObjectFeatureRelationship(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_relationship\" where \"object_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureRelationship")
	}
	defer results.Close()

	var resultSlice []*FeatureRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureRelationship")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.ObjectFeatureRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.ObjectID {
				local.R.ObjectFeatureRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadSubjectFeatureRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadSubjectFeatureRelationship(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_relationship\" where \"subject_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureRelationship")
	}
	defer results.Close()

	var resultSlice []*FeatureRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureRelationship")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.SubjectFeatureRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.SubjectID {
				local.R.SubjectFeatureRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureloc allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadFeatureloc(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"featureloc\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Featureloc")
	}
	defer results.Close()

	var resultSlice []*Featureloc
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Featureloc")
	}

	if len(featureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Featureloc = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.FeatureID {
				local.R.Featureloc = foreign
				break
			}
		}
	}

	return nil
}

// LoadSrcfeatureFeaturelocs allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (featureL) LoadSrcfeatureFeaturelocs(e boil.Executor, singular bool, maybeFeature interface{}) error {
	var slice []*Feature
	var object *Feature

	count := 1
	if singular {
		object = maybeFeature.(*Feature)
	} else {
		slice = *maybeFeature.(*FeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &featureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &featureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"featureloc\" where \"srcfeature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load featureloc")
	}
	defer results.Close()

	var resultSlice []*Featureloc
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice featureloc")
	}

	if len(featurelocAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.SrcfeatureFeaturelocs = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.SrcfeatureID.Int {
				local.R.SrcfeatureFeaturelocs = append(local.R.SrcfeatureFeaturelocs, foreign)
				break
			}
		}
	}

	return nil
}

// SetOrganism of the feature to the related item.
// Sets o.R.Organism to related.
// Adds o to related.R.Feature.
func (o *Feature) SetOrganism(exec boil.Executor, insert bool, related *Organism) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"organism_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurePrimaryKeyColumns),
	)
	values := []interface{}{related.OrganismID, o.FeatureID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrganismID = related.OrganismID

	if o.R == nil {
		o.R = &featureR{
			Organism: related,
		}
	} else {
		o.R.Organism = related
	}

	if related.R == nil {
		related.R = &organismR{
			Feature: o,
		}
	} else {
		related.R.Feature = o
	}

	return nil
}

// SetType of the feature to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeFeature.
func (o *Feature) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurePrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.FeatureID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &featureR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeFeature: o,
		}
	} else {
		related.R.TypeFeature = o
	}

	return nil
}

// SetDbxref of the feature to the related item.
// Sets o.R.Dbxref to related.
// Adds o to related.R.Features.
func (o *Feature) SetDbxref(exec boil.Executor, insert bool, related *Dbxref) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"feature\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
		strmangle.WhereClause("\"", "\"", 2, featurePrimaryKeyColumns),
	)
	values := []interface{}{related.DbxrefID, o.FeatureID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DbxrefID.Int = related.DbxrefID
	o.DbxrefID.Valid = true

	if o.R == nil {
		o.R = &featureR{
			Dbxref: related,
		}
	} else {
		o.R.Dbxref = related
	}

	if related.R == nil {
		related.R = &dbxrefR{
			Features: FeatureSlice{o},
		}
	} else {
		related.R.Features = append(related.R.Features, o)
	}

	return nil
}

// RemoveDbxref relationship.
// Sets o.R.Dbxref to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Feature) RemoveDbxref(exec boil.Executor, related *Dbxref) error {
	var err error

	o.DbxrefID.Valid = false
	if err = o.Update(exec, "dbxref_id"); err != nil {
		o.DbxrefID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Dbxref = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Features {
		if o.DbxrefID.Int != ri.DbxrefID.Int {
			continue
		}

		ln := len(related.R.Features)
		if ln > 1 && i < ln-1 {
			related.R.Features[i] = related.R.Features[ln-1]
		}
		related.R.Features = related.R.Features[:ln-1]
		break
	}
	return nil
}

// SetFeatureDbxref of the feature to the related item.
// Sets o.R.FeatureDbxref to related.
// Adds o to related.R.Feature.
func (o *Feature) SetFeatureDbxref(exec boil.Executor, insert bool, related *FeatureDbxref) error {
	var err error

	if insert {
		related.FeatureID = o.FeatureID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureDbxrefPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureID, related.FeatureDbxrefID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureID = o.FeatureID

	}

	if o.R == nil {
		o.R = &featureR{
			FeatureDbxref: related,
		}
	} else {
		o.R.FeatureDbxref = related
	}

	if related.R == nil {
		related.R = &featureDbxrefR{
			Feature: o,
		}
	} else {
		related.R.Feature = o
	}
	return nil
}

// SetFeaturePhenotype of the feature to the related item.
// Sets o.R.FeaturePhenotype to related.
// Adds o to related.R.Feature.
func (o *Feature) SetFeaturePhenotype(exec boil.Executor, insert bool, related *FeaturePhenotype) error {
	var err error

	if insert {
		related.FeatureID = o.FeatureID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_phenotype\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurePhenotypePrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureID, related.FeaturePhenotypeID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureID = o.FeatureID

	}

	if o.R == nil {
		o.R = &featureR{
			FeaturePhenotype: related,
		}
	} else {
		o.R.FeaturePhenotype = related
	}

	if related.R == nil {
		related.R = &featurePhenotypeR{
			Feature: o,
		}
	} else {
		related.R.Feature = o
	}
	return nil
}

// SetAnalysisfeature of the feature to the related item.
// Sets o.R.Analysisfeature to related.
// Adds o to related.R.Feature.
func (o *Feature) SetAnalysisfeature(exec boil.Executor, insert bool, related *Analysisfeature) error {
	var err error

	if insert {
		related.FeatureID = o.FeatureID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"analysisfeature\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
			strmangle.WhereClause("\"", "\"", 2, analysisfeaturePrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureID, related.AnalysisfeatureID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureID = o.FeatureID

	}

	if o.R == nil {
		o.R = &featureR{
			Analysisfeature: related,
		}
	} else {
		o.R.Analysisfeature = related
	}

	if related.R == nil {
		related.R = &analysisfeatureR{
			Feature: o,
		}
	} else {
		related.R.Feature = o
	}
	return nil
}

// SetFeatureprop of the feature to the related item.
// Sets o.R.Featureprop to related.
// Adds o to related.R.Feature.
func (o *Feature) SetFeatureprop(exec boil.Executor, insert bool, related *Featureprop) error {
	var err error

	if insert {
		related.FeatureID = o.FeatureID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"featureprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurepropPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureID, related.FeaturepropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureID = o.FeatureID

	}

	if o.R == nil {
		o.R = &featureR{
			Featureprop: related,
		}
	} else {
		o.R.Featureprop = related
	}

	if related.R == nil {
		related.R = &featurepropR{
			Feature: o,
		}
	} else {
		related.R.Feature = o
	}
	return nil
}

// SetFeatureCvterm of the feature to the related item.
// Sets o.R.FeatureCvterm to related.
// Adds o to related.R.Feature.
func (o *Feature) SetFeatureCvterm(exec boil.Executor, insert bool, related *FeatureCvterm) error {
	var err error

	if insert {
		related.FeatureID = o.FeatureID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureCvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureID, related.FeatureCvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureID = o.FeatureID

	}

	if o.R == nil {
		o.R = &featureR{
			FeatureCvterm: related,
		}
	} else {
		o.R.FeatureCvterm = related
	}

	if related.R == nil {
		related.R = &featureCvtermR{
			Feature: o,
		}
	} else {
		related.R.Feature = o
	}
	return nil
}

// SetChromosomeFeatureGenotype of the feature to the related item.
// Sets o.R.ChromosomeFeatureGenotype to related.
// Adds o to related.R.Chromosome.
func (o *Feature) SetChromosomeFeatureGenotype(exec boil.Executor, insert bool, related *FeatureGenotype) error {
	var err error

	if insert {
		related.ChromosomeID.Int = o.FeatureID
		related.ChromosomeID.Valid = true

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_genotype\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"chromosome_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureGenotypePrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureID, related.FeatureGenotypeID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.ChromosomeID.Int = o.FeatureID
		related.ChromosomeID.Valid = true
	}

	if o.R == nil {
		o.R = &featureR{
			ChromosomeFeatureGenotype: related,
		}
	} else {
		o.R.ChromosomeFeatureGenotype = related
	}

	if related.R == nil {
		related.R = &featureGenotypeR{
			Chromosome: o,
		}
	} else {
		related.R.Chromosome = o
	}
	return nil
}

// RemoveChromosomeFeatureGenotype relationship.
// Sets o.R.ChromosomeFeatureGenotype to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Feature) RemoveChromosomeFeatureGenotype(exec boil.Executor, related *FeatureGenotype) error {
	var err error

	related.ChromosomeID.Valid = false
	if err = related.Update(exec, "chromosome_id"); err != nil {
		related.ChromosomeID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.ChromosomeFeatureGenotype = nil
	if related == nil || related.R == nil {
		return nil
	}

	related.R.Chromosome = nil
	return nil
}

// SetFeatureGenotype of the feature to the related item.
// Sets o.R.FeatureGenotype to related.
// Adds o to related.R.Feature.
func (o *Feature) SetFeatureGenotype(exec boil.Executor, insert bool, related *FeatureGenotype) error {
	var err error

	if insert {
		related.FeatureID = o.FeatureID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_genotype\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureGenotypePrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureID, related.FeatureGenotypeID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureID = o.FeatureID

	}

	if o.R == nil {
		o.R = &featureR{
			FeatureGenotype: related,
		}
	} else {
		o.R.FeatureGenotype = related
	}

	if related.R == nil {
		related.R = &featureGenotypeR{
			Feature: o,
		}
	} else {
		related.R.Feature = o
	}
	return nil
}

// SetFeaturePub of the feature to the related item.
// Sets o.R.FeaturePub to related.
// Adds o to related.R.Feature.
func (o *Feature) SetFeaturePub(exec boil.Executor, insert bool, related *FeaturePub) error {
	var err error

	if insert {
		related.FeatureID = o.FeatureID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurePubPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureID, related.FeaturePubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureID = o.FeatureID

	}

	if o.R == nil {
		o.R = &featureR{
			FeaturePub: related,
		}
	} else {
		o.R.FeaturePub = related
	}

	if related.R == nil {
		related.R = &featurePubR{
			Feature: o,
		}
	} else {
		related.R.Feature = o
	}
	return nil
}

// SetFeatureSynonym of the feature to the related item.
// Sets o.R.FeatureSynonym to related.
// Adds o to related.R.Feature.
func (o *Feature) SetFeatureSynonym(exec boil.Executor, insert bool, related *FeatureSynonym) error {
	var err error

	if insert {
		related.FeatureID = o.FeatureID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_synonym\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureSynonymPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureID, related.FeatureSynonymID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureID = o.FeatureID

	}

	if o.R == nil {
		o.R = &featureR{
			FeatureSynonym: related,
		}
	} else {
		o.R.FeatureSynonym = related
	}

	if related.R == nil {
		related.R = &featureSynonymR{
			Feature: o,
		}
	} else {
		related.R.Feature = o
	}
	return nil
}

// SetObjectFeatureRelationship of the feature to the related item.
// Sets o.R.ObjectFeatureRelationship to related.
// Adds o to related.R.Object.
func (o *Feature) SetObjectFeatureRelationship(exec boil.Executor, insert bool, related *FeatureRelationship) error {
	var err error

	if insert {
		related.ObjectID = o.FeatureID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"object_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureID, related.FeatureRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.ObjectID = o.FeatureID

	}

	if o.R == nil {
		o.R = &featureR{
			ObjectFeatureRelationship: related,
		}
	} else {
		o.R.ObjectFeatureRelationship = related
	}

	if related.R == nil {
		related.R = &featureRelationshipR{
			Object: o,
		}
	} else {
		related.R.Object = o
	}
	return nil
}

// SetSubjectFeatureRelationship of the feature to the related item.
// Sets o.R.SubjectFeatureRelationship to related.
// Adds o to related.R.Subject.
func (o *Feature) SetSubjectFeatureRelationship(exec boil.Executor, insert bool, related *FeatureRelationship) error {
	var err error

	if insert {
		related.SubjectID = o.FeatureID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"subject_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureID, related.FeatureRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.SubjectID = o.FeatureID

	}

	if o.R == nil {
		o.R = &featureR{
			SubjectFeatureRelationship: related,
		}
	} else {
		o.R.SubjectFeatureRelationship = related
	}

	if related.R == nil {
		related.R = &featureRelationshipR{
			Subject: o,
		}
	} else {
		related.R.Subject = o
	}
	return nil
}

// SetFeatureloc of the feature to the related item.
// Sets o.R.Featureloc to related.
// Adds o to related.R.Feature.
func (o *Feature) SetFeatureloc(exec boil.Executor, insert bool, related *Featureloc) error {
	var err error

	if insert {
		related.FeatureID = o.FeatureID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"featureloc\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurelocPrimaryKeyColumns),
		)
		values := []interface{}{o.FeatureID, related.FeaturelocID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.FeatureID = o.FeatureID

	}

	if o.R == nil {
		o.R = &featureR{
			Featureloc: related,
		}
	} else {
		o.R.Featureloc = related
	}

	if related.R == nil {
		related.R = &featurelocR{
			Feature: o,
		}
	} else {
		related.R.Feature = o
	}
	return nil
}

// AddSrcfeatureFeaturelocs adds the given related objects to the existing relationships
// of the feature, optionally inserting them as new records.
// Appends related to o.R.SrcfeatureFeaturelocs.
// Sets related.R.Srcfeature appropriately.
func (o *Feature) AddSrcfeatureFeaturelocs(exec boil.Executor, insert bool, related ...*Featureloc) error {
	var err error
	for _, rel := range related {
		rel.SrcfeatureID.Int = o.FeatureID
		rel.SrcfeatureID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "srcfeature_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &featureR{
			SrcfeatureFeaturelocs: related,
		}
	} else {
		o.R.SrcfeatureFeaturelocs = append(o.R.SrcfeatureFeaturelocs, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &featurelocR{
				Srcfeature: o,
			}
		} else {
			rel.R.Srcfeature = o
		}
	}
	return nil
}

// SetSrcfeatureFeaturelocs removes all previously related items of the
// feature replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Srcfeature's SrcfeatureFeaturelocs accordingly.
// Replaces o.R.SrcfeatureFeaturelocs with related.
// Sets related.R.Srcfeature's SrcfeatureFeaturelocs accordingly.
func (o *Feature) SetSrcfeatureFeaturelocs(exec boil.Executor, insert bool, related ...*Featureloc) error {
	query := "update \"featureloc\" set \"srcfeature_id\" = null where \"srcfeature_id\" = $1"
	values := []interface{}{o.FeatureID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.SrcfeatureFeaturelocs {
			rel.SrcfeatureID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Srcfeature = nil
		}

		o.R.SrcfeatureFeaturelocs = nil
	}
	return o.AddSrcfeatureFeaturelocs(exec, insert, related...)
}

// RemoveSrcfeatureFeaturelocs relationships from objects passed in.
// Removes related items from R.SrcfeatureFeaturelocs (uses pointer comparison, removal does not keep order)
// Sets related.R.Srcfeature.
func (o *Feature) RemoveSrcfeatureFeaturelocs(exec boil.Executor, related ...*Featureloc) error {
	var err error
	for _, rel := range related {
		rel.SrcfeatureID.Valid = false
		if rel.R != nil {
			rel.R.Srcfeature = nil
		}
		if err = rel.Update(exec, "srcfeature_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.SrcfeatureFeaturelocs {
			if rel != ri {
				continue
			}

			ln := len(o.R.SrcfeatureFeaturelocs)
			if ln > 1 && i < ln-1 {
				o.R.SrcfeatureFeaturelocs[i] = o.R.SrcfeatureFeaturelocs[ln-1]
			}
			o.R.SrcfeatureFeaturelocs = o.R.SrcfeatureFeaturelocs[:ln-1]
			break
		}
	}

	return nil
}

// FeaturesG retrieves all records.
func FeaturesG(mods ...qm.QueryMod) featureQuery {
	return Features(boil.GetDB(), mods...)
}

// Features retrieves all the records using an executor.
func Features(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	mods = append(mods, qm.From("\"feature\""))
	return featureQuery{NewQuery(exec, mods...)}
}

// FindFeatureG retrieves a single record by ID.
func FindFeatureG(featureID int, selectCols ...string) (*Feature, error) {
	return FindFeature(boil.GetDB(), featureID, selectCols...)
}

// FindFeatureGP retrieves a single record by ID, and panics on error.
func FindFeatureGP(featureID int, selectCols ...string) *Feature {
	retobj, err := FindFeature(boil.GetDB(), featureID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindFeature retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeature(exec boil.Executor, featureID int, selectCols ...string) (*Feature, error) {
	featureObj := &Feature{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"feature\" where \"feature_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, featureID)

	err := q.Bind(featureObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from feature")
	}

	return featureObj, nil
}

// FindFeatureP retrieves a single record by ID with an executor, and panics on error.
func FindFeatureP(exec boil.Executor, featureID int, selectCols ...string) *Feature {
	retobj, err := FindFeature(exec, featureID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Feature) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Feature) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Feature) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Feature) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	featureInsertCacheMut.RLock()
	cache, cached := featureInsertCache[key]
	featureInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			featureColumns,
			featureColumnsWithDefault,
			featureColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(featureType, featureMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featureType, featureMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"feature\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into feature")
	}

	if !cached {
		featureInsertCacheMut.Lock()
		featureInsertCache[key] = cache
		featureInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Feature record. See Update for
// whitelist behavior description.
func (o *Feature) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Feature record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Feature) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Feature, and panics on error.
// See Update for whitelist behavior description.
func (o *Feature) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Feature.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Feature) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	featureUpdateCacheMut.RLock()
	cache, cached := featureUpdateCache[key]
	featureUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(featureColumns, featurePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update feature, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"feature\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, featurePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featureType, featureMapping, append(wl, featurePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update feature row")
	}

	if !cached {
		featureUpdateCacheMut.Lock()
		featureUpdateCache[key] = cache
		featureUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q featureQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q featureQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for feature")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o FeatureSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o FeatureSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o FeatureSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FeatureSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"feature\" SET %s WHERE (\"feature_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurePrimaryKeyColumns), len(colNames)+1, len(featurePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in feature slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Feature) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Feature) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Feature) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Feature) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no feature provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(featureColumnsWithDefault, o)

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

	featureUpsertCacheMut.RLock()
	cache, cached := featureUpsertCache[key]
	featureUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			featureColumns,
			featureColumnsWithDefault,
			featureColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			featureColumns,
			featurePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert feature, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(featurePrimaryKeyColumns))
			copy(conflict, featurePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"feature\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(featureType, featureMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featureType, featureMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for feature")
	}

	if !cached {
		featureUpsertCacheMut.Lock()
		featureUpsertCache[key] = cache
		featureUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Feature record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Feature) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Feature record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Feature) DeleteG() error {
	if o == nil {
		return errors.New("models: no Feature provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Feature record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Feature) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Feature record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Feature) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Feature provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featurePrimaryKeyMapping)
	sql := "DELETE FROM \"feature\" WHERE \"feature_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from feature")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q featureQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q featureQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no featureQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from feature")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o FeatureSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o FeatureSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Feature slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o FeatureSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FeatureSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Feature slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(featureBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"feature\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(featurePrimaryKeyColumns), 1, len(featurePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from feature slice")
	}

	if len(featureAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Feature) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Feature) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Feature) ReloadG() error {
	if o == nil {
		return errors.New("models: no Feature provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Feature) Reload(exec boil.Executor) error {
	ret, err := FindFeature(exec, o.FeatureID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *FeatureSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty FeatureSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FeatureSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	features := FeatureSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featurePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"feature\".* FROM \"feature\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, featurePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(featurePrimaryKeyColumns), 1, len(featurePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&features)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FeatureSlice")
	}

	*o = features

	return nil
}

// FeatureExists checks if the Feature row exists.
func FeatureExists(exec boil.Executor, featureID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"feature\" where \"feature_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, featureID)
	}

	row := exec.QueryRow(sql, featureID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if feature exists")
	}

	return exists, nil
}

// FeatureExistsG checks if the Feature row exists.
func FeatureExistsG(featureID int) (bool, error) {
	return FeatureExists(boil.GetDB(), featureID)
}

// FeatureExistsGP checks if the Feature row exists. Panics on error.
func FeatureExistsGP(featureID int) bool {
	e, err := FeatureExists(boil.GetDB(), featureID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// FeatureExistsP checks if the Feature row exists. Panics on error.
func FeatureExistsP(exec boil.Executor, featureID int) bool {
	e, err := FeatureExists(exec, featureID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
