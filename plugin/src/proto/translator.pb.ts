/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "./fetch.pb"
export type TranslatorRequest = {
  num?: string
}

export type TranslatorResponse = {
  ans?: number
}

export type HealthCheckResponse = {
  health?: string
}

export type HealthCheckRequest = {
}

export class TranslatorService {
  static HealthCheck(req: HealthCheckRequest, initReq?: fm.InitReq): Promise<HealthCheckResponse> {
    return fm.fetchReq<HealthCheckRequest, HealthCheckResponse>(`/translator.TranslatorService/HealthCheck`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}