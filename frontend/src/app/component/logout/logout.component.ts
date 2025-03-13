import { Component, Inject } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { DOCUMENT } from '@angular/common';

@Component({
  selector: 'app-auth-button',
  template: ` <button (click)="auth.loginWithRedirect()">Log in</button> `,
  standalone: true,
})
export class LogoutComponent {
  constructor(
    @Inject(DOCUMENT) public document: Document,
    public auth: AuthService,
  ) {}
}
