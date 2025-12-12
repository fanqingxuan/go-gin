package db

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"go-gin/internal/errorx"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrMissingWhereClause = errors.New("missing WHERE clause for UPDATE/DELETE operation")

type Model struct {
	ctx        context.Context
	db         *gorm.DB
	table      string
	primaryKey string
	fields     []string
	fieldsEx   []string
	wheres     []whereClause
	data       any
	orderBy    string
	groupBy    string
	having     string
	distinct   bool
	limit      int
	offset     int
	unscoped   bool
}

type whereClause struct {
	query any
	args  []any
	or    bool
}

func newModel(ctx context.Context, table string) *Model {
	return &Model{
		ctx:        ctx,
		db:         instance.WithContext(ctx),
		table:      table,
		primaryKey: "id",
	}
}

// SetPrimaryKey sets the primary key for the model.
func (m *Model) SetPrimaryKey(pk string) *Model {
	model := m.clone()
	model.primaryKey = pk
	return model
}

func (m *Model) clone() *Model {
	newModel := *m
	newModel.fields = append([]string{}, m.fields...)
	newModel.fieldsEx = append([]string{}, m.fieldsEx...)
	newModel.wheres = append([]whereClause{}, m.wheres...)
	return &newModel
}

func (m *Model) Where(where any, args ...any) *Model {
	model := m.clone()
	model.wheres = append(model.wheres, whereClause{query: where, args: args})
	return model
}

func (m *Model) WhereOr(where any, args ...any) *Model {
	model := m.clone()
	model.wheres = append(model.wheres, whereClause{query: where, args: args, or: true})
	return model
}

// WherePri uses primary key as the condition.
func (m *Model) WherePri(value any) *Model {
	return m.Where(m.primaryKey+" = ?", value)
}

// WhereLT builds `column < value` statement.
func (m *Model) WhereLT(column string, value any) *Model {
	return m.Where(column+" < ?", value)
}

// WhereLTE builds `column <= value` statement.
func (m *Model) WhereLTE(column string, value any) *Model {
	return m.Where(column+" <= ?", value)
}

// WhereGT builds `column > value` statement.
func (m *Model) WhereGT(column string, value any) *Model {
	return m.Where(column+" > ?", value)
}

// WhereGTE builds `column >= value` statement.
func (m *Model) WhereGTE(column string, value any) *Model {
	return m.Where(column+" >= ?", value)
}

// WhereBetween builds `column BETWEEN min AND max` statement.
func (m *Model) WhereBetween(column string, min, max any) *Model {
	return m.Where(column+" BETWEEN ? AND ?", min, max)
}

// WhereNotBetween builds `column NOT BETWEEN min AND max` statement.
func (m *Model) WhereNotBetween(column string, min, max any) *Model {
	return m.Where(column+" NOT BETWEEN ? AND ?", min, max)
}

// WhereLike builds `column LIKE like` statement.
func (m *Model) WhereLike(column string, like string) *Model {
	return m.Where(column+" LIKE ?", like)
}

// WhereNotLike builds `column NOT LIKE like` statement.
func (m *Model) WhereNotLike(column string, like string) *Model {
	return m.Where(column+" NOT LIKE ?", like)
}

// WhereIn builds `column IN (in)` statement.
func (m *Model) WhereIn(column string, in any) *Model {
	return m.Where(column+" IN (?)", in)
}

// WhereNotIn builds `column NOT IN (in)` statement.
func (m *Model) WhereNotIn(column string, in any) *Model {
	return m.Where(column+" NOT IN (?)", in)
}

// WhereNull builds `column IS NULL` statement.
func (m *Model) WhereNull(column string) *Model {
	return m.Where(column + " IS NULL")
}

// WhereNotNull builds `column IS NOT NULL` statement.
func (m *Model) WhereNotNull(column string) *Model {
	return m.Where(column + " IS NOT NULL")
}

// WhereNot builds `column != value` statement.
func (m *Model) WhereNot(column string, value any) *Model {
	return m.Where(column+" != ?", value)
}

// Wheref builds condition string using fmt.Sprintf and target arguments.
func (m *Model) Wheref(format string, args ...any) *Model {
	return m.Where(fmt.Sprintf(format, args...))
}

func (m *Model) Data(data any) *Model {
	model := m.clone()
	model.data = data
	return model
}

func (m *Model) Fields(fields ...string) *Model {
	model := m.clone()
	model.fields = append(model.fields, fields...)
	return model
}

func (m *Model) FieldsEx(fields ...string) *Model {
	model := m.clone()
	model.fieldsEx = append(model.fieldsEx, fields...)
	return model
}

func (m *Model) Order(order string) *Model {
	model := m.clone()
	model.orderBy = order
	return model
}

func (m *Model) Group(group string) *Model {
	model := m.clone()
	model.groupBy = group
	return model
}

func (m *Model) Having(having string) *Model {
	model := m.clone()
	model.having = having
	return model
}

func (m *Model) Limit(limit int) *Model {
	model := m.clone()
	model.limit = limit
	return model
}

func (m *Model) Offset(offset int) *Model {
	model := m.clone()
	model.offset = offset
	return model
}

func (m *Model) Page(page, size int) *Model {
	model := m.clone()
	if page < 1 {
		page = 1
	}
	model.offset = (page - 1) * size
	model.limit = size
	return model
}

func (m *Model) Unscoped() *Model {
	model := m.clone()
	model.unscoped = true
	return model
}

func (m *Model) buildQuery() *gorm.DB {
	query := m.db.Table(m.table)

	if m.distinct {
		query = query.Distinct()
	}

	if len(m.fields) > 0 {
		query = query.Select(m.fields)
	}

	for _, w := range m.wheres {
		conditions := structToMap(w.query)
		if conditions != nil {
			if w.or {
				query = query.Or(conditions, w.args...)
			} else {
				query = query.Where(conditions, w.args...)
			}
		} else {
			if w.or {
				query = query.Or(w.query, w.args...)
			} else {
				query = query.Where(w.query, w.args...)
			}
		}
	}

	if m.orderBy != "" {
		query = query.Order(m.orderBy)
	}
	if m.groupBy != "" {
		query = query.Group(m.groupBy)
	}
	if m.having != "" {
		query = query.Having(m.having)
	}
	if m.limit > 0 {
		query = query.Limit(m.limit)
	}
	if m.offset > 0 {
		query = query.Offset(m.offset)
	}
	if m.unscoped {
		query = query.Unscoped()
	}

	return query
}

func (m *Model) One(dest any) error {
	err := m.buildQuery().First(dest).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return errorx.TryToDBError(err)
}

func (m *Model) All(dest any) error {
	return errorx.TryToDBError(m.buildQuery().Find(dest).Error)
}

func (m *Model) Scan(dest any) error {
	return errorx.TryToDBError(m.buildQuery().Scan(dest).Error)
}

func (m *Model) Count() (int64, error) {
	var count int64
	err := m.buildQuery().Count(&count).Error
	return count, errorx.TryToDBError(err)
}

func (m *Model) Exist() (bool, error) {
	count, err := m.Limit(1).Count()
	return count > 0, err
}

func (m *Model) Insert() (int64, error) {
	data := m.prepareData()
	result := m.db.Table(m.table).Create(data)
	return result.RowsAffected, errorx.TryToDBError(result.Error)
}

func (m *Model) InsertAndGetId() (int64, error) {
	data := m.prepareData()
	result := m.db.Table(m.table).Create(data)
	if result.Error != nil {
		return 0, errorx.TryToDBError(result.Error)
	}
	if dataMap, ok := data.(map[string]any); ok {
		if id, exists := dataMap["id"]; exists {
			switch v := id.(type) {
			case int64:
				return v, nil
			case int:
				return int64(v), nil
			}
		}
	}
	return result.RowsAffected, nil
}

func (m *Model) Update() (int64, error) {
	if len(m.wheres) == 0 {
		return 0, ErrMissingWhereClause
	}
	data := m.prepareData()
	result := m.buildQuery().Updates(data)
	return result.RowsAffected, errorx.TryToDBError(result.Error)
}

func (m *Model) Delete() (int64, error) {
	if len(m.wheres) == 0 {
		return 0, ErrMissingWhereClause
	}
	result := m.buildQuery().Delete(nil)
	return result.RowsAffected, errorx.TryToDBError(result.Error)
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement.
func (m *Model) InsertIgnore() (int64, error) {
	data := m.prepareData()
	result := m.db.Table(m.table).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(data)
	return result.RowsAffected, errorx.TryToDBError(result.Error)
}

// Replace does "REPLACE INTO ..." statement.
func (m *Model) Replace() (int64, error) {
	data := m.prepareData()
	result := m.db.Table(m.table).Clauses(clause.Insert{Modifier: "REPLACE"}).Create(data)
	return result.RowsAffected, errorx.TryToDBError(result.Error)
}

// UpdateAndGetAffected performs update and returns affected rows.
func (m *Model) UpdateAndGetAffected() (int64, error) {
	return m.Update()
}

func (m *Model) Pluck(column string, dest any) error {
	return errorx.TryToDBError(m.buildQuery().Pluck(column, dest).Error)
}

func (m *Model) Value(column string, dest any) error {
	return errorx.TryToDBError(m.buildQuery().Select(column).Limit(1).Scan(dest).Error)
}

// Min does "SELECT MIN(column) FROM ..." statement.
func (m *Model) Min(column string) (float64, error) {
	var result float64
	err := m.buildQuery().Select("MIN(" + column + ")").Scan(&result).Error
	return result, errorx.TryToDBError(err)
}

// Max does "SELECT MAX(column) FROM ..." statement.
func (m *Model) Max(column string) (float64, error) {
	var result float64
	err := m.buildQuery().Select("MAX(" + column + ")").Scan(&result).Error
	return result, errorx.TryToDBError(err)
}

// Avg does "SELECT AVG(column) FROM ..." statement.
func (m *Model) Avg(column string) (float64, error) {
	var result float64
	err := m.buildQuery().Select("AVG(" + column + ")").Scan(&result).Error
	return result, errorx.TryToDBError(err)
}

// Sum does "SELECT SUM(column) FROM ..." statement.
func (m *Model) Sum(column string) (float64, error) {
	var result float64
	err := m.buildQuery().Select("SUM(" + column + ")").Scan(&result).Error
	return result, errorx.TryToDBError(err)
}

// Distinct sets DISTINCT for the query.
func (m *Model) Distinct() *Model {
	model := m.clone()
	model.distinct = true
	return model
}

// ScanAndCount scans records and returns total count.
func (m *Model) ScanAndCount(dest any, totalCount *int64) error {
	count, err := m.Count()
	if err != nil {
		return err // already wrapped
	}
	*totalCount = count
	if count == 0 {
		return nil
	}
	return m.All(dest) // already wrapped
}

// ChunkHandler is a function that handles chunk results.
type ChunkHandler func(result []map[string]any, err error) bool

// Chunk iterates the query result with given size and handler function.
func (m *Model) Chunk(size int, handler ChunkHandler) {
	page := 1
	for {
		var result []map[string]any
		err := m.Page(page, size).Scan(&result)
		if err != nil {
			handler(nil, err)
			break
		}
		if len(result) == 0 {
			break
		}
		if !handler(result, nil) {
			break
		}
		if len(result) < size {
			break
		}
		page++
	}
}

// Array returns column values as slice.
func (m *Model) Array(column string) ([]any, error) {
	var result []any
	err := m.buildQuery().Pluck(column, &result).Error
	return result, errorx.TryToDBError(err)
}

// Increment increments a column's value by a given amount.
func (m *Model) Increment(column string, amount any) (int64, error) {
	if len(m.wheres) == 0 {
		return 0, ErrMissingWhereClause
	}
	result := m.buildQuery().UpdateColumn(column, gorm.Expr(column+" + ?", amount))
	return result.RowsAffected, errorx.TryToDBError(result.Error)
}

// Decrement decrements a column's value by a given amount.
func (m *Model) Decrement(column string, amount any) (int64, error) {
	if len(m.wheres) == 0 {
		return 0, ErrMissingWhereClause
	}
	result := m.buildQuery().UpdateColumn(column, gorm.Expr(column+" - ?", amount))
	return result.RowsAffected, errorx.TryToDBError(result.Error)
}

// Save does "INSERT ... ON DUPLICATE KEY UPDATE ..." statement.
func (m *Model) Save() (int64, error) {
	data := m.prepareData()
	result := m.db.Table(m.table).Save(data)
	return result.RowsAffected, errorx.TryToDBError(result.Error)
}

func (m *Model) prepareData() any {
	data := structToMap(m.data)
	if data != nil {
		return data
	}
	return m.data
}

func structToMap(v any) map[string]any {
	if v == nil {
		return nil
	}
	if m, ok := v.(map[string]any); ok {
		return filterNilValues(m)
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil
	}

	rt := rv.Type()
	result := make(map[string]any)

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}
		if field.Kind() == reflect.Interface && field.IsNil() {
			continue
		}
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		fieldName := toSnakeCase(fieldType.Name)
		result[fieldName] = field.Interface()
	}

	if len(result) == 0 {
		return nil
	}
	return result
}

func filterNilValues(m map[string]any) map[string]any {
	result := make(map[string]any)
	for k, v := range m {
		if v != nil {
			result[k] = v
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func toSnakeCase(s string) string {
	var result []byte
	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, byte(c+'a'-'A'))
		} else {
			result = append(result, byte(c))
		}
	}
	return string(result)
}
