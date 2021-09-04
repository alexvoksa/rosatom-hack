import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from './auth.service';

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.scss'],
})
export class AuthComponent implements OnInit {

  isLoginMode = true;
  constructor(private authService: AuthService, private router: Router) { }

  onSwitchMode(){
    this.isLoginMode = !this.isLoginMode;
  }

  onSubmit(form: NgForm){
    if (!form.valid){
      return;
    }
    const email = form.value.email;
    const password = form.value.password;

    if(this.isLoginMode){
        this.authService.login(email, password).subscribe(response => {
          console.log(response);
          this.router.navigate(['/profile']);
        }, error => {
          console.log(error);
        })
    } else {
        this.authService.signUp(email, password).subscribe(response => {
          console.log(response);
          this.router.navigate(['/edit-profile']);
        }, error => {
          console.log(error);
        });
    }
    form.reset();
  }

  ngOnInit() {}

}