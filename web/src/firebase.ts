import {initializeApp} from 'firebase/app'

export function loadFirebaseConfig() {
    // from https://firebase.google.com/docs/hosting/reserved-urls?authuser=2
    // in dev env, web/public/__/firebase/init.json will be loaded
    return fetch('/__/firebase/init.json').then(async response => {
        initializeApp(await response.json())
    })
}