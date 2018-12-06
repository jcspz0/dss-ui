import { Component, OnInit, Input } from '@angular/core';
import { DocService } from '../doc.service';
import { Doc } from '../model/doc';

import { DocListComponent } from '../doc-list/doc-list.component';

@Component({
  selector: 'document-details',
  templateUrl: './doc-details.component.html',
  styleUrls: ['./doc-details.component.css']
})
export class DocDetailsComponent implements OnInit {

  @Input() doc: Doc;

  constructor(private docService: DocService, private listComponent: DocListComponent) { }

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

  deleteDoc() {
    this.docService.deleteDoc(this.doc.ID)
      .subscribe(
        data => {
          console.log(data);
          this.listComponent.reloadData();
        },
        error => console.log(error));
  }

}
