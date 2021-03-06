import { HttpInterceptor, HttpRequest, HttpHandler, HttpEvent } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment";
import { AuthenticationService } from "../service/authentication.service";

@Injectable({
    providedIn: 'root'
})

export class TokenInterceptor implements HttpInterceptor {
    constructor(public authService: AuthenticationService) { }

    intercept(
        request: HttpRequest<any>,
        next: HttpHandler
      ): Observable<HttpEvent<any>> {
        // add auth header with jwt if user is logged in and request is to api url
        const accessToken = this.authService.getToken();
        const isLoggedIn = this.authService.isLoggedIn();
        const isApiUrl = request.url.startsWith(environment.apiUrl);
        if (isLoggedIn && isApiUrl) {
          request = request.clone({
            setHeaders: {
              'Accept': 'application/json',
              'Access-Control-Allow-Origin': ['http://localhost:4200', 'http://localhost:4201'],
              'Authorization': `${accessToken}`,
              'Content-Type': 'application/json'
            },
          });
        }
    
        return next.handle(request);
      }
}