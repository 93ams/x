package model

type (
	KeySpace struct {
		Name        string            `db:"keyspace_name"`
		Durable     bool              `db:"durable_writes"`
		Replication map[string]string `db:"replication"`
	}
	Table struct {
		Id                      string            `db:"id"`
		Name                    string            `db:"table_name"`
		Keyspace                string            `db:"keyspace_name"`
		Comment                 string            `db:"comment"`
		SpeculativeRetry        string            `db:"speculative_retry"`
		DefaultTTL              int               `db:"default_time_to_live"`
		Gc                      int               `db:"gc_grace_seconds"`
		MaxIndexInterval        int               `db:"max_index_interval"`
		MinIndexInterval        int               `db:"min_index_interval"`
		FlushPeriod             int               `db:"memtable_flush_period_in_ms"`
		CrcCheckChance          float64           `db:"crc_check_chance"`
		ReadRepairChance        float64           `db:"read_repair_chance"`
		DclocalReadRepairChance float64           `db:"dclocal_read_repair_chance"`
		BloomFilterFpChance     float64           `db:"bloom_filter_fp_chance"`
		Caching                 map[string]string `db:"caching"`
		Compression             map[string]string `db:"compression"`
		Compaction              map[string]string `db:"compaction"`
		Flags                   []string          `db:"flags"`
		Extensions              map[string][]byte `db:"extensions"`
	}
)
