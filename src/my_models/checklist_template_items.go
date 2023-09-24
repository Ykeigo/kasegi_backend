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

// ChecklistTemplateItem is an object representing the database table.
type ChecklistTemplateItem struct {
	ID                  string `boil:"id" json:"id" toml:"id" yaml:"id"`
	ChecklistTemplateID string `boil:"checklist_template_id" json:"checklist_template_id" toml:"checklist_template_id" yaml:"checklist_template_id"`
	Title               string `boil:"title" json:"title" toml:"title" yaml:"title"`
	IsChecked           bool   `boil:"is_checked" json:"is_checked" toml:"is_checked" yaml:"is_checked"`

	R *checklistTemplateItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L checklistTemplateItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChecklistTemplateItemColumns = struct {
	ID                  string
	ChecklistTemplateID string
	Title               string
	IsChecked           string
}{
	ID:                  "id",
	ChecklistTemplateID: "checklist_template_id",
	Title:               "title",
	IsChecked:           "is_checked",
}

var ChecklistTemplateItemTableColumns = struct {
	ID                  string
	ChecklistTemplateID string
	Title               string
	IsChecked           string
}{
	ID:                  "checklist_template_items.id",
	ChecklistTemplateID: "checklist_template_items.checklist_template_id",
	Title:               "checklist_template_items.title",
	IsChecked:           "checklist_template_items.is_checked",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var ChecklistTemplateItemWhere = struct {
	ID                  whereHelperstring
	ChecklistTemplateID whereHelperstring
	Title               whereHelperstring
	IsChecked           whereHelperbool
}{
	ID:                  whereHelperstring{field: "\"checklist_template_items\".\"id\""},
	ChecklistTemplateID: whereHelperstring{field: "\"checklist_template_items\".\"checklist_template_id\""},
	Title:               whereHelperstring{field: "\"checklist_template_items\".\"title\""},
	IsChecked:           whereHelperbool{field: "\"checklist_template_items\".\"is_checked\""},
}

// ChecklistTemplateItemRels is where relationship names are stored.
var ChecklistTemplateItemRels = struct {
}{}

// checklistTemplateItemR is where relationships are stored.
type checklistTemplateItemR struct {
}

// NewStruct creates a new relationship struct
func (*checklistTemplateItemR) NewStruct() *checklistTemplateItemR {
	return &checklistTemplateItemR{}
}

// checklistTemplateItemL is where Load methods for each relationship are stored.
type checklistTemplateItemL struct{}

var (
	checklistTemplateItemAllColumns            = []string{"id", "checklist_template_id", "title", "is_checked"}
	checklistTemplateItemColumnsWithoutDefault = []string{"id", "checklist_template_id", "title", "is_checked"}
	checklistTemplateItemColumnsWithDefault    = []string{}
	checklistTemplateItemPrimaryKeyColumns     = []string{"id"}
	checklistTemplateItemGeneratedColumns      = []string{}
)

type (
	// ChecklistTemplateItemSlice is an alias for a slice of pointers to ChecklistTemplateItem.
	// This should almost always be used instead of []ChecklistTemplateItem.
	ChecklistTemplateItemSlice []*ChecklistTemplateItem
	// ChecklistTemplateItemHook is the signature for custom ChecklistTemplateItem hook methods
	ChecklistTemplateItemHook func(context.Context, boil.ContextExecutor, *ChecklistTemplateItem) error

	checklistTemplateItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	checklistTemplateItemType                 = reflect.TypeOf(&ChecklistTemplateItem{})
	checklistTemplateItemMapping              = queries.MakeStructMapping(checklistTemplateItemType)
	checklistTemplateItemPrimaryKeyMapping, _ = queries.BindMapping(checklistTemplateItemType, checklistTemplateItemMapping, checklistTemplateItemPrimaryKeyColumns)
	checklistTemplateItemInsertCacheMut       sync.RWMutex
	checklistTemplateItemInsertCache          = make(map[string]insertCache)
	checklistTemplateItemUpdateCacheMut       sync.RWMutex
	checklistTemplateItemUpdateCache          = make(map[string]updateCache)
	checklistTemplateItemUpsertCacheMut       sync.RWMutex
	checklistTemplateItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var checklistTemplateItemAfterSelectHooks []ChecklistTemplateItemHook

var checklistTemplateItemBeforeInsertHooks []ChecklistTemplateItemHook
var checklistTemplateItemAfterInsertHooks []ChecklistTemplateItemHook

var checklistTemplateItemBeforeUpdateHooks []ChecklistTemplateItemHook
var checklistTemplateItemAfterUpdateHooks []ChecklistTemplateItemHook

var checklistTemplateItemBeforeDeleteHooks []ChecklistTemplateItemHook
var checklistTemplateItemAfterDeleteHooks []ChecklistTemplateItemHook

var checklistTemplateItemBeforeUpsertHooks []ChecklistTemplateItemHook
var checklistTemplateItemAfterUpsertHooks []ChecklistTemplateItemHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ChecklistTemplateItem) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checklistTemplateItemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ChecklistTemplateItem) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checklistTemplateItemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ChecklistTemplateItem) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checklistTemplateItemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ChecklistTemplateItem) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checklistTemplateItemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ChecklistTemplateItem) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checklistTemplateItemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ChecklistTemplateItem) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checklistTemplateItemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ChecklistTemplateItem) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checklistTemplateItemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ChecklistTemplateItem) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checklistTemplateItemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ChecklistTemplateItem) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range checklistTemplateItemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChecklistTemplateItemHook registers your hook function for all future operations.
func AddChecklistTemplateItemHook(hookPoint boil.HookPoint, checklistTemplateItemHook ChecklistTemplateItemHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		checklistTemplateItemAfterSelectHooks = append(checklistTemplateItemAfterSelectHooks, checklistTemplateItemHook)
	case boil.BeforeInsertHook:
		checklistTemplateItemBeforeInsertHooks = append(checklistTemplateItemBeforeInsertHooks, checklistTemplateItemHook)
	case boil.AfterInsertHook:
		checklistTemplateItemAfterInsertHooks = append(checklistTemplateItemAfterInsertHooks, checklistTemplateItemHook)
	case boil.BeforeUpdateHook:
		checklistTemplateItemBeforeUpdateHooks = append(checklistTemplateItemBeforeUpdateHooks, checklistTemplateItemHook)
	case boil.AfterUpdateHook:
		checklistTemplateItemAfterUpdateHooks = append(checklistTemplateItemAfterUpdateHooks, checklistTemplateItemHook)
	case boil.BeforeDeleteHook:
		checklistTemplateItemBeforeDeleteHooks = append(checklistTemplateItemBeforeDeleteHooks, checklistTemplateItemHook)
	case boil.AfterDeleteHook:
		checklistTemplateItemAfterDeleteHooks = append(checklistTemplateItemAfterDeleteHooks, checklistTemplateItemHook)
	case boil.BeforeUpsertHook:
		checklistTemplateItemBeforeUpsertHooks = append(checklistTemplateItemBeforeUpsertHooks, checklistTemplateItemHook)
	case boil.AfterUpsertHook:
		checklistTemplateItemAfterUpsertHooks = append(checklistTemplateItemAfterUpsertHooks, checklistTemplateItemHook)
	}
}

// One returns a single checklistTemplateItem record from the query.
func (q checklistTemplateItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ChecklistTemplateItem, error) {
	o := &ChecklistTemplateItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for checklist_template_items")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ChecklistTemplateItem records from the query.
func (q checklistTemplateItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (ChecklistTemplateItemSlice, error) {
	var o []*ChecklistTemplateItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ChecklistTemplateItem slice")
	}

	if len(checklistTemplateItemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ChecklistTemplateItem records in the query.
func (q checklistTemplateItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count checklist_template_items rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q checklistTemplateItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if checklist_template_items exists")
	}

	return count > 0, nil
}

// ChecklistTemplateItems retrieves all the records using an executor.
func ChecklistTemplateItems(mods ...qm.QueryMod) checklistTemplateItemQuery {
	mods = append(mods, qm.From("\"checklist_template_items\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"checklist_template_items\".*"})
	}

	return checklistTemplateItemQuery{q}
}

// FindChecklistTemplateItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChecklistTemplateItem(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*ChecklistTemplateItem, error) {
	checklistTemplateItemObj := &ChecklistTemplateItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"checklist_template_items\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, checklistTemplateItemObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from checklist_template_items")
	}

	if err = checklistTemplateItemObj.doAfterSelectHooks(ctx, exec); err != nil {
		return checklistTemplateItemObj, err
	}

	return checklistTemplateItemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ChecklistTemplateItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no checklist_template_items provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(checklistTemplateItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	checklistTemplateItemInsertCacheMut.RLock()
	cache, cached := checklistTemplateItemInsertCache[key]
	checklistTemplateItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			checklistTemplateItemAllColumns,
			checklistTemplateItemColumnsWithDefault,
			checklistTemplateItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(checklistTemplateItemType, checklistTemplateItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(checklistTemplateItemType, checklistTemplateItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"checklist_template_items\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"checklist_template_items\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into checklist_template_items")
	}

	if !cached {
		checklistTemplateItemInsertCacheMut.Lock()
		checklistTemplateItemInsertCache[key] = cache
		checklistTemplateItemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ChecklistTemplateItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ChecklistTemplateItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	checklistTemplateItemUpdateCacheMut.RLock()
	cache, cached := checklistTemplateItemUpdateCache[key]
	checklistTemplateItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			checklistTemplateItemAllColumns,
			checklistTemplateItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update checklist_template_items, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"checklist_template_items\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, checklistTemplateItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(checklistTemplateItemType, checklistTemplateItemMapping, append(wl, checklistTemplateItemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update checklist_template_items row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for checklist_template_items")
	}

	if !cached {
		checklistTemplateItemUpdateCacheMut.Lock()
		checklistTemplateItemUpdateCache[key] = cache
		checklistTemplateItemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q checklistTemplateItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for checklist_template_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for checklist_template_items")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChecklistTemplateItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), checklistTemplateItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"checklist_template_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, checklistTemplateItemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in checklistTemplateItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all checklistTemplateItem")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ChecklistTemplateItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no checklist_template_items provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(checklistTemplateItemColumnsWithDefault, o)

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

	checklistTemplateItemUpsertCacheMut.RLock()
	cache, cached := checklistTemplateItemUpsertCache[key]
	checklistTemplateItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			checklistTemplateItemAllColumns,
			checklistTemplateItemColumnsWithDefault,
			checklistTemplateItemColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			checklistTemplateItemAllColumns,
			checklistTemplateItemPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert checklist_template_items, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(checklistTemplateItemPrimaryKeyColumns))
			copy(conflict, checklistTemplateItemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"checklist_template_items\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(checklistTemplateItemType, checklistTemplateItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(checklistTemplateItemType, checklistTemplateItemMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert checklist_template_items")
	}

	if !cached {
		checklistTemplateItemUpsertCacheMut.Lock()
		checklistTemplateItemUpsertCache[key] = cache
		checklistTemplateItemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ChecklistTemplateItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ChecklistTemplateItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ChecklistTemplateItem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), checklistTemplateItemPrimaryKeyMapping)
	sql := "DELETE FROM \"checklist_template_items\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from checklist_template_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for checklist_template_items")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q checklistTemplateItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no checklistTemplateItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from checklist_template_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for checklist_template_items")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChecklistTemplateItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(checklistTemplateItemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), checklistTemplateItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"checklist_template_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, checklistTemplateItemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from checklistTemplateItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for checklist_template_items")
	}

	if len(checklistTemplateItemAfterDeleteHooks) != 0 {
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
func (o *ChecklistTemplateItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindChecklistTemplateItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChecklistTemplateItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChecklistTemplateItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), checklistTemplateItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"checklist_template_items\".* FROM \"checklist_template_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, checklistTemplateItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChecklistTemplateItemSlice")
	}

	*o = slice

	return nil
}

// ChecklistTemplateItemExists checks if the ChecklistTemplateItem row exists.
func ChecklistTemplateItemExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"checklist_template_items\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if checklist_template_items exists")
	}

	return exists, nil
}

// Exists checks if the ChecklistTemplateItem row exists.
func (o *ChecklistTemplateItem) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ChecklistTemplateItemExists(ctx, exec, o.ID)
}
