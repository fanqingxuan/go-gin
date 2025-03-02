package g

type (
	M          = map[string]any
	Map        = map[string]any    // Map is alias of frequently-used map type map[string]any.
	MapAnyAny  = map[any]any       // MapAnyAny is alias of frequently-used map type map[any]any.
	MapAnyStr  = map[any]string    // MapAnyStr is alias of frequently-used map type map[any]string.
	MapAnyInt  = map[any]int       // MapAnyInt is alias of frequently-used map type map[any]int.
	MapStrAny  = map[string]any    // MapStrAny is alias of frequently-used map type map[string]any.
	MapStrStr  = map[string]string // MapStrStr is alias of frequently-used map type map[string]string.
	MapStrInt  = map[string]int    // MapStrInt is alias of frequently-used map type map[string]int.
	MapIntAny  = map[int]any       // MapIntAny is alias of frequently-used map type map[int]any.
	MapIntStr  = map[int]string    // MapIntStr is alias of frequently-used map type map[int]string.
	MapIntInt  = map[int]int       // MapIntInt is alias of frequently-used map type map[int]int.
	MapAnyBool = map[any]bool      // MapAnyBool is alias of frequently-used map type map[any]bool.
	MapStrBool = map[string]bool   // MapStrBool is alias of frequently-used map type map[string]bool.
	MapIntBool = map[int]bool      // MapIntBool is alias of frequently-used map type map[int]bool.
)

type (
	List        = []Map        // List is alias of frequently-used slice type []Map.
	ListAnyAny  = []MapAnyAny  // ListAnyAny is alias of frequently-used slice type []MapAnyAny.
	ListAnyStr  = []MapAnyStr  // ListAnyStr is alias of frequently-used slice type []MapAnyStr.
	ListAnyInt  = []MapAnyInt  // ListAnyInt is alias of frequently-used slice type []MapAnyInt.
	ListStrAny  = []MapStrAny  // ListStrAny is alias of frequently-used slice type []MapStrAny.
	ListStrStr  = []MapStrStr  // ListStrStr is alias of frequently-used slice type []MapStrStr.
	ListStrInt  = []MapStrInt  // ListStrInt is alias of frequently-used slice type []MapStrInt.
	ListIntAny  = []MapIntAny  // ListIntAny is alias of frequently-used slice type []MapIntAny.
	ListIntStr  = []MapIntStr  // ListIntStr is alias of frequently-used slice type []MapIntStr.
	ListIntInt  = []MapIntInt  // ListIntInt is alias of frequently-used slice type []MapIntInt.
	ListAnyBool = []MapAnyBool // ListAnyBool is alias of frequently-used slice type []MapAnyBool.
	ListStrBool = []MapStrBool // ListStrBool is alias of frequently-used slice type []MapStrBool.
	ListIntBool = []MapIntBool // ListIntBool is alias of frequently-used slice type []MapIntBool.
)

type (
	Slice    = []any    // Slice is alias of frequently-used slice type []any.
	SliceAny = []any    // SliceAny is alias of frequently-used slice type []any.
	SliceStr = []string // SliceStr is alias of frequently-used slice type []string.
	SliceInt = []int    // SliceInt is alias of frequently-used slice type []int.
)

type (
	Array    = []any    // Array is alias of frequently-used slice type []any.
	ArrayAny = []any    // ArrayAny is alias of frequently-used slice type []any.
	ArrayStr = []string // ArrayStr is alias of frequently-used slice type []string.
	ArrayInt = []int    // ArrayInt is alias of frequently-used slice type []int.
)

// Scalar 接口用于限制参数为标量类型或基于标量的自定义类型
type Scalar interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~bool | ~string // 允许基于标量类型的自定义类型
}
