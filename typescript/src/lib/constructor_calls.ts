import { ExecuteResponse } from "../kurtosis_lambda_rpc_api_bindings/api_lambda_service_pb";

export function newExecuteResponse(responseJson: string): ExecuteResponse {
    const result: ExecuteResponse = new ExecuteResponse();
    result.setResponseJson(responseJson);
    return result;
}