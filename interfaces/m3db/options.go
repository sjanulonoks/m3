	"os"
	"github.com/m3db/m3x/log"
	"github.com/m3db/m3x/metrics"

	tchannel "github.com/uber/tchannel-go"
	EncodingTszPooled(bufferBucketAllocSize, databaseBlockAllocSize int) DatabaseOptions

	// EncodingTsz sets tsz encoding and returns a new DatabaseOptions
	EncodingTsz() DatabaseOptions
	Logger(value xlog.Logger) DatabaseOptions
	GetLogger() xlog.Logger
	MetricsScope(value xmetrics.Scope) DatabaseOptions
	GetMetricsScope() xmetrics.Scope
	// DatabaseBlockAllocSize sets the databaseBlockAllocSize and returns a new DatabaseOptions
	DatabaseBlockAllocSize(value int) DatabaseOptions

	// GetDatabaseBlockAllocSize returns the databaseBlockAllocSize
	GetDatabaseBlockAllocSize() int

	// CommitLogFlushSize sets the commit log flush size and returns a new DatabaseOptions
	CommitLogFlushSize(value int) DatabaseOptions

	// GetCommitLogFlushSize returns the commit log flush size
	GetCommitLogFlushSize() int

	// CommitLogFlushInterval sets the commit log flush interval and returns a new DatabaseOptions
	CommitLogFlushInterval(value time.Duration) DatabaseOptions

	// GetCommitLogFlushInterval returns the commit log flush interval
	GetCommitLogFlushInterval() time.Duration

	// CommitLogBacklogQueueSize sets the commit log backlog queue size and returns a new DatabaseOptions
	CommitLogBacklogQueueSize(value int) DatabaseOptions

	// GetCommitLogBacklogQueueSize returns the commit log backlog queue size
	GetCommitLogBacklogQueueSize() int

	// CommitLogStrategy sets the commit log strategy and returns a new DatabaseOptions
	CommitLogStrategy(value CommitLogStrategy) DatabaseOptions

	// GetCommitLogStrategy returns the commit log strategy
	GetCommitLogStrategy() CommitLogStrategy

	// DatabaseBlockPool sets the databaseBlockPool and returns a new DatabaseOptions
	DatabaseBlockPool(value DatabaseBlockPool) DatabaseOptions

	// GetDatabaseBlockPool returns the databaseBlockPool
	GetDatabaseBlockPool() DatabaseBlockPool

	// SegmentReaderPool sets the segment reader pool.
	SegmentReaderPool(value SegmentReaderPool) DatabaseOptions

	// GetSegmentReaderPool returns the segment reader pool.
	GetSegmentReaderPool() SegmentReaderPool

	// ReaderIteratorPool sets the readerIteratorPool and returns a new DatabaseOptions
	ReaderIteratorPool(value ReaderIteratorPool) DatabaseOptions

	// GetReaderIteratorPool returns the readerIteratorPool
	GetReaderIteratorPool() ReaderIteratorPool
	// MultiReaderIteratorPool sets the multiReaderIteratorPool and returns a new DatabaseOptions
	MultiReaderIteratorPool(value MultiReaderIteratorPool) DatabaseOptions

	// GetMultiReaderIteratorPool returns the multiReaderIteratorPool
	GetMultiReaderIteratorPool() MultiReaderIteratorPool
	// FileWriterOptions sets the file writer options.
	FileWriterOptions(value FileWriterOptions) DatabaseOptions

	// GetFileWriterOptions returns the file writer options.
	GetFileWriterOptions() FileWriterOptions

	// NewFileSetReaderFn sets the function for creating a new fileset reader.
	NewFileSetReaderFn(value NewFileSetReaderFn) DatabaseOptions

	// GetNewFileSetReaderFn returns the function for creating a new fileset reader.
	GetNewFileSetReaderFn() NewFileSetReaderFn


	// NewPersistenceManagerFn sets the function for creating a new persistence manager.
	NewPersistenceManagerFn(value NewPersistenceManagerFn) DatabaseOptions

	// GetNewPersistenceManagerFn returns the function for creating a new persistence manager.
	GetNewPersistenceManagerFn() NewPersistenceManagerFn

	// WriterBufferSize sets the buffer size for writing TSDB files.
	WriterBufferSize(value int) DatabaseOptions

	// GetWriterBufferSize returns the buffer size for writing TSDB files.
	GetWriterBufferSize() int

	// ReaderBufferSize sets the buffer size for reading TSDB files.
	ReaderBufferSize(value int) DatabaseOptions

	// GetReaderBufferSize returns the buffer size for reading TSDB files.
	GetReaderBufferSize() int
}

// ClientOptions is a set of client options
type ClientOptions interface {
	// Validate validates the options
	Validate() error

	// EncodingTsz sets tsz encoding and returns a new ClientOptions
	EncodingTsz() ClientOptions

	// Logger sets the logger and returns a new ClientOptions
	Logger(value xlog.Logger) ClientOptions

	// GetLogger returns the logger
	GetLogger() xlog.Logger

	// MetricsScope sets the metricsScope and returns a new ClientOptions
	MetricsScope(value xmetrics.Scope) ClientOptions

	// GetMetricsScope returns the metricsScope
	GetMetricsScope() xmetrics.Scope

	// TopologyType sets the topologyType and returns a new ClientOptions
	TopologyType(value TopologyType) ClientOptions

	// GetTopologyType returns the topologyType
	GetTopologyType() TopologyType

	// ConsistencyLevel sets the consistencyLevel and returns a new ClientOptions
	ConsistencyLevel(value ConsistencyLevel) ClientOptions

	// GetConsistencyLevel returns the consistencyLevel
	GetConsistencyLevel() ConsistencyLevel

	// ChannelOptions sets the channelOptions and returns a new ClientOptions
	ChannelOptions(value *tchannel.ChannelOptions) ClientOptions

	// GetChannelOptions returns the channelOptions
	GetChannelOptions() *tchannel.ChannelOptions

	// NowFn sets the nowFn and returns a new ClientOptions
	NowFn(value NowFn) ClientOptions

	// GetNowFn returns the nowFn
	GetNowFn() NowFn

	// MaxConnectionCount sets the maxConnectionCount and returns a new ClientOptions
	MaxConnectionCount(value int) ClientOptions

	// GetMaxConnectionCount returns the maxConnectionCount
	GetMaxConnectionCount() int

	// MinConnectionCount sets the minConnectionCount and returns a new ClientOptions
	MinConnectionCount(value int) ClientOptions

	// GetMinConnectionCount returns the minConnectionCount
	GetMinConnectionCount() int

	// HostConnectTimeout sets the hostConnectTimeout and returns a new ClientOptions
	HostConnectTimeout(value time.Duration) ClientOptions

	// GetHostConnectTimeout returns the hostConnectTimeout
	GetHostConnectTimeout() time.Duration

	// ClusterConnectTimeout sets the clusterConnectTimeout and returns a new ClientOptions
	ClusterConnectTimeout(value time.Duration) ClientOptions

	// GetClusterConnectTimeout returns the clusterConnectTimeout
	GetClusterConnectTimeout() time.Duration

	// ClusterConnectConsistencyLevel sets the clusterConnectConsistencyLevel and returns a new ClientOptions
	ClusterConnectConsistencyLevel(value ConsistencyLevel) ClientOptions

	// GetClusterConnectConsistencyLevel returns the clusterConnectConsistencyLevel
	GetClusterConnectConsistencyLevel() ConsistencyLevel

	// WriteRequestTimeout sets the writeRequestTimeout and returns a new ClientOptions
	WriteRequestTimeout(value time.Duration) ClientOptions

	// GetWriteRequestTimeout returns the writeRequestTimeout
	GetWriteRequestTimeout() time.Duration

	// FetchRequestTimeout sets the fetchRequestTimeout and returns a new ClientOptions
	FetchRequestTimeout(value time.Duration) ClientOptions

	// GetFetchRequestTimeout returns the fetchRequestTimeout
	GetFetchRequestTimeout() time.Duration

	// BackgroundConnectInterval sets the backgroundConnectInterval and returns a new ClientOptions
	BackgroundConnectInterval(value time.Duration) ClientOptions

	// GetBackgroundConnectInterval returns the backgroundConnectInterval
	GetBackgroundConnectInterval() time.Duration

	// BackgroundConnectStutter sets the backgroundConnectStutter and returns a new ClientOptions
	BackgroundConnectStutter(value time.Duration) ClientOptions

	// GetBackgroundConnectStutter returns the backgroundConnectStutter
	GetBackgroundConnectStutter() time.Duration

	// BackgroundHealthCheckInterval sets the backgroundHealthCheckInterval and returns a new ClientOptions
	BackgroundHealthCheckInterval(value time.Duration) ClientOptions

	// GetBackgroundHealthCheckInterval returns the backgroundHealthCheckInterval
	GetBackgroundHealthCheckInterval() time.Duration

	// BackgroundHealthCheckStutter sets the backgroundHealthCheckStutter and returns a new ClientOptions
	BackgroundHealthCheckStutter(value time.Duration) ClientOptions

	// GetBackgroundHealthCheckStutter returns the backgroundHealthCheckStutter
	GetBackgroundHealthCheckStutter() time.Duration

	// WriteOpPoolSize sets the writeOpPoolSize and returns a new ClientOptions
	WriteOpPoolSize(value int) ClientOptions

	// GetWriteOpPoolSize returns the writeOpPoolSize
	GetWriteOpPoolSize() int

	// FetchBatchOpPoolSize sets the fetchBatchOpPoolSize and returns a new ClientOptions
	FetchBatchOpPoolSize(value int) ClientOptions

	// GetFetchBatchOpPoolSize returns the fetchBatchOpPoolSize
	GetFetchBatchOpPoolSize() int

	// WriteBatchSize sets the writeBatchSize and returns a new ClientOptions
	// NB(r): for a write only application load this should match the host
	// queue ops flush size so that each time a host queue is flushed it can
	// fit the entire flushed write ops into a single batch.
	WriteBatchSize(value int) ClientOptions

	// GetWriteBatchSize returns the writeBatchSize
	GetWriteBatchSize() int

	// FetchBatchSize sets the fetchBatchSize and returns a new ClientOptions
	// NB(r): for a fetch only application load this should match the host
	// queue ops flush size so that each time a host queue is flushed it can
	// fit the entire flushed fetch ops into a single batch.
	FetchBatchSize(value int) ClientOptions

	// GetFetchBatchSize returns the fetchBatchSize
	GetFetchBatchSize() int

	// HostQueueOpsFlushSize sets the hostQueueOpsFlushSize and returns a new ClientOptions
	HostQueueOpsFlushSize(value int) ClientOptions

	// GetHostQueueOpsFlushSize returns the hostQueueOpsFlushSize
	GetHostQueueOpsFlushSize() int

	// HostQueueOpsFlushInterval sets the hostQueueOpsFlushInterval and returns a new ClientOptions
	HostQueueOpsFlushInterval(value time.Duration) ClientOptions

	// GetHostQueueOpsFlushInterval returns the hostQueueOpsFlushInterval
	GetHostQueueOpsFlushInterval() time.Duration

	// HostQueueOpsArrayPoolSize sets the hostQueueOpsArrayPoolSize and returns a new ClientOptions
	HostQueueOpsArrayPoolSize(value int) ClientOptions

	// GetHostQueueOpsArrayPoolSize returns the hostQueueOpsArrayPoolSize
	GetHostQueueOpsArrayPoolSize() int

	// SeriesIteratorPoolSize sets the seriesIteratorPoolSize and returns a new ClientOptions
	SeriesIteratorPoolSize(value int) ClientOptions

	// GetSeriesIteratorPoolSize returns the seriesIteratorPoolSize
	GetSeriesIteratorPoolSize() int

	// SeriesIteratorArrayPoolBuckets sets the seriesIteratorArrayPoolBuckets and returns a new ClientOptions
	SeriesIteratorArrayPoolBuckets(value []PoolBucket) ClientOptions

	// GetSeriesIteratorArrayPoolBuckets returns the seriesIteratorArrayPoolBuckets
	GetSeriesIteratorArrayPoolBuckets() []PoolBucket

	// ReaderIteratorAllocate sets the readerIteratorAllocate and returns a new ClientOptions
	ReaderIteratorAllocate(value ReaderIteratorAllocate) ClientOptions

	// GetReaderIteratorAllocate returns the readerIteratorAllocate
	GetReaderIteratorAllocate() ReaderIteratorAllocate
}

// TopologyTypeOptions is a set of static topology type options
type TopologyTypeOptions interface {
	// Validate validates the options
	Validate() error

	// ShardScheme sets the shardScheme and returns a new TopologyTypeOptions
	ShardScheme(value ShardScheme) TopologyTypeOptions

	// GetShardScheme returns the shardScheme
	GetShardScheme() ShardScheme

	// Replicas sets the replicas and returns a new TopologyTypeOptions
	Replicas(value int) TopologyTypeOptions

	// GetReplicas returns the replicas
	GetReplicas() int

	// HostShardSets sets the hostShardSets and returns a new TopologyTypeOptions
	HostShardSets(value []HostShardSet) TopologyTypeOptions

	// GetHostShardSets returns the hostShardSets
	GetHostShardSets() []HostShardSet
}

// FileWriterOptions is a set of file writing options for a file writer
type FileWriterOptions interface {
	// NewFileMode sets the new file mode.
	NewFileMode(value os.FileMode) FileWriterOptions

	// GetNewFileMode returns the new file mode.
	GetNewFileMode() os.FileMode

	// NewDirectoryMode sets the new directory mode.
	NewDirectoryMode(value os.FileMode) FileWriterOptions

	// GetNewDirectoryMode returns the new directory mode.
	GetNewDirectoryMode() os.FileMode