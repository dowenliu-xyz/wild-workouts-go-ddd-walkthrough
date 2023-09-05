import {Configuration as UserApiCfg, DefaultApi as UsersApi} from '@/repositories/clients/users'
import {Auth, setApiClientsAuth} from './auth'
import {serverSettings, tokenProvider} from './config'

export const Trainer = 'trainer'
export const Attendee = 'attendee'

const getBasePath = (): string => {
  return import.meta.env.VITE_USERS_API_BASE_PATH ?? (serverSettings.hostname + '/api')
}

const usersApi = new UsersApi(new UserApiCfg({
  basePath: getBasePath(),
  accessToken: tokenProvider,
}))

export const getUserRole = (): string => {
  return localStorage.getItem('role') ?? ''
}

export const getTrainingBalance = async (): Promise<number> => {
  return usersApi.getCurrentUser()
    .then((response) => response.data.balance)
}

export const loginUser = (login: string, password: string): Promise<void> => {
  return Auth.login(login, password)
    .then(() => {
      return Auth.waitForAuthReady()
    })
    .then(() => {
      return Auth.getJwtToken(false)
    })
    .then((token) => {
      setApiClientsAuth(token)
    })
    .then(() => {
      return usersApi.getCurrentUser().then((response) => response.data)
    })
    .then((data) => {
      localStorage.setItem('role', data.role)
    })
}

export interface User {
  uuid: string;
  login: string;
  password: string;
  role: string;
  name: string;
}

export const getTestUsers = (): User[] => {
  return [
    {
      uuid: '1',
      login: 'trainer@threedots.tech',
      password: '123456',
      role: 'trainer',
      name: 'Trainer',
    },
    {
      uuid: '2',
      login: 'attendee@threedots.tech',
      password: '123456',
      role: 'attendee',
      name: 'Mock Arnie',
    },
  ]
}