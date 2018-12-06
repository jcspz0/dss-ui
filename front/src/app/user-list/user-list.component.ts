import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';

import { UserService } from '../user.service';
import { User } from '../model/user';

@Component({
  selector: 'user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css']
})
export class UserListComponent implements OnInit {


  users: Observable<User[]>;

  user: Array<User>;

  constructor(private userService: UserService) { }

  ngOnInit() {
    this.reloadData();
  }

  deleteUser() {
    this.userService.deleteAll()
      .subscribe(
        data => {
          console.log(data);
          this.reloadData();
        },
        error => console.log('ERROR: ' + error));
  }

  reloadData() {
    this.userService.getUserList()
    .subscribe(user => {
      console.log(this.user = user);
      return this.user = user;
    });

    console.log(this.user);
  }

}
