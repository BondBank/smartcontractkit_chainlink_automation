package directrequestocr_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/core/logger"
	drocr_serv "github.com/smartcontractkit/chainlink/core/services/directrequestocr"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/directrequestocr"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func PreparePlugin(t *testing.T) (types.ReportingPlugin, drocr_serv.ORM) {
	ocrLogger := logger.NewOCRWrapper(logger.TestLogger(t), true, func(msg string) {})

	orm := drocr_serv.NewInMemoryORM()
	factory := directrequestocr.DirectRequestReportingPluginFactory{
		Logger:    ocrLogger,
		PluginORM: orm,
	}

	plugin, _, _ := factory.NewReportingPlugin(types.ReportingPluginConfig{})
	return plugin, orm
}

func intToByte32(id uint) [32]byte {
	byteArr := (*[32]byte)([]byte(fmt.Sprintf("%032d\n", id)))
	return *byteArr
}

func TestDRReporting_Query(t *testing.T) {
	t.Parallel()
	plugin, orm := PreparePlugin(t)
	reqId := intToByte32(13)
	txHash := common.HexToHash("0xabc")

	id, err := orm.CreateRequest(reqId, time.Now(), &txHash)
	require.NoError(t, err)
	orm.SetResult(id, 1, []byte{}, time.Now())
	q, err := plugin.Query(testutils.Context(t), types.ReportTimestamp{})
	require.NoError(t, err)

	queryProto := &directrequestocr.Query{}
	err = proto.Unmarshal(q, queryProto)
	require.NoError(t, err)

	require.Equal(t, len(queryProto.RequestIDs), 1)
	require.Equal(t, queryProto.RequestIDs[0], reqId[:])
}
