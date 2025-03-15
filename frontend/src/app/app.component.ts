import { Component, Inject } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { RouterOutlet } from '@angular/router';
import { DOCUMENT, CommonModule } from '@angular/common';

import { LoginComponent } from './component/login/login.component';
import { LoadingComponent } from './component/loading/loading.component';

import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSidenavModule } from '@angular/material/sidenav';

@Component({
  selector: 'app-root',
  imports: [
    RouterOutlet,
    LoginComponent,
    CommonModule,
    MatButtonModule,
    MatIconModule,
    MatToolbarModule,
    MatSidenavModule,
    LoadingComponent,
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent {
  constructor(
    @Inject(DOCUMENT) public document: Document,
    public auth: AuthService,
  ) {}

  title = 'webAppPractice';
  opened = true;
}
