import { Injectable } from "@angular/core";
import { ApiService } from "./api.service";
import { Observable, tap } from "rxjs";
import { toast } from "ngx-sonner";

export interface ILoginData {
  email: string;
  password: string;
}

@Injectable({
  providedIn: "root",
})
export class AuthService {
  private KeyAuth = "token";

  constructor(private api: ApiService) {}

  login(data: ILoginData): Observable<any> {
    return this.api
      .post<{ token: string }, ILoginData>("/api/auth/login", data)
      .pipe(
        tap((response) => {
          sessionStorage.setItem(this.KeyAuth, response.token);
        })
      );
  }

  logout() {
    sessionStorage.removeItem(this.KeyAuth);
  }

  getToken() {
    return sessionStorage.getItem(this.KeyAuth);
  }

  getKeyAuth() {
    return this.KeyAuth;
  }
}
