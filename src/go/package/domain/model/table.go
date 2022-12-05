package model

type (
	TableKey struct {
		KeySpace string
		Name     string
	}
	Table struct {
		TableKey
		Id                      string
		Comment                 string
		SpeculativeRetry        string
		DefaultTTL              int
		Gc                      int
		MaxIndexInterval        int
		MinIndexInterval        int
		FlushPeriod             int
		CrcCheckChance          float64
		ReadRepairChance        float64
		DclocalReadRepairChance float64
		BloomFilterFpChance     float64
		Caching                 map[string]string
		Compression             map[string]string
		Compaction              map[string]string
		Flags                   []string
		Extensions              map[string][]byte
		Columns                 []Column
	}
)

func (k TableKey) String() string { return k.KeySpace + "." + k.Name }
func (k TableKey) Raw() any       { return k }
