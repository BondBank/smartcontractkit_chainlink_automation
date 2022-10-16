package directrequestocr

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/smartcontractkit/chainlink/core/services/directrequestocr"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"google.golang.org/protobuf/proto"
)

type DirectRequestReportingPluginFactory struct {
	Logger    commontypes.Logger
	PluginORM directrequestocr.ORM
	JobID     uuid.UUID
}

var _ types.ReportingPluginFactory = (*DirectRequestReportingPluginFactory)(nil)

type directRequestReporting struct {
	logger    commontypes.Logger
	pluginORM directrequestocr.ORM
	jobID     uuid.UUID

	/*offchainConfig           OffchainConfig
	onchainConfig            OnchainConfig
	contractTransmitter      MedianContract
	dataSource               DataSource
	logger                   loghelper.LoggerWithContext
	reportCodec              ReportCodec
	configDigest             types.ConfigDigest
	latestAcceptedEpochRound epochRound*/
}

var _ types.ReportingPlugin = (*directRequestReporting)(nil)

// NewReportingPlugin complies with ReportingPluginFactory
func (f DirectRequestReportingPluginFactory) NewReportingPlugin(types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	info := types.ReportingPluginInfo{
		Name:          "directRequestReporting",
		UniqueReports: false, // todo change?
		Limits: types.ReportingPluginLimits{
			MaxQueryLength:       1000,
			MaxObservationLength: 1000,
			MaxReportLength:      1000,
		},
	}
	plugin := directRequestReporting{
		logger:    f.Logger,
		pluginORM: f.PluginORM,
		jobID:     f.JobID,
	}
	return plugin, info, nil
}

// Query() complies with ReportingPlugin
func (r directRequestReporting) Query(ctx context.Context, ts types.ReportTimestamp) (types.Query, error) {
	str := fmt.Sprintf("DRPLUGIN QUERY(ep: %d, rnd: %d) jobID: %s", ts.Epoch, ts.Round, r.jobID.String())
	r.logger.Info(str, commontypes.LogFields{})
	results, err := r.pluginORM.FindOldestEntriesByState(directrequestocr.RESULT_READY, 1000) // todo put limit in config
	if err != nil {
		return nil, nil
	}

	queryProto := Query{}
	for _, value := range results {
		queryProto.RequestIDs = append(queryProto.RequestIDs, value.ContractRequestID[:])
	}
	marshalled, err := proto.Marshal(&queryProto)
	return marshalled, err
}

// Observation() complies with ReportingPlugin
func (r directRequestReporting) Observation(ctx context.Context, ts types.ReportTimestamp, query types.Query) (types.Observation, error) {
	str := fmt.Sprintf("DRPLUGIN OBSERVATION(ep: %d, rnd: %d) received query: %s, myID: %s", ts.Epoch, ts.Round, string(query), r.jobID.String())
	r.logger.Info(str, commontypes.LogFields{})
	//iterate over report, fetch every item and check if it's ready
	return []byte("abc"), nil
}

// Report() complies with ReportingPlugin
func (r directRequestReporting) Report(ctx context.Context, ts types.ReportTimestamp, query types.Query, obs []types.AttributedObservation) (bool, types.Report, error) {
	str := fmt.Sprintf("DRPLUGIN REPORT(ep: %d, rnd: %d) , myID: %s, obCount: %d  |    ", ts.Epoch, ts.Round, r.jobID.String(), len(obs))
	/*for _, atObs := range obs {
		str = str + "(id: " + string(atObs.Observer) + " , obs: " + string(atObs.Observation) + "),  "
	}*/
	r.logger.Info(str, commontypes.LogFields{})
	return true, []byte("abc"), nil
}

// ShouldAcceptFinalizedReport() complies with ReportingPlugin
func (r directRequestReporting) ShouldAcceptFinalizedReport(ctx context.Context, ts types.ReportTimestamp, report types.Report) (bool, error) {
	str := fmt.Sprintf("DRPLUGIN SHOULD_ACCEPT(ep: %d, rnd: %d) , myID: %s, report: %s", ts.Epoch, ts.Round, r.jobID.String(), string(report))
	r.logger.Info(str, commontypes.LogFields{})
	return true, nil
}

// ShouldTransmitAcceptedReport() complies with ReportingPlugin
func (r directRequestReporting) ShouldTransmitAcceptedReport(ctx context.Context, ts types.ReportTimestamp, report types.Report) (bool, error) {
	str := fmt.Sprintf("DRPLUGIN SHOULD_TRASMIT(ep: %d, rnd: %d) , myID: %s, report: %s", ts.Epoch, ts.Round, r.jobID.String(), string(report))
	r.logger.Info(str, commontypes.LogFields{})
	return true, nil
}

// Close() complies with ReportingPlugin
func (r directRequestReporting) Close() error {
	r.logger.Error("DRPLUGIN  CLOSE", commontypes.LogFields{})
	return nil
}
