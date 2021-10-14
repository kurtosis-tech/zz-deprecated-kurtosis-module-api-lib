import { ExecuteResponse } from "../kurtosis_module_rpc_api_bindings/executable_module_service_pb";

export function newExecuteResponse(responseJson: string): ExecuteResponse {
    const result: ExecuteResponse = new ExecuteResponse();
    result.setResponseJson(responseJson);
    return result;
}