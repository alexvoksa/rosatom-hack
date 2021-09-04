import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import axios from "axios";
import { from, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TenderProvidersService {

  hasura: null;

  constructor() {
    this.hasura = axios.create({
      baseURL: "http://195.133.48.245:8081",
      withCredentials: true,
      headers: { "x-hasura-admin-secret": "qwerty" },
    });
  }

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  applyFilter(reqObj: Object) {
    if (this.hasura !== null) {
      return from(
        this.hasura.post("/", {
          query: require("../gql/qSearchTenderProviders.js").default
        })
      )
    }
  }

}
