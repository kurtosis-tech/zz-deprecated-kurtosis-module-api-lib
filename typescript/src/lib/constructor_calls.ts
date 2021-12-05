import { ExecuteResponse } from "kurtosis-core-api-lib"

export function newExecuteResponse(responseJson: string): ExecuteResponse {
    const result: ExecuteResponse = new ExecuteResponse();
    result.setResponseJson(responseJson);
    return result;
}