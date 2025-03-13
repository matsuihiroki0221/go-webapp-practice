import { Component } from '@angular/core';
import { MatCardModule } from '@angular/material/card';
import { AuthService } from '@auth0/auth0-angular';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-login',
  imports: [MatCardModule, MatButtonModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
})
export class LoginComponent {
  // Inject the authentication service into your component through the constructor
  constructor(public auth: AuthService) {}
}
