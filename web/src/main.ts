import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import './assets/main.css'

import {createApp} from 'vue'
import {createPinia} from 'pinia'

import App from './App.vue'
import router from './router'
import {loadFirebaseConfig} from '@/firebase'
import {Auth, setApiClientsAuth} from '@/repositories/auth'

const app = createApp(App)

app.use(createPinia())
app.use(router)

loadFirebaseConfig()
  .then(() => {
    return Auth.waitForAuthReady()
  })
  .then(() => {
    return Auth.getJwtToken(false)
  })
  .then((token) => {
    setApiClientsAuth(token)
  })
  .finally(() => {
    app.mount('#app')
  })
