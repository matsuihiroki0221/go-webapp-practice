import { Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent }, // /login で表示
  { path: '', redirectTo: '/login', pathMatch: 'full' }, // デフォルトで /login にリダイレクト
];

export default routes;
