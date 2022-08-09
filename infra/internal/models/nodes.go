// Code generated by SQLBoiler 4.12.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Node is an object representing the database table.
type Node struct {
	ID              int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	State           string     `boil:"state" json:"state" toml:"state" yaml:"state"`
	StartedAt       int64      `boil:"started_at" json:"startedAt" toml:"startedAt" yaml:"startedAt"`
	StoppedAt       null.Int64 `boil:"stopped_at" json:"stoppedAt,omitempty" toml:"stoppedAt" yaml:"stoppedAt,omitempty"`
	ProviderName    string     `boil:"provider_name" json:"providerName" toml:"providerName" yaml:"providerName"`
	ProviderID      string     `boil:"provider_id" json:"providerID" toml:"providerID" yaml:"providerID"`
	Name            string     `boil:"name" json:"name" toml:"name" yaml:"name"`
	Memory          int        `boil:"memory" json:"memory" toml:"memory" yaml:"memory"`
	CPUs            int        `boil:"cpus" json:"cpus" toml:"cpus" yaml:"cpus"`
	Disk            int        `boil:"disk" json:"disk" toml:"disk" yaml:"disk"`
	IPV4            string     `boil:"ip_v4" json:"ipV4" toml:"ipV4" yaml:"ipV4"`
	IPV6            string     `boil:"ip_v6" json:"ipV6" toml:"ipV6" yaml:"ipV6"`
	RegionName      string     `boil:"region_name" json:"regionName" toml:"regionName" yaml:"regionName"`
	RegionLat       float64    `boil:"region_lat" json:"regionLat" toml:"regionLat" yaml:"regionLat"`
	RegionLng       float64    `boil:"region_lng" json:"regionLNG" toml:"regionLNG" yaml:"regionLNG"`
	SKUName         string     `boil:"sku_name" json:"skuName" toml:"skuName" yaml:"skuName"`
	SKUNetworkCap   int        `boil:"sku_network_cap" json:"skuNetworkCap" toml:"skuNetworkCap" yaml:"skuNetworkCap"`
	SKUNetworkSpeed int        `boil:"sku_network_speed" json:"skuNetworkSpeed" toml:"skuNetworkSpeed" yaml:"skuNetworkSpeed"`
	SKUPriceMonthly float32    `boil:"sku_price_monthly" json:"skuPriceMonthly" toml:"skuPriceMonthly" yaml:"skuPriceMonthly"`
	SKUPriceHourly  float32    `boil:"sku_price_hourly" json:"skuPriceHourly" toml:"skuPriceHourly" yaml:"skuPriceHourly"`
	WireguardKey    string     `boil:"wireguard_key" json:"wireguardKey" toml:"wireguardKey" yaml:"wireguardKey"`
	User            string     `boil:"user" json:"user" toml:"user" yaml:"user"`
	Type            string     `boil:"type" json:"type" toml:"type" yaml:"type"`

	R *nodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L nodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NodeColumns = struct {
	ID              string
	State           string
	StartedAt       string
	StoppedAt       string
	ProviderName    string
	ProviderID      string
	Name            string
	Memory          string
	CPUs            string
	Disk            string
	IPV4            string
	IPV6            string
	RegionName      string
	RegionLat       string
	RegionLng       string
	SKUName         string
	SKUNetworkCap   string
	SKUNetworkSpeed string
	SKUPriceMonthly string
	SKUPriceHourly  string
	WireguardKey    string
	User            string
	Type            string
}{
	ID:              "id",
	State:           "state",
	StartedAt:       "started_at",
	StoppedAt:       "stopped_at",
	ProviderName:    "provider_name",
	ProviderID:      "provider_id",
	Name:            "name",
	Memory:          "memory",
	CPUs:            "cpus",
	Disk:            "disk",
	IPV4:            "ip_v4",
	IPV6:            "ip_v6",
	RegionName:      "region_name",
	RegionLat:       "region_lat",
	RegionLng:       "region_lng",
	SKUName:         "sku_name",
	SKUNetworkCap:   "sku_network_cap",
	SKUNetworkSpeed: "sku_network_speed",
	SKUPriceMonthly: "sku_price_monthly",
	SKUPriceHourly:  "sku_price_hourly",
	WireguardKey:    "wireguard_key",
	User:            "user",
	Type:            "type",
}

var NodeTableColumns = struct {
	ID              string
	State           string
	StartedAt       string
	StoppedAt       string
	ProviderName    string
	ProviderID      string
	Name            string
	Memory          string
	CPUs            string
	Disk            string
	IPV4            string
	IPV6            string
	RegionName      string
	RegionLat       string
	RegionLng       string
	SKUName         string
	SKUNetworkCap   string
	SKUNetworkSpeed string
	SKUPriceMonthly string
	SKUPriceHourly  string
	WireguardKey    string
	User            string
	Type            string
}{
	ID:              "nodes.id",
	State:           "nodes.state",
	StartedAt:       "nodes.started_at",
	StoppedAt:       "nodes.stopped_at",
	ProviderName:    "nodes.provider_name",
	ProviderID:      "nodes.provider_id",
	Name:            "nodes.name",
	Memory:          "nodes.memory",
	CPUs:            "nodes.cpus",
	Disk:            "nodes.disk",
	IPV4:            "nodes.ip_v4",
	IPV6:            "nodes.ip_v6",
	RegionName:      "nodes.region_name",
	RegionLat:       "nodes.region_lat",
	RegionLng:       "nodes.region_lng",
	SKUName:         "nodes.sku_name",
	SKUNetworkCap:   "nodes.sku_network_cap",
	SKUNetworkSpeed: "nodes.sku_network_speed",
	SKUPriceMonthly: "nodes.sku_price_monthly",
	SKUPriceHourly:  "nodes.sku_price_hourly",
	WireguardKey:    "nodes.wireguard_key",
	User:            "nodes.user",
	Type:            "nodes.type",
}

// Generated where

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelperfloat64 struct{ field string }

func (w whereHelperfloat64) EQ(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat64) NEQ(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat64) LT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat64) LTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat64) GT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat64) GTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat64) NIN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperfloat32 struct{ field string }

func (w whereHelperfloat32) EQ(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat32) NEQ(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat32) LT(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat32) LTE(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat32) GT(x float32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat32) GTE(x float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat32) IN(slice []float32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat32) NIN(slice []float32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var NodeWhere = struct {
	ID              whereHelperint64
	State           whereHelperstring
	StartedAt       whereHelperint64
	StoppedAt       whereHelpernull_Int64
	ProviderName    whereHelperstring
	ProviderID      whereHelperstring
	Name            whereHelperstring
	Memory          whereHelperint
	CPUs            whereHelperint
	Disk            whereHelperint
	IPV4            whereHelperstring
	IPV6            whereHelperstring
	RegionName      whereHelperstring
	RegionLat       whereHelperfloat64
	RegionLng       whereHelperfloat64
	SKUName         whereHelperstring
	SKUNetworkCap   whereHelperint
	SKUNetworkSpeed whereHelperint
	SKUPriceMonthly whereHelperfloat32
	SKUPriceHourly  whereHelperfloat32
	WireguardKey    whereHelperstring
	User            whereHelperstring
	Type            whereHelperstring
}{
	ID:              whereHelperint64{field: "\"nodes\".\"id\""},
	State:           whereHelperstring{field: "\"nodes\".\"state\""},
	StartedAt:       whereHelperint64{field: "\"nodes\".\"started_at\""},
	StoppedAt:       whereHelpernull_Int64{field: "\"nodes\".\"stopped_at\""},
	ProviderName:    whereHelperstring{field: "\"nodes\".\"provider_name\""},
	ProviderID:      whereHelperstring{field: "\"nodes\".\"provider_id\""},
	Name:            whereHelperstring{field: "\"nodes\".\"name\""},
	Memory:          whereHelperint{field: "\"nodes\".\"memory\""},
	CPUs:            whereHelperint{field: "\"nodes\".\"cpus\""},
	Disk:            whereHelperint{field: "\"nodes\".\"disk\""},
	IPV4:            whereHelperstring{field: "\"nodes\".\"ip_v4\""},
	IPV6:            whereHelperstring{field: "\"nodes\".\"ip_v6\""},
	RegionName:      whereHelperstring{field: "\"nodes\".\"region_name\""},
	RegionLat:       whereHelperfloat64{field: "\"nodes\".\"region_lat\""},
	RegionLng:       whereHelperfloat64{field: "\"nodes\".\"region_lng\""},
	SKUName:         whereHelperstring{field: "\"nodes\".\"sku_name\""},
	SKUNetworkCap:   whereHelperint{field: "\"nodes\".\"sku_network_cap\""},
	SKUNetworkSpeed: whereHelperint{field: "\"nodes\".\"sku_network_speed\""},
	SKUPriceMonthly: whereHelperfloat32{field: "\"nodes\".\"sku_price_monthly\""},
	SKUPriceHourly:  whereHelperfloat32{field: "\"nodes\".\"sku_price_hourly\""},
	WireguardKey:    whereHelperstring{field: "\"nodes\".\"wireguard_key\""},
	User:            whereHelperstring{field: "\"nodes\".\"user\""},
	Type:            whereHelperstring{field: "\"nodes\".\"type\""},
}

// NodeRels is where relationship names are stored.
var NodeRels = struct {
}{}

// nodeR is where relationships are stored.
type nodeR struct {
}

// NewStruct creates a new relationship struct
func (*nodeR) NewStruct() *nodeR {
	return &nodeR{}
}

// nodeL is where Load methods for each relationship are stored.
type nodeL struct{}

var (
	nodeAllColumns            = []string{"id", "state", "started_at", "stopped_at", "provider_name", "provider_id", "name", "memory", "cpus", "disk", "ip_v4", "ip_v6", "region_name", "region_lat", "region_lng", "sku_name", "sku_network_cap", "sku_network_speed", "sku_price_monthly", "sku_price_hourly", "wireguard_key", "user", "type"}
	nodeColumnsWithoutDefault = []string{"started_at", "provider_name", "provider_id", "name", "memory", "cpus", "disk", "ip_v4", "ip_v6", "region_name", "region_lat", "region_lng", "sku_name", "sku_network_cap", "sku_network_speed", "sku_price_monthly", "sku_price_hourly", "wireguard_key", "user", "type"}
	nodeColumnsWithDefault    = []string{"id", "state", "stopped_at"}
	nodePrimaryKeyColumns     = []string{"id"}
	nodeGeneratedColumns      = []string{}
)

type (
	// NodeSlice is an alias for a slice of pointers to Node.
	// This should almost always be used instead of []Node.
	NodeSlice []*Node

	nodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	nodeType                 = reflect.TypeOf(&Node{})
	nodeMapping              = queries.MakeStructMapping(nodeType)
	nodePrimaryKeyMapping, _ = queries.BindMapping(nodeType, nodeMapping, nodePrimaryKeyColumns)
	nodeInsertCacheMut       sync.RWMutex
	nodeInsertCache          = make(map[string]insertCache)
	nodeUpdateCacheMut       sync.RWMutex
	nodeUpdateCache          = make(map[string]updateCache)
	nodeUpsertCacheMut       sync.RWMutex
	nodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single node record from the query using the global executor.
func (q nodeQuery) OneG(ctx context.Context) (*Node, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single node record from the query.
func (q nodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Node, error) {
	o := &Node{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for nodes")
	}

	return o, nil
}

// AllG returns all Node records from the query using the global executor.
func (q nodeQuery) AllG(ctx context.Context) (NodeSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Node records from the query.
func (q nodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (NodeSlice, error) {
	var o []*Node

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Node slice")
	}

	return o, nil
}

// CountG returns the count of all Node records in the query using the global executor
func (q nodeQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Node records in the query.
func (q nodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count nodes rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q nodeQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q nodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if nodes exists")
	}

	return count > 0, nil
}

// Nodes retrieves all the records using an executor.
func Nodes(mods ...qm.QueryMod) nodeQuery {
	mods = append(mods, qm.From("\"nodes\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"nodes\".*"})
	}

	return nodeQuery{q}
}

// FindNodeG retrieves a single record by ID.
func FindNodeG(ctx context.Context, iD int64, selectCols ...string) (*Node, error) {
	return FindNode(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindNode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNode(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Node, error) {
	nodeObj := &Node{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"nodes\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, nodeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from nodes")
	}

	return nodeObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Node) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Node) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no nodes provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(nodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	nodeInsertCacheMut.RLock()
	cache, cached := nodeInsertCache[key]
	nodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			nodeAllColumns,
			nodeColumnsWithDefault,
			nodeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(nodeType, nodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(nodeType, nodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"nodes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"nodes\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into nodes")
	}

	if !cached {
		nodeInsertCacheMut.Lock()
		nodeInsertCache[key] = cache
		nodeInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Node record using the global executor.
// See Update for more documentation.
func (o *Node) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Node.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Node) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	nodeUpdateCacheMut.RLock()
	cache, cached := nodeUpdateCache[key]
	nodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			nodeAllColumns,
			nodePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update nodes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"nodes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, nodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(nodeType, nodeMapping, append(wl, nodePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update nodes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for nodes")
	}

	if !cached {
		nodeUpdateCacheMut.Lock()
		nodeUpdateCache[key] = cache
		nodeUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q nodeQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q nodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for nodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for nodes")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o NodeSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"nodes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, nodePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in node slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all node")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Node) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Node) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no nodes provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(nodeColumnsWithDefault, o)

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

	nodeUpsertCacheMut.RLock()
	cache, cached := nodeUpsertCache[key]
	nodeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			nodeAllColumns,
			nodeColumnsWithDefault,
			nodeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			nodeAllColumns,
			nodePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert nodes, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(nodePrimaryKeyColumns))
			copy(conflict, nodePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"nodes\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(nodeType, nodeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(nodeType, nodeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert nodes")
	}

	if !cached {
		nodeUpsertCacheMut.Lock()
		nodeUpsertCache[key] = cache
		nodeUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single Node record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Node) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Node record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Node) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Node provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), nodePrimaryKeyMapping)
	sql := "DELETE FROM \"nodes\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from nodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for nodes")
	}

	return rowsAff, nil
}

func (q nodeQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q nodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no nodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from nodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for nodes")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o NodeSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"nodes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, nodePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from node slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for nodes")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Node) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Node provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Node) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNode(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NodeSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty NodeSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), nodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"nodes\".* FROM \"nodes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, nodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NodeSlice")
	}

	*o = slice

	return nil
}

// NodeExistsG checks if the Node row exists.
func NodeExistsG(ctx context.Context, iD int64) (bool, error) {
	return NodeExists(ctx, boil.GetContextDB(), iD)
}

// NodeExists checks if the Node row exists.
func NodeExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"nodes\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if nodes exists")
	}

	return exists, nil
}
