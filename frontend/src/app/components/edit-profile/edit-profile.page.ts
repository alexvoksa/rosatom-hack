import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, NgForm, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ManageProfileService } from './manage-profile-data.service';
import { Organisation } from './organisation.interface';

@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.page.html',
  styleUrls: ['./edit-profile.page.scss'],
})
export class EditProfilePage implements OnInit {

  profileForm: FormGroup;
  profile: Organisation;
  constructor(public formBuilder: FormBuilder, public profileSerice: ManageProfileService, private router: Router) { }

  ngOnInit() {
    this.createForm();
    
  }

  onSubmit(form: NgForm){
    if(!form.valid){
      return;
    }
    const id = JSON.parse(localStorage.getItem('userData')).id;
    this.profile = form.value;
    console.log(this.profile);
    this.profileSerice.postProfileData(this.profile, id).subscribe(response => {
      console.log(response);
      this.router.navigate(['/profile'])
    })

  }

  createForm(){
    this.profileForm = this.formBuilder.group({
      orgName: ['', Validators.required],
      ogrn: ['', Validators.required],
      ULinn: ['', Validators.required],
      address: ['', Validators.required],
      name: ['', Validators.required],
      FLinn: ['', Validators.required],
      passportSeries: ['', Validators.required],
      passportNumber: ['', Validators.required],
      snils: ['', Validators.required],
      okved: ['', Validators.required],
     
    });
    // this.profileForm.valueChanges.subscribe((data) => {
    //   console.log(data);
    // });
  }

}
