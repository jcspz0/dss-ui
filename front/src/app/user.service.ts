import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from './model/user';
import { HttpClient } from '@angular/common/http';
import 'rxjs/add/operator/map';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private ip = 'localhost';
  private port = '9001';

  private baseUrl = `http://${this.ip}:${this.port}/dss/api`;

  constructor(private http: HttpClient) { }

  getUser(id: number): Observable<Object> {
    return this.http.get(`${this.baseUrl}/users/${id}`);
  }

  createUser(user: User): Observable<Object> {
    return this.http.post(`${this.baseUrl}` + `/users`, user);
  }

  updateUser(id: number, value: any): Observable<Object> {
    return this.http.put(`${this.baseUrl}/${id}`, value);
  }

  deleteUser(id: number): Observable<any> {
    return this.http.delete(`${this.baseUrl}` + `/users` + `/${id}`, { responseType: 'text' });
  }

  getUserList(): Observable<User[]> {
    return this.http.get<User[]>(`${this.baseUrl}` + `/users`).map(data => {
      console.log(<User[]>data);
      return <User[]>data;
    });
 }

  deleteAll(): Observable<any> {
    return this.http.delete(`${this.baseUrl}` + `/users/delete`, { responseType: 'text' });
  }

  getUserByEmail(email: string): Observable<any> {
    return this.http.get(`${this.baseUrl}/users/email/${email}`);
  }

  getUserByName(name: string): Observable<any> {
    return this.http.get(`${this.baseUrl}/users/${name}`);
  }

}
