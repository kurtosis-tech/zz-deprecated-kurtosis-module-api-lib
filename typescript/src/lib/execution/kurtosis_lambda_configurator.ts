import { Result } from "neverthrow";
import { KurtosisLambda } from "../kurtosis_lambda/kurtosis_lambda";

// Docs available at https://docs.kurtosistech.com/kurtosis-lambda-api-lib/lib-documentation
export interface KurtosisLambdaConfigurator {
    parseParamsAndCreateKurtosisLambda(serializedCustomParamsStr: string): Result<KurtosisLambda, Error>;
}
