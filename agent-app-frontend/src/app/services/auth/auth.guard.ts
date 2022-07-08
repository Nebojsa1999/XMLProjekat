import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot } from '@angular/router';

import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(private authService: AuthService, private router: Router) { }

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean | Promise<boolean> {
    var isAuthenticated = this.authService.getAuthStatus();
    var hasAgentAppTokenExpired = this.authService.hasAgentAppTokenExpired();

    if (!isAuthenticated || hasAgentAppTokenExpired) {
      this.router.navigate(['/login']);

      return false;
    }

    return true;
  }
}
