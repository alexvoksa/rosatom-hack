import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { BehaviorSubject, Subject } from "rxjs";
import { User } from "./user.modal";
import { tap } from "rxjs/operators"
import { Router } from "@angular/router";

interface AuthResponseData{
    idToken: string,
    email: string,
    refreshTocken: string,
    expiresIn: string,
    localId: string
}

@Injectable({providedIn: 'root'})

export class AuthService{
    user = new BehaviorSubject<User>(null);
    private tokenExpirationTimer: any;
    constructor(private http: HttpClient, private router: Router){}

    signUp(email: string, password: string){
        return this.http.post<AuthResponseData>('https://identitytoolkit.googleapis.com/v1/accounts:signUp?key=AIzaSyBdJS9BFKKAvJQi0fl4vhOPo24yNErGvR8',
        {
            email: email,
            password: password,
            returnSecureToken: true
        }).pipe(tap(resData => {
           this.handleAuth(resData.email, resData.localId, resData.idToken, +resData.expiresIn)
        }))
    }

    login(email: string, password: string){
        return this.http.post<any>('https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=AIzaSyBdJS9BFKKAvJQi0fl4vhOPo24yNErGvR8', {
            email: email,
            password: password,
            returnSecureToken: true 
        }).pipe(tap(resData => {
            this.handleAuth(resData.email, resData.localId, resData.idToken, +resData.expiresIn)
        }));
    }

    autoLogin(){
        const userData: {
            email: string;
            id: string;
            _token: string;
            _tokenExpirationDate: string;
        } = JSON.parse(localStorage.getItem('userData'));
        if (!userData){
            return;
        }
        const loadedUser = new User(userData.email, userData.id, userData._token, new Date(userData._tokenExpirationDate));
        if(loadedUser.token){
            this.user.next(loadedUser);
            const expirationDuration = new Date(userData._tokenExpirationDate).getTime() - new Date().getTime();
            this.autoLogout(expirationDuration);
        }
    }

    logout(){
        this.user.next(null);
        this.router.navigate(['/login']);
        localStorage.removeItem('userData');
        localStorage.removeItem('userType');
        if (this.tokenExpirationTimer){
            clearTimeout(this.tokenExpirationTimer);
        }
        this.tokenExpirationTimer = null;
    }

    autoLogout(expirationDuration: number){
        this.tokenExpirationTimer = setTimeout(() => {
            this.logout();
        }, expirationDuration)
    }

    private handleAuth(email: string, userId: string, token: string, expiresIn: number){
        const expirationDate = new Date(new Date().getTime() + expiresIn * 1000);
        const user = new User(email, userId, token, expirationDate);
        this.user.next(user);
         this.autoLogout(expiresIn * 1000);
        localStorage.setItem('userData', JSON.stringify(user));
    }
}