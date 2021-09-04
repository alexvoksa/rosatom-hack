import { HttpClient, HttpParams } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { exhaustMap, map, take } from "rxjs/operators";
import { AuthService } from "src/app/auth/auth.service";

@Injectable({providedIn: 'root'})


export class ManageProfileService{

    constructor(private http: HttpClient, private authService: AuthService){}


    postProfileData(data, userId){
        return this.authService.user.pipe(
            take(1), 
            exhaustMap(user => {
                return this.http.post<any>(`https://mozhaysk-4acc5-default-rtdb.firebaseio.com/user/${userId}.json`, data,
                {
                    params: new HttpParams().set('auth', user.token)
                });
        }))
    }

    getProfileData(userId){
        return this.authService.user.pipe(
            take(1), 
            exhaustMap(user => {
                return this.http.get<any>(`https://mozhaysk-4acc5-default-rtdb.firebaseio.com/user/${userId}.json`, 
                {
                    params: new HttpParams().set('auth', user.token)
                }).pipe(
                    map(responseData => {
                        const array = [];
                        for (const key in responseData){
                            if(responseData.hasOwnProperty(key)){
                                array.push({...responseData[key], id: key});
                            }
                        }
                        return array;
                    })
                );
        }))
    }
    
}