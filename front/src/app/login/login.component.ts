import { Component, OnInit } from '@angular/core';

import { User } from '../model/user';
import { UserService } from '../user.service';
import {Globals} from '../globals';

@Component({
  selector: 'login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  user: User = new User();
  submitted = false;
  errorU = false;
  private logged: string;

  constructor(private userService: UserService, private globals: Globals) {
    this.logged = globals.login;
  }

  ngOnInit() {
  }

  updateLogin(): void {
    this.submitted = false;
    this.errorU = false;
  }

  save() {
    this.userService.getUserByName(this.logged)
      .subscribe(data => {
        console.log(data);
        this.submitted = true;
        this.globals.login = this.logged;
      }, error => {
        console.log(error);
        this.submitted = false;
        this.logged = 'admin';
        this.globals.login = this.logged;
        this.errorU = true;
      });
    this.user = new User();
  }

  onSubmit() {
    this.submitted = true;
    this.save();
  }

}
