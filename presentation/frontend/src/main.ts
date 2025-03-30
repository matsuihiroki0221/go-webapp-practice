import { bootstrapApplication } from '@angular/platform-browser';
import { createAppProvidersConfig, getAuthConfig } from './app/app.config';
import { AppComponent } from './app/app.component';

console.log('main.ts');
getAuthConfig()
  .then((config) => {
    const providers = createAppProvidersConfig(config);
    bootstrapApplication(AppComponent, providers).catch((err) =>
      console.error(err),
    );
  })
  .catch((err) => console.error(err));
