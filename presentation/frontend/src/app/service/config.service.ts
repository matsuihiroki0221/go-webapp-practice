import { HttpClient } from '@angular/common/http';
import { Injectable, InjectionToken } from '@angular/core';
import { firstValueFrom } from 'rxjs';

export const CONFIG_DATA = new InjectionToken<AuthConfig>('CONFIG_DATA');

@Injectable({
  providedIn: 'root',
})
export class ConfigService {
  private config: AuthConfig | null = null;
  constructor(private http: HttpClient) {}

  async loadConfig() {
    this.config = await firstValueFrom(
      this.http.get<AuthConfig>('http://localhost:8080/api/config'),
      // this.http.get<AuthConfig>('/api/config'),
    );
  }

  getConfig() {
    return this.config;
  }
}

interface AuthConfig {
  domain: string;
  clientId: string;
  audience: string;
}
