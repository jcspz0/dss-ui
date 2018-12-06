import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Doc } from './model/doc';
import { HttpClient , HttpHeaders} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class DocService {

  private ip = 'localhost';
  private port = '9001';

  private baseUrl = `http://${this.ip}:${this.port}/dss/api`;

  constructor(private http: HttpClient) { }

  getDoc(id: number): Observable<Object> {
    return this.http.get(`${this.baseUrl}/documents/${id}`);
  }

  createDoc(doc: Object): Observable<Object> {
    return this.http.post(`${this.baseUrl}` + `/documents`, doc);
  }

  updateDoc(id: number, value: any): Observable<Object> {
    return this.http.put(`${this.baseUrl}/${id}`, value);
  }

  deleteDoc(id: number): Observable<any> {
    return this.http.delete(`${this.baseUrl}` + `/documents` + `/${id}`, { responseType: 'text' });
  }

  getDocList(): Observable<any> {
    return this.http.get(`${this.baseUrl}` + `/documents`);
  }

  deleteAll(): Observable<any> {
    return this.http.delete(`${this.baseUrl}` + `/documents/delete`, { responseType: 'text' });
  }

  getDocByName(name: string): Observable<any> {
    return this.http.get(`${this.baseUrl}/documents/name/${name}`);
  }

  uploadFile(files: FileList) {
    const input = new FormData();

    input.append('file', files[0], files[0].name);
    const HttpUploadOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'multipart/form-data' })
    };
    console.log(input);
    return this.http.post(`${this.baseUrl}` + `/documents`, input);
  }

}
