// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
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

// GameMatch is an object representing the database table.
type GameMatch struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	GameID    int       `boil:"game_id" json:"game_id" toml:"game_id" yaml:"game_id"`
	UserID    string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *gameMatchR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L gameMatchL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GameMatchColumns = struct {
	ID        string
	GameID    string
	UserID    string
	CreatedAt string
}{
	ID:        "id",
	GameID:    "game_id",
	UserID:    "user_id",
	CreatedAt: "created_at",
}

var GameMatchTableColumns = struct {
	ID        string
	GameID    string
	UserID    string
	CreatedAt string
}{
	ID:        "game_matches.id",
	GameID:    "game_matches.game_id",
	UserID:    "game_matches.user_id",
	CreatedAt: "game_matches.created_at",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var GameMatchWhere = struct {
	ID        whereHelperint
	GameID    whereHelperint
	UserID    whereHelperstring
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"game_matches\".\"id\""},
	GameID:    whereHelperint{field: "\"game_matches\".\"game_id\""},
	UserID:    whereHelperstring{field: "\"game_matches\".\"user_id\""},
	CreatedAt: whereHelpertime_Time{field: "\"game_matches\".\"created_at\""},
}

// GameMatchRels is where relationship names are stored.
var GameMatchRels = struct {
}{}

// gameMatchR is where relationships are stored.
type gameMatchR struct {
}

// NewStruct creates a new relationship struct
func (*gameMatchR) NewStruct() *gameMatchR {
	return &gameMatchR{}
}

// gameMatchL is where Load methods for each relationship are stored.
type gameMatchL struct{}

var (
	gameMatchAllColumns            = []string{"id", "game_id", "user_id", "created_at"}
	gameMatchColumnsWithoutDefault = []string{"game_id", "user_id", "created_at"}
	gameMatchColumnsWithDefault    = []string{"id"}
	gameMatchPrimaryKeyColumns     = []string{"id"}
	gameMatchGeneratedColumns      = []string{}
)

type (
	// GameMatchSlice is an alias for a slice of pointers to GameMatch.
	// This should almost always be used instead of []GameMatch.
	GameMatchSlice []*GameMatch
	// GameMatchHook is the signature for custom GameMatch hook methods
	GameMatchHook func(context.Context, boil.ContextExecutor, *GameMatch) error

	gameMatchQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	gameMatchType                 = reflect.TypeOf(&GameMatch{})
	gameMatchMapping              = queries.MakeStructMapping(gameMatchType)
	gameMatchPrimaryKeyMapping, _ = queries.BindMapping(gameMatchType, gameMatchMapping, gameMatchPrimaryKeyColumns)
	gameMatchInsertCacheMut       sync.RWMutex
	gameMatchInsertCache          = make(map[string]insertCache)
	gameMatchUpdateCacheMut       sync.RWMutex
	gameMatchUpdateCache          = make(map[string]updateCache)
	gameMatchUpsertCacheMut       sync.RWMutex
	gameMatchUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var gameMatchAfterSelectHooks []GameMatchHook

var gameMatchBeforeInsertHooks []GameMatchHook
var gameMatchAfterInsertHooks []GameMatchHook

var gameMatchBeforeUpdateHooks []GameMatchHook
var gameMatchAfterUpdateHooks []GameMatchHook

var gameMatchBeforeDeleteHooks []GameMatchHook
var gameMatchAfterDeleteHooks []GameMatchHook

var gameMatchBeforeUpsertHooks []GameMatchHook
var gameMatchAfterUpsertHooks []GameMatchHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GameMatch) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMatchAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GameMatch) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMatchBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GameMatch) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMatchAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GameMatch) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMatchBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GameMatch) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMatchAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GameMatch) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMatchBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GameMatch) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMatchAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GameMatch) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMatchBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GameMatch) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gameMatchAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGameMatchHook registers your hook function for all future operations.
func AddGameMatchHook(hookPoint boil.HookPoint, gameMatchHook GameMatchHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		gameMatchAfterSelectHooks = append(gameMatchAfterSelectHooks, gameMatchHook)
	case boil.BeforeInsertHook:
		gameMatchBeforeInsertHooks = append(gameMatchBeforeInsertHooks, gameMatchHook)
	case boil.AfterInsertHook:
		gameMatchAfterInsertHooks = append(gameMatchAfterInsertHooks, gameMatchHook)
	case boil.BeforeUpdateHook:
		gameMatchBeforeUpdateHooks = append(gameMatchBeforeUpdateHooks, gameMatchHook)
	case boil.AfterUpdateHook:
		gameMatchAfterUpdateHooks = append(gameMatchAfterUpdateHooks, gameMatchHook)
	case boil.BeforeDeleteHook:
		gameMatchBeforeDeleteHooks = append(gameMatchBeforeDeleteHooks, gameMatchHook)
	case boil.AfterDeleteHook:
		gameMatchAfterDeleteHooks = append(gameMatchAfterDeleteHooks, gameMatchHook)
	case boil.BeforeUpsertHook:
		gameMatchBeforeUpsertHooks = append(gameMatchBeforeUpsertHooks, gameMatchHook)
	case boil.AfterUpsertHook:
		gameMatchAfterUpsertHooks = append(gameMatchAfterUpsertHooks, gameMatchHook)
	}
}

// One returns a single gameMatch record from the query.
func (q gameMatchQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GameMatch, error) {
	o := &GameMatch{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for game_matches")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all GameMatch records from the query.
func (q gameMatchQuery) All(ctx context.Context, exec boil.ContextExecutor) (GameMatchSlice, error) {
	var o []*GameMatch

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to GameMatch slice")
	}

	if len(gameMatchAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all GameMatch records in the query.
func (q gameMatchQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count game_matches rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q gameMatchQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if game_matches exists")
	}

	return count > 0, nil
}

// GameMatches retrieves all the records using an executor.
func GameMatches(mods ...qm.QueryMod) gameMatchQuery {
	mods = append(mods, qm.From("\"game_matches\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"game_matches\".*"})
	}

	return gameMatchQuery{q}
}

// FindGameMatch retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGameMatch(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*GameMatch, error) {
	gameMatchObj := &GameMatch{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"game_matches\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, gameMatchObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from game_matches")
	}

	if err = gameMatchObj.doAfterSelectHooks(ctx, exec); err != nil {
		return gameMatchObj, err
	}

	return gameMatchObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GameMatch) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no game_matches provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gameMatchColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	gameMatchInsertCacheMut.RLock()
	cache, cached := gameMatchInsertCache[key]
	gameMatchInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			gameMatchAllColumns,
			gameMatchColumnsWithDefault,
			gameMatchColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(gameMatchType, gameMatchMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(gameMatchType, gameMatchMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"game_matches\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"game_matches\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into game_matches")
	}

	if !cached {
		gameMatchInsertCacheMut.Lock()
		gameMatchInsertCache[key] = cache
		gameMatchInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the GameMatch.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GameMatch) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	gameMatchUpdateCacheMut.RLock()
	cache, cached := gameMatchUpdateCache[key]
	gameMatchUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			gameMatchAllColumns,
			gameMatchPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update game_matches, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"game_matches\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, gameMatchPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(gameMatchType, gameMatchMapping, append(wl, gameMatchPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update game_matches row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for game_matches")
	}

	if !cached {
		gameMatchUpdateCacheMut.Lock()
		gameMatchUpdateCache[key] = cache
		gameMatchUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q gameMatchQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for game_matches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for game_matches")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GameMatchSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gameMatchPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"game_matches\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, gameMatchPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in gameMatch slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all gameMatch")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GameMatch) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no game_matches provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gameMatchColumnsWithDefault, o)

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

	gameMatchUpsertCacheMut.RLock()
	cache, cached := gameMatchUpsertCache[key]
	gameMatchUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			gameMatchAllColumns,
			gameMatchColumnsWithDefault,
			gameMatchColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			gameMatchAllColumns,
			gameMatchPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert game_matches, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(gameMatchPrimaryKeyColumns))
			copy(conflict, gameMatchPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"game_matches\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(gameMatchType, gameMatchMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(gameMatchType, gameMatchMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert game_matches")
	}

	if !cached {
		gameMatchUpsertCacheMut.Lock()
		gameMatchUpsertCache[key] = cache
		gameMatchUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single GameMatch record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GameMatch) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no GameMatch provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), gameMatchPrimaryKeyMapping)
	sql := "DELETE FROM \"game_matches\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from game_matches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for game_matches")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q gameMatchQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no gameMatchQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from game_matches")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for game_matches")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GameMatchSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(gameMatchBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gameMatchPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"game_matches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gameMatchPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from gameMatch slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for game_matches")
	}

	if len(gameMatchAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *GameMatch) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGameMatch(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GameMatchSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GameMatchSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gameMatchPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"game_matches\".* FROM \"game_matches\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gameMatchPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GameMatchSlice")
	}

	*o = slice

	return nil
}

// GameMatchExists checks if the GameMatch row exists.
func GameMatchExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"game_matches\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if game_matches exists")
	}

	return exists, nil
}

// Exists checks if the GameMatch row exists.
func (o *GameMatch) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return GameMatchExists(ctx, exec, o.ID)
}
