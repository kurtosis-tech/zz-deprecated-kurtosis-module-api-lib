import { Result } from "neverthrow";
import { ExecutableKurtosisModule } from "../kurtosis_modules/executable_kurtosis_modules";

// Docs available at https://docs.kurtosistech.com/kurtosis-module-api-lib/lib-documentation
export interface KurtosisModuleConfigurator {
    parseParamsAndCreateExecutableModule(serializedCustomParamsStr: string): Result<ExecutableKurtosisModule, Error>;
}
