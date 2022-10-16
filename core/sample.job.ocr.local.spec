type = "offchainreporting2"
schemaVersion = 1
name = "ocr2 dr"
forwardingAllowed = false
maxTaskDuration = "1s"
contractID = "0xf434fe2b4bb1f130e93c4e77dc1517dc16271ef3"
ocrKeyBundleID = "c690b835f6f3001c9fcee1d148bf041d38e70351775885a7196f5e59d8d58a51"
p2pv2Bootstrappers = [
  "12D3KooWKUZgN346Wsk2ahC34zPWYu6CQNM135dvE95mTcX2KPso@127.0.0.1:6690"
]
relay = "evm"
pluginType = "directrequest"
transmitterID = "0xDA43dABB62fAE4906CCaeA0B9c47BA4D32b887Ac"
observationSource = """
    decode_log   [type="ethabidecodelog" abi="OracleRequest(bytes32 requestId, bytes data)" data="$(jobRun.logData)" topics="$(jobRun.logTopics)"]
    decode_cbor  [type="cborparse" data="$(decode_log.data)"]
    run_computation    [type="bridge" name="ea_bridge" requestData="{\\"id\\": $(jobSpec.externalJobID), \\"data\\": $(decode_cbor)}" timeout="20s"]
    parse_result [type=jsonparse data="$(run_computation)" path="result,result"]
    parse_error  [type=jsonparse data="$(run_computation)" path="result,error"]

    encode_tx    [type="ethabiencode" abi="fulfillRequest(bytes32 requestId, bytes calldata response, bytes calldata err)" data="{\\"requestId\\": $(decode_log.requestId), \\"response\\": $(parse_result), \\"err\\": $(parse_error)}"]
    submit_tx    [type="ethtx" to="0xf434fe2b4bb1f130e93c4e77dc1517dc16271ef3" data="$(encode_tx)"]

decode_log -> decode_cbor -> run_computation -> parse_result -> parse_error -> encode_tx -> submit_tx
"""

[pluginConfig]
minIncomingConfirmations = 2
maxRequestsPerOCRRound = 5
requestE2eTimeoutMillis = 10000

[relayConfig]
chainID = 5