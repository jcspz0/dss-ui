import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';

import { DocService } from '../doc.service';
import { Doc } from '../model/doc';


@Component({
  selector: 'doc-list',
  templateUrl: './doc-list.component.html',
  styleUrls: ['./doc-list.component.css']
})
export class DocListComponent implements OnInit {

  docs: Observable<Doc[]>;

  constructor(private docService: DocService) { }

  ngOnInit() {
    this.reloadData();
  }

  deleteDoc() {
    this.docService.deleteAll()
      .subscribe(
        data => {
          console.log(data);
          this.reloadData();
        },
        error => console.log('ERROR: ' + error));
  }

  reloadData() {
    this.docs = this.docService.getDocList();
    console.log(this.docs);
  }

  uploadFile(files: FileList) {

    this.docService.uploadFile(files).subscribe(filename => console.log(files[0].name));
    this.reloadData();
  }


}


