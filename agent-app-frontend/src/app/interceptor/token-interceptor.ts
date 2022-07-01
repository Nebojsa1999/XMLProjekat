import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";

import { AuthService } from "../services/auth/auth.service";

@Injectable({
    providedIn: 'root'
})
export class TokenInterceptor implements HttpInterceptor {
    constructor(public authService: AuthService) { }

    intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        if (this.authService.getAuthStatus()) {
            req = req.clone({
                setHeaders: {
                    'Accept': 'application/json',
                    'Access-Control-Allow-Origin': 'http://localhost:4200',
                    'Authorization': `Bearer ${this.authService.getAgentAppToken()}`,
                    'Content-Type': 'application/json'
                }
            });
        }

        return next.handle(req);
    }
}
