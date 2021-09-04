import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import axios from "axios";
import { from, Observable } from 'rxjs';
import { resolve } from 'dns';

const hasura = axios.create({
  baseURL: PRODUCTION
    ? "https://hasura.lumiprobe.com/v1/graphql"
    : "https://hasura.lumiprobe.nmr.su/v1/graphql",
  withCredentials: true,
  headers: { "x-hasura-admin-secret": "Q8vr92A-r#n_3b" },
});


@Injectable({
  providedIn: 'root'
})
export class TenderProvidersService {

  constructor() { }

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  applyFilter(reqObj: Object) {
    return from(
      hasura.post("/", {
        query: require("../gql/qSearchTenderProviders.js").default
      })
    )
  }

}
