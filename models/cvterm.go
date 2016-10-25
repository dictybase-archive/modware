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

// Cvterm is an object representing the database table.
type Cvterm struct {
	CvtermID           int         `boil:"cvterm_id" json:"cvterm_id" toml:"cvterm_id" yaml:"cvterm_id"`
	CVID               int         `boil:"cv_id" json:"cv_id" toml:"cv_id" yaml:"cv_id"`
	Name               string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Definition         null.String `boil:"definition" json:"definition,omitempty" toml:"definition" yaml:"definition,omitempty"`
	DbxrefID           int         `boil:"dbxref_id" json:"dbxref_id" toml:"dbxref_id" yaml:"dbxref_id"`
	IsObsolete         int         `boil:"is_obsolete" json:"is_obsolete" toml:"is_obsolete" yaml:"is_obsolete"`
	IsRelationshiptype int         `boil:"is_relationshiptype" json:"is_relationshiptype" toml:"is_relationshiptype" yaml:"is_relationshiptype"`

	R *cvtermR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cvtermL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// cvtermR is where relationships are stored.
type cvtermR struct {
	CV                          *CV
	Dbxref                      *Dbxref
	TypeSynonym                 *Synonym
	TypePhenotypeprop           *Phenotypeprop
	TypeStock                   *Stock
	TypePubRelationship         *PubRelationship
	TypePubprop                 *Pubprop
	Cvtermsynonym               *Cvtermsynonym
	TypeDbxrefprop              *Dbxrefprop
	TypeFeature                 *Feature
	TypeStockDbxrefprop         *StockDbxrefprop
	EnvironmentCvterm           *EnvironmentCvterm
	StockCvterm                 *StockCvterm
	TypeStockCvtermprop         *StockCvtermprop
	TypeStockcollection         *Stockcollection
	TypeAnalysisfeatureprop     *Analysisfeatureprop
	CvtermDbxref                *CvtermDbxref
	TypeFeatureprop             *Featureprop
	TypeAnalysisprop            *Analysisprop
	FeatureCvterm               *FeatureCvterm
	TypeChadoprop               *Chadoprop
	TypeFeatureCvtermprop       *FeatureCvtermprop
	FeatureGenotype             *FeatureGenotype
	TypeContactRelationship     *ContactRelationship
	TypeFeaturePubprop          *FeaturePubprop
	TypeCvprop                  *Cvprop
	TypeStockRelationship       *StockRelationship
	TypeFeatureRelationshipprop *FeatureRelationshipprop
	TypeFeatureRelationship     *FeatureRelationship
	TypeGenotypeprop            *Genotypeprop
	TypeOrganismprop            *Organismprop
	TypePhendesc                *Phendesc
	TypePhenstatement           *Phenstatement
	TypeStockcollectionprop     *Stockcollectionprop
	TypeStockprop               *Stockprop
	ObjectCvtermRelationship    *CvtermRelationship
	SubjectCvtermRelationship   *CvtermRelationship
	TypeCvtermRelationship      *CvtermRelationship
	PhenotypeCvterm             *PhenotypeCvterm
	PhenotypeComparisonCvterm   *PhenotypeComparisonCvterm
	ObjectCvtermpath            *Cvtermpath
	SubjectCvtermpath           *Cvtermpath
	TypeCvtermpath              *Cvtermpath
	TypeUserRelationship        *UserRelationship
	ObjectUserRelationship      *UserRelationship
	SubjectUserRelationship     *UserRelationship
	Cvtermprop                  *Cvtermprop
	TypeCvtermprop              *Cvtermprop
	AssayPhenotypes             PhenotypeSlice
	AttrPhenotypes              PhenotypeSlice
	CvaluePhenotypes            PhenotypeSlice
	ObservablePhenotypes        PhenotypeSlice
	TypeCvtermsynonyms          CvtermsynonymSlice
	TypeJbrowseTracks           JbrowseTrackSlice
	TypePubs                    PubSlice
	StockRelationshipCvterms    StockRelationshipCvtermSlice
	TypeContacts                ContactSlice
	TypeGenotypes               GenotypeSlice
}

// cvtermL is where Load methods for each relationship are stored.
type cvtermL struct{}

var (
	cvtermColumns               = []string{"cvterm_id", "cv_id", "name", "definition", "dbxref_id", "is_obsolete", "is_relationshiptype"}
	cvtermColumnsWithoutDefault = []string{"cv_id", "name", "definition", "dbxref_id"}
	cvtermColumnsWithDefault    = []string{"cvterm_id", "is_obsolete", "is_relationshiptype"}
	cvtermPrimaryKeyColumns     = []string{"cvterm_id"}
)

type (
	// CvtermSlice is an alias for a slice of pointers to Cvterm.
	// This should generally be used opposed to []Cvterm.
	CvtermSlice []*Cvterm
	// CvtermHook is the signature for custom Cvterm hook methods
	CvtermHook func(boil.Executor, *Cvterm) error

	cvtermQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cvtermType                 = reflect.TypeOf(&Cvterm{})
	cvtermMapping              = queries.MakeStructMapping(cvtermType)
	cvtermPrimaryKeyMapping, _ = queries.BindMapping(cvtermType, cvtermMapping, cvtermPrimaryKeyColumns)
	cvtermInsertCacheMut       sync.RWMutex
	cvtermInsertCache          = make(map[string]insertCache)
	cvtermUpdateCacheMut       sync.RWMutex
	cvtermUpdateCache          = make(map[string]updateCache)
	cvtermUpsertCacheMut       sync.RWMutex
	cvtermUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var cvtermBeforeInsertHooks []CvtermHook
var cvtermBeforeUpdateHooks []CvtermHook
var cvtermBeforeDeleteHooks []CvtermHook
var cvtermBeforeUpsertHooks []CvtermHook

var cvtermAfterInsertHooks []CvtermHook
var cvtermAfterSelectHooks []CvtermHook
var cvtermAfterUpdateHooks []CvtermHook
var cvtermAfterDeleteHooks []CvtermHook
var cvtermAfterUpsertHooks []CvtermHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Cvterm) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Cvterm) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Cvterm) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Cvterm) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Cvterm) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Cvterm) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Cvterm) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Cvterm) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Cvterm) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCvtermHook registers your hook function for all future operations.
func AddCvtermHook(hookPoint boil.HookPoint, cvtermHook CvtermHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cvtermBeforeInsertHooks = append(cvtermBeforeInsertHooks, cvtermHook)
	case boil.BeforeUpdateHook:
		cvtermBeforeUpdateHooks = append(cvtermBeforeUpdateHooks, cvtermHook)
	case boil.BeforeDeleteHook:
		cvtermBeforeDeleteHooks = append(cvtermBeforeDeleteHooks, cvtermHook)
	case boil.BeforeUpsertHook:
		cvtermBeforeUpsertHooks = append(cvtermBeforeUpsertHooks, cvtermHook)
	case boil.AfterInsertHook:
		cvtermAfterInsertHooks = append(cvtermAfterInsertHooks, cvtermHook)
	case boil.AfterSelectHook:
		cvtermAfterSelectHooks = append(cvtermAfterSelectHooks, cvtermHook)
	case boil.AfterUpdateHook:
		cvtermAfterUpdateHooks = append(cvtermAfterUpdateHooks, cvtermHook)
	case boil.AfterDeleteHook:
		cvtermAfterDeleteHooks = append(cvtermAfterDeleteHooks, cvtermHook)
	case boil.AfterUpsertHook:
		cvtermAfterUpsertHooks = append(cvtermAfterUpsertHooks, cvtermHook)
	}
}

// OneP returns a single cvterm record from the query, and panics on error.
func (q cvtermQuery) OneP() *Cvterm {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single cvterm record from the query.
func (q cvtermQuery) One() (*Cvterm, error) {
	o := &Cvterm{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cvterm")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Cvterm records from the query, and panics on error.
func (q cvtermQuery) AllP() CvtermSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Cvterm records from the query.
func (q cvtermQuery) All() (CvtermSlice, error) {
	var o CvtermSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Cvterm slice")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Cvterm records in the query, and panics on error.
func (q cvtermQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Cvterm records in the query.
func (q cvtermQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cvterm rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q cvtermQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q cvtermQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cvterm exists")
	}

	return count > 0, nil
}

// CVG pointed to by the foreign key.
func (o *Cvterm) CVG(mods ...qm.QueryMod) cvQuery {
	return o.CV(boil.GetDB(), mods...)
}

// CV pointed to by the foreign key.
func (o *Cvterm) CV(exec boil.Executor, mods ...qm.QueryMod) cvQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cv_id=$1", o.CVID),
	}

	queryMods = append(queryMods, mods...)

	query := CVS(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cv\"")

	return query
}

// DbxrefG pointed to by the foreign key.
func (o *Cvterm) DbxrefG(mods ...qm.QueryMod) dbxrefQuery {
	return o.Dbxref(boil.GetDB(), mods...)
}

// Dbxref pointed to by the foreign key.
func (o *Cvterm) Dbxref(exec boil.Executor, mods ...qm.QueryMod) dbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("dbxref_id=$1", o.DbxrefID),
	}

	queryMods = append(queryMods, mods...)

	query := Dbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"dbxref\"")

	return query
}

// TypeSynonymG pointed to by the foreign key.
func (o *Cvterm) TypeSynonymG(mods ...qm.QueryMod) synonymQuery {
	return o.TypeSynonym(boil.GetDB(), mods...)
}

// TypeSynonym pointed to by the foreign key.
func (o *Cvterm) TypeSynonym(exec boil.Executor, mods ...qm.QueryMod) synonymQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Synonyms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"synonym\"")

	return query
}

// TypePhenotypepropG pointed to by the foreign key.
func (o *Cvterm) TypePhenotypepropG(mods ...qm.QueryMod) phenotypepropQuery {
	return o.TypePhenotypeprop(boil.GetDB(), mods...)
}

// TypePhenotypeprop pointed to by the foreign key.
func (o *Cvterm) TypePhenotypeprop(exec boil.Executor, mods ...qm.QueryMod) phenotypepropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Phenotypeprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotypeprop\"")

	return query
}

// TypeStockG pointed to by the foreign key.
func (o *Cvterm) TypeStockG(mods ...qm.QueryMod) stockQuery {
	return o.TypeStock(boil.GetDB(), mods...)
}

// TypeStock pointed to by the foreign key.
func (o *Cvterm) TypeStock(exec boil.Executor, mods ...qm.QueryMod) stockQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Stocks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock\"")

	return query
}

// TypePubRelationshipG pointed to by the foreign key.
func (o *Cvterm) TypePubRelationshipG(mods ...qm.QueryMod) pubRelationshipQuery {
	return o.TypePubRelationship(boil.GetDB(), mods...)
}

// TypePubRelationship pointed to by the foreign key.
func (o *Cvterm) TypePubRelationship(exec boil.Executor, mods ...qm.QueryMod) pubRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := PubRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub_relationship\"")

	return query
}

// TypePubpropG pointed to by the foreign key.
func (o *Cvterm) TypePubpropG(mods ...qm.QueryMod) pubpropQuery {
	return o.TypePubprop(boil.GetDB(), mods...)
}

// TypePubprop pointed to by the foreign key.
func (o *Cvterm) TypePubprop(exec boil.Executor, mods ...qm.QueryMod) pubpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pubprop\"")

	return query
}

// CvtermsynonymG pointed to by the foreign key.
func (o *Cvterm) CvtermsynonymG(mods ...qm.QueryMod) cvtermsynonymQuery {
	return o.Cvtermsynonym(boil.GetDB(), mods...)
}

// Cvtermsynonym pointed to by the foreign key.
func (o *Cvterm) Cvtermsynonym(exec boil.Executor, mods ...qm.QueryMod) cvtermsynonymQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvtermsynonyms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvtermsynonym\"")

	return query
}

// TypeDbxrefpropG pointed to by the foreign key.
func (o *Cvterm) TypeDbxrefpropG(mods ...qm.QueryMod) dbxrefpropQuery {
	return o.TypeDbxrefprop(boil.GetDB(), mods...)
}

// TypeDbxrefprop pointed to by the foreign key.
func (o *Cvterm) TypeDbxrefprop(exec boil.Executor, mods ...qm.QueryMod) dbxrefpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Dbxrefprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"dbxrefprop\"")

	return query
}

// TypeFeatureG pointed to by the foreign key.
func (o *Cvterm) TypeFeatureG(mods ...qm.QueryMod) featureQuery {
	return o.TypeFeature(boil.GetDB(), mods...)
}

// TypeFeature pointed to by the foreign key.
func (o *Cvterm) TypeFeature(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// TypeStockDbxrefpropG pointed to by the foreign key.
func (o *Cvterm) TypeStockDbxrefpropG(mods ...qm.QueryMod) stockDbxrefpropQuery {
	return o.TypeStockDbxrefprop(boil.GetDB(), mods...)
}

// TypeStockDbxrefprop pointed to by the foreign key.
func (o *Cvterm) TypeStockDbxrefprop(exec boil.Executor, mods ...qm.QueryMod) stockDbxrefpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := StockDbxrefprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_dbxrefprop\"")

	return query
}

// EnvironmentCvtermG pointed to by the foreign key.
func (o *Cvterm) EnvironmentCvtermG(mods ...qm.QueryMod) environmentCvtermQuery {
	return o.EnvironmentCvterm(boil.GetDB(), mods...)
}

// EnvironmentCvterm pointed to by the foreign key.
func (o *Cvterm) EnvironmentCvterm(exec boil.Executor, mods ...qm.QueryMod) environmentCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := EnvironmentCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"environment_cvterm\"")

	return query
}

// StockCvtermG pointed to by the foreign key.
func (o *Cvterm) StockCvtermG(mods ...qm.QueryMod) stockCvtermQuery {
	return o.StockCvterm(boil.GetDB(), mods...)
}

// StockCvterm pointed to by the foreign key.
func (o *Cvterm) StockCvterm(exec boil.Executor, mods ...qm.QueryMod) stockCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := StockCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_cvterm\"")

	return query
}

// TypeStockCvtermpropG pointed to by the foreign key.
func (o *Cvterm) TypeStockCvtermpropG(mods ...qm.QueryMod) stockCvtermpropQuery {
	return o.TypeStockCvtermprop(boil.GetDB(), mods...)
}

// TypeStockCvtermprop pointed to by the foreign key.
func (o *Cvterm) TypeStockCvtermprop(exec boil.Executor, mods ...qm.QueryMod) stockCvtermpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := StockCvtermprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_cvtermprop\"")

	return query
}

// TypeStockcollectionG pointed to by the foreign key.
func (o *Cvterm) TypeStockcollectionG(mods ...qm.QueryMod) stockcollectionQuery {
	return o.TypeStockcollection(boil.GetDB(), mods...)
}

// TypeStockcollection pointed to by the foreign key.
func (o *Cvterm) TypeStockcollection(exec boil.Executor, mods ...qm.QueryMod) stockcollectionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Stockcollections(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stockcollection\"")

	return query
}

// TypeAnalysisfeaturepropG pointed to by the foreign key.
func (o *Cvterm) TypeAnalysisfeaturepropG(mods ...qm.QueryMod) analysisfeaturepropQuery {
	return o.TypeAnalysisfeatureprop(boil.GetDB(), mods...)
}

// TypeAnalysisfeatureprop pointed to by the foreign key.
func (o *Cvterm) TypeAnalysisfeatureprop(exec boil.Executor, mods ...qm.QueryMod) analysisfeaturepropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Analysisfeatureprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"analysisfeatureprop\"")

	return query
}

// CvtermDbxrefG pointed to by the foreign key.
func (o *Cvterm) CvtermDbxrefG(mods ...qm.QueryMod) cvtermDbxrefQuery {
	return o.CvtermDbxref(boil.GetDB(), mods...)
}

// CvtermDbxref pointed to by the foreign key.
func (o *Cvterm) CvtermDbxref(exec boil.Executor, mods ...qm.QueryMod) cvtermDbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := CvtermDbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm_dbxref\"")

	return query
}

// TypeFeaturepropG pointed to by the foreign key.
func (o *Cvterm) TypeFeaturepropG(mods ...qm.QueryMod) featurepropQuery {
	return o.TypeFeatureprop(boil.GetDB(), mods...)
}

// TypeFeatureprop pointed to by the foreign key.
func (o *Cvterm) TypeFeatureprop(exec boil.Executor, mods ...qm.QueryMod) featurepropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Featureprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"featureprop\"")

	return query
}

// TypeAnalysispropG pointed to by the foreign key.
func (o *Cvterm) TypeAnalysispropG(mods ...qm.QueryMod) analysispropQuery {
	return o.TypeAnalysisprop(boil.GetDB(), mods...)
}

// TypeAnalysisprop pointed to by the foreign key.
func (o *Cvterm) TypeAnalysisprop(exec boil.Executor, mods ...qm.QueryMod) analysispropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Analysisprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"analysisprop\"")

	return query
}

// FeatureCvtermG pointed to by the foreign key.
func (o *Cvterm) FeatureCvtermG(mods ...qm.QueryMod) featureCvtermQuery {
	return o.FeatureCvterm(boil.GetDB(), mods...)
}

// FeatureCvterm pointed to by the foreign key.
func (o *Cvterm) FeatureCvterm(exec boil.Executor, mods ...qm.QueryMod) featureCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_cvterm\"")

	return query
}

// TypeChadopropG pointed to by the foreign key.
func (o *Cvterm) TypeChadopropG(mods ...qm.QueryMod) chadopropQuery {
	return o.TypeChadoprop(boil.GetDB(), mods...)
}

// TypeChadoprop pointed to by the foreign key.
func (o *Cvterm) TypeChadoprop(exec boil.Executor, mods ...qm.QueryMod) chadopropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Chadoprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"chadoprop\"")

	return query
}

// TypeFeatureCvtermpropG pointed to by the foreign key.
func (o *Cvterm) TypeFeatureCvtermpropG(mods ...qm.QueryMod) featureCvtermpropQuery {
	return o.TypeFeatureCvtermprop(boil.GetDB(), mods...)
}

// TypeFeatureCvtermprop pointed to by the foreign key.
func (o *Cvterm) TypeFeatureCvtermprop(exec boil.Executor, mods ...qm.QueryMod) featureCvtermpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureCvtermprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_cvtermprop\"")

	return query
}

// FeatureGenotypeG pointed to by the foreign key.
func (o *Cvterm) FeatureGenotypeG(mods ...qm.QueryMod) featureGenotypeQuery {
	return o.FeatureGenotype(boil.GetDB(), mods...)
}

// FeatureGenotype pointed to by the foreign key.
func (o *Cvterm) FeatureGenotype(exec boil.Executor, mods ...qm.QueryMod) featureGenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureGenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_genotype\"")

	return query
}

// TypeContactRelationshipG pointed to by the foreign key.
func (o *Cvterm) TypeContactRelationshipG(mods ...qm.QueryMod) contactRelationshipQuery {
	return o.TypeContactRelationship(boil.GetDB(), mods...)
}

// TypeContactRelationship pointed to by the foreign key.
func (o *Cvterm) TypeContactRelationship(exec boil.Executor, mods ...qm.QueryMod) contactRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := ContactRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"contact_relationship\"")

	return query
}

// TypeFeaturePubpropG pointed to by the foreign key.
func (o *Cvterm) TypeFeaturePubpropG(mods ...qm.QueryMod) featurePubpropQuery {
	return o.TypeFeaturePubprop(boil.GetDB(), mods...)
}

// TypeFeaturePubprop pointed to by the foreign key.
func (o *Cvterm) TypeFeaturePubprop(exec boil.Executor, mods ...qm.QueryMod) featurePubpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := FeaturePubprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_pubprop\"")

	return query
}

// TypeCvpropG pointed to by the foreign key.
func (o *Cvterm) TypeCvpropG(mods ...qm.QueryMod) cvpropQuery {
	return o.TypeCvprop(boil.GetDB(), mods...)
}

// TypeCvprop pointed to by the foreign key.
func (o *Cvterm) TypeCvprop(exec boil.Executor, mods ...qm.QueryMod) cvpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvprop\"")

	return query
}

// TypeStockRelationshipG pointed to by the foreign key.
func (o *Cvterm) TypeStockRelationshipG(mods ...qm.QueryMod) stockRelationshipQuery {
	return o.TypeStockRelationship(boil.GetDB(), mods...)
}

// TypeStockRelationship pointed to by the foreign key.
func (o *Cvterm) TypeStockRelationship(exec boil.Executor, mods ...qm.QueryMod) stockRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := StockRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_relationship\"")

	return query
}

// TypeFeatureRelationshippropG pointed to by the foreign key.
func (o *Cvterm) TypeFeatureRelationshippropG(mods ...qm.QueryMod) featureRelationshippropQuery {
	return o.TypeFeatureRelationshipprop(boil.GetDB(), mods...)
}

// TypeFeatureRelationshipprop pointed to by the foreign key.
func (o *Cvterm) TypeFeatureRelationshipprop(exec boil.Executor, mods ...qm.QueryMod) featureRelationshippropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureRelationshipprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_relationshipprop\"")

	return query
}

// TypeFeatureRelationshipG pointed to by the foreign key.
func (o *Cvterm) TypeFeatureRelationshipG(mods ...qm.QueryMod) featureRelationshipQuery {
	return o.TypeFeatureRelationship(boil.GetDB(), mods...)
}

// TypeFeatureRelationship pointed to by the foreign key.
func (o *Cvterm) TypeFeatureRelationship(exec boil.Executor, mods ...qm.QueryMod) featureRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_relationship\"")

	return query
}

// TypeGenotypepropG pointed to by the foreign key.
func (o *Cvterm) TypeGenotypepropG(mods ...qm.QueryMod) genotypepropQuery {
	return o.TypeGenotypeprop(boil.GetDB(), mods...)
}

// TypeGenotypeprop pointed to by the foreign key.
func (o *Cvterm) TypeGenotypeprop(exec boil.Executor, mods ...qm.QueryMod) genotypepropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Genotypeprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"genotypeprop\"")

	return query
}

// TypeOrganismpropG pointed to by the foreign key.
func (o *Cvterm) TypeOrganismpropG(mods ...qm.QueryMod) organismpropQuery {
	return o.TypeOrganismprop(boil.GetDB(), mods...)
}

// TypeOrganismprop pointed to by the foreign key.
func (o *Cvterm) TypeOrganismprop(exec boil.Executor, mods ...qm.QueryMod) organismpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Organismprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"organismprop\"")

	return query
}

// TypePhendescG pointed to by the foreign key.
func (o *Cvterm) TypePhendescG(mods ...qm.QueryMod) phendescQuery {
	return o.TypePhendesc(boil.GetDB(), mods...)
}

// TypePhendesc pointed to by the foreign key.
func (o *Cvterm) TypePhendesc(exec boil.Executor, mods ...qm.QueryMod) phendescQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Phendescs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phendesc\"")

	return query
}

// TypePhenstatementG pointed to by the foreign key.
func (o *Cvterm) TypePhenstatementG(mods ...qm.QueryMod) phenstatementQuery {
	return o.TypePhenstatement(boil.GetDB(), mods...)
}

// TypePhenstatement pointed to by the foreign key.
func (o *Cvterm) TypePhenstatement(exec boil.Executor, mods ...qm.QueryMod) phenstatementQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Phenstatements(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenstatement\"")

	return query
}

// TypeStockcollectionpropG pointed to by the foreign key.
func (o *Cvterm) TypeStockcollectionpropG(mods ...qm.QueryMod) stockcollectionpropQuery {
	return o.TypeStockcollectionprop(boil.GetDB(), mods...)
}

// TypeStockcollectionprop pointed to by the foreign key.
func (o *Cvterm) TypeStockcollectionprop(exec boil.Executor, mods ...qm.QueryMod) stockcollectionpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Stockcollectionprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stockcollectionprop\"")

	return query
}

// TypeStockpropG pointed to by the foreign key.
func (o *Cvterm) TypeStockpropG(mods ...qm.QueryMod) stockpropQuery {
	return o.TypeStockprop(boil.GetDB(), mods...)
}

// TypeStockprop pointed to by the foreign key.
func (o *Cvterm) TypeStockprop(exec boil.Executor, mods ...qm.QueryMod) stockpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Stockprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stockprop\"")

	return query
}

// ObjectCvtermRelationshipG pointed to by the foreign key.
func (o *Cvterm) ObjectCvtermRelationshipG(mods ...qm.QueryMod) cvtermRelationshipQuery {
	return o.ObjectCvtermRelationship(boil.GetDB(), mods...)
}

// ObjectCvtermRelationship pointed to by the foreign key.
func (o *Cvterm) ObjectCvtermRelationship(exec boil.Executor, mods ...qm.QueryMod) cvtermRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("object_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := CvtermRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm_relationship\"")

	return query
}

// SubjectCvtermRelationshipG pointed to by the foreign key.
func (o *Cvterm) SubjectCvtermRelationshipG(mods ...qm.QueryMod) cvtermRelationshipQuery {
	return o.SubjectCvtermRelationship(boil.GetDB(), mods...)
}

// SubjectCvtermRelationship pointed to by the foreign key.
func (o *Cvterm) SubjectCvtermRelationship(exec boil.Executor, mods ...qm.QueryMod) cvtermRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("subject_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := CvtermRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm_relationship\"")

	return query
}

// TypeCvtermRelationshipG pointed to by the foreign key.
func (o *Cvterm) TypeCvtermRelationshipG(mods ...qm.QueryMod) cvtermRelationshipQuery {
	return o.TypeCvtermRelationship(boil.GetDB(), mods...)
}

// TypeCvtermRelationship pointed to by the foreign key.
func (o *Cvterm) TypeCvtermRelationship(exec boil.Executor, mods ...qm.QueryMod) cvtermRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := CvtermRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm_relationship\"")

	return query
}

// PhenotypeCvtermG pointed to by the foreign key.
func (o *Cvterm) PhenotypeCvtermG(mods ...qm.QueryMod) phenotypeCvtermQuery {
	return o.PhenotypeCvterm(boil.GetDB(), mods...)
}

// PhenotypeCvterm pointed to by the foreign key.
func (o *Cvterm) PhenotypeCvterm(exec boil.Executor, mods ...qm.QueryMod) phenotypeCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := PhenotypeCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_cvterm\"")

	return query
}

// PhenotypeComparisonCvtermG pointed to by the foreign key.
func (o *Cvterm) PhenotypeComparisonCvtermG(mods ...qm.QueryMod) phenotypeComparisonCvtermQuery {
	return o.PhenotypeComparisonCvterm(boil.GetDB(), mods...)
}

// PhenotypeComparisonCvterm pointed to by the foreign key.
func (o *Cvterm) PhenotypeComparisonCvterm(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := PhenotypeComparisonCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_comparison_cvterm\"")

	return query
}

// ObjectCvtermpathG pointed to by the foreign key.
func (o *Cvterm) ObjectCvtermpathG(mods ...qm.QueryMod) cvtermpathQuery {
	return o.ObjectCvtermpath(boil.GetDB(), mods...)
}

// ObjectCvtermpath pointed to by the foreign key.
func (o *Cvterm) ObjectCvtermpath(exec boil.Executor, mods ...qm.QueryMod) cvtermpathQuery {
	queryMods := []qm.QueryMod{
		qm.Where("object_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvtermpaths(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvtermpath\"")

	return query
}

// SubjectCvtermpathG pointed to by the foreign key.
func (o *Cvterm) SubjectCvtermpathG(mods ...qm.QueryMod) cvtermpathQuery {
	return o.SubjectCvtermpath(boil.GetDB(), mods...)
}

// SubjectCvtermpath pointed to by the foreign key.
func (o *Cvterm) SubjectCvtermpath(exec boil.Executor, mods ...qm.QueryMod) cvtermpathQuery {
	queryMods := []qm.QueryMod{
		qm.Where("subject_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvtermpaths(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvtermpath\"")

	return query
}

// TypeCvtermpathG pointed to by the foreign key.
func (o *Cvterm) TypeCvtermpathG(mods ...qm.QueryMod) cvtermpathQuery {
	return o.TypeCvtermpath(boil.GetDB(), mods...)
}

// TypeCvtermpath pointed to by the foreign key.
func (o *Cvterm) TypeCvtermpath(exec boil.Executor, mods ...qm.QueryMod) cvtermpathQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvtermpaths(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvtermpath\"")

	return query
}

// TypeUserRelationshipG pointed to by the foreign key.
func (o *Cvterm) TypeUserRelationshipG(mods ...qm.QueryMod) userRelationshipQuery {
	return o.TypeUserRelationship(boil.GetDB(), mods...)
}

// TypeUserRelationship pointed to by the foreign key.
func (o *Cvterm) TypeUserRelationship(exec boil.Executor, mods ...qm.QueryMod) userRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := UserRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"user_relationship\"")

	return query
}

// ObjectUserRelationshipG pointed to by the foreign key.
func (o *Cvterm) ObjectUserRelationshipG(mods ...qm.QueryMod) userRelationshipQuery {
	return o.ObjectUserRelationship(boil.GetDB(), mods...)
}

// ObjectUserRelationship pointed to by the foreign key.
func (o *Cvterm) ObjectUserRelationship(exec boil.Executor, mods ...qm.QueryMod) userRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("object_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := UserRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"user_relationship\"")

	return query
}

// SubjectUserRelationshipG pointed to by the foreign key.
func (o *Cvterm) SubjectUserRelationshipG(mods ...qm.QueryMod) userRelationshipQuery {
	return o.SubjectUserRelationship(boil.GetDB(), mods...)
}

// SubjectUserRelationship pointed to by the foreign key.
func (o *Cvterm) SubjectUserRelationship(exec boil.Executor, mods ...qm.QueryMod) userRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("subject_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := UserRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"user_relationship\"")

	return query
}

// CvtermpropG pointed to by the foreign key.
func (o *Cvterm) CvtermpropG(mods ...qm.QueryMod) cvtermpropQuery {
	return o.Cvtermprop(boil.GetDB(), mods...)
}

// Cvtermprop pointed to by the foreign key.
func (o *Cvterm) Cvtermprop(exec boil.Executor, mods ...qm.QueryMod) cvtermpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvtermprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvtermprop\"")

	return query
}

// TypeCvtermpropG pointed to by the foreign key.
func (o *Cvterm) TypeCvtermpropG(mods ...qm.QueryMod) cvtermpropQuery {
	return o.TypeCvtermprop(boil.GetDB(), mods...)
}

// TypeCvtermprop pointed to by the foreign key.
func (o *Cvterm) TypeCvtermprop(exec boil.Executor, mods ...qm.QueryMod) cvtermpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("type_id=$1", o.CvtermID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvtermprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvtermprop\"")

	return query
}

// AssayPhenotypesG retrieves all the phenotype's phenotype via assay_id column.
func (o *Cvterm) AssayPhenotypesG(mods ...qm.QueryMod) phenotypeQuery {
	return o.AssayPhenotypes(boil.GetDB(), mods...)
}

// AssayPhenotypes retrieves all the phenotype's phenotype with an executor via assay_id column.
func (o *Cvterm) AssayPhenotypes(exec boil.Executor, mods ...qm.QueryMod) phenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"assay_id\"=$1", o.CvtermID),
	)

	query := Phenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype\" as \"a\"")
	return query
}

// AttrPhenotypesG retrieves all the phenotype's phenotype via attr_id column.
func (o *Cvterm) AttrPhenotypesG(mods ...qm.QueryMod) phenotypeQuery {
	return o.AttrPhenotypes(boil.GetDB(), mods...)
}

// AttrPhenotypes retrieves all the phenotype's phenotype with an executor via attr_id column.
func (o *Cvterm) AttrPhenotypes(exec boil.Executor, mods ...qm.QueryMod) phenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"attr_id\"=$1", o.CvtermID),
	)

	query := Phenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype\" as \"a\"")
	return query
}

// CvaluePhenotypesG retrieves all the phenotype's phenotype via cvalue_id column.
func (o *Cvterm) CvaluePhenotypesG(mods ...qm.QueryMod) phenotypeQuery {
	return o.CvaluePhenotypes(boil.GetDB(), mods...)
}

// CvaluePhenotypes retrieves all the phenotype's phenotype with an executor via cvalue_id column.
func (o *Cvterm) CvaluePhenotypes(exec boil.Executor, mods ...qm.QueryMod) phenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"cvalue_id\"=$1", o.CvtermID),
	)

	query := Phenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype\" as \"a\"")
	return query
}

// ObservablePhenotypesG retrieves all the phenotype's phenotype via observable_id column.
func (o *Cvterm) ObservablePhenotypesG(mods ...qm.QueryMod) phenotypeQuery {
	return o.ObservablePhenotypes(boil.GetDB(), mods...)
}

// ObservablePhenotypes retrieves all the phenotype's phenotype with an executor via observable_id column.
func (o *Cvterm) ObservablePhenotypes(exec boil.Executor, mods ...qm.QueryMod) phenotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"observable_id\"=$1", o.CvtermID),
	)

	query := Phenotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype\" as \"a\"")
	return query
}

// TypeCvtermsynonymsG retrieves all the cvtermsynonym's cvtermsynonym via type_id column.
func (o *Cvterm) TypeCvtermsynonymsG(mods ...qm.QueryMod) cvtermsynonymQuery {
	return o.TypeCvtermsynonyms(boil.GetDB(), mods...)
}

// TypeCvtermsynonyms retrieves all the cvtermsynonym's cvtermsynonym with an executor via type_id column.
func (o *Cvterm) TypeCvtermsynonyms(exec boil.Executor, mods ...qm.QueryMod) cvtermsynonymQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"type_id\"=$1", o.CvtermID),
	)

	query := Cvtermsynonyms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvtermsynonym\" as \"a\"")
	return query
}

// TypeJbrowseTracksG retrieves all the jbrowse_track's jbrowse track via type_id column.
func (o *Cvterm) TypeJbrowseTracksG(mods ...qm.QueryMod) jbrowseTrackQuery {
	return o.TypeJbrowseTracks(boil.GetDB(), mods...)
}

// TypeJbrowseTracks retrieves all the jbrowse_track's jbrowse track with an executor via type_id column.
func (o *Cvterm) TypeJbrowseTracks(exec boil.Executor, mods ...qm.QueryMod) jbrowseTrackQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"type_id\"=$1", o.CvtermID),
	)

	query := JbrowseTracks(exec, queryMods...)
	queries.SetFrom(query.Query, "\"jbrowse_track\" as \"a\"")
	return query
}

// TypePubsG retrieves all the pub's pub via type_id column.
func (o *Cvterm) TypePubsG(mods ...qm.QueryMod) pubQuery {
	return o.TypePubs(boil.GetDB(), mods...)
}

// TypePubs retrieves all the pub's pub with an executor via type_id column.
func (o *Cvterm) TypePubs(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"type_id\"=$1", o.CvtermID),
	)

	query := Pubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub\" as \"a\"")
	return query
}

// StockRelationshipCvtermsG retrieves all the stock_relationship_cvterm's stock relationship cvterm.
func (o *Cvterm) StockRelationshipCvtermsG(mods ...qm.QueryMod) stockRelationshipCvtermQuery {
	return o.StockRelationshipCvterms(boil.GetDB(), mods...)
}

// StockRelationshipCvterms retrieves all the stock_relationship_cvterm's stock relationship cvterm with an executor.
func (o *Cvterm) StockRelationshipCvterms(exec boil.Executor, mods ...qm.QueryMod) stockRelationshipCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"cvterm_id\"=$1", o.CvtermID),
	)

	query := StockRelationshipCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_relationship_cvterm\" as \"a\"")
	return query
}

// TypeContactsG retrieves all the contact's contact via type_id column.
func (o *Cvterm) TypeContactsG(mods ...qm.QueryMod) contactQuery {
	return o.TypeContacts(boil.GetDB(), mods...)
}

// TypeContacts retrieves all the contact's contact with an executor via type_id column.
func (o *Cvterm) TypeContacts(exec boil.Executor, mods ...qm.QueryMod) contactQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"type_id\"=$1", o.CvtermID),
	)

	query := Contacts(exec, queryMods...)
	queries.SetFrom(query.Query, "\"contact\" as \"a\"")
	return query
}

// TypeGenotypesG retrieves all the genotype's genotype via type_id column.
func (o *Cvterm) TypeGenotypesG(mods ...qm.QueryMod) genotypeQuery {
	return o.TypeGenotypes(boil.GetDB(), mods...)
}

// TypeGenotypes retrieves all the genotype's genotype with an executor via type_id column.
func (o *Cvterm) TypeGenotypes(exec boil.Executor, mods ...qm.QueryMod) genotypeQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"type_id\"=$1", o.CvtermID),
	)

	query := Genotypes(exec, queryMods...)
	queries.SetFrom(query.Query, "\"genotype\" as \"a\"")
	return query
}

// LoadCV allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadCV(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CVID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CVID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cv\" where \"cv_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CV")
	}
	defer results.Close()

	var resultSlice []*CV
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CV")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.CV = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CVID == foreign.CVID {
				local.R.CV = foreign
				break
			}
		}
	}

	return nil
}

// LoadDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadDbxref(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.DbxrefID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
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

	if len(cvtermAfterSelectHooks) != 0 {
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
			if local.DbxrefID == foreign.DbxrefID {
				local.R.Dbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeSynonym allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeSynonym(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"synonym\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Synonym")
	}
	defer results.Close()

	var resultSlice []*Synonym
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Synonym")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeSynonym = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeSynonym = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypePhenotypeprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypePhenotypeprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotypeprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Phenotypeprop")
	}
	defer results.Close()

	var resultSlice []*Phenotypeprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Phenotypeprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypePhenotypeprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypePhenotypeprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeStock allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeStock(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Stock")
	}
	defer results.Close()

	var resultSlice []*Stock
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Stock")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeStock = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeStock = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypePubRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypePubRelationship(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"pub_relationship\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PubRelationship")
	}
	defer results.Close()

	var resultSlice []*PubRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PubRelationship")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypePubRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypePubRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypePubprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypePubprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"pubprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Pubprop")
	}
	defer results.Close()

	var resultSlice []*Pubprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Pubprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypePubprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypePubprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadCvtermsynonym allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadCvtermsynonym(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvtermsynonym\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvtermsynonym")
	}
	defer results.Close()

	var resultSlice []*Cvtermsynonym
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvtermsynonym")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Cvtermsynonym = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.CvtermID {
				local.R.Cvtermsynonym = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeDbxrefprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeDbxrefprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"dbxrefprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Dbxrefprop")
	}
	defer results.Close()

	var resultSlice []*Dbxrefprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Dbxrefprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeDbxrefprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeDbxrefprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeFeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeFeature(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature\" where \"type_id\" in (%s)",
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

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeFeature = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeFeature = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeStockDbxrefprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeStockDbxrefprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_dbxrefprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockDbxrefprop")
	}
	defer results.Close()

	var resultSlice []*StockDbxrefprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockDbxrefprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeStockDbxrefprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeStockDbxrefprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadEnvironmentCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadEnvironmentCvterm(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"environment_cvterm\" where \"cvterm_id\" in (%s)",
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

	if len(cvtermAfterSelectHooks) != 0 {
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
			if local.CvtermID == foreign.CvtermID {
				local.R.EnvironmentCvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadStockCvterm(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_cvterm\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockCvterm")
	}
	defer results.Close()

	var resultSlice []*StockCvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockCvterm")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockCvterm = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.CvtermID {
				local.R.StockCvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeStockCvtermprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeStockCvtermprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_cvtermprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockCvtermprop")
	}
	defer results.Close()

	var resultSlice []*StockCvtermprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockCvtermprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeStockCvtermprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeStockCvtermprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeStockcollection allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeStockcollection(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stockcollection\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Stockcollection")
	}
	defer results.Close()

	var resultSlice []*Stockcollection
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Stockcollection")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeStockcollection = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeStockcollection = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeAnalysisfeatureprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeAnalysisfeatureprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"analysisfeatureprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Analysisfeatureprop")
	}
	defer results.Close()

	var resultSlice []*Analysisfeatureprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Analysisfeatureprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeAnalysisfeatureprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeAnalysisfeatureprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadCvtermDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadCvtermDbxref(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm_dbxref\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CvtermDbxref")
	}
	defer results.Close()

	var resultSlice []*CvtermDbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CvtermDbxref")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.CvtermDbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.CvtermID {
				local.R.CvtermDbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeFeatureprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeFeatureprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"featureprop\" where \"type_id\" in (%s)",
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

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeFeatureprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeFeatureprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeAnalysisprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeAnalysisprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"analysisprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Analysisprop")
	}
	defer results.Close()

	var resultSlice []*Analysisprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Analysisprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeAnalysisprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeAnalysisprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadFeatureCvterm(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_cvterm\" where \"cvterm_id\" in (%s)",
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

	if len(cvtermAfterSelectHooks) != 0 {
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
			if local.CvtermID == foreign.CvtermID {
				local.R.FeatureCvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeChadoprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeChadoprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"chadoprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Chadoprop")
	}
	defer results.Close()

	var resultSlice []*Chadoprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Chadoprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeChadoprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeChadoprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeFeatureCvtermprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeFeatureCvtermprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_cvtermprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureCvtermprop")
	}
	defer results.Close()

	var resultSlice []*FeatureCvtermprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureCvtermprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeFeatureCvtermprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeFeatureCvtermprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureGenotype allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadFeatureGenotype(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_genotype\" where \"cvterm_id\" in (%s)",
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

	if len(cvtermAfterSelectHooks) != 0 {
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
			if local.CvtermID == foreign.CvtermID {
				local.R.FeatureGenotype = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeContactRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeContactRelationship(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"contact_relationship\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ContactRelationship")
	}
	defer results.Close()

	var resultSlice []*ContactRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ContactRelationship")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeContactRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeContactRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeFeaturePubprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeFeaturePubprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_pubprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeaturePubprop")
	}
	defer results.Close()

	var resultSlice []*FeaturePubprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeaturePubprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeFeaturePubprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeFeaturePubprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeCvprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeCvprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvprop")
	}
	defer results.Close()

	var resultSlice []*Cvprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeCvprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeCvprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeStockRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeStockRelationship(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_relationship\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockRelationship")
	}
	defer results.Close()

	var resultSlice []*StockRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockRelationship")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeStockRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeStockRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeFeatureRelationshipprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeFeatureRelationshipprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_relationshipprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureRelationshipprop")
	}
	defer results.Close()

	var resultSlice []*FeatureRelationshipprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureRelationshipprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeFeatureRelationshipprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeFeatureRelationshipprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeFeatureRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeFeatureRelationship(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_relationship\" where \"type_id\" in (%s)",
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

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeFeatureRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeFeatureRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeGenotypeprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeGenotypeprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"genotypeprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Genotypeprop")
	}
	defer results.Close()

	var resultSlice []*Genotypeprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Genotypeprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeGenotypeprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeGenotypeprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeOrganismprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeOrganismprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"organismprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Organismprop")
	}
	defer results.Close()

	var resultSlice []*Organismprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Organismprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeOrganismprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeOrganismprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypePhendesc allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypePhendesc(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phendesc\" where \"type_id\" in (%s)",
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

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypePhendesc = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypePhendesc = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypePhenstatement allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypePhenstatement(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenstatement\" where \"type_id\" in (%s)",
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

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypePhenstatement = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypePhenstatement = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeStockcollectionprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeStockcollectionprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stockcollectionprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Stockcollectionprop")
	}
	defer results.Close()

	var resultSlice []*Stockcollectionprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Stockcollectionprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeStockcollectionprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeStockcollectionprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeStockprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeStockprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stockprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Stockprop")
	}
	defer results.Close()

	var resultSlice []*Stockprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Stockprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeStockprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeStockprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadObjectCvtermRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadObjectCvtermRelationship(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm_relationship\" where \"object_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CvtermRelationship")
	}
	defer results.Close()

	var resultSlice []*CvtermRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CvtermRelationship")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.ObjectCvtermRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.ObjectID {
				local.R.ObjectCvtermRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadSubjectCvtermRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadSubjectCvtermRelationship(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm_relationship\" where \"subject_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CvtermRelationship")
	}
	defer results.Close()

	var resultSlice []*CvtermRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CvtermRelationship")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.SubjectCvtermRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.SubjectID {
				local.R.SubjectCvtermRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeCvtermRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeCvtermRelationship(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm_relationship\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CvtermRelationship")
	}
	defer results.Close()

	var resultSlice []*CvtermRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CvtermRelationship")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeCvtermRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeCvtermRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenotypeCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadPhenotypeCvterm(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_cvterm\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PhenotypeCvterm")
	}
	defer results.Close()

	var resultSlice []*PhenotypeCvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PhenotypeCvterm")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.PhenotypeCvterm = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.CvtermID {
				local.R.PhenotypeCvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenotypeComparisonCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadPhenotypeComparisonCvterm(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_comparison_cvterm\" where \"cvterm_id\" in (%s)",
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

	if len(cvtermAfterSelectHooks) != 0 {
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
			if local.CvtermID == foreign.CvtermID {
				local.R.PhenotypeComparisonCvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadObjectCvtermpath allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadObjectCvtermpath(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvtermpath\" where \"object_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvtermpath")
	}
	defer results.Close()

	var resultSlice []*Cvtermpath
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvtermpath")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.ObjectCvtermpath = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.ObjectID {
				local.R.ObjectCvtermpath = foreign
				break
			}
		}
	}

	return nil
}

// LoadSubjectCvtermpath allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadSubjectCvtermpath(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvtermpath\" where \"subject_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvtermpath")
	}
	defer results.Close()

	var resultSlice []*Cvtermpath
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvtermpath")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.SubjectCvtermpath = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.SubjectID {
				local.R.SubjectCvtermpath = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeCvtermpath allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeCvtermpath(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvtermpath\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvtermpath")
	}
	defer results.Close()

	var resultSlice []*Cvtermpath
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvtermpath")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeCvtermpath = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID.Int {
				local.R.TypeCvtermpath = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeUserRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeUserRelationship(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"user_relationship\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load UserRelationship")
	}
	defer results.Close()

	var resultSlice []*UserRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice UserRelationship")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeUserRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeUserRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadObjectUserRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadObjectUserRelationship(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"user_relationship\" where \"object_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load UserRelationship")
	}
	defer results.Close()

	var resultSlice []*UserRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice UserRelationship")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.ObjectUserRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.ObjectID {
				local.R.ObjectUserRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadSubjectUserRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadSubjectUserRelationship(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"user_relationship\" where \"subject_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load UserRelationship")
	}
	defer results.Close()

	var resultSlice []*UserRelationship
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice UserRelationship")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.SubjectUserRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.SubjectID {
				local.R.SubjectUserRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadCvtermprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadCvtermprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvtermprop\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvtermprop")
	}
	defer results.Close()

	var resultSlice []*Cvtermprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvtermprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Cvtermprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.CvtermID {
				local.R.Cvtermprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadTypeCvtermprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeCvtermprop(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvtermprop\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvtermprop")
	}
	defer results.Close()

	var resultSlice []*Cvtermprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvtermprop")
	}

	if len(cvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.TypeCvtermprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeCvtermprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadAssayPhenotypes allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadAssayPhenotypes(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype\" where \"assay_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load phenotype")
	}
	defer results.Close()

	var resultSlice []*Phenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice phenotype")
	}

	if len(phenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.AssayPhenotypes = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.AssayID.Int {
				local.R.AssayPhenotypes = append(local.R.AssayPhenotypes, foreign)
				break
			}
		}
	}

	return nil
}

// LoadAttrPhenotypes allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadAttrPhenotypes(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype\" where \"attr_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load phenotype")
	}
	defer results.Close()

	var resultSlice []*Phenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice phenotype")
	}

	if len(phenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.AttrPhenotypes = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.AttrID.Int {
				local.R.AttrPhenotypes = append(local.R.AttrPhenotypes, foreign)
				break
			}
		}
	}

	return nil
}

// LoadCvaluePhenotypes allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadCvaluePhenotypes(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype\" where \"cvalue_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load phenotype")
	}
	defer results.Close()

	var resultSlice []*Phenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice phenotype")
	}

	if len(phenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.CvaluePhenotypes = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.CvalueID.Int {
				local.R.CvaluePhenotypes = append(local.R.CvaluePhenotypes, foreign)
				break
			}
		}
	}

	return nil
}

// LoadObservablePhenotypes allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadObservablePhenotypes(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype\" where \"observable_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load phenotype")
	}
	defer results.Close()

	var resultSlice []*Phenotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice phenotype")
	}

	if len(phenotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ObservablePhenotypes = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.ObservableID.Int {
				local.R.ObservablePhenotypes = append(local.R.ObservablePhenotypes, foreign)
				break
			}
		}
	}

	return nil
}

// LoadTypeCvtermsynonyms allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeCvtermsynonyms(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvtermsynonym\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load cvtermsynonym")
	}
	defer results.Close()

	var resultSlice []*Cvtermsynonym
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice cvtermsynonym")
	}

	if len(cvtermsynonymAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.TypeCvtermsynonyms = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID.Int {
				local.R.TypeCvtermsynonyms = append(local.R.TypeCvtermsynonyms, foreign)
				break
			}
		}
	}

	return nil
}

// LoadTypeJbrowseTracks allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeJbrowseTracks(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"jbrowse_track\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load jbrowse_track")
	}
	defer results.Close()

	var resultSlice []*JbrowseTrack
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice jbrowse_track")
	}

	if len(jbrowseTrackAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.TypeJbrowseTracks = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID.Int {
				local.R.TypeJbrowseTracks = append(local.R.TypeJbrowseTracks, foreign)
				break
			}
		}
	}

	return nil
}

// LoadTypePubs allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypePubs(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"pub\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load pub")
	}
	defer results.Close()

	var resultSlice []*Pub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice pub")
	}

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.TypePubs = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypePubs = append(local.R.TypePubs, foreign)
				break
			}
		}
	}

	return nil
}

// LoadStockRelationshipCvterms allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadStockRelationshipCvterms(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_relationship_cvterm\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load stock_relationship_cvterm")
	}
	defer results.Close()

	var resultSlice []*StockRelationshipCvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice stock_relationship_cvterm")
	}

	if len(stockRelationshipCvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.StockRelationshipCvterms = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.CvtermID {
				local.R.StockRelationshipCvterms = append(local.R.StockRelationshipCvterms, foreign)
				break
			}
		}
	}

	return nil
}

// LoadTypeContacts allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeContacts(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"contact\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load contact")
	}
	defer results.Close()

	var resultSlice []*Contact
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice contact")
	}

	if len(contactAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.TypeContacts = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID.Int {
				local.R.TypeContacts = append(local.R.TypeContacts, foreign)
				break
			}
		}
	}

	return nil
}

// LoadTypeGenotypes allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermL) LoadTypeGenotypes(e boil.Executor, singular bool, maybeCvterm interface{}) error {
	var slice []*Cvterm
	var object *Cvterm

	count := 1
	if singular {
		object = maybeCvterm.(*Cvterm)
	} else {
		slice = *maybeCvterm.(*CvtermSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermR{}
		args[0] = object.CvtermID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermR{}
			args[i] = obj.CvtermID
		}
	}

	query := fmt.Sprintf(
		"select * from \"genotype\" where \"type_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load genotype")
	}
	defer results.Close()

	var resultSlice []*Genotype
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice genotype")
	}

	if len(genotypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.TypeGenotypes = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CvtermID == foreign.TypeID {
				local.R.TypeGenotypes = append(local.R.TypeGenotypes, foreign)
				break
			}
		}
	}

	return nil
}

// SetCV of the cvterm to the related item.
// Sets o.R.CV to related.
// Adds o to related.R.Cvterm.
func (o *Cvterm) SetCV(exec boil.Executor, insert bool, related *CV) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"cv_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.CVID, o.CvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CVID = related.CVID

	if o.R == nil {
		o.R = &cvtermR{
			CV: related,
		}
	} else {
		o.R.CV = related
	}

	if related.R == nil {
		related.R = &cvR{
			Cvterm: o,
		}
	} else {
		related.R.Cvterm = o
	}

	return nil
}

// SetDbxref of the cvterm to the related item.
// Sets o.R.Dbxref to related.
// Adds o to related.R.Cvterm.
func (o *Cvterm) SetDbxref(exec boil.Executor, insert bool, related *Dbxref) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvterm\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"dbxref_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermPrimaryKeyColumns),
	)
	values := []interface{}{related.DbxrefID, o.CvtermID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DbxrefID = related.DbxrefID

	if o.R == nil {
		o.R = &cvtermR{
			Dbxref: related,
		}
	} else {
		o.R.Dbxref = related
	}

	if related.R == nil {
		related.R = &dbxrefR{
			Cvterm: o,
		}
	} else {
		related.R.Cvterm = o
	}

	return nil
}

// SetTypeSynonym of the cvterm to the related item.
// Sets o.R.TypeSynonym to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeSynonym(exec boil.Executor, insert bool, related *Synonym) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"synonym\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, synonymPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.SynonymID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeSynonym: related,
		}
	} else {
		o.R.TypeSynonym = related
	}

	if related.R == nil {
		related.R = &synonymR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypePhenotypeprop of the cvterm to the related item.
// Sets o.R.TypePhenotypeprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypePhenotypeprop(exec boil.Executor, insert bool, related *Phenotypeprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenotypeprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenotypepropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.PhenotypepropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypePhenotypeprop: related,
		}
	} else {
		o.R.TypePhenotypeprop = related
	}

	if related.R == nil {
		related.R = &phenotypepropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeStock of the cvterm to the related item.
// Sets o.R.TypeStock to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeStock(exec boil.Executor, insert bool, related *Stock) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.StockID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeStock: related,
		}
	} else {
		o.R.TypeStock = related
	}

	if related.R == nil {
		related.R = &stockR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypePubRelationship of the cvterm to the related item.
// Sets o.R.TypePubRelationship to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypePubRelationship(exec boil.Executor, insert bool, related *PubRelationship) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"pub_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, pubRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.PubRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypePubRelationship: related,
		}
	} else {
		o.R.TypePubRelationship = related
	}

	if related.R == nil {
		related.R = &pubRelationshipR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypePubprop of the cvterm to the related item.
// Sets o.R.TypePubprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypePubprop(exec boil.Executor, insert bool, related *Pubprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"pubprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, pubpropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.PubpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypePubprop: related,
		}
	} else {
		o.R.TypePubprop = related
	}

	if related.R == nil {
		related.R = &pubpropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetCvtermsynonym of the cvterm to the related item.
// Sets o.R.Cvtermsynonym to related.
// Adds o to related.R.Cvterm.
func (o *Cvterm) SetCvtermsynonym(exec boil.Executor, insert bool, related *Cvtermsynonym) error {
	var err error

	if insert {
		related.CvtermID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvtermsynonym\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvtermsynonymPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.CvtermsynonymID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.CvtermID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			Cvtermsynonym: related,
		}
	} else {
		o.R.Cvtermsynonym = related
	}

	if related.R == nil {
		related.R = &cvtermsynonymR{
			Cvterm: o,
		}
	} else {
		related.R.Cvterm = o
	}
	return nil
}

// SetTypeDbxrefprop of the cvterm to the related item.
// Sets o.R.TypeDbxrefprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeDbxrefprop(exec boil.Executor, insert bool, related *Dbxrefprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"dbxrefprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, dbxrefpropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.DbxrefpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeDbxrefprop: related,
		}
	} else {
		o.R.TypeDbxrefprop = related
	}

	if related.R == nil {
		related.R = &dbxrefpropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeFeature of the cvterm to the related item.
// Sets o.R.TypeFeature to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeFeature(exec boil.Executor, insert bool, related *Feature) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurePrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.FeatureID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeFeature: related,
		}
	} else {
		o.R.TypeFeature = related
	}

	if related.R == nil {
		related.R = &featureR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeStockDbxrefprop of the cvterm to the related item.
// Sets o.R.TypeStockDbxrefprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeStockDbxrefprop(exec boil.Executor, insert bool, related *StockDbxrefprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_dbxrefprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockDbxrefpropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.StockDbxrefpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeStockDbxrefprop: related,
		}
	} else {
		o.R.TypeStockDbxrefprop = related
	}

	if related.R == nil {
		related.R = &stockDbxrefpropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetEnvironmentCvterm of the cvterm to the related item.
// Sets o.R.EnvironmentCvterm to related.
// Adds o to related.R.Cvterm.
func (o *Cvterm) SetEnvironmentCvterm(exec boil.Executor, insert bool, related *EnvironmentCvterm) error {
	var err error

	if insert {
		related.CvtermID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"environment_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
			strmangle.WhereClause("\"", "\"", 2, environmentCvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.EnvironmentCvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.CvtermID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			EnvironmentCvterm: related,
		}
	} else {
		o.R.EnvironmentCvterm = related
	}

	if related.R == nil {
		related.R = &environmentCvtermR{
			Cvterm: o,
		}
	} else {
		related.R.Cvterm = o
	}
	return nil
}

// SetStockCvterm of the cvterm to the related item.
// Sets o.R.StockCvterm to related.
// Adds o to related.R.Cvterm.
func (o *Cvterm) SetStockCvterm(exec boil.Executor, insert bool, related *StockCvterm) error {
	var err error

	if insert {
		related.CvtermID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockCvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.StockCvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.CvtermID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			StockCvterm: related,
		}
	} else {
		o.R.StockCvterm = related
	}

	if related.R == nil {
		related.R = &stockCvtermR{
			Cvterm: o,
		}
	} else {
		related.R.Cvterm = o
	}
	return nil
}

// SetTypeStockCvtermprop of the cvterm to the related item.
// Sets o.R.TypeStockCvtermprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeStockCvtermprop(exec boil.Executor, insert bool, related *StockCvtermprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_cvtermprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockCvtermpropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.StockCvtermpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeStockCvtermprop: related,
		}
	} else {
		o.R.TypeStockCvtermprop = related
	}

	if related.R == nil {
		related.R = &stockCvtermpropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeStockcollection of the cvterm to the related item.
// Sets o.R.TypeStockcollection to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeStockcollection(exec boil.Executor, insert bool, related *Stockcollection) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stockcollection\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockcollectionPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.StockcollectionID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeStockcollection: related,
		}
	} else {
		o.R.TypeStockcollection = related
	}

	if related.R == nil {
		related.R = &stockcollectionR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeAnalysisfeatureprop of the cvterm to the related item.
// Sets o.R.TypeAnalysisfeatureprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeAnalysisfeatureprop(exec boil.Executor, insert bool, related *Analysisfeatureprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"analysisfeatureprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, analysisfeaturepropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.AnalysisfeaturepropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeAnalysisfeatureprop: related,
		}
	} else {
		o.R.TypeAnalysisfeatureprop = related
	}

	if related.R == nil {
		related.R = &analysisfeaturepropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetCvtermDbxref of the cvterm to the related item.
// Sets o.R.CvtermDbxref to related.
// Adds o to related.R.Cvterm.
func (o *Cvterm) SetCvtermDbxref(exec boil.Executor, insert bool, related *CvtermDbxref) error {
	var err error

	if insert {
		related.CvtermID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvterm_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvtermDbxrefPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.CvtermDbxrefID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.CvtermID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			CvtermDbxref: related,
		}
	} else {
		o.R.CvtermDbxref = related
	}

	if related.R == nil {
		related.R = &cvtermDbxrefR{
			Cvterm: o,
		}
	} else {
		related.R.Cvterm = o
	}
	return nil
}

// SetTypeFeatureprop of the cvterm to the related item.
// Sets o.R.TypeFeatureprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeFeatureprop(exec boil.Executor, insert bool, related *Featureprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"featureprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurepropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.FeaturepropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeFeatureprop: related,
		}
	} else {
		o.R.TypeFeatureprop = related
	}

	if related.R == nil {
		related.R = &featurepropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeAnalysisprop of the cvterm to the related item.
// Sets o.R.TypeAnalysisprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeAnalysisprop(exec boil.Executor, insert bool, related *Analysisprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"analysisprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, analysispropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.AnalysispropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeAnalysisprop: related,
		}
	} else {
		o.R.TypeAnalysisprop = related
	}

	if related.R == nil {
		related.R = &analysispropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetFeatureCvterm of the cvterm to the related item.
// Sets o.R.FeatureCvterm to related.
// Adds o to related.R.Cvterm.
func (o *Cvterm) SetFeatureCvterm(exec boil.Executor, insert bool, related *FeatureCvterm) error {
	var err error

	if insert {
		related.CvtermID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureCvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.FeatureCvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.CvtermID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			FeatureCvterm: related,
		}
	} else {
		o.R.FeatureCvterm = related
	}

	if related.R == nil {
		related.R = &featureCvtermR{
			Cvterm: o,
		}
	} else {
		related.R.Cvterm = o
	}
	return nil
}

// SetTypeChadoprop of the cvterm to the related item.
// Sets o.R.TypeChadoprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeChadoprop(exec boil.Executor, insert bool, related *Chadoprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"chadoprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, chadopropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.ChadopropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeChadoprop: related,
		}
	} else {
		o.R.TypeChadoprop = related
	}

	if related.R == nil {
		related.R = &chadopropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeFeatureCvtermprop of the cvterm to the related item.
// Sets o.R.TypeFeatureCvtermprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeFeatureCvtermprop(exec boil.Executor, insert bool, related *FeatureCvtermprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_cvtermprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureCvtermpropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.FeatureCvtermpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeFeatureCvtermprop: related,
		}
	} else {
		o.R.TypeFeatureCvtermprop = related
	}

	if related.R == nil {
		related.R = &featureCvtermpropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetFeatureGenotype of the cvterm to the related item.
// Sets o.R.FeatureGenotype to related.
// Adds o to related.R.Cvterm.
func (o *Cvterm) SetFeatureGenotype(exec boil.Executor, insert bool, related *FeatureGenotype) error {
	var err error

	if insert {
		related.CvtermID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_genotype\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureGenotypePrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.FeatureGenotypeID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.CvtermID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			FeatureGenotype: related,
		}
	} else {
		o.R.FeatureGenotype = related
	}

	if related.R == nil {
		related.R = &featureGenotypeR{
			Cvterm: o,
		}
	} else {
		related.R.Cvterm = o
	}
	return nil
}

// SetTypeContactRelationship of the cvterm to the related item.
// Sets o.R.TypeContactRelationship to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeContactRelationship(exec boil.Executor, insert bool, related *ContactRelationship) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"contact_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, contactRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.ContactRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeContactRelationship: related,
		}
	} else {
		o.R.TypeContactRelationship = related
	}

	if related.R == nil {
		related.R = &contactRelationshipR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeFeaturePubprop of the cvterm to the related item.
// Sets o.R.TypeFeaturePubprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeFeaturePubprop(exec boil.Executor, insert bool, related *FeaturePubprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_pubprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurePubpropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.FeaturePubpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeFeaturePubprop: related,
		}
	} else {
		o.R.TypeFeaturePubprop = related
	}

	if related.R == nil {
		related.R = &featurePubpropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeCvprop of the cvterm to the related item.
// Sets o.R.TypeCvprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeCvprop(exec boil.Executor, insert bool, related *Cvprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvpropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.CvpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeCvprop: related,
		}
	} else {
		o.R.TypeCvprop = related
	}

	if related.R == nil {
		related.R = &cvpropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeStockRelationship of the cvterm to the related item.
// Sets o.R.TypeStockRelationship to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeStockRelationship(exec boil.Executor, insert bool, related *StockRelationship) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.StockRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeStockRelationship: related,
		}
	} else {
		o.R.TypeStockRelationship = related
	}

	if related.R == nil {
		related.R = &stockRelationshipR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeFeatureRelationshipprop of the cvterm to the related item.
// Sets o.R.TypeFeatureRelationshipprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeFeatureRelationshipprop(exec boil.Executor, insert bool, related *FeatureRelationshipprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_relationshipprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureRelationshippropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.FeatureRelationshippropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeFeatureRelationshipprop: related,
		}
	} else {
		o.R.TypeFeatureRelationshipprop = related
	}

	if related.R == nil {
		related.R = &featureRelationshippropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeFeatureRelationship of the cvterm to the related item.
// Sets o.R.TypeFeatureRelationship to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeFeatureRelationship(exec boil.Executor, insert bool, related *FeatureRelationship) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.FeatureRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeFeatureRelationship: related,
		}
	} else {
		o.R.TypeFeatureRelationship = related
	}

	if related.R == nil {
		related.R = &featureRelationshipR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeGenotypeprop of the cvterm to the related item.
// Sets o.R.TypeGenotypeprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeGenotypeprop(exec boil.Executor, insert bool, related *Genotypeprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"genotypeprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, genotypepropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.GenotypepropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeGenotypeprop: related,
		}
	} else {
		o.R.TypeGenotypeprop = related
	}

	if related.R == nil {
		related.R = &genotypepropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeOrganismprop of the cvterm to the related item.
// Sets o.R.TypeOrganismprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeOrganismprop(exec boil.Executor, insert bool, related *Organismprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"organismprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, organismpropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.OrganismpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeOrganismprop: related,
		}
	} else {
		o.R.TypeOrganismprop = related
	}

	if related.R == nil {
		related.R = &organismpropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypePhendesc of the cvterm to the related item.
// Sets o.R.TypePhendesc to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypePhendesc(exec boil.Executor, insert bool, related *Phendesc) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phendesc\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, phendescPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.PhendescID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypePhendesc: related,
		}
	} else {
		o.R.TypePhendesc = related
	}

	if related.R == nil {
		related.R = &phendescR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypePhenstatement of the cvterm to the related item.
// Sets o.R.TypePhenstatement to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypePhenstatement(exec boil.Executor, insert bool, related *Phenstatement) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenstatement\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenstatementPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.PhenstatementID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypePhenstatement: related,
		}
	} else {
		o.R.TypePhenstatement = related
	}

	if related.R == nil {
		related.R = &phenstatementR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeStockcollectionprop of the cvterm to the related item.
// Sets o.R.TypeStockcollectionprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeStockcollectionprop(exec boil.Executor, insert bool, related *Stockcollectionprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stockcollectionprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockcollectionpropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.StockcollectionpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeStockcollectionprop: related,
		}
	} else {
		o.R.TypeStockcollectionprop = related
	}

	if related.R == nil {
		related.R = &stockcollectionpropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetTypeStockprop of the cvterm to the related item.
// Sets o.R.TypeStockprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeStockprop(exec boil.Executor, insert bool, related *Stockprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stockprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockpropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.StockpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeStockprop: related,
		}
	} else {
		o.R.TypeStockprop = related
	}

	if related.R == nil {
		related.R = &stockpropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetObjectCvtermRelationship of the cvterm to the related item.
// Sets o.R.ObjectCvtermRelationship to related.
// Adds o to related.R.Object.
func (o *Cvterm) SetObjectCvtermRelationship(exec boil.Executor, insert bool, related *CvtermRelationship) error {
	var err error

	if insert {
		related.ObjectID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvterm_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"object_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvtermRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.CvtermRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.ObjectID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			ObjectCvtermRelationship: related,
		}
	} else {
		o.R.ObjectCvtermRelationship = related
	}

	if related.R == nil {
		related.R = &cvtermRelationshipR{
			Object: o,
		}
	} else {
		related.R.Object = o
	}
	return nil
}

// SetSubjectCvtermRelationship of the cvterm to the related item.
// Sets o.R.SubjectCvtermRelationship to related.
// Adds o to related.R.Subject.
func (o *Cvterm) SetSubjectCvtermRelationship(exec boil.Executor, insert bool, related *CvtermRelationship) error {
	var err error

	if insert {
		related.SubjectID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvterm_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"subject_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvtermRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.CvtermRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.SubjectID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			SubjectCvtermRelationship: related,
		}
	} else {
		o.R.SubjectCvtermRelationship = related
	}

	if related.R == nil {
		related.R = &cvtermRelationshipR{
			Subject: o,
		}
	} else {
		related.R.Subject = o
	}
	return nil
}

// SetTypeCvtermRelationship of the cvterm to the related item.
// Sets o.R.TypeCvtermRelationship to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeCvtermRelationship(exec boil.Executor, insert bool, related *CvtermRelationship) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvterm_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvtermRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.CvtermRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeCvtermRelationship: related,
		}
	} else {
		o.R.TypeCvtermRelationship = related
	}

	if related.R == nil {
		related.R = &cvtermRelationshipR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetPhenotypeCvterm of the cvterm to the related item.
// Sets o.R.PhenotypeCvterm to related.
// Adds o to related.R.Cvterm.
func (o *Cvterm) SetPhenotypeCvterm(exec boil.Executor, insert bool, related *PhenotypeCvterm) error {
	var err error

	if insert {
		related.CvtermID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenotype_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenotypeCvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.PhenotypeCvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.CvtermID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			PhenotypeCvterm: related,
		}
	} else {
		o.R.PhenotypeCvterm = related
	}

	if related.R == nil {
		related.R = &phenotypeCvtermR{
			Cvterm: o,
		}
	} else {
		related.R.Cvterm = o
	}
	return nil
}

// SetPhenotypeComparisonCvterm of the cvterm to the related item.
// Sets o.R.PhenotypeComparisonCvterm to related.
// Adds o to related.R.Cvterm.
func (o *Cvterm) SetPhenotypeComparisonCvterm(exec boil.Executor, insert bool, related *PhenotypeComparisonCvterm) error {
	var err error

	if insert {
		related.CvtermID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenotype_comparison_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonCvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.PhenotypeComparisonCvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.CvtermID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			PhenotypeComparisonCvterm: related,
		}
	} else {
		o.R.PhenotypeComparisonCvterm = related
	}

	if related.R == nil {
		related.R = &phenotypeComparisonCvtermR{
			Cvterm: o,
		}
	} else {
		related.R.Cvterm = o
	}
	return nil
}

// SetObjectCvtermpath of the cvterm to the related item.
// Sets o.R.ObjectCvtermpath to related.
// Adds o to related.R.Object.
func (o *Cvterm) SetObjectCvtermpath(exec boil.Executor, insert bool, related *Cvtermpath) error {
	var err error

	if insert {
		related.ObjectID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvtermpath\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"object_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvtermpathPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.CvtermpathID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.ObjectID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			ObjectCvtermpath: related,
		}
	} else {
		o.R.ObjectCvtermpath = related
	}

	if related.R == nil {
		related.R = &cvtermpathR{
			Object: o,
		}
	} else {
		related.R.Object = o
	}
	return nil
}

// SetSubjectCvtermpath of the cvterm to the related item.
// Sets o.R.SubjectCvtermpath to related.
// Adds o to related.R.Subject.
func (o *Cvterm) SetSubjectCvtermpath(exec boil.Executor, insert bool, related *Cvtermpath) error {
	var err error

	if insert {
		related.SubjectID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvtermpath\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"subject_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvtermpathPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.CvtermpathID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.SubjectID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			SubjectCvtermpath: related,
		}
	} else {
		o.R.SubjectCvtermpath = related
	}

	if related.R == nil {
		related.R = &cvtermpathR{
			Subject: o,
		}
	} else {
		related.R.Subject = o
	}
	return nil
}

// SetTypeCvtermpath of the cvterm to the related item.
// Sets o.R.TypeCvtermpath to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeCvtermpath(exec boil.Executor, insert bool, related *Cvtermpath) error {
	var err error

	if insert {
		related.TypeID.Int = o.CvtermID
		related.TypeID.Valid = true

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvtermpath\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvtermpathPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.CvtermpathID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID.Int = o.CvtermID
		related.TypeID.Valid = true
	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeCvtermpath: related,
		}
	} else {
		o.R.TypeCvtermpath = related
	}

	if related.R == nil {
		related.R = &cvtermpathR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// RemoveTypeCvtermpath relationship.
// Sets o.R.TypeCvtermpath to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Cvterm) RemoveTypeCvtermpath(exec boil.Executor, related *Cvtermpath) error {
	var err error

	related.TypeID.Valid = false
	if err = related.Update(exec, "type_id"); err != nil {
		related.TypeID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.TypeCvtermpath = nil
	if related == nil || related.R == nil {
		return nil
	}

	related.R.Type = nil
	return nil
}

// SetTypeUserRelationship of the cvterm to the related item.
// Sets o.R.TypeUserRelationship to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeUserRelationship(exec boil.Executor, insert bool, related *UserRelationship) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"user_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, userRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.UserRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeUserRelationship: related,
		}
	} else {
		o.R.TypeUserRelationship = related
	}

	if related.R == nil {
		related.R = &userRelationshipR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// SetObjectUserRelationship of the cvterm to the related item.
// Sets o.R.ObjectUserRelationship to related.
// Adds o to related.R.Object.
func (o *Cvterm) SetObjectUserRelationship(exec boil.Executor, insert bool, related *UserRelationship) error {
	var err error

	if insert {
		related.ObjectID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"user_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"object_id"}),
			strmangle.WhereClause("\"", "\"", 2, userRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.UserRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.ObjectID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			ObjectUserRelationship: related,
		}
	} else {
		o.R.ObjectUserRelationship = related
	}

	if related.R == nil {
		related.R = &userRelationshipR{
			Object: o,
		}
	} else {
		related.R.Object = o
	}
	return nil
}

// SetSubjectUserRelationship of the cvterm to the related item.
// Sets o.R.SubjectUserRelationship to related.
// Adds o to related.R.Subject.
func (o *Cvterm) SetSubjectUserRelationship(exec boil.Executor, insert bool, related *UserRelationship) error {
	var err error

	if insert {
		related.SubjectID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"user_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"subject_id"}),
			strmangle.WhereClause("\"", "\"", 2, userRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.UserRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.SubjectID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			SubjectUserRelationship: related,
		}
	} else {
		o.R.SubjectUserRelationship = related
	}

	if related.R == nil {
		related.R = &userRelationshipR{
			Subject: o,
		}
	} else {
		related.R.Subject = o
	}
	return nil
}

// SetCvtermprop of the cvterm to the related item.
// Sets o.R.Cvtermprop to related.
// Adds o to related.R.Cvterm.
func (o *Cvterm) SetCvtermprop(exec boil.Executor, insert bool, related *Cvtermprop) error {
	var err error

	if insert {
		related.CvtermID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvtermprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"cvterm_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvtermpropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.CvtermpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.CvtermID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			Cvtermprop: related,
		}
	} else {
		o.R.Cvtermprop = related
	}

	if related.R == nil {
		related.R = &cvtermpropR{
			Cvterm: o,
		}
	} else {
		related.R.Cvterm = o
	}
	return nil
}

// SetTypeCvtermprop of the cvterm to the related item.
// Sets o.R.TypeCvtermprop to related.
// Adds o to related.R.Type.
func (o *Cvterm) SetTypeCvtermprop(exec boil.Executor, insert bool, related *Cvtermprop) error {
	var err error

	if insert {
		related.TypeID = o.CvtermID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"cvtermprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
			strmangle.WhereClause("\"", "\"", 2, cvtermpropPrimaryKeyColumns),
		)
		values := []interface{}{o.CvtermID, related.CvtermpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.TypeID = o.CvtermID

	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeCvtermprop: related,
		}
	} else {
		o.R.TypeCvtermprop = related
	}

	if related.R == nil {
		related.R = &cvtermpropR{
			Type: o,
		}
	} else {
		related.R.Type = o
	}
	return nil
}

// AddAssayPhenotypes adds the given related objects to the existing relationships
// of the cvterm, optionally inserting them as new records.
// Appends related to o.R.AssayPhenotypes.
// Sets related.R.Assay appropriately.
func (o *Cvterm) AddAssayPhenotypes(exec boil.Executor, insert bool, related ...*Phenotype) error {
	var err error
	for _, rel := range related {
		rel.AssayID.Int = o.CvtermID
		rel.AssayID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "assay_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &cvtermR{
			AssayPhenotypes: related,
		}
	} else {
		o.R.AssayPhenotypes = append(o.R.AssayPhenotypes, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &phenotypeR{
				Assay: o,
			}
		} else {
			rel.R.Assay = o
		}
	}
	return nil
}

// SetAssayPhenotypes removes all previously related items of the
// cvterm replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Assay's AssayPhenotypes accordingly.
// Replaces o.R.AssayPhenotypes with related.
// Sets related.R.Assay's AssayPhenotypes accordingly.
func (o *Cvterm) SetAssayPhenotypes(exec boil.Executor, insert bool, related ...*Phenotype) error {
	query := "update \"phenotype\" set \"assay_id\" = null where \"assay_id\" = $1"
	values := []interface{}{o.CvtermID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.AssayPhenotypes {
			rel.AssayID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Assay = nil
		}

		o.R.AssayPhenotypes = nil
	}
	return o.AddAssayPhenotypes(exec, insert, related...)
}

// RemoveAssayPhenotypes relationships from objects passed in.
// Removes related items from R.AssayPhenotypes (uses pointer comparison, removal does not keep order)
// Sets related.R.Assay.
func (o *Cvterm) RemoveAssayPhenotypes(exec boil.Executor, related ...*Phenotype) error {
	var err error
	for _, rel := range related {
		rel.AssayID.Valid = false
		if rel.R != nil {
			rel.R.Assay = nil
		}
		if err = rel.Update(exec, "assay_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.AssayPhenotypes {
			if rel != ri {
				continue
			}

			ln := len(o.R.AssayPhenotypes)
			if ln > 1 && i < ln-1 {
				o.R.AssayPhenotypes[i] = o.R.AssayPhenotypes[ln-1]
			}
			o.R.AssayPhenotypes = o.R.AssayPhenotypes[:ln-1]
			break
		}
	}

	return nil
}

// AddAttrPhenotypes adds the given related objects to the existing relationships
// of the cvterm, optionally inserting them as new records.
// Appends related to o.R.AttrPhenotypes.
// Sets related.R.Attr appropriately.
func (o *Cvterm) AddAttrPhenotypes(exec boil.Executor, insert bool, related ...*Phenotype) error {
	var err error
	for _, rel := range related {
		rel.AttrID.Int = o.CvtermID
		rel.AttrID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "attr_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &cvtermR{
			AttrPhenotypes: related,
		}
	} else {
		o.R.AttrPhenotypes = append(o.R.AttrPhenotypes, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &phenotypeR{
				Attr: o,
			}
		} else {
			rel.R.Attr = o
		}
	}
	return nil
}

// SetAttrPhenotypes removes all previously related items of the
// cvterm replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Attr's AttrPhenotypes accordingly.
// Replaces o.R.AttrPhenotypes with related.
// Sets related.R.Attr's AttrPhenotypes accordingly.
func (o *Cvterm) SetAttrPhenotypes(exec boil.Executor, insert bool, related ...*Phenotype) error {
	query := "update \"phenotype\" set \"attr_id\" = null where \"attr_id\" = $1"
	values := []interface{}{o.CvtermID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.AttrPhenotypes {
			rel.AttrID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Attr = nil
		}

		o.R.AttrPhenotypes = nil
	}
	return o.AddAttrPhenotypes(exec, insert, related...)
}

// RemoveAttrPhenotypes relationships from objects passed in.
// Removes related items from R.AttrPhenotypes (uses pointer comparison, removal does not keep order)
// Sets related.R.Attr.
func (o *Cvterm) RemoveAttrPhenotypes(exec boil.Executor, related ...*Phenotype) error {
	var err error
	for _, rel := range related {
		rel.AttrID.Valid = false
		if rel.R != nil {
			rel.R.Attr = nil
		}
		if err = rel.Update(exec, "attr_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.AttrPhenotypes {
			if rel != ri {
				continue
			}

			ln := len(o.R.AttrPhenotypes)
			if ln > 1 && i < ln-1 {
				o.R.AttrPhenotypes[i] = o.R.AttrPhenotypes[ln-1]
			}
			o.R.AttrPhenotypes = o.R.AttrPhenotypes[:ln-1]
			break
		}
	}

	return nil
}

// AddCvaluePhenotypes adds the given related objects to the existing relationships
// of the cvterm, optionally inserting them as new records.
// Appends related to o.R.CvaluePhenotypes.
// Sets related.R.Cvalue appropriately.
func (o *Cvterm) AddCvaluePhenotypes(exec boil.Executor, insert bool, related ...*Phenotype) error {
	var err error
	for _, rel := range related {
		rel.CvalueID.Int = o.CvtermID
		rel.CvalueID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "cvalue_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &cvtermR{
			CvaluePhenotypes: related,
		}
	} else {
		o.R.CvaluePhenotypes = append(o.R.CvaluePhenotypes, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &phenotypeR{
				Cvalue: o,
			}
		} else {
			rel.R.Cvalue = o
		}
	}
	return nil
}

// SetCvaluePhenotypes removes all previously related items of the
// cvterm replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Cvalue's CvaluePhenotypes accordingly.
// Replaces o.R.CvaluePhenotypes with related.
// Sets related.R.Cvalue's CvaluePhenotypes accordingly.
func (o *Cvterm) SetCvaluePhenotypes(exec boil.Executor, insert bool, related ...*Phenotype) error {
	query := "update \"phenotype\" set \"cvalue_id\" = null where \"cvalue_id\" = $1"
	values := []interface{}{o.CvtermID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.CvaluePhenotypes {
			rel.CvalueID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Cvalue = nil
		}

		o.R.CvaluePhenotypes = nil
	}
	return o.AddCvaluePhenotypes(exec, insert, related...)
}

// RemoveCvaluePhenotypes relationships from objects passed in.
// Removes related items from R.CvaluePhenotypes (uses pointer comparison, removal does not keep order)
// Sets related.R.Cvalue.
func (o *Cvterm) RemoveCvaluePhenotypes(exec boil.Executor, related ...*Phenotype) error {
	var err error
	for _, rel := range related {
		rel.CvalueID.Valid = false
		if rel.R != nil {
			rel.R.Cvalue = nil
		}
		if err = rel.Update(exec, "cvalue_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.CvaluePhenotypes {
			if rel != ri {
				continue
			}

			ln := len(o.R.CvaluePhenotypes)
			if ln > 1 && i < ln-1 {
				o.R.CvaluePhenotypes[i] = o.R.CvaluePhenotypes[ln-1]
			}
			o.R.CvaluePhenotypes = o.R.CvaluePhenotypes[:ln-1]
			break
		}
	}

	return nil
}

// AddObservablePhenotypes adds the given related objects to the existing relationships
// of the cvterm, optionally inserting them as new records.
// Appends related to o.R.ObservablePhenotypes.
// Sets related.R.Observable appropriately.
func (o *Cvterm) AddObservablePhenotypes(exec boil.Executor, insert bool, related ...*Phenotype) error {
	var err error
	for _, rel := range related {
		rel.ObservableID.Int = o.CvtermID
		rel.ObservableID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "observable_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &cvtermR{
			ObservablePhenotypes: related,
		}
	} else {
		o.R.ObservablePhenotypes = append(o.R.ObservablePhenotypes, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &phenotypeR{
				Observable: o,
			}
		} else {
			rel.R.Observable = o
		}
	}
	return nil
}

// SetObservablePhenotypes removes all previously related items of the
// cvterm replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Observable's ObservablePhenotypes accordingly.
// Replaces o.R.ObservablePhenotypes with related.
// Sets related.R.Observable's ObservablePhenotypes accordingly.
func (o *Cvterm) SetObservablePhenotypes(exec boil.Executor, insert bool, related ...*Phenotype) error {
	query := "update \"phenotype\" set \"observable_id\" = null where \"observable_id\" = $1"
	values := []interface{}{o.CvtermID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.ObservablePhenotypes {
			rel.ObservableID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Observable = nil
		}

		o.R.ObservablePhenotypes = nil
	}
	return o.AddObservablePhenotypes(exec, insert, related...)
}

// RemoveObservablePhenotypes relationships from objects passed in.
// Removes related items from R.ObservablePhenotypes (uses pointer comparison, removal does not keep order)
// Sets related.R.Observable.
func (o *Cvterm) RemoveObservablePhenotypes(exec boil.Executor, related ...*Phenotype) error {
	var err error
	for _, rel := range related {
		rel.ObservableID.Valid = false
		if rel.R != nil {
			rel.R.Observable = nil
		}
		if err = rel.Update(exec, "observable_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ObservablePhenotypes {
			if rel != ri {
				continue
			}

			ln := len(o.R.ObservablePhenotypes)
			if ln > 1 && i < ln-1 {
				o.R.ObservablePhenotypes[i] = o.R.ObservablePhenotypes[ln-1]
			}
			o.R.ObservablePhenotypes = o.R.ObservablePhenotypes[:ln-1]
			break
		}
	}

	return nil
}

// AddTypeCvtermsynonyms adds the given related objects to the existing relationships
// of the cvterm, optionally inserting them as new records.
// Appends related to o.R.TypeCvtermsynonyms.
// Sets related.R.Type appropriately.
func (o *Cvterm) AddTypeCvtermsynonyms(exec boil.Executor, insert bool, related ...*Cvtermsynonym) error {
	var err error
	for _, rel := range related {
		rel.TypeID.Int = o.CvtermID
		rel.TypeID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "type_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeCvtermsynonyms: related,
		}
	} else {
		o.R.TypeCvtermsynonyms = append(o.R.TypeCvtermsynonyms, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &cvtermsynonymR{
				Type: o,
			}
		} else {
			rel.R.Type = o
		}
	}
	return nil
}

// SetTypeCvtermsynonyms removes all previously related items of the
// cvterm replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Type's TypeCvtermsynonyms accordingly.
// Replaces o.R.TypeCvtermsynonyms with related.
// Sets related.R.Type's TypeCvtermsynonyms accordingly.
func (o *Cvterm) SetTypeCvtermsynonyms(exec boil.Executor, insert bool, related ...*Cvtermsynonym) error {
	query := "update \"cvtermsynonym\" set \"type_id\" = null where \"type_id\" = $1"
	values := []interface{}{o.CvtermID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.TypeCvtermsynonyms {
			rel.TypeID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Type = nil
		}

		o.R.TypeCvtermsynonyms = nil
	}
	return o.AddTypeCvtermsynonyms(exec, insert, related...)
}

// RemoveTypeCvtermsynonyms relationships from objects passed in.
// Removes related items from R.TypeCvtermsynonyms (uses pointer comparison, removal does not keep order)
// Sets related.R.Type.
func (o *Cvterm) RemoveTypeCvtermsynonyms(exec boil.Executor, related ...*Cvtermsynonym) error {
	var err error
	for _, rel := range related {
		rel.TypeID.Valid = false
		if rel.R != nil {
			rel.R.Type = nil
		}
		if err = rel.Update(exec, "type_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.TypeCvtermsynonyms {
			if rel != ri {
				continue
			}

			ln := len(o.R.TypeCvtermsynonyms)
			if ln > 1 && i < ln-1 {
				o.R.TypeCvtermsynonyms[i] = o.R.TypeCvtermsynonyms[ln-1]
			}
			o.R.TypeCvtermsynonyms = o.R.TypeCvtermsynonyms[:ln-1]
			break
		}
	}

	return nil
}

// AddTypeJbrowseTracks adds the given related objects to the existing relationships
// of the cvterm, optionally inserting them as new records.
// Appends related to o.R.TypeJbrowseTracks.
// Sets related.R.Type appropriately.
func (o *Cvterm) AddTypeJbrowseTracks(exec boil.Executor, insert bool, related ...*JbrowseTrack) error {
	var err error
	for _, rel := range related {
		rel.TypeID.Int = o.CvtermID
		rel.TypeID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "type_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeJbrowseTracks: related,
		}
	} else {
		o.R.TypeJbrowseTracks = append(o.R.TypeJbrowseTracks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &jbrowseTrackR{
				Type: o,
			}
		} else {
			rel.R.Type = o
		}
	}
	return nil
}

// SetTypeJbrowseTracks removes all previously related items of the
// cvterm replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Type's TypeJbrowseTracks accordingly.
// Replaces o.R.TypeJbrowseTracks with related.
// Sets related.R.Type's TypeJbrowseTracks accordingly.
func (o *Cvterm) SetTypeJbrowseTracks(exec boil.Executor, insert bool, related ...*JbrowseTrack) error {
	query := "update \"jbrowse_track\" set \"type_id\" = null where \"type_id\" = $1"
	values := []interface{}{o.CvtermID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.TypeJbrowseTracks {
			rel.TypeID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Type = nil
		}

		o.R.TypeJbrowseTracks = nil
	}
	return o.AddTypeJbrowseTracks(exec, insert, related...)
}

// RemoveTypeJbrowseTracks relationships from objects passed in.
// Removes related items from R.TypeJbrowseTracks (uses pointer comparison, removal does not keep order)
// Sets related.R.Type.
func (o *Cvterm) RemoveTypeJbrowseTracks(exec boil.Executor, related ...*JbrowseTrack) error {
	var err error
	for _, rel := range related {
		rel.TypeID.Valid = false
		if rel.R != nil {
			rel.R.Type = nil
		}
		if err = rel.Update(exec, "type_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.TypeJbrowseTracks {
			if rel != ri {
				continue
			}

			ln := len(o.R.TypeJbrowseTracks)
			if ln > 1 && i < ln-1 {
				o.R.TypeJbrowseTracks[i] = o.R.TypeJbrowseTracks[ln-1]
			}
			o.R.TypeJbrowseTracks = o.R.TypeJbrowseTracks[:ln-1]
			break
		}
	}

	return nil
}

// AddTypePubs adds the given related objects to the existing relationships
// of the cvterm, optionally inserting them as new records.
// Appends related to o.R.TypePubs.
// Sets related.R.Type appropriately.
func (o *Cvterm) AddTypePubs(exec boil.Executor, insert bool, related ...*Pub) error {
	var err error
	for _, rel := range related {
		rel.TypeID = o.CvtermID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "type_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &cvtermR{
			TypePubs: related,
		}
	} else {
		o.R.TypePubs = append(o.R.TypePubs, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &pubR{
				Type: o,
			}
		} else {
			rel.R.Type = o
		}
	}
	return nil
}

// AddStockRelationshipCvterms adds the given related objects to the existing relationships
// of the cvterm, optionally inserting them as new records.
// Appends related to o.R.StockRelationshipCvterms.
// Sets related.R.Cvterm appropriately.
func (o *Cvterm) AddStockRelationshipCvterms(exec boil.Executor, insert bool, related ...*StockRelationshipCvterm) error {
	var err error
	for _, rel := range related {
		rel.CvtermID = o.CvtermID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "cvterm_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &cvtermR{
			StockRelationshipCvterms: related,
		}
	} else {
		o.R.StockRelationshipCvterms = append(o.R.StockRelationshipCvterms, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &stockRelationshipCvtermR{
				Cvterm: o,
			}
		} else {
			rel.R.Cvterm = o
		}
	}
	return nil
}

// AddTypeContacts adds the given related objects to the existing relationships
// of the cvterm, optionally inserting them as new records.
// Appends related to o.R.TypeContacts.
// Sets related.R.Type appropriately.
func (o *Cvterm) AddTypeContacts(exec boil.Executor, insert bool, related ...*Contact) error {
	var err error
	for _, rel := range related {
		rel.TypeID.Int = o.CvtermID
		rel.TypeID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "type_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeContacts: related,
		}
	} else {
		o.R.TypeContacts = append(o.R.TypeContacts, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &contactR{
				Type: o,
			}
		} else {
			rel.R.Type = o
		}
	}
	return nil
}

// SetTypeContacts removes all previously related items of the
// cvterm replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Type's TypeContacts accordingly.
// Replaces o.R.TypeContacts with related.
// Sets related.R.Type's TypeContacts accordingly.
func (o *Cvterm) SetTypeContacts(exec boil.Executor, insert bool, related ...*Contact) error {
	query := "update \"contact\" set \"type_id\" = null where \"type_id\" = $1"
	values := []interface{}{o.CvtermID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.TypeContacts {
			rel.TypeID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Type = nil
		}

		o.R.TypeContacts = nil
	}
	return o.AddTypeContacts(exec, insert, related...)
}

// RemoveTypeContacts relationships from objects passed in.
// Removes related items from R.TypeContacts (uses pointer comparison, removal does not keep order)
// Sets related.R.Type.
func (o *Cvterm) RemoveTypeContacts(exec boil.Executor, related ...*Contact) error {
	var err error
	for _, rel := range related {
		rel.TypeID.Valid = false
		if rel.R != nil {
			rel.R.Type = nil
		}
		if err = rel.Update(exec, "type_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.TypeContacts {
			if rel != ri {
				continue
			}

			ln := len(o.R.TypeContacts)
			if ln > 1 && i < ln-1 {
				o.R.TypeContacts[i] = o.R.TypeContacts[ln-1]
			}
			o.R.TypeContacts = o.R.TypeContacts[:ln-1]
			break
		}
	}

	return nil
}

// AddTypeGenotypes adds the given related objects to the existing relationships
// of the cvterm, optionally inserting them as new records.
// Appends related to o.R.TypeGenotypes.
// Sets related.R.Type appropriately.
func (o *Cvterm) AddTypeGenotypes(exec boil.Executor, insert bool, related ...*Genotype) error {
	var err error
	for _, rel := range related {
		rel.TypeID = o.CvtermID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "type_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &cvtermR{
			TypeGenotypes: related,
		}
	} else {
		o.R.TypeGenotypes = append(o.R.TypeGenotypes, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &genotypeR{
				Type: o,
			}
		} else {
			rel.R.Type = o
		}
	}
	return nil
}

// CvtermsG retrieves all records.
func CvtermsG(mods ...qm.QueryMod) cvtermQuery {
	return Cvterms(boil.GetDB(), mods...)
}

// Cvterms retrieves all the records using an executor.
func Cvterms(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	mods = append(mods, qm.From("\"cvterm\""))
	return cvtermQuery{NewQuery(exec, mods...)}
}

// FindCvtermG retrieves a single record by ID.
func FindCvtermG(cvtermID int, selectCols ...string) (*Cvterm, error) {
	return FindCvterm(boil.GetDB(), cvtermID, selectCols...)
}

// FindCvtermGP retrieves a single record by ID, and panics on error.
func FindCvtermGP(cvtermID int, selectCols ...string) *Cvterm {
	retobj, err := FindCvterm(boil.GetDB(), cvtermID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCvterm retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCvterm(exec boil.Executor, cvtermID int, selectCols ...string) (*Cvterm, error) {
	cvtermObj := &Cvterm{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"cvterm\" where \"cvterm_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, cvtermID)

	err := q.Bind(cvtermObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cvterm")
	}

	return cvtermObj, nil
}

// FindCvtermP retrieves a single record by ID with an executor, and panics on error.
func FindCvtermP(exec boil.Executor, cvtermID int, selectCols ...string) *Cvterm {
	retobj, err := FindCvterm(exec, cvtermID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Cvterm) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Cvterm) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Cvterm) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Cvterm) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no cvterm provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvtermColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	cvtermInsertCacheMut.RLock()
	cache, cached := cvtermInsertCache[key]
	cvtermInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			cvtermColumns,
			cvtermColumnsWithDefault,
			cvtermColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(cvtermType, cvtermMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cvtermType, cvtermMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"cvterm\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into cvterm")
	}

	if !cached {
		cvtermInsertCacheMut.Lock()
		cvtermInsertCache[key] = cache
		cvtermInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Cvterm record. See Update for
// whitelist behavior description.
func (o *Cvterm) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Cvterm record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Cvterm) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Cvterm, and panics on error.
// See Update for whitelist behavior description.
func (o *Cvterm) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Cvterm.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Cvterm) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	cvtermUpdateCacheMut.RLock()
	cache, cached := cvtermUpdateCache[key]
	cvtermUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(cvtermColumns, cvtermPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update cvterm, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, cvtermPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cvtermType, cvtermMapping, append(wl, cvtermPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update cvterm row")
	}

	if !cached {
		cvtermUpdateCacheMut.Lock()
		cvtermUpdateCache[key] = cache
		cvtermUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q cvtermQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q cvtermQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for cvterm")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CvtermSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CvtermSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CvtermSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CvtermSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"cvterm\" SET %s WHERE (\"cvterm_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvtermPrimaryKeyColumns), len(colNames)+1, len(cvtermPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in cvterm slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Cvterm) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Cvterm) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Cvterm) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Cvterm) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no cvterm provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvtermColumnsWithDefault, o)

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

	cvtermUpsertCacheMut.RLock()
	cache, cached := cvtermUpsertCache[key]
	cvtermUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			cvtermColumns,
			cvtermColumnsWithDefault,
			cvtermColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			cvtermColumns,
			cvtermPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert cvterm, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(cvtermPrimaryKeyColumns))
			copy(conflict, cvtermPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"cvterm\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(cvtermType, cvtermMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cvtermType, cvtermMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for cvterm")
	}

	if !cached {
		cvtermUpsertCacheMut.Lock()
		cvtermUpsertCache[key] = cache
		cvtermUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Cvterm record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Cvterm) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Cvterm record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Cvterm) DeleteG() error {
	if o == nil {
		return errors.New("models: no Cvterm provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Cvterm record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Cvterm) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Cvterm record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Cvterm) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Cvterm provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cvtermPrimaryKeyMapping)
	sql := "DELETE FROM \"cvterm\" WHERE \"cvterm_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from cvterm")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q cvtermQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q cvtermQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no cvtermQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from cvterm")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CvtermSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CvtermSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Cvterm slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CvtermSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CvtermSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Cvterm slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(cvtermBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"cvterm\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvtermPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvtermPrimaryKeyColumns), 1, len(cvtermPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from cvterm slice")
	}

	if len(cvtermAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Cvterm) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Cvterm) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Cvterm) ReloadG() error {
	if o == nil {
		return errors.New("models: no Cvterm provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Cvterm) Reload(exec boil.Executor) error {
	ret, err := FindCvterm(exec, o.CvtermID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvtermSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvtermSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvtermSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty CvtermSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvtermSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	cvterms := CvtermSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"cvterm\".* FROM \"cvterm\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvtermPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(cvtermPrimaryKeyColumns), 1, len(cvtermPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&cvterms)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CvtermSlice")
	}

	*o = cvterms

	return nil
}

// CvtermExists checks if the Cvterm row exists.
func CvtermExists(exec boil.Executor, cvtermID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"cvterm\" where \"cvterm_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, cvtermID)
	}

	row := exec.QueryRow(sql, cvtermID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cvterm exists")
	}

	return exists, nil
}

// CvtermExistsG checks if the Cvterm row exists.
func CvtermExistsG(cvtermID int) (bool, error) {
	return CvtermExists(boil.GetDB(), cvtermID)
}

// CvtermExistsGP checks if the Cvterm row exists. Panics on error.
func CvtermExistsGP(cvtermID int) bool {
	e, err := CvtermExists(boil.GetDB(), cvtermID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CvtermExistsP checks if the Cvterm row exists. Panics on error.
func CvtermExistsP(exec boil.Executor, cvtermID int) bool {
	e, err := CvtermExists(exec, cvtermID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
