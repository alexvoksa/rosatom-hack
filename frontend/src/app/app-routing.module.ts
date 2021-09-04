import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';
import { AuthComponent } from './auth/auth.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'login',
    pathMatch: 'full'
  },
  {
    path: 'login',
    component: AuthComponent
  },
  {
    path: 'profile',
    loadChildren: () => import('./components/profile/profile.module').then( m => m.ProfilePageModule)
  },
  {
    path: 'purchases',
    loadChildren: () => import('./components/purchases/purchases.module').then( m => m.PurchasesPageModule)
  },
  {
    path: 'edit-profile',
    loadChildren: () => import('./components/edit-profile/edit-profile.module').then( m => m.EditProfilePageModule)
  }
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, { preloadingStrategy: PreloadAllModules })
  ],
  exports: [RouterModule]
})
export class AppRoutingModule {}
