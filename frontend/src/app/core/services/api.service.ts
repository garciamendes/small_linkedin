import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http";
import { inject, Injectable } from "@angular/core";
import { catchError, Observable, throwError } from "rxjs";
import { AuthService } from "./auth.service";

@Injectable({
  providedIn: "root",
})
export class ApiService {
  private baseUrl = "http://localhost:8080";
  private http = inject(HttpClient);

  get<T>(
    endpoint: string,
    params?: Record<string, string>,
    headers?: Record<string, string>
  ): Observable<T> {
    return this.http
      .get<T>(this.buildUrl(endpoint), {
        headers: this.buildHeaders(headers),
        params: this.buildParams(params),
      })
      .pipe(catchError(this.handleError));
  }

  post<T, D = any>(
    endpoint: string,
    data: D,
    headers?: Record<string, string>
  ): Observable<T> {
    return this.http
      .post<T>(this.buildUrl(endpoint), data, {
        headers: this.buildHeaders(headers),
      })
      .pipe(catchError(this.handleError));
  }

  put<T, D = any>(
    endpoint: string,
    data: D,
    headers?: Record<string, string>
  ): Observable<T> {
    return this.http
      .put<T>(this.buildUrl(endpoint), data, {
        headers: this.buildHeaders(headers),
      })
      .pipe(catchError(this.handleError));
  }

  delete<T>(endpoint: string, headers?: Record<string, string>): Observable<T> {
    return this.http
      .delete<T>(this.buildUrl(endpoint), {
        headers: this.buildHeaders(headers),
      })
      .pipe(catchError(this.handleError));
  }

  private buildUrl(endpoint: string): string {
    return `${this.baseUrl}${endpoint}`;
  }

  private buildHeaders(customheaders?: Record<string, string>): HttpHeaders {
    let headers = new HttpHeaders({
      "Content-Type": "application/json",
      ...customheaders,
    });

    return headers;
  }

  private buildParams(params?: Record<string, string>): HttpParams {
    let httpParams = new HttpParams();

    if (params) {
      Object.entries(params).forEach(([key, value]) => {
        if (value !== undefined && value !== null) {
          httpParams = httpParams.set(key, String(value));
        }
      });
    }

    return httpParams;
  }

  private handleError(error: any) {
    return throwError(() => error);
  }
}
