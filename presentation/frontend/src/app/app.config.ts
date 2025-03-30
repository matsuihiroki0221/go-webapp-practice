import { ApplicationConfig, provideZoneChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideAuth0 } from '@auth0/auth0-angular';
import routes from './app.routes';

// export const appConfig: ApplicationConfig = {
//   providers: [
//     provideZoneChangeDetection({ eventCoalescing: true }),
//     provideRouter(routes),
//     provideAuth0({
//       domain: 'dev-23fbe27gcpbfmo8v.us.auth0.com',
//       clientId: '0xFiPYMay6RWK9bB06RpgfeuXavOnp58',
//       authorizationParams: {
//         redirect_uri: window.location.origin,
//         audience: 'practice-backend',
//       },
//     }),
//   ],
// };

export function createAppProvidersConfig(
  config: AuthConfig,
): ApplicationConfig {
  return {
    providers: [
      provideZoneChangeDetection({ eventCoalescing: true }),
      provideRouter(routes),
      provideAuth0({
        domain: config.domain,
        clientId: config.clientId,
        authorizationParams: {
          redirect_uri: window.location.origin,
          audience: config.audience,
        },
      }),
    ],
  };
}

export async function getAuthConfig(): Promise<AuthConfig> {
  const response = await fetch('http://localhost:8080/api/config');
  const data = await response.json();

  // 型チェックを行う（念のため）
  if (!data.domain || !data.clientId || !data.audience) {
    throw new Error('Invalid AuthConfig response');
  }

  return data as AuthConfig;
}

interface AuthConfig {
  domain: string;
  clientId: string;
  audience: string;
}
