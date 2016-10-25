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

// Pub is an object representing the database table.
type Pub struct {
	PubID       int         `boil:"pub_id" json:"pub_id" toml:"pub_id" yaml:"pub_id"`
	Title       null.String `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Volumetitle null.String `boil:"volumetitle" json:"volumetitle,omitempty" toml:"volumetitle" yaml:"volumetitle,omitempty"`
	Volume      null.String `boil:"volume" json:"volume,omitempty" toml:"volume" yaml:"volume,omitempty"`
	SeriesName  null.String `boil:"series_name" json:"series_name,omitempty" toml:"series_name" yaml:"series_name,omitempty"`
	Issue       null.String `boil:"issue" json:"issue,omitempty" toml:"issue" yaml:"issue,omitempty"`
	Pyear       null.String `boil:"pyear" json:"pyear,omitempty" toml:"pyear" yaml:"pyear,omitempty"`
	Pages       null.String `boil:"pages" json:"pages,omitempty" toml:"pages" yaml:"pages,omitempty"`
	Miniref     null.String `boil:"miniref" json:"miniref,omitempty" toml:"miniref" yaml:"miniref,omitempty"`
	Uniquename  string      `boil:"uniquename" json:"uniquename" toml:"uniquename" yaml:"uniquename"`
	TypeID      int         `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	IsObsolete  null.Bool   `boil:"is_obsolete" json:"is_obsolete,omitempty" toml:"is_obsolete" yaml:"is_obsolete,omitempty"`
	Publisher   null.String `boil:"publisher" json:"publisher,omitempty" toml:"publisher" yaml:"publisher,omitempty"`
	Pubplace    null.String `boil:"pubplace" json:"pubplace,omitempty" toml:"pubplace" yaml:"pubplace,omitempty"`

	R *pubR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pubL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// pubR is where relationships are stored.
type pubR struct {
	Type                       *Cvterm
	StockPub                   *StockPub
	FeaturelocPub              *FeaturelocPub
	PubDbxref                  *PubDbxref
	ObjectPubRelationship      *PubRelationship
	SubjectPubRelationship     *PubRelationship
	Pubauthor                  *Pubauthor
	Pubprop                    *Pubprop
	FeatureCvtermPub           *FeatureCvtermPub
	StockCvterm                *StockCvterm
	StockRelationshipPub       *StockRelationshipPub
	FeatureCvterm              *FeatureCvterm
	FeaturePub                 *FeaturePub
	FeatureSynonym             *FeatureSynonym
	FeatureRelationshippropPub *FeatureRelationshippropPub
	FeaturepropPub             *FeaturepropPub
	FeatureRelationshipPub     *FeatureRelationshipPub
	Phendesc                   *Phendesc
	Phenstatement              *Phenstatement
	StockpropPub               *StockpropPub
	PhenotypeComparison        *PhenotypeComparison
	StockRelationshipCvterms   StockRelationshipCvtermSlice
	PhenotypeComparisonCvterms PhenotypeComparisonCvtermSlice
}

// pubL is where Load methods for each relationship are stored.
type pubL struct{}

var (
	pubColumns               = []string{"pub_id", "title", "volumetitle", "volume", "series_name", "issue", "pyear", "pages", "miniref", "uniquename", "type_id", "is_obsolete", "publisher", "pubplace"}
	pubColumnsWithoutDefault = []string{"title", "volumetitle", "volume", "series_name", "issue", "pyear", "pages", "miniref", "uniquename", "type_id", "publisher", "pubplace"}
	pubColumnsWithDefault    = []string{"pub_id", "is_obsolete"}
	pubPrimaryKeyColumns     = []string{"pub_id"}
)

type (
	// PubSlice is an alias for a slice of pointers to Pub.
	// This should generally be used opposed to []Pub.
	PubSlice []*Pub
	// PubHook is the signature for custom Pub hook methods
	PubHook func(boil.Executor, *Pub) error

	pubQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pubType                 = reflect.TypeOf(&Pub{})
	pubMapping              = queries.MakeStructMapping(pubType)
	pubPrimaryKeyMapping, _ = queries.BindMapping(pubType, pubMapping, pubPrimaryKeyColumns)
	pubInsertCacheMut       sync.RWMutex
	pubInsertCache          = make(map[string]insertCache)
	pubUpdateCacheMut       sync.RWMutex
	pubUpdateCache          = make(map[string]updateCache)
	pubUpsertCacheMut       sync.RWMutex
	pubUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var pubBeforeInsertHooks []PubHook
var pubBeforeUpdateHooks []PubHook
var pubBeforeDeleteHooks []PubHook
var pubBeforeUpsertHooks []PubHook

var pubAfterInsertHooks []PubHook
var pubAfterSelectHooks []PubHook
var pubAfterUpdateHooks []PubHook
var pubAfterDeleteHooks []PubHook
var pubAfterUpsertHooks []PubHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Pub) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range pubBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Pub) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range pubBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Pub) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range pubBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Pub) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range pubBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Pub) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range pubAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Pub) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range pubAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Pub) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range pubAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Pub) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range pubAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Pub) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range pubAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPubHook registers your hook function for all future operations.
func AddPubHook(hookPoint boil.HookPoint, pubHook PubHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		pubBeforeInsertHooks = append(pubBeforeInsertHooks, pubHook)
	case boil.BeforeUpdateHook:
		pubBeforeUpdateHooks = append(pubBeforeUpdateHooks, pubHook)
	case boil.BeforeDeleteHook:
		pubBeforeDeleteHooks = append(pubBeforeDeleteHooks, pubHook)
	case boil.BeforeUpsertHook:
		pubBeforeUpsertHooks = append(pubBeforeUpsertHooks, pubHook)
	case boil.AfterInsertHook:
		pubAfterInsertHooks = append(pubAfterInsertHooks, pubHook)
	case boil.AfterSelectHook:
		pubAfterSelectHooks = append(pubAfterSelectHooks, pubHook)
	case boil.AfterUpdateHook:
		pubAfterUpdateHooks = append(pubAfterUpdateHooks, pubHook)
	case boil.AfterDeleteHook:
		pubAfterDeleteHooks = append(pubAfterDeleteHooks, pubHook)
	case boil.AfterUpsertHook:
		pubAfterUpsertHooks = append(pubAfterUpsertHooks, pubHook)
	}
}

// OneP returns a single pub record from the query, and panics on error.
func (q pubQuery) OneP() *Pub {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single pub record from the query.
func (q pubQuery) One() (*Pub, error) {
	o := &Pub{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pub")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Pub records from the query, and panics on error.
func (q pubQuery) AllP() PubSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Pub records from the query.
func (q pubQuery) All() (PubSlice, error) {
	var o PubSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Pub slice")
	}

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Pub records in the query, and panics on error.
func (q pubQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Pub records in the query.
func (q pubQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pub rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q pubQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q pubQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pub exists")
	}

	return count > 0, nil
}

// TypeG pointed to by the foreign key.
func (o *Pub) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *Pub) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// StockPubG pointed to by the foreign key.
func (o *Pub) StockPubG(mods ...qm.QueryMod) stockPubQuery {
	return o.StockPub(boil.GetDB(), mods...)
}

// StockPub pointed to by the foreign key.
func (o *Pub) StockPub(exec boil.Executor, mods ...qm.QueryMod) stockPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := StockPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_pub\"")

	return query
}

// FeaturelocPubG pointed to by the foreign key.
func (o *Pub) FeaturelocPubG(mods ...qm.QueryMod) featurelocPubQuery {
	return o.FeaturelocPub(boil.GetDB(), mods...)
}

// FeaturelocPub pointed to by the foreign key.
func (o *Pub) FeaturelocPub(exec boil.Executor, mods ...qm.QueryMod) featurelocPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := FeaturelocPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"featureloc_pub\"")

	return query
}

// PubDbxrefG pointed to by the foreign key.
func (o *Pub) PubDbxrefG(mods ...qm.QueryMod) pubDbxrefQuery {
	return o.PubDbxref(boil.GetDB(), mods...)
}

// PubDbxref pointed to by the foreign key.
func (o *Pub) PubDbxref(exec boil.Executor, mods ...qm.QueryMod) pubDbxrefQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := PubDbxrefs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub_dbxref\"")

	return query
}

// ObjectPubRelationshipG pointed to by the foreign key.
func (o *Pub) ObjectPubRelationshipG(mods ...qm.QueryMod) pubRelationshipQuery {
	return o.ObjectPubRelationship(boil.GetDB(), mods...)
}

// ObjectPubRelationship pointed to by the foreign key.
func (o *Pub) ObjectPubRelationship(exec boil.Executor, mods ...qm.QueryMod) pubRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("object_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := PubRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub_relationship\"")

	return query
}

// SubjectPubRelationshipG pointed to by the foreign key.
func (o *Pub) SubjectPubRelationshipG(mods ...qm.QueryMod) pubRelationshipQuery {
	return o.SubjectPubRelationship(boil.GetDB(), mods...)
}

// SubjectPubRelationship pointed to by the foreign key.
func (o *Pub) SubjectPubRelationship(exec boil.Executor, mods ...qm.QueryMod) pubRelationshipQuery {
	queryMods := []qm.QueryMod{
		qm.Where("subject_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := PubRelationships(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pub_relationship\"")

	return query
}

// PubauthorG pointed to by the foreign key.
func (o *Pub) PubauthorG(mods ...qm.QueryMod) pubauthorQuery {
	return o.Pubauthor(boil.GetDB(), mods...)
}

// Pubauthor pointed to by the foreign key.
func (o *Pub) Pubauthor(exec boil.Executor, mods ...qm.QueryMod) pubauthorQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubauthors(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pubauthor\"")

	return query
}

// PubpropG pointed to by the foreign key.
func (o *Pub) PubpropG(mods ...qm.QueryMod) pubpropQuery {
	return o.Pubprop(boil.GetDB(), mods...)
}

// Pubprop pointed to by the foreign key.
func (o *Pub) Pubprop(exec boil.Executor, mods ...qm.QueryMod) pubpropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Pubprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pubprop\"")

	return query
}

// FeatureCvtermPubG pointed to by the foreign key.
func (o *Pub) FeatureCvtermPubG(mods ...qm.QueryMod) featureCvtermPubQuery {
	return o.FeatureCvtermPub(boil.GetDB(), mods...)
}

// FeatureCvtermPub pointed to by the foreign key.
func (o *Pub) FeatureCvtermPub(exec boil.Executor, mods ...qm.QueryMod) featureCvtermPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureCvtermPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_cvterm_pub\"")

	return query
}

// StockCvtermG pointed to by the foreign key.
func (o *Pub) StockCvtermG(mods ...qm.QueryMod) stockCvtermQuery {
	return o.StockCvterm(boil.GetDB(), mods...)
}

// StockCvterm pointed to by the foreign key.
func (o *Pub) StockCvterm(exec boil.Executor, mods ...qm.QueryMod) stockCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := StockCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_cvterm\"")

	return query
}

// StockRelationshipPubG pointed to by the foreign key.
func (o *Pub) StockRelationshipPubG(mods ...qm.QueryMod) stockRelationshipPubQuery {
	return o.StockRelationshipPub(boil.GetDB(), mods...)
}

// StockRelationshipPub pointed to by the foreign key.
func (o *Pub) StockRelationshipPub(exec boil.Executor, mods ...qm.QueryMod) stockRelationshipPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := StockRelationshipPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_relationship_pub\"")

	return query
}

// FeatureCvtermG pointed to by the foreign key.
func (o *Pub) FeatureCvtermG(mods ...qm.QueryMod) featureCvtermQuery {
	return o.FeatureCvterm(boil.GetDB(), mods...)
}

// FeatureCvterm pointed to by the foreign key.
func (o *Pub) FeatureCvterm(exec boil.Executor, mods ...qm.QueryMod) featureCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_cvterm\"")

	return query
}

// FeaturePubG pointed to by the foreign key.
func (o *Pub) FeaturePubG(mods ...qm.QueryMod) featurePubQuery {
	return o.FeaturePub(boil.GetDB(), mods...)
}

// FeaturePub pointed to by the foreign key.
func (o *Pub) FeaturePub(exec boil.Executor, mods ...qm.QueryMod) featurePubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := FeaturePubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_pub\"")

	return query
}

// FeatureSynonymG pointed to by the foreign key.
func (o *Pub) FeatureSynonymG(mods ...qm.QueryMod) featureSynonymQuery {
	return o.FeatureSynonym(boil.GetDB(), mods...)
}

// FeatureSynonym pointed to by the foreign key.
func (o *Pub) FeatureSynonym(exec boil.Executor, mods ...qm.QueryMod) featureSynonymQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureSynonyms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_synonym\"")

	return query
}

// FeatureRelationshippropPubG pointed to by the foreign key.
func (o *Pub) FeatureRelationshippropPubG(mods ...qm.QueryMod) featureRelationshippropPubQuery {
	return o.FeatureRelationshippropPub(boil.GetDB(), mods...)
}

// FeatureRelationshippropPub pointed to by the foreign key.
func (o *Pub) FeatureRelationshippropPub(exec boil.Executor, mods ...qm.QueryMod) featureRelationshippropPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureRelationshippropPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_relationshipprop_pub\"")

	return query
}

// FeaturepropPubG pointed to by the foreign key.
func (o *Pub) FeaturepropPubG(mods ...qm.QueryMod) featurepropPubQuery {
	return o.FeaturepropPub(boil.GetDB(), mods...)
}

// FeaturepropPub pointed to by the foreign key.
func (o *Pub) FeaturepropPub(exec boil.Executor, mods ...qm.QueryMod) featurepropPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := FeaturepropPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"featureprop_pub\"")

	return query
}

// FeatureRelationshipPubG pointed to by the foreign key.
func (o *Pub) FeatureRelationshipPubG(mods ...qm.QueryMod) featureRelationshipPubQuery {
	return o.FeatureRelationshipPub(boil.GetDB(), mods...)
}

// FeatureRelationshipPub pointed to by the foreign key.
func (o *Pub) FeatureRelationshipPub(exec boil.Executor, mods ...qm.QueryMod) featureRelationshipPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := FeatureRelationshipPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature_relationship_pub\"")

	return query
}

// PhendescG pointed to by the foreign key.
func (o *Pub) PhendescG(mods ...qm.QueryMod) phendescQuery {
	return o.Phendesc(boil.GetDB(), mods...)
}

// Phendesc pointed to by the foreign key.
func (o *Pub) Phendesc(exec boil.Executor, mods ...qm.QueryMod) phendescQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Phendescs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phendesc\"")

	return query
}

// PhenstatementG pointed to by the foreign key.
func (o *Pub) PhenstatementG(mods ...qm.QueryMod) phenstatementQuery {
	return o.Phenstatement(boil.GetDB(), mods...)
}

// Phenstatement pointed to by the foreign key.
func (o *Pub) Phenstatement(exec boil.Executor, mods ...qm.QueryMod) phenstatementQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := Phenstatements(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenstatement\"")

	return query
}

// StockpropPubG pointed to by the foreign key.
func (o *Pub) StockpropPubG(mods ...qm.QueryMod) stockpropPubQuery {
	return o.StockpropPub(boil.GetDB(), mods...)
}

// StockpropPub pointed to by the foreign key.
func (o *Pub) StockpropPub(exec boil.Executor, mods ...qm.QueryMod) stockpropPubQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := StockpropPubs(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stockprop_pub\"")

	return query
}

// PhenotypeComparisonG pointed to by the foreign key.
func (o *Pub) PhenotypeComparisonG(mods ...qm.QueryMod) phenotypeComparisonQuery {
	return o.PhenotypeComparison(boil.GetDB(), mods...)
}

// PhenotypeComparison pointed to by the foreign key.
func (o *Pub) PhenotypeComparison(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonQuery {
	queryMods := []qm.QueryMod{
		qm.Where("pub_id=$1", o.PubID),
	}

	queryMods = append(queryMods, mods...)

	query := PhenotypeComparisons(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_comparison\"")

	return query
}

// StockRelationshipCvtermsG retrieves all the stock_relationship_cvterm's stock relationship cvterm.
func (o *Pub) StockRelationshipCvtermsG(mods ...qm.QueryMod) stockRelationshipCvtermQuery {
	return o.StockRelationshipCvterms(boil.GetDB(), mods...)
}

// StockRelationshipCvterms retrieves all the stock_relationship_cvterm's stock relationship cvterm with an executor.
func (o *Pub) StockRelationshipCvterms(exec boil.Executor, mods ...qm.QueryMod) stockRelationshipCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"pub_id\"=$1", o.PubID),
	)

	query := StockRelationshipCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_relationship_cvterm\" as \"a\"")
	return query
}

// PhenotypeComparisonCvtermsG retrieves all the phenotype_comparison_cvterm's phenotype comparison cvterm.
func (o *Pub) PhenotypeComparisonCvtermsG(mods ...qm.QueryMod) phenotypeComparisonCvtermQuery {
	return o.PhenotypeComparisonCvterms(boil.GetDB(), mods...)
}

// PhenotypeComparisonCvterms retrieves all the phenotype_comparison_cvterm's phenotype comparison cvterm with an executor.
func (o *Pub) PhenotypeComparisonCvterms(exec boil.Executor, mods ...qm.QueryMod) phenotypeComparisonCvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"pub_id\"=$1", o.PubID),
	)

	query := PhenotypeComparisonCvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"phenotype_comparison_cvterm\" as \"a\"")
	return query
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadType(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
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

	if len(pubAfterSelectHooks) != 0 {
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

// LoadStockPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadStockPub(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_pub\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockPub")
	}
	defer results.Close()

	var resultSlice []*StockPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockPub")
	}

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.StockPub = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeaturelocPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadFeaturelocPub(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"featureloc_pub\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeaturelocPub")
	}
	defer results.Close()

	var resultSlice []*FeaturelocPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeaturelocPub")
	}

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeaturelocPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.FeaturelocPub = foreign
				break
			}
		}
	}

	return nil
}

// LoadPubDbxref allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadPubDbxref(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"pub_dbxref\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PubDbxref")
	}
	defer results.Close()

	var resultSlice []*PubDbxref
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PubDbxref")
	}

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.PubDbxref = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.PubDbxref = foreign
				break
			}
		}
	}

	return nil
}

// LoadObjectPubRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadObjectPubRelationship(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"pub_relationship\" where \"object_id\" in (%s)",
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

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.ObjectPubRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.ObjectID {
				local.R.ObjectPubRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadSubjectPubRelationship allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadSubjectPubRelationship(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"pub_relationship\" where \"subject_id\" in (%s)",
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

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.SubjectPubRelationship = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.SubjectID {
				local.R.SubjectPubRelationship = foreign
				break
			}
		}
	}

	return nil
}

// LoadPubauthor allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadPubauthor(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"pubauthor\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Pubauthor")
	}
	defer results.Close()

	var resultSlice []*Pubauthor
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Pubauthor")
	}

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Pubauthor = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.Pubauthor = foreign
				break
			}
		}
	}

	return nil
}

// LoadPubprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadPubprop(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"pubprop\" where \"pub_id\" in (%s)",
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

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Pubprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.Pubprop = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureCvtermPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadFeatureCvtermPub(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_cvterm_pub\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureCvtermPub")
	}
	defer results.Close()

	var resultSlice []*FeatureCvtermPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureCvtermPub")
	}

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureCvtermPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.FeatureCvtermPub = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadStockCvterm(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_cvterm\" where \"pub_id\" in (%s)",
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

	if len(pubAfterSelectHooks) != 0 {
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
			if local.PubID == foreign.PubID {
				local.R.StockCvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockRelationshipPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadStockRelationshipPub(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_relationship_pub\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockRelationshipPub")
	}
	defer results.Close()

	var resultSlice []*StockRelationshipPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockRelationshipPub")
	}

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockRelationshipPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.StockRelationshipPub = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureCvterm allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadFeatureCvterm(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_cvterm\" where \"pub_id\" in (%s)",
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

	if len(pubAfterSelectHooks) != 0 {
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
			if local.PubID == foreign.PubID {
				local.R.FeatureCvterm = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeaturePub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadFeaturePub(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_pub\" where \"pub_id\" in (%s)",
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

	if len(pubAfterSelectHooks) != 0 {
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
			if local.PubID == foreign.PubID {
				local.R.FeaturePub = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureSynonym allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadFeatureSynonym(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_synonym\" where \"pub_id\" in (%s)",
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

	if len(pubAfterSelectHooks) != 0 {
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
			if local.PubID == foreign.PubID {
				local.R.FeatureSynonym = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureRelationshippropPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadFeatureRelationshippropPub(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_relationshipprop_pub\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureRelationshippropPub")
	}
	defer results.Close()

	var resultSlice []*FeatureRelationshippropPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureRelationshippropPub")
	}

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureRelationshippropPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.FeatureRelationshippropPub = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeaturepropPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadFeaturepropPub(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"featureprop_pub\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeaturepropPub")
	}
	defer results.Close()

	var resultSlice []*FeaturepropPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeaturepropPub")
	}

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeaturepropPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.FeaturepropPub = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeatureRelationshipPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadFeatureRelationshipPub(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature_relationship_pub\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FeatureRelationshipPub")
	}
	defer results.Close()

	var resultSlice []*FeatureRelationshipPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FeatureRelationshipPub")
	}

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.FeatureRelationshipPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.FeatureRelationshipPub = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhendesc allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadPhendesc(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phendesc\" where \"pub_id\" in (%s)",
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

	if len(pubAfterSelectHooks) != 0 {
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
			if local.PubID == foreign.PubID {
				local.R.Phendesc = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenstatement allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadPhenstatement(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenstatement\" where \"pub_id\" in (%s)",
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

	if len(pubAfterSelectHooks) != 0 {
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
			if local.PubID == foreign.PubID {
				local.R.Phenstatement = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockpropPub allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadStockpropPub(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stockprop_pub\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StockpropPub")
	}
	defer results.Close()

	var resultSlice []*StockpropPub
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StockpropPub")
	}

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.StockpropPub = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.StockpropPub = foreign
				break
			}
		}
	}

	return nil
}

// LoadPhenotypeComparison allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadPhenotypeComparison(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_comparison\" where \"pub_id\" in (%s)",
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

	if len(pubAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.PhenotypeComparison = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.PhenotypeComparison = foreign
				break
			}
		}
	}

	return nil
}

// LoadStockRelationshipCvterms allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadStockRelationshipCvterms(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_relationship_cvterm\" where \"pub_id\" in (%s)",
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
			if local.PubID == foreign.PubID.Int {
				local.R.StockRelationshipCvterms = append(local.R.StockRelationshipCvterms, foreign)
				break
			}
		}
	}

	return nil
}

// LoadPhenotypeComparisonCvterms allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (pubL) LoadPhenotypeComparisonCvterms(e boil.Executor, singular bool, maybePub interface{}) error {
	var slice []*Pub
	var object *Pub

	count := 1
	if singular {
		object = maybePub.(*Pub)
	} else {
		slice = *maybePub.(*PubSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &pubR{}
		args[0] = object.PubID
	} else {
		for i, obj := range slice {
			obj.R = &pubR{}
			args[i] = obj.PubID
		}
	}

	query := fmt.Sprintf(
		"select * from \"phenotype_comparison_cvterm\" where \"pub_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load phenotype_comparison_cvterm")
	}
	defer results.Close()

	var resultSlice []*PhenotypeComparisonCvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice phenotype_comparison_cvterm")
	}

	if len(phenotypeComparisonCvtermAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.PhenotypeComparisonCvterms = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PubID == foreign.PubID {
				local.R.PhenotypeComparisonCvterms = append(local.R.PhenotypeComparisonCvterms, foreign)
				break
			}
		}
	}

	return nil
}

// SetType of the pub to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypePubs.
func (o *Pub) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"pub\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, pubPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.PubID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &pubR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypePubs: PubSlice{o},
		}
	} else {
		related.R.TypePubs = append(related.R.TypePubs, o)
	}

	return nil
}

// SetStockPub of the pub to the related item.
// Sets o.R.StockPub to related.
// Adds o to related.R.Pub.
func (o *Pub) SetStockPub(exec boil.Executor, insert bool, related *StockPub) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockPubPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.StockPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			StockPub: related,
		}
	} else {
		o.R.StockPub = related
	}

	if related.R == nil {
		related.R = &stockPubR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetFeaturelocPub of the pub to the related item.
// Sets o.R.FeaturelocPub to related.
// Adds o to related.R.Pub.
func (o *Pub) SetFeaturelocPub(exec boil.Executor, insert bool, related *FeaturelocPub) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"featureloc_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurelocPubPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.FeaturelocPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			FeaturelocPub: related,
		}
	} else {
		o.R.FeaturelocPub = related
	}

	if related.R == nil {
		related.R = &featurelocPubR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetPubDbxref of the pub to the related item.
// Sets o.R.PubDbxref to related.
// Adds o to related.R.Pub.
func (o *Pub) SetPubDbxref(exec boil.Executor, insert bool, related *PubDbxref) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"pub_dbxref\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, pubDbxrefPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.PubDbxrefID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			PubDbxref: related,
		}
	} else {
		o.R.PubDbxref = related
	}

	if related.R == nil {
		related.R = &pubDbxrefR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetObjectPubRelationship of the pub to the related item.
// Sets o.R.ObjectPubRelationship to related.
// Adds o to related.R.Object.
func (o *Pub) SetObjectPubRelationship(exec boil.Executor, insert bool, related *PubRelationship) error {
	var err error

	if insert {
		related.ObjectID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"pub_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"object_id"}),
			strmangle.WhereClause("\"", "\"", 2, pubRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.PubRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.ObjectID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			ObjectPubRelationship: related,
		}
	} else {
		o.R.ObjectPubRelationship = related
	}

	if related.R == nil {
		related.R = &pubRelationshipR{
			Object: o,
		}
	} else {
		related.R.Object = o
	}
	return nil
}

// SetSubjectPubRelationship of the pub to the related item.
// Sets o.R.SubjectPubRelationship to related.
// Adds o to related.R.Subject.
func (o *Pub) SetSubjectPubRelationship(exec boil.Executor, insert bool, related *PubRelationship) error {
	var err error

	if insert {
		related.SubjectID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"pub_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"subject_id"}),
			strmangle.WhereClause("\"", "\"", 2, pubRelationshipPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.PubRelationshipID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.SubjectID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			SubjectPubRelationship: related,
		}
	} else {
		o.R.SubjectPubRelationship = related
	}

	if related.R == nil {
		related.R = &pubRelationshipR{
			Subject: o,
		}
	} else {
		related.R.Subject = o
	}
	return nil
}

// SetPubauthor of the pub to the related item.
// Sets o.R.Pubauthor to related.
// Adds o to related.R.Pub.
func (o *Pub) SetPubauthor(exec boil.Executor, insert bool, related *Pubauthor) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"pubauthor\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, pubauthorPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.PubauthorID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			Pubauthor: related,
		}
	} else {
		o.R.Pubauthor = related
	}

	if related.R == nil {
		related.R = &pubauthorR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetPubprop of the pub to the related item.
// Sets o.R.Pubprop to related.
// Adds o to related.R.Pub.
func (o *Pub) SetPubprop(exec boil.Executor, insert bool, related *Pubprop) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"pubprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, pubpropPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.PubpropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			Pubprop: related,
		}
	} else {
		o.R.Pubprop = related
	}

	if related.R == nil {
		related.R = &pubpropR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetFeatureCvtermPub of the pub to the related item.
// Sets o.R.FeatureCvtermPub to related.
// Adds o to related.R.Pub.
func (o *Pub) SetFeatureCvtermPub(exec boil.Executor, insert bool, related *FeatureCvtermPub) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_cvterm_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureCvtermPubPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.FeatureCvtermPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			FeatureCvtermPub: related,
		}
	} else {
		o.R.FeatureCvtermPub = related
	}

	if related.R == nil {
		related.R = &featureCvtermPubR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetStockCvterm of the pub to the related item.
// Sets o.R.StockCvterm to related.
// Adds o to related.R.Pub.
func (o *Pub) SetStockCvterm(exec boil.Executor, insert bool, related *StockCvterm) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockCvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.StockCvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			StockCvterm: related,
		}
	} else {
		o.R.StockCvterm = related
	}

	if related.R == nil {
		related.R = &stockCvtermR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetStockRelationshipPub of the pub to the related item.
// Sets o.R.StockRelationshipPub to related.
// Adds o to related.R.Pub.
func (o *Pub) SetStockRelationshipPub(exec boil.Executor, insert bool, related *StockRelationshipPub) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stock_relationship_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockRelationshipPubPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.StockRelationshipPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			StockRelationshipPub: related,
		}
	} else {
		o.R.StockRelationshipPub = related
	}

	if related.R == nil {
		related.R = &stockRelationshipPubR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetFeatureCvterm of the pub to the related item.
// Sets o.R.FeatureCvterm to related.
// Adds o to related.R.Pub.
func (o *Pub) SetFeatureCvterm(exec boil.Executor, insert bool, related *FeatureCvterm) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_cvterm\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureCvtermPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.FeatureCvtermID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			FeatureCvterm: related,
		}
	} else {
		o.R.FeatureCvterm = related
	}

	if related.R == nil {
		related.R = &featureCvtermR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetFeaturePub of the pub to the related item.
// Sets o.R.FeaturePub to related.
// Adds o to related.R.Pub.
func (o *Pub) SetFeaturePub(exec boil.Executor, insert bool, related *FeaturePub) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurePubPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.FeaturePubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			FeaturePub: related,
		}
	} else {
		o.R.FeaturePub = related
	}

	if related.R == nil {
		related.R = &featurePubR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetFeatureSynonym of the pub to the related item.
// Sets o.R.FeatureSynonym to related.
// Adds o to related.R.Pub.
func (o *Pub) SetFeatureSynonym(exec boil.Executor, insert bool, related *FeatureSynonym) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_synonym\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureSynonymPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.FeatureSynonymID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			FeatureSynonym: related,
		}
	} else {
		o.R.FeatureSynonym = related
	}

	if related.R == nil {
		related.R = &featureSynonymR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetFeatureRelationshippropPub of the pub to the related item.
// Sets o.R.FeatureRelationshippropPub to related.
// Adds o to related.R.Pub.
func (o *Pub) SetFeatureRelationshippropPub(exec boil.Executor, insert bool, related *FeatureRelationshippropPub) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_relationshipprop_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureRelationshippropPubPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.FeatureRelationshippropPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			FeatureRelationshippropPub: related,
		}
	} else {
		o.R.FeatureRelationshippropPub = related
	}

	if related.R == nil {
		related.R = &featureRelationshippropPubR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetFeaturepropPub of the pub to the related item.
// Sets o.R.FeaturepropPub to related.
// Adds o to related.R.Pub.
func (o *Pub) SetFeaturepropPub(exec boil.Executor, insert bool, related *FeaturepropPub) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"featureprop_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, featurepropPubPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.FeaturepropPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			FeaturepropPub: related,
		}
	} else {
		o.R.FeaturepropPub = related
	}

	if related.R == nil {
		related.R = &featurepropPubR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetFeatureRelationshipPub of the pub to the related item.
// Sets o.R.FeatureRelationshipPub to related.
// Adds o to related.R.Pub.
func (o *Pub) SetFeatureRelationshipPub(exec boil.Executor, insert bool, related *FeatureRelationshipPub) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"feature_relationship_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, featureRelationshipPubPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.FeatureRelationshipPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			FeatureRelationshipPub: related,
		}
	} else {
		o.R.FeatureRelationshipPub = related
	}

	if related.R == nil {
		related.R = &featureRelationshipPubR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetPhendesc of the pub to the related item.
// Sets o.R.Phendesc to related.
// Adds o to related.R.Pub.
func (o *Pub) SetPhendesc(exec boil.Executor, insert bool, related *Phendesc) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phendesc\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, phendescPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.PhendescID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			Phendesc: related,
		}
	} else {
		o.R.Phendesc = related
	}

	if related.R == nil {
		related.R = &phendescR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetPhenstatement of the pub to the related item.
// Sets o.R.Phenstatement to related.
// Adds o to related.R.Pub.
func (o *Pub) SetPhenstatement(exec boil.Executor, insert bool, related *Phenstatement) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenstatement\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenstatementPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.PhenstatementID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			Phenstatement: related,
		}
	} else {
		o.R.Phenstatement = related
	}

	if related.R == nil {
		related.R = &phenstatementR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetStockpropPub of the pub to the related item.
// Sets o.R.StockpropPub to related.
// Adds o to related.R.Pub.
func (o *Pub) SetStockpropPub(exec boil.Executor, insert bool, related *StockpropPub) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"stockprop_pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, stockpropPubPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.StockpropPubID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			StockpropPub: related,
		}
	} else {
		o.R.StockpropPub = related
	}

	if related.R == nil {
		related.R = &stockpropPubR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// SetPhenotypeComparison of the pub to the related item.
// Sets o.R.PhenotypeComparison to related.
// Adds o to related.R.Pub.
func (o *Pub) SetPhenotypeComparison(exec boil.Executor, insert bool, related *PhenotypeComparison) error {
	var err error

	if insert {
		related.PubID = o.PubID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"phenotype_comparison\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"pub_id"}),
			strmangle.WhereClause("\"", "\"", 2, phenotypeComparisonPrimaryKeyColumns),
		)
		values := []interface{}{o.PubID, related.PhenotypeComparisonID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PubID = o.PubID

	}

	if o.R == nil {
		o.R = &pubR{
			PhenotypeComparison: related,
		}
	} else {
		o.R.PhenotypeComparison = related
	}

	if related.R == nil {
		related.R = &phenotypeComparisonR{
			Pub: o,
		}
	} else {
		related.R.Pub = o
	}
	return nil
}

// AddStockRelationshipCvterms adds the given related objects to the existing relationships
// of the pub, optionally inserting them as new records.
// Appends related to o.R.StockRelationshipCvterms.
// Sets related.R.Pub appropriately.
func (o *Pub) AddStockRelationshipCvterms(exec boil.Executor, insert bool, related ...*StockRelationshipCvterm) error {
	var err error
	for _, rel := range related {
		rel.PubID.Int = o.PubID
		rel.PubID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "pub_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &pubR{
			StockRelationshipCvterms: related,
		}
	} else {
		o.R.StockRelationshipCvterms = append(o.R.StockRelationshipCvterms, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &stockRelationshipCvtermR{
				Pub: o,
			}
		} else {
			rel.R.Pub = o
		}
	}
	return nil
}

// SetStockRelationshipCvterms removes all previously related items of the
// pub replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Pub's StockRelationshipCvterms accordingly.
// Replaces o.R.StockRelationshipCvterms with related.
// Sets related.R.Pub's StockRelationshipCvterms accordingly.
func (o *Pub) SetStockRelationshipCvterms(exec boil.Executor, insert bool, related ...*StockRelationshipCvterm) error {
	query := "update \"stock_relationship_cvterm\" set \"pub_id\" = null where \"pub_id\" = $1"
	values := []interface{}{o.PubID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.StockRelationshipCvterms {
			rel.PubID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Pub = nil
		}

		o.R.StockRelationshipCvterms = nil
	}
	return o.AddStockRelationshipCvterms(exec, insert, related...)
}

// RemoveStockRelationshipCvterms relationships from objects passed in.
// Removes related items from R.StockRelationshipCvterms (uses pointer comparison, removal does not keep order)
// Sets related.R.Pub.
func (o *Pub) RemoveStockRelationshipCvterms(exec boil.Executor, related ...*StockRelationshipCvterm) error {
	var err error
	for _, rel := range related {
		rel.PubID.Valid = false
		if rel.R != nil {
			rel.R.Pub = nil
		}
		if err = rel.Update(exec, "pub_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.StockRelationshipCvterms {
			if rel != ri {
				continue
			}

			ln := len(o.R.StockRelationshipCvterms)
			if ln > 1 && i < ln-1 {
				o.R.StockRelationshipCvterms[i] = o.R.StockRelationshipCvterms[ln-1]
			}
			o.R.StockRelationshipCvterms = o.R.StockRelationshipCvterms[:ln-1]
			break
		}
	}

	return nil
}

// AddPhenotypeComparisonCvterms adds the given related objects to the existing relationships
// of the pub, optionally inserting them as new records.
// Appends related to o.R.PhenotypeComparisonCvterms.
// Sets related.R.Pub appropriately.
func (o *Pub) AddPhenotypeComparisonCvterms(exec boil.Executor, insert bool, related ...*PhenotypeComparisonCvterm) error {
	var err error
	for _, rel := range related {
		rel.PubID = o.PubID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "pub_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &pubR{
			PhenotypeComparisonCvterms: related,
		}
	} else {
		o.R.PhenotypeComparisonCvterms = append(o.R.PhenotypeComparisonCvterms, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &phenotypeComparisonCvtermR{
				Pub: o,
			}
		} else {
			rel.R.Pub = o
		}
	}
	return nil
}

// PubsG retrieves all records.
func PubsG(mods ...qm.QueryMod) pubQuery {
	return Pubs(boil.GetDB(), mods...)
}

// Pubs retrieves all the records using an executor.
func Pubs(exec boil.Executor, mods ...qm.QueryMod) pubQuery {
	mods = append(mods, qm.From("\"pub\""))
	return pubQuery{NewQuery(exec, mods...)}
}

// FindPubG retrieves a single record by ID.
func FindPubG(pubID int, selectCols ...string) (*Pub, error) {
	return FindPub(boil.GetDB(), pubID, selectCols...)
}

// FindPubGP retrieves a single record by ID, and panics on error.
func FindPubGP(pubID int, selectCols ...string) *Pub {
	retobj, err := FindPub(boil.GetDB(), pubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindPub retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPub(exec boil.Executor, pubID int, selectCols ...string) (*Pub, error) {
	pubObj := &Pub{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pub\" where \"pub_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, pubID)

	err := q.Bind(pubObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pub")
	}

	return pubObj, nil
}

// FindPubP retrieves a single record by ID with an executor, and panics on error.
func FindPubP(exec boil.Executor, pubID int, selectCols ...string) *Pub {
	retobj, err := FindPub(exec, pubID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Pub) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Pub) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Pub) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Pub) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no pub provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pubColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	pubInsertCacheMut.RLock()
	cache, cached := pubInsertCache[key]
	pubInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			pubColumns,
			pubColumnsWithDefault,
			pubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(pubType, pubMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pubType, pubMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"pub\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into pub")
	}

	if !cached {
		pubInsertCacheMut.Lock()
		pubInsertCache[key] = cache
		pubInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Pub record. See Update for
// whitelist behavior description.
func (o *Pub) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Pub record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Pub) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Pub, and panics on error.
// See Update for whitelist behavior description.
func (o *Pub) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Pub.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Pub) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	pubUpdateCacheMut.RLock()
	cache, cached := pubUpdateCache[key]
	pubUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(pubColumns, pubPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update pub, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pub\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pubPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pubType, pubMapping, append(wl, pubPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update pub row")
	}

	if !cached {
		pubUpdateCacheMut.Lock()
		pubUpdateCache[key] = cache
		pubUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q pubQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q pubQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for pub")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PubSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o PubSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o PubSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PubSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"pub\" SET %s WHERE (\"pub_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(pubPrimaryKeyColumns), len(colNames)+1, len(pubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in pub slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Pub) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Pub) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Pub) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Pub) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no pub provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pubColumnsWithDefault, o)

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

	pubUpsertCacheMut.RLock()
	cache, cached := pubUpsertCache[key]
	pubUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			pubColumns,
			pubColumnsWithDefault,
			pubColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			pubColumns,
			pubPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert pub, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pubPrimaryKeyColumns))
			copy(conflict, pubPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"pub\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(pubType, pubMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pubType, pubMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for pub")
	}

	if !cached {
		pubUpsertCacheMut.Lock()
		pubUpsertCache[key] = cache
		pubUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Pub record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Pub) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Pub record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Pub) DeleteG() error {
	if o == nil {
		return errors.New("models: no Pub provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Pub record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Pub) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Pub record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Pub) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Pub provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pubPrimaryKeyMapping)
	sql := "DELETE FROM \"pub\" WHERE \"pub_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from pub")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q pubQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q pubQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no pubQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from pub")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o PubSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o PubSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Pub slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o PubSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PubSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Pub slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(pubBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, pubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(pubPrimaryKeyColumns), 1, len(pubPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from pub slice")
	}

	if len(pubAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Pub) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Pub) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Pub) ReloadG() error {
	if o == nil {
		return errors.New("models: no Pub provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Pub) Reload(exec boil.Executor) error {
	ret, err := FindPub(exec, o.PubID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PubSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PubSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PubSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty PubSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PubSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	pubs := PubSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pubPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"pub\".* FROM \"pub\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, pubPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(pubPrimaryKeyColumns), 1, len(pubPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&pubs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PubSlice")
	}

	*o = pubs

	return nil
}

// PubExists checks if the Pub row exists.
func PubExists(exec boil.Executor, pubID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"pub\" where \"pub_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, pubID)
	}

	row := exec.QueryRow(sql, pubID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pub exists")
	}

	return exists, nil
}

// PubExistsG checks if the Pub row exists.
func PubExistsG(pubID int) (bool, error) {
	return PubExists(boil.GetDB(), pubID)
}

// PubExistsGP checks if the Pub row exists. Panics on error.
func PubExistsGP(pubID int) bool {
	e, err := PubExists(boil.GetDB(), pubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// PubExistsP checks if the Pub row exists. Panics on error.
func PubExistsP(exec boil.Executor, pubID int) bool {
	e, err := PubExists(exec, pubID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
