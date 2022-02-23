// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Artist is an object representing the database table.
type Artist struct {
	ID   int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Slug string `boil:"slug" json:"slug" toml:"slug" yaml:"slug"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *artistR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L artistL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ArtistColumns = struct {
	ID   string
	Slug string
	Name string
}{
	ID:   "id",
	Slug: "slug",
	Name: "name",
}

var ArtistTableColumns = struct {
	ID   string
	Slug string
	Name string
}{
	ID:   "artist.id",
	Slug: "artist.slug",
	Name: "artist.name",
}

// Generated where

var ArtistWhere = struct {
	ID   whereHelperint64
	Slug whereHelperstring
	Name whereHelperstring
}{
	ID:   whereHelperint64{field: "\"artist\".\"id\""},
	Slug: whereHelperstring{field: "\"artist\".\"slug\""},
	Name: whereHelperstring{field: "\"artist\".\"name\""},
}

// ArtistRels is where relationship names are stored.
var ArtistRels = struct {
	Archives string
}{
	Archives: "Archives",
}

// artistR is where relationships are stored.
type artistR struct {
	Archives ArchiveSlice `boil:"Archives" json:"Archives" toml:"Archives" yaml:"Archives"`
}

// NewStruct creates a new relationship struct
func (*artistR) NewStruct() *artistR {
	return &artistR{}
}

// artistL is where Load methods for each relationship are stored.
type artistL struct{}

var (
	artistAllColumns            = []string{"id", "slug", "name"}
	artistColumnsWithoutDefault = []string{}
	artistColumnsWithDefault    = []string{"id", "slug", "name"}
	artistPrimaryKeyColumns     = []string{"id"}
	artistGeneratedColumns      = []string{}
)

type (
	// ArtistSlice is an alias for a slice of pointers to Artist.
	// This should almost always be used instead of []Artist.
	ArtistSlice []*Artist
	// ArtistHook is the signature for custom Artist hook methods
	ArtistHook func(boil.Executor, *Artist) error

	artistQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	artistType                 = reflect.TypeOf(&Artist{})
	artistMapping              = queries.MakeStructMapping(artistType)
	artistPrimaryKeyMapping, _ = queries.BindMapping(artistType, artistMapping, artistPrimaryKeyColumns)
	artistInsertCacheMut       sync.RWMutex
	artistInsertCache          = make(map[string]insertCache)
	artistUpdateCacheMut       sync.RWMutex
	artistUpdateCache          = make(map[string]updateCache)
	artistUpsertCacheMut       sync.RWMutex
	artistUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var artistAfterSelectHooks []ArtistHook

var artistBeforeInsertHooks []ArtistHook
var artistAfterInsertHooks []ArtistHook

var artistBeforeUpdateHooks []ArtistHook
var artistAfterUpdateHooks []ArtistHook

var artistBeforeDeleteHooks []ArtistHook
var artistAfterDeleteHooks []ArtistHook

var artistBeforeUpsertHooks []ArtistHook
var artistAfterUpsertHooks []ArtistHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Artist) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range artistAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Artist) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range artistBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Artist) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range artistAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Artist) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range artistBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Artist) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range artistAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Artist) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range artistBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Artist) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range artistAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Artist) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range artistBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Artist) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range artistAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddArtistHook registers your hook function for all future operations.
func AddArtistHook(hookPoint boil.HookPoint, artistHook ArtistHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		artistAfterSelectHooks = append(artistAfterSelectHooks, artistHook)
	case boil.BeforeInsertHook:
		artistBeforeInsertHooks = append(artistBeforeInsertHooks, artistHook)
	case boil.AfterInsertHook:
		artistAfterInsertHooks = append(artistAfterInsertHooks, artistHook)
	case boil.BeforeUpdateHook:
		artistBeforeUpdateHooks = append(artistBeforeUpdateHooks, artistHook)
	case boil.AfterUpdateHook:
		artistAfterUpdateHooks = append(artistAfterUpdateHooks, artistHook)
	case boil.BeforeDeleteHook:
		artistBeforeDeleteHooks = append(artistBeforeDeleteHooks, artistHook)
	case boil.AfterDeleteHook:
		artistAfterDeleteHooks = append(artistAfterDeleteHooks, artistHook)
	case boil.BeforeUpsertHook:
		artistBeforeUpsertHooks = append(artistBeforeUpsertHooks, artistHook)
	case boil.AfterUpsertHook:
		artistAfterUpsertHooks = append(artistAfterUpsertHooks, artistHook)
	}
}

// OneG returns a single artist record from the query using the global executor.
func (q artistQuery) OneG() (*Artist, error) {
	return q.One(boil.GetDB())
}

// One returns a single artist record from the query.
func (q artistQuery) One(exec boil.Executor) (*Artist, error) {
	o := &Artist{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for artist")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Artist records from the query using the global executor.
func (q artistQuery) AllG() (ArtistSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Artist records from the query.
func (q artistQuery) All(exec boil.Executor) (ArtistSlice, error) {
	var o []*Artist

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Artist slice")
	}

	if len(artistAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Artist records in the query, and panics on error.
func (q artistQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Artist records in the query.
func (q artistQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count artist rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q artistQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q artistQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if artist exists")
	}

	return count > 0, nil
}

// Archives retrieves all the archive's Archives with an executor.
func (o *Artist) Archives(mods ...qm.QueryMod) archiveQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"archive_artists\" on \"archive\".\"id\" = \"archive_artists\".\"archive_id\""),
		qm.Where("\"archive_artists\".\"artist_id\"=?", o.ID),
	)

	query := Archives(queryMods...)
	queries.SetFrom(query.Query, "\"archive\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"archive\".*"})
	}

	return query
}

// LoadArchives allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (artistL) LoadArchives(e boil.Executor, singular bool, maybeArtist interface{}, mods queries.Applicator) error {
	var slice []*Artist
	var object *Artist

	if singular {
		object = maybeArtist.(*Artist)
	} else {
		slice = *maybeArtist.(*[]*Artist)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &artistR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &artistR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("\"archive\".id, \"archive\".path, \"archive\".created_at, \"archive\".updated_at, \"archive\".published_at, \"archive\".title, \"archive\".slug, \"archive\".pages, \"archive\".size, \"archive\".circle_id, \"archive\".magazine_id, \"archive\".parody_id, \"a\".\"artist_id\""),
		qm.From("\"archive\""),
		qm.InnerJoin("\"archive_artists\" as \"a\" on \"archive\".\"id\" = \"a\".\"archive_id\""),
		qm.WhereIn("\"a\".\"artist_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load archive")
	}

	var resultSlice []*Archive

	var localJoinCols []int64
	for results.Next() {
		one := new(Archive)
		var localJoinCol int64

		err = results.Scan(&one.ID, &one.Path, &one.CreatedAt, &one.UpdatedAt, &one.PublishedAt, &one.Title, &one.Slug, &one.Pages, &one.Size, &one.CircleID, &one.MagazineID, &one.ParodyID, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for archive")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice archive")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on archive")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for archive")
	}

	if len(archiveAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Archives = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &archiveR{}
			}
			foreign.R.Artists = append(foreign.R.Artists, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Archives = append(local.R.Archives, foreign)
				if foreign.R == nil {
					foreign.R = &archiveR{}
				}
				foreign.R.Artists = append(foreign.R.Artists, local)
				break
			}
		}
	}

	return nil
}

// AddArchivesG adds the given related objects to the existing relationships
// of the artist, optionally inserting them as new records.
// Appends related to o.R.Archives.
// Sets related.R.Artists appropriately.
// Uses the global database handle.
func (o *Artist) AddArchivesG(insert bool, related ...*Archive) error {
	return o.AddArchives(boil.GetDB(), insert, related...)
}

// AddArchives adds the given related objects to the existing relationships
// of the artist, optionally inserting them as new records.
// Appends related to o.R.Archives.
// Sets related.R.Artists appropriately.
func (o *Artist) AddArchives(exec boil.Executor, insert bool, related ...*Archive) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"archive_artists\" (\"artist_id\", \"archive_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, query)
			fmt.Fprintln(boil.DebugWriter, values)
		}
		_, err = exec.Exec(query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &artistR{
			Archives: related,
		}
	} else {
		o.R.Archives = append(o.R.Archives, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &archiveR{
				Artists: ArtistSlice{o},
			}
		} else {
			rel.R.Artists = append(rel.R.Artists, o)
		}
	}
	return nil
}

// SetArchivesG removes all previously related items of the
// artist replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Artists's Archives accordingly.
// Replaces o.R.Archives with related.
// Sets related.R.Artists's Archives accordingly.
// Uses the global database handle.
func (o *Artist) SetArchivesG(insert bool, related ...*Archive) error {
	return o.SetArchives(boil.GetDB(), insert, related...)
}

// SetArchives removes all previously related items of the
// artist replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Artists's Archives accordingly.
// Replaces o.R.Archives with related.
// Sets related.R.Artists's Archives accordingly.
func (o *Artist) SetArchives(exec boil.Executor, insert bool, related ...*Archive) error {
	query := "delete from \"archive_artists\" where \"artist_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeArchivesFromArtistsSlice(o, related)
	if o.R != nil {
		o.R.Archives = nil
	}
	return o.AddArchives(exec, insert, related...)
}

// RemoveArchivesG relationships from objects passed in.
// Removes related items from R.Archives (uses pointer comparison, removal does not keep order)
// Sets related.R.Artists.
// Uses the global database handle.
func (o *Artist) RemoveArchivesG(related ...*Archive) error {
	return o.RemoveArchives(boil.GetDB(), related...)
}

// RemoveArchives relationships from objects passed in.
// Removes related items from R.Archives (uses pointer comparison, removal does not keep order)
// Sets related.R.Artists.
func (o *Artist) RemoveArchives(exec boil.Executor, related ...*Archive) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	query := fmt.Sprintf(
		"delete from \"archive_artists\" where \"artist_id\" = $1 and \"archive_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err = exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeArchivesFromArtistsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Archives {
			if rel != ri {
				continue
			}

			ln := len(o.R.Archives)
			if ln > 1 && i < ln-1 {
				o.R.Archives[i] = o.R.Archives[ln-1]
			}
			o.R.Archives = o.R.Archives[:ln-1]
			break
		}
	}

	return nil
}

func removeArchivesFromArtistsSlice(o *Artist, related []*Archive) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Artists {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Artists)
			if ln > 1 && i < ln-1 {
				rel.R.Artists[i] = rel.R.Artists[ln-1]
			}
			rel.R.Artists = rel.R.Artists[:ln-1]
			break
		}
	}
}

// Artists retrieves all the records using an executor.
func Artists(mods ...qm.QueryMod) artistQuery {
	mods = append(mods, qm.From("\"artist\""))
	return artistQuery{NewQuery(mods...)}
}

// FindArtistG retrieves a single record by ID.
func FindArtistG(iD int64, selectCols ...string) (*Artist, error) {
	return FindArtist(boil.GetDB(), iD, selectCols...)
}

// FindArtist retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindArtist(exec boil.Executor, iD int64, selectCols ...string) (*Artist, error) {
	artistObj := &Artist{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"artist\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, artistObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from artist")
	}

	if err = artistObj.doAfterSelectHooks(exec); err != nil {
		return artistObj, err
	}

	return artistObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Artist) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Artist) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no artist provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(artistColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	artistInsertCacheMut.RLock()
	cache, cached := artistInsertCache[key]
	artistInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			artistAllColumns,
			artistColumnsWithDefault,
			artistColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(artistType, artistMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(artistType, artistMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"artist\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"artist\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
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
		return errors.Wrap(err, "models: unable to insert into artist")
	}

	if !cached {
		artistInsertCacheMut.Lock()
		artistInsertCache[key] = cache
		artistInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Artist record using the global executor.
// See Update for more documentation.
func (o *Artist) UpdateG(columns boil.Columns) error {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Artist.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Artist) Update(exec boil.Executor, columns boil.Columns) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(columns, nil)
	artistUpdateCacheMut.RLock()
	cache, cached := artistUpdateCache[key]
	artistUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			artistAllColumns,
			artistPrimaryKeyColumns,
		)
		if len(wl) == 0 {
			return errors.New("models: unable to update artist, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"artist\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, artistPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(artistType, artistMapping, append(wl, artistPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update artist row")
	}

	if !cached {
		artistUpdateCacheMut.Lock()
		artistUpdateCache[key] = cache
		artistUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q artistQuery) UpdateAllG(cols M) error {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q artistQuery) UpdateAll(exec boil.Executor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for artist")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ArtistSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ArtistSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), artistPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"artist\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, artistPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in artist slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Artist) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Artist) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no artist provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(artistColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
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
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	artistUpsertCacheMut.RLock()
	cache, cached := artistUpsertCache[key]
	artistUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			artistAllColumns,
			artistColumnsWithDefault,
			artistColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			artistAllColumns,
			artistPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert artist, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(artistPrimaryKeyColumns))
			copy(conflict, artistPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"artist\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(artistType, artistMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(artistType, artistMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert artist")
	}

	if !cached {
		artistUpsertCacheMut.Lock()
		artistUpsertCache[key] = cache
		artistUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single Artist record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Artist) DeleteG() error {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Artist record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Artist) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Artist provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), artistPrimaryKeyMapping)
	sql := "DELETE FROM \"artist\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from artist")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

func (q artistQuery) DeleteAllG() error {
	return q.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all matching rows.
func (q artistQuery) DeleteAll(exec boil.Executor) error {
	if q.Query == nil {
		return errors.New("models: no artistQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec(exec)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from artist")
	}

	return nil
}

// DeleteAllG deletes all rows in the slice.
func (o ArtistSlice) DeleteAllG() error {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ArtistSlice) DeleteAll(exec boil.Executor) error {
	if len(o) == 0 {
		return nil
	}

	if len(artistBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), artistPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"artist\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, artistPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from artist slice")
	}

	if len(artistAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Artist) ReloadG() error {
	if o == nil {
		return errors.New("models: no Artist provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Artist) Reload(exec boil.Executor) error {
	ret, err := FindArtist(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ArtistSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ArtistSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ArtistSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ArtistSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), artistPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"artist\".* FROM \"artist\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, artistPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ArtistSlice")
	}

	*o = slice

	return nil
}

// ArtistExistsG checks if the Artist row exists.
func ArtistExistsG(iD int64) (bool, error) {
	return ArtistExists(boil.GetDB(), iD)
}

// ArtistExists checks if the Artist row exists.
func ArtistExists(exec boil.Executor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"artist\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if artist exists")
	}

	return exists, nil
}
