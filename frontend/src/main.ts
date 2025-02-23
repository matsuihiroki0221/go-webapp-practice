import { createApp } from 'vue'
import './style.css'
import { createAuth0 } from '@auth0/auth0-vue';

import App from './App.vue'

const app = createApp(App);

app.use(
  createAuth0({
    domain: import.meta.env.VITE_AUTH0_DOMAIN,
    clientId: import.meta.env.VITE_WEBAPP_CLIENT_ID,
    authorizationParams: {
      redirect_uri: window.location.origin,
      audience: import.meta.env.VITE_AUTH0_API_AUDIENCE,
    }
  })
);


app.mount('#app');