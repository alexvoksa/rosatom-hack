import { Component, OnInit } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { filter } from 'rxjs/operators';
@Component({
  selector: 'app-root',
  templateUrl: 'app.component.html',
  styleUrls: ['app.component.scss'],
})
export class AppComponent implements OnInit {
  public appPages = [
    { title: 'Профиль', url: 'profile', icon: 'person' },
    { title: 'Закупки', url: 'purchases', icon: 'newspaper' },
    
  ];
  public labels = ['Андрей', 'Александр', 'Джордж', 'Димас', 'Андрей'];

  public showSideMenu = false;
  currentRoute: string;

  constructor(private router: Router){}

  ngOnInit(){
    this.router.events
      .pipe(filter(e => e instanceof NavigationEnd))
      .subscribe((e: NavigationEnd) => {
        this.currentRoute = e.url;
        console.log(this.currentRoute);
        if(this.currentRoute === '/login' || this.currentRoute === '/' || this.currentRoute === '/edit-profile'){
          this.showSideMenu = false;
        }
        else {
          this.showSideMenu = true;
        }   
    });
  }
}

