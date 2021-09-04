import { Component, OnInit } from '@angular/core';
import { ManageProfileService } from '../edit-profile/manage-profile-data.service';
import { Organisation } from '../edit-profile/organisation.interface';
import { AuthService } from '../../auth/auth.service'

@Component({
  selector: 'app-profile',
  templateUrl: './profile.page.html',
  styleUrls: ['./profile.page.scss'],
})
export class ProfilePage implements OnInit {

  constructor(public profileService: ManageProfileService, public authSerive: AuthService) { }
  profile: Organisation;
  ngOnInit() {
    const userId = JSON.parse(localStorage.getItem('userData')).id;
    this.profileService.getProfileData(userId).subscribe(response => {
      this.profile = response[0];
      console.log(this.profile);
    })
  }

  logout(){
    this.authSerive.logout();
  }

}