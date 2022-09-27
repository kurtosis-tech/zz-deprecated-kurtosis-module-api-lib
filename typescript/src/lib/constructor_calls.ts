import { ExecuteResponse } from "kurtosis-sdk"

export function newExecuteResponse(responseJson: string): ExecuteResponse {
    const result: ExecuteResponse = new ExecuteResponse();
    result.setResponseJson(responseJson);
    return result;
}
