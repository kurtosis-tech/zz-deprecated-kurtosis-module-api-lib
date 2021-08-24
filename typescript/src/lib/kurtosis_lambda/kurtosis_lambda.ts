import { NetworkContext } from "kurtosis-core-api-lib";
import { Result } from "neverthrow";

// Docs available at https://docs.kurtosistech.com/kurtosis-lambda-api-lib/lib-documentation
export interface KurtosisLambda {
	execute(
        networkCtx: NetworkContext, 
        serializedParams: string
    ): Promise<Result<string, Error>>
}