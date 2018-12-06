import { Component, OnInit, Input } from '@angular/core';
import { UserService } from '../user.service';
import { User } from '../model/user';

import { UserListComponent } from '../user-list/user-list.component';

@Component({
  selector: 'user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['./user-details.component.css']
})
export class UserDetailsComponent implements OnInit {

  @Input() user: User;

  constructor(private userService: UserService, private listComponent: UserListComponent) { }

  ngOnInit() {
  }

  /*updateActive(isActive: boolean) {
    this.personService.updatePerson(this.person.id,
      { name: this.person.name, lastname: this.person.lastname })
      .subscribe(
        data => {
          console.log(data);
          this.person = data as Person;
        },
        error => console.log(error));
  }*/

  deleteUser() {
    this.userService.deleteUser(this.user.ID)
      .subscribe(
        data => {
          console.log(data);
          this.listComponent.reloadData();
        },
        error => console.log(error));
  }

}
