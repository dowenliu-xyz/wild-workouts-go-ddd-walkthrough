/* tslint:disable */
/* eslint-disable */
/**
 * Wild Workouts trainings
 * TODO
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from './configuration';
import type { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
import type { RequestArgs } from './base';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, BaseAPI, RequiredError } from './base';

/**
 * 
 * @export
 * @interface ModelError
 */
export interface ModelError {
    /**
     * 
     * @type {string}
     * @memberof ModelError
     */
    'slug': string;
    /**
     * 
     * @type {string}
     * @memberof ModelError
     */
    'message': string;
}
/**
 * 
 * @export
 * @interface PostTraining
 */
export interface PostTraining {
    /**
     * 
     * @type {string}
     * @memberof PostTraining
     */
    'notes': string;
    /**
     * 
     * @type {string}
     * @memberof PostTraining
     */
    'time': string;
}
/**
 * 
 * @export
 * @interface Training
 */
export interface Training {
    /**
     * 
     * @type {string}
     * @memberof Training
     */
    'uuid': string;
    /**
     * 
     * @type {string}
     * @memberof Training
     */
    'user': string;
    /**
     * 
     * @type {string}
     * @memberof Training
     */
    'userUuid': string;
    /**
     * 
     * @type {string}
     * @memberof Training
     */
    'notes': string;
    /**
     * 
     * @type {string}
     * @memberof Training
     */
    'time': string;
    /**
     * 
     * @type {boolean}
     * @memberof Training
     */
    'canBeCancelled': boolean;
    /**
     * 
     * @type {boolean}
     * @memberof Training
     */
    'moveRequiresAccept': boolean;
    /**
     * 
     * @type {string}
     * @memberof Training
     */
    'proposedTime'?: string;
    /**
     * 
     * @type {string}
     * @memberof Training
     */
    'moveProposedBy'?: string;
}
/**
 * 
 * @export
 * @interface Trainings
 */
export interface Trainings {
    /**
     * 
     * @type {Array<Training>}
     * @memberof Trainings
     */
    'trainings': Array<Training>;
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        approveRescheduleTraining: async (trainingUUID: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'trainingUUID' is not null or undefined
            assertParamExists('approveRescheduleTraining', 'trainingUUID', trainingUUID)
            const localVarPath = `/trainings/{trainingUUID}/approve-reschedule`
                .replace(`{${"trainingUUID"}}`, encodeURIComponent(String(trainingUUID)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication bearerAuth required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        cancelTraining: async (trainingUUID: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'trainingUUID' is not null or undefined
            assertParamExists('cancelTraining', 'trainingUUID', trainingUUID)
            const localVarPath = `/trainings/{trainingUUID}`
                .replace(`{${"trainingUUID"}}`, encodeURIComponent(String(trainingUUID)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'DELETE', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication bearerAuth required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {PostTraining} postTraining todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createTraining: async (postTraining: PostTraining, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'postTraining' is not null or undefined
            assertParamExists('createTraining', 'postTraining', postTraining)
            const localVarPath = `/trainings`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication bearerAuth required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(postTraining, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getTrainings: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/trainings`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication bearerAuth required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        rejectRescheduleTraining: async (trainingUUID: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'trainingUUID' is not null or undefined
            assertParamExists('rejectRescheduleTraining', 'trainingUUID', trainingUUID)
            const localVarPath = `/trainings/{trainingUUID}/reject-reschedule`
                .replace(`{${"trainingUUID"}}`, encodeURIComponent(String(trainingUUID)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication bearerAuth required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {PostTraining} postTraining todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        requestRescheduleTraining: async (trainingUUID: string, postTraining: PostTraining, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'trainingUUID' is not null or undefined
            assertParamExists('requestRescheduleTraining', 'trainingUUID', trainingUUID)
            // verify required parameter 'postTraining' is not null or undefined
            assertParamExists('requestRescheduleTraining', 'postTraining', postTraining)
            const localVarPath = `/trainings/{trainingUUID}/request-reschedule`
                .replace(`{${"trainingUUID"}}`, encodeURIComponent(String(trainingUUID)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication bearerAuth required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(postTraining, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {PostTraining} postTraining todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        rescheduleTraining: async (trainingUUID: string, postTraining: PostTraining, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'trainingUUID' is not null or undefined
            assertParamExists('rescheduleTraining', 'trainingUUID', trainingUUID)
            // verify required parameter 'postTraining' is not null or undefined
            assertParamExists('rescheduleTraining', 'postTraining', postTraining)
            const localVarPath = `/trainings/{trainingUUID}/reschedule`
                .replace(`{${"trainingUUID"}}`, encodeURIComponent(String(trainingUUID)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication bearerAuth required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(postTraining, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * DefaultApi - functional programming interface
 * @export
 */
export const DefaultApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = DefaultApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async approveRescheduleTraining(trainingUUID: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.approveRescheduleTraining(trainingUUID, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async cancelTraining(trainingUUID: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.cancelTraining(trainingUUID, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {PostTraining} postTraining todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createTraining(postTraining: PostTraining, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.createTraining(postTraining, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getTrainings(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Trainings>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getTrainings(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async rejectRescheduleTraining(trainingUUID: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.rejectRescheduleTraining(trainingUUID, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {PostTraining} postTraining todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async requestRescheduleTraining(trainingUUID: string, postTraining: PostTraining, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.requestRescheduleTraining(trainingUUID, postTraining, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {PostTraining} postTraining todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async rescheduleTraining(trainingUUID: string, postTraining: PostTraining, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.rescheduleTraining(trainingUUID, postTraining, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * DefaultApi - factory interface
 * @export
 */
export const DefaultApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = DefaultApiFp(configuration)
    return {
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        approveRescheduleTraining(trainingUUID: string, options?: any): AxiosPromise<void> {
            return localVarFp.approveRescheduleTraining(trainingUUID, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        cancelTraining(trainingUUID: string, options?: any): AxiosPromise<void> {
            return localVarFp.cancelTraining(trainingUUID, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {PostTraining} postTraining todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createTraining(postTraining: PostTraining, options?: any): AxiosPromise<void> {
            return localVarFp.createTraining(postTraining, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getTrainings(options?: any): AxiosPromise<Trainings> {
            return localVarFp.getTrainings(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        rejectRescheduleTraining(trainingUUID: string, options?: any): AxiosPromise<void> {
            return localVarFp.rejectRescheduleTraining(trainingUUID, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {PostTraining} postTraining todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        requestRescheduleTraining(trainingUUID: string, postTraining: PostTraining, options?: any): AxiosPromise<void> {
            return localVarFp.requestRescheduleTraining(trainingUUID, postTraining, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} trainingUUID todo
         * @param {PostTraining} postTraining todo
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        rescheduleTraining(trainingUUID: string, postTraining: PostTraining, options?: any): AxiosPromise<void> {
            return localVarFp.rescheduleTraining(trainingUUID, postTraining, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * DefaultApi - object-oriented interface
 * @export
 * @class DefaultApi
 * @extends {BaseAPI}
 */
export class DefaultApi extends BaseAPI {
    /**
     * 
     * @param {string} trainingUUID todo
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public approveRescheduleTraining(trainingUUID: string, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).approveRescheduleTraining(trainingUUID, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} trainingUUID todo
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public cancelTraining(trainingUUID: string, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).cancelTraining(trainingUUID, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {PostTraining} postTraining todo
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public createTraining(postTraining: PostTraining, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).createTraining(postTraining, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getTrainings(options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getTrainings(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} trainingUUID todo
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public rejectRescheduleTraining(trainingUUID: string, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).rejectRescheduleTraining(trainingUUID, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} trainingUUID todo
     * @param {PostTraining} postTraining todo
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public requestRescheduleTraining(trainingUUID: string, postTraining: PostTraining, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).requestRescheduleTraining(trainingUUID, postTraining, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} trainingUUID todo
     * @param {PostTraining} postTraining todo
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public rescheduleTraining(trainingUUID: string, postTraining: PostTraining, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).rescheduleTraining(trainingUUID, postTraining, options).then((request) => request(this.axios, this.basePath));
    }
}

