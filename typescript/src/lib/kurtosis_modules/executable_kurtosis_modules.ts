import { EnclaveContext } from "kurtosis-sdk";
import { Result } from "neverthrow";

// Docs available at https://docs.kurtosistech.com/kurtosis-module-api-lib/lib-documentation
export interface ExecutableKurtosisModule {
	execute(
        enclaveCtx: EnclaveContext,
        serializedParams: string
    ): Promise<Result<string, Error>>
}
