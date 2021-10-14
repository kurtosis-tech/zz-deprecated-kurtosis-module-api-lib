import { NetworkContext } from "kurtosis-core-api-lib";
import { Result } from "neverthrow";

// Docs available at https://docs.kurtosistech.com/kurtosis-module-api-lib/lib-documentation
export interface ExecutableKurtosisModule {
	execute(
        networkCtx: NetworkContext, 
        serializedParams: string
    ): Promise<Result<string, Error>>
}