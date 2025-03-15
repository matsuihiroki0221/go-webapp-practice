import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';

// Import the authentication guard
// import { AuthGuard } from '@auth0/auth0-angular';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: '**', redirectTo: '' },
];

export default routes;
