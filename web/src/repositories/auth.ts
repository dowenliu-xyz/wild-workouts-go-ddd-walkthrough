import {connectAuthEmulator, getAuth, onAuthStateChanged, signInWithEmailAndPassword} from 'firebase/auth'
import {getTestUsers, type User} from '@/repositories/user'
import {SignJWT} from 'jose'

interface Auth {
  login(login: string, password: string): Promise<void>;

  waitForAuthReady(): Promise<void>;

  getJwtToken(required: boolean): Promise<string>

  logout(): Promise<void>

  isLoggedIn(): boolean
}

class FirebaseAuth implements Auth {
  login(login: string, password: string): Promise<void> {
    return new Promise((resolve) => {
      signInWithEmailAndPassword(this.getAuth(), login, password).then(() => {
        resolve()
      })
    })
  }

  waitForAuthReady(): Promise<void> {
    return new Promise((resolve) => {
      return onAuthStateChanged(this.getAuth(), () => {
        resolve()
      })
    })
  }

  getJwtToken(required: boolean): Promise<string> {
    return new Promise((resolve, reject) => {
      const auth = this.getAuth()
      if (!auth.currentUser) {
        if (required) {
          reject('no user found')
        } else {
          resolve('')
        }
        return
      }

      auth.currentUser.getIdToken(false)
        .then((token) => resolve(token))
        .catch((error) => reject(error))
    })
  }

  logout(): Promise<void> {
    return new Promise((resolve) => {
      const auth = this.getAuth()
      if (!auth.currentUser) {
        resolve()
        return
      }

      return auth.signOut()
    })
  }

  isLoggedIn(): boolean {
    return this.getAuth().currentUser != null
  }

  private getAuth() {
    const auth = getAuth()
    if (import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST) {
      connectAuthEmulator(auth, import.meta.env.VITE_FIREBASE_AUTH_EMULATOR_HOST)
    }
    return auth
  }
}

class MockAuth implements Auth {
  login(login: string, password: string): Promise<void> {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        const found = getTestUsers().filter(u => u.login === login && u.password === password)

        if (found) {
          localStorage.setItem('_mock_user', JSON.stringify(found[0]))
          resolve()
        } else {
          reject('invalid login or password')
        }
      }, 500)
    })
  }

  waitForAuthReady(): Promise<void> {
    return new Promise((resolve) => {
      setTimeout(resolve, 50)
    })
  }

  getJwtToken(required: boolean): Promise<string> {
    return new Promise((resolve, reject) => {
      const user = this.currentMockUser()

      if (!user) {
        if (required) {
          reject('no user found')
        } else {
          resolve('')
        }
        return
      }

      const signJWT = new SignJWT({
        'user_uuid': user.uuid,
        'email': user.login,
        'role': user.role,
        'name': user.name,
      }).setProtectedHeader({alg: 'HS256'})
      const secret = new TextEncoder().encode('mock_secret')
      signJWT.sign(secret).then((token) => resolve(token))
    })
  }

  private currentMockUser: () => User | null = () => {
    const userStr = localStorage.getItem('_mock_user')
    if (!userStr) {
      return null
    }

    let user: User = {
      uuid: '',
      login: '',
      password: '',
      role: '',
      name: '',
    }
    try {
      user = Object.assign(user, JSON.parse(userStr))
    } catch (e) {
      console.error('invalid _mock_user', userStr, e)
      return null
    }
    return user
  }

  logout(): Promise<void> {
    return new Promise((resolve) => {
      localStorage.removeItem('_mock_user')
      setTimeout(resolve, 50)
    })
  }

  isLoggedIn(): boolean {
    return this.currentMockUser() !== null
  }
}

export let Auth: Auth

if (import.meta.env.DEV) {
  Auth = new MockAuth()
} else {
  Auth = new FirebaseAuth()
}

export const setApiClientsAuth = (token: string): void => {
  localStorage.setItem('token', token)
}
