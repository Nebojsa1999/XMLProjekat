import { HttpInterceptor, HttpRequest, HttpHandler, HttpEvent } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { AuthenticationService } from "../service/authentication.service";

@Injectable({
    providedIn: 'root'
})

export class TokenInterceptor implements HttpInterceptor {
    constructor(public authService: AuthenticationService) { }

    intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        if (this.authService.getAuthStatus()) {
            req = req.clone({
                setHeaders: {
                    'Accept': 'application/json',
                    'Access-Control-Allow-Origin': 'http://localhost:4200',
                    'Authorization': 'Bearer ${this.authService.getDislinktAppToken()}',
                    'Content-Type': 'application/json'
                }
            });
        }

        return next.handle(req);
    }
}