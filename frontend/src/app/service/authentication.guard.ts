import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { AuthenticationService } from './authentication.service';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationGuard implements CanActivate {
  constructor(private authService: AuthenticationService, private router: Router) { }

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean | Promise<boolean> {
    var isAuthenticated = this.authService.getAuthStatus();
    var hasDislinktAppTokenExpired = this.authService.hasDislinktAppTokenExpired();

    if (!isAuthenticated || hasDislinktAppTokenExpired) {
      this.router.navigate(['/login']);

      return false;
    }

    /*
    if (route.url.toString()) {
      this.router.navigate(['/login']);
      return false;
    }
    */

    return true;
  }
}