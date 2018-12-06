export class User {
  ID: number;
  Name: string;
  Email: string;

fromJSON(obj: any) {
 this.ID = obj.ID;
 this.Name = obj.Name;
 this.Email = obj.Email;
 return this;
}

}
